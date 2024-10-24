package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Service struct {
	Port    string `json:"port"`
	State   string `json:"state"`
	Service string `json:"service"`
	Version string `json:"version"`
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

	http.HandleFunc("/services", func(w http.ResponseWriter, r *http.Request) {
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
				State:   "open",
				Service: "mysql",
				Version: "MySQL 5.5.64-MariaDB-1~trusty",
			},
			{
				Port:    "8080/tcp",
				State:   "open",
				Service: "http",
				Version: "Apache httpd 2.4.57 ((Debian))",
			},
			{
				Port:    "8081/tcp",
				State:   "open",
				Service: "http",
				Version: "Apache httpd 2.4.62 ((Debian))",
			},
		}
	}

	logs := coll.ExportedLogs()
	if logs == nil {
		return nil
	}

	var res []Service
	for _, rl := range logs {
		for _, sl := range rl.ScopeLogs {
			for _, lr := range sl.LogRecords {
				nmapEvent := false
				for _, a := range lr.Attributes {
					if a.Key == "event.name" && a.Value.GetStringValue() == "nmap.run" {
						nmapEvent = true
						break
					}
				}
				if !nmapEvent {
					break
				}
				body := lr.Body
				for _, bv := range body.GetArrayValue().Values {
					var port, protocol, state, serviceName, product, version string
					for _, f := range bv.GetKvlistValue().Values {
						switch f.Key {
						case "port_id":
							port = strconv.Itoa(int(f.Value.GetIntValue()))
						case "protocol":
							protocol = f.Value.GetStringValue()
						case "state":
							state = f.Value.GetStringValue()
						case "service":
							for _, sf := range f.Value.GetKvlistValue().Values {
								switch sf.Key {
								case "name":
									serviceName = sf.Value.GetStringValue()
								case "product":
									product = sf.Value.GetStringValue()
								case "version":
									version = sf.Value.GetStringValue()
								}
							}
						}
					}
					srvc := Service{
						Port:    port + "/" + protocol,
						State:   state,
						Service: serviceName,
						Version: product + " " + version,
					}
					res = append(res, srvc)
				}
			}
		}
	}
	return res
}
