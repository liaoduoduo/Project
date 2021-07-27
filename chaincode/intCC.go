package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

const DOC_TYPE = "fileObj"

// PutFile 保存情报
// args: Intelligence
func PutFile(stub shim.ChaincodeStubInterface, file Intelligence) ([]byte, bool) {

	file.ObjectType = DOC_TYPE

	b, err := json.Marshal(file)
	if err != nil {
		return nil, false
	}

	// 保存file状态
	err = stub.PutState(file.EntityID, b)
	if err != nil {
		return nil, false
	}

	return b, true
}

//上传情报
//args:Intelligence结构体
func (t *IntelligenceChaincode) addFile(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 2 {
		return shim.Error("给定的参数个数不符合要求")
	}
	var file Intelligence
	err := json.Unmarshal([]byte(args[0]), &file)
	if err != nil {
		return shim.Error("反序列化信息时发生错误")
	}

	_, bl := PutFile(stub, file)
	if !bl {
		return shim.Error("保存信息时发生错误")
	}
	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte("信息添加成功"))
}

// 根据姓名及ID查询历史情报 （溯源）
// args: id, name
func (t *IntelligenceChaincode) queryFileByEntityIDAndName(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("给定的参数个数不符合要求")
	}
	id := args[0]
	name := args[1]
	queryString := fmt.Sprintf("{\"selector\":{\"EntityID\":\"%s\", \"Name\":\"%s\"}}", id, name)
	result, err := getFileByQueryString(stub, queryString)
	if err != nil {
		return shim.Error("根据ID及姓名查询信息时发生错误")
	}
	if result == nil {
		return shim.Error("根据ID及姓名没有查询到相关的信息")
	}

	var file Intelligence
	err = json.Unmarshal(result, &file)
	if err != nil {
		return shim.Error("反序列化file信息失败")
	}
	iterator, err := stub.GetHistoryForKey(file.EntityID)
	if err != nil {
		return shim.Error("根据指定的ID查询对应的历史变更数据失败")
	}
	defer iterator.Close()

	// 迭代处理
	var historys []HistoryItem
	var hisFile Intelligence
	for iterator.HasNext() {
		hisData, err := iterator.Next()
		if err != nil {
			return shim.Error("获取ID的历史变更数据失败")
		}

		var historyItem HistoryItem
		historyItem.TxId = hisData.TxId
		json.Unmarshal(hisData.Value, &hisFile)

		if hisData.Value == nil {
			var empty Intelligence
			historyItem.Intelligence = empty
		} else {
			historyItem.Intelligence = hisFile
		}

		historys = append(historys, historyItem)

	}

	file.History = historys

	// 返回
	result, err = json.Marshal(file)
	if err != nil {
		return shim.Error("序列化file信息时发生错误")
	}
	return shim.Success(result)
}

// 根据指定的查询字符串实现富查询
func getFileByQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}

		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		bArrayMemberAlreadyWritten = true
	}

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil

}

func GetFileInfo(stub shim.ChaincodeStubInterface, entityID string) (Intelligence, bool) {
	var file Intelligence
	// 根据ID查询信息状态
	b, err := stub.GetState(entityID)
	if err != nil {
		return file, false
	}

	if b == nil {
		return file, false
	}

	// 对查询到的状态进行反序列化
	err = json.Unmarshal(b, &file)
	if err != nil {
		return file, false
	}

	// 返回结果
	return file, true
}

// 根据ID更新情报文件
// args: IntelligenceObject
func (t *IntelligenceChaincode) updateFile(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("给定的参数个数不符合要求")
	}

	var info Intelligence
	err := json.Unmarshal([]byte(args[0]), &info)
	if err != nil {
		return shim.Error("反序列化edu信息失败")
	}

	// 根据身份证号码查询信息

	result, bl := GetFileInfo(stub, info.EntityID)
	if !bl {
		return shim.Error("根据ID查询信息时发生错误")
	}

	result.Name = info.Name
	result.EntityID = info.EntityID
	result.FileHash = info.FileHash
	result.FileType = info.FileType
	result.Desc = info.Desc
	result.Company = info.Company

	_, bl = PutFile(stub, result)
	if !bl {
		return shim.Error("保存信息信息时发生错误")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("信息更新成功"))
}
