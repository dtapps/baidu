package baidu

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
	"net/http"
)

type MapReverseGeocodingResponse struct {
	Status int `json:"status"`
	Result struct {
		Location struct {
			Lng float64 `json:"lng"`
			Lat float64 `json:"lat"`
		} `json:"location"`
		FormattedAddress string `json:"formatted_address"`
		Business         string `json:"business"`
		AddressComponent struct {
			Country         string `json:"country"`
			CountryCode     int    `json:"country_code"`
			CountryCodeIso  string `json:"country_code_iso"`
			CountryCodeIso2 string `json:"country_code_iso2"`
			Province        string `json:"province"`
			City            string `json:"city"`
			CityLevel       int    `json:"city_level"`
			District        string `json:"district"`
			Town            string `json:"town"`
			TownCode        string `json:"town_code"`
			Distance        string `json:"distance"`
			Direction       string `json:"direction"`
			Adcode          string `json:"adcode"`
			Street          string `json:"street"`
			StreetNumber    string `json:"street_number"`
		} `json:"addressComponent"`
		Pois               []interface{} `json:"pois"`
		Roads              []interface{} `json:"roads"`
		PoiRegions         []interface{} `json:"poiRegions"`
		SematicDescription string        `json:"sematic_description"`
		CityCode           int           `json:"cityCode"`
	} `json:"result"`
}

type MapReverseGeocodingResult struct {
	Result MapReverseGeocodingResponse // 结果
	Body   []byte                      // 内容
	Http   gorequest.Response          // 请求
}

func newMapReverseGeocodingResult(result MapReverseGeocodingResponse, body []byte, http gorequest.Response) *MapReverseGeocodingResult {
	return &MapReverseGeocodingResult{Result: result, Body: body, Http: http}
}

// V3 全球逆地理编码
// https://lbsyun.baidu.com/faq/api?title=webapi/guide/webservice-geocoding-abroad-base
func (mcg *MapReverseGeocoding) V3(ctx context.Context, location string, notMustParams ...gorequest.Params) (*MapReverseGeocodingResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("ak", mcg.mc.ak)
	params.Set("location", location)
	params.Set("output", "json")
	// 请求
	request, err := mcg.mc.request(ctx, "/reverse_geocoding/v3/", params, http.MethodGet)
	if err != nil {
		return newMapReverseGeocodingResult(MapReverseGeocodingResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response MapReverseGeocodingResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newMapReverseGeocodingResult(response, request.ResponseBody, request), err
}
