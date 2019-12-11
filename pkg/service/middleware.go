package service

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/lukaszozimek/organization_service/pkg/model"
)

// Middleware describes a service middleware.
type Middleware func(OrganizationService) OrganizationService

type loggingMiddleware struct {
	logger log.Logger
	next   OrganizationService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a OrganizationService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next OrganizationService) OrganizationService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) CreateUserOrganizationById(ctx context.Context, organization model.Organization) (organization model.Organization, err error) {
	defer func() {
		l.logger.Log("method", "CreateUserOrganizationById", "organization", organization, "organization", organization, "err", err)
	}()
	return l.next.CreateUserOrganizationById(ctx, organization)
}
func (l loggingMiddleware) DeleteUserOrganizationById(ctx context.Context, organizationId string) (organization model.Organization, err error) {
	defer func() {
		l.logger.Log("method", "DeleteUserOrganizationById", "organizationId", organizationId, "organization", organization, "err", err)
	}()
	return l.next.DeleteUserOrganizationById(ctx, organizationId)
}
func (l loggingMiddleware) GetUserOrganizationById(ctx context.Context, organizationId string) (organization model.Organization, err error) {
	defer func() {
		l.logger.Log("method", "GetUserOrganizationById", "organizationId", organizationId, "organization", organization, "err", err)
	}()
	return l.next.GetUserOrganizationById(ctx, organizationId)
}
func (l loggingMiddleware) GetUserOrganizations(ctx context.Context) (organization []model.Organization, err error) {
	defer func() {
		l.logger.Log("method", "GetUserOrganizations", "organization", organization, "err", err)
	}()
	return l.next.GetUserOrganizations(ctx)
}
func (l loggingMiddleware) UpdateUserOrganizationById(ctx context.Context, organization model.Organization) (organization model.Organization, err error) {
	defer func() {
		l.logger.Log("method", "UpdateUserOrganizationById", "organization", organization, "organization", organization, "err", err)
	}()
	return l.next.UpdateUserOrganizationById(ctx, organization)
}
