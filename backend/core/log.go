package core

import (
	"time"

	"github.com/ochanoco/proxy/ent"
)

func logRawCommunication(header string, body []byte, proxy *OchanocoProxy) (*ent.ServiceLog, error) {
	time := time.Now()

	l := proxy.Database.CreateServiceLog(time, header, body)
	return proxy.Database.SaveServiceLog(l)
}
