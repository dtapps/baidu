package baidu

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
	"net/http"
)

type MapGeocodingResponse struct {
	Status int `json:"status"`
	Result struct {
		Location struct {
			Lng float64 `json:"lng"`
			Lat float64 `json:"lat"`
		} `json:"location"`
		Precise       int    `json:"precise"`
		Confidence    int    `json:"confidence"`
		Comprehension int    `json:"comprehension"`
		Level         string `json:"level"`
	} `json:"result"`
}

type MapGeocodingResult struct {
	Result MapGeocodingResponse // 结果
	Body   []byte               // 内容
	Http   gorequest.Response   // 请求
}

func newMapGeocodingResult(result MapGeocodingResponse, body []byte, http gorequest.Response) *MapGeocodingResult {
	return &MapGeocodingResult{Result: result, Body: body, Http: http}
}

// V3 地理编码
// https://lbsyun.baidu.com/faq/api?title=webapi/guide/webservice-geocoding-base
func (mcg *MapGeocoding) V3(ctx context.Context, address string, notMustParams ...gorequest.Params) (*MapGeocodingResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("ak", mcg.mc.ak)
	params.Set("address", address)
	params.Set("output", "json")
	// 请求
	request, err := mcg.mc.request(ctx, "/geocoding/v3/", params, http.MethodGet)
	if err != nil {
		return newMapGeocodingResult(MapGeocodingResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response MapGeocodingResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newMapGeocodingResult(response, request.ResponseBody, request), err
}
