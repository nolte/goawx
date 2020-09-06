package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type CredentialsService struct {
	client *Client
}

type ListCredentialsResponse struct {
	Pagination
	Results []*Credential `json:"results"`
}

const credentialsAPIEndpoint = "/api/v2/credentials/"

func (cs *CredentialsService) ListCredentials(params map[string]string) ([]*Credential,
	*ListCredentialsResponse,
	error) {
	result := new(ListCredentialsResponse)
	resp, err := cs.client.Requester.GetJSON(credentialsAPIEndpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

func (cs *CredentialsService) CreateCredentials(data map[string]interface{}, params map[string]string) (*Credential, error) {
	result := new(Credential)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := cs.client.Requester.PostJSON(credentialsAPIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (cs *CredentialsService) GetCredentialsByID(id int, params map[string]string) (*Credential, error) {
	result := new(Credential)
	endpoint := fmt.Sprintf("%s%d", credentialsAPIEndpoint, id)
	resp, err := cs.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (cs *CredentialsService) UpdateCredentialsByID(id int, data map[string]interface{},
	params map[string]string) (*Credential, error) {
	result := new(Credential)
	endpoint := fmt.Sprintf("%s%d", credentialsAPIEndpoint, id)

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := cs.client.Requester.PatchJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (cs *CredentialsService) DeleteCredentialsByID(id int, params map[string]string) error {
	endpoint := fmt.Sprintf("%s%d", credentialsAPIEndpoint, id)
	resp, err := cs.client.Requester.Delete(endpoint, nil, params)
	if err != nil {
		return err
	}

	err = CheckResponse(resp)
	if err != nil {
		return err
	}

	return nil
}
