package domains

import (
	"log"
	"net"

	"github.com/miekg/dns"
)

var (
	records map[string]net.IP
)

func init() {
	records = make(map[string]net.IP)
	records["service1.newt.example."] = net.IPv4(127, 0, 0, 1)
}

func Lookup(w dns.ResponseWriter, req *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(req)
	m.Authoritative = true

	if len(req.Question) != 1 {
		return
	}
	q := req.Question[0]
	log.Printf("Lookup: question: %v\n", q)

	ip, ok := records[q.Name]
	if !ok {
		log.Printf("Lookup: record %v not found\n", q.Name)
	}

	rr := &dns.A{
		Hdr: dns.RR_Header{Name: q.Name, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 0},
		A:   ip.To4(),
	}
	m.Answer = append(m.Answer, rr)

	w.WriteMsg(m)
}

func StartServer() error {
	dns.HandleFunc("newt.example.", Lookup)

	srv := &dns.Server{Addr: ":8053", Net: "udp"}
	return srv.ListenAndServe()
}
