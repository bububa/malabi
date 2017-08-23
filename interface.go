package malabi

import (
	"github.com/levigross/grequests"
)

const (
	GATEWAY       = "http://api.malabi.co/Camera51Server/processImage"
	GATEWAY_ASYNC = "http://api.malabi.co/Camera51Server/processImageAsync"
)

type ResultCode uint

const (
	SUCCESS_CODE_ASYNC       ResultCode = 1
	UNKNOWN_ERROR_ASYNC      ResultCode = 2
	LOW_CONTRAST_ERROR_ASYNC ResultCode = 3
	BAD_LIGHTING_ERROR_ASYNC ResultCode = 4
	NOT_SUITABLE_ERROR_ASYNC ResultCode = 5
	DONT_NEED_ERROR_ASYNC    ResultCode = 6
	SUCCESS_CODE             ResultCode = 7
	UNKNOWN_ERROR            ResultCode = 8
	LOW_CONTRAST_ERROR       ResultCode = 9
	BAD_LIGHTING_ERROR       ResultCode = 10
	NOT_SUITABLE_ERROR       ResultCode = 11
	DONT_NEED_ERROR          ResultCode = 12
	BAD_IMAGE_ERROR          ResultCode = 101
	TOO_SMALL_ERROR          ResultCode = 103
)

type Request struct {
	OriginalImageURL string `json:"originalImageURL,omitempty" codec:"originalImageURL,omitempty" url:"originalImageURL"`
	CustomerId       uint64 `json:"customerId,omitempty" codec:"customerId,omitempty" url:"customerId"`
	Token            string `json:"token,omitempty" codec:"token,omitempty" url:"token"`
	Shadow           bool   `json:"shadow,omitempty" codec:"shadow,omitempty" url:"shadow"`
	Transparent      bool   `json:"transparent,omitempty" codec:"transparent,omitempty" url:"transparent"`
	SignedURL        string `json:"signedURL,omitempty" codec:"signedURL,omitempty" url:"signedURL"`
	TrackId          string `json:"trackId,omitempty" codec:"trackId,omitempty" url:"trackId"`
	CallbackURL      string `json:"callbackURL,omitempty" codec:"callbackURL,omitempty" url:"callbackURL"`
}

type Response struct {
	Result *Result `json:"response,omitempty"`
}

type Result struct {
	ImageCopyURL         string     `json:"imageCopyURL,omitempty" codec:"imageCopyURL,omitempty"`
	ResultImageURL       string     `json:"resultImageURL,omitempty" codec:"resultImageURL,omitempty"`
	ResultEditMaskURL    string     `json:"resultEditMaskURL,omitempty" codec:"resultEditMaskURL,omitempty"`
	TrackId              string     `json:"trackId,omitempty" codec:"trackId,omitempty"`
	SessionId            string     `json:"sessionId,omitempty" codec:"sessionId,omitempty"`
	ProcessingResultCode ResultCode `json:"processingResultCode,omitempty" codec:"processingResultCode,omitempty"`
	Errors               []string   `json:"errors,omitempty" codec:"errors,omitempty"`
}

type Client struct {
	Session    *grequests.Session
	CustomerId uint64
	Token      string
}
