package main

import (
	_ "deploy/config"
	"deploy/route"
	"fmt"
)

func main() {
	r := route.InitRouter()
	err := r.Run() // default 8080
	if err != nil {
		fmt.Printf("启动失败 %s", err.Error())
	}
}
