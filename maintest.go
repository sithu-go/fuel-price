package main

import (
	"fuel-price/pkg/service"
	"log"
)

func main() {

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	service.CrawlFuelPrices()
}
