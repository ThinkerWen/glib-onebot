package api

import (
	"onebot/events"
	"onebot/messages"
)

func (a *ActionStruct) SendGroupMsg(groupId int64, autoEscape bool, msgs ...messages.Message) IDo {
	a.action = "send_group_msg"
	params := make(map[string]interface{})

	params["group_id"] = groupId
	params["message"] = msgs
	params["auto_escape"] = autoEscape
	a.params = params
	return a
}

func (a *ActionStruct) SendPrivateMsg(userId int64, autoEscape bool, msgs ...messages.Message) IDo {
	a.action = "send_private_msg"
	params := make(map[string]interface{})

	params["user_id"] = userId
	params["message"] = msgs
	params["auto_escape"] = autoEscape
	a.params = params
	return a
}

func (a *ActionStruct) SendMsg(messageType string, userId, groupId int64, autoEscape bool, msgs ...messages.Message) IDo {
	a.action = "send_msg"
	params := make(map[string]interface{})

	params["message_type"] = messageType
	params["user_id"] = userId
	params["group_id"] = groupId
	params["message"] = msgs
	params["auto_escape"] = autoEscape
	a.params = params
	return a
}

func (a *ActionStruct) GetMsg(messageId string) IDo {
	a.action = "get_msg"
	params := make(map[string]interface{})

	params["message_id"] = messageId
	a.params = params
	return a
}

func (a *ActionStruct) GetForwardMsg(id string) IDo {
	a.action = "get_forward_msg"
	params := make(map[string]interface{})

	params["id"] = id
	a.params = params
	return a
}

func (a *ActionStruct) SendLike(userId int64, times int) IDo {
	a.action = "send_like"
	params := make(map[string]interface{})

	params["user_id"] = userId
	params["times"] = times
	a.params = params
	return a
}

func (a *ActionStruct) SetGroupKick(groupId, userId int64, rejectAddRequest bool) IDo {
	a.action = "set_group_kick"
	params := make(map[string]interface{})

	params["group_id"] = groupId
	params["user_id"] = userId
	params["reject_add_request"] = rejectAddRequest
	a.params = params
	return a
}

func (a *ActionStruct) SetGroupBan(groupId, userId int64, duration int) IDo {
	a.action = "set_group_ban"
	params := make(map[string]interface{})

	params["group_id"] = groupId
	params["user_id"] = userId
	params["duration"] = duration
	a.params = params
	return a
}

func (a *ActionStruct) SetGroupAnonymousBan(groupId int64, anonymous events.Anonymous, flag string, duration int) IDo {
	a.action = "set_group_anonymous_ban"
	params := make(map[string]interface{})

	params["group_id"] = groupId
	params["anonymous"] = anonymous
	params["flag"] = flag
	params["duration"] = duration
	a.params = params
	return a
}

func (a *ActionStruct) SetGroupWholeBan(groupId int64, enable bool) IDo {
	a.action = "set_group_whole_ban"
	params := make(map[string]interface{})

	params["group_id"] = groupId
	params["enable"] = enable
	a.params = params
	return a
}

func (a *ActionStruct) SetGroupAdmin(groupId, userId int64, enable bool) IDo {
	a.action = "set_group_admin"
	params := make(map[string]interface{})

	params["group_id"] = groupId
	params["user_id"] = userId
	params["enable"] = enable
	a.params = params
	return a
}

func (a *ActionStruct) SetGroupAnonymous(groupId, enable bool) IDo {
	a.action = "set_group_anonymous"
	params := make(map[string]interface{})

	params["group_id"] = groupId
	params["enable"] = enable
	a.params = params
	return a
}

func (a *ActionStruct) SetGroupCard(groupId, userId int64, card string) IDo {
	a.action = "set_group_card"
	params := make(map[string]interface{})

	params["group_id"] = groupId
	params["user_id"] = userId
	params["card"] = card
	a.params = params
	return a
}

func (a *ActionStruct) SetGroupName(groupId int64, groupName string) IDo {
	a.action = "set_group_name"
	params := make(map[string]interface{})

	params["group_id"] = groupId
	params["group_name"] = groupName
	a.params = params
	return a
}

func (a *ActionStruct) SetGroupLeave(groupId int64, isDismiss bool) IDo {
	a.action = "set_group_leave"
	params := make(map[string]interface{})

	params["group_id"] = groupId
	params["is_dismiss"] = isDismiss
	a.params = params
	return a
}

