package utils

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var UnauthorizedErrorTag = "failed to authorize users"
var FailedToSplitErrorTag = "failed to split error tag"

var errorStatusMap = map[string]int{
	UnauthorizedErrorTag: http.StatusUnauthorized,
}

func MakeError(e error, tag string) error {
	return fmt.Errorf("%s: %v", tag, e)
}

func SplitErrorTag(err error) (string, error) {
	errMsg := err.Error()

	splited := strings.Split(errMsg, ":")
	if len(splited) < 1 {
		return "", MakeError(err, FailedToSplitErrorTag)
	}

	return splited[0], nil
}

func FindStatusCodeByErr(err *error) int {
	var statusCode = http.StatusInternalServerError

	tag, splitErr := SplitErrorTag(*err)
	if splitErr != nil {
		return statusCode
	}

	if val, ok := errorStatusMap[tag]; ok {
		statusCode = val
	}

	return statusCode
}

func AbordGin(err error, c *gin.Context) {
	statusCode := FindStatusCodeByErr(&err)
	tag, _ := SplitErrorTag(err)
	fmt.Printf("error: %d, %v, %v", statusCode, err, tag)

	c.Status(statusCode)
	c.Writer.WriteString(Scripts)
	c.Writer.WriteString(BackHTML)
	c.Abort()
}
