package api

import "onebot/messages"

func (a *ActionStruct) SendPrivateMsg(userId int64, autoEscape bool, msgs ...messages.Message) IDo {
	a.action = "send_private_msg"
	params := make(map[string]interface{})
	message := make([]messages.Message, 0)

	params["user_id"] = userId
	params["message"] = append(message, msgs...)
	a.params = params
	return a
}
