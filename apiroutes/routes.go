package apiroutes

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

var HotelInitRoute, HotelShowRoute string
var GuestCreateRoute, GuestCheckoutRoute,GuestAllRoute, GuestGetRoute, GuestDeleteRoute string

func init() {
	envMap, err := godotenv.Read(".dev.env")
	if err != nil {
		log.Fatal(err)
	}
	apiBase := envMap["API_BASE"]
	// API_BASE=127.0.0.1:4000
	HotelInitRoute = fmt.Sprintf("http://%s/hotel", apiBase)
	HotelShowRoute = HotelInitRoute

	GuestAllRoute = fmt.Sprintf("http://%s/guests", apiBase)
	
	GuestCheckoutRoute = fmt.Sprintf("http://%s/guest", apiBase)
	GuestCreateRoute = GuestCheckoutRoute
	GuestDeleteRoute = GuestCheckoutRoute
	GuestDeleteRoute = GuestCheckoutRoute
	GuestGetRoute = GuestCheckoutRoute
}
