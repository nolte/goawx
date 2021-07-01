package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// TeamsService implements awx teams apis.
type TeamsService struct {
	client *Client
}

// ListTeamsResponse represents `ListTeams` endpoint response.
type ListTeamsResponse struct {
	Pagination
	Results []*Teams `json:"results"`
}

const teamsAPIEndpoint = "/api/v2/teams/"

// GetTeamByID shows the details of an awx team.
func (t *TeamsService) GetTeamByID(id int, params map[string]string) (*Teams, error) {
	result := new(Teams)
	endpoint := fmt.Sprintf("%s%d/", teamsAPIEndpoint, id)
	resp, err := t.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// ListTeams shows list of awx teams.
func (t *TeamsService) ListTeams(params map[string]string) ([]*Teams, *ListTeamsResponse, error) {
	result := new(ListTeamsResponse)
	resp, err := t.client.Requester.GetJSON(teamsAPIEndpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

// CreateTeam creates an awx team.
func (t *TeamsService) CreateTeam(data map[string]interface{}, params map[string]string) (*Teams, error) {
	mandatoryFields = []string{"name", "organization"}
	validate, status := ValidateParams(data, mandatoryFields)

	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}

	result := new(Teams)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Add check if team exists and return proper error

	resp, err := t.client.Requester.PostJSON(teamsAPIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateTeam update an awx team
func (t *TeamsService) UpdateTeam(id int, data map[string]interface{}, params map[string]string) (*Teams, error) {
	result := new(Teams)
	endpoint := fmt.Sprintf("%s%d", teamsAPIEndpoint, id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := t.client.Requester.PatchJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// GetTeam retrives the team information from its ID or Name
func (t *TeamsService) GetTeam(id int, params map[string]string) (*Teams, error) {
	endpoint := fmt.Sprintf("%s%d", teamsAPIEndpoint, id)
	result := new(Teams)
	resp, err := t.client.Requester.GetJSON(endpoint, result, map[string]string{})
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteTeam delete an team from AWX
func (t *TeamsService) DeleteTeam(id int) (*Teams, error) {
	result := new(Teams)
	endpoint := fmt.Sprintf("%s%d", teamsAPIEndpoint, id)

	resp, err := t.client.Requester.Delete(endpoint, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
