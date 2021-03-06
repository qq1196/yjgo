package error

import (
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/service/utils/response"
)

func Unauth(r *ghttp.Request) {
	response.BuildTpl(r, "error/unauth.html").WriteTpl()
}

func Error(r *ghttp.Request) {
	response.BuildTpl(r, "error/500.html").WriteTpl()
}

func NotFound(r *ghttp.Request) {
	response.BuildTpl(r, "error/404.html").WriteTpl()
}
