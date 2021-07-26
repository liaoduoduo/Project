package main

import (
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

type IntelligenceChaincode struct {

}

func (t *IntelligenceChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response{

	return shim.Success(nil)
}

func (t *IntelligenceChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response{
	// 获取用户意图
	fun, args := stub.GetFunctionAndParameters()

	if fun == "addFile"{
		return t.addFile(stub, args)		// 上传情报
	}else if fun == "queryFileByEntityIDAndName" {
		return t.queryFileByEntityIDAndName(stub, args)		// 根据上传者ID与姓名查询情报哈希
	}else if fun == "updateFile" {
		return t.updateFile(stub, args)
	}
		return shim.Error("指定的函数名称错误")

}

func main() {
	err := shim.Start(new(IntelligenceChaincode))
	if err != nil {
		fmt.Printf("启动Chaincode时发生错误: %s", err)
	}
}