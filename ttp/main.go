package main

import (
	"fmt"
	"log"

	"github.com/miekg/dns"
	"github.com/ochanoco/ttp/certs"
	"github.com/ochanoco/ttp/domains"
)

func main() {
	// skip attesttation
	verified, err := VerifyLogger("")
	if err != nil {
		panic(err)
	}
	if !verified {
		err = fmt.Errorf("main: failed to verify the logger")
		panic(err)
	}

	c := new(dns.Client)
	m := new(dns.Msg)
	m.SetQuestion("google.com.", dns.TypeA)
	m.RecursionDesired = true
	r, _, err := c.Exchange(m, "8.8.8.8:53")

	if err != nil {
		panic(err)
	}

	for _, a := range r.Answer {
		if a, ok := a.(*dns.A); ok {
			fmt.Printf("%s\n", a.String())
		}
	}

	msg := []byte("hello, world")
	signature, err := certs.Sign(msg)
	if err != nil {
		panic(err)
	}
	signaturePem := certs.EncodeSignaturePem(signature)
	log.Printf("main: signature:\n%v", string(signaturePem))

	ok := certs.Verify(msg, signature)
	log.Printf("main: is_valid_signature: %v", ok)

	err = domains.StartServer()
	if err != nil {
		e := fmt.Errorf("failed to serve: %w", err)

		panic(e)
	}
}
