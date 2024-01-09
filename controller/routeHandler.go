package controller

import (
	"GikBazar/database"
	"GikBazar/model"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

var db, err = database.ConnectDB()

func SetUser(c *fiber.Ctx) error {
	time.Sleep(time.Millisecond * 500)
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Added an user",
	}

	record := new(model.User)
	if err := c.BodyParser(record); err != nil {
		log.Println("Error in parsing request.")
		context["statusText"] = ""
		context["msg"] = "Something went wrong."
	}

	result := db.Create(record)

	if result.Error != nil {
		log.Println("Error in saving data.")
		context["statusText"] = ""
		context["msg"] = "Something went wrong."
	}

	context["msg"] = "Record is saved successully."
	context["data"] = record

	c.Status(201)
	return c.JSON(context)
}

func GetUsers(c *fiber.Ctx) error {
	c.Status(400)
	context := fiber.Map{
		"statusText": "",
		"msg":        "",
	}
	var record []model.User
	db.Find(&record)
	context["record"] = record
	c.Status(200)
	return c.JSON(context)

}

func GetUserById(c *fiber.Ctx) error {
	c.Status(400)
	context := fiber.Map{
		"statusText": "",
		"msg":        "",
	}
	var record model.User
	id := c.Params("id")
	if err := db.First(&record, id).Error; err != nil {
		log.Println("Error in finding the item:", err)
		context["statusText"] = ""
		context["msg"] = "Item not found."
		c.Status(404)
		return c.JSON(context)
	}
	context["data"] = record
	return c.JSON(context)
}

func GetItems(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Items List",
	}
	var records []model.Item
	db.Find(&records)
	context["shop_items"] = records
	c.Status(200)
	return c.JSON(context)
}

func PostItem(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Added an item",
	}

	record := new(model.Item)
	if err := c.BodyParser(record); err != nil {
		log.Println("Error in parsing request.")
		context["statusText"] = ""
		context["msg"] = "Something went wrong."
	}
	log.Println(record)
	result := db.Create(record)
	log.Println(result.Statement.SQL.String())
	if result.Error != nil {
		log.Println("Error in saving data.")
		context["statusText"] = ""
		context["msg"] = "Something went wrong."
	}

	context["msg"] = "Record is saved successully."
	context["data"] = record

	c.Status(201)
	return c.JSON(context)
}

func UpdateItem(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Updated an item",
	}

	// Получаем идентификатор товара из параметров запроса
	itemID := c.Params("id")
	if itemID == "" {
		context["statusText"] = ""
		context["msg"] = "Invalid item ID."
		c.Status(400)
		return c.JSON(context)
	}

	// Поиск товара по идентификатору
	var record model.Item
	if err := db.First(&record, itemID).Error; err != nil {
		log.Println("Error in finding the item:", err)
		context["statusText"] = ""
		context["msg"] = "Item not found."
		c.Status(404)
		return c.JSON(context)
	}

	// Парсинг данных запроса и обновление записи
	if err := c.BodyParser(&record); err != nil {
		log.Println("Error in parsing request.")
		context["statusText"] = ""
		context["msg"] = "Something went wrong."
		c.Status(400)
		return c.JSON(context)
	}

	// Сохранение обновленной записи
	result := db.Save(record)
	if result.Error != nil {
		log.Println("Error in updating data:", result.Error)
		context["statusText"] = ""
		context["msg"] = "Something went wrong."
		c.Status(500)
		return c.JSON(context)
	}

	context["msg"] = "Record is updated successfully."
	context["data"] = record

	c.Status(200)
	return c.JSON(context)
}

func DeleteItem(c *fiber.Ctx) error {
	c.Status(400)
	context := fiber.Map{
		"statusText": "",
		"msg":        "",
	}

	id := c.Params("id")

	var record model.Item
	db.First(&record, id)
	if record.ID == 0 {
		log.Println("Record not Found.")
		context["msg"] = "Record not Found."

		return c.JSON(context)
	}

	result := db.Delete(record)
	if result.Error != nil {
		context["msg"] = "Something went wrong."
		return c.JSON(context)
	}

	context["statusText"] = "Ok"
	context["msg"] = "Record deleted successfully."
	c.Status(200)
	return c.JSON(context)
}

func GetItemById(c *fiber.Ctx) error {
	c.Status(400)
	context := fiber.Map{
		"statusText": "",
		"msg":        "",
	}

	id := c.Params("id")

	var record model.Item
	db.First(&record, id)
	if record.ID == 0 {
		log.Println("Record not Found.")
		context["msg"] = "Record not Found."

		return c.JSON(context)
	}
	context["statusText"] = "Ok"
	context["data"] = record
	return c.JSON(context)
}

func Authorisation(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Login successful",
	}

	// Декодируем JSON из тела запроса в структуру Credentials
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&credentials); err != nil {
		context["statusText"] = ""
		context["msg"] = "Invalid request payload"
		c.Status(400)
		return c.JSON(context)
	}
	context["token"] = "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NSIsIm5hbWUiOiJOaWtpdGEiLCJhZG1pbiI6dHJ1ZX0.Zw_nC_6v0lWdCR6s6vqP29RYy652qQg2_g9RocXE92DBC4X3lEyFvj6BIVcSB0iap8-yTWuWQhLoYj9KoeKwcQ"
	return c.JSON(context)
}

func ShoperById(c *fiber.Ctx) error {
	c.Status(400)
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Login successful",
	}
	id := c.Params("id")
	var record model.Shoper
	db.First(&record, id)
	if record.ID == 0 {
		log.Println("Record not Found.")
		context["msg"] = "Record not Found."

		return c.JSON(context)
	}
	context["statusText"] = "Ok"
	context["data"] = record
	return c.JSON(context)
}

func UpdateShoper(c *fiber.Ctx) error {
	c.Status(400)
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Login successful",
	}
	id := c.Params("id")
	var record model.Shoper
	db.First(&record, id)
	if record.ID == 0 {
		log.Println("Record not Found.")
		context["msg"] = "Record not Found."

		return c.JSON(context)
	}

	if err := c.BodyParser(&record); err != nil {
		log.Println("Error in parsing request.")
		context["statusText"] = ""
		context["msg"] = "Something went wrong."
		c.Status(400)
		return c.JSON(context)
	}

	// Сохранение обновленной записи
	result := db.Save(record)
	if result.Error != nil {
		log.Println("Error in updating data:", result.Error)
		context["statusText"] = ""
		context["msg"] = "Something went wrong."
		c.Status(500)
		return c.JSON(context)
	}

	context["msg"] = "Shoper is updated successfully."
	context["data"] = record

	c.Status(200)
	return c.JSON(context)
}

func SetItemToShoper(c *fiber.Ctx) error {
	time.Sleep(time.Millisecond * 500)
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Added an shoper",
	}

	record := new(model.Shoper)
	if err := c.BodyParser(record); err != nil {
		log.Println("Error in parsing request.")
		context["statusText"] = ""
		context["msg"] = "Something went wrong."
	}

	result := db.Create(record)

	if result.Error != nil {
		log.Println("Error in saving data.")
		context["statusText"] = ""
		context["msg"] = "Something went wrong."
	}

	context["msg"] = "Record is saved successully."
	context["data"] = record

	c.Status(201)
	return c.JSON(context)
}
