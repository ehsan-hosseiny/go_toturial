package main

import (
	"fmt"
	"log"

	"github.com/ehsan-hosseiny/gorm/config"
	"github.com/ehsan-hosseiny/gorm/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)

	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable Timezone=Asia/Tehran",
		config.AppConfig.Database.PostgresHost,
		config.AppConfig.Database.PostgresUser,
		config.AppConfig.Database.PostgresPassword,
		config.AppConfig.Database.PostgresDb,
		config.AppConfig.Database.PostgresPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)

	}

	var users []models.User
	err = db.Model(&models.User{}).Where("id",1).Find(&users).Error
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range users {
		fmt.Println("id : ",v.ID,"first name : ", v.FirstName, "last name : ", v.LastName)
	}

	// fmt.Println(users)

	// db.Exec("CREATE TYPE gender AS ENUM('Male','Female')")
	// err = db.AutoMigrate(&models.User{})
	// if err != nil {
	// 	log.Fatal(err)

	// }

	// email := "ehsanhossin@gmail.com"
	// err = db.Create(&models.User{
	// 	FirstName: "ehsan",
	// 	LastName:  "hosseiny",
	// 	Mobile:    "09122438871",
	// 	Email:     &email,
	// 	Age:       36,
	// 	Gender:    "Male",
	// }).Error

	// if err != nil {
	// 	log.Fatal(err)
	// }

}
