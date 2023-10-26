package main

import (
	"github.com/kataras/iris"
	"log"
)

func main() {
	var err error
	//创建server实例
	app := iris.New()
	//设置错误模式,以及错误级别
	app.Logger().SetLevel("debug")
	//注册模板
	template := iris.HTML("./backend/web/views", ".html").
		Layout("shared/layout.html").Reload(true)
	app.RegisterView(template)
	//设置模板目标
	app.StaticWeb("/assets", "./backend/web/assets")
	//出现异常跳转到指定页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "Access error!"))
		ctx.ViewLayout("")
		err = ctx.View("shared/error.html")
		if err != nil {
			log.Fatalf("iris View error.html err:%v", err)
		}
	})
	//注册控制器
	//启动server
	err = app.Run(
		iris.Addr("localhost:8000"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
	if err != nil {
		log.Fatalf("app Run err:%v", err)
	}
}
