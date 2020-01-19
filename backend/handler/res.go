package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"crud-api-class/domain"
)

func List(c *gin.Context) {
	urls, err := domain.GetAll()
	if err != nil {
		log.Fatal("fail get: ", err)
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"urls": urls,
	})
}
