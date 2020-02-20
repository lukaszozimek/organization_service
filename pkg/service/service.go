package service

import (
	"context"
	"github.com/lukaszozimek/organization_service/pkg/model"
	"strconv"
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
	db := model.GetDB()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return model.Organization{}, err
	}

	if err := tx.Create(&organization).Error; err != nil {
		tx.Rollback()
		return model.Organization{}, err
	}

	tx.Commit()
	return organization, err
}
func (b *basicOrganizationService) DeleteUserOrganizationById(ctx context.Context, organizationId string) (res model.Organization, err error) {
	db := model.GetDB()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return model.Organization{}, err
	}
	var organization model.Organization
	tx.Where("ID = ?", organizationId).Find(&organization)
	if err := tx.Delete(organization).Error; err != nil {
		tx.Rollback()
		return model.Organization{}, err
	}
	tx.Commit()
	return model.Organization{}, err
}
func (b *basicOrganizationService) GetUserOrganizationById(ctx context.Context, organizationId string) (res model.Organization, err error) {
	db := model.GetDB()
	i, err := strconv.ParseInt(organizationId, 10, 64)
	var organization model.Organization
	db.First(&organization, i)
	return res, err
}
func (b *basicOrganizationService) GetUserOrganizations(ctx context.Context) (organization []model.Organization, err error) {
	db := model.GetDB()
	var organizations []model.Organization
	db.Find(&organizations)
	return organizations, nil
}
func (b *basicOrganizationService) UpdateUserOrganizationById(ctx context.Context, organization model.Organization) (res model.Organization, err error) {
	db := model.GetDB()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return model.Organization{}, err
	}
	if err := tx.Update(&organization).Error; err != nil {
		tx.Rollback()
		return model.Organization{}, err
	}
	tx.Commit()
	return organization, err
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
