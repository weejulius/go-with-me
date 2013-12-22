package main

import (
	"github.com/astaxie/beego"
	"github.com/weejulius/go-with-me/controllers"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Run()
}
