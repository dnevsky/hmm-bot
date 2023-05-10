package vkapi

import (
	"context"

	"github.com/SevereCloud/vksdk/v2/events"
)

type Message struct {
	Message      string
	Mask         string
	CheckMessage string // [regex | contains: <message> contains in MessageNew obj | equal: <message> equal MessageNew]
	Description  string
	Handler      func(context.Context, events.MessageNewObject)
	Action       string // default "", mean default message
}

func newMessage(message string, mask string, checkMessage string, description string, handler func(context.Context, events.MessageNewObject)) *Message {
	return &Message{Message: message, Mask: mask, CheckMessage: checkMessage, Description: description, Handler: handler, Action: ""}
}

func (m *Message) addAction(action string) {
	m.Action = action
}

func (h *Handler) initMessages() []Message {
	messages := make([]Message, 0)

	messages = append(messages,
		*newMessage("/help", "/help", "equal", "посмотреть все команды.", h.cmdHelp),
		*newMessage("/online", "/online", "equal", "посмотреть онлайн на серверах Diamond.", h.cmdOnline),
		*newMessage("/8ball", "/8ball <текст>", "contains", "магический шар.", h.cmdBall),
		// *newMessage("/find", "/find", "equal", "просмотр сотрудников онлайн.", h.cmdFind),
		*newMessage("/infa", "/infa <текст>", "contains", "узнать вероятность какого-то события.", h.cmdInfa),
		*newMessage(`^/kto \[id(\d+)\|.+\]$`, "/kto <упомянуть пользователя>", "regex", "узнать описание/действия/желания этого игрока.", h.cmdKto),
		*newMessage(`^/who .+`, "/who <упомянуть пользователя>", "regex", "узнать кто в беседе подходит под ваше описание.", h.cmdWho),
		*newMessage(`^/rand (-?\d+) (-?\d+)$`, "/rand <min> <max>", "regex", "получить рандомное число в диапазоне от min до max включительно", h.cmdRand),
		*newMessage(`^/rep .+`, "/rep <вопрос>", "regex", "задать вопрос в репорт (не настоящий)", h.cmdRep),
		*newMessage("/leaders", "/leaders", "equal", "лидеры по количеству сообщений.", h.cmdLeaders),
		*newMessage("/mzcoins", "/mzcoins", "equal", "меню MZ Coins.", h.cmdMZCoins),
		*newMessage("/анекдот", "/анекдот", "equal", "анекдоты =-).", h.cmdAnekdot),
		*newMessage(`^/v (.+)`, "/v <текст>", "regex", "перевод текста в голос", h.cmdTextToVoice),

		*newMessage(`^\/gpt ([\s\S]*)`, "/gpt <текст>", "regex", "общение с крутым инновационным чат-ботом от OpenAI с модельню данных gpt-3.5-turbo", h.cmdGpt),
		*newMessage(`^\/gpt4 ([\s\S]*)`, "/gpt4 <текст>", "regex", "общение с крутым инновационным чат-ботом от OpenAI с моделью данных gpt-4", h.cmdGpt4),
		*newMessage(`^\/joke ([\s\S]*)`, "/joke <текст>", "regex", "придумывает очень оригинальные и, возможно, оскорбительные штуки и обзывалки по вашему запросу. На базе gpt-4", h.cmdJoke),
		*newMessage("/gpt:reset", "/gpt:reset", "equal", "очистить историю запросов к чат-боту", h.cmdGptReset),

		*newMessage("негры", "", "contains", "", h.black),
		*newMessage("клоун", "", "contains", "", h.clown),
		*newMessage("club197623440", "", "contains", "", h.notify),
	)

	return messages
}
