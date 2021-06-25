package main

import (
	"fmt"
	"github.com/rebay1982/gonasa/internal/app/gonasa"
)

var apiKey = ""
var apiEmail = ""
var apiAccount = ""

func main() {

	fmt.Println("[ GoNASA ]")
	marsClient := gonasa.NewMarsAPIClient(apiKey)

	rovers, err := marsClient.MarsManifestAPI()

	if err != nil {
		fmt.Printf("Unable to retrieve rover details")
	}

	for _, rover := range rovers.Rovers {
		fmt.Printf("[ %s ]\n  Launch:\t%s\n  Landing:\t%s\n  Sols:\t\t%d\n\n",
			rover.Name,
			rover.LaunchDate,
			rover.LandingDate,
			rover.MaxSol)
	}

	fmt.Println("[ Done ]")
}
