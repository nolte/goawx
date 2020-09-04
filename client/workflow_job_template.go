package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// JobTemplateService implements awx job template apis.
type WorkflowJobTemplateService struct {
	client *Client
}

// ListJobTemplatesResponse represents `ListJobTemplates` endpoint response.
type ListWorkflowJobTemplatesResponse struct {
	Pagination
	Results []*WorkflowJobTemplate `json:"results"`
}

const workflowJobTemplateAPIEndpoint = "/api/v2/workflow_job_templates/"

// GetJobTemplateByID shows the details of a job template.
func (jt *WorkflowJobTemplateService) GetWorkflowJobTemplateByID(id int, params map[string]string) (*WorkflowJobTemplate, error) {
	result := new(WorkflowJobTemplate)
	endpoint := fmt.Sprintf("%s%d/", workflowJobTemplateAPIEndpoint, id)
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
func (jt *WorkflowJobTemplateService) ListWorkflowJobTemplates(params map[string]string) ([]*WorkflowJobTemplate, *ListWorkflowJobTemplatesResponse, error) {
	result := new(ListWorkflowJobTemplatesResponse)
	resp, err := jt.client.Requester.GetJSON(workflowJobTemplateAPIEndpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

// CreateJobTemplate creates a job template
func (jt *WorkflowJobTemplateService) CreateWorkflowJobTemplate(data map[string]interface{}, params map[string]string) (*WorkflowJobTemplate, error) {
	result := new(WorkflowJobTemplate)
	mandatoryFields = []string{"name"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := jt.client.Requester.PostJSON(workflowJobTemplateAPIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}
	if err := CheckResponse(resp); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateJobTemplate updates a job template
func (jt *WorkflowJobTemplateService) UpdateWorkflowJobTemplate(id int, data map[string]interface{}, params map[string]string) (*WorkflowJobTemplate, error) {
	result := new(WorkflowJobTemplate)
	endpoint := fmt.Sprintf("%s%d", workflowJobTemplateAPIEndpoint, id)
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
func (jt *WorkflowJobTemplateService) DeleteWorkflowJobTemplate(id int) (*WorkflowJobTemplate, error) {
	result := new(WorkflowJobTemplate)
	endpoint := fmt.Sprintf("%s%d", workflowJobTemplateAPIEndpoint, id)

	resp, err := jt.client.Requester.Delete(endpoint, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
