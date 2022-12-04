package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const BASE_NEXT_PATH = "../app/out/"
const OCHANOCO_LOGIN_URL = "http://localhost:8080/login"

func main() {
	secret := []byte("secret")

	r := gin.Default()
	r.LoadHTMLGlob(BASE_NEXT_PATH + "/*.html")
	r.Static("/_next/", BASE_NEXT_PATH+"_next/")

	store := cookie.NewStore(secret)
	r.Use(sessions.Sessions("mysession", store))

	InitProxyWeb(r)
	InitIdPWeb(r)

	r.Run()
}
