/**
  @Prject: goProjects
  @Dev_Software: GoLand
  @File : controllerHandler
  @Time : 2018/10/18 14:31
  @Author : hanxiaodong
*/

package controller

import (
	"Project/service"
	"encoding/json"
	"fmt"
	"net/http"
)

var cuser User

func (app *Application) LoginView(w http.ResponseWriter, r *http.Request) {

	ShowView(w, r, "login.html", nil)
}

func (app *Application) Index(w http.ResponseWriter, r *http.Request) {
	ShowView(w, r, "index.html", nil)
}

func (app *Application) Help(w http.ResponseWriter, r *http.Request) {
	data := &struct {
		CurrentUser User
	}{
		CurrentUser: cuser,
	}
	ShowView(w, r, "help.html", data)
}

// Login 用户登录
func (app *Application) Login(w http.ResponseWriter, r *http.Request) {
	loginName := r.FormValue("loginName")
	password := r.FormValue("password")

	var flag bool
	for _, user := range users {
		if user.LoginName == loginName && user.Password == password {
			cuser = user
			flag = true
			break
		}
	}

	data := &struct {
		CurrentUser User
		Flag        bool
	}{
		CurrentUser: cuser,
		Flag:        false,
	}

	if flag {
		// 登录成功
		ShowView(w, r, "index.html", data)
	} else {
		// 登录失败
		data.Flag = true
		data.CurrentUser.LoginName = loginName
		ShowView(w, r, "login.html", data)
	}
}

// LoginOut 用户登出
func (app *Application) LoginOut(w http.ResponseWriter, r *http.Request) {
	cuser = User{}
	ShowView(w, r, "login.html", nil)
}

// AddFileShow 显示添加信息页面
func (app *Application) AddFileShow(w http.ResponseWriter, r *http.Request) {
	data := &struct {
		CurrentUser User
		Msg         string
		Flag        bool
	}{
		CurrentUser: cuser,
		Msg:         "",
		Flag:        false,
	}
	ShowView(w, r, "addFile.html", data)
}

// AddFile 添加信息
func (app *Application) AddFile(w http.ResponseWriter, r *http.Request) {

	file := service.Intelligence{
		Name:     r.FormValue("Name"),
		EntityID: r.FormValue("EntityID"),
		FileHash: r.FormValue("FileHash"),
		FileType: r.FormValue("FileType"),
		Desc:     r.FormValue("Desc"),
		Company:  r.FormValue("Company"),

		PermissionsInfo:         r.FormValue("PermissionsInfo"),
		MachineInfo:             r.FormValue("MachineInfo"),
		MergeInfo:               r.FormValue("MergeInfo"),
		GradeInfo:               r.FormValue("GradeInfo"),
		AllocationInfo:          r.FormValue("AllocationInfo"),
		SpecialItemInfo:         r.FormValue("SpecialItemInfo"),
		SpecialGroupInfo:        r.FormValue("SpecialGroupInfo"),
		TaskEstablishAssignInfo: r.FormValue("TaskEstablishAssignInfo"),
		TaskReceivedInfo:        r.FormValue("TaskReceivedInfo"),
		ResearchLogInfo:         r.FormValue("ResearchLogInfo"),
		ResultFeedbackInfo:      r.FormValue("ResultFeedbackInfo"),
		EvaluationInfo:          r.FormValue("EvaluationInfo"),
		ReportInfo:              r.FormValue("ReportInfo"),
		ResultAuditInfo:         r.FormValue("ResultAuditInfo"),
	}

	app.Setup.SaveFile(file)

	r.Form.Set("EntityID", file.EntityID)
	r.Form.Set("Name", file.Name)
	app.FindCertByNoAndName(w, r)
}

func (app *Application) QueryPage(w http.ResponseWriter, r *http.Request) {
	data := &struct {
		CurrentUser User
		Msg         string
		Flag        bool
	}{
		CurrentUser: cuser,
		Msg:         "",
		Flag:        false,
	}
	ShowView(w, r, "query.html", data)
}

