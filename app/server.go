package app

import (
	"net/http"
	"fmt"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello myroute!")
}
func Push(path string, method string) {
	mux := &MyMux{}
	http.ListenAndServe(":9091", mux)
	//http.HandleFunc("/", controller.Index)
	//err := http.ListenAndServe(":9090", nil) //设置监听的端口
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}
}
