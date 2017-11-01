package route

import (
	"blog/app"
	"fmt"

)

/*
	路由分发
*/
func RegisterController(path string, controller string) {


	fmt.Println(path,controller)
	app.Push(path,controller)

}
