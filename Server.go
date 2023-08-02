package main

import (
	"html/template"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

var events []string = make([]string, 0)

func main() {
	engine := html.New("./", ".html")
	engine.AddFunc(
		"unescape", func(s string) template.HTML {
			return template.HTML(s)
		},
	)
	engine.Reload(true)
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use("/wsrecord", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	app.Use("/wsplay", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	app.Get("/example1record", func(c *fiber.Ctx) error {
		return c.Render("html-example/record", fiber.Map{})
	})
	app.Get("/example1play", func(c *fiber.Ctx) error {
		return c.Render("html-example/play", fiber.Map{})
	})
	app.Get("/wsrecord", websocket.New(func(c *websocket.Conn) {
		var (
			msg []byte
			err error
		)
		for {
			if _, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}
			events = append(events, string(msg))
		}
	}))
	app.Get("/wsplay", websocket.New(func(c *websocket.Conn) {
		for _, e := range events {
			c.WriteMessage(websocket.TextMessage, []byte(e))
		}
	}))
	log.Fatal(app.Listen(":80"))
}
