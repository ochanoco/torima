package utils

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"runtime"
)

type FlowLog struct {
	name   string
	result string
}

type FlowLogger struct {
	logs []FlowLog
}

func NewFlowLogger() FlowLogger {
	return FlowLogger{
		logs: []FlowLog{},
	}
}

func (logger *FlowLogger) Add(extension any, result int) {
	rv1 := reflect.ValueOf(extension)
	ptr1 := rv1.Pointer()

	rv2 := reflect.ValueOf(result)
	ptr2 := rv2.Pointer()

	extensionName := runtime.FuncForPC(ptr1).Name()
	resultName := runtime.FuncForPC(ptr2).Name()

	newLog := FlowLog{
		extensionName,
		resultName,
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
