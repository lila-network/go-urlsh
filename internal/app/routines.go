package app

import (
	"codeberg.org/lauralani/go-urlsh/models"
	"context"
	"github.com/jasonlvhit/gocron"
	"log"
)

func CleanupLoginsCronJob() {
	err := gocron.Every(1).Hours().Do(CleanupLogins)
	if err != nil {
		return
	}
	<-gocron.Start()
}

func CleanupLogins() {
	_, err := models.DB.NewDelete().Table("logins").Where("expires < NOW()").Exec(context.Background())
	if err != nil {
		log.Printf("[CleanupLogins] Error deleting logins: %v\n", err)
	}
	log.Printf("Cleaned up logins")
}
