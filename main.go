package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stclaird/greenmap/location"
)

func main() {

	r := gin.Default()
	v1 := r.Group("/api")

	location.LocationsRegister(v1.Group("/location"))

	r.Run()
}
