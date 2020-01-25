package main

import (
	"ncov-statistics/router"
)

func main() {
	r := router.InitRouter()
	r.Run("0.0.0.0:80") // listen and serve on 0.0.0.0:8080
}
