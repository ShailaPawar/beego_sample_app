package routers

import (
	"beego_sample_app/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/data/", &controllers.ViewDataController{}, "get:ViewData")
}
