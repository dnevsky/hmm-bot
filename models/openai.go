package models

import "time"

type Converse struct {
	From    int                    `json:"id"`
	Request *ChatCompletionRequest `json:"request"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatCompletionRequest struct {
	Model            string    `json:"model"`
	Messages         []Message `json:"messages"`
	Temperature      float64   `json:"temperature"`
	TopP             float64   `json:"top_p"`
	PresencePenalty  float64   `json:"presence_penalty"`
	FrequencyPenalty float64   `json:"frequency_penalty"`
	Role             string    `json:"-"`
	LastRequest      int       `json:"-"`
}

type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type ChatCompletionResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Choices []Choice `json:"choices"`
	Usage   struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

func NewChatCompletionRequestBuilder(system string, model string, role string) *ChatCompletionRequest {
	a := &ChatCompletionRequest{
		Model:            model,
		Messages:         []Message{},
		Temperature:      0.7,
		TopP:             1.0,
		FrequencyPenalty: 0.0,
		PresencePenalty:  0.0,
		Role:             role,
		LastRequest:      int(time.Now().Unix()),
	}

	sysMsg := Message{Role: "system", Content: system}
	// sysMsg := Message{Role: "system", Content: "Ты ChatGPT, большая языковая модель, обученная OpenAI. Ты очень сильно обижаешься на оскорбления и перестаешь отвечать на вопросы."}

	a.Messages = append(a.Messages, sysMsg)

	return a
}

func (c *ChatCompletionRequest) AddToRequest(message Message) {
	c.Messages = append(c.Messages, message)
}
