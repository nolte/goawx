package awx

import (
	"fmt"
	"net/http"
)

// This variable is mandatory and to be populated for creating services API
var mandatoryFields []string

// AWX represents awx api endpoints with services, and using
// client to communicate with awx server.
type AWX struct {
	client *Client

	PingService                           *PingService
	InventoriesService                    *InventoriesService
	JobService                            *JobService
	JobTemplateService                    *JobTemplateService
	ProjectService                        *ProjectService
	ProjectUpdatesService                 *ProjectUpdatesService
	UserService                           *UserService
	GroupService                          *GroupService
	HostService                           *HostService
	CredentialsService                    *CredentialsService
	CredentialInputSourceService          *CredentialInputSourceService
	InventorySourcesService               *InventorySourcesService
	InventoryGroupService                 *InventoryGroupService
	OrganizationsService                  *OrganizationsService
	WorkflowJobTemplateService            *WorkflowJobTemplateService
	WorkflowJobTemplateNodeService        *WorkflowJobTemplateNodeService
	WorkflowJobTemplateNodeAllwaysService *WorkflowJobTemplateNodeStepService
	WorkflowJobTemplateNodeFailureService *WorkflowJobTemplateNodeStepService
	WorkflowJobTemplateNodeSuccessService *WorkflowJobTemplateNodeStepService
}

// Client implement http client.
type Client struct {
	BaseURL   string
	Requester *Requester
}

// CheckResponse do http response check, and return err if not in [200, 300).
func CheckResponse(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	return fmt.Errorf("responsed with %d, resp: %v", resp.StatusCode, resp)
}

// ValidateParams is to validate the input to use the services.
func ValidateParams(data map[string]interface{}, mandatoryFields []string) (notfound []string, status bool) {
	status = true
	for _, key := range mandatoryFields {
		_, exists := data[key]

		if !exists {
			notfound = append(notfound, key)
			status = false
		}
	}
	return notfound, status
}

// NewAWX news an awx handler with basic auth support, you could customize the http
// transport by passing custom client.
func NewAWX(baseURL, userName, passwd string, client *http.Client) (*AWX, error) {
	r := &Requester{Base: baseURL, BasicAuth: &BasicAuth{Username: userName, Password: passwd}, Client: client}
	if r.Client == nil {
		r.Client = http.DefaultClient
	}

	awxClient := &Client{
		BaseURL:   baseURL,
		Requester: r,
	}

	newAWX := &AWX{
		client: awxClient,

		PingService: &PingService{
			client: awxClient,
		},
		InventoriesService: &InventoriesService{
			client: awxClient,
		},
		JobService: &JobService{
			client: awxClient,
		},
		JobTemplateService: &JobTemplateService{
			client: awxClient,
		},
		ProjectService: &ProjectService{
			client: awxClient,
		},
		ProjectUpdatesService: &ProjectUpdatesService{
			client: awxClient,
		},
		UserService: &UserService{
			client: awxClient,
		},
		GroupService: &GroupService{
			client: awxClient,
		},
		HostService: &HostService{
			client: awxClient,
		},
		CredentialsService: &CredentialsService{
			client: awxClient,
		},
		CredentialInputSourceService: &CredentialInputSourceService{
			client: awxClient,
		},
		InventorySourcesService: &InventorySourcesService{
			client: awxClient,
		},
		InventoryGroupService: &InventoryGroupService{
			client: awxClient,
		},
		OrganizationsService: &OrganizationsService{
			client: awxClient,
		},
		WorkflowJobTemplateService: &WorkflowJobTemplateService{
			client: awxClient,
		},
		WorkflowJobTemplateNodeService: &WorkflowJobTemplateNodeService{
			client: awxClient,
		},
		WorkflowJobTemplateNodeSuccessService: &WorkflowJobTemplateNodeStepService{
			endpoint: fmt.Sprintf("%s%s", workflowJobTemplateNodeAPIEndpoint, "%d/success_nodes/"),
			client:   awxClient,
		},
		WorkflowJobTemplateNodeFailureService: &WorkflowJobTemplateNodeStepService{
			endpoint: fmt.Sprintf("%s%s", workflowJobTemplateNodeAPIEndpoint, "%d/failure_nodes/"),
			client:   awxClient,
		},
		WorkflowJobTemplateNodeAllwaysService: &WorkflowJobTemplateNodeStepService{
			endpoint: fmt.Sprintf("%s%s", workflowJobTemplateNodeAPIEndpoint, "%d/allways_nodes/"),
			client:   awxClient,
		},
	}

	// test the connection and return and error if there's an issue
	_, err := newAWX.PingService.Ping()
	if err != nil {
		return nil, err
	}

	return newAWX, nil
}
