package extension

import (
	gin_ninsho "github.com/ochanoco/ninsho/extension/gin"
)

var AUTH_PATH = gin_ninsho.NinshoGinPath{
	Unauthorized: "/auth/login",
	Callback:     "/auth/callback",
	AfterAuth:    "/_torima/back",
}
