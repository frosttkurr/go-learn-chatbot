package models

type WebhookReqBody struct {
	Message Message `json:"message"`
}
