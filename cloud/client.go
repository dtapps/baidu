package cloud

import (
	"context"
	"go.dtapp.net/golog"
)

type Client struct {
	apiKey      string
	secretKey   string
	accessToken string
	gormLog     struct {
		status bool           // 状态
		client *golog.ApiGorm // 日志服务
	}
	mongoLog struct {
		status bool            // 状态
		client *golog.ApiMongo // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(ctx context.Context, apiKey string, secretKey string) (*Client, error) {
	return &Client{apiKey: apiKey, secretKey: secretKey}, nil
}
