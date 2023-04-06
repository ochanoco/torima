package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"runtime"
)

type FlowLog struct {
	name   string
	result bool
}

type FlowLogs struct {
	logs []FlowLog
}

func NewFlowLogs() FlowLogs {
	return FlowLogs{
		logs: []FlowLog{},
	}
}

func (flowLogs *FlowLogs) Add(f any, result bool) {
	rv := reflect.ValueOf(f)
	ptr := rv.Pointer()
	name := runtime.FuncForPC(ptr).Name()

	newLog := FlowLog{
		name,
		result,
	}

	flowLogs.logs = append(flowLogs.logs, newLog)
}

func (flowLogs *FlowLogs) Show() {
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
