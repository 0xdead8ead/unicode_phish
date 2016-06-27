package main

import (
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

func ServerStatus(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "UPUPUP\n")
}

/*
func Index(resp http.ResponseWriter, req *http.Request) {
	io.WriteString(resp, "INDEX\n")
}
*/

func Redirect(resp http.ResponseWriter, req *http.Request) {
	io.WriteString(resp, "Redirecting to SSO ;) \n")
}

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/status", ServerStatus)
	//rtr.HandleFunc("/index", Index)
	rtr.HandleFunc("/{rest:.*}", Redirect)

	http.Handle("/", rtr)

	//err := http.ListenAndServeTLS(":4443", "cert.pem", "key.pem", nil)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
