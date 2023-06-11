/*
Copyright © 2023 HAIDER ALI <haider.lee23@gmail.com>
*/
package ai

import (
	"context"
	"fmt"
	"os"

	"github.com/franciscoescher/goopenai"
)

func Connect() *goopenai.Client {
	apiKey := os.Getenv("API_KEY")
	organization := os.Getenv("API_ORG")

	if apiKey == "" {
		fmt.Println(
			`No API_KEY found!
Create a new API Key -> https://platform.openai.com/account/api-keys`,
		)
		os.Exit(1)
	}

	return goopenai.NewClient(apiKey, organization)
}

func PromptResult(client *goopenai.Client, prompt string) string {
	r := goopenai.CreateCompletionsRequest{
		Model: "gpt-3.5-turbo",
		Messages: []goopenai.Message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
		Temperature: 0.7,
	}

	completions, err := client.CreateCompletions(context.Background(), r)
	if err != nil {
		fmt.Println("AI failure: " + err.Error())
		os.Exit(1)
	}

	return completions.Choices[0].Message.Content
}
