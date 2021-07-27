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
	ObjectType string `json:"docType"`
	Name       string `json:"Name"`
	EntityID   string `json:"EntityID"`
	FileHash   string `json:"FileHash"`
	FileType   string `json:"FileType"`
	Desc       string `json:"Desc"`
	Company    string `json:"Company"`

	History []HistoryItem
}

type HistoryItem struct {
	TxId         string
	Intelligence Intelligence
}
