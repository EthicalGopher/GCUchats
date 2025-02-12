package main

import (

	"genai/template"
	"github.com/gofiber/template/html/v2"
	"github.com/gofiber/fiber/v2"

)





func main(){
	engine := html.New("./frontend", ".html")


app:=fiber.New(fiber.Config{
	Views: engine,
})


	app.Static("/","./frontend")
	app.Post("/send-message", template.Botresponse)
	defer app.Listen(":8080")
}