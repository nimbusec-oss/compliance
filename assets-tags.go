package compliance

import "net/http"

type AssetTagService service

func (srv AssetTagService) List(assetID string) ([]AssetTag, error) {
	tags := []AssetTag{}
	err := srv.client.Do(http.MethodGet, "/api/v1/assets/"+assetID+"/tags", nil, &tags)
	return tags, err
}

func (srv AssetTagService) Create(assetID string, create []Tag) ([]Tag, error) {
	tags := []Tag{}
	err := srv.client.Do(http.MethodPost, "/api/v1/assets/"+assetID+"/tags", create, &tags)
	return tags, err
}

func (srv AssetTagService) Link(assetID, tagID string) error {
	err := srv.client.Do(http.MethodPut, "/api/v1/assets/"+assetID+"/tags/"+tagID, nil, nil)
	return err
}

func (srv AssetTagService) Unlink(assetID, tagID string) error {
	err := srv.client.Do(http.MethodDelete, "/api/v1/assets/"+assetID+"/tags/"+tagID, nil, nil)
	return err
}
