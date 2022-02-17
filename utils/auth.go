package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"strconv"
	"time"
)

//GenerateNewAccessToken func for generate a new Access token.
func GenerateNewAccessToken(id int) string {

	secret := os.Getenv("JWT_SECRET_KEY")

	exp := time.Now().Add(time.Minute * 30).Unix()

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	userId := strconv.Itoa(id)

	claims["user_id"] = userId

	claims["exp"] = exp

	t, err := token.SignedString([]byte(secret))

	if err != nil {
		return err.Error()
	}

	return t
}

func GetUserID(c *fiber.Ctx) int {
	jwtUser := c.Locals("user").(*jwt.Token)

	claims := jwtUser.Claims.(jwt.MapClaims)

	id := claims["user_id"].(string)

	userId, _ := strconv.Atoi(id)

	return userId
}
