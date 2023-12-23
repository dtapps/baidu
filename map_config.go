package baidu

import "go.dtapp.net/golog"

const (
	MapApiUrl   = "https://api.map.baidu.com" // 	接口地址
	MapLogTable = "baidu"                     // 日志表
)

const ()

// ConfigApiGormFun 接口日志配置
func (mc *MapClient) ConfigApiGormFun(apiClientFun golog.ApiGormFun) {
	client := apiClientFun()
	if client != nil {
		mc.gormLog.client = client
		mc.gormLog.status = true
	}
}
