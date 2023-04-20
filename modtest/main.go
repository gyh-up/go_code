package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	ginServ := gin.Default()
	ginServ.Any("/higin", WebRoot)
	ginServ.Run(":8888")
}

func WebRoot(context *gin.Context) {
	context.String(http.StatusOK, "hello, world")
}