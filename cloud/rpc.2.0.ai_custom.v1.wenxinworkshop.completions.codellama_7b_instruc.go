package cloud

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
	"net/http"
)

type Rpc2AiCustomV1WenxinworkshopCompletionsCodellama7bInstructResponse struct {
	Id               string `json:"id"`                    // 本轮对话的id
	Object           string `json:"object"`                // 回包类型
	Created          int    `json:"created"`               // 时间戳
	SentenceId       int    `json:"sentence_id,omitempty"` // 表示当前子句的序号。只有在流式接口模式下会返回该字段
	IsEnd            int    `json:"is_end,omitempty"`      // 表示当前子句是否是最后一句。只有在流式接口模式下会返回该字段
	IsTruncated      bool   `json:"is_truncated"`          // 当前生成的结果是否被截断
	Result           string `json:"result"`                // 对话返回结果
	NeedClearHistory bool   `json:"need_clear_history"`    // 表示用户输入是否存在安全，是否关闭当前会话，清理历史会话信息  true：是，表示用户输入存在安全风险，建议关闭当前会话，清理历史会话信息 false：否，表示用户输入无安全风险
	BanRound         int    `json:"ban_round"`             // 当need_clear_history为true时，此字段会告知第几轮对话有敏感信息，如果是当前问题，ban_round=-1
	Usage            struct {
		PromptTokens     int `json:"prompt_tokens"`     // 问题tokens数
		CompletionTokens int `json:"completion_tokens"` // 回答tokens数
		TotalTokens      int `json:"total_tokens"`      // tokens总数
	} `json:"usage"` // token统计信息
}

type Rpc2AiCustomV1WenxinworkshopCompletionsCodellama7bInstructResult struct {
	Result Rpc2AiCustomV1WenxinworkshopCompletionsCodellama7bInstructResponse // 结果
	Body   []byte                                                             // 内容
	Http   gorequest.Response                                                 // 请求
}

func newRpc2AiCustomV1WenxinworkshopCompletionsCodellama7bInstructResult(result Rpc2AiCustomV1WenxinworkshopCompletionsCodellama7bInstructResponse, body []byte, http gorequest.Response) *Rpc2AiCustomV1WenxinworkshopCompletionsCodellama7bInstructResult {
	return &Rpc2AiCustomV1WenxinworkshopCompletionsCodellama7bInstructResult{Result: result, Body: body, Http: http}
}

// Rpc2AiCustomV1WenxinworkshopCompletionsCodellama7bInstruct Llama-2-13B
// https://console.bce.baidu.com/qianfan/modelcenter/model/buildIn/detail/8139
// https://cloud.baidu.com/doc/WENXINWORKSHOP/s/ylo47d03k
func (c *Client) Rpc2AiCustomV1WenxinworkshopCompletionsCodellama7bInstruct(ctx context.Context, notMustParams ...gorequest.Params) (*Rpc2AiCustomV1WenxinworkshopCompletionsCodellama7bInstructResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, "/rpc/2.0/ai_custom/v1/wenxinworkshop/completions/codellama_7b_instruct?access_token="+c.accessToken, params, http.MethodPost, "JSON")
	if err != nil {
		return newRpc2AiCustomV1WenxinworkshopCompletionsCodellama7bInstructResult(Rpc2AiCustomV1WenxinworkshopCompletionsCodellama7bInstructResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response Rpc2AiCustomV1WenxinworkshopCompletionsCodellama7bInstructResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newRpc2AiCustomV1WenxinworkshopCompletionsCodellama7bInstructResult(response, request.ResponseBody, request), err
}