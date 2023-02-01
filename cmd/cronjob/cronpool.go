package cronjob

import (
	"fuel-price/pkg/repository"
	"fuel-price/pkg/service"
	"log"

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
	log.Println("HEHE")

	go func() {
		fuelPrices, err := service.CrawlFuelPrices()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(fuelPrices["Shan State"])
	}()

}