// FindCertByNoAndName 根据ID与姓名查询信息
func (app *Application) FindCertByNoAndName(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("EntityID")
	name := r.FormValue("Name")
	result, err := app.Setup.FindFileByIDAndName(id, name)
	var file = service.Intelligence{}
	json.Unmarshal(result, &file)

	fmt.Println("根据ID与姓名查询信息成功：")
	fmt.Println(file)

	data := &struct {
		File        service.Intelligence
		CurrentUser User
		Msg         string
		Flag        bool
		History     bool
	}{
		File:        file,
		CurrentUser: cuser,
		Msg:         "",
		Flag:        false,
		History:     true,
	}

	if err != nil {
		data.Msg = err.Error()
		data.Flag = true
	}

	ShowView(w, r, "queryResult.html", data)
}

func (app *Application) QueryPage2(w http.ResponseWriter, r *http.Request) {
	data := &struct {
		CurrentUser User
		Msg         string
		Flag        bool
	}{
		CurrentUser: cuser,
		Msg:         "",
		Flag:        false,
	}
	ShowView(w, r, "query2.html", data)
}

// ModifyShow 修改/添加新信息
func (app *Application) ModifyShow(w http.ResponseWriter, r *http.Request) {
	// 根据ID与姓名查询信息
	id := r.FormValue("EntityID")
	name := r.FormValue("Name")
	result, err := app.Setup.FindFileByIDAndName(id, name)

	var file = service.Intelligence{}
	json.Unmarshal(result, &file)

	data := &struct {
		File        service.Intelligence
		CurrentUser User
		Msg         string
		Flag        bool
	}{
		File:        file,
		CurrentUser: cuser,
		Flag:        true,
		Msg:         "",
	}

	if err != nil {
		data.Msg = err.Error()
		data.Flag = true
	}

	ShowView(w, r, "modify.html", data)
}

// Modify 修改/添加新信息
func (app *Application) Modify(w http.ResponseWriter, r *http.Request) {
	file := service.Intelligence{
		Name:     r.FormValue("Name"),
		EntityID: r.FormValue("EntityID"),
		FileHash: r.FormValue("FileHash"),
		FileType: r.FormValue("FileType"),
		Desc:     r.FormValue("Desc"),
		Company:  r.FormValue("Company"),

		PermissionsInfo:         r.FormValue("PermissionsInfo"),
		MachineInfo:             r.FormValue("MachineInfo"),
		MergeInfo:               r.FormValue("MergeInfo"),
		GradeInfo:               r.FormValue("GradeInfo"),
		AllocationInfo:          r.FormValue("AllocationInfo"),
		SpecialItemInfo:         r.FormValue("SpecialItemInfo"),
		SpecialGroupInfo:        r.FormValue("SpecialGroupInfo"),
		TaskEstablishAssignInfo: r.FormValue("TaskEstablishAssignInfo"),
		TaskReceivedInfo:        r.FormValue("TaskReceivedInfo"),
		ResearchLogInfo:         r.FormValue("ResearchLogInfo"),
		ResultFeedbackInfo:      r.FormValue("ResultFeedbackInfo"),
		EvaluationInfo:          r.FormValue("EvaluationInfo"),
		ReportInfo:              r.FormValue("ReportInfo"),
		ResultAuditInfo:         r.FormValue("ResultAuditInfo"),
	}

	//transactionID, err := app.Setup.ModifyEdu(edu)
	app.Setup.ModifyFile(file)

	/*data := &struct {
		Edu service.Education
		CurrentUser User
		Msg string
		Flag bool
	}{
		CurrentUser:cuser,
		Flag:true,
		Msg:"",
	}

	if err != nil {
		data.Msg = err.Error()
	}else{
		data.Msg = "新信息添加成功:" + transactionID
	}

	ShowView(w, r, "modify.html", data)
	*/

	r.Form.Set("EntityID", file.EntityID)
	r.Form.Set("Name", file.Name)
	app.FindCertByNoAndName(w, r)
}
