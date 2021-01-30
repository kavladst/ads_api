package main

import (
	"log"

	"github.com/kavladst/ads_api/internal/app/api"
)

func main() {
	apiApplication, err := api.New()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Fatal(apiApplication.Run())
	}
}
