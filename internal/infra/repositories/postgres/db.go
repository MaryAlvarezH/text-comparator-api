package postgres

import (
	"fmt"
	"sync"
	"time"

	"github.com/MaryAlvarezH/text-comparator/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var onceDBLoad sync.Once

var tables = []interface{}{
	&entity.User{},
	&entity.TextComparison{},
}

func connect() *gorm.DB {
	onceDBLoad.Do(func() {
		source := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s",
			"0.0.0.0",
			"root",
			"root",
			"text_comparator",
			"5432",
		)

		var i int
		for {
			var err error
			if i >= 30 {
				panic("failed to connect: " + source)
			}
			time.Sleep(3 * time.Second)
			db, err = gorm.Open(postgres.Open(source), &gorm.Config{})
			if err != nil {
				log.Info("Retrying connection...", err)
				i++
				continue
			}
			break
		}
		migrate()
		log.Info("Connected to db!")
	})

	return db
}

func migrate() {
	for _, table := range tables {
		db.AutoMigrate(table)
	}
}
