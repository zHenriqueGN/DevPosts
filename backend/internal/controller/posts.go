package controller

import (
	"api/internal/auth"
	"api/internal/database"
	"api/internal/models"
	"api/internal/repositories"

	"github.com/gofiber/fiber/v2"
)

// CreatePost create a user in database
func CreatePost(c *fiber.Ctx) error {
	authorization := c.GetReqHeaders()["Authorization"]
	tokenUserID, err := auth.GetTokenUserID(authorization)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var post models.Post
	err = c.BodyParser(&post)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	err = post.Prepare()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	post.AuthorID = tokenUserID

	db, err := database.ConnectToDB()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	postID, err := repository.Create(post)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	post.ID = postID

	return c.Status(fiber.StatusCreated).JSON(post)
}

// FetchPosts fetch all the posts in database
func FetchPosts(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{"error": "method not implemented yet"})
}

// FetchPost fetch a post in database
func FetchPost(c *fiber.Ctx) error {
	ID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	db, err := database.ConnectToDB()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	post, err := repository.GetById(ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if post.ID != ID {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(post)
}

// UpdatePost update a post in database
func UpdatePost(c *fiber.Ctx) error {
	ID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	authorization := c.GetReqHeaders()["Authorization"]

	tokenUserID, err := auth.GetTokenUserID(authorization)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var post models.Post
	err = c.BodyParser(&post)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	post.ID = ID

	db, err := database.ConnectToDB()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	tempPost, err := repository.GetById(ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if tempPost.ID != ID {
		return c.SendStatus(fiber.StatusNotFound)
	}

	if tempPost.AuthorID != tokenUserID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "You can't update someone else's post"})
	}

	err = repository.Update(post)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// DeletePost delete a post from database
func DeletePost(c *fiber.Ctx) error {
	ID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	authorization := c.GetReqHeaders()["Authorization"]

	tokenUserID, err := auth.GetTokenUserID(authorization)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	db, err := database.ConnectToDB()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	tempPost, err := repository.GetById(ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if tempPost.ID != ID {
		return c.SendStatus(fiber.StatusNotFound)
	}

	if tempPost.AuthorID != tokenUserID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "You can't delete someone else's post"})
	}

	err = repository.Delete(ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
