package dns

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"switchboard-server/db"

	"github.com/miekg/dns"
)

var records = map[string]string{
	"test.service.": "192.168.0.2",
}

func parseQuery(m *dns.Msg) {
	for _, q := range m.Question {
		switch q.Qtype {
		case dns.TypeA:
			log.Printf("A Record query for %s\n", q.Name)
			ip := records[q.Name]
			if ip != "" {
				rr, err := dns.NewRR(fmt.Sprintf("%s A %s", q.Name, ip))
				if err == nil {
					m.Answer = append(m.Answer, rr)
				}
			}
		default:
			log.Printf("CNAME Record Query for %s\n", q.Name)
			lastDotRegExp := regexp.MustCompile(".$")
			record, _ := db.FindRecordByValue(lastDotRegExp.ReplaceAllString(q.Name, ""))
			if record.Host != "" {
				rr, err := dns.NewRR(fmt.Sprintf("%s CNAME %s", q.Name, record.Host))
				if err == nil {
					m.Answer = append(m.Answer, rr)
				}
			}
		}
	}
}

func handleDnsRequest(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	m.Compress = false

	switch r.Opcode {
	case dns.OpcodeQuery:
		parseQuery(m)
	}

	w.WriteMsg(m)
}

func getPort() string {
	port := os.Getenv("DNS_PORT")
	if port == "" {
		port = "53"
	}
	return port
}

func DnsServer() {
	dns.HandleFunc(".", handleDnsRequest)
	port := getPort()

	dnsServer := dns.Server{
		Addr: ":" + port,
		Net:  "udp",
	}

	log.Printf("Starting DNS server at %s port\n", port)
	err := dnsServer.ListenAndServe()
	defer dnsServer.Shutdown()
	if err != nil {
		log.Fatalf("Failed to start server: %s\n ", err.Error())
	}
}
