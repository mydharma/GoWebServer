package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	//io.WriteString(w, "Yes, it works")
	io.WriteString(w,
			`<!DOCTYPE html>
			<html>
			<body>

			<h1>
					   It works!!!
			</h1>

			</body>
			</html>
			`)
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
