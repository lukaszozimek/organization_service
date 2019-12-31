package endpoint_test

import (
	"github.com/lukaszozimek/organization_service/pkg/endpoint"
	"github.com/lukaszozimek/organization_service/pkg/service"
	"testing"
)

var mock = service.NewBasicOrganizationServiceServiceMock()

func TestMakeCreateUserOrganizationByIdEndpoint(t *testing.T) {

	endpoint.MakeCreateUserOrganizationByIdEndpoint(mock)
}

func TestMakeDeleteUserOrganizationByIdEndpoint(t *testing.T) {

	endpoint.MakeDeleteUserOrganizationByIdEndpoint(mock)
}

func TestMakeGetUserOrganizationByIdEndpoint(t *testing.T) {

	endpoint.MakeGetUserOrganizationByIdEndpoint(mock)
}

func TestMakeGetUserOrganizationsEndpoint(t *testing.T) {

	endpoint.MakeGetUserOrganizationsEndpoint(mock)
}

func TestMakeUpdateUserOrganizationByIdEndpoint(t *testing.T) {
	endpoint.MakeUpdateUserOrganizationByIdEndpoint(mock)
}

func TestMakeHealthEndpoint(t *testing.T) {
	endpoint.MakeHealthEndpoint(mock)
}
