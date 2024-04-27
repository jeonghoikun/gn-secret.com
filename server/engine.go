package server

import (
	"github.com/dustin/go-humanize"
	"github.com/gofiber/template/html"
)

func comma(n int) string { return humanize.Comma(int64(n)) }

func engine() *html.Engine {
	e := html.New("./views", ".html")
	e.Reload(true)
	e.AddFunc("Comma", comma)
	return e
}
