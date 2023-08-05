package core

import (
	"bytes"
	"time"

	"github.com/ochanoco/proxy/ent"
)

func logRawCommunication(_type string, message []byte, proxy *OchanocoProxy) (*ent.CommunicationLog, error) {
	time := time.Now()

	splited := bytes.Split(message, []byte("\r\n\r\n"))

	header := splited[0]
	headerLen := len(header)

	body := message[headerLen:]

	l := proxy.Database.CommunicationLog(_type, time, string(header), body)
	return proxy.Database.SaveCommunicateLog(l)
}
