package main

/**
警员姓名：张三
警员ID：101
情报hash(IPFS的CID):QmdKjne7dhQ99GxMZ5DqcN3FqYB2TrQN3a8hefFiouAC2a
情报文件类型：jpg
文件描述：2021年7月，广东省广州市某小区公寓房内有多人聚众吸毒照片
所属组织：广州市花都区花山镇禁毒办
*/

type Intelligence struct {
	ObjectType              string `json:"docType"`
	Name                    string `json:"Name"`                    //警员姓名
	EntityID                string `json:"EntityID"`                //ID、工号
	FileHash                string `json:"FileHash"`                //文件哈希
	FileType                string `json:"FileType"`                //文件类型
	Desc                    string `json:"Desc"`                    //线索描述
	Company                 string `json:"Company"`                 //基本信息
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
