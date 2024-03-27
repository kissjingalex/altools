package response

import jsoniter "github.com/json-iterator/go"

type ImLogicHttpReplyWrapper struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    jsoniter.RawMessage `json:"data"`
}
