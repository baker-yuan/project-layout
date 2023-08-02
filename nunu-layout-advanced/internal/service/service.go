package service

import (
	"nunu-layout-advanced/internal/middleware"
	"nunu-layout-advanced/pkg/helper/sid"
	"nunu-layout-advanced/pkg/log"
)

type Service struct {
	logger *log.Logger
	sid    *sid.Sid
	jwt    *middleware.JWT
}

func NewService(logger *log.Logger, sid *sid.Sid, jwt *middleware.JWT) *Service {
	return &Service{
		logger: logger,
		sid:    sid,
		jwt:    jwt,
	}
}
