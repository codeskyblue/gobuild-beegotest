package main

import (
	"github.com/astaxie/beego"
	"github.com/shxsun/gobuild-beegotest/controllers"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Run()
}
