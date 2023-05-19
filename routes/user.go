package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/mergen1212/fiber-api-go/datebase"
	"github.com/mergen1212/fiber-api-go/models"
)

type User struct {
	//serialazer
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUser(user models.User) User {
	return User{ID: user.ID, FirstName: user.FirstName, LastName: user.LastName}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	datebase.Database.DB.Create(&user)
	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)

}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	datebase.Database.DB.Find(&users)
	responseUsers := []User{}
	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)

	}
	return c.Status(200).JSON(responseUsers)
}

func findUser(id int, user *models.User) error {
	datebase.Database.DB.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("User does not exist")

	}
	return nil
}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User
	if err != nil {
		return c.Status(400).JSON(("Please ensure that : id is an integer"))
	}
	if err := findUser(id, &user); err != nil {
		c.Status(400).JSON(err.Error())
	}
	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}
