package compliance

import "net/http"

type AssetExternalService service

func (srv AssetExternalService) List(assetID string) ([]AssetExternal, error) {
	externals := []AssetExternal{}
	err := srv.client.Do(http.MethodGet, "/api/v1/assets/"+assetID+"/externals", nil, &externals)
	return externals, err
}

func (srv AssetExternalService) Create(assetID string, create []AssetExternal) ([]AssetExternal, error) {
	externals := []AssetExternal{}
	err := srv.client.Do(http.MethodPost, "/api/v1/assets/"+assetID+"/externals", create, &externals)
	return externals, err
}

func (srv AssetExternalService) Delete(assetID, externalID string) error {
	err := srv.client.Do(http.MethodDelete, "/api/v1/assets/"+assetID+"/externals/"+externalID, nil, nil)
	return err
}
