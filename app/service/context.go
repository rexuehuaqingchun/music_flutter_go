package service

import (
	"context"
	"github.com/gogf/gf/net/ghttp"
	"music/app/model"
)

//这里定义一个空的结构体，用来后面进行关联方法 方法即为某一结构体的可执行函数
type contextService struct{}

// Context 一个空的结构体实例  作为service的上下文实例
var Context = contextService{}

// Init 此方法的作用是将model的上下文用model中key注册到request的上下文对象中 通过这种方法可以通过
//service实例Context拿到model对象
func (s *contextService) Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(model.ContextKey, customCtx)
}

func (s *contextService) Get(ctx context.Context) *model.Context {
	value := ctx.Value(model.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}
