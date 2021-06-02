package service

import (
	"context"
	"music/app/model"
)

type sessionService struct{}

var Session = sessionService{}

const (
	sessionKeyUser = "SessionKeyUser"
)

func SetUser(ctx context.Context, user *model.User) error {
	return Context.Get(ctx).Session.Set(sessionKeyUser, user)
}

func (s *sessionService) getUser(ctx context.Context) *model.User {
	customCtx := Context.Get(ctx)
	if customCtx != nil {
		if v := customCtx.Session.GetVar(sessionKeyUser); !v.IsNil() {
			var user *model.User
			_ = v.Struct(&user)
			return user
		}
	}
	return nil
}

func (s *sessionService) RemoveUser(ctx context.Context) error {
	customCtx := Context.Get(ctx)
	if Context.Get(ctx) != nil {
		return customCtx.Session.Remove(sessionKeyUser)
	}
	return nil
}
