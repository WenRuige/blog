package app

import (
	"net/http"
	"log"
)

func push() {
	http.HandleFunc("/", path)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
