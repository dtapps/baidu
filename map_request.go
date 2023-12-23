package baidu

import (
	"context"
	"go.dtapp.net/gorequest"
)

func (mc *MapClient) request(ctx context.Context, url string, param gorequest.Params, method string) (gorequest.Response, error) {

	// 创建请求
	client := gorequest.NewHttp()

	// 设置请求地址
	client.SetUri(MapApiUrl + url)

	// 设置方式
	client.SetMethod(method)

	// 设置格式
	client.SetContentTypeJson()

	// 设置参数
	client.SetParams(param)

	// 发起请求
	request, err := client.Request(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 记录日志
	if mc.gormLog.status {
		go mc.gormLog.client.Middleware(ctx, request)
	}

	return request, err
}
