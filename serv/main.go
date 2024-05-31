package main

import (
	"fmt"

	"github.com/ochanoco/torima/core"
	"github.com/ochanoco/torima/proxy"
)

const NAME = "line"

func Run() (*core.TorimaProxy, error) {
	proxyServ, err := proxy.ProxyServer()
	return proxyServ, err
}

func main() {
	proxyServ, err := Run()
	if err != nil {
		panic(err)
	}

	port := fmt.Sprintf(":%d", proxyServ.Config.Port)
	proxyServ.Engine.Run(port)
}
