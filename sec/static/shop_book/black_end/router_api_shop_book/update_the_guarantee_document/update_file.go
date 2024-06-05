package updatetheguaranteedocument

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func UploadImage(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		log.Println("Error in uploading Image: ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	// Use the original file name
	image := file.Filename

	// Ensure the directory exists
	dir := `D:\TEST_golang\testvar\sec\static\shop_book\black_end\router_api_shop_book\update_the_guarantee_document\images`
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			log.Println("Error creating directory:", err)
			return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
		}
	}

	// Save the file
	filePath := filepath.Join(dir, image)
	err = c.SaveFile(file, filePath)
	if err != nil {
		log.Println("Error in saving Image:", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	imageUrl := fmt.Sprintf("http://localhost:8080/images/%s", image)
	data := map[string]interface{}{
		"imageName": image,
		"imageUrl":  imageUrl,
		"header":    file.Header,
		"size":      file.Size,
	}

	return c.JSON(fiber.Map{"status": 201, "message": "Image uploaded successfully", "data": data})
}

func ViewImage(c *fiber.Ctx) error {
	imageName := c.Params("imageName")
	imageUrl := fmt.Sprintf("http://localhost:8080/images/%s", imageName)

	tmpl, err := template.ParseFiles("view.html")
	if err != nil {
		log.Println("Error parsing template:", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	data := map[string]string{
		"ImageUrl": imageUrl,
	}

	return tmpl.Execute(c.Response().BodyWriter(), data)
}
