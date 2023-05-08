package pkg

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dnevsky/hmm-bot/models"
)

const (
	ChatGPTHTTPAddress = "https://api.openai.com/v1/chat/completions"
)

type OpenAI struct {
	convs []models.Converse
	token string
}

func NewOpenAI(token string) *OpenAI {
	return &OpenAI{
		convs: make([]models.Converse, 0),
		token: token,
	}
}

// model - модель данных, role - роль, которую будет выполнять программа. Нужно это для того, что-бы сохранять историю запросов
// конкретно для разных ролей. /joke (придумать шутку) и /gpt4 (обычный чат с gpt-4) вдвоем используют модель gpt-4,
// но очевидно что история запросов у них будет разная, и очищать историю поиска тоже нужно для них по разному
func (o *OpenAI) ConverseWithOpenAI(from int, prompt string, system string, model string, role string) (models.ChatCompletionResponse, error) {
	req := models.NewChatCompletionRequestBuilder(system, model, role)

	// здесь имплементирована проверка, чтобы неактуальная история запросов для определенной команды (role) очищалась спустя час
	// если разница между последним запросом и текущим временем меньше заданного (3600), то мы в только созданный массив
	// записываем из истории, если час прошел то просто начинаем новый чат
	for _, v := range o.convs {
		if v.From == from {
			if role == v.Request.Role {
				if int(time.Now().Unix())-v.Request.LastRequest < 3600 {
					req = v.Request
				}
			}
		}
	}

	req.AddToRequest(models.Message{
		Role:    "user",
		Content: prompt,
	})

	jsonValue, _ := json.Marshal(req)
	fmt.Printf("%v\n", string(jsonValue))

	re, _ := http.NewRequest("POST", ChatGPTHTTPAddress, bytes.NewBuffer(jsonValue))
	re.Header.Set("Authorization", "Bearer "+o.token)
	re.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(re)
	if err != nil {
		return models.ChatCompletionResponse{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.ChatCompletionResponse{}, errors.New(fmt.Sprintf("error while read response body\n%#v\n%#v", err, resp.Body))
	}

	if resp.StatusCode != 200 {
		return models.ChatCompletionResponse{}, errors.New(fmt.Sprintf("status != 200\n%s", body))
	}

	var chatCompletionResponse models.ChatCompletionResponse
	if err := json.Unmarshal(body, &chatCompletionResponse); err != nil {
		return models.ChatCompletionResponse{}, errors.New(fmt.Sprintf("error while unmarshall response body to json\n%#v\n%#v", err, resp.Body))
	}

	if chatCompletionResponse.Choices[0].FinishReason == "length" {
		for k, v := range o.convs {
			if v.From == from {
				o.convs = removeIndex(o.convs, k)
				return chatCompletionResponse, errors.New("length")
			}
		}
	}

	req.Messages = append(req.Messages, chatCompletionResponse.Choices[0].Message)
	req.LastRequest = int(time.Now().Unix())

	for k, v := range o.convs {
		if v.From == from {
			o.convs[k].Request = req
			return chatCompletionResponse, nil
		}
	}

	o.convs = append(o.convs, models.Converse{From: from, Request: req})

	return chatCompletionResponse, nil
}

func (o *OpenAI) ClearConverseById(from_id int) {
	for k, v := range o.convs {
		if v.From == from_id {
			o.convs = removeIndex(o.convs, k)
		}
	}
}

func removeIndex(s []models.Converse, index int) []models.Converse {
	s[index] = s[len(s)-1]
	s[len(s)-1] = models.Converse{}
	s = s[:len(s)-1]
	return s
}
