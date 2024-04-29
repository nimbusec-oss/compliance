package compliance

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type IssueService service

func (srv IssueService) List(filter *IssueFilter) ([]Issue, error) {
	v, err := query.Values(filter)
	if err != nil {
		return nil, err
	}
	u := url.URL{
		Path:     "/api/v1/issues",
		RawQuery: v.Encode(),
	}

	issues := []tempIssue{}
	err = srv.client.Do(http.MethodGet, u.String(), nil, &issues)
	if err != nil {
		return nil, err
	}

	translated := make([]Issue, len(issues))
	for i := 0; i < len(issues); i++ {
		issue := issues[i]
		err := issue.unmarshalDetails()
		if err != nil {
			return nil, err
		}
		translated[i] = issue.Issue
	}

	return translated, err
}

func (srv IssueService) PatchList(patches []IssuePatch) ([]Issue, error) {
	issues := []tempIssue{}
	err := srv.client.Do(http.MethodPatch, "/api/v1/issues", patches, &issues)
	if err != nil {
		return nil, err
	}

	translated := make([]Issue, len(issues))
	for i := 0; i < len(issues); i++ {
		issue := issues[i]
		err := issue.unmarshalDetails()
		if err != nil {
			return nil, err
		}
		translated[i] = issue.Issue
	}

	return translated, nil
}

func (srv IssueService) Get(issueID string) (Issue, error) {
	issue := tempIssue{}
	err := srv.client.Do(http.MethodGet, "/api/v1/issues/"+issueID, nil, &issue)
	if err != nil {
		return issue.Issue, err
	}

	err = issue.unmarshalDetails()
	return issue.Issue, err
}

func (srv IssueService) Patch(issueID string, patch IssuePatch) (Issue, error) {
	patch.IssueID = issueID
	issues, err := srv.PatchList([]IssuePatch{patch})
	if err != nil {
		return Issue{}, err
	}

	return issues[0], nil
}
