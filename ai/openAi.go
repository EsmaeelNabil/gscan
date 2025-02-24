package ai

import (
	"context"
	"fmt"

	"github.com/openai/openai-go"
)

type OpenAiAgent struct{}

func (a OpenAiAgent) Process(prompt, model, systemMessage string) string {
	return processOpenAiCompletion(prompt, systemMessage)
}

func processOpenAiCompletion(prompt, systemMessage string) string {
	client := openai.NewClient()

	completion, error := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
			openai.SystemMessage(systemMessage),
		}),
		Model: openai.F(openai.ChatModelGPT4oMini),
	})

	if error != nil {
		return fmt.Sprintf("Couldn't do the completion! but here is your input! : %s", prompt)
	}

	return completion.Choices[0].Message.Content
}
