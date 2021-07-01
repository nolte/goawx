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
	Results []*Team `json:"results"`
}

const teamsAPIEndpoint = "/api/v2/teams/"

// GetTeamByID shows the details of an awx team.
func (i *TeamsService) GetTeamByID(id int, params map[string]string) (*Team, error) {
	result := new(Team)
	endpoint := fmt.Sprintf("%s%d/", teamsAPIEndpoint, id)
	resp, err := i.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// ListTeams shows list of awx teams.
func (i *TeamsService) ListTeams(params map[string]string) ([]*Team, *ListTeamsResponse, error) {
	result := new(ListTeamsResponse)
	resp, err := i.client.Requester.GetJSON(teamsAPIEndpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

// CreateTeam creates an awx team.
func (i *TeamsService) CreateTeam(data map[string]interface{}, params map[string]string) (*Team, error) {
	mandatoryFields = []string{"name", "organization"}
	validate, status := ValidateParams(data, mandatoryFields)

	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}

	result := new(Team)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Add check if team exists and return proper error

	resp, err := i.client.Requester.PostJSON(teamsAPIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateTeam update an awx team
func (i *TeamsService) UpdateTeam(id int, data map[string]interface{}, params map[string]string) (*Team, error) {
	result := new(Team)
	endpoint := fmt.Sprintf("%s%d", teamsAPIEndpoint, id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := i.client.Requester.PatchJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// GetTeam retrives the team information from its ID or Name
func (i *TeamsService) GetTeam(id int, params map[string]string) (*Team, error) {
	endpoint := fmt.Sprintf("%s%d", teamsAPIEndpoint, id)
	result := new(Team)
	resp, err := i.client.Requester.GetJSON(endpoint, result, map[string]string{})
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteTeam delete an team from AWX
func (i *TeamsService) DeleteTeam(id int) (*Team, error) {
	result := new(Team)
	endpoint := fmt.Sprintf("%s%d", teamsAPIEndpoint, id)

	resp, err := i.client.Requester.Delete(endpoint, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
