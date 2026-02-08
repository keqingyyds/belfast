package middleware

import (
	"https://gh.xmly.dev/github.com/kataras/iris/v12"
	"https://gh.xmly.dev/github.com/kataras/iris/v12/middleware/recover"
)

func Recover() iris.Handler {
	return recover.New()
}
