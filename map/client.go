package _map

import (
	"context"
	"go.dtapp.net/golog"
)

type Client struct {
	ak      string
	gormLog struct {
		status bool           // 状态
		client *golog.ApiGorm // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(ctx context.Context, ak string) (*Client, error) {
	return &Client{ak: ak}, nil
}
