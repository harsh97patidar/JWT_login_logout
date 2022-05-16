package service

import (
	"JWT_auth/go-auth/database"
	"JWT_auth/go-auth/models"
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user := models.User{
		Name:     data["name"],
		Password: data["password"],
		Email:    data["email"],
	}

	item := `INSERT INTO user_data ("Id", "Name", "Email", "Password") VALUES ($1, $2, $3, $4)`

	res, err := database.DB.Exec(item, 2, user.Name, user.Email, user.Password)

	if err != nil {
		panic(err)
	}

	numDeleted, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	print(numDeleted)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	query := `SELECT "Email","Id" FROM user_data WHERE "Email"=$1;`

	row := database.DB.QueryRow(query, data["email"])

	var Email string
	var Id int

	switch err := row.Scan(&Email, &Id); err {
	case sql.ErrNoRows:
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	case nil:
		return c.JSON(fiber.Map{
			"message": "success",
			"id":      Id,
		})
	default:
		panic(err)
	}
}

func Logout(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func User(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	id := c.Params("id")

	query := `SELECT * FROM user_data WHERE "Id"=$1;`

	row := database.DB.QueryRow(query, id)

	var user models.User

	switch err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password); err {
	case sql.ErrNoRows:
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	case nil:
		return c.JSON(user)
	default:
		panic(err)
	}
}
