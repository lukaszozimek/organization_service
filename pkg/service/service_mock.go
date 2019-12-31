package service

import (
	"context"
	"github.com/lukaszozimek/organization_service/pkg/model"
)

type basicOrganizationServiceMock struct{}

func (b *basicOrganizationServiceMock) CreateUserOrganizationById(ctx context.Context, organization model.Organization) (res model.Organization, err error) {
	return organization, err
}
func (b *basicOrganizationServiceMock) DeleteUserOrganizationById(ctx context.Context, organizationId string) (res model.Organization, err error) {
	var organization model.Organization
	return organization, err
}
func (b *basicOrganizationServiceMock) GetUserOrganizationById(ctx context.Context, organizationId string) (res model.Organization, err error) {
	return res, err
}
func (b *basicOrganizationServiceMock) GetUserOrganizations(ctx context.Context) (organization []model.Organization, err error) {
	var organizations []model.Organization
	return organizations, nil
}
func (b *basicOrganizationServiceMock) UpdateUserOrganizationById(ctx context.Context, organization model.Organization) (res model.Organization, err error) {
	return organization, err
}
func (b *basicOrganizationServiceMock) Health(ctx context.Context) (rs string, err error) {
	return "true", err
}

// NewBasicOrganizationServiceService returns a naive, stateless implementation of OrganizationService.
func NewBasicOrganizationServiceServiceMock() OrganizationService {
	return &basicOrganizationServiceMock{}
}

// New returns a OrganizationService with all of the expected middleware wired in.
func NewMock(middleware []Middleware) OrganizationService {
	var svc OrganizationService = NewBasicOrganizationServiceServiceMock()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
