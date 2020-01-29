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

	r.LoadHTMLGlob("./views/*.html")

	apiGroup := r.Group("/api")

	apiGroup.GET("/province", Province)
	apiGroup.GET("/trend", Trend)
	apiGroup.GET("/areas", Area)
	apiGroup.GET("/map/info", Map)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "武汉加油",
		})
	})

	r.GET("/trend", func(c *gin.Context) {
		c.HTML(http.StatusOK, "trend.html", gin.H{
			"title": "武汉加油",
		})
	})

	r.GET("/map", func(c *gin.Context) {
		c.HTML(http.StatusOK, "map.html", gin.H{
			"title": "武汉加油",
		})
	})

	return r
}

func Province(c *gin.Context) {
	var (
		appG = Gin{C: c}

		provinceName = c.Query("province_name")
	)

	appG.Response(SUCCESS, service.Province(provinceName))
}

func Trend(c *gin.Context) {
	var (
		appG = Gin{C: c}

		provinceName = c.Query("province_name")
	)

	appG.Response(SUCCESS, service.Trend(provinceName))
}

func Area(c *gin.Context) {
	var (
		appG = Gin{C: c}
	)

	appG.Response(SUCCESS, service.GetAllData())
}

func Map(c *gin.Context) {
	var (
		appG = Gin{C: c}
		provinceName = c.Query("province_name")
	)

	appG.Response(SUCCESS, service.Map(provinceName))
}