func (a *ActionStruct) SetGroupSpecialTitle(groupId, userId int64, specialTitle string, duration int) IDo {
	a.action = "set_group_special_title"
	params := make(map[string]interface{})

	params["group_id"] = groupId
	params["user_id"] = userId
	params["special_title"] = specialTitle
	params["duration"] = duration
	a.params = params
	return a
}

func (a *ActionStruct) SetFriendAddRequest(flag string, approve bool, remark string) IDo {
	a.action = "set_friend_add_request"
	params := make(map[string]interface{})

	params["flag"] = flag
	params["approve"] = approve
	params["remark"] = remark
	a.params = params
	return a
}

func (a *ActionStruct) SetGroupAddRequest(flag, subType string, approve bool, reason string) IDo {
	a.action = "set_group_add_request"
	params := make(map[string]interface{})

	params["flag"] = flag
	params["sub_type"] = subType
	params["approve"] = approve
	params["reason"] = reason
	a.params = params
	return a
}

func (a *ActionStruct) GetLoginInfo() IDo {
	a.action = "get_login_info"
	params := make(map[string]interface{})

	a.params = params
	return a
}

func (a *ActionStruct) GetStrangerInfo(userId int64, noCache bool) IDo {
	a.action = "get_stranger_info"
	params := make(map[string]interface{})

	params["user_id"] = userId
	params["no_cache"] = noCache
	a.params = params
	return a
}

func (a *ActionStruct) GetFriendList() IDo {
	a.action = "get_friend_list"
	params := make(map[string]interface{})

	a.params = params
	return a
}

func (a *ActionStruct) GetGroupInfo(groupId int64, noCache bool) IDo {
	a.action = "get_group_info"
	params := make(map[string]interface{})

	params["group_id"] = groupId
	params["no_cache"] = noCache
	a.params = params
	return a
}

func (a *ActionStruct) GetGroupList() IDo {
	a.action = "get_group_list"
	params := make(map[string]interface{})

	a.params = params
	return a
}

func (a *ActionStruct) GetGroupMemberInfo(groupId, userId int64, noCache bool) IDo {
	a.action = "get_group_member_info"
	params := make(map[string]interface{})

	params["group_id"] = groupId
	params["user_id"] = userId
	params["no_cache"] = noCache
	a.params = params
	return a
}

func (a *ActionStruct) GetGroupMemberList(groupId int64) IDo {
	a.action = "get_group_member_list"
	params := make(map[string]interface{})

	params["group_id"] = groupId
	a.params = params
	return a
}

func (a *ActionStruct) GetGroupHonorInfo(groupId int64, honorType string) IDo {
	a.action = "get_group_honor_info"
	params := make(map[string]interface{})

	params["group_id"] = groupId
	params["type"] = honorType
	a.params = params
	return a
}

func (a *ActionStruct) GetCookies(domain string) IDo {
	a.action = "get_cookies"
	params := make(map[string]interface{})

	params["domain"] = domain
	a.params = params
	return a
}

func (a *ActionStruct) GetCsrfToken(token int) IDo {
	a.action = "get_csrf_token"
	params := make(map[string]interface{})

	params["token"] = token
	a.params = params
	return a
}

func (a *ActionStruct) GetCredentials(domain string) IDo {
	a.action = "get_credentials"
	params := make(map[string]interface{})

	params["domain"] = domain
	a.params = params
	return a
}

func (a *ActionStruct) GetRecord(file, outFormat string) IDo {
	a.action = "get_record"
	params := make(map[string]interface{})

	params["file"] = file
	params["out_format"] = outFormat
	a.params = params
	return a
}

func (a *ActionStruct) GetImage(file string) IDo {
	a.action = "get_image"
	params := make(map[string]interface{})

	params["file"] = file
	a.params = params
	return a
}

func (a *ActionStruct) CanSendImage() IDo {
	a.action = "can_send_image"
	params := make(map[string]interface{})

	a.params = params
	return a
}

func (a *ActionStruct) CanSendRecord() IDo {
	a.action = "can_send_record"
	params := make(map[string]interface{})

	a.params = params
	return a
}

func (a *ActionStruct) GetStatus() IDo {
	a.action = "get_status"
	params := make(map[string]interface{})

	a.params = params
	return a
}

func (a *ActionStruct) GetVersionInfo() IDo {
	a.action = "get_version_info"
	params := make(map[string]interface{})

	a.params = params
	return a
}

func (a *ActionStruct) SetRestart(delay int) IDo {
	a.action = "set_restart"
	params := make(map[string]interface{})

	params["delay"] = delay
	a.params = params
	return a
}
