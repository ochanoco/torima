package main

import (
	"github.com/ochanoco/proxy/auth/cloud"
	"github.com/ochanoco/proxy/auth/line"
	"github.com/ochanoco/proxy/example"
)

func main() {
	if VER == cloud.NAME {
		cloud.Main()
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
