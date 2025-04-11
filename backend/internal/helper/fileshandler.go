package helper

import (
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/zeromicro/go-zero/rest"
)

func DirHandler(pater, fileDir string, isWithGzip bool) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // 指定哪些站点可以参与跨站资源共享
		w.Header().Add("Access-Control-Allow-Headers", "Content-Encoding,Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Add("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		w.Header().Add("Access-Control-Allow-Credentials", "true")

		if isWithGzip {
			w.Header().Add("Content-Encoding", "gzip") // 设置响应体的MIME类型
			w.Header().Add("Content-Type", "application/gzip")
		}
		handler := http.StripPrefix(pater, http.FileServer(http.Dir(fileDir)))
		handler.ServeHTTP(w, req)
	}
}

// FileRegisterHandlers 文件注册handler
func FileRegisterHandlers(engine *rest.Server, pater, dirPath string, isWithGzip bool) {
	/*
	   pater: 请求路径格式/assets/:1,/assets/:1/:2
	   fileDir: 映射对应的文件夹./assets/
	*/
	dirLevel := make([]string, 0)
	for i := 1; i < 20; i++ {
		dirLevel = append(dirLevel, fmt.Sprintf(":%d", i))
	}
	for i := 1; i <= len(dirLevel); i++ {
		pathStr := strings.Join(dirLevel[:i], "/")
		r := rest.Route{
			Method:  http.MethodGet,
			Path:    path.Join(pater, pathStr),
			Handler: DirHandler(pater, dirPath, isWithGzip),
		}
		engine.AddRoute(r)
	}
}
