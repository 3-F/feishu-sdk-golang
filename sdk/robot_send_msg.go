package sdk

import (
	"github.com/3-F/feishu-sdk-golang/core/consts"
	"github.com/3-F/feishu-sdk-golang/core/model/vo"
	"github.com/3-F/feishu-sdk-golang/core/util/http"
	"github.com/3-F/feishu-sdk-golang/core/util/json"
	"github.com/3-F/feishu-sdk-golang/core/util/log"
)

//发送消息卡片 https://open.feishu.cn/document/ukTMukTMukTM/uYTNwUjL2UDM14iN1ATN
func (t Tenant) SendMessage(msg vo.MsgVo) (*vo.MsgResp, error) {
	reqBody := json.ToJsonIgnoreError(msg)
	respBody, err := http.Post(consts.ApiRobotSendMessage, nil, reqBody, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.MsgResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//发送消息卡片V1: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/create
func (t Tenant) SendMessageV1() {

}

//发送消息卡片 https://open.feishu.cn/document/ukTMukTMukTM/uYTNwUjL2UDM14iN1ATN
func (t Tenant) SendMessageBatch(msg vo.BatchMsgVo) (*vo.MsgResp, error) {
	reqBody := json.ToJsonIgnoreError(msg)
	respBody, err := http.Post(consts.ApiRobotSendBatchMessage, nil, reqBody, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.MsgResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}
