package service

import (
	"context"
)

type SiteMock struct {
	CreateFunc func(ctx context.Context, s *AddParams) (*Site, error)
	DeleteFunc func(ctx context.Context, siteID int) error
	GetFunc    func(ctx context.Context, siteID int) (*Site, error)
	ListFunc   func(ctx context.Context) (*ListResponse, error)
}

func (m *SiteMock) Add(ctx context.Context, s *AddParams) (*Site, error) {
	return m.CreateFunc(ctx, s)
}

func (m *SiteMock) Delete(ctx context.Context, siteID int) error {
	return m.DeleteFunc(ctx, siteID)
}

func (m *SiteMock) Get(ctx context.Context, siteID int) (*Site, error) {
	return m.GetFunc(ctx, siteID)
}

func (m *SiteMock) List(ctx context.Context) (*ListResponse, error) {
	return m.ListFunc(ctx)
}
