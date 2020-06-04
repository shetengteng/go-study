package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

// 上传测试页面
func testPageHandler(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	t, err := template.ParseFiles("./videos/upload.html")
	if err != nil {
		log.Printf("error testpage %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "internal error")
	}
	t.Execute(w, nil)
}

// 下载视频
func downloadHandler(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	vidDir := VIDEO_DIR + vid
	videoFile, err := os.Open(vidDir)
	if err != nil {
		log.Printf("error when try to open file: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "internal error")
		return
	}
	// 给头部设置信息，用于解析文件，否则会通过头部的一部分字节信息分析结果
	w.Header().Set("Content-Type", "video/mp4")
	// 读取请求中的数据到本地
	http.ServeContent(w, req, "", time.Now(), videoFile)

	defer videoFile.Close()
}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "file is to big")
		return
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("error when try to get file :%v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("read file error:%v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "internal error")
		return
	}
	fn := p.ByName("vid-id")
	// 注意权限不要设置太大，不要0777
	err = ioutil.WriteFile(VIDEO_DIR+fn, data, 0666)
	if err != nil {
		log.Printf("writer file error :%v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "interal error")
		return
	}
	// 201
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "uploaded successfully")
}
