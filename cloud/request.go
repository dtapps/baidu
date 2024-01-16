package cloud

import (
	"context"
	"go.dtapp.net/gorequest"
)

func (c *Client) request(ctx context.Context, url string, param gorequest.Params, method string, contentType string) (gorequest.Response, error) {

	// 创建请求
	client := gorequest.NewHttp()

	// 设置请求地址
	client.SetUri(apiUrl + url)

	// 设置方式
	client.SetMethod(method)

	// 设置格式
	switch contentType {
	case "JSON":
		client.SetContentTypeJson()
	case "FORM":
		client.SetContentTypeForm()
	}

	// 设置参数
	client.SetParams(param)

	// 发起请求
	request, err := client.Request(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 记录日志
	if c.gormLog.status {
		go c.gormLog.client.Middleware(ctx, request)
	}
	if c.mongoLog.status {
		go c.mongoLog.client.Middleware(ctx, request)
	}

	return request, err
}
