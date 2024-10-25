package nmapreceiver

import (
	"fmt"
	"strings"

	"github.com/Ullaakut/nmap/v3"
	"go.uber.org/zap"
)

// Scanner is a wrapper around the [nmap.Scanner].
type Scanner struct {
	scanner *nmap.Scanner
	logger  *zap.Logger
}

// NewScanner creates a new instance of the [Scanner].
func NewScanner(scanner *nmap.Scanner, logger *zap.Logger) *Scanner {
	return &Scanner{
		scanner: scanner,
		logger:  logger,
	}
}

// Run executes the nmap scan and returns the result.
func (s *Scanner) Run() (Scan, error) {
	res, warnings, err := s.scanner.Run()
	if len(*warnings) > 0 {
		s.logger.Warn("nmap scan returned warnings", zap.Strings("warnings", *warnings))
	}
	if err != nil {
		return Scan{}, fmt.Errorf("unable to run nmap scan: %w", err)
	}

	var scan Scan
	for _, host := range res.Hosts {
		for _, port := range host.Ports {
			service := Service{
				Name:    port.Service.Name,
				Product: port.Service.Product,
				Version: port.Service.Version,
			}

			for _, script := range port.Scripts {
				service.Vulnerabilities = append(service.Vulnerabilities, getVulnerabilities(s.logger, script)...)
			}

			scan.Ports = append(scan.Ports, Port{
				ID:       port.ID,
				Protocol: port.Protocol,
				State:    port.State.State,
				Service:  service,
			})
		}
	}

	return scan, nil
}

func getVulnerabilities(logger *zap.Logger, script nmap.Script) []Vulnerability {
	if script.ID != "vulners" || len(script.Tables) == 0 {
		return nil
	}

	var vulns []Vulnerability
	for _, t := range script.Tables[0].Tables {
		var (
			id, vType, isExploit string
		)
		v := Vulnerability{}
		for _, elem := range t.Elements {
			switch elem.Key {
			case "id":
				id = elem.Value
			case "cvss":
				val := elem.Value
				if val != "0.0" {
					v.CVSS = val
				}
			case "is_exploit":
				isExploit = elem.Value
			case "type":
				vType = elem.Value
			}
		}
		if strings.ToLower(isExploit) == "true" {
			v.ExploitID = id
		} else {
			v.CVE = id
		}
		v.URL = fmt.Sprintf("https://vulners.com/%s/%s", vType, id)
		vulns = append(vulns, v)
	}

	return vulns
}
