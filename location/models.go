package locations

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Location struct {
	Id                    string
	LocationType          string
	LocationName          string
	AddressFirstLine      string
	AddressTown           string
	AddressPostCode       string
	AddressState          string
	AddressCountry        string
	AddressCountryCodeISO string
	LocationPoint         LocationPoint
}

type LocationPoint struct {
	coordinates []uint64
}

func LocationsRegister(router *gin.RouterGroup) {
	router.POST("/", LocationGet)
}

func LocationGet(c *gin.Context) {
	c.JSON(http.StatusOK, locations)
}

var locations = []Location{
    {
		Id: "1", 
		LocationType: "shop", 
		LocationName:          "Eco Freaks Emporium",
		AddressFirstLine:      "Havant High Street",
		AddressTown:           "Havant",
		AddressPostCode:       "PO91QD"
		AddressState:          "Hampshire",
		AddressCountry:        "UK",
		AddressCountryCodeISO: "GB",
	},
}