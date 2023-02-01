package cronjob

import (
	"fuel-price/pkg/model"
	"fuel-price/pkg/repository"
	"fuel-price/pkg/service"
	"log"
	"strings"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

type CronConfig struct {
	DB   *gorm.DB
	Repo *repository.Repository
}
type cronPool struct {
	cronJob *cron.Cron
	db      *gorm.DB
	repo    *repository.Repository
}

func NewCronPool(cronConfig *CronConfig) *cronPool {
	cronJob := cron.New()

	return &cronPool{
		cronJob: cronJob,
		db:      cronConfig.DB,
		repo:    cronConfig.Repo,
	}
}

func (c *cronPool) StartCronPool() {
	c.cronJob.AddFunc("0 */4 * * *", c.crawlFuelPrices) // at every 4th hour

	c.cronJob.Start()
}

func (c *cronPool) crawlFuelPrices() {
	go func() {
		fuelPrices, err := service.CrawlFuelPrices()
		if err != nil {
			log.Println(err)
			return
		}

		for divisonName, stations := range fuelPrices {
			_ = stations

			// division part
			var divison model.Division
			if err := c.db.Where("name", divisonName).First(&divison).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					dNameArr := strings.Split(divisonName, " ")
					if err := c.db.Where("name LIKE ?", "%"+dNameArr[0]+"%").First(&divison).Error; err != nil {
						if err == gorm.ErrRecordNotFound {
							divison.Name = divisonName
							if err := c.db.Create(&divison).Error; err != nil {
								log.Println(err)
								return
							}
							log.Println("divison created ", divison)
						} else {
							log.Println(err)
							return
						}
					} else {
						divison.Name = divisonName
						if err := c.db.Debug().Model(&model.Division{}).Where("id", divison.ID).UpdateColumn("name", divisonName).Error; err != nil {
							log.Println(err)
							return
						}
					}
				} else {
					log.Println(err)
					return
				}
			}

			// station part
			for _, sInfo := range stations {
				var station model.Station
				if err := c.db.Where("name", sInfo.Name).First(&station).Error; err != nil {
					if err == gorm.ErrRecordNotFound {
						sNameArr := strings.Split(sInfo.Name, " ")
						if err := c.db.Where("name LIKE ?", "%"+sNameArr[0]+"%").First(&station).Error; err != nil {
							if err == gorm.ErrRecordNotFound {
								station.Name = sInfo.Name
								station.DivisionId = divison.ID
								if err := c.db.Create(&station).Error; err != nil {
									log.Println(err)
									return
								}
							} else {
								log.Println(err)
								return
							}
						} else {
							station.Name = sInfo.Name
							if err := c.db.Debug().Model(&model.Station{}).Where("id", station.ID).UpdateColumn("name", sInfo.Name).Error; err != nil {
								log.Println(err)
								return
							}
						}
					} else {
						log.Println(err)
						return
					}
				}

				log.Println(station, "HEHE")
				if station.ID == 0 {
					log.Fatal(station.ID, "station")
				}

				// Fuel log part
				fuelLog := model.FuelLog{
					StationId:     station.ID,
					DiselPrice:    sInfo.Diesel,
					PreDiselPrice: sInfo.PremiumDiesel,
					O95Price:      sInfo.Octance95,
					O92Price:      sInfo.Octance92,
				}

				if err := c.db.Debug().Create(&fuelLog).Error; err != nil {
					log.Println(err)
					return
				}
			}

		}
	}()

}
