package directors

import (
	"bytes"
	"net/http/httputil"

	"github.com/ochanoco/torima/core"
	"github.com/ochanoco/torima/utils"
)

func MakeLogDirector(flag string) core.TorimaDirector {
	return func(c *core.TorimaDirectorPackageContext) (core.TorimaPackageStatus, error) {
		request, err := httputil.DumpRequest(c.Target, true)

		if err != nil {
			err = utils.MakeError(err, "failed to dump headers to json: ")
			return core.ForceStop, err
		}

		splited := bytes.Split(request, []byte("\r\n\r\n"))

		header := splited[0]
		headerLen := len(header)

		body := request[headerLen:]

		l := c.Proxy.Database.CreateRequestLog(string(header), body, flag)
		_, err = l.Save(c.Proxy.Database.Ctx)

		if err != nil {
			err = utils.MakeError(err, "failed to save request: ")
			return core.ForceStop, err
		}

		return core.Keep, err
	}
}

var BeforeLogDirector = MakeLogDirector("before")
var AfterLogDirector = MakeLogDirector("after")
