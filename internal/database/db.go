package database

import (
	"job-portal/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open() (*gorm.DB, error) {
	dsn := "host=postgres user=postgres password=admin dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// err = db.Migrator().DropTable(&models.User{}, &models.Company{}, &models.Job{})
	// if err != nil {
	// 	// If there is an error while migrating, log the error message and stop the program
	// 	return nil, err
	// }
	err = db.Migrator().AutoMigrate(&models.User{}, &models.Company{}, &models.Job{})
	if err != nil {
		// If there is an error while migrating, log the error message and stop the program
		return nil, err
	}
	return db, nil
}
