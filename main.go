package main

import (
	"blog/app/util"
	"blog/app/route"
)

func init() {

}

func main() {
	util.Killprocess()
	route.Test();
}
