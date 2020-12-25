package uimserver

import "context"

type ChatPostMessageParameters struct {
	AppId     string      `json:"appId"`
	Id        string      `json:"id"`
	UserId    string      `json:"userId"`
	ToUserId  string      `json:"toUserId"`
	Content   *MsgContent `json:"content"`
	IsDeleted bool        `json:"isDeleted"`
	IsRevoked bool        `json:"isRevoked"`
	SendAt    int64       `json:"sendAt"`
	RevokeAt  int64       `json:"revokeAt"`
	DeleteAt  int64       `json"deleteAt"`
	CreateAt  int64       `json:"createAt"`
	UpdateAt  int64       `json:"updateAt"`
}

type ChatPostMessageResponse struct {
	Response
	Message *Message `json:"message"`
}

func (api *Client) ChatPostMessage(req *ChatPostMessageParameters) error {
	return api.ChatPostMessageContext(context.Background(), req)
}

func (api *Client) ChatPostMessageContext(ctx context.Context, req *ChatPostMessageParameters) error {
	response := Response{}
	err := api.postJSONMethod(ctx, "/chat/postMessage", req, &response)
	if err != nil {
		return err
	}
	return response.Err()
}
