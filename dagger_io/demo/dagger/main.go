// A generated module for Demo functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/slack-go/slack"
)

type Demo struct {
	Name    string
	Message string
}

// Hello World Function
func (m *Demo) Hello(greeting string, name string) string {
	message := fmt.Sprintf("%s, %s", greeting, name)
	return message
}

// Sets The Name To Mention In Slack Message
func (m *Demo) SetName(
	ctx context.Context,
	// +optional
	name string,
) (*Demo, error) {
	if name != "" {
		m.Name = name
	} else {
		m.Name = "Generic Name"
	}
	return m, nil
}

// Sets The Message Content For A Slack Message
func (m *Demo) SetMessage(
	ctx context.Context,
	// +optional
	message string,
) (*Demo, error) {
	if message == "" {
		m.Message = "<!channel> Hi From Dagger.io and GoLang"
	} else {
		m.Message = message
	}
	return m, nil
}

// Sends A Slack Message Via A WebHook
func (m *Demo) SendSlackMessage(webhook string) {
	attachment := slack.Attachment{
		Color:         "good",
		AuthorName:    m.Name,
		AuthorSubname: "veradigm.com",
		Text:          m.Message,
		Ts:            json.Number(strconv.FormatInt(time.Now().Unix(), 10)),
	}

	msg := slack.WebhookMessage{
		Attachments: []slack.Attachment{attachment},
	}

	err := slack.PostWebhook(webhook, &msg)
	if err != nil {
		fmt.Println(err)
	}
}
