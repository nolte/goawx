package awx

import (
	"fmt"
)

// InventoryGroupService implements awx inventory group apis.
type InventoryGroupService struct {
	client *Client
}

// ListInventoryGroups shows list of awx groups in some inventory.
func (i *InventoryGroupService) ListInventoryGroups(id int, params map[string]string) ([]*Group, *ListGroupsResponse, error) {
	result := new(ListGroupsResponse)
	endpoint := fmt.Sprintf("%s%d/groups/", inventoriesAPIEndpoint, id)
	resp, err := i.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}
