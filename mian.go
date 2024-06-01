package main

import (
	"database/sql"
	"log"
	"testvar/sec/data"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Product struct {
	ID       int     `json:"id"`
	Price    float64 `json:"price"`
	VAT      float64 `json:"vat"`
	NetPrice float64 `json:"netPrice"`
}

var db *sql.DB

func calculateVAT(price float64) (float64, float64) {
	vat := price * 7 / 100
	netPrice := price + vat
	return vat, netPrice
}

func getProducts(c *fiber.Ctx) error {
	rows, err := db.Query("SELECT id, price, vat, netprice FROM vat_alculator")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()

	products := []Product{}
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Price, &p.VAT, &p.NetPrice); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		products = append(products, p)
	}

	return c.JSON(products)
}

func upvat(c *fiber.Ctx) error {
	var p Product
	if err := c.BodyParser(&p); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	p.VAT, p.NetPrice = calculateVAT(p.Price)

	result, err := db.Exec("INSERT INTO vat_alculator (id, price, vat, netprice) VALUES (?, ?, ?, ?)", p.ID, p.Price, p.VAT, p.NetPrice)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	p.ID = int(id)

	return c.JSON(p)
}

func main() {
	var err error
	db, err = data.DBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST",
	}))

	app.Static("/", "./sec/static")
	app.Static("/up", "./sec/static/upvat")

	app.Get("/api/products", getProducts)
	app.Post("/api/products1", upvat)

	log.Fatal(app.Listen("0.0.0.0:8080"))
}
