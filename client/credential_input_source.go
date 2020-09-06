package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type CredentialInputSourceService struct {
	client *Client
}

type ListCredentialInputSourceResponse struct {
	Pagination
	Results []*CredentialInputSource `json:"results"`
}

const credentialInputSourceAPIEndpoint = "/api/v2/credential_input_sources/"

func (cs *CredentialInputSourceService) ListCredentialInputSources(params map[string]string) ([]*CredentialInputSource,
	*ListCredentialInputSourceResponse,
	error) {
	result := new(ListCredentialInputSourceResponse)
	resp, err := cs.client.Requester.GetJSON(credentialInputSourceAPIEndpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

func (cs *CredentialInputSourceService) CreateCredentialInputSource(data map[string]interface{}, params map[string]string) (*CredentialInputSource, error) {
	result := new(CredentialInputSource)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := cs.client.Requester.PostJSON(credentialInputSourceAPIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (cs *CredentialInputSourceService) GetCredentialInputSourceByID(id int, params map[string]string) (*CredentialInputSource, error) {
	result := new(CredentialInputSource)
	endpoint := fmt.Sprintf("%s%d", credentialInputSourceAPIEndpoint, id)
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

func (cs *CredentialInputSourceService) UpdateCredentialInputSourceByID(id int, data map[string]interface{},
	params map[string]string) (*CredentialInputSource, error) {
	result := new(CredentialInputSource)
	endpoint := fmt.Sprintf("%s%d", credentialInputSourceAPIEndpoint, id)

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

func (cs *CredentialInputSourceService) DeleteCredentialInputSourceByID(id int, params map[string]string) error {
	endpoint := fmt.Sprintf("%s%d", credentialInputSourceAPIEndpoint, id)
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
