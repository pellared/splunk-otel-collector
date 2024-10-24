package nmapreceiver

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
)

// Scan represents the nmap scan result.
type Scan struct {
	Ports []Port
}

// Port represents a discovered port
// in the nmap scan result.
type Port struct {
	ID       uint16
	Protocol string
	State    string
	Service  Service
}

// Service represents the discovered service.
type Service struct {
	Name    string
	Product string
	Version string
}

// ToLogRecord converts the Scan to a [plog.LogRecord].
func (s Scan) ToLogRecord(dest plog.LogRecord) {
	dest.Attributes().PutStr("event.name", "nmap.run")
	body := dest.Body()
	ps := body.SetEmptySlice()
	for _, p := range s.Ports {
		dest := ps.AppendEmpty().SetEmptyMap()
		p.toMap(dest)
	}
}

func (p Port) toMap(dest pcommon.Map) {
	dest.PutInt("port_id", int64(p.ID))
	dest.PutStr("protocol", p.Protocol)
	dest.PutStr("state", p.State)
	serviceMap := dest.PutEmptyMap("service")
	p.Service.toMap(serviceMap)
}

func (s Service) toMap(dest pcommon.Map) {
	dest.PutStr("name", s.Name)
	dest.PutStr("product", s.Product)
	dest.PutStr("version", s.Version)
}