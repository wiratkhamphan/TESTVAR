package routerapishopbook

import (
	apishopbook "testvar/sec/static/shop_book/black_end/router_api_shop_book/api_shop_book"
	updatetheguaranteedocument "testvar/sec/static/shop_book/black_end/router_api_shop_book/update_the_guarantee_document"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter_shop_book(app *fiber.App) {

	app.Get("/Get_book", apishopbook.Get_book)
	app.Post("/appbook", apishopbook.Get_book)

	app.Post("/post_upfile", updatetheguaranteedocument.UploadImage)
	app.Get("/view/:imageName", updatetheguaranteedocument.ViewImage)
}
