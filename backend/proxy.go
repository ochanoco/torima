package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var patternSpecialPath = regexp.MustCompile(`^\/ochanoco\/`)

const CONTINUE = true
const FINISHED = false



type OchanocoPackageArgument interface{ *http.Request | *http.Response }

func runAllPackage[T OchanocoPackageArgument](
	pkgs []func(*OchanocoProxy, T, *gin.Context) bool,
	args T, proxy *OchanocoProxy, c *gin.Context) {

	flogLogs := NewFlowLogs()

	for _, pkg := range pkgs {
		isContinuing := pkg(proxy, args, c)

		flogLogs.Add(pkg, isContinuing)

		if !isContinuing {
			break
		}
	}

	flogLogs.Show()
}

/**
 * Directors is a list of functions that modify the
 * request before it is sent to the target server.
 **/
func (proxy *OchanocoProxy) Director(req *http.Request, c *gin.Context) {
	req.URL.Scheme = "http"

	runAllPackage(proxy.Directors, req, proxy, c)
	LogReq(req)
}

/**
  * ModifyResponses is a list of functions that modify the
  * response before it is sent to the client.
**/
func (proxy *OchanocoProxy) ModifyResponse(res *http.Response, c *gin.Context) error {
	runAllPackage(proxy.ModifyResponses, res, proxy, c)
	return nil
}

func LoginPathDirector(proxy *OchanocoProxy, req *http.Request, c *gin.Context) bool {
	fmt.Printf("Path: %v\n", req.URL.Path)

	if req.URL.Path != "/ochanoco/redirect" {
		return CONTINUE
	}

	req.URL.Host = PROXYWEB_HOST

	callback_path := req.URL.Query().Get("callback_path")

	if callback_path == "" {
		callback_path = "/"
	}

	session := sessions.Default(c)
	session.Set("callback_path", callback_path)

	return FINISHED
}

func CallbackPathDirector(proxy *OchanocoProxy, req *http.Request, c *gin.Context) bool {
	if req.URL.Path != "/ochanoco/callback" {
		return CONTINUE
	}

	token := req.URL.Query().Get("authorization_code")
	if token == "" {
		panic("failed to get authorization_code")
	}

	session := sessions.Default(c)
	session.Set("authorization_code", token)

	if session.Save() != nil {
		panic("failed to save authorization_code")
	}

	callbackPath := session.Get("callback_path")
	if callbackPath == nil {
		callbackPath = "/"
	}

	req.URL.Path = callbackPath.(string)
	req.URL.Host = req.Host
	req.URL.RawQuery = ""

	return FINISHED
}

func NextStaticFileDirector(proxy *OchanocoProxy, req *http.Request, c *gin.Context) bool {
	isOchanocoPath := patternSpecialPath.Match([]byte(req.URL.Path))

	if isOchanocoPath {
		req.URL.Host = PROXYWEB_HOST

		return FINISHED
	}

	return CONTINUE
}

func AuthDirector(proxy *OchanocoProxy, req *http.Request, c *gin.Context) bool {
	loginRedirectURL, err := url.Parse(PROXY_REDIRECT_URL)

	if err != nil {
		GoToErrorPage("failed parse", err, req)
		return FINISHED
	}

	isCleanContent := passIfCleanContent(req)
	isAuthed := authenticateRequest(req)
	mustRedirect := !isCleanContent && !isAuthed

	fmt.Printf("must auth redirect: %v\ncleanContent: %v, authed %v\n\n", mustRedirect, isCleanContent, isAuthed)

	if mustRedirect {
		req.URL.Scheme = loginRedirectURL.Scheme
		req.URL.Host = loginRedirectURL.Host
		req.URL.Path = "/ochanoco/redirect"

		return FINISHED
	}

	return CONTINUE
}

func DefaultDirector(proxy *OchanocoProxy, req *http.Request, c *gin.Context) bool {
	project, err := proxy.Database.FindServiceProviderByHost(req.Host)

	if err != nil {
		msg := fmt.Sprintf("failed to get destination site (%s)", req.Host)
		GoToErrorPage(msg, err, req)
		return FINISHED
	}

	req.URL.Host = project.DestinationIP

	req.Header.Set("User-Agent", "ochanoco")
	req.Header.Set("X-Ochanoco-Proxy-Token", "<proxy_token>")

	return CONTINUE
}

func MainModifyResponse(proxy *OchanocoProxy, resp *http.Response) {
	fmt.Printf("=> %v\n", resp.Request.URL)
}
