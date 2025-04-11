package handler

import (
	"embed"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"

	"lowcode/internal/svc"
)

//go:embed swagger
var f embed.FS

type fileType string

// fileType
// https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Basics_of_HTTP/MIME_types/Common_types
const (
	JS   fileType = "js"   //.js
	CSS  fileType = "css"  //.css
	HTML fileType = "html" //.html and .map
	JSON fileType = "json" //.json
	PNG  fileType = "png"  //.png
)

func addSwagger(filepath string, filetype fileType, svc *svc.ServiceContext) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		body, err := f.ReadFile(filepath)
		if err != nil {
			panic("swagger file is not exist!")
		}
		switch filetype {
		case JS:
			writer.Header().Set("content-type", "application/javascript")
			fmt.Fprint(writer, string(body))
		case CSS:
			writer.Header().Set("content-type", "text/css")
			fmt.Fprint(writer, string(body))
		case JSON:
			var result map[string]interface{}
			err = json.Unmarshal(body, &result)
			if err != nil {
				logx.Info("swagger parse err: ", err)
			}
			result["host"] = svc.Config.Swagger.Host
			body, err = json.Marshal(result)
			if err != nil {
				logx.Info("swagger parse err: ", err)
			}
			writer.Header().Set("content-type", "application/json")
			fmt.Fprint(writer, string(body))
		case PNG:
			writer.Header().Set("content-type", "image/png")
			fmt.Fprint(writer, body)
		case HTML:
			fmt.Fprint(writer, string(body))
		}
	}
}

func RegisterSwaggerHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/swagger/swagger-ui-bundle.js",
				Handler: addSwagger("swagger/swagger-ui-bundle.js", JS, serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/swagger/swagger-ui.css",
				Handler: addSwagger("swagger/swagger-ui.css", CSS, serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/swagger/swagger-ui-standalone-preset.js",
				Handler: addSwagger("swagger/swagger-ui-standalone-preset.js", JS, serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/swagger/favicon-32x32.png",
				Handler: addSwagger("swagger/favicon-32x32.png", PNG, serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/swagger/index.html",
				Handler: addSwagger("swagger/index.html", HTML, serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/swagger/lowcode.json",
				Handler: addSwagger("swagger/lowcode.json", JSON, serverCtx),
			},
		},
	)
}
