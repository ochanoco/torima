package proxy

// func TestRemoveToken(t *testing.T) {
// 	msg := "hello"
// 	token := "hello"

// 	directors := []func(req *http.Request){}
// 	modifyRespes := []func(req *http.Response){}

// 	directors = append(directors, RemoveToken)
// 	proxy := NewOchancoProxy(directors, modifyRespes)

// 	t.Run("enc/dec token", func(t *testing.T) {
// 		simpleServ := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
// 			if req.URL.Query().Get("token") == token {
// 				t.Errorf("Error: removing token (1)")
// 			}

// 			_, err := req.Cookie("token")

// 			if err != nil {
// 				t.Errorf("Error: removing token (2)")
// 			}

// 			fmt.Fprintln(writer, msg)
// 		}))

// 		simpleDirector := func(req *http.Request) {
// 			url, _ := url.Parse(simpleServ.URL)
// 			req.URL.Scheme = url.Scheme
// 			req.URL.Host = url.Host
// 			req.URL.Path = "/"
// 		}

// 		directors = append(directors, simpleDirector)
// 		targetServ := httptest.NewServer(proxy.ReverseProxy)

// 		url := fmt.Sprintf("%s?token=%s", targetServ.URL, token)
// 		req, err := http.NewRequest(http.MethodGet, url, nil)

// 		if err != nil {
// 			t.Error(err)
// 		}

// 		req.AddCookie(&http.Cookie{Name: "token", Value: token})

// 		_, err = new(http.Client).Do(req)
// 		if err != nil {
// 			t.Error(err)
// 		}
// 	})
// }
