package line

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/ochanoco/proxy/core"
)

func AttachUserIdDirector(ochanocoProxy *core.OchanocoProxy, req *http.Request, c *gin.Context) (bool, error) {
	session := sessions.Default(c)
	userId := session.Get("userId")

	switch userId.(type) {
	case string:
		req.Header.Set("X-Ochanoco-UserID", userId.(string))

	default:
		req.Header.Set("X-Ochanoco-UserID", "")
	}

	return core.CONTINUE, nil
}
