package authControllers

import (
	"fiber/app/models"
	"fiber/app/requests"
	"fiber/config"
	"fiber/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

// ForgetPassword step 1
func ForgetPassword(c *fiber.Ctx) error {
	var data map[string]string

	// get the data from request
	if err := c.BodyParser(&data); err != nil {
		return utils.ErrorJSON(c, fiber.StatusBadRequest, "error", "please try to put some data")
	}

	// user object
	user := models.User{
		Email: data["email"],
	}

	// validate Error
	if err := requests.ForgetPasswordValidate(user.Email); err != nil {
		return utils.ErrorJSON(c, fiber.StatusBadRequest, "validation errors", err)
	}

	// check if email already exist
	config.DB.Where("email = ?", user.Email).First(&user)

	if user.ID == 0 {
		return utils.ErrorJSON(c, fiber.StatusUnprocessableEntity, "error", models.ErrorAuth{
			Email: "This email does not exits in our credential.",
		})
	}

	// create random numbers
	code := utils.RandNumbers()

	// reset password object
	ResetPasswordData := models.ResetPassword{
		Email: user.Email,
		Code:  code,
	}

	// delete old records
	config.DB.Where("email = ?", ResetPasswordData.Email).Delete(&ResetPasswordData)

	// create a new one
	config.DB.Create(&ResetPasswordData)

	// send an email
	go config.SendMail(user.Email, code, "reset", "Reset Password Code")

	// response success
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Sent Email Successfully",
	})
}

// CheckCode step 2
func CheckCode(c *fiber.Ctx) error {
	var data map[string]string

	// get the data from request
	if err := c.BodyParser(&data); err != nil {
		return utils.ErrorJSON(c, fiber.StatusBadRequest, "error", "please try to put some data")
	}

	// resetPassword object
	resetPassword := models.ResetPassword{
		Code: data["code"],
	}

	// validate Error
	if err := requests.CheckCodeValidate(resetPassword.Code); err != nil {
		return utils.ErrorJSON(c, fiber.StatusBadRequest, "validation errors", err)
	}

	config.DB.Where("code = ?", resetPassword.Code).First(&resetPassword)

	// check if  the code is there
	if resetPassword.ID == 0 {
		return utils.ErrorJSON(c, fiber.StatusUnprocessableEntity, "error", models.ErrorAuth{
			Code: "this code does not exits in our credential.",
		})
	}

	// check if  the code is expired
	if time.Now().Add(time.Hour * 1).Before(resetPassword.CreatedAt) {
		return utils.ErrorJSON(c, fiber.StatusUnprocessableEntity, "error", models.ErrorAuth{
			Code: "this code is expired",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "the code is valid",
	})
}

// ResetPassword step 3
func ResetPassword(c *fiber.Ctx) error {
	var data map[string]string

	// get the data from request
	if err := c.BodyParser(&data); err != nil {
		return utils.ErrorJSON(c, fiber.StatusBadRequest, "error", "please try to put some data")
	}

	// userControllers object
	user := models.User{
		Email:    data["email"],
		Password: data["password"],
	}

	// reset password object
	ResetPasswordData := models.ResetPassword{
		Email: user.Email,
	}

	config.DB.
		Where("email = ?", ResetPasswordData.Email).
		Find(&ResetPasswordData)

	if ResetPasswordData.ID == 0 {
		return utils.ErrorJSON(c, fiber.StatusUnprocessableEntity, "error", "Invalid email, try to reset again")
	}

	// validations
	if err := requests.ResetPasswordValidate(data["email"], data["password"], data["password_confirmation"]); err != nil {
		return utils.ErrorJSON(c, fiber.StatusBadRequest, "validation errors", err)
	}

	// check if email already exist
	config.DB.Where("email = ?", user.Email).First(&user)
	if user.ID == 0 {
		return utils.ErrorJSON(c, fiber.StatusUnprocessableEntity, "error", models.ErrorAuth{
			Email: "This email does not exits in our credential.",
		})
	}

	// hash password
	user.HashPassword([]byte(data["password"]))

	// update password
	config.DB.
		Model(&user).
		Where("email = ?", user.Email).
		Update("password", user.Password)

	// delete old records
	config.DB.
		Where("email = ?", ResetPasswordData.Email).
		Delete(&ResetPasswordData)

	// response success
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "password changed successfully",
	})
}
