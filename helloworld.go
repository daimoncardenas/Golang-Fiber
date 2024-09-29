package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"hello-golang/utils"
)

func main() {
    app := fiber.New()

    // Ruta GET para verificar que el servidor está funcionando
    app.Get("/ping", func(c *fiber.Ctx) error {
		fmt.Println("Request Body:", string(c.Body()))
        return c.JSON(fiber.Map{
            "message": "pong",
        })
    })

    // Ruta POST que recibe datos en el cuerpo de la solicitud
    app.Post("/data", func(c *fiber.Ctx) error {
        // Imprimir el cuerpo de la solicitud
        fmt.Println("Request Body:", string(c.Body()))

        // Definir la estructura para los datos esperados
        type Data struct {
            Name string `json:"name"`
            Age  int    `json:"age"`
        }

        // Crear una instancia de la estructura Data
        data := new(Data)

        // Parsear el cuerpo de la solicitud al struct data
        if err := c.BodyParser(data); err != nil {
            // Si hay un error, devolver un mensaje de error al cliente
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "Invalid request body",
            })
        }

        // Devolver los datos recibidos como respuesta
        return c.JSON(data)
    })

     // Llamar a la función Saludar del paquete utils
     utils.Saludar()

    // Escuchar en el puerto 8080
    app.Listen(":8080")
}
