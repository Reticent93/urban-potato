package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
	} else if r.URL.Path == "/contact" {
	fmt.Fprint(w, "To get in touch, please send email to <a href=\"mailto.support@urban-potato.com\">support@urban-potato.com</a>.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>We could not find the page you were looking for :(</h1><p>Please email us if you keep being sent to an invalid page</p>")
	}
}


func main() {
	
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
	
}