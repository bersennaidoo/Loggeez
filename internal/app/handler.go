package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *App) handleProduce(c *gin.Context) {
	req := ProduceRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	off, err := a.Log.Append(req.Record)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	res := ProduceResponse{Offset: off}

	c.JSON(http.StatusOK, res)
}

func (a *App) handleConsume(c *gin.Context) {

	var req ConsumeRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record, err := a.Log.Read(req.Offset)
	if err == ErrOffsetNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := ConsumeResponse{Record: record}

	c.JSON(http.StatusOK, res)
}
