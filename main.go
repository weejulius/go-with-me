package main

import (
	"github.com/astaxie/beego"
	"github.com/weejulius/go-with/controllers"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Run()
}
