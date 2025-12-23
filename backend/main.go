package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type Note struct {
	Baslik string `json:"baslik"`
	Icerik string `json:"icerik"`
}

func main() {
	// Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	// Fiber app
	app := fiber.New()

	// CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST",
	}))

	// Ana route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "Not Defteri Backend Çalışıyor 🚀",
		})
	})

	// Not ekleme
	app.Post("/ekle", func(c *fiber.Ctx) error {
		start := time.Now()

		var note Note
		if err := c.BodyParser(&note); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Geçersiz veri",
			})
		}

		if note.Baslik == "" || note.Icerik == "" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Başlık ve içerik boş olamaz",
			})
		}

		if err := rdb.Set(ctx, note.Baslik, note.Icerik, 0).Err(); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Not kaydedilemedi",
			})
		}

		duration := time.Since(start)
		log.Println("Not eklendi | Başlık:", note.Baslik, "| Süre:", duration)

		return c.JSON(fiber.Map{
			"message":  "Not başarıyla kaydedildi",
			"duration": duration.String(),
		})
	})

	// Not okuma
	app.Get("/oku", func(c *fiber.Ctx) error {
		start := time.Now()
		baslik := c.Query("baslik")

		val, err := rdb.Get(ctx, baslik).Result()
		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"error": "Not bulunamadı",
			})
		}

		duration := time.Since(start)
		log.Println("Not okundu | Başlık:", baslik, "| Süre:", duration)

		return c.JSON(fiber.Map{
			"icerik":   val,
			"duration": duration.String(),
		})
	})

	// Server start
	log.Fatal(app.Listen(":3000"))
}
