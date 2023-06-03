package main

import (
	"fmt"

	"github.com/miekg/dns"
)

func main() {
	fmt.Println("hello, world")
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
}
