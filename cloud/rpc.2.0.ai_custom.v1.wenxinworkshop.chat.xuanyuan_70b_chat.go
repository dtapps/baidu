package cloud

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
	"net/http"
)

type Rpc2AiCustomV1WenxinworkshopChatXuanyuan70bChatResponse struct {
	LogId  string `json:"log_id"`
	Result struct {
		TemplateId         int         `json:"templateId"`        // prompt工程里面对应的模板id
		TemplateName       string      `json:"templateName"`      // 模板名称
		TemplateContent    string      `json:"templateContent"`   // 模板原始内容
		TemplateVariables  string      `json:"templateVariables"` // 模板变量插值
		Content            string      `json:"content"`           // 将变量插值填充到模板原始内容后得到的模板内容
		VariableIdentifier string      `json:"variableIdentifier"`
		Labels             interface{} `json:"labels"`
		CreatorName        string      `json:"creatorName"`
		Type               int         `json:"type"`
		SceneType          int         `json:"sceneType"`
		FrameworkType      int         `json:"frameworkType"`
	} `json:"result"`
	Status  int  `json:"status"`
	Success bool `json:"success"`
}

type Rpc2AiCustomV1WenxinworkshopChatXuanyuan70bChatResult struct {
	Result Rpc2AiCustomV1WenxinworkshopChatXuanyuan70bChatResponse // 结果
	Body   []byte                                                  // 内容
	Http   gorequest.Response                                      // 请求
}

func newRpc2AiCustomV1WenxinworkshopChatXuanyuan70bChatResult(result Rpc2AiCustomV1WenxinworkshopChatXuanyuan70bChatResponse, body []byte, http gorequest.Response) *Rpc2AiCustomV1WenxinworkshopChatXuanyuan70bChatResult {
	return &Rpc2AiCustomV1WenxinworkshopChatXuanyuan70bChatResult{Result: result, Body: body, Http: http}
}

// Rpc2AiCustomV1WenxinworkshopChatXuanyuan70bChat XuanYuan-70B
// https://console.bce.baidu.com/qianfan/modelcenter/model/buildIn/detail/9257
// https://cloud.baidu.com/doc/WENXINWORKSHOP/s/Ylp88e5jc
func (c *Client) Rpc2AiCustomV1WenxinworkshopChatXuanyuan70bChat(ctx context.Context, notMustParams ...gorequest.Params) (*Rpc2AiCustomV1WenxinworkshopChatXuanyuan70bChatResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, "/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/xuanyuan_70b_chat?access_token="+c.accessToken, params, http.MethodPost, "JSON")
	if err != nil {
		return newRpc2AiCustomV1WenxinworkshopChatXuanyuan70bChatResult(Rpc2AiCustomV1WenxinworkshopChatXuanyuan70bChatResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response Rpc2AiCustomV1WenxinworkshopChatXuanyuan70bChatResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newRpc2AiCustomV1WenxinworkshopChatXuanyuan70bChatResult(response, request.ResponseBody, request), err
}
