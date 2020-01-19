package handler

import (
	"log"
	"net/http"
	"strconv"

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

func Edit(c *gin.Context) {
	id := c.Query("id")
	idu, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Fatal("fail convert id: ", err)
	}

	url, err := domain.GetElementByID(idu)
	if err != nil {
		log.Fatal("fail get by id: ", err)
	}

	c.HTML(http.StatusOK, "edit.html", gin.H{
		"ID":   id,
		"Link": url.Link,
	})
}
