package main

import (
	"fmt"
	"rsc.io/quote"
	"github.com/beego/beego/v2/server/web"
)
func main() {
	fmt.Println("hello there!")
	fmt.Println(quote.Go())
	web.Run()
}