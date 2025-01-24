package adapters

import (
	entities "chater/internal/domain/entity"
	"context"
	"fmt"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type OpenAIAdapter struct {
	client      *openai.Client
	modelName   string
	basePrompt  string
	chatHistory []openai.ChatCompletionMessage
}

// Поддерживаемые модели OpenAI
var supportedModels = map[string]bool{
	openai.ChatModelGPT4oMini: true,
	openai.ChatModelGPT4o:     true,
}

func getSupportedModels() []string {
	models := make([]string, 0, len(supportedModels))
	for model := range supportedModels {
		models = append(models, model)
	}
	return models
}

// NewOpenAIAdapter создаёт адаптер для работы с OpenAI
func NewOpenAIAdapter(apiKey, modelName, basePrompt string) (entities.MLModelApi, error) {
	if !supportedModels[modelName] {
		return nil, fmt.Errorf("unsupported model: %s. Supported models are: %v", modelName, getSupportedModels())
	}
	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)
	return &OpenAIAdapter{
		client:      client,
		modelName:   modelName,
		basePrompt:  basePrompt,
		chatHistory: []openai.ChatCompletionMessage{},
	}, nil
}

// SendRequest отправляет запрос к OpenAI с учётом истории диалога
func (o *OpenAIAdapter) SendRequest(input string) (string, error) {
	// Если история пустая, добавляем системный промт
	if len(o.chatHistory) == 0 && o.basePrompt != "" {
		o.chatHistory = append(o.chatHistory, openai.ChatCompletionMessage{
			Role:    openai.ChatCompletionMessageRole(openai.ChatCompletionMessageParamRoleSystem),
			Content: o.basePrompt,
		})
	}

	// Добавляем сообщение пользователя в историю
	o.chatHistory = append(o.chatHistory, openai.ChatCompletionMessage{
		Role:    openai.ChatCompletionMessageRole(openai.ChatCompletionMessageParamRoleUser),
		Content: input,
	})

	// Формируем запрос для чата
	req := openai.{
		Model:    openai.GPT3Dot5Turbo, // Вы можете заменить на "gpt-4", если это требуется
		Messages: o.chatHistory,
	}

	// Отправляем запрос к OpenAI
	resp, err := o.client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", fmt.Errorf("failed to get OpenAI response: %w", err)
	}

	// Проверяем, есть ли ответ
	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response from OpenAI")
	}

	// Получаем ответ ассистента
	assistantMessage := resp.Choices[0].Message

	// Добавляем сообщение ассистента в историю
	o.chatHistory = append(o.chatHistory, assistantMessage)

	return assistantMessage.Content, nil
}
