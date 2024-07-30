package compliance

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/oauth2/clientcredentials"
)

type service struct {
	client *Client
}

type Client struct {
	http *http.Client
	base string
	dry  bool

	common         service // Reuse a single struct instead of allocating one for each service on the heap.
	Tags           *TagService
	Bundles        *BundleService
	Contacts       *ContactService
	Issues         *IssueService
	IssueExternals *IssueExternalService
	Metadata       *MetadataService
	Assets         *AssetService
	DisabledAssets *AssetDisabledService
	AssetTags      *AssetTagService
	AssetExternals *AssetExternalService
	AssetComments  *AssetCommentService
}

type Error struct {
	Message    string            `json:"message"`
	Validation map[string]string `json:"validation,omitempty"`
}

func (e Error) Error() string {
	return e.Message
}

type ClientConfig struct {
	ClientID     string
	ClientSecret string
	Endpoint     Endpoint
	Dry          bool
}

type Endpoint struct {
	AuthURL   string
	ServerURL string
}

var DefaultEndpoint = Endpoint{
	AuthURL:   "https://auth.nimbusec.com/oauth2/token",
	ServerURL: "https://compliance.nimbusec.com",
}

func NewClient(config ClientConfig) *Client {
	oauth := &clientcredentials.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		TokenURL:     config.Endpoint.AuthURL,
	}

	c := &Client{http: oauth.Client(context.Background()), base: config.Endpoint.ServerURL, dry: config.Dry}
	c.common = service{client: c}
	c.Tags = (*TagService)(&c.common)
	c.Bundles = (*BundleService)(&c.common)
	c.Contacts = (*ContactService)(&c.common)
	c.Issues = (*IssueService)(&c.common)
	c.IssueExternals = (*IssueExternalService)(&c.common)
	c.Metadata = (*MetadataService)(&c.common)
	c.Assets = (*AssetService)(&c.common)
	c.DisabledAssets = (*AssetDisabledService)(&c.common)
	c.AssetTags = (*AssetTagService)(&c.common)
	c.AssetExternals = (*AssetExternalService)(&c.common)
	c.AssetComments = (*AssetCommentService)(&c.common)

	return c
}

func (client *Client) Ping() error {
	err := client.Do(http.MethodHead, "/health", nil, nil)
	return err
}

func (client *Client) Do(method, path string, in, out interface{}) error {
	var req *http.Request
	var err error

	// encode `in` object as json body if provided
	if in != nil {
		body := &bytes.Buffer{}
		err = json.NewEncoder(body).Encode(in)
		if err != nil {
			return fmt.Errorf("Encode request body: %w", err)
		}

		req, err = http.NewRequest(method, client.base+path, body)
	} else {
		req, err = http.NewRequest(method, client.base+path, nil)
	}

	if err != nil {
		return fmt.Errorf("NewRequest: %w", err)
	}

	// dry mode active, do not make an action except for read-only GET requests
	if client.dry && method != http.MethodGet {
		return nil
	}

	// perform request
	req = req.WithContext(context.Background())
	req.Close = true
	resp, err := client.http.Do(req)
	if err != nil {
		return fmt.Errorf("Client.Do: %w", err)
	}

	// decode response body into `out` object if provided
	defer resp.Body.Close()

	// check for api errors, try to decode in error object
	if resp.StatusCode >= 400 {
		buf := []byte{}
		_, err = io.ReadFull(resp.Body, buf)
		if err != nil {
			return fmt.Errorf("%d | %w", resp.StatusCode, err)
		}

		msg := Error{}
		err = json.Unmarshal(buf, &msg)
		if err != nil {
			return fmt.Errorf("message %s could not be decoded into err: %w | Statuscode was: %d", string(buf), err, resp.StatusCode)
		}
		return fmt.Errorf("%d | %w", resp.StatusCode, err)
	}

	if out != nil {
		err = json.NewDecoder(resp.Body).Decode(&out)
		if err != nil {
			return fmt.Errorf("%d | %w", resp.StatusCode, err)
		}
	}

	return nil
}
