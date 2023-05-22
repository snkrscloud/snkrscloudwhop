package webhook

import "github.com/gofiber/fiber/v2"

func AddWebhookRoutes(app *fiber.App, route string, controller *WebhookController) {
	app.Post(route, controller.handleWebhook)
}
