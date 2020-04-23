package main

import "github.com/gofiber/fiber"

func main()  {
    app := fiber.New()

    app.Settings.StrictRouting= true

    app.Static("/", "./views/static/index.html")

    app.Listen(443)
}
