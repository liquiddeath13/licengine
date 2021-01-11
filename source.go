package main

import (
	"net/http"
)

func main() {
	initUsers()
	initWareList()
	http.HandleFunc("/rl", middleware(libHost))
	http.HandleFunc("/q", middleware(check))
	http.HandleFunc("/sl", middleware(softwareList))
	http.ListenAndServe(":1337", nil)
}
