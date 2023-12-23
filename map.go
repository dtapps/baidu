package baidu

import (
	"context"
	"go.dtapp.net/golog"
)

type MapClient struct {
	ak      string
	gormLog struct {
		status bool           // 状态
		client *golog.ApiGorm // 日志服务
	}
}

func NewMapClient(ctx context.Context, ak string) (*MapClient, error) {
	return &MapClient{ak: ak}, nil
}
