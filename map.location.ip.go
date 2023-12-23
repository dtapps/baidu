package baidu

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
	"net/http"
)

type MapLocationIpResponse struct {
	Address string `json:"address"` // 详细地址信息
	Content struct {
		AddressDetail struct {
			Province     string `json:"province"` // 省份
			City         string `json:"city"`     // 城市
			District     string `json:"district"`
			Street       string `json:"street"`
			StreetNumber string `json:"street_number"`
			CityCode     int64  `json:"city_code"` // 百度城市代码
			Adcode       string `json:"adcode"`
		} `json:"address_detail"`
		Address string `json:"address"` // 简要地址信息
		Point   struct {
			X string `json:"x"` // 当前城市中心点经度
			Y string `json:"y"` // 当前城市中心点纬度
		} `json:"point"`
	} `json:"content"`
	Status int `json:"status"`
}

type MapLocationIpResult struct {
	Result MapLocationIpResponse // 结果
	Body   []byte                // 内容
	Http   gorequest.Response    // 请求
}

func newMapLocationIpResult(result MapLocationIpResponse, body []byte, http gorequest.Response) *MapLocationIpResult {
	return &MapLocationIpResult{Result: result, Body: body, Http: http}
}

// Ip 普通IP定位
// https://lbsyun.baidu.com/faq/api?title=webapi/ip-api-base
func (mcl *MapLocation) Ip(ctx context.Context, ip string, notMustParams ...gorequest.Params) (*MapLocationIpResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("ak", mcl.mc.ak)
	params.Set("ip", ip)
	// 请求
	request, err := mcl.mc.request(ctx, "/location/ip", params, http.MethodGet)
	if err != nil {
		return newMapLocationIpResult(MapLocationIpResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response MapLocationIpResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newMapLocationIpResult(response, request.ResponseBody, request), err
}
