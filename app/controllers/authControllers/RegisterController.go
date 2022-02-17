package authControllers

import (
	"errors"
	"fiber/app/models"
	"fiber/app/requests"
	"fiber/config"
	"fiber/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	// get the data from request
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// userControllers object
	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: data["password"],
	}

	// validate Error
	if err := requests.RegisterValidate(user.Name, user.Email, user.Password); err != nil {
		return utils.ErrorJSON(c, fiber.StatusUnprocessableEntity, "validation errors", err)
	}
	// check if the user is exits
	result := config.DB.Where("email = ?", user.Email).Take(&user)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return utils.ErrorJSON(c, fiber.StatusBadRequest, "error", models.ErrorAuth{
			Email: "the email has already been taken.",
		})
	}

	// hashed password
	user.HashPassword([]byte(user.Password))

	// store userControllers in DB
	config.DB.Create(&user)

	// response success
	return c.JSON(fiber.Map{
		"message": "User Registered Successfully",
	})
}
