package proxy

import (
	"fmt"

	"github.com/ochanoco/torima/core"
)

func printConfig(config *core.TorimaConfig) {
	fmt.Println("default_origin:", config.DefaultOrigin)
	fmt.Println("host:", config.Host)
	fmt.Println("port:", config.Port)
	fmt.Println("scheme:", config.Scheme)

	fmt.Println("skip_auth_list:", config.SkipAuthList)
	fmt.Println("protection_scope:", config.ProtectionScope)
	fmt.Println("web_root:", config.WebRoot)
}
