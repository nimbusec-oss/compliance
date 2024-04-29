package compliance

const (
	IssueTypeComplianceContentPolicy IssueType = "compliance_content_policy"
	IssueTypeComplianceCookieOptIn   IssueType = "compliance_cookie_opt_in"
	IssueTypeComplianceCookiePolicy  IssueType = "compliance_cookie_policy"
	IssueTypeComplianceFormHTTP      IssueType = "compliance_form_http"
	IssueTypeComplianceFormExternal  IssueType = "compliance_form_external"
	IssueTypeComplianceFormSensitive IssueType = "compliance_form_sensitive"
	IssueTypeComplianceStatusCode    IssueType = "compliance_status_code"
	IssueTypeComplianceTrackerOptIn  IssueType = "compliance_tracker_opt_in"
	IssueTypeComplianceTrackerPolicy IssueType = "compliance_tracker_policy"
	IssueTypeDanglingDNS             IssueType = "dangling-dns"
)

const (
	IssueStatusPending IssueStatus = 1
	IssueStatusOK      IssueStatus = 2 // acknowledged status
	IssueStatusIgnored IssueStatus = 3
	IssueStatusAutoOK  IssueStatus = 7
)

const (
	AssetInfoTypeASN         AssetInfoType = "asn"
	AssetInfoTypeValidation  AssetInfoType = "validation"
	AssetInfoTypeDiscovery   AssetInfoType = "discovery"
	AssetInfoTypeApplication AssetInfoType = "application"
	AssetInfoTypeLinks       AssetInfoType = "links"
)
