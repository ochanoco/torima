package core

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var unauthorizedErrorTag = "failed to authorize users"
var failedToSplitErrorTag = "failed to split error tag"

var errorStatusMap = map[string]int{
	unauthorizedErrorTag: http.StatusUnauthorized,
}

func makeError(e error, tag string) error {
	if e == nil {
		return nil
	}

	return fmt.Errorf("%s: %v", tag, e)
}

func splitErrorTag(err error) (string, error) {
	errMsg := err.Error()

	splited := strings.Split(errMsg, ":")
	if len(splited) < 1 {
		return "", makeError(err, failedToSplitErrorTag)
	}

	return splited[0], nil
}

func findStatusCodeByErr(err *error) int {
	var statusCode = http.StatusInternalServerError

	tag, splitErr := splitErrorTag(*err)
	if splitErr != nil {
		return statusCode
	}

	if val, ok := errorStatusMap[tag]; ok {
		statusCode = val
	}

	return statusCode
}

func abordGin(err error, c *gin.Context) {
	statusCode := findStatusCodeByErr(&err)
	tag, _ := splitErrorTag(err)
	fmt.Printf("error: %d, %v, %v", statusCode, err, tag)

	c.Status(statusCode)
	c.Writer.WriteString(scripts)
	c.Writer.WriteString(backHTML)
	c.Abort()
}
