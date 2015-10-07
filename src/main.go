package main

import (
	_ "github.com/vladiibine/go-beego-already/src/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {
    beego.Get("/petre", func(ctx *context.Context){
            var a string
            var x *int
            a = string(*x)
            ctx.Output.Body([]byte("hello petre!!!"))
            ctx.Output.Body([]byte(a))
        })
	beego.Run()
}

