package routers

import (
	"github.com/vladiibine/go-beego-already/src/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
