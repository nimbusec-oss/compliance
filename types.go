package compliance

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

type Asset struct {
	ID        string    `json:"id"`
	IDfy      string    `json:"idfy"`
	Origin    string    `json:"origin"`
	CBundleID *string   `json:"cBundleID"`
	CreatedAt time.Time `json:"createdAt"`
	ContactID *string   `json:"contactID"`
}

type AssetPatch struct {
	CBundleID *string `json:"cBundleID"`
	Origin    string  `json:"origin"`
	ContactID *string `json:"contactID"`
}

type AssetFilter struct {
	OnlySecurity   bool `json:"onlySecurity"`   // only security assets
	OnlyCompliance bool `json:"onlyCompliance"` // only compliance assets

	Name      []string `json:"name"` // idfy
	Origin    []string `json:"origin"`
	CreatedAt []string `json:"createdAt"`
}

type DisabledAsset struct {
	ID             string    `json:"id"`
	IDfy           string    `json:"idfy"`
	Origin         string    `json:"origin"`
	CreatedAt      time.Time `json:"createdAt"`
	DisabledAt     time.Time `json:"disabledAt"`
	DisabledReason *string   `json:"disabledReason"`
}

type DisabledAssetFilter struct {
	Name           []string `json:"name,omitempty"` // idfy
	Origin         []string `json:"origin,omitempty"`
	CreatedAt      []string `json:"createdAt,omitempty"`
	DisabledAt     []string `json:"disabledAt,omitempty"`
	DisabledReason []string `json:"disabledReason,omitempty"`
}

type DisablePatch struct {
	Reason string `json:"reason"`
}

type AssetTag struct {
	AssetID   string    `json:"assetID"`
	TagID     string    `json:"tagID"`
	CreatedAt time.Time `json:"createdAt"`
}

type AssetExternal struct {
	ID       string  `json:"id"`
	AssetID  string  `json:"assetID"`
	External string  `json:"external"`
	Comment  *string `json:"comment"`
}

type AssetComment struct {
	ID        string     `json:"id"`
	Value     string     `json:"value"`
	CreatedBy string     `json:"createdBy"`
	CreatedAt time.Time  `json:"createdAt"`
	ChangedAt *time.Time `json:"changedAt"`
	AssetID   *string    `json:"assetID"`
}

type Bundle struct {
	ID                     string `json:"id"`
	Name                   string `json:"name"`
	ComplianceScanInterval int    `json:"scanInterval"`
	Disabled               bool   `json:"disabled"`
}

type Contact struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type ASN int64

func (a ASN) String() string {
	return fmt.Sprintf("AS%d", a)
}

type IPAS struct {
	IP           string `json:"ip"`
	CIDR         string `json:"cidr"`
	ASN          ASN    `json:"asn"`
	Organization string `json:"organization"`
	CountryCode  string `json:"countryCode"`
	Country      string `json:"country"`
	IsEU         bool   `json:"isEU"`
}

type AssetInfoASN struct {
	IPASs      []IPAS   `json:"ipASs"`
	Nameserver []string `json:"nameserver"`
	IsEU       *bool    `json:"isEU,omitempty"`
}

type AssetInfoValidation struct {
	Redirects      []Redirect `json:"redirects"`
	HasBadRedirect bool       `json:"hasBadRedirect"`
	Scheme         string     `json:"scheme"`
	StatusCode     int        `json:"statusCode"`
	HasHTTP        bool       `json:"hasHTTP"`
	HasHTTPS       bool       `json:"hasHTTPS"`
	HostsDifferent bool       `json:"hostsDifferent"`
	ContentEqual   bool       `json:"contentEqual"`
	HTTPDowngrade  bool       `json:"httpDowngrade"`
}

type DiscoveryInput struct {
	Name  string `json:"name"`
	Src   string `json:"src"`
	Value string `json:"value"`
}

type Redirect struct {
	URL        string `json:"url"`
	Address    string `json:"address"`
	StatusCode int    `json:"statusCode"`
}

type AssetInfoDiscovery struct {
	Sources   []DiscoveryInput `json:"sources"`
	Redirects []Redirect       `json:"redirects"`
}

type AssetInfoApplications []AssetInfoApplication
type AssetInfoApplication struct {
	Name       string `json:"name"`
	Version    string `json:"version,omitempty"`
	Path       string `json:"path,omitempty"`
	Category   string `json:"category"`
	Source     string `json:"source"`
	Latest     bool   `json:"latest"`
	Vulnerable bool   `json:"vulnerable"`
}

type AssetInfoType string
type AssetInfo struct {
	AssetID   string        `json:"assetID"`
	Type      AssetInfoType `json:"type"`
	Details   interface{}   `json:"details"`
	ChangedAt time.Time     `json:"changedAt"`
}

type Metadata struct {
	AssetID string                      `json:"assetID"`
	Infos   map[AssetInfoType]AssetInfo `json:"infos"`
}

type MetadataFilter struct {
	Types []string `url:"type,omitempty"`
}

type IssueExternal struct {
	IssueID    string  `json:"issueID"`
	ExternalID string  `json:"externalID"`
	Comment    *string `json:"comment"`
}

