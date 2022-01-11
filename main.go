package main

import "net/http"

func main() {
	app := app{Username: "Username"}

	http.HandleFunc("/path", app.handle)
	http.ListenAndServe(":8090", nil)
}
