package compliance

import (
	"net/http"
)

type TagService service

func (srv TagService) List() ([]Tag, error) {
	tags := []Tag{}
	err := srv.client.Do(http.MethodGet, "/api/v1/tags", nil, &tags)
	return tags, err
}

func (srv TagService) Create(create Tag) (Tag, error) {
	tag := Tag{}
	err := srv.client.Do(http.MethodPost, "/api/v1/tags", create, &tag)
	return tag, err
}

func (srv TagService) Get(tagID string) (Tag, error) {
	tag := Tag{}
	err := srv.client.Do(http.MethodGet, "/api/v1/tags/"+tagID, nil, &tag)
	return tag, err
}

func (srv TagService) Update(tagID string, update Tag) (Tag, error) {
	tag := Tag{}
	err := srv.client.Do(http.MethodPut, "/api/v1/tags/"+tagID, update, &tag)
	return tag, err
}

func (srv TagService) Delete(tagID string) error {
	err := srv.client.Do(http.MethodDelete, "/api/v1tags/"+tagID, nil, nil)
	return err
}
