package route

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jeonghoikun/gn-secret.com/siteconfig"
)

type indexHandler struct{}

func (*indexHandler) robots(c *fiber.Ctx) error {
	var ss []string
	ss = append(ss, "User-agent: *")
	ss = append(ss, "Allow: /")
	ss = append(ss, fmt.Sprintf("Sitemap: %s/sitemap.txt", siteconfig.Host()))
	return c.Status(http.StatusOK).SendString(strings.Join(ss, "\n"))
}

func (*indexHandler) sitemap(c *fiber.Ctx) error {
	var ss []string
	ss = append(ss, fmt.Sprintf("%s/", siteconfig.Host()))
	return c.Status(http.StatusOK).SendString(strings.Join(ss, "\n"))
}

func (*indexHandler) index(c *fiber.Ctx) error {
	s := siteconfig.Config
	m := PageConfig(c, &PageConfigParams{
		Author:        s.Author,
		Title:         s.Title,
		Description:   s.Description,
		Keywords:      s.Keywords,
		Thumbnail:     fmt.Sprintf("%s%s", siteconfig.Host(), s.Thumbnail),
		DatePublished: s.DatePublished,
		DateModified:  s.DateModified,
	})
	return c.Status(http.StatusOK).Render("index", m, "layout/index")
}

// BaseURL = /
func HandleIndex(r fiber.Router) {
	h := &indexHandler{}
	r.Get("/robots.txt", h.robots)
	r.Get("/sitemap.txt", h.sitemap)
	r.Get("/", h.index)
}
