package service

import (
	"LightningOnOmni/bean"
	"LightningOnOmni/dao"
	"LightningOnOmni/tool"
	"encoding/json"
	"errors"
	"github.com/asdine/storm/q"
	"log"
	"strconv"
)

type htlcQueryManager struct{}

var HtlcQueryService = htlcQueryManager{}

//非R创建者，通过承诺交易获取R
func (service *htlcQueryManager) GetRFromCommitmentTx(msgData string, user bean.User) (r string, err error) {

	if tool.CheckIsString(&msgData) == false {
		return r, errors.New("error input data")
	}
	reqData := bean.ChannelIdReq{}
	err = json.Unmarshal([]byte(msgData), &reqData)
	if err != nil {
		return r, err
	}
	if bean.ChannelIdService.IsEmpty(reqData.ChannelId) {
		return r, errors.New("error ChannelId ")
	}

	commitmentTxInfo, err := getLatestCommitmentTx(reqData.ChannelId, user.PeerId)
	if err != nil {
		return r, err
	}
	if commitmentTxInfo.TxType != dao.CommitmentTransactionType_Htlc {
		return r, errors.New("error tx type")
	}
	r = commitmentTxInfo.HtlcR
	if tool.CheckIsString(&r) == false {
		err = errors.New("empty R")
	}
	return r, err
}

//通过H获取路径信息
func (service *htlcQueryManager) GetPathInfoByH(msgData string, user bean.User) (pathInfo *dao.HtlcSingleHopPathInfo, err error) {
	if tool.CheckIsString(&msgData) == false {
		return nil, errors.New("error input data")
	}

	pathInfo = &dao.HtlcSingleHopPathInfo{}
	err = db.Select(q.Eq("H", msgData)).First(pathInfo)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return pathInfo, nil
}

//R创建者通过H获取R
func (service *htlcQueryManager) GetRByHOfOwner(msgData string, user bean.User) (r string, err error) {
	if tool.CheckIsString(&msgData) == false {
		return "", errors.New("error input data")
	}

	info := &dao.HtlcRAndHInfo{}
	err = db.Select(q.Eq("H", msgData), q.Eq("SignBy", user.PeerId)).First(info)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return info.R, nil
}

//转账发起方获取H列表
func (service *htlcHMessageManager) GetHtlcCreatedRandHInfoList(user *bean.User) (data interface{}, err error) {
	var rAndHInfoList []dao.HtlcRAndHInfo
	err = db.Select(q.Eq("CreateBy", user.PeerId)).OrderBy("CreateAt").Reverse().Find(&rAndHInfoList)
	if err != nil {
		return nil, err
	}
	for _, item := range rAndHInfoList {
		item.R = ""
	}
	return rAndHInfoList, nil
}

//转账发起方通过Id获取H结构体详细
func (service *htlcHMessageManager) GetHtlcCreatedRandHInfoItemById(msgData string, user *bean.User) (data interface{}, err error) {
	id, err := strconv.Atoi(msgData)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var rAndHInfo dao.HtlcRAndHInfo
	err = db.Select(q.Eq("Id", id), q.Eq("CreateBy", user.PeerId)).First(&rAndHInfo)
	if err != nil {
		return nil, err
	}
	rAndHInfo.R = ""
	return rAndHInfo, nil
}

//htlc最终收款方获取HR的结构体列表
func (service *htlcHMessageManager) GetHtlcSignedRandHInfoList(user *bean.User) (data interface{}, err error) {
	var rAndHInfoList []dao.HtlcRAndHInfo
	err = db.Select(q.Eq("RecipientPeerId", user.PeerId), q.Eq("SignBy", user.PeerId)).OrderBy("CreateAt").Reverse().Find(&rAndHInfoList)
	if err != nil {
		return nil, err
	}
	return rAndHInfoList, nil
}

//htlc最终收款方通过id获取HR详情
func (service *htlcHMessageManager) GetHtlcSignedRandHInfoItem(msgData string, user *bean.User) (data interface{}, err error) {
	id, err := strconv.Atoi(msgData)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var rAndHInfo dao.HtlcRAndHInfo
	err = db.Select(q.Eq("Id", id), q.Eq("SignBy", user.PeerId)).First(&rAndHInfo)
	if err != nil {
		return nil, err
	}
	return rAndHInfo, nil
}