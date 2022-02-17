package userControllers

import (
	"fiber/app/models"
	"fiber/config"
	utils2 "fiber/utils"
	"github.com/gofiber/fiber/v2"
)

// UserInfo get user info
func UserInfo(c *fiber.Ctx) error {
	var user models.User

	config.DB.Where("id = ?", utils2.GetUserID(c)).First(&user)

	return utils2.SuccessJSON(c, fiber.StatusOK, "user info!", user)
}

//UpdateUser update user info
//func UpdateUser(c *fiber.Ctx) error {
//	var user models.User
//
//	database.DB.Where("id = ?", utils.GetUserID(c)).First(&user)
//
//	return utils.SuccessJSON(c, fiber.StatusOK, "user info!", user)
//}
