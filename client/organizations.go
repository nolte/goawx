package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// OrganizationsService implements awx organizations apis.
type OrganizationsService struct {
	client *Client
}

// ListOrganizationsResponse represents `ListOrganizations` endpoint response.
type ListOrganizationsResponse struct {
	Pagination
	Results []*Organizations `json:"results"`
}

const organizationsAPIEndpoint = "/api/v2/organizations/"

// ListOrganizations shows list of awx organizations.
func (p *OrganizationsService) ListOrganizations(params map[string]string) ([]*Organizations, *ListOrganizationsResponse, error) {
	result := new(ListOrganizationsResponse)
	resp, err := p.client.Requester.GetJSON(organizationsAPIEndpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

// GetOrganizationsByID shows the details of a Organization.
func (p *OrganizationsService) GetOrganizationsByID(id int, params map[string]string) (*Organizations, error) {
	result := new(Organizations)
	endpoint := fmt.Sprintf("%s%d/", organizationsAPIEndpoint, id)
	resp, err := p.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// CreateOrganization creates an awx Organization.
func (p *OrganizationsService) CreateOrganization(data map[string]interface{}, params map[string]string) (*Organizations, error) {
	mandatoryFields = []string{"name"}
	validate, status := ValidateParams(data, mandatoryFields)

	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}

	result := new(Organizations)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Requester.PostJSON(organizationsAPIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateOrganization update an awx Organization.
func (p *OrganizationsService) UpdateOrganization(id int, data map[string]interface{}, params map[string]string) (*Organizations, error) {
	result := new(Organizations)
	endpoint := fmt.Sprintf("%s%d", organizationsAPIEndpoint, id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Requester.PatchJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteOrganization delete an awx Organization.
func (p *OrganizationsService) DeleteOrganization(id int) (*Organizations, error) {
	result := new(Organizations)
	endpoint := fmt.Sprintf("%s%d", organizationsAPIEndpoint, id)

	resp, err := p.client.Requester.Delete(endpoint, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
