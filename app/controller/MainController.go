package controller

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm();
	fmt.Println(r.Form)
	fmt.Println("hello world")
}
