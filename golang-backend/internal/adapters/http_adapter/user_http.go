package adapters_http

import (
	"time"

	"github.com/ArmNonthakon/Minor-Cineplex/internal/core/domain"
	ports "github.com/ArmNonthakon/Minor-Cineplex/internal/core/ports/user"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type UserServiceIml struct {
	service ports.UserService
}

func NewUserService(service ports.UserService) *UserServiceIml {
	return &UserServiceIml{service: service}
}

func (h *UserServiceIml) Login(c *fiber.Ctx) error {
	input := domain.User{}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON("EROR INPUT!!")
	}
	if err := h.service.GetInputToLogin(input.UserName, input.Password); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON("UserName or Password wrong!!")
	}
	claims := jwt.MapClaims{
		"name": input.UserName,
		"role": "customer",
		"exp":  time.Now().Add(time.Hour * 8).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("MinorCineplex"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = t
	cookie.Expires = time.Now().Add(8 * time.Hour)

	c.Cookie(cookie)
	return c.Status(fiber.StatusAccepted).JSON("Login Success!!")
}

func (h *UserServiceIml) Register(c *fiber.Ctx) error {
	input := domain.User{}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON("EROR INPUT!!")
	}
	if err := h.service.GetInputToRegister(input.UserName, input.Email, input.Password); err != nil {
		return c.SendStatus(fiber.StatusNoContent)
	}
	return c.Status(fiber.StatusCreated).JSON("Register Success!!")
}
