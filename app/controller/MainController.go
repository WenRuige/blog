package controller

import (
	"fmt"
	"net/http"
)

func Test2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
}
