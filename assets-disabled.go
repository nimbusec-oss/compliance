package compliance

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type AssetDisabledService service

func (srv AssetDisabledService) List(filter *DisabledAssetFilter) ([]DisabledAsset, error) {
	v, err := query.Values(filter)
	if err != nil {
		return nil, err
	}
	u := url.URL{
		Path:     "/api/v1/assets-disabled",
		RawQuery: v.Encode(),
	}

	disabled := []DisabledAsset{}
	err = srv.client.Do(http.MethodGet, u.String(), nil, &disabled)
	return disabled, err
}

func (srv AssetDisabledService) Enable(assetID string) error {
	return srv.client.Do(http.MethodPost, "/api/v1/assets-disabled/"+assetID+"/enable", nil, nil)
}
