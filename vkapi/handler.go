package vkapi

import (
	"VKBotAPI/pkg"
	"VKBotAPI/storage"
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/SevereCloud/vksdk/v2/longpoll-bot"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	vk      *api.VK
	lp      *longpoll.LongPoll
	storage *storage.Storage
	openai  *pkg.OpenAI
}

func NewHandler(vk *api.VK, lp *longpoll.LongPoll, storage *storage.Storage, openai *pkg.OpenAI) *Handler {
	return &Handler{vk: vk, lp: lp, storage: storage, openai: openai}
}

func (h *Handler) InitHandler() {
	messages := h.initMessages()
	h.lp.MessageNew(func(c context.Context, obj events.MessageNewObject) {
		logrus.Printf("peer: %d | from: %d | action: %s | text: %s", obj.Message.PeerID, obj.Message.FromID, obj.Message.Action.Type, obj.Message.Text)

		if obj.Message.Action.Type == "chat_invite_user" {
			msg := fmt.Sprintf("Welcome to the club, [id%d|buddy]", obj.Message.Action.MemberID)

			if obj.Message.Action.MemberID == -1*h.lp.GroupID {
				msg = "Всем ку, олды на месте.\nЭто я, олд.\n\nНе забудьте дать доступ к переписке, а в идеале админку."
			}

			attachment := "video-197623440_456239017"

			sendMessage(msg, obj.Message.PeerID, false, attachment)
			return
		}

		if obj.Message.Action.Type == "chat_kick_user" {
			var msg string

			user, err := getUsersInfo([]string{fmt.Sprintf("%d", obj.Message.Action.MemberID)})
			if err == nil {
				msg = fmt.Sprintf("[id%d|%s] покинул(-а) нас.", obj.Message.Action.MemberID, user[0].FirstName)
			} else {
				msg = fmt.Sprintf("Покинул(-а) нас.\n%d", obj.Message.Action.MemberID)
			}

			sendMessage(msg, obj.Message.PeerID, false, "")
			return
		}

		for _, v := range messages {
			if v.Action == "" { // it's message
				if v.CheckMessage == "regex" {
					if regexMatchString(v.Message, obj.Message.Text) {
						re := regexp.MustCompile(v.Message)
						res := re.FindAllStringSubmatch(obj.Message.Text, -1)

						сtx := context.WithValue(c, "rgx", res[0])

						v.Handler(сtx, obj)
						return
					} else if commandContainsInString(v.Message, obj.Message.Text) {
						sendMessage(v.Mask, obj.Message.PeerID, false, "")
						return
					}

				}

				if v.CheckMessage == "contains" {
					if strings.Contains(obj.Message.Text, v.Message) {
						v.Handler(c, obj)
						return
					}

				}

				if v.CheckMessage == "equal" {
					if v.Message == obj.Message.Text {
						v.Handler(c, obj)
						return
					}
				}
				// тут сделать чтобы считало кол-во сообщений
			}
			// other actions
		}
	})
}

func (h *Handler) Run() {
	if err := h.lp.Run(); err != nil {
		logrus.Fatalf(err.Error())
	}
}

func regexMatchString(pattern string, s string) (result bool) {
	ok, _ := regexp.MatchString(pattern, s)
	return ok
}

// функция принимает в себя две строки. Внутри она делит строки в массив с разделителем по пробелу
// а потом проверяет первый элемент первого массива и первый элемент второго массива на схожесть
// /kto *kecksic - нужный шаблон
// /kto sdihsdir - то что мы передали
// равны, поэтому в обработчике мы выведем что мы ошиблись с коммандой
// /kto *kecksic - нужный шаблон
// /who *kecksic - то что мы передали
// не равны, поэтому вернем false и ничего пользователю не покажем
func commandContainsInString(pattern string, check string) bool {
	patternArr := strings.Split(pattern, " ")
	checkArr := strings.Split(check, " ")

	if len(pattern) == 0 || len(check) == 0 { // избегаем паники
		return false
	}

	if len(patternArr) == 0 || len(checkArr) == 0 { // избегаем паники
		return false
	}

	// поскольку в некоторых regex паттернах есть символ в начале строки - "^", нам нужно его убрать.
	// поэтому мы делаем массив сразу без этого символа
	if pattern[0] == '^' {
		patternArr = strings.Split(pattern[1:], " ")
	}

	if patternArr[0] == checkArr[0] {
		return true
	}

	return false
}
