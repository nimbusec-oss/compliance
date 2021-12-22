package compliance

import "net/http"

type IssueExternalService service

func (srv IssueExternalService) List(issueID string) ([]IssueExternal, error) {
	externals := []IssueExternal{}
	err := srv.client.Do(http.MethodGet, "/api/v1/issues/"+issueID+"/externals", nil, &externals)
	return externals, err
}

func (srv IssueExternalService) Create(issueID string, create []IssueExternal) ([]IssueExternal, error) {
	externals := []IssueExternal{}
	err := srv.client.Do(http.MethodPost, "/api/v1/issues/"+issueID+"/externals", create, &externals)
	return externals, err
}

func (srv IssueExternalService) Delete(issueID string) error {
	err := srv.client.Do(http.MethodDelete, "/api/v1/issues/"+issueID+"/externals", nil, nil)
	return err
}
