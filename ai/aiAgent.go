package ai

type AiAgent interface {
	Process(prompt, model, systemMessage string) string
}

type LocalAgent struct{}
