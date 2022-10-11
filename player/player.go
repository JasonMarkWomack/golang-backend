package player

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:admin@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"

type Player struct {
	gorm.Model
	FirstName string `json:"fisrtname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to Database")
	}
	DB.AutoMigrate(&Player{})
}

func GetPlayer(c *fiber.Ctx) error {
	id := c.Params("id")
	var player Player
	DB.Find(&player, id)
	return c.JSON(&player)
}

func SavePlayer(c *fiber.Ctx) error {
	player := new(Player)
	if err := c.BodyParser(player); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Create(&player)
	return c.JSON(&player)
}

func DeletePlayer(c *fiber.Ctx) error {
	id := c.Params("id")
	var player Player
	DB.First(&player, id)
	if player.Email == "" {
		return c.Status(500).SendString("Player not available")
	}

	DB.Delete(&player)
	return c.SendString("Player is deleted!!!")
}

func UpdatePlayer(c *fiber.Ctx) error {
	id := c.Params("id")
	player := new(Player)
	DB.First(&player, id)
	if player.Email == "" {
		return c.Status(500).SendString("Player not available")
	}
	if err := c.BodyParser(&player); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Save(&player)
	return c.JSON(&player)
}
