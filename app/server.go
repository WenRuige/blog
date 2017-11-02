package app

import (
	"blog/app/controller"
	"net/http"
	"log"
)

func Push(path string, method string) {
	http.HandleFunc("/", controller.Index)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
