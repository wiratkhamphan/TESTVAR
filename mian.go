package main

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"testvar/sec/data"
	routerapishopbook "testvar/sec/static/shop_book/black_end/router_api_shop_book"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Product struct {
	ID       int            `json:"id"`
	Price    float64        `json:"price"`
	VAT      float64        `json:"vat"`
	NetPrice float64        `json:"netPrice"`
	Img      sql.NullString `json:"img"`
}

var db *sql.DB

func calculateVAT(price float64) (float64, float64) {
	vat := price * 7 / 100
	netPrice := price + vat
	return vat, netPrice
}

func getProducts(c *fiber.Ctx) error {
	rows, err := db.Query("SELECT id, price, vat, netprice, img FROM vat_alculator")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()

	products := []map[string]interface{}{}
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Price, &p.VAT, &p.NetPrice, &p.Img); err != nil {
			return c.Status(500).SendString(err.Error())
		}

		product := map[string]interface{}{
			"id":       p.ID,
			"price":    p.Price,
			"vat":      p.VAT,
			"netPrice": p.NetPrice,
			"img":      p.Img.String,
			"imageUrl": "http://localhost:8080/images/" + p.Img.String,
		}

		products = append(products, product)
	}

	return c.JSON(products)
}

func UploadImage(c *fiber.Ctx) (string, error) {
	file, err := c.FormFile("image")
	if err != nil {
		log.Println("Error in uploading Image:", err)
		return "", err
	}

	// Use the original file name
	image := file.Filename

	// Ensure the directory exists
	dir := `D:\TEST_golang\testvar\sec\static\shop_book\black_end\router_api_shop_book\update_the_guarantee_document\images`
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			log.Println("Error creating directory:", err)
			return "", err
		}
	}

	// Save the file
	filePath := filepath.Join(dir, image)
	err = c.SaveFile(file, filePath)
	if err != nil {
		log.Println("Error in saving Image:", err)
		return "", err
	}

	return image, nil
}

func upvat(c *fiber.Ctx) error {
	// Upload the image and get the file path
	imagePath, err := UploadImage(c)
	if err != nil {
		return c.Status(500).SendString("Error uploading image: " + err.Error())
	}

	var p Product
	if err := c.BodyParser(&p); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	p.Img = sql.NullString{String: imagePath, Valid: true}
	p.VAT, p.NetPrice = calculateVAT(p.Price)

	result, err := db.Exec("INSERT INTO vat_alculator (price, vat, netprice, img) VALUES (?, ?, ?, ?)", p.Price, p.VAT, p.NetPrice, p.Img.String)
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
	app.Static("/images", "D:/TEST_golang/testvar/sec/static/shop_book/black_end/router_api_shop_book/update_the_guarantee_document/images")

	app.Get("/api/products", getProducts)
	app.Post("/api/products1", upvat)
	routerapishopbook.SetupRouter_shop_book(app)

	if err := app.Listen("0.0.0.0:8080"); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
