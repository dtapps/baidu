package cloud

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
	"net/http"
)

type Rest2WenxinworkshopApiV1TemplateInfoResponse struct {
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

type Rest2WenxinworkshopApiV1TemplateInfoResult struct {
	Result Rest2WenxinworkshopApiV1TemplateInfoResponse // 结果
	Body   []byte                                       // 内容
	Http   gorequest.Response                           // 请求
}

func newRest2WenxinworkshopApiV1TemplateInfoResult(result Rest2WenxinworkshopApiV1TemplateInfoResponse, body []byte, http gorequest.Response) *Rest2WenxinworkshopApiV1TemplateInfoResult {
	return &Rest2WenxinworkshopApiV1TemplateInfoResult{Result: result, Body: body, Http: http}
}

// Rest2WenxinworkshopApiV1TemplateInfo Rest2WenxinworkshopApiV1TemplateInfo
func (c *Client) Rest2WenxinworkshopApiV1TemplateInfo(ctx context.Context, notMustParams ...gorequest.Params) (*Rest2WenxinworkshopApiV1TemplateInfoResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("access_token", c.accessToken)
	// 请求
	request, err := c.request(ctx, "/rest/2.0/wenxinworkshop/api/v1/template/info", params, http.MethodGet, "FORM")
	if err != nil {
		return newRest2WenxinworkshopApiV1TemplateInfoResult(Rest2WenxinworkshopApiV1TemplateInfoResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response Rest2WenxinworkshopApiV1TemplateInfoResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newRest2WenxinworkshopApiV1TemplateInfoResult(response, request.ResponseBody, request), err
}
