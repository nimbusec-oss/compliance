package compliance

import (
	"net/http"
)

type BundleService service

func (srv BundleService) List() ([]Bundle, error) {
	bundles := []Bundle{}
	err := srv.client.Do(http.MethodGet, "/api/v1/bundles", nil, &bundles)
	return bundles, err
}

func (srv BundleService) Get(bundleID string) (Bundle, error) {
	bundle := Bundle{}
	err := srv.client.Do(http.MethodGet, "/api/v1/bundles/"+bundleID, nil, &bundle)
	return bundle, err
}
