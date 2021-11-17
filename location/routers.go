package location

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LocationsRegister(router *gin.RouterGroup) {
	router.GET("/", LocationGet)
}

func LocationGet(c *gin.Context) {
	locations, err := GetAllLocations()
	if err != nil {
		fmt.Printf("err %i", err)
	}
	c.JSON(http.StatusOK, locations)
}

// var locations = []Location{
// 	{
// 		Id:                    "1",
// 		LocationType:          "shop",
// 		LocationName:          "Eco Freaks Emporium",
// 		AddressFirstLine:      "Havant High Street",
// 		AddressTown:           "Havant",
// 		AddressPostCode:       "PO9 1QD",
// 		AddressState:          "Hampshire",
// 		AddressCountry:        "UK",
// 		AddressCountryCodeISO: "GB",
// 		LocationPoint: LocationPoint{
// 			Lat: 50.8530324,
// 			Lon: -0.985761,
// 		},
// 	},
// 	{
// 		Id:                    "1",
// 		LocationType:          "vehiclecharge",
// 		LocationName:          "Tesco Extra",
// 		AddressFirstLine:      "Solent Rd",
// 		AddressTown:           "Havant",
// 		AddressPostCode:       "PO9 1TR",
// 		AddressState:          "Hampshire",
// 		AddressCountry:        "UK",
// 		AddressCountryCodeISO: "GB",
// 		LocationPoint: LocationPoint{
// 			Lat: 50.8508107,
// 			Lon: -0.9887221,
// 		},
// 	},
// }
