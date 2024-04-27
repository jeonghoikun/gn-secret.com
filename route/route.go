package route

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jeonghoikun/gn-secret.com/siteconfig"
)

type page struct {
	Author string

	Title       string
	Description string
	Keywords    string

	Thumbnail string

	DatePublished time.Time
	DateModified  time.Time

	Path string
	URL  string
}

type PageConfigParams struct {
	Author string

	Title       string
	Description string
	Keywords    string

	Thumbnail string

	DatePublished time.Time
	DateModified  time.Time
}

func PageConfig(c *fiber.Ctx, arg *PageConfigParams) fiber.Map {
	m := fiber.Map{}
	p := &page{}
	p.Author = arg.Author
	p.Title = arg.Title
	p.Description = arg.Description
	p.Keywords = arg.Keywords
	p.Thumbnail = arg.Thumbnail
	p.DatePublished = arg.DatePublished
	p.DateModified = arg.DateModified
	p.Path = c.Path()
	p.URL = fmt.Sprintf("%s%s", siteconfig.Host(), c.Path())
	m["Page"] = p
	return m
}
