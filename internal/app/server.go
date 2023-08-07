package app

import "github.com/gin-gonic/gin"

func (a *App) RunAPI(addr string) error {
	r := gin.Default()

	r.POST("/", a.handleProduce)
	r.GET("/", a.handleConsume)

	return r.Run(addr)
}
