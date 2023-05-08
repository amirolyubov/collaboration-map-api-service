package categories

import "github.com/gofiber/fiber/v2"

func GetCategories(c *fiber.Ctx) error {
	return c.SendString("get categories");
}
