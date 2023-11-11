package handler

import "github.com/gofiber/fiber/v2"

func HealthCheck(c *fiber.Ctx) error {
	type HealthRes struct {
		Status string `json:"status"`
	}

	res := HealthRes{
		Status: "alive",
	}
	return c.JSON(res)
}
