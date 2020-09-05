package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// JobTemplateService implements awx job template apis.
type WorkflowJobTemplateNodeService struct {
	client *Client
}

// ListJobTemplatesResponse represents `ListJobTemplates` endpoint response.
type ListWorkflowJobTemplateNodesResponse struct {
	Pagination
	Results []*WorkflowJobTemplateNode `json:"results"`
}

const workflowJobTemplateNodeAPIEndpoint = "/api/v2/workflow_job_template_nodes/"

// GetJobTemplateByID shows the details of a job template.
func (jt *WorkflowJobTemplateNodeService) GetWorkflowJobTemplateNodeByID(id int, params map[string]string) (*WorkflowJobTemplateNode, error) {
	result := new(WorkflowJobTemplateNode)
	endpoint := fmt.Sprintf("%s%d/", workflowJobTemplateNodeAPIEndpoint, id)
	resp, err := jt.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// ListJobTemplates shows a list of job templates.
func (jt *WorkflowJobTemplateNodeService) ListWorkflowJobTemplateNodes(params map[string]string) ([]*WorkflowJobTemplateNode, *ListWorkflowJobTemplateNodesResponse, error) {
	result := new(ListWorkflowJobTemplateNodesResponse)

	resp, err := jt.client.Requester.GetJSON(workflowJobTemplateNodeAPIEndpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

// CreateJobTemplate creates a job template
func (jt *WorkflowJobTemplateNodeService) CreateWorkflowJobTemplateNode(data map[string]interface{}, params map[string]string) (*WorkflowJobTemplateNode, error) {
	result := new(WorkflowJobTemplateNode)
	mandatoryFields = []string{"workflow_job_template", "unified_job_template", "identifier"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := jt.client.Requester.PostJSON(workflowJobTemplateNodeAPIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}
	if err := CheckResponse(resp); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateJobTemplate updates a job template
func (jt *WorkflowJobTemplateNodeService) UpdateWorkflowJobTemplateNode(id int, data map[string]interface{}, params map[string]string) (*WorkflowJobTemplateNode, error) {
	result := new(WorkflowJobTemplateNode)
	endpoint := fmt.Sprintf("%s%d", workflowJobTemplateNodeAPIEndpoint, id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := jt.client.Requester.PatchJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}
	if err := CheckResponse(resp); err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteJobTemplate deletes a job template
func (jt *WorkflowJobTemplateNodeService) DeleteWorkflowJobTemplateNode(id int) (*WorkflowJobTemplateNode, error) {
	result := new(WorkflowJobTemplateNode)
	endpoint := fmt.Sprintf("%s%d", workflowJobTemplateNodeAPIEndpoint, id)

	resp, err := jt.client.Requester.Delete(endpoint, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
