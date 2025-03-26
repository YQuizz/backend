package database

import (
	"errors"
	"log"
	"os"
	"sync"
	"time"
	"yquiz_back/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func InitDatabase() error {
	var err error

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return errors.New("DATABASE_URL is not set")
	}

	const maxAttempts = 5

	for i := 0; i < maxAttempts; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("connected to database")
			//TODO: uncomment this when the database is ready
			SyncDatabase()
			return nil
		}
		log.Printf("failed to connect database: %v", err)
		time.Sleep(2 * time.Second)
	}
	return errors.New("Failed to connect to database after multiple attempts" + err.Error())
}

func SyncDatabase() {
	tables := []interface{}{
		&models.User{}, &models.Session{}, &models.SessionParticipant{},
		&models.Question{}, &models.Answer{}, &models.Quiz{},
		&models.MonitoringLog{}, &models.UserAnswer{}, &models.Class{},
	}

	for _, table := range tables {
		if err := DB.AutoMigrate(table); err != nil {
			log.Fatalf("Failed to sync table %T: %v", table, err)
		}
	}

	log.Println("Database synced successfully")
}

func GetDB() (*gorm.DB, error) {
	var err error

	once.Do(func() {
		err = InitDatabase()
	})

	if err != nil {
		return nil, err
	}

	return DB, nil
}
