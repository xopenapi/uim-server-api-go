package uimserver

import "context"

type Contact struct {
	Id            string                 `json:"id"`
	AccountId     string                 `json:"accountId"`
	UserId        string                 `json:"userId"`
	Nickname      string                 `json:"nickname"`
	Avatar        string                 `json:"avatar"`
	Gender        int32                  `json:"gender"`
	Mobile        string                 `json:"mobile"`
	Country       string                 `json:"country"`
	State         string                 `json:"state"`
	City          string                 `json:"city"`
	Signature     string                 `json:"signature"`
	Alias         string                 `json:"alias"`
	Tags          []string               `json:"tags"`
	Source        string                 `json:"source"`
	IsBlackListed bool                   `json:"isBlackListed"`
	ExtendProps   map[string]interface{} `json:"extendProps"`
	CreateAt      int64                  `json:"createAt"`
	UpdateAt      int64                  `json:"updateAt"`
}

type ContactAddParameters struct {
	AppId         string                 `json:"appId"`
	AccountId     string                 `json:"accountId"`
	UserId        string                 `json:"userId"`
	Nickname      string                 `json:"nickname"`
	Avatar        string                 `json:"avatar"`
	Gender        int32                  `json:"gender"`
	Mobile        string                 `json:"mobile"`
	Country       string                 `json:"country"`
	State         string                 `json:"state"`
	City          string                 `json:"city"`
	Signature     string                 `json:"signature"`
	Alias         string                 `json:"alias"`
	Tags          []string               `json:"tags"`
	Source        string                 `json:"source"`
	IsBlackListed bool                   `json:"isBlackListed"`
	ExtendProps   map[string]interface{} `json:"extendProps"`
}

type ContactAddResponse struct {
	Response
}

type ContactUpdateParameters struct {
	AppId         string                 `json:"appId"`
	AccountId     string                 `json:"accountId"`
	UserId        string                 `json:"userId"`
	Nickname      string                 `json:"nickname"`
	Avatar        string                 `json:"avatar"`
	Gender        int32                  `json:"gender"`
	Mobile        string                 `json:"mobile"`
	Country       string                 `json:"country"`
	State         string                 `json:"state"`
	City          string                 `json:"city"`
	Signature     string                 `json:"signature"`
	Alias         string                 `json:"alias"`
	Tags          []string               `json:"tags"`
	Source        string                 `json:"source"`
	IsBlackListed bool                   `json:"isBlackListed"`
	ExtendProps   map[string]interface{} `json:"extendProps"`
}

type ContactUpdateResponse struct {
	Response
	Contact Contact `json:"contact"`
}

func (api *Client) ContactAdd(req *ContactAddParameters) error {
	return api.ContactAddContext(context.Background(), req)
}

func (api *Client) ContactAddContext(ctx context.Context, req *ContactAddParameters) error {
	response := Response{}
	err := api.postJSONMethod(ctx, "/contact/add", req, &response)
	if err != nil {
		return err
	}
	return response.Err()
}

func (api *Client) ContactUpdate(req *ContactUpdateParameters) error {
	return api.ContactUpdateContext(context.Background(), req)
}

func (api *Client) ContactUpdateContext(ctx context.Context, req *ContactUpdateParameters) error {
	response := Response{}
	err := api.postJSONMethod(ctx, "/contact/updateFromProvider", req, &response)
	if err != nil {
		return err
	}
	return response.Err()
}
