package adapter

import (
	"github.com/gin-gonic/gin"
	"github.com/phuhao00/wddd/blog/internal/app/adapter/respository"
	"github.com/phuhao00/wddd/blog/internal/app/adapter/service"
	"github.com/phuhao00/wddd/blog/internal/app/application/usecase"
	"github.com/phuhao00/wddd/blog/internal/app/domain/valueobject"
	"net/http"
)

var (
	blogSvr  = &service.Blog{}
	blogRepo = respository.BlogRepo{}
)

// Controller is a controller
type Controller struct{}

// Router is routing settings
func Router() *gin.Engine {
	r := gin.Default()
	ctrl := Controller{}
	r.LoadHTMLGlob("internal/app/adapter/view/*")
	r.GET("/", ctrl.index)
	r.POST("/update", ctrl.update)
	r.POST("/delete", ctrl.delete)
	r.POST("/get", ctrl.get)
	return r
}

func (ctrl Controller) index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Hello wddd",
	})
}

func (ctrl Controller) update(c *gin.Context) {
	id := c.Query("id")
	title := c.PostForm("title")
	content := c.PostForm("content")
	usecase.Update(blogSvr, &valueobject.Post{
		ID:      id,
		Title:   title,
		Content: content,
	})
}

func (ctrl Controller) delete(c *gin.Context) {
	id := c.Query("id")
	usecase.Delete(blogSvr, id)
}

func (ctrl Controller) get(c *gin.Context) {
	id := c.Query("id")
	usecase.Get(blogSvr, id)
}
