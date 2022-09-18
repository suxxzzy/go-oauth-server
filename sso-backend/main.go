package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var corsHandler = cors.New(cors.Options{
	AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:3000"},
	AllowedMethods:   []string{http.MethodGet,http.MethodPost,http.MethodPut,http.MethodDelete,http.MethodOptions},
	AllowedHeaders:   []string{"Origin", "Accept", "Content-Type", "X-Requested-With"},
	AllowCredentials: true,
	MaxAge:           0,
	Debug:            true,
})

var handler = corsHandler.Handler(MakeWebHandler())

func handleServeLoginPage(w http.ResponseWriter, r *http.Request){
	fmt.Println("redirect started")
	http.Redirect(w, r, "/static/build/", http.StatusSeeOther)
}

func handleLogin(w http.ResponseWriter, r *http.Request){
	fmt.Println("handling login request...")
	//로그인 요청 처리: LDAP 서버에 인증 요청!
}


func MakeWebHandler() http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/loginpage", handleServeLoginPage).Methods("GET")
	mux.HandleFunc("/login", handleLogin).Methods("POST")
	fs := http.FileServer(http.Dir("./static/build"))
    mux.PathPrefix("/static/build/").Handler(http.StripPrefix("/static/build/", fs))
	return mux
}

func main() {
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal(err)
	}
}
