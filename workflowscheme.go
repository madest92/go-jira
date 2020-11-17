package jira

import (
	"context"
	"fmt"
)

type WorkflowSchemeService struct {
	client *Client
}

// GetWithContext returns a full workflow scheme of the project for the given project key
func (s *WorkflowSchemeService) GetWithContext(ctx context.Context, issue string) (*string, *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/scriptrunner/latest/custom/getWfName?key=%s", issue)
	req, err := s.client.NewRequestWithContext(ctx, "GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	workflowScheme := new(string)
	resp, err := s.client.Do(req, workflowScheme)

	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}

	return workflowScheme, resp, nil
}

// Get wraps GetWorkflowSchemeWithContext using the background context.
func (s *WorkflowSchemeService) Get(issue string) (*string, *Response, error) {
	return s.GetWithContext(context.Background(), issue)
}
