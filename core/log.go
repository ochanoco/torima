package core

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"runtime"
)

type FlowLog struct {
	name   string
	result int
}

type FlowLogger struct {
	logs []FlowLog
}

func NewFlowLogger() FlowLogger {
	return FlowLogger{
		logs: []FlowLog{},
	}
}

func (logger *FlowLogger) Add(f any, result int) {
	rv := reflect.ValueOf(f)
	ptr := rv.Pointer()
	name := runtime.FuncForPC(ptr).Name()

	newLog := FlowLog{
		name,
		result,
	}

	logger.logs = append(logger.logs, newLog)
}

func (flowLogs *FlowLogger) Show() {
	log.Println("\n--- start ----")

	for _, v := range flowLogs.logs {
		log.Printf("name: %v\n", v.name)
		log.Printf("result: %v\n", v.result)
	}

	fmt.Println("---  end  ----")
}

/**
 * LogReq is the function that logs the request.
**/
func LogReq(req *http.Request) {
	fmt.Printf("[%s] %s%s\n=> %s%s\n\n", req.Method, req.Host, req.RequestURI, req.URL.Host, req.URL.Path)
}
