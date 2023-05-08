package vkapi

import (
	"math/rand"
	"time"

	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/sirupsen/logrus"
)

func sendMessage(msg string, peer int, mention bool, attach string) {
	p := params.NewMessagesSendBuilder()
	rand.Seed(time.Now().UnixNano())

	p.Message(msg)
	p.RandomID(rand.Int())
	p.PeerID(peer)
	p.DisableMentions(mention)
	if attach != "" {
		p.Attachment(attach)
	}

	_, err := vk.MessagesSend(p.Params)

	if err != nil {
		logrus.Printf("error while send message: %s", err.Error())
	}
}

func sendMessageError(err error, peer int) {
	sendMessage("Во время выполнения запроса возникла ошибка.\n"+err.Error(), peer, false, "")
	logrus.Println(err.Error())
}
