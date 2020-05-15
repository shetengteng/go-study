package main

import (
	"net/http"
	respError "video-server/api-server/domain/error"
	"video-server/api-server/session"
)

// X-打头的header是自定义header
var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FIELD_UNAME = "X-User-Name"

// 验证session有没有相应的sessionid，得到sessionid查看是否过期，没有过期则更换为name
// 暂时不用
func validateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0 {
		return false
	}
	username, isOk := session.IsSessionExpired(sid)
	if isOk {
		return false
	}
	// 验证通过，添加username
	r.Header.Set(HEADER_FIELD_UNAME, username)
	return true
}

func addUsernameInSession(r *http.Request) bool {
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) != 0 {
		if username, isOk := session.IsSessionExpired(sid); !isOk {
			// 验证通过，添加username
			r.Header.Set(HEADER_FIELD_UNAME, username)
		}
	}
	return true
}

// 后续的方法需要验证session中是否有name
func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	username := r.Header.Get(HEADER_FIELD_UNAME)
	if len(username) == 0 {
		sendErrorResponse(w, respError.ErrorNotAuthUser)
		return false
	}
	return true
}
