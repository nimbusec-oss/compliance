package compliance

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type MetadataService service

func (srv MetadataService) List(filter *MetadataFilter) ([]Metadata, error) {
	v, err := query.Values(filter)
	if err != nil {
		return nil, err
	}
	u := url.URL{
		Path:     "/api/v1/assets/metadata",
		RawQuery: v.Encode(),
	}

	metadatas := []Metadata{}
	err = srv.client.Do(http.MethodGet, u.String(), nil, &metadatas)
	return metadatas, err
}

func (srv MetadataService) Get(assetID string) (Metadata, error) {
	metadata := Metadata{}
	err := srv.client.Do(http.MethodGet, "/assets/"+assetID+"/metadata", nil, &metadata)
	return metadata, err
}
