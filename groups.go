package uimserver

type GroupPostMessageParameters struct {
	UserId    string      `json:"userId"`
	ToGroupId string      `json:"toGroupId"`
	Content   *MsgContent `json:"content"`
}

type GroupPostMessageResponse struct {
	Response
	Message *Message `json:"message"`
}
