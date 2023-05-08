package vkapi

import (
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/longpoll-bot"
)

var vk *api.VK

func NewVK(token string) (*api.VK, error) {
	// init vk
	vkApi := api.NewVK(token)

	// check work
	_, err := vkApi.GroupsGetByID(nil)
	if err != nil {
		return nil, err
	}

	vk = vkApi

	return vkApi, nil
}

func InitLongPool(vk *api.VK) (*longpoll.LongPoll, error) {
	group, err := vk.GroupsGetByID(nil)
	if err != nil {
		return nil, err
	}

	lp, err := longpoll.NewLongPoll(vk, group[0].ID)
	if err != nil {
		return nil, err
	}

	lp.Goroutine(true)

	return lp, nil
}

func getUsersInfo(users []string) (api.UsersGetResponse, error) {
	p := params.NewUsersGetBuilder()

	p.UserIDs(users)
	p.Fields([]string{"sex"})
	res, err := vk.UsersGet(p.Params)

	return res, err
}
