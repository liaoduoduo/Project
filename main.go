package main

import (
	"Project/sdkInit"
	"Project/service"
	"Project/web"
	"Project/web/controller"
	"encoding/json"
	"fmt"
	"os"
)

const (
	configFile  = "config.yaml"
	initialized = false
	CC          = "lddcc"
)

func main() {
	initInfo := &sdkInit.InitInfo{

		ChannelID:     "kevinkongyixueyuan",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/liaoduoduo/Project/fixtures/artifacts/channel.tx",

		OrgAdmin:       "Admin",
		OrgName:        "Org1",
		OrdererOrgName: "orderer.kevin.kongyixueyuan.com",

		ChaincodeID:     CC,
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "github.com/liaoduoduo/Project/chaincode/",
		UserName:        "User1",
	}
	sdk, err := sdkInit.SetupSDK(configFile, initialized)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	defer sdk.Close()

	err = sdkInit.CreateChannel(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	channelClient, err := sdkInit.InstallAndInstantiateCC(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(channelClient)

	//===========================================//

	serviceSetup := service.ServiceSetup{
		ChaincodeID: CC,
		Client:      channelClient,
	}

	file := service.Intelligence{
		Name:                    "廖多越",
		EntityID:                "123",
		FileHash:                "QmdKjne7dhQ99GxMZ5DqcN3FqYB2TrQN3a8hefFiouAC2a",
		FileType:                "jpg",
		Desc:                    "2021年7月，广东省广州市某小区公寓房内有多人聚众吸毒照片",
		Company:                 "广州市花都区花山镇禁毒办",
		PermissionsInfo:         "N/A",
		MachineInfo:             "N/A",
		MergeInfo:               "N/A",
		GradeInfo:               "N/A",
		AllocationInfo:          "N/A",
		SpecialItemInfo:         "N/A",
		SpecialGroupInfo:        "N/A",
		TaskEstablishAssignInfo: "N/A",
		TaskReceivedInfo:        "N/A",
		ResearchLogInfo:         "N/A",
		ResultFeedbackInfo:      "N/A",
		EvaluationInfo:          "N/A",
		ReportInfo:              "N/A",
		ResultAuditInfo:         "N/A",
	}

	msg, err := serviceSetup.SaveFile(file)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("信息发布成功, 交易编号为: " + msg)
	}

	// 根据证书编号与名称查询信息
	result, err := serviceSetup.FindFileByIDAndName("123", "廖多越")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var file service.Intelligence
		json.Unmarshal(result, &file)
		fmt.Println("根据ID与姓名查询信息成功：")
		fmt.Println(file)
	}

	app := controller.Application{
		Setup: &serviceSetup,
	}
	web.WebStart(app)
}
