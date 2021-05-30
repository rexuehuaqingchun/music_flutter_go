package service

type sessionService struct{}

var Session = sessionService{}

const (
	sessionKeyUser = "SessionKeyUser"
)
