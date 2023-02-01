package service

import (
	"log"
	"strconv"

	"github.com/gocolly/colly"
)

type StationFuelPrice struct {
	Name          string
	Diesel        uint64
	PremiumDiesel uint64
	Octance92     uint64
	Octance95     uint64
}

func CrawlFuelPrices() {

	fuelPrices := map[string][]StationFuelPrice{}
	c := colly.NewCollector()

	var division string
	c.OnHTML("div table tbody tr", func(h *colly.HTMLElement) {
		var stationName string
		var dieselPrice uint64
		var dieselPremiumPrice uint64
		var octance92Price uint64
		var octance95Price uint64

		h.ForEach("td", func(i int, h *colly.HTMLElement) {
			switch i {
			case 0:
				// fuelPrices[h.Text] = []*StationFuelPrice{}
				// log.Println(len(h.Text), h.Text, "DI")
				if len(h.Text) != 0 {
					division = h.Text
				}
			case 1:
				stationName = h.Text
			case 2:
				price, _ := strconv.ParseUint(h.Text, 10, 64)
				dieselPrice = price
			case 3:
				price, _ := strconv.ParseUint(h.Text, 10, 64)
				dieselPremiumPrice = price
			case 4:
				price, _ := strconv.ParseUint(h.Text, 10, 64)
				octance92Price = price
			case 5:
				price, _ := strconv.ParseUint(h.Text, 10, 64)
				octance95Price = price
			}

		})

		stationFuelPrices := StationFuelPrice{
			Name:          stationName,
			Diesel:        dieselPrice,
			PremiumDiesel: dieselPremiumPrice,
			Octance92:     octance92Price,
			Octance95:     octance95Price,
		}
		fuelPrices[division] = append(fuelPrices[division], stationFuelPrices)

	})

	c.Visit("https://denkomyanmar.com/all-denko-station-daily-fuel-rates/")

	for _, k := range fuelPrices["Shan State"] {
		log.Println(k, "HEHE")
	}
}
