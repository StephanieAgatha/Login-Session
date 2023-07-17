package main

import (
	"fmt"
	"learn-session-login-logout/config"
	"learn-session-login-logout/controller"
	"net/http"
)

func main() {
	config.InitDB()

	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/register", controller.Register)
	http.HandleFunc("/logout", controller.Logout)

	fmt.Println("Webservice running...")
	http.ListenAndServe("localhost:3000", nil)
}
