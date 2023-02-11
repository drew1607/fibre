package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "host=localhost user=dnerd password= dbname=dnerd port=5433 sslmode=disable TimeZone=Asia/Shanghai"

type Books struct {
	FirstName string `json:"fisrtname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Country   string `json:"country"`
	State     string `json:"state"`
	Model     string `json:"Model"`
	Rep       string `json:"rep"`
}

func main() {
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to Database")
	}
	DB.AutoMigrate(&Books{})
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":8081"))
}
