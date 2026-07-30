package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/swaggo/swag"

	_ "github.com/osamah22/evently/docs"
)

// swaggerUIPage renders a minimal Swagger UI page against /swagger/doc.json.
// Fiber v3 has no official swagger-ui middleware yet (gofiber/contrib/swagger
// only supports Fiber v2), so this hand-rolls it using the swagger-ui-dist
// CDN bundle instead of pulling in an incompatible fiber/v2 dependency tree.
const swaggerUIPage = `<!DOCTYPE html>
<html>
<head>
  <title>Evently API Docs</title>
  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css" />
</head>
<body>
  <div id="swagger-ui"></div>
  <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
  <script>
    window.onload = () => {
      window.ui = SwaggerUIBundle({
        url: "/swagger/doc.json",
        dom_id: "#swagger-ui",
      });
    };
  </script>
</body>
</html>`

func registerSwagger(app *fiber.App) {
	app.Get("/swagger/doc.json", func(c fiber.Ctx) error {
		doc, err := swag.ReadDoc()
		if err != nil {
			return err
		}
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		return c.SendString(doc)
	})

	app.Get("/swagger", func(c fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return c.SendString(swaggerUIPage)
	})
}
