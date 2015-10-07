package main                                                                   
                                                                               
import (                                                                       
    "github.com/astaxie/beego"                                                 
    "github.com/astaxie/beego/context"                                         
    )                                                                          
                                                                               
func gheorghe() string{                                                        
    return "hello world!!"                                                              
}                                                                              
                                                                               
func main() {                                                                  
    beego.Get("/", func(ctx *context.Context){                                 
        ctx.Output.Body([]byte(gheorghe()))                                    
    })                                                                         
    beego.Run()                                                                
}          
