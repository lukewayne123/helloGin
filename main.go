package main

import (
	//"github.com/gin-gonic/gin"
	"helloGin/routers"
)

func main() {
	r := routers.Init()
	//r.Run("10.0.2.15:8080")
	r.Run(":8080")
}
