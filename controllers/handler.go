package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"learn-chatbot/models"
	"net/http"
)

func Handler(res http.ResponseWriter, req *http.Request) {
	body := &models.WebhookReqBody{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		fmt.Println("Could not decode request body", err)
		return
	}

	if err := sayPolo(body.Message.Chat.ID); err != nil {
		fmt.Println("Error in sending reply:", err)
		return
	}
	fmt.Println("Reply sent")
}

func sayPolo(chatID int64) error {
	reqBody := &models.SendMessageReqBody{
		ChatID: chatID,
		Text:   "Polo!!",
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	res, err := http.Post("https://api.telegram.org/bot5949952682:AAENHuLEy1-RXdtfDK0W9aakkr2gUSxLiqY/sendMessage", "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("Unexpected status" + res.Status)
	}
	return nil
}
