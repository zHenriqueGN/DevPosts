package controller

import (
	"api/internal/auth"
	"api/internal/database"
	"api/internal/models"
	"api/internal/repositories"
	"api/internal/security"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// CreateUser create a user in database
func CreateUser(c *fiber.Ctx) error {
	var user models.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	if err = user.Prepare("register"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	db, err := database.ConnectToDB()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	ID, err := repository.Create(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	user.ID = ID

	return c.Status(fiber.StatusCreated).JSON(user)
}

// FetchUsers fetch all the users in database
func FetchUsers(c *fiber.Ctx) error {
	userName := c.Query("username")

	db, err := database.ConnectToDB()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	users, err := repository.FilterByUserName(userName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

// FetchUser fetch an user in database
func FetchUser(c *fiber.Ctx) error {
	ID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	db, err := database.ConnectToDB()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	user, err := repository.GetById(ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if user.ID != ID {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

// UpdateUser update an user in database
func UpdateUser(c *fiber.Ctx) error {
	ID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	authorization := c.GetReqHeaders()["Authorization"]

	tokenUserID, err := auth.GetTokenUserID(authorization)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if ID != tokenUserID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "You can't update someone else's user"})
	}

	var user models.User
	err = c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	user.ID = ID

	if err = user.Prepare("update"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	db, err := database.ConnectToDB()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)

	tempUser, err := repository.GetById(ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if tempUser.ID != ID {
		return c.SendStatus(fiber.StatusNotFound)
	}

	if err := repository.Update(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// DeleteUser delete an user from database
func DeleteUser(c *fiber.Ctx) error {
	ID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	authorization := c.GetReqHeaders()["Authorization"]

	tokenUserID, err := auth.GetTokenUserID(authorization)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if ID != tokenUserID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "You can't update someone else's user"})
	}

	db, err := database.ConnectToDB()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	tempUser, err := repository.GetById(ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if tempUser.ID != ID {
		return c.SendStatus(fiber.StatusNotFound)
	}

	err = repository.Delete(ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// FollowUser allow an user to follow another user
func FollowUser(c *fiber.Ctx) error {
	ID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	authorization := c.GetReqHeaders()["Authorization"]

	tokenUserID, err := auth.GetTokenUserID(authorization)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if ID == tokenUserID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "You can't follow yourself"})
	}

	db, err := database.ConnectToDB()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)

	tempUser, err := repository.GetById(ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if tempUser.ID != ID {
		return c.SendStatus(fiber.StatusNotFound)
	}

	err = repository.Follow(ID, tokenUserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// UnfollowUser allow an user to unfollow another user
func UnfollowUser(c *fiber.Ctx) error {
	ID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	authorization := c.GetReqHeaders()["Authorization"]

	tokenUserID, err := auth.GetTokenUserID(authorization)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if ID == tokenUserID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "You can't unfollow yourself"})
	}

	db, err := database.ConnectToDB()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)

	tempUser, err := repository.GetById(ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if tempUser.ID != ID {
		return c.SendStatus(fiber.StatusNotFound)
	}

	err = repository.Unfollow(ID, tokenUserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// GetFollowers get all the followers of a giver user
func GetFollowers(c *fiber.Ctx) error {
	ID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	db, err := database.ConnectToDB()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)

	tempUser, err := repository.GetById(ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if tempUser.ID != ID {
		return c.SendStatus(fiber.StatusNotFound)
	}

	followers, err := repository.GetFollowers(ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(followers)
}

// GetFollowings get all the followings of a giver user
func GetFollowings(c *fiber.Ctx) error {
	ID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	db, err := database.ConnectToDB()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	followings, err := repository.GetFollowings(ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if followings == nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(followings)
}

// UpdatePassword updates an user passoword
func UpdatePassword(c *fiber.Ctx) error {
	ID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	authorization := c.GetReqHeaders()["Authorization"]
	tokenUserID, err := auth.GetTokenUserID(authorization)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if ID != tokenUserID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "You can't update someone else's password"})
	}

	var passwordUpdate models.PasswordUpdate
	err = c.BodyParser(&passwordUpdate)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	if passwordUpdate.OldPassword == "" || passwordUpdate.NewPassword == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "The password can't be empty"})
	}

	db, err := database.ConnectToDB()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	dbUserPassword, err := repository.FetchPassword(ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	err = security.ComparePasswordWithHash(dbUserPassword, passwordUpdate.OldPassword)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Wrong old password"})
	}

	passwordWithHash, err := security.HashPassword(passwordUpdate.NewPassword)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err = repository.UpdatePassword(ID, string(passwordWithHash))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
