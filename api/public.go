package api

import "onebot/messages"

func (a *ActionStruct) SendGroupMsg(groupId int64, autoEscape bool, msgs ...messages.Message) IDo {
	a.action = "send_group_msg"
	params := make(map[string]interface{})
	message := make([]messages.Message, 0)

	params["group_id"] = groupId
	params["message"] = append(message, msgs...)
	a.params = params
	return a
}

func (a *ActionStruct) SendPrivateMsg(userId int64, autoEscape bool, msgs ...messages.Message) IDo {
	a.action = "send_private_msg"
	params := make(map[string]interface{})
	message := make([]messages.Message, 0)

	params["user_id"] = userId
	params["message"] = append(message, msgs...)
	a.params = params
	return a
}
