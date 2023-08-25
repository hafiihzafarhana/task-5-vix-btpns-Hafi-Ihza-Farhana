package database

import (
	"log"
	"rakamin/projectfinal/config"
	"rakamin/projectfinal/models"
	"rakamin/projectfinal/util"

	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	db, err = gorm.Open(config.GetDBConfig())
	if err != nil {
		log.Fatalln(err.Error())
	}

	if err = db.AutoMigrate(&models.RoleModel{}, &models.UserModel{}, &models.PhotoModel{}); err != nil {
		log.Fatalln(err.Error())
	}

	// if err = SeedData(db); err != nil {
	// 	log.Fatalln(err.Error())
	// }

	log.Println("Connected to DB!")
}

func GetPostgresInstance() *gorm.DB {
	return db
}

func SeedData(db *gorm.DB) error {
	roles := []models.RoleModel{
		{ID: 1, Role: "Admin"},
		{ID: 2, Role: "User"},
	}

	// Create roles
	for _, role := range roles {
		if err := db.Create(&role).Error; err != nil {
			return err
		}
	}

	hashedPassword, _ := util.HashPassword("adminpassword")

	// Add more seed data as needed...
	adminUser := models.UserModel{
		Username: "adminuser",
		Email:    "admin@example.com",
		Password: hashedPassword,
		RoleID:   1, // Assign the role ID from your manually defined roles
	}

	if err := db.Create(&adminUser).Error; err != nil {
		return err
	}

	return nil
}
