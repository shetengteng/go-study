package main

import (
	"encoding/json"
	"io"
	"net/http"
	respError "video-server/api-server/domain/error"
)

// 发送错误消息
func sendErrorResponse(w http.ResponseWriter, resp respError.Response) {
	w.WriteHeader(resp.HttpSC)
	res, _ := json.Marshal(&resp.Error)
	io.WriteString(w, string(res))
}

// 发送正常的消息
func sendResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
