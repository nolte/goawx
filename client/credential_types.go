package awx

type CredentialTypeService struct {
	client *Client
}

type ListCredentialTypeResponse struct {
	Pagination
	Results []*CredentialType `json:"results"`
}

const credentialTypesAPIEndpoint = "/api/v2/credential_types/"

func (cs *CredentialTypeService) ListCredentialsTypes(params map[string]string) ([]*CredentialType,
	*ListCredentialTypeResponse, error) {
	result := new(ListCredentialTypeResponse)
	resp, err := cs.client.Requester.GetJSON(credentialTypesAPIEndpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}
