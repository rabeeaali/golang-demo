package utils

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"math"
)

func Paginate(db *gorm.DB, entity Entity, page int) fiber.Map {
	limit := 10
	offset := (page - 1) * limit
	data := entity.Take(db, limit, offset)
	total := entity.Count(db)

	return fiber.Map{
		"data": data,
		"meta": fiber.Map{
			"page":      page,
			"total":     total,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	}
}
