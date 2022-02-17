package authControllers

import (
	"fiber/app/models"
	"fiber/app/requests"
	"fiber/config"
	"fiber/utils"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var data map[string]string

	// get the data from request
	if err := c.BodyParser(&data); err != nil {
		return utils.ErrorJSON(c, fiber.StatusBadRequest, "error", "please try to put some data")
	}

	var user = models.User{
		Email:    data["email"],
		Password: data["password"],
	}

	// validate Error
	if err := requests.LoginValidate(user.Email, user.Password); err != nil {
		return utils.ErrorJSON(c, fiber.StatusUnprocessableEntity, "validation errors", err)
	}

	// check if email already exist
	config.DB.Where("email = ?", user.Email).First(&user)
	if user.ID == 0 {
		return utils.ErrorJSON(c, fiber.StatusUnprocessableEntity, "error", models.ErrorAuth{
			Email: "This email does not exits in our credential.",
		})
	}

	// check if password is correct
	if err := user.CheckPassword([]byte(data["password"])); err != nil {
		return utils.ErrorJSON(c, fiber.StatusUnprocessableEntity, "error", models.ErrorAuth{
			Password: "password is not correct.",
		})
	}

	// generate access token
	token := utils.GenerateNewAccessToken(int(user.ID))

	// return response
	return utils.SuccessJSON(c, fiber.StatusOK, "OK!", models.LoginResult{
		UserInfo: user,
		Token:    token,
	})
}
