package example

import (
	"fmt"
	"net/http"
)

func targetServ(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("X-Ochanoco-UserID")
	fmt.Fprintf(w, 
		"<p>Hello! %v</p><br><a href='%v'>link</a>",
		userId,
		"/ochanoco/login?callback_path=/hello")
}
