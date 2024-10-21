package controllers

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kingomez21/app_task_deno_go/database"
	"github.com/kingomez21/app_task_deno_go/models"
)

func GetUsers(c *fiber.Ctx) error {

	var users []models.User

	rows, err := database.DB.Query("SELECT id, firstname, lastname, email, address FROM user")

	if err != nil {
		log.Println("Error ejecutando consulta:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error al obtener usuarios")
	}
	defer rows.Close()

	for rows.Next() {
		var ID, firstname, lastname, email, address string
		if err := rows.Scan(&ID, &firstname, &lastname, &email, &address); err != nil {
			log.Println("Error escaneando filas:", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Error al procesar usuarios")
		}
		users = append(users, models.User{ID: ID, Firstname: firstname, Lastname: lastname, Email: email, Address: address})
	}

	return c.JSON(users)
}

func GetOneUser(c *fiber.Ctx) error {

	var user models.User
	id := c.Params("id")

	result, err := database.DB.Query("SELECT id, firstname, lastname, email, address FROM user WHERE id = ?", id)

	if err != nil {
		return err
	}

	defer result.Close()

	for result.Next() {
		var ID, firstname, lastname, email, address string
		if err := result.Scan(&ID, &firstname, &lastname, &email, &address); err != nil {
			log.Println("Error escaneando filas:", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Error al procesar usuarios")
		}
		user = models.User{ID: ID, Firstname: firstname, Lastname: lastname, Email: email, Address: address}
	}

	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {

	formData := new(models.User)
	var newUserId int64
	var errinsert error

	err := c.BodyParser(formData)

	if err != nil {
		return err
	}

	insert := "INSERT INTO user (firstname, lastname, email, address) VALUES (?,?,?,?)"

	result, errsql := database.DB.Exec(insert, formData.Firstname, formData.Lastname, formData.Email, formData.Address)

	if errsql != nil {
		return errsql
	}

	newUserId, errinsert = result.LastInsertId()

	if errinsert != nil {
		return errinsert
	}

	return c.JSON(fiber.Map{
		"msg":    "Usuario " + strconv.FormatInt(newUserId, 10) + " creado satisfactoriamente",
		"status": true,
	})

}

func UpdatedUser(c *fiber.Ctx) error {

	id := c.Params("id")
	formData := new(models.User)
	query := "UPDATE user SET firstname=?, lastname=?, email=?, address=? WHERE id = ?"

	if err := c.BodyParser(formData); err != nil {
		return err
	}

	_, err := database.DB.Exec(query, formData.Firstname, formData.Lastname, formData.Email, formData.Address, id)

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"msg": "Usuario actualizado con exito",
	})
}

func DeleteUser(c *fiber.Ctx) error {

	id := c.Params("id")

	query := "DELETE FROM user WHERE id = ?"

	_, err := database.DB.Exec(query, id)

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"msg": "Usuario eliminado exitosamente",
	})

}
