package location

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LocationsRegister(router *gin.RouterGroup) {
	router.GET("/", LocationGet)
	router.POST("/search", LocationSearch)
}

func LocationGet(c *gin.Context) {
	locations, err := GetLocations()
	if err != nil {
		fmt.Printf("err %i", err)
	}
	c.JSON(http.StatusOK, locations)
}

func LocationSearch(c *gin.Context) {
	locations, err := GetLocations()
	if err != nil {
		fmt.Printf("err %i", err)
	}
	c.JSON(http.StatusOK, locations)
}
