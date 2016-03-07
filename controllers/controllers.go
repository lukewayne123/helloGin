package controllers

import (
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(200, "index.tpl", gin.H{
            "Website": "beego.me",
	"Email": "astaxie@gmail.com",
        })	
}

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
            "message": "hello",
        })
}
