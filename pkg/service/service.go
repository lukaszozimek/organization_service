package service

import (
	"context"
	"organization_service/pkg/model"
)

// OrganizationServiceService describes the service.
type OrganizationServiceService interface {
	CreateUserOrganizationById(ctx context.Context, organization model.Organization) (organization model.Organization, err error)
	DeleteUserOrganizationById(ctx context.Context, organizationId string) (organization model.Organization, err error)
	GetUserOrganizationById(ctx context.Context, organizationId string) (organization model.Organization, err error)
	GetUserOrganizations(ctx context.Context) (organization []model.Organization, err error)
	UpdateUserOrganizationById(ctx context.Context, organization model.Organization) (organization model.Organization, err error)
}

type basicOrganizationServiceService struct{}

func (b *basicOrganizationServiceService) CreateUserOrganizationById(ctx context.Context, organization model.Organization) (organization model.Organization, err error) {
	// TODO implement the business logic of CreateUserOrganizationById
	return organization, err
}
func (b *basicOrganizationServiceService) DeleteUserOrganizationById(ctx context.Context, organizationId string) (organization model.Organization, err error) {
	// TODO implement the business logic of DeleteUserOrganizationById
	return organization, err
}
func (b *basicOrganizationServiceService) GetUserOrganizationById(ctx context.Context, organizationId string) (organization model.Organization, err error) {
	// TODO implement the business logic of GetUserOrganizationById
	return organization, err
}
func (b *basicOrganizationServiceService) GetUserOrganizations(ctx context.Context) (organization []model.Organization, err error) {
	// TODO implement the business logic of GetUserOrganizations
	return organization, err
}
func (b *basicOrganizationServiceService) UpdateUserOrganizationById(ctx context.Context, organization model.Organization) (organization model.Organization, err error) {
	// TODO implement the business logic of UpdateUserOrganizationById
	return organization, err
}

// NewBasicOrganizationServiceService returns a naive, stateless implementation of OrganizationServiceService.
func NewBasicOrganizationServiceService() OrganizationServiceService {
	return &basicOrganizationServiceService{}
}

// New returns a OrganizationServiceService with all of the expected middleware wired in.
func New(middleware []Middleware) OrganizationServiceService {
	var svc OrganizationServiceService = NewBasicOrganizationServiceService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
