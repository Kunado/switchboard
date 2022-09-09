package dns

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"switchboard/db"

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
		case dns.TypeCNAME:
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

func DnsServer() {
	dns.HandleFunc("service.", handleDnsRequest)
	port := 53530

	dnsServer := dns.Server{
		Addr: ":" + strconv.Itoa(port),
		Net:  "udp",
	}

	log.Printf("Starting DNS server at %d port\n", port)
	err := dnsServer.ListenAndServe()
	defer dnsServer.Shutdown()
	if err != nil {
		log.Fatalf("Failed to start server: %s\n ", err.Error())
	}
}
