package compliance

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type AssetService service

func (srv AssetService) List(filter *AssetFilter) ([]Asset, error) {
	v, err := query.Values(filter)
	if err != nil {
		return nil, err
	}
	u := url.URL{
		Path:     "/api/v1/assets",
		RawQuery: v.Encode(),
	}

	assets := []Asset{}
	err = srv.client.Do(http.MethodGet, u.String(), nil, &assets)
	return assets, err
}

func (srv AssetService) Get(assetID string) (Asset, error) {
	asset := Asset{}
	err := srv.client.Do(http.MethodGet, "/api/v1/assets/"+assetID, nil, &asset)
	return asset, err
}

func (srv AssetService) Create(create Asset) (Asset, error) {
	asset := Asset{}
	err := srv.client.Do(http.MethodPost, "/api/v1/assets", create, &asset)
	return asset, err
}

func (srv AssetService) Patch(assetID string, patch AssetPatch) (Asset, error) {
	asset := Asset{}
	err := srv.client.Do(http.MethodPatch, "/api/v1/assets/"+assetID, patch, &asset)
	return asset, err
}

func (srv AssetService) Delete(assetID string) error {
	err := srv.client.Do(http.MethodDelete, "/api/v1/assets/"+assetID, nil, nil)
	return err
}

func (srv AssetService) GetIssues(assetID string) ([]Issue, error) {
	issues := []tempIssue{}
	err := srv.client.Do(http.MethodGet, "/api/v1/assets/"+assetID+"/issues", nil, &issues)
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

func (srv AssetService) Disable(assetID string, patch *DisablePatch) error {
	err := srv.client.Do(http.MethodPost, "/api/v1/assets/"+assetID+"/disable", patch, nil)
	return err
}
