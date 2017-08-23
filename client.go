package malabi

import (
	"github.com/levigross/grequests"
	"strconv"
)

func NewClient(customerId uint64, token string) *Client {
	ro := &grequests.RequestOptions{
		Headers: map[string]string{
			"Cache-Control": "no-cache",
			"Content-Type":  "application/x-www-form-urlencoded",
		},
	}
	session := grequests.NewSession(ro)
	return &Client{Session: session, CustomerId: customerId, Token: token}
}

func (this *Client) ProcessImage(req *Request) (*Result, error) {
	if req.CustomerId == 0 {
		req.CustomerId = this.CustomerId
	}
	if req.Token == "" {
		req.Token = this.Token
	}
	data := make(map[string]string)
	data["originalImageURL"] = req.OriginalImageURL
	data["customerId"] = strconv.FormatUint(req.CustomerId, 10)
	data["token"] = req.Token
	if req.Shadow {
		data["shadow"] = "true"
	}
	if req.Transparent {
		data["transparent"] = "true"
	}
	if req.SignedURL != "" {
		data["signedURL"] = req.SignedURL
	}
	if req.TrackId != "" {
		data["trackId"] = req.TrackId
	}
	if req.CallbackURL != "" {
		data["callbackURL"] = req.CallbackURL
	}
	ro := &grequests.RequestOptions{
		Data: data,
	}
	res, err := this.Session.Post(GATEWAY, ro)
	if err != nil {
		return nil, err
	}
	var response Response
	err = res.JSON(&response)
	if err != nil {
		return nil, err
	}
	return response.Result, nil
}

func (this *Client) ProcessImageAsync(req *Request) (*Result, error) {
	if req.CustomerId == 0 {
		req.CustomerId = this.CustomerId
	}
	if req.Token == "" {
		req.Token = this.Token
	}
	data := make(map[string]string)
	data["originalImageURL"] = req.OriginalImageURL
	data["customerId"] = strconv.FormatUint(req.CustomerId, 10)
	data["token"] = req.Token
	if req.Shadow {
		data["shadow"] = "true"
	}
	if req.Transparent {
		data["transparent"] = "true"
	}
	if req.SignedURL != "" {
		data["signedURL"] = req.SignedURL
	}
	if req.TrackId != "" {
		data["trackId"] = req.TrackId
	}
	if req.CallbackURL != "" {
		data["callbackURL"] = req.CallbackURL
	}
	ro := &grequests.RequestOptions{
		Data: data,
	}
	res, err := this.Session.Post(GATEWAY_ASYNC, ro)
	if err != nil {
		return nil, err
	}
	var response Response
	err = res.JSON(&response)
	if err != nil {
		return nil, err
	}
	return response.Result, nil
}
