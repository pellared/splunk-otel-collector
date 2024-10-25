package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Service struct {
	Port            string          `json:"port"`
	State           string          `json:"state"`
	Host            string          `json:"host"`
	Service         string          `json:"service"`
	Version         string          `json:"version"`
	Vulnerabilities []Vulnerability `json:"vulnerabilities"`
}

type Vulnerability struct {
	CVE       string `json:"cve"`
	ExploitID string `json:"exploit_id"`
	URL       string `json:"url"`
	CVSS      string `json:"cvss"`
}

func main() {
	endpoint := os.Getenv("OTEL_RECEIVER_OTLP_ENDPOINT")
	if endpoint == "" {
		endpoint = ":4317"
	}
	coll := &collector{
		Endpoint: endpoint,
	}
	err := coll.Start()
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/services", func(w http.ResponseWriter, _ *http.Request) {
		srvs := services(coll)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(srvs); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	port := "8090"

	http.ListenAndServe(":"+port, nil)
}

func services(coll *collector) []Service {
	if os.Getenv("FAKE") != "" {
		return []Service{
			{
				Port:    "3306/tcp",
				Service: "mysql",
				Version: "MySQL 5.5.64-MariaDB-1~trusty",
			},
			{
				Port:    "8080/tcp",
				Service: "http",
				Version: "Apache httpd 2.4.57 ((Debian))",
			},
			{
				Port:    "8081/tcp",
				Service: "http",
				Version: "Apache httpd 2.4.62 ((Debian))",
			},
		}
	}

	logs := coll.ExportedLogs()

	var res []Service
	for _, rl := range logs {
		var host string
		for _, attr := range rl.GetResource().GetAttributes() {
			if attr.GetKey() == "host.name" {
				host = attr.GetValue().GetStringValue()
				break
			}
		}
		for _, sl := range rl.GetScopeLogs() {
			for _, lr := range sl.GetLogRecords() {
				nmapEvent := false
				for _, a := range lr.GetAttributes() {
					if a.Key == "event.name" && a.GetValue().GetStringValue() == "nmap.run" {
						nmapEvent = true
						break
					}
				}
				if !nmapEvent {
					break
				}

				for _, bv := range lr.GetBody().GetArrayValue().GetValues() {
					var (
						port, protocol, state, serviceName, product, version string
						vulns                                                []Vulnerability
					)
					for _, f := range bv.GetKvlistValue().GetValues() {
						switch f.Key {
						case "port":
							port = strconv.Itoa(int(f.GetValue().GetIntValue()))
						case "protocol":
							protocol = f.GetValue().GetStringValue()
						case "state":
							state = f.GetValue().GetStringValue()
						case "service":
							for _, sf := range f.GetValue().GetKvlistValue().GetValues() {
								switch sf.Key {
								case "name":
									serviceName = sf.GetValue().GetStringValue()
								case "product":
									product = sf.GetValue().GetStringValue()
								case "version":
									version = sf.GetValue().GetStringValue()
								case "vulnerabilities":
									for _, vf := range sf.GetValue().GetArrayValue().GetValues() {
										var cve, exploitID, url, cvss string
										for _, vf := range vf.GetKvlistValue().GetValues() {
											switch vf.GetKey() {
											case "cve":
												cve = vf.GetValue().GetStringValue()
											case "exploit_id":
												exploitID = vf.GetValue().GetStringValue()
											case "url":
												url = vf.GetValue().GetStringValue()
											case "cvss":
												cvss = vf.GetValue().GetStringValue()
											}
										}
										vulns = append(vulns, Vulnerability{
											CVE:       cve,
											ExploitID: exploitID,
											URL:       url,
											CVSS:      cvss,
										})
									}
								}
							}
						}
					}
					srvc := Service{
						Port:            port + "/" + protocol,
						Host:            host,
						State:           state,
						Service:         serviceName,
						Version:         product + " " + version,
						Vulnerabilities: vulns,
					}
					res = append(res, srvc)
				}
			}
		}
	}
	return res
}
