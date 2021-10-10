package service

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"time"
)

type Intelligence struct {
	ObjectType string `json:"docType"`
	Name       string `json:"Name"`
	EntityID   string `json:"EntityID"`
	FileHash   string `json:"FileHash"`
	FileType   string `json:"FileType"`
	Desc       string `json:"Desc"`
	Company    string `json:"Company"`

	PermissionsInfo         string `json:"PermissionsInfo"`         //权限信息：只读、更改、完全控制
	MachineInfo             string `json:"MachineInfo"`             //智能研判信息
	MergeInfo               string `json:"MergeInfo"`               //合并信息
	GradeInfo               string `json:"GradeInfo"`               //分类定级信息
	AllocationInfo          string `json:"AllocationInfo"`          //分配信息
	SpecialItemInfo         string `json:"SpecialItemInfo"`         //涉毒研判专项建立信息
	SpecialGroupInfo        string `json:"SpecialGroupInfo"`        //线索研判小组建立信息
	TaskEstablishAssignInfo string `json:"TaskEstablishAssignInfo"` //研判任务建立和分配信息
	TaskReceivedInfo        string `json:"TaskReceivedInfo"`        //研判任务签收信息
	ResearchLogInfo         string `json:"ResearchLogInfo"`         //研判日志信息
	ResultFeedbackInfo      string `json:"ResultFeedbackInfo"`      //研判结果反馈信息
	EvaluationInfo          string `json:"EvaluationInfo"`          //研判任务评价信息
	ReportInfo              string `json:"ReportInfo"`              //研判报告生成信息
	ResultAuditInfo         string `json:"ResultAuditInfo"`         //研判结果审核信息

	History []HistoryItem
}

type HistoryItem struct {
	TxId         string
	Intelligence Intelligence
}
type ServiceSetup struct {
	ChaincodeID string
	Client      *channel.Client
}

func regitserEvent(client *channel.Client, chaincodeID, eventID string) (fab.Registration, <-chan *fab.CCEvent) {

	reg, notifier, err := client.RegisterChaincodeEvent(chaincodeID, eventID)
	if err != nil {
		fmt.Println("注册链码事件失败: %s", err)
	}
	return reg, notifier
}

func eventResult(notifier <-chan *fab.CCEvent, eventID string) error {
	select {
	case ccEvent := <-notifier:
		fmt.Printf("接收到链码事件: %v\n", ccEvent)
	case <-time.After(time.Second * 20):
		return fmt.Errorf("不能根据指定的事件ID接收到相应的链码事件(%s)", eventID)
	}
	return nil
}
