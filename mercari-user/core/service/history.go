package service

import (
	"context"
	"sagungw/mercari/core/datastore"
)

type LoginHistoryService interface {
	GetLoginHistory(ctx context.Context, userID string, page int64) ([]*datastore.LoginHistory, error)
}

type loginHistoryService struct {
	loginHistoryDatastore datastore.HistoryDatastore
}

func NewLoginHistoryService(loginHistoryDatastore datastore.HistoryDatastore) *loginHistoryService {
	return &loginHistoryService{
		loginHistoryDatastore: loginHistoryDatastore,
	}
}

func (l *loginHistoryService) GetLoginHistory(ctx context.Context, userID string, page int64) ([]*datastore.LoginHistory, error) {
	return l.loginHistoryDatastore.GetLoginHistory(ctx, userID, page)
}
