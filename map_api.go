package baidu

// MapGeocoding 地理编码
// https://lbsyun.baidu.com/faq/api?title=webapi/guide/webservice-geocoding
type MapGeocoding struct {
	mc *MapClient
}

func (mc *MapClient) Geocoding() *MapGeocoding {
	return &MapGeocoding{mc: mc}
}

// MapReverseGeocoding 全球逆地理编码
// https://lbsyun.baidu.com/faq/api?title=webapi/guide/webservice-geocoding-abroad
type MapReverseGeocoding struct {
	mc *MapClient
}

func (mc *MapClient) ReverseGeocoding() *MapReverseGeocoding {
	return &MapReverseGeocoding{mc: mc}
}

// MapWeather 天气查询
// https://lbsyun.baidu.com/faq/api?title=webapi/weather
type MapWeather struct {
	mc *MapClient
}

func (mc *MapClient) Weather() *MapWeather {
	return &MapWeather{mc: mc}
}

// MapLocation 普通IP定位
// https://lbsyun.baidu.com/faq/api?title=webapi/ip-api
type MapLocation struct {
	mc *MapClient
}

func (mc *MapClient) Location() *MapLocation {
	return &MapLocation{mc: mc}
}
