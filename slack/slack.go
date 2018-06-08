package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"../healthevent"
)

type webhookAttachment struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	Color string `json:"color"`
}

type webhookArgs struct {
	Title       string              `json:"title"`
	Text        string              `json:"text"`
	Attachments []webhookAttachment `json:"attachments"`
}

// Send health event to slack
func Send(event *healthevent.EventDetail) error {
	args := generateArgs(event)
	res, err := post(args)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return err
}

func generateArgs(event *healthevent.EventDetail) *webhookArgs {
	attachment := webhookAttachment{
		Title: event.TypeCode,
		Text:  "",
		Color: "DBE511",
	}
	for _, desc := range event.Description {
		if desc.Language == "ja_JP" || desc.Language == "en_US" {
			attachment.Text = desc.Latest + "\n\n  from: " + event.StartTime + "\n  to: " + event.EndTime
			break
		}
	}
	if attachment.Text == "" {
		attachment.Text = event.Description[0].Latest + "\nfrom " + event.StartTime + " to " + event.EndTime
	}
	args := webhookArgs{
		Title:       "AWS Health Event Notification",
		Text:        "",
		Attachments: []webhookAttachment{attachment},
	}
	return &args
}

func post(args *webhookArgs) (*http.Response, error) {
	url := os.Getenv("SLACK_WEBHOOK_URL")
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(args)
	return http.Post(url, "application/json", buffer)
}
