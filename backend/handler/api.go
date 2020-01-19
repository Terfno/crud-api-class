package handler

import (
	"log"
	"net/http"
	"strconv"

	"crud-api-class/domain"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	c.Request.ParseForm()
	url := c.Request.Form["url"][0]

	err := domain.CreateNewElement(url)
	if err != nil {
		log.Fatal("fail create: ", err)
	}

	c.Redirect(http.StatusFound, "/")
}

func Update(c *gin.Context) {
	c.Request.ParseForm()
	url := c.Request.Form["url"][0]

	ids := c.Query("id")
	id, err := strconv.ParseUint(ids, 10, 64)
	if err != nil {
		log.Fatal("fail convert id: ", err)
	}

	err = domain.UpdateElement(id, url)
	if err != nil {
		log.Fatal("fail update: ", err)
	}

	c.Redirect(http.StatusFound, "/")
}

func Delete(c *gin.Context) {
	ids := c.Query("id")
	id, err := strconv.ParseUint(ids, 10, 64)
	if err != nil {
		log.Fatal("fail convert id: ", err)
	}

	err = domain.DeleteElement(id)
	if err != nil {
		log.Fatal("fail delete: ", err)
	}

	c.Redirect(http.StatusFound, "/")
}
