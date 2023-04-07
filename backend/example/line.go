package example

import (
	"github.com/ochanoco/proxy/core"
)

func Main() {
	proxyServ := prepare()
	proxyServ.Engine.Run(core.PROXY_PORT)
}

const LINE_NAME = "line_example"
