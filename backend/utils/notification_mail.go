package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sygap_new_knowledge_management/backend/email_template"

	"github.com/gofiber/fiber/v2"
)

func SendEmailNotification(notif_type string, content_list []string, c *fiber.Ctx) bool {
	if notif_type == "feedback" {
		content := email_template.Feedback(
			content_list[0], content_list[1], content_list[2], content_list[3], content_list[4], content_list[5], content_list[6],
		)
		receiver := []string{os.Getenv("DEFAULT_RECEIVER"), os.Getenv("SECOND_RECEIVER")}
		header := "Knowledge Management " + content_list[0] + " Rated By " + content_list[4]

		CurlSendEmail(content, receiver, header, c)
	} else if notif_type == "reported" {
		content := email_template.Report(
			content_list[0], content_list[1], content_list[2], content_list[3], content_list[4],
		)
		receiver := []string{os.Getenv("DEFAULT_RECEIVER"), os.Getenv("SECOND_RECEIVER")}
		header := "Knowledge Management " + content_list[0] + " Reported By " + content_list[3]

		CurlSendEmail(content, receiver, header, c)
	}
	return true
}

func CurlSendEmail(content string, receiver []string, header string, c *fiber.Ctx) bool {
	base_url := os.Getenv("URL_MAILER")
	url := base_url + "/api/v1/notification-manager/send-email/create-ticket"

	// Create the data structure to hold the JSON payload
	data := map[string]interface{}{
		"receiver":    receiver,
		"content":     content,
		"header":      header,
		"ticket_code": header,
	}

	// Marshal the map into a JSON byte slice
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return false
	}

	// Create a new POST request with the JSON payload
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return false
	}

	// Set the appropriate headers for the request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.Get("Authorization"))
	// Perform the request using the http.Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return false
	}
	defer resp.Body.Close()

	// Check the response status
	fmt.Println("Response status:", resp.Status)
	return true
}
