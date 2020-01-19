package main

import (
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"

	"crud-api-class/handler"
)

func main() {
	r := gin.Default()

	r.HTMLRender = loadTemplates("./templates")

	r.GET("/", handler.List)
	r.GET("/edit", handler.Edit)

	apiRouter := r.Group("/api")
	{
		apiRouter.POST("/create", handler.Create)
		apiRouter.POST("/update", handler.Update)
		apiRouter.GET("/delete", handler.Delete)
	}

	r.Run(":8080")
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/includes/*.html")
	if err != nil {
		panic(err.Error())
	}

	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
