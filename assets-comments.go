package compliance

import "net/http"

type AssetCommentService service

func (srv AssetCommentService) List(assetID string) ([]AssetComment, error) {
	comments := []AssetComment{}
	err := srv.client.Do(http.MethodGet, "/api/v1/assets/"+assetID+"/comments", nil, &comments)
	return comments, err
}

func (srv AssetCommentService) Create(assetID string, create []AssetComment) ([]AssetComment, error) {
	comments := []AssetComment{}
	err := srv.client.Do(http.MethodPost, "/api/v1/assets/"+assetID+"/comments", create, &comments)
	return comments, err
}

func (srv AssetCommentService) Delete(assetID, commentID string) error {
	err := srv.client.Do(http.MethodDelete, "/api/v1/assets/"+assetID+"/comments/"+commentID, nil, nil)
	return err
}
