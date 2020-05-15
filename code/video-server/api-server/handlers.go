package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"video-server/api-server/dao"
	"video-server/api-server/domain/dto"
	"video-server/api-server/domain/entity"
	respError "video-server/api-server/domain/error"
	"video-server/api-server/session"

	"github.com/julienschmidt/httprouter"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 从request将body读取出
	res, _ := ioutil.ReadAll(r.Body)

	user := &entity.User{}
	if err := json.Unmarshal(res, user); err != nil {
		sendErrorResponse(w, respError.ErrorRequestBodyParseFailed)
		return
	}
	err := dao.AddUser(user)
	if err != nil {
		sendErrorResponse(w, respError.ErrorDBOps)
		return
	}
	// 添加成功后放入session中
	id, _ := session.GenerateNewSessionID(user.Username)
	// 返回登录成功的消息
	su := &dto.SignedUp{true, id}

	if res, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, respError.ErrorInteralFailed)
	} else {
		sendResponse(w, string(res), 201)
	}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	username := p.ByName("username")
	io.WriteString(w, username)
}
