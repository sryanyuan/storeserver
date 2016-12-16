package main

const HTTPStatusError = 480

const (
	HTTPRspCodeOK = iota
	HTTPRspCodeInternalError
)

type HTTPResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
