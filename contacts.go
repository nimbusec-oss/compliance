package compliance

import (
	"net/http"
)

type ContactService service

func (srv ContactService) List() ([]Contact, error) {
	contacts := []Contact{}
	err := srv.client.Do(http.MethodGet, "/api/v1/contacts", nil, &contacts)
	return contacts, err
}

func (srv ContactService) Get(contactID string) (Contact, error) {
	contact := Contact{}
	err := srv.client.Do(http.MethodGet, "/api/v1/contacts/"+contactID, nil, &contact)
	return contact, err
}

func (srv ContactService) Create(create Contact) (Contact, error) {
	contact := Contact{}
	err := srv.client.Do(http.MethodPost, "/api/v1/contacts", create, &contact)
	return contact, err
}

func (srv ContactService) Update(contactID string, update Contact) (Contact, error) {
	contact := Contact{}
	err := srv.client.Do(http.MethodPut, "/api/v1/contacts/"+contactID, update, &contact)
	return contact, err
}

func (srv ContactService) Delete(contactID string) error {
	err := srv.client.Do(http.MethodDelete, "/api/v1/contacts/"+contactID, nil, nil)
	return err
}
