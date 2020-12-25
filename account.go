package uimserver

import "context"

type Account struct {
	Id          string                 `json:"id"`
	AccountId   string                 `json:"accountId"`
	CustomId    string                 `json:"customId"`
	Nickname    string                 `json:"nickname"`
	Avatar      string                 `json:"avatar"`
	Gender      int32                  `json:"gender"`
	Mobile      string                 `json:"mobile"`
	Country     string                 `json:"country"`
	State       string                 `json:"state"`
	City        string                 `json:"city"`
	Signature   string                 `json:"signature"`
	Alias       string                 `json:"alias"`
	Qrcode      string                 `json:"qrcode"`
	ExtendProps map[string]interface{} `json:"extendProps"`
	IsDeleted   bool                   `json:"isDeleted"`
	CreateAt    int64                  `json:"createAt"`
	UpdateAt    int64                  `json:"updateAt"`
}

type AccountUpateParameters struct {
	AppId        string                 `json:"appId"`
	AccountId    string                 `json:"accountId"`
	CustomId     string                 `json:"customeId"`
	Nickname     string                 `json:"nickname"`
	Avatar       string                 `json:"avatar"`
	Gender       int32                  `json:"gender"`
	Mobile       string                 `json:"mobile"`
	Country      string                 `json:"country"`
	State        string                 `json:"state"`
	City         string                 `json:"city"`
	Alias        string                 `json:"alias"`
	Qrcode       string                 `json:"qrcode"`
	Signature    string                 `json:"signature"`
	TenantId     string                 `json:"tenantId"`
	TenantUserId string                 `json:"tenantUserId"`
	ExtendProps  map[string]interface{} `json:"extendProps"`
	IsDeleted    bool                   `json:"isDeleted"`
}

type AccountAddParameters struct {
	AppId        string                 `json:"appId"`
	AccountId    string                 `json:"accountId"`
	CustomId     string                 `json:"customId"`
	Nickname     string                 `json:"nickname"`
	Avatar       string                 `json:"avatar"`
	Gender       int32                  `json:"gender"`
	Mobile       string                 `json:"mobile"`
	Country      string                 `json:"country"`
	State        string                 `json:"state"`
	City         string                 `json:"city"`
	Alias        string                 `json:"alias"`
	Qrcode       string                 `json:"qrcode"`
	Signature    string                 `json:"signature"`
	TenantId     string                 `json:"tenantId"`
	TenantUserId string                 `json:"tenantUserId"`
	ExtendProps  map[string]interface{} `json:"extendProps"`
}

func (api *Client) AccountAdd(req *AccountAddParameters) error {
	return api.AccountAddContext(context.Background(), req)
}

func (api *Client) AccountAddContext(ctx context.Context, req *AccountAddParameters) error {
	response := Response{}
	err := api.postJSONMethod(ctx, "/account/add", req, &response)
	if err != nil {
		return err
	}
	return response.Err()
}

func (api *Client) AccountUpdate(req *AccountUpateParameters) error {
	return api.AccountUpdateContext(context.Background(), req)
}

func (api *Client) AccountUpdateContext(ctx context.Context, req *AccountUpateParameters) error {
	response := Response{}
	err := api.postJSONMethod(ctx, "/account/update", req, &response)
	if err != nil {
		return err
	}
	return response.Err()
}
