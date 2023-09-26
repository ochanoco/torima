package serv

import (
	"fmt"

	"github.com/ochanoco/proxy/core"
)

const NAME = "line"

func Run() (*core.OchanocoProxy, error) {
	proxyServ, err := core.ProxyServer()
	return proxyServ, err
}

func Main() {
	proxyServ, err := Run()
	if err != nil {
		panic(err)
	}

	port := fmt.Sprintf(":%d", proxyServ.Config.Port)
	proxyServ.Engine.Run(port)
}
