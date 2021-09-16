package example

import (
	"fmt"
	"github.com/micro/services/clients/go/time"
	"os"
)

// Get the timezone info for a specific location
func GetTheTimezoneInfoForAspecificLocation() {
	timeService := time.NewTimeService(os.Getenv("MICRO_API_TOKEN"))
	rsp, err := timeService.Zone(&time.ZoneRequest{
		Location: "London",
	})
	fmt.Println(rsp, err)
}