type IssueType string
type Severity int
type IssueStatus int
type Issue struct {
	ID              string      `json:"id"`
	AssetID         string      `json:"assetID"`
	Type            IssueType   `json:"type"`
	Severity        Severity    `json:"severity"`
	Status          IssueStatus `json:"status"`
	StatusChangedAt time.Time   `json:"statusChangedAt"`
	CreatedAt       time.Time   `json:"createdAt"`
	LastSeenAt      time.Time   `json:"lastSeenAt"`
	Details         interface{} `json:"details"`
	CommentID       *string     `json:"commentID"`
}

// tempIssue is an temporary struct to work with the raw byte response from the API
type tempIssue struct {
	Issue
	RawDetails json.RawMessage `json:"details"`
}

// unmarshalDetails converts the given details value from the API to the correct details type
func (issue *tempIssue) unmarshalDetails() error {
	switch issue.Type {
	case IssueTypeComplianceContentPolicy:
		d := ContentPolicyDetails{}
		err := json.Unmarshal(issue.RawDetails, &d)
		if err != nil {
			return err
		}
		issue.Details = d
	case IssueTypeComplianceCookieOptIn:
		d := CookieDetails{}
		err := json.Unmarshal(issue.RawDetails, &d)
		if err != nil {
			return err
		}
		issue.Details = d
	case IssueTypeComplianceCookiePolicy:
		d := CookiePolicyDetails{}
		err := json.Unmarshal(issue.RawDetails, &d)
		if err != nil {
			return err
		}
		issue.Details = d
	case IssueTypeComplianceFormExternal,
		IssueTypeComplianceFormHTTP,
		IssueTypeComplianceFormSensitive:
		d := FormDetails{}
		err := json.Unmarshal(issue.RawDetails, &d)
		if err != nil {
			return err
		}
		issue.Details = d
	case IssueTypeComplianceStatusCode:
		d := StatusCodeDetails{}
		err := json.Unmarshal(issue.RawDetails, &d)
		if err != nil {
			return err
		}
		issue.Details = d
	case IssueTypeComplianceTrackerOptIn:
		d := TrackerDetails{}
		err := json.Unmarshal(issue.RawDetails, &d)
		if err != nil {
			return err
		}
		issue.Details = d
	case IssueTypeComplianceTrackerPolicy:
		d := TrackerPolicyDetails{}
		err := json.Unmarshal(issue.RawDetails, &d)
		if err != nil {
			return err
		}
		issue.Details = d
	default:
		return errors.Errorf("unknown issue type '%s' for id '%s'", issue.Type, issue.ID)
	}

	return nil
}

type IssueFilter struct {
	Statuses  []IssueStatus `json:"statuses"`
	Externals []string      `json:"externals"`
}

type TrackerCategory string
type TrackerDetails struct {
	Name       string          `json:"name"`
	Category   TrackerCategory `json:"category"`
	Hostname   string          `json:"hostname"`
	Initiators [][]Initiator   `json:"initiators"`
	Visited    bool            `json:"visited"`
}

type TrackerPolicyDetails struct {
	TrackerDetails
	Rule string `json:"rule"`
}

// Initiator holds all information about the initiator of an request/malware issue
type Initiator struct {
	InitType string `json:"type"`
	URL      string `json:"url"`
	Line     string `json:"line"`
	Column   string `json:"column"`
}

type StatusCodeDetails struct {
	URL          string `json:"url"`
	StatusCode   int    `json:"statusCode"`
	ExpectedCode int    `json:"expected"`
}

type FormDetails struct {
	URLs             []string       `json:"urls"`
	Name             string         `json:"name"`
	FormID           string         `json:"formID,omitempty"`
	Action           string         `json:"action"`
	Method           string         `json:"method"`
	InputFields      []InputDetails `json:"inputDetails"`
	HTTPTransmit     bool           `json:"httpTransmit"`
	ExternalTransmit bool           `json:"externalTransmit"`
	// if it contains input fields that violates either art 9 or 10 once
	Art9  bool `json:"art9"`
	Art10 bool `json:"art10"`
}

type InputDetails struct {
	InputName         string `json:"inputName"`
	InputType         string `json:"inputType"`
	InputID           string `json:"inputID,omitempty"`
	InputAutocomplete string `json:"inputAutocomplete"`
	InputPlaceholder  string `json:"inputPlaceholder"`
	Category          string `json:"category"`
	Art9              bool   `json:"art9"`
	Art10             bool   `json:"art10"`
}

type CookieDetails struct {
	Name              string     `json:"name"`
	URLValues         []URLValue `json:"urlValues"` // URLValues[value]{url,url...}
	Domain            string     `json:"domain"`
	HTTPOnly          bool       `json:"httpOnly"`
	Lifetime          int64      `json:"lifetime"`
	SameSite          string     `json:"sameSite"`
	Secure            bool       `json:"secure"`
	ThirdParty        bool       `json:"thirdParty"`
	InGlobalWhitelist bool       `json:"inGlobalWhitelist"`
	InCustomWhitelist bool       `json:"inCustomWhitelist"`
	InServerResponse  bool       `json:"inServerResponse"`
}

type URLValue struct {
	URL   string `json:"url"`
	Value string `json:"value"`
}

type CookiePolicyDetails struct {
	CookieDetails
}

type ContentPolicyDetails struct {
	URL  string `json:"url"`
	Rule string `json:"rule"`
}

type IssuePatch struct {
	IssueID string      `json:"issueID"`
	Status  IssueStatus `json:"status"`
	Comment *string     `json:"comment"`
}

type TagType string
type Tag struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Type        TagType `json:"type"`
}
