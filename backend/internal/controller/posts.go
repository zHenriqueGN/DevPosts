package controller

import (
	"github.com/gofiber/fiber/v2"
)

// CreatePost create a user in database
func CreatePost(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "method not implemented yet"})
}

// FetchPosts fetch all the posts in database
func FetchPosts(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "method not implemented yet"})
}

// FetchPost fetch a post in database
func FetchPost(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "method not implemented yet"})
}

// UpdatePost update a post in database
func UpdatePost(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "method not implemented yet"})
}

// DeletePost delete a post from database
func DeletePost(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "method not implemented yet"})
}
