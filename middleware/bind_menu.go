package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jeonghoikun/gn-secret.com/menu"
)

func BindMenu(c *fiber.Ctx) error {
	m := fiber.Map{}
	m["Whiskies"] = menu.Whiskies
	m["TCs"] = menu.TCs
	m["Foods"] = menu.Foods
	m["Menu"] = menu.Menus
	if err := c.Bind(fiber.Map{"Menu": m}); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.Next()
}
