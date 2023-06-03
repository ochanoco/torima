package main

import (
	"github.com/ochanoco/proxy/example"
	"github.com/ochanoco/proxy/serv/cloud"
	"github.com/ochanoco/proxy/serv/line"
	"github.com/ochanoco/proxy/serv/tee"
)

func main() {
	if VER == cloud.NAME {
		cloud.Main()
	} else if VER == tee.NAME {
		tee.Main()
	} else if VER == line.NAME {
		line.Main()
	} else if VER == example.LINE_NAME {
		example.Main()
	} else {
		println("please set tags")
		println("go build --tags cloud")
		println("or")
		println("go build --tags line")
	}
}
