package postControllers

import (
	"fiber/app/models"
	"fiber/app/requests"
	"fiber/config"
	"fiber/utils"
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
		return err
	}

	post := models.Post{
		Title:  data["title"],
		Desc:   data["description"],
		UserID: utils.GetUserID(c),
	}

	// validate Error
	if err := requests.StoreValidate(post.Title, post.Desc); err != nil {
		return utils.ErrorJSON(c, fiber.StatusUnprocessableEntity, "validation errors", err)
	}

	// store data
	config.DB.Create(&post)

	return c.JSON(fiber.Map{
		"message": "Added Post Successfully",
	})
}

//func ShowPost(c *gin.Context) {
//	id := c.Param("id")
//	var post models.Post
//	if err := database.DB.First(&post, id).Error; err != nil {
//		util.ErrorJSON(c, http.StatusBadRequest, "record not found!")
//		return
//	}
//	util.SuccessJSON(c, http.StatusOK, "OK!", post)
//}
//
//func UpdatePost(c *gin.Context) {
//	var post models.Post
//	if err := database.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
//		util.ErrorJSON(c, http.StatusBadRequest, "record not found!")
//		return
//	}
//	// Decode request data
//	if err := c.ShouldBind(&post); err != nil {
//		util.ErrorJSON(c, http.StatusBadRequest, err.Error())
//		return
//	}
//	// Validate error
//	if err := post.Validate(); err != nil {
//		util.ErrorJSON(c, http.StatusBadRequest, err.Error())
//		return
//	}
//	// update post
//	postData := models.Post{Name: post.Name, Desc: post.Desc}
//	database.DB.Model(&post).Updates(postData)
//	util.ResponseSuccess(c, http.StatusOK, "Updated Post Successfully")
//}
//
//func DeletePost(c *gin.Context) {
//	// Get model if exist
//	var post models.Post
//	if err := database.DB.First(&post, c.Param("id")).Error; err != nil {
//		util.ErrorJSON(c, http.StatusBadRequest, "record not found!")
//		return
//	}
//	database.DB.Delete(&post)
//	util.ResponseSuccess(c, http.StatusOK, "Deleted Post Successfully")
//}
