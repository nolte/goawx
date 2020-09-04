package awx

// OrganizationsService implements awx projects apis.
type OrganizationsService struct {
	client *Client
}

// ListOrganizationsResponse represents `ListProjects` endpoint response.
type ListOrganizationsResponse struct {
	Pagination
	Results []*Organizations `json:"results"`
}

const organizationsAPIEndpoint = "/api/v2/organizations/"

// ListOrganizations shows list of awx projects.
func (p *OrganizationsService) ListOrganizations(params map[string]string) ([]*Organizations, *ListOrganizationsResponse, error) {
	result := new(ListOrganizationsResponse)
	resp, err := p.client.Requester.GetJSON(projectsAPIEndpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}
