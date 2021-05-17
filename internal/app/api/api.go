package api

import (
	"container/list"
	"flow/internal/app/view"
	"flow/internal/conf"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

func Index(rw http.ResponseWriter, r *http.Request) {
	view.Show(rw, "index.html", nil)
}

func FileBrowse(rw http.ResponseWriter, r *http.Request) {
	pathPrefix := "/file/browse/"
	fileRoot := conf.ShareFileRoot
	fpath := fileRoot + "/" + strings.TrimPrefix(r.URL.Path, pathPrefix)
	finfo, err := os.Stat(path.Clean(fpath))
	if err != nil {
		log.Println(err)
		return
	}
	fileServe := func() {
		http.StripPrefix(pathPrefix, http.FileServer(http.Dir(fileRoot))).ServeHTTP(rw, r)
	}
	if finfo.IsDir() {
		html := `<!DOCTYPE html><html><head><meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0"></head><style type="text/css">a {font-size: 1rem;}</style><body>`
		io.WriteString(rw, html)
		fileServe()
		html = "</body></html>"
		io.WriteString(rw, html)
	} else {
		fileServe()
	}
}

func FileUpload(rw http.ResponseWriter, r *http.Request) {
	srcFile, header, err := r.FormFile("file")
	if err != nil {
		log.Println(err)
		return
	}
	defer srcFile.Close()

	dstFile, err := os.Create("./" + header.Filename)
	if err != nil {
		log.Println(err)
		return
	}
	defer dstFile.Close()
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("收到新文件：" + header.Filename)
	io.WriteString(rw, "ok")
}

var textMessage strings.Builder

func TextPost(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	message := r.Form.Get("message")
	log.Println("客户端内容：" + message)
	// messageList.PushBack(message)
	textMessage.WriteString(message)
	textMessage.WriteRune('\n')
	io.WriteString(rw, message)
}

func TextGet(rw http.ResponseWriter, r *http.Request) {
	text := textMessage.String()
	io.WriteString(rw, text)
	list.New()
}
