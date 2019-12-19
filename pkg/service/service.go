package service

import (
	"context"
	"github.com/lukaszozimek/organization_service/pkg/model"
)

// OrganizationService describes the service.
type OrganizationService interface {
	CreateUserOrganizationById(ctx context.Context, organization model.Organization) (res model.Organization, err error)
	DeleteUserOrganizationById(ctx context.Context, organizationId string) (res model.Organization, err error)
	GetUserOrganizationById(ctx context.Context, organizationId string) (res model.Organization, err error)
	GetUserOrganizations(ctx context.Context) (organization []model.Organization, err error)
	UpdateUserOrganizationById(ctx context.Context, organization model.Organization) (res model.Organization, err error)
	Health(ctx context.Context) (rs string, err error)
}

type basicOrganizationService struct{}

func (b *basicOrganizationService) CreateUserOrganizationById(ctx context.Context, organization model.Organization) (res model.Organization, err error) {
	// TODO implement the business logic of CreateUserOrganizationById
	return organization, err
}
func (b *basicOrganizationService) DeleteUserOrganizationById(ctx context.Context, organizationId string) (res model.Organization, err error) {
	// TODO implement the business logic of DeleteUserOrganizationById
	return res, err
}
func (b *basicOrganizationService) GetUserOrganizationById(ctx context.Context, organizationId string) (res model.Organization, err error) {
	// TODO implement the business logic of GetUserOrganizationById
	return res, err
}
func (b *basicOrganizationService) GetUserOrganizations(ctx context.Context) (organization []model.Organization, err error) {
	// TODO implement the business logic of GetUserOrganizations
	return organization, err
}
func (b *basicOrganizationService) UpdateUserOrganizationById(ctx context.Context, organization model.Organization) (res model.Organization, err error) {
	// TODO implement the business logic of UpdateUserOrganizationById
	return res, err
}
func (b *basicOrganizationService) Health(ctx context.Context) (rs string, err error) {
	return "true", err
}

// NewBasicOrganizationServiceService returns a naive, stateless implementation of OrganizationService.
func NewBasicOrganizationServiceService() OrganizationService {
	return &basicOrganizationService{}
}

// New returns a OrganizationService with all of the expected middleware wired in.
func New(middleware []Middleware) OrganizationService {
	var svc OrganizationService = NewBasicOrganizationServiceService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
