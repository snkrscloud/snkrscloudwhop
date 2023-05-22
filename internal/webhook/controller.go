package webhook

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type WebhookController struct {
	storage *WebhookStorage
}

func NewWebhookController(storage *WebhookStorage) *WebhookController {
	return &WebhookController{
		storage: storage,
	}
}

type WebhookEvent struct {
	Action string `json:"action"`
	Data   struct {
		ID   string `json:"id"`
		User struct {
			ID            string `json:"id"`
			Username      string `json:"username"`
			Email         string `json:"email"`
			ProfilePicURL string `json:"profile_pic_url"`
		} `json:"user"`
		Status string `json:"status"`
		Valid  bool   `json:"valid"`
	} `json:"data"`
}

func (t *WebhookController) handleWebhook(c *fiber.Ctx) error {
	// parse the request body
	var webhook WebhookEvent
	if err := c.BodyParser(&webhook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	fmt.Println("New webhook event: ", webhook.Action)

	data := webhook.Data

	err := t.storage.updateUser(data.User.ID, data.User.Username, data.User.Email, data.User.ProfilePicURL, data.Valid, c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update User",
		})
	}

	return c.SendStatus(fiber.StatusOK)
}
