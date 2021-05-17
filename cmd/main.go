package main

import (
	"flag"
	"flow/internal/app/api"
	"flow/internal/conf"
	"flow/internal/pkg/tool"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
)

func handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			// catch handler painc
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()
		serveMux := http.NewServeMux()
		// 默认首页
		// http.HandleFunc("/", index.Index)
		serveMux.HandleFunc("/", api.Index)
		serveMux.HandleFunc("/favicon.ico", http.NotFound)
		// 文件浏览
		serveMux.HandleFunc("/file/browse/", api.FileBrowse)
		// 上传文件
		serveMux.HandleFunc("/file/upload", api.FileUpload)
		// 文本内容
		serveMux.HandleFunc("/text/post", api.TextPost)
		serveMux.HandleFunc("/text/get", api.TextGet)
		serveMux.ServeHTTP(w, r)
	})
}

func useArgs() (port int) {
	defaultPort := conf.Port
	flag.IntVar(&port, "port", defaultPort, "the server listen port")
	flag.Parse()
	return
}

func main() {
	port := useArgs()
	ip, err := tool.LanIp()
	if err != nil {
		log.Fatal(err)
	}
	addr := ip + ":" + strconv.Itoa(port)
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("访问服务地址：http://" + addr)
	log.Println("文件保存目录：" + path.Dir(pwd))

	if err := http.ListenAndServe(addr, handler()); err != nil {
		log.Fatal(err)
	}
}
