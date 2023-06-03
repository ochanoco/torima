package domains

import (
	"log"
	"net"

	"github.com/miekg/dns"
)

var (
	records map[string]net.IP
)

func Lookup(w dns.ResponseWriter, req *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(req)
	m.Authoritative = true

	if len(req.Question) != 1 {
		return
	}
	q := req.Question[0]
	log.Printf("Lookup: question: %v\n", q)

	w.WriteMsg(m)
}

func StartServer() error {
	dns.HandleFunc("newt.example.", Lookup)

	srv := &dns.Server{Addr: ":8053", Net: "udp"}
	return srv.ListenAndServe()
}
