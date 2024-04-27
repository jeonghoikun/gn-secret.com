package middleware

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jeonghoikun/gn-secret.com/siteconfig"
)

func BindSiteConfig(c *fiber.Ctx) error {
	m := fiber.Map{}
	m["Site"] = siteconfig.Config
	m["TimeNow"] = time.Now()
	if err := c.Bind(m); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.Next()
}
