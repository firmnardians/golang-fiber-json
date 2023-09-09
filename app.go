package main

import (
    "github.com/gofiber/fiber/v2"
) 

type SampleResponse struct {
	Field1 string `json:"message"`
	Field2 int `json:"code"`
	Field3 string `json:"status"`
}

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        tr := SampleResponse{"This is Golang API", 200, "ok"}
        return c.JSON(tr)
    })

    app.Listen(":5000")
}