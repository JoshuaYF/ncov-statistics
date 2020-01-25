package router

import (
	"github.com/gin-gonic/gin"
	"ncov-statistics/service"
	"net/http"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(code int, data interface{}) {
	if code == SYSTEM_ERROR {
		g.C.JSON(http.StatusInternalServerError, Response{
			Code: code,
			Msg:  GetMsg(code),
			Data: data,
		})
		return
	}
	g.C.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  GetMsg(code),
		Data: data,
	})
	return
}

func InitRouter() *gin.Engine {

	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.LoadHTMLGlob("./views/*")

	apiGroup := r.Group("/api")

	apiGroup.GET("/all", All)
	apiGroup.GET("/areas", Area)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "hello Go",
		})
	})
	return r
}

func All(c *gin.Context) {
	var (
		appG = Gin{C: c}
	)

	appG.Response(SUCCESS, service.GetAll())
}

func Area(c *gin.Context) {
	var (
		appG = Gin{C: c}
	)

	appG.Response(SUCCESS, service.GetAllData())
}
