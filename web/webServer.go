package web

import (
	"Project/web/controller"
	"fmt"
	"net/http"
)

// WebStart 启动Web服务并指定路由信息
func WebStart(app controller.Application)  {

	fs:= http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// 指定路由信息(匹配请求)
	http.HandleFunc("/", app.LoginView)
	http.HandleFunc("/login", app.Login)
	http.HandleFunc("/loginout", app.LoginOut)

	http.HandleFunc("/index", app.Index)
	http.HandleFunc("/help", app.Help)

	http.HandleFunc("/addFileInfo", app.AddFileShow)	// 显示添加信息页面
	http.HandleFunc("/addFile", app.AddFile)	// 提交信息请求

	http.HandleFunc("/queryPage", app.QueryPage)	// 转至根据ID与姓名查询信息页面
	http.HandleFunc("/query", app.FindCertByNoAndName)	// 根据ID与姓名查询信息

	http.HandleFunc("/modifyPage", app.ModifyShow)	// 修改信息页面
	http.HandleFunc("/modify", app.Modify)	//  修改信息

	fmt.Println("启动Web服务, 监听端口号为: 9000")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("Web服务启动失败: %v", err)
	}

}
