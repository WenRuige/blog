package route

import (
	"blog/app"
)

var routerConfig map[string]string
/*
	路由分发
*/
func RegisterController(path string, controller string) {

	app.Push(path, controller)

}
