package main

import (
	_ "beeblog/models"
	_ "beeblog/routers"
	_ "beeblog/utils"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego"
)

func main() {
	//utils.AllLogsGet()
	beego.Run()
}
