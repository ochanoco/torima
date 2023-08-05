package core

import (
	"time"

	"github.com/ochanoco/proxy/ent"
)

func logRawCommunication(_type string, header string, body []byte, proxy *OchanocoProxy) (*ent.CommunicationLog, error) {
	time := time.Now()

	l := proxy.Database.CommunicationLog(_type, time, header, body)
	return proxy.Database.SaveCommunicateLog(l)
}
