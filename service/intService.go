package service

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func (t *ServiceSetup) SaveFile(file Intelligence) (string, error) {

	eventID := "eventAddFile"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	// 将对象序列化成为字节数组
	b, err := json.Marshal(file)
	if err != nil {
		return "", fmt.Errorf("指定的file对象序列化时发生错误")
	}

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "addFile", Args: [][]byte{b, []byte(eventID)}}
	response, err := t.Client.Execute(req)
	if err != nil {
		return "", err
	}

	err = eventResult(notifier, eventID)
	if err != nil {
		return "", err
	}

	return string(response.TransactionID), nil
}

func (t *ServiceSetup) FindFileByIDAndName(id, name string) ([]byte, error){

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "queryFileByEntityIDAndName", Args: [][]byte{[]byte(id), []byte(name)}}
	response, err := t.Client.Query(req)
	if err != nil {
		return []byte{0x00}, err
	}

	return response.Payload, nil
}

func (t *ServiceSetup) ModifyFile(file Intelligence) (string, error) {

	eventID := "eventModifyFile"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	// 将edu对象序列化成为字节数组
	b, err := json.Marshal(file)
	if err != nil {
		return "", fmt.Errorf("指定的File对象序列化时发生错误")
	}

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "updateFile", Args: [][]byte{b, []byte(eventID)}}
	respone, err := t.Client.Execute(req)
	if err != nil {
		return "", err
	}

	err = eventResult(notifier, eventID)
	if err != nil {
		return "", err
	}

	return string(respone.TransactionID), nil
}