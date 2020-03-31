package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"learngo/errhandling/filelistingserver/filelisting"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

// 返回一个panic异常
func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(122)
}

func errUserError(writer http.ResponseWriter, request *http.Request) error {
	return filelisting.UserError("user error")
}

func errNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

func errNotPermission(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

func errUnknown(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("Unknown error")
}

func noError(writer http.ResponseWriter, request *http.Request) error {
	fmt.Fprintln(writer, "no error")
	return nil
}

var tests = []struct {
	h       appHandler // 函数类型，放入函数
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errNotPermission, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}

func TestErrWrapper(t *testing.T) {
	// httptest.NewRecorder()本质上是一个responseWriter

	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet,
			"http://www.baidu.com",
			nil,
		)
		// 将结果返回到response中
		f(response, request)

		b, _ := ioutil.ReadAll(response.Body)
		body := strings.Trim(string(b), "\n")
		if response.Code != tt.code || body != tt.message {
			t.Errorf("expect (%d ,%s) got (%d, %s)",
				tt.code, tt.message, response.Code, body)
		}
	}
}

// 使用启动一个server来测试，使用假的response和request
func TestErrWrapperInServer(t *testing.T) {

	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		fmt.Println(server.URL)
		resp, _ := http.Get(server.URL)

		b, _ := ioutil.ReadAll(resp.Body)
		body := strings.Trim(string(b), "\n")
		if resp.StatusCode != tt.code || body != tt.message {
			t.Errorf("expect (%d ,%s) got (%d, %s)",
				tt.code, tt.message, resp.StatusCode, body)
		}
	}
}
