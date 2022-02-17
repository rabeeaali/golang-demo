package postControllers

import (
	"fiber/app/models"
	"fiber/app/requests"
	"fiber/config"
	"fiber/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func Index(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(utils.Paginate(config.DB, &models.Post{}, page))
}

func StorePost(c *fiber.Ctx) error {
	var data map[string]string

	// get the data from request
	if err := c.BodyParser(&data); err != nil {
		return utils.ErrorJSON(c, fiber.StatusBadRequest, "error", "please try to put some data")
	}

	post := models.Post{
		Title:  data["title"],
		Desc:   data["description"],
		UserID: utils.GetUserID(c),
	}

	// validate Error
	if err := requests.PostStoreValidate(post.Title, post.Desc); err != nil {
		return utils.ErrorJSON(c, fiber.StatusUnprocessableEntity, "validation errors", err)
	}

	// store data
	config.DB.Create(&post)

	return c.JSON(fiber.Map{
		"message": "Added Post Successfully",
	})
}

func ShowPost(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println(id)
	var post models.Post

	if err := config.DB.First(&post, id).Error; err != nil {
		return utils.ErrorJSON(c, fiber.StatusBadRequest, "error", "Post not found")
	}

	return utils.SuccessJSON(c, fiber.StatusOK, "OK!", post)
}

func UpdatePost(c *fiber.Ctx) error {
	var data map[string]string

	post := models.Post{
		Title:  data["title"],
		Desc:   data["description"],
		UserID: utils.GetUserID(c),
	}

	// Decode request data
	if err := c.BodyParser(&post); err != nil {
		return utils.ErrorJSON(c, fiber.StatusBadRequest, "error", "please try to put some data")
	}

	// Validate error
	if err := requests.PostUpdateValidate(post.Title, post.Desc); err != nil {
		return utils.ErrorJSON(c, fiber.StatusUnprocessableEntity, "validation errors", err)
	}

	// Find the post data by id
	if err := config.DB.Where("id = ?", c.Params("id")).First(&post).Error; err != nil {
		return utils.ErrorJSON(c, fiber.StatusBadRequest, "error", "Post not found")
	}

	// Update the post data
	config.DB.Model(&post).Updates(post)

	return utils.SuccessJSON(c, fiber.StatusOK, "Updated Post Successfully", "")
}

func DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.Post

	if err := config.DB.First(&post, id).Error; err != nil {
		return utils.ErrorJSON(c, fiber.StatusBadRequest, "error", "Post not found")
	}
	config.DB.Delete(&post)
	return utils.SuccessJSON(c, fiber.StatusOK, "Deleted Post Successfully", "")
}
