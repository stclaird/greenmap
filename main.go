package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stclaird/greenmap/locations"
)

func main() {

	r := gin.Default()
	v1 := r.Group("/api")

	locations.ArticlesRegister(v1.Group("/location"))

	r.run()
}
