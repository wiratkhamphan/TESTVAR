package routerapishopbook

import (
	"database/sql"
	"log"
	data "testvar/sec/data/data_shop_book"

	"github.com/gofiber/fiber/v2"
)

var db *sql.DB

type Book struct {
	ID      int     `json:"id"`
	B_book  int     `json:"b_book"`
	B_name  string  `json:"b_name"`
	B_price float64 `json:"b_price"`
	B_file  string  `json:"b_file"`
	B_stock int     `json:"b_stock"`
}

func AppBook(c *fiber.Ctx) error {
	var B Book
	return c.JSON(B)
}

func Get_book(c *fiber.Ctx) error {
	var err error
	db, err = data.DBConnection_book()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, id_book, b_name, price, file, stock FROM app_by")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()

	Books := []Book{}
	for rows.Next() {
		var B Book

		if err := rows.Scan(
			&B.ID,
			&B.B_book,
			&B.B_name,
			&B.B_price,
			&B.B_file,
			&B.B_stock); err != nil {
			return c.Status(500).SendString(err.Error())
		}

		Books = append(Books, B)
	}
	return c.JSON(Books)
}

func SetupRouter_shop_book(app *fiber.App) {
	app.Get("/Get_book", Get_book)
	app.Post("/appbook", AppBook)
}
