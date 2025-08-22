package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ApplicationStatusRequest represents the ApplicationStatusRequest schema from the OpenAPI specification
type ApplicationStatusRequest struct {
	Build_version string `json:"build_version,omitempty"` // The version of build for which real time status is to be fetched.
	Bundle_identifier string `json:"bundle_identifier"` // Bundle Identifier of application in Apple Itunes portal.
	Password string `json:"password"` // The password for the Apple Developer account.
	Team_identifier string `json:"team_identifier,omitempty"` // Identifier of the team to use when logged in.
	Track_identifier string `json:"track_identifier"` // Track Identifier for which the status is to be fetched.
	Train_version string `json:"train_version,omitempty"` // The Train version for which the status is to be fetched.
	Username string `json:"username"` // The username for the Apple Developer account.
}

// TestGDPRPipelineTest represents the TestGDPRPipelineTest schema from the OpenAPI specification
type TestGDPRPipelineTest struct {
	App_upload_id string `json:"app_upload_id,omitempty"`
	Test_parameters map[string]interface{} `json:"test_parameters,omitempty"`
}

// ReleaseDetailsUpdateRequest represents the ReleaseDetailsUpdateRequest schema from the OpenAPI specification
type ReleaseDetailsUpdateRequest struct {
	Enabled bool `json:"enabled,omitempty"` // Toggle this release to be enable distribute/download or not.
	Release_notes string `json:"release_notes,omitempty"` // Release notes for this release.
	Build map[string]interface{} `json:"build,omitempty"` // Contains metadata about the build that produced the release being uploaded
}

// AppleCertificateSecretDetails represents the AppleCertificateSecretDetails schema from the OpenAPI specification
type AppleCertificateSecretDetails struct {
	Certificatevaliditystartdate string `json:"certificateValidityStartDate"` // The date-time from which the certificate is valid
	Displayname string `json:"displayName"` // The display name (CN) of the certificate
	Password string `json:"password"` // The password for the certificate
	Base64certificate string `json:"base64Certificate"` // The certificate contents in base 64 encoded string
	Certificatevalidityenddate string `json:"certificateValidityEndDate"` // The date-time till which the certificate is valid
}

// PackageHashToBlobInfoMap represents the PackageHashToBlobInfoMap schema from the OpenAPI specification
type PackageHashToBlobInfoMap struct {
}

// UserEmailOrgRoleRequest represents the UserEmailOrgRoleRequest schema from the OpenAPI specification
type UserEmailOrgRoleRequest struct {
	Role string `json:"role,omitempty"` // The user's role
	User_email string `json:"user_email"` // The user's email address
}

// DeploymentModification represents the DeploymentModification schema from the OpenAPI specification
type DeploymentModification struct {
	Name string `json:"name"`
}

// StartServiceLog represents the StartServiceLog schema from the OpenAPI specification
type StartServiceLog struct {
	Install_id string `json:"install_id"` // Install ID.
	Timestamp string `json:"timestamp"` // Log creation timestamp.
	TypeField string `json:"type"` // Log type.
	Device map[string]interface{} `json:"device"` // Device characteristics.
}

// InternalBulkAppResponse represents the InternalBulkAppResponse schema from the OpenAPI specification
type InternalBulkAppResponse struct {
	App_name string `json:"app_name,omitempty"` // The name of the app
	Owner_display_name string `json:"owner_display_name,omitempty"` // The display name of the owner
}

// LegacyAppResponse represents the LegacyAppResponse schema from the OpenAPI specification
type LegacyAppResponse struct {
	App interface{} `json:"app,omitempty"`
}

// CodePushReleaseModification represents the CodePushReleaseModification schema from the OpenAPI specification
type CodePushReleaseModification struct {
	Is_disabled bool `json:"is_disabled,omitempty"`
	Is_mandatory bool `json:"is_mandatory,omitempty"`
	Rollout int `json:"rollout,omitempty"`
	Target_binary_range string `json:"target_binary_range,omitempty"`
	Description string `json:"description,omitempty"`
}

// IntuneAppsResponse represents the IntuneAppsResponse schema from the OpenAPI specification
type IntuneAppsResponse struct {
	Created_month string `json:"created_month,omitempty"` // PartitionKey year-month
	Refreshstatus string `json:"refreshStatus,omitempty"` // Refresh Status
	App_id string `json:"app_id,omitempty"` // App id
}

// ErrorLog represents the ErrorLog schema from the OpenAPI specification
type ErrorLog struct {
	Device map[string]interface{} `json:"device"` // Device characteristics.
	Install_id string `json:"install_id"` // Install ID.
	Timestamp string `json:"timestamp"` // Log creation timestamp.
	TypeField string `json:"type"` // Log type.
}

// AADTenantResponse represents the AADTenantResponse schema from the OpenAPI specification
type AADTenantResponse struct {
	Aad_tenant_id string `json:"aad_tenant_id"` // The AAD tenant id
	Display_name string `json:"display_name"` // The name of the AAD Tenant
}

// AppleCertificateNonSecretDetailsResponse represents the AppleCertificateNonSecretDetailsResponse schema from the OpenAPI specification
type AppleCertificateNonSecretDetailsResponse struct {
	Displayname string `json:"displayName"` // display name of shared connection
	Servicetype string `json:"serviceType"` // service type of shared connection can be apple|gitlab|googleplay|jira
	Credentialtype string `json:"credentialType,omitempty"` // the type of the credential
}

// TeamAppResponse represents the TeamAppResponse schema from the OpenAPI specification
type TeamAppResponse struct {
	Id string `json:"id"` // The internal unique id (UUID) of the team.
	Name string `json:"name"` // The name of the team
	Description string `json:"description,omitempty"` // The description of the team
	Display_name string `json:"display_name"` // The display name of the team
	Permissions []string `json:"permissions,omitempty"` // The permissions the team has for the app
}

// AlertUserEmailSettingsResult represents the AlertUserEmailSettingsResult schema from the OpenAPI specification
type AlertUserEmailSettingsResult struct {
	Request_id string `json:"request_id"` // Unique request identifier for tracking
	Enabled bool `json:"enabled"` // Allows to forcefully disable emails on app or user level
	Settings []map[string]interface{} `json:"settings"` // The settings the user has for the app
	Userid string `json:"userId,omitempty"` // The unique id (UUID) of the user
	Etag string `json:"eTag,omitempty"` // The ETag of the entity
}

// CreateReleaseUploadResponse represents the CreateReleaseUploadResponse schema from the OpenAPI specification
type CreateReleaseUploadResponse struct {
	Id string `json:"id"` // The ID for the newly created upload. It is going to be required later in the process.
	Package_asset_id string `json:"package_asset_id"` // The associated asset ID in the file management service associated with this uploaded.
	Token string `json:"token"` // The access token used for upload permissions.
	Upload_domain string `json:"upload_domain"` // The URL domain used to upload the release.
	Url_encoded_token string `json:"url_encoded_token"` // The access token used for upload permissions (URL encoded to use as a single query parameter).
}

// TestSeriesName represents the TestSeriesName schema from the OpenAPI specification
type TestSeriesName struct {
	Name string `json:"name"` // Name of the new test series
}

// LogFlowLogWithProperties represents the LogFlowLogWithProperties schema from the OpenAPI specification
type LogFlowLogWithProperties struct {
	Timestamp string `json:"timestamp"` // Log creation timestamp.
	TypeField string `json:"type"` // Log type.
	Device map[string]interface{} `json:"device"` // Device characteristics.
	Install_id string `json:"install_id"` // Install ID.
}

// AADUser represents the AADUser schema from the OpenAPI specification
type AADUser struct {
	Object_id string `json:"object_id"` // the aad user's id
	Tenant_id string `json:"tenant_id"` // the aad user's tenant id
}

// TeamAppUpdateRequest represents the TeamAppUpdateRequest schema from the OpenAPI specification
type TeamAppUpdateRequest struct {
	Permissions []string `json:"permissions"` // The permissions all members of the team have on the app
}

// BuildServiceStatus represents the BuildServiceStatus schema from the OpenAPI specification
type BuildServiceStatus struct {
	Os string `json:"os,omitempty"`
	Service string `json:"service,omitempty"`
	Status string `json:"status,omitempty"`
	Url string `json:"url,omitempty"`
	Valid_until int `json:"valid_until,omitempty"`
	Message string `json:"message,omitempty"`
}

// AccessKeyListResponse represents the AccessKeyListResponse schema from the OpenAPI specification
type AccessKeyListResponse struct {
	Accesskeys []interface{} `json:"accessKeys,omitempty"` // Array containing the list of existing AccessKeys
}

// AzureSubscriptionUpdateBillableRequest represents the AzureSubscriptionUpdateBillableRequest schema from the OpenAPI specification
type AzureSubscriptionUpdateBillableRequest struct {
	Is_billable bool `json:"is_billable"` // Billable status of the subscription
}

// Exception represents the Exception schema from the OpenAPI specification
type Exception struct {
	Inner_exceptions []Exception `json:"inner_exceptions,omitempty"`
	Platform string `json:"platform,omitempty"` // SDK/Platform this thread is beeing generated from
	Reason string `json:"reason,omitempty"` // Reason of the exception
	Relevant bool `json:"relevant,omitempty"` // relevant exception (crashed)
	TypeField string `json:"type,omitempty"` // Type of the exception (NSSomethingException, NullPointerException)
	Frames []interface{} `json:"frames"` // frames of the excetpion
}

// FeatureResponse represents the FeatureResponse schema from the OpenAPI specification
type FeatureResponse struct {
	State int `json:"state"` // The state (unset, enabled, disabled) of the feature
	Updated_at string `json:"updated_at,omitempty"` // The date the feature was last updated at
	Created_at string `json:"created_at,omitempty"` // The creation date of the feature
	Description string `json:"description,omitempty"` // The description of the feature
	Display_name string `json:"display_name"` // The friendly name of the feature
	Name string `json:"name"` // The unique name of the feature
}

// ErrorCounts represents the ErrorCounts schema from the OpenAPI specification
type ErrorCounts struct {
	Count int64 `json:"count,omitempty"` // total error count
	Errors []interface{} `json:"errors,omitempty"` // the total error count for day
}

// DateTimeDownloadReleaseCount represents the DateTimeDownloadReleaseCount schema from the OpenAPI specification
type DateTimeDownloadReleaseCount struct {
	Unique int64 `json:"unique,omitempty"`
	Datetime string `json:"datetime,omitempty"` // The ISO 8601 datetime.
	Total int64 `json:"total,omitempty"`
}

// PageLog represents the PageLog schema from the OpenAPI specification
type PageLog struct {
	TypeField string `json:"type"` // Log type.
	Device map[string]interface{} `json:"device"` // Device characteristics.
	Install_id string `json:"install_id"` // Install ID.
	Timestamp string `json:"timestamp"` // Log creation timestamp.
	Properties map[string]interface{} `json:"properties,omitempty"` // Additional key/value pair parameters.
}

// RawCrashLog represents the RawCrashLog schema from the OpenAPI specification
type RawCrashLog struct {
}

// AggregatedBillingInformation represents the AggregatedBillingInformation schema from the OpenAPI specification
type AggregatedBillingInformation struct {
	Usage map[string]interface{} `json:"usage,omitempty"` // Usage section in the Billing Information
	Version string `json:"version,omitempty"` // Version of the Billing Information schema
	Azuresubscriptionid string `json:"azureSubscriptionId,omitempty"` // Unique identifier for the Azure subscription used for billing
	Azuresubscriptionstate string `json:"azureSubscriptionState,omitempty"` // State of the Azure subscription used for billing
	Billingplans map[string]interface{} `json:"billingPlans,omitempty"` // Billing Plans section in the Billing Information
	Id string `json:"id,omitempty"` // ID of the user or organization
	Timestamp string `json:"timestamp,omitempty"` // The ISO 8601 datetime of last modification
}

// DataSubjectRightOperation represents the DataSubjectRightOperation schema from the OpenAPI specification
type DataSubjectRightOperation struct {
	Requestid string `json:"requestId"` // Unique request identifier
	Requesttype string `json:"requestType"` // Request type
	Status string `json:"status"` // Operation status
	Appid string `json:"appId,omitempty"` // Application identifier if applicable
	Context string `json:"context"` // JSON object decribing what to delete (TODO - make separate definition?)
	Operationid string `json:"operationId"` // Unique operation identifier
	Participant string `json:"participant"` // Participant to execute the response
	Participantdata string `json:"participantData,omitempty"` // String field to be used by participant for any intermediate statuses or data they need to save
}

// ValidatedApiTokenResponse represents the ValidatedApiTokenResponse schema from the OpenAPI specification
type ValidatedApiTokenResponse struct {
	Principal_type string `json:"principal_type"` // Indicates the type of the principal (app or user)
	Token_id string `json:"token_id"` // The token's unique id (UUID)
	Token_scope []string `json:"token_scope"` // The token's scope. A list of allowed roles.
	Claims []interface{} `json:"claims"` // Collection of attributes that describe the principal of the specified API Token
	Principal_id string `json:"principal_id"` // The ID of the owner of the API Token (user_id or app_id)
}

// ItunesTeamsResponse represents the ItunesTeamsResponse schema from the OpenAPI specification
type ItunesTeamsResponse struct {
	Teamid string `json:"teamId,omitempty"` // Itunes team id.
	Teamname string `json:"teamName,omitempty"` // Itunes Team Name
}

// MalwareScanResultPayload represents the MalwareScanResultPayload schema from the OpenAPI specification
type MalwareScanResultPayload struct {
	Requestid string `json:"requestId"` // Scan request identifier
	Result string `json:"result"` // Scan result
}

// TestGDPRHashFile represents the TestGDPRHashFile schema from the OpenAPI specification
type TestGDPRHashFile struct {
	Id string `json:"id,omitempty"`
	Filename string `json:"filename,omitempty"`
}

// AppleCredentialsMultifactorSecretRequest represents the AppleCredentialsMultifactorSecretRequest schema from the OpenAPI specification
type AppleCredentialsMultifactorSecretRequest struct {
	Data map[string]interface{} `json:"data,omitempty"` // shared connection details
	Displayname string `json:"displayName,omitempty"` // display name of shared connection
	Servicetype string `json:"serviceType"` // service type of shared connection can be apple|gitlab|googleplay|jira|applecertificate
	Credentialtype string `json:"credentialType,omitempty"` // credential type of the shared connection. Values can be credentials|certificate
}

// CrashGroupLanguage represents the CrashGroupLanguage schema from the OpenAPI specification
type CrashGroupLanguage struct {
	Crash_count int64 `json:"crash_count,omitempty"` // Count of languages.
	Language_name string `json:"language_name,omitempty"` // Language name.
}

// Branch represents the Branch schema from the OpenAPI specification
type Branch struct {
	Commit interface{} `json:"commit"`
	Name string `json:"name"` // The branch name
}

// CodePushReleaseLabel represents the CodePushReleaseLabel schema from the OpenAPI specification
type CodePushReleaseLabel struct {
	Label string `json:"label,omitempty"`
}

// PatchReleaseAssetIdRequest represents the PatchReleaseAssetIdRequest schema from the OpenAPI specification
type PatchReleaseAssetIdRequest struct {
	Ipa_uuids string `json:"ipa_uuids,omitempty"` // The ipa UUIDs for this release, as a serialized JSON array
	Package_asset_id string `json:"package_asset_id"` // The release new package id in ACFUS
	Upload_id string `json:"upload_id"` // The release upload id used to upload the release
}

// TestGDPRResourceList represents the TestGDPRResourceList schema from the OpenAPI specification
type TestGDPRResourceList struct {
	Resources []map[string]interface{} `json:"resources,omitempty"`
}

// RepoConfigPostRequest represents the RepoConfigPostRequest schema from the OpenAPI specification
type RepoConfigPostRequest struct {
	External_user_id string `json:"external_user_id,omitempty"` // The external user id from the repository provider. Required for GitLab.com repositories
	Repo_id string `json:"repo_id,omitempty"` // The repository id from the repository provider. Required for repositories connected from GitHub App and GitLab.com
	Repo_url string `json:"repo_url,omitempty"` // The repository's git url, must be a HTTPS URL
	Service_connection_id string `json:"service_connection_id,omitempty"` // The id of the service connection (private). Required for GitLab self-hosted repositories
}

// Thread represents the Thread schema from the OpenAPI specification
type Thread struct {
	Crashed bool `json:"crashed,omitempty"` // True if this thread crashed
	Exception Exception `json:"exception,omitempty"` // a exception
	Frames []interface{} `json:"frames"` // frames of that thread
	Platform string `json:"platform,omitempty"` // SDK/Platform this thread is beeing generated from
	Relevant bool `json:"relevant,omitempty"` // Shows if a thread is relevant or not. Is false if all frames are non relevant, otherwise true
	Title string `json:"title"` // name of the thread
}

// FileValidationDetails represents the FileValidationDetails schema from the OpenAPI specification
type FileValidationDetails struct {
	Certificateuploadid string `json:"certificateUploadId,omitempty"`
	P12password string `json:"p12password"`
}

// Error represents the Error schema from the OpenAPI specification
type Error struct {
	ErrorField map[string]interface{} `json:"error,omitempty"`
}

// XamarinBranchConfigurationProperties represents the XamarinBranchConfigurationProperties schema from the OpenAPI specification
type XamarinBranchConfigurationProperties struct {
	Slnpath string `json:"slnPath,omitempty"`
	P12pwd string `json:"p12Pwd,omitempty"`
	Symlink string `json:"symlink,omitempty"` // Symlink of the SDK Bundle and Mono installation. The build will use the associated Mono bundled with related Xamarin SDK. If both symlink and monoVersion or sdkBundle are passed, the symlink is taking precedence. If non-existing symlink is passed, the current stable Mono version will be configured for building.
	P12file string `json:"p12File,omitempty"`
	Provprofile string `json:"provProfile,omitempty"`
	Monoversion string `json:"monoVersion,omitempty"`
	Sdkbundle string `json:"sdkBundle,omitempty"`
	Args string `json:"args,omitempty"`
	Configuration string `json:"configuration,omitempty"`
	Issimbuild bool `json:"isSimBuild,omitempty"`
}

// Module represents the Module schema from the OpenAPI specification
type Module struct {
}

// AzureSubscriptionPatchRequest represents the AzureSubscriptionPatchRequest schema from the OpenAPI specification
type AzureSubscriptionPatchRequest struct {
	Is_billing bool `json:"is_billing"` // If the subscription is used for billing
}

// SubmitDsrOperationRequest represents the SubmitDsrOperationRequest schema from the OpenAPI specification
type SubmitDsrOperationRequest struct {
	Operationid string `json:"operationId,omitempty"` // The DSR operation ID provided by the GDPR coordinator. Used for tracking only.
	Request string `json:"request,omitempty"`
	Requestid string `json:"requestId,omitempty"` // Request ID provided by the GDPR coordinator. Used for tracking.
	Accountid string `json:"accountId,omitempty"`
	Appid string `json:"appId,omitempty"`
}

// DistributionGroupAppsDeleteRequest represents the DistributionGroupAppsDeleteRequest schema from the OpenAPI specification
type DistributionGroupAppsDeleteRequest struct {
	Apps []interface{} `json:"apps,omitempty"` // The list of apps to delete from the distribution group
}

// AppGroupResponse represents the AppGroupResponse schema from the OpenAPI specification
type AppGroupResponse struct {
	Group_id string `json:"group_id"` // The unique ID (UUID) of the group that the app belongs to
	Id string `json:"id"` // The unique ID (UUID) of the app
	Name string `json:"name"` // The name of the app used in URLs
	Os string `json:"os"` // The OS the app will be running on
	Platform string `json:"platform"` // The platform of the app
	Release_type string `json:"release_type,omitempty"` // A one-word descriptive release-type value that starts with a capital letter but is otherwise lowercase
	Display_name string `json:"display_name,omitempty"` // The display name of the app
}

// ErrorDetails represents the ErrorDetails schema from the OpenAPI specification
type ErrorDetails struct {
	Code string `json:"code"`
	Message string `json:"message"`
}

// GetReleasesContainer represents the GetReleasesContainer schema from the OpenAPI specification
type GetReleasesContainer struct {
	Releases []interface{} `json:"releases"`
}

// AnalyticsReleasesResponse represents the AnalyticsReleasesResponse schema from the OpenAPI specification
type AnalyticsReleasesResponse struct {
	Releases []map[string]interface{} `json:"releases,omitempty"`
}

// BuildParams represents the BuildParams schema from the OpenAPI specification
type BuildParams struct {
	Sourceversion string `json:"sourceVersion,omitempty"` // Version to build which represents the full Git commit reference
	Debug bool `json:"debug,omitempty"` // Run build in debug mode
}

// ErrorsSearchResult represents the ErrorsSearchResult schema from the OpenAPI specification
type ErrorsSearchResult struct {
	Errors []map[string]interface{} `json:"errors,omitempty"`
	Hasmoreresults bool `json:"hasMoreResults,omitempty"`
}

// LegacyDeploymentMetricsResponse represents the LegacyDeploymentMetricsResponse schema from the OpenAPI specification
type LegacyDeploymentMetricsResponse struct {
	Metrics map[string]interface{} `json:"metrics,omitempty"` // Object containing a property named after each release label, which contains an object that contains that release's metrics.
}

// DeviceInfoResponse represents the DeviceInfoResponse schema from the OpenAPI specification
type DeviceInfoResponse struct {
	Status string `json:"status"` // The provisioning status of the device.
	Registered_at string `json:"registered_at,omitempty"` // Timestamp of when the device was registered in ISO format.
	Serial string `json:"serial,omitempty"` // The device's serial number. Always empty or undefined at present.
	Os_build string `json:"os_build"` // The last known OS version running on the device
	Os_version string `json:"os_version"` // The last known OS version running on the device
	Owner_id string `json:"owner_id,omitempty"` // The user ID of the device owner.
	Udid string `json:"udid"` // The Unique Device IDentifier of the device
	Device_name string `json:"device_name"` // The device description, in the format "iPhone 7 Plus (A1784)"
	Full_device_name string `json:"full_device_name,omitempty"` // A combination of the device model name and the owner name.
	Model string `json:"model"` // The model identifier of the device, in the format iDeviceM,N
	Imei string `json:"imei,omitempty"` // The device's International Mobile Equipment Identity number. Always empty or undefined at present.
}

// UserNameUpdateRequest represents the UserNameUpdateRequest schema from the OpenAPI specification
type UserNameUpdateRequest struct {
	Name string `json:"name,omitempty"` // The new, unique name that is used to identify.
}

// V2MissingSymbolCrashGroupsInfoResponse represents the V2MissingSymbolCrashGroupsInfoResponse schema from the OpenAPI specification
type V2MissingSymbolCrashGroupsInfoResponse struct {
	Total_crash_count int `json:"total_crash_count"` // total number of crashes for all missing symbol groups
}

// EventProperties represents the EventProperties schema from the OpenAPI specification
type EventProperties struct {
	Event_properties []string `json:"event_properties,omitempty"`
}

// DistributionGroupPatchRequest represents the DistributionGroupPatchRequest schema from the OpenAPI specification
type DistributionGroupPatchRequest struct {
	Is_public bool `json:"is_public,omitempty"` // Whether the distribution group is public
	Name string `json:"name,omitempty"` // The name of the distribution group
}

// ReleaseProvisionResponse represents the ReleaseProvisionResponse schema from the OpenAPI specification
type ReleaseProvisionResponse struct {
	Status_url string `json:"status_url,omitempty"` // The url to check provisioning status.
}

// StoresDetails represents the StoresDetails schema from the OpenAPI specification
type StoresDetails struct {
	Id string `json:"id,omitempty"` // ID identifying a unique distribution store.
	Is_latest bool `json:"is_latest,omitempty"` // Is the containing release the latest one in this distribution store.
	Name string `json:"name,omitempty"` // A name identifying a unique distribution store.
	Publishing_status string `json:"publishing_status,omitempty"` // A status identifying the status of release in the distribution store.
	TypeField string `json:"type,omitempty"` // A type identifying the type of distribution store.
}

// PrivateJiraConnectionSecretResponse represents the PrivateJiraConnectionSecretResponse schema from the OpenAPI specification
type PrivateJiraConnectionSecretResponse struct {
	Id string `json:"id"` // id of the shared connection
	Is2fa bool `json:"is2FA,omitempty"` // if the account is a 2FA account or not
	Isvalid bool `json:"isValid,omitempty"` // whether the credentials are valid or not
	Servicetype string `json:"serviceType"` // service type of shared connection can be apple|gitlab|googleplay|jira|applecertificate
	Displayname string `json:"displayName,omitempty"` // display name of shared connection
}

// LogFlowDevice represents the LogFlowDevice schema from the OpenAPI specification
type LogFlowDevice struct {
	App_version string `json:"app_version"` // Application version name, e.g. 1.1.0
	Carrier_country string `json:"carrier_country,omitempty"` // Carrier country.
	Live_update_deployment_key string `json:"live_update_deployment_key,omitempty"` // Identifier of environment that current application release belongs to, deployment key then maps to environment like Production, Staging.
	Wrapper_sdk_version string `json:"wrapper_sdk_version,omitempty"` // Version of the wrapper SDK in semver format. When the SDK is embedding another base SDK (for example Xamarin.Android wraps Android), the Xamarin specific version is populated into this field while sdkVersion refers to the original Android SDK.
	Model string `json:"model,omitempty"` // Device model (example: iPad2,3).
	App_build string `json:"app_build"` // The app's build number, e.g. 42.
	Os_api_level int `json:"os_api_level,omitempty"` // API level when applicable like in Android (example: 15).
	Time_zone_offset int `json:"time_zone_offset"` // The offset in minutes from UTC for the device time zone, including daylight savings time.
	Oem_name string `json:"oem_name,omitempty"` // Device manufacturer (example: HTC).
	Sdk_name string `json:"sdk_name"` // Name of the SDK. Consists of the name of the SDK and the platform, e.g. "appcenter.ios", "hockeysdk.android".
	Live_update_release_label string `json:"live_update_release_label,omitempty"` // Label that is used to identify application code 'version' released via Live Update beacon running on device
	Os_version string `json:"os_version"` // OS version (example: 9.3.0).
	Screen_size string `json:"screen_size,omitempty"` // Screen size of the device in pixels (example: 640x480).
	Carrier_code string `json:"carrier_code,omitempty"` // Carrier country code (for mobile devices).
	Carrier_name string `json:"carrier_name,omitempty"` // Carrier name (for mobile devices).
	Live_update_package_hash string `json:"live_update_package_hash,omitempty"` // Hash of all files (ReactNative or Cordova) deployed to device via LiveUpdate beacon. Helps identify the Release version on device or need to download updates in future.
	Os_name string `json:"os_name"` // OS name (example: iOS). The following OS names are standardized (non-exclusive): Android, iOS, macOS, tvOS, Windows.
	Wrapper_runtime_version string `json:"wrapper_runtime_version,omitempty"` // Version of the wrapper technology framework (Xamarin runtime version or ReactNative or Cordova etc...). See wrapper_sdk_name to see if this version refers to Xamarin or ReactNative or other.
	Locale string `json:"locale"` // Language code (example: en_US).
	Wrapper_sdk_name string `json:"wrapper_sdk_name,omitempty"` // Name of the wrapper SDK. Consists of the name of the SDK and the wrapper platform, e.g. "appcenter.xamarin", "hockeysdk.cordova".
	Sdk_version string `json:"sdk_version"` // Version of the SDK in semver format, e.g. "1.2.0" or "0.12.3-alpha.1".
	Os_build string `json:"os_build,omitempty"` // OS build code (example: LMY47X).
	App_namespace string `json:"app_namespace,omitempty"` // The bundle identifier, package identifier, or namespace, depending on what the individual plattforms use, .e.g com.microsoft.example.
}

// ToolsetProjects represents the ToolsetProjects schema from the OpenAPI specification
type ToolsetProjects struct {
	Commit string `json:"commit,omitempty"` // The commit hash of the analyzed commit
	Javascript interface{} `json:"javascript,omitempty"`
	Testcloud interface{} `json:"testcloud,omitempty"`
	Uwp interface{} `json:"uwp,omitempty"`
	Xamarin interface{} `json:"xamarin,omitempty"`
	Xcode interface{} `json:"xcode,omitempty"`
	Android interface{} `json:"android,omitempty"`
	Buildscripts interface{} `json:"buildscripts,omitempty"` // A collection of detected pre/post buildscripts for current platform toolset
}

// DistributionGroupsUserVerifyRequest represents the DistributionGroupsUserVerifyRequest schema from the OpenAPI specification
type DistributionGroupsUserVerifyRequest struct {
	Distribution_group_ids []string `json:"distribution_group_ids"` // An array of distribution group ids
}

// NodeVersion represents the NodeVersion schema from the OpenAPI specification
type NodeVersion struct {
	Current bool `json:"current,omitempty"` // If the Node version is default for AppCenter
	Name string `json:"name,omitempty"` // The version name
}

// TestCloudStartTestRunOptions represents the TestCloudStartTestRunOptions schema from the OpenAPI specification
type TestCloudStartTestRunOptions struct {
	Test_series string `json:"test_series,omitempty"` // Name of the test series.
	Device_selection string `json:"device_selection"` // Device selection string.
	Language string `json:"language,omitempty"` // Language that should be used to run tests.
	Locale string `json:"locale,omitempty"` // Locale that should be used to run tests.
	Test_framework string `json:"test_framework"` // Test framework used by tests.
	Test_parameters map[string]interface{} `json:"test_parameters,omitempty"` // A JSON dictionary with additional test parameters
}

// ErrorDateTimePercentages represents the ErrorDateTimePercentages schema from the OpenAPI specification
type ErrorDateTimePercentages struct {
	Datetime string `json:"datetime,omitempty"` // the ISO 8601 datetime
	Percentage float64 `json:"percentage,omitempty"` // percentage of the object
}

// ReleaseDailySessions represents the ReleaseDailySessions schema from the OpenAPI specification
type ReleaseDailySessions struct {
	Sessions []map[string]interface{} `json:"sessions,omitempty"` // Sessions per day.
	Totalsessioncounts int64 `json:"totalSessionCounts,omitempty"`
	Avgsessionsperday float64 `json:"avgSessionsPerDay,omitempty"`
}

// AppleCertificateDetails represents the AppleCertificateDetails schema from the OpenAPI specification
type AppleCertificateDetails struct {
	Password string `json:"password"` // The password for the certificate
	Base64certificate string `json:"base64Certificate"` // The certificate contents in base 64 encoded string
}

// EmailVerificationRequest represents the EmailVerificationRequest schema from the OpenAPI specification
type EmailVerificationRequest struct {
	Token string `json:"token"` // The verification token that was sent to the user
}

// LegacyUpdateCheckResponse represents the LegacyUpdateCheckResponse schema from the OpenAPI specification
type LegacyUpdateCheckResponse struct {
	Updateinfo interface{} `json:"updateInfo"`
}

// OrganizationRequest represents the OrganizationRequest schema from the OpenAPI specification
type OrganizationRequest struct {
	Name string `json:"name,omitempty"` // The name of the organization used in URLs
	Display_name string `json:"display_name,omitempty"` // The display name of the organization
}

// AvailableVersions represents the AvailableVersions schema from the OpenAPI specification
type AvailableVersions struct {
	Total_count int64 `json:"total_count,omitempty"` // The full number of versions across all pages.
	Versions []string `json:"versions,omitempty"` // List of available versions.
}

// AppleSecretDetailsResponse represents the AppleSecretDetailsResponse schema from the OpenAPI specification
type AppleSecretDetailsResponse struct {
	Username string `json:"username"` // username to connect to apple store
}

// GDPRPurgeVerifyRequest represents the GDPRPurgeVerifyRequest schema from the OpenAPI specification
type GDPRPurgeVerifyRequest struct {
	Key string `json:"key"` // deployment key
	Id string `json:"id"` // deployment id
}

// DistributionStoreWithoutIsLatest represents the DistributionStoreWithoutIsLatest schema from the OpenAPI specification
type DistributionStoreWithoutIsLatest struct {
	Id string `json:"id"` // ID identifying a unique distribution store.
	Name string `json:"name,omitempty"` // A name identifying a unique distribution store.
	Publishing_status string `json:"publishing_status,omitempty"` // publishing status of the release in the store.
	TypeField string `json:"type,omitempty"` // type of the distribution store currently stores type can be intune, googleplay or windows.
}

// Toolsets represents the Toolsets schema from the OpenAPI specification
type Toolsets struct {
	Node []interface{} `json:"node,omitempty"` // A list of Node versions
	Xamarin []interface{} `json:"xamarin,omitempty"` // A list of Xamarin SDK bundles
	Xcode []interface{} `json:"xcode,omitempty"` // A list of Xcode versions
}

// AlertingJiraBugtrackerSettings represents the AlertingJiraBugtrackerSettings schema from the OpenAPI specification
type AlertingJiraBugtrackerSettings struct {
	Callback_url string `json:"callback_url,omitempty"`
	Owner_name string `json:"owner_name"`
	TypeField string `json:"type"` // type of bugtracker
}

// UserProfileResponsev2 represents the UserProfileResponsev2 schema from the OpenAPI specification
type UserProfileResponsev2 struct {
	Settings map[string]interface{} `json:"settings,omitempty"` // The user's settings
	Created_at string `json:"created_at,omitempty"` // The created date of the user
	Name string `json:"name"` // The unique name that is used to identify the user.
	Next_nps_survey_date string `json:"next_nps_survey_date,omitempty"` // The date in the future when the user should be checked again for NPS eligibility
	Feature_flags []string `json:"feature_flags,omitempty"` // The feature flags that are enabled for this user
	Id string `json:"id"` // The unique id (UUID) of the user
	Session_hash string `json:"session_hash,omitempty"` // The session hash of the user
	Avatar_url string `json:"avatar_url,omitempty"` // The avatar URL of the user
	Can_change_password bool `json:"can_change_password,omitempty"` // User is required to send an old password in order to change the password.
	Display_name string `json:"display_name"` // The full name of the user. Might for example be first and last name
	Origin string `json:"origin"` // The creation origin of this user
	Admin_role string `json:"admin_role,omitempty"` // The new admin_role
	Email string `json:"email"` // The email address of the user
}

// AvailableAppBuilds represents the AvailableAppBuilds schema from the OpenAPI specification
type AvailableAppBuilds struct {
	App_builds []string `json:"app_builds,omitempty"` // List of available app builds.
}

// HandledErrorDetails represents the HandledErrorDetails schema from the OpenAPI specification
type HandledErrorDetails struct {
	Country string `json:"country,omitempty"`
	Errorid string `json:"errorId,omitempty"`
	Ostype string `json:"osType,omitempty"`
	Hasbreadcrumbs bool `json:"hasBreadcrumbs,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Userid string `json:"userId,omitempty"`
	Hasattachments bool `json:"hasAttachments,omitempty"`
	Language string `json:"language,omitempty"`
	Osversion string `json:"osVersion,omitempty"`
	Devicename string `json:"deviceName,omitempty"`
}

// HandledErrorLog represents the HandledErrorLog schema from the OpenAPI specification
type HandledErrorLog struct {
	Device map[string]interface{} `json:"device"` // Device characteristics.
	Install_id string `json:"install_id"` // Install ID.
	Timestamp string `json:"timestamp"` // Log creation timestamp.
	TypeField string `json:"type"` // Log type.
}

// RepoConfig represents the RepoConfig schema from the OpenAPI specification
type RepoConfig struct {
	Service_connection_id string `json:"service_connection_id,omitempty"` // The id of the service connection (private). Required for GitLab self-hosted repositories
	External_user_id string `json:"external_user_id,omitempty"` // The external user id from the repository provider. Required for GitLab.com repositories
	Repo_id string `json:"repo_id,omitempty"` // The repository id from the repository provider. Required for repositories connected from GitHub App and GitLab.com
	Repo_url string `json:"repo_url,omitempty"` // The repository's git url, must be a HTTPS URL
	Installation_id string `json:"installation_id,omitempty"` // The GitHub App Installation id. Required for repositories connected from GitHub App
}

// EventLogDiagnostics represents the EventLogDiagnostics schema from the OpenAPI specification
type EventLogDiagnostics struct {
	Install_id string `json:"install_id"` // Install ID.
	Timestamp string `json:"timestamp"` // Log creation timestamp.
	TypeField string `json:"type"` // Log type.
	Device map[string]interface{} `json:"device"` // Device characteristics.
	Properties map[string]interface{} `json:"properties,omitempty"` // Additional key/value pair parameters.
}

// LogFlowCustomProperty represents the LogFlowCustomProperty schema from the OpenAPI specification
type LogFlowCustomProperty struct {
	TypeField string `json:"type"`
	Name string `json:"name"`
}

// AppDeleteRequest represents the AppDeleteRequest schema from the OpenAPI specification
type AppDeleteRequest struct {
	Name string `json:"name"` // The name of the app to be deleted from the distribution group
}

// IntuneTargetAudienceResponse represents the IntuneTargetAudienceResponse schema from the OpenAPI specification
type IntuneTargetAudienceResponse struct {
	Id string `json:"id,omitempty"` // ID for the target audience/group.
	Name string `json:"name,omitempty"` // display name for the target audience/group
}

// UserInteractionMetricsResponse represents the UserInteractionMetricsResponse schema from the OpenAPI specification
type UserInteractionMetricsResponse struct {
	Less_than_100_apps bool `json:"less_than_100_apps,omitempty"` // check if the user has less than 100 apps.
	Has_more_than_1_release bool `json:"has_more_than_1_release,omitempty"` // check if the user's whole apps has more than 1 releases.
}

// IntuneAppCategory represents the IntuneAppCategory schema from the OpenAPI specification
type IntuneAppCategory struct {
	Name string `json:"name,omitempty"` // display name for the app category
}

// AlertBugTrackerReposResult represents the AlertBugTrackerReposResult schema from the OpenAPI specification
type AlertBugTrackerReposResult struct {
	Repo_type string `json:"repo_type,omitempty"`
	Repositories []map[string]interface{} `json:"repositories"`
}

// DiagnosticsException represents the DiagnosticsException schema from the OpenAPI specification
type DiagnosticsException struct {
	Reason string `json:"reason,omitempty"` // Reason of the exception
	Relevant bool `json:"relevant,omitempty"` // relevant exception (crashed)
	TypeField string `json:"type,omitempty"` // Type of the exception (NSSomethingException, NullPointerException)
	Frames []map[string]interface{} `json:"frames"` // frames of the excetpion
	Inner_exceptions []DiagnosticsException `json:"inner_exceptions,omitempty"`
	Platform string `json:"platform,omitempty"` // SDK/Platform this thread is beeing generated from
}

// DestinationId represents the DestinationId schema from the OpenAPI specification
type DestinationId struct {
	Id string `json:"id,omitempty"` // Id of a distribution group / store. The release will be associated with this distribution group / store. If the distribution group / store doesn't exist a 400 is returned. If both distribution group / store name and id are passed, the id is taking precedence.
	Name string `json:"name,omitempty"` // Name of a distribution group / distribution store. The release will be associated with this distribution group or store. If the distribution group / store doesn't exist a 400 is returned. If both distribution group / store name and id are passed, the id is taking precedence.
}

// DeviceFrameDefinition represents the DeviceFrameDefinition schema from the OpenAPI specification
type DeviceFrameDefinition struct {
	Frameurl string `json:"frameUrl,omitempty"`
	Height float64 `json:"height,omitempty"`
	Screen []float64 `json:"screen,omitempty"`
	Width float64 `json:"width,omitempty"`
}

// AdministeredOrgsResponse represents the AdministeredOrgsResponse schema from the OpenAPI specification
type AdministeredOrgsResponse struct {
	Organizations interface{} `json:"organizations"`
}

// Modules represents the Modules schema from the OpenAPI specification
type Modules struct {
	Modules map[string]interface{} `json:"modules,omitempty"`
}

// DistributionGroupAADGroupsDeleteRequest represents the DistributionGroupAADGroupsDeleteRequest schema from the OpenAPI specification
type DistributionGroupAADGroupsDeleteRequest struct {
	Aad_group_ids []string `json:"aad_group_ids,omitempty"` // The list of aad group ids
}

// Event represents the Event schema from the OpenAPI specification
type Event struct {
	Previous_device_count int64 `json:"previous_device_count,omitempty"` // The device count of previous time range of the event.
	Count int64 `json:"count,omitempty"`
	Count_per_device float64 `json:"count_per_device,omitempty"`
	Count_per_session float64 `json:"count_per_session,omitempty"`
	Device_count int64 `json:"device_count,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Previous_count int64 `json:"previous_count,omitempty"` // The event count of previous time range of the event.
}

// ErrorResponsev2 represents the ErrorResponsev2 schema from the OpenAPI specification
type ErrorResponsev2 struct {
	ErrorField interface{} `json:"error"`
}

// DataSubjectRightQueueInfo represents the DataSubjectRightQueueInfo schema from the OpenAPI specification
type DataSubjectRightQueueInfo struct {
	Queuename string `json:"queueName"`
	Sasuri string `json:"sasUri"`
	Expiresat string `json:"expiresAt"`
}

// ApiTokenResponse represents the ApiTokenResponse schema from the OpenAPI specification
type ApiTokenResponse struct {
	Encrypted_token string `json:"encrypted_token,omitempty"` // The encrypted value of a token. This value will only be returned for token of type in_app_update.
	Id string `json:"id"` // The unique id (UUID) of the api token
	Scope []string `json:"scope,omitempty"` // The token's scope. A list of allowed roles.
	Created_at string `json:"created_at"` // The creation time
	Description string `json:"description,omitempty"` // The description of the token
}

// AvailableVersionsDiagnostics represents the AvailableVersionsDiagnostics schema from the OpenAPI specification
type AvailableVersionsDiagnostics struct {
	Total_count int64 `json:"total_count,omitempty"` // The full number of versions across all pages.
	Versions []string `json:"versions,omitempty"` // List of available versions.
}

// BillingPlansChangeTypeResponse represents the BillingPlansChangeTypeResponse schema from the OpenAPI specification
type BillingPlansChangeTypeResponse struct {
	Result string `json:"result,omitempty"`
}

// EventCount represents the EventCount schema from the OpenAPI specification
type EventCount struct {
	Total_count int64 `json:"total_count,omitempty"`
	Count []interface{} `json:"count,omitempty"`
	Previous_total_count int64 `json:"previous_total_count,omitempty"`
}

// GooglePlayConnectionSecretResponse represents the GooglePlayConnectionSecretResponse schema from the OpenAPI specification
type GooglePlayConnectionSecretResponse struct {
	Credentialtype string `json:"credentialType,omitempty"` // the type of the credential
	Displayname string `json:"displayName"` // display name of shared connection
	Servicetype string `json:"serviceType"` // service type of shared connection can be apple|gitlab|googleplay|jira
}

// GitHubPullRequestLite represents the GitHubPullRequestLite schema from the OpenAPI specification
type GitHubPullRequestLite struct {
	Base interface{} `json:"base,omitempty"` // The lite version of GitHub branch
	Head interface{} `json:"head,omitempty"` // The lite version of GitHub branch
}

// Owner represents the Owner schema from the OpenAPI specification
type Owner struct {
	Id string `json:"id"` // The unique id (UUID) of the owner
	Name string `json:"name"` // The unique name that used to identify the owner
	TypeField string `json:"type"` // The owner type. Can either be 'org' or 'user'
	Avatar_url string `json:"avatar_url,omitempty"` // The avatar URL of the owner
	Display_name string `json:"display_name"` // The owner's display name
	Email string `json:"email,omitempty"` // The owner's email address
}

// BillingPlan represents the BillingPlan schema from the OpenAPI specification
type BillingPlan struct {
	Parentid string `json:"parentId,omitempty"`
	Paymentsource string `json:"paymentSource,omitempty"` // Service that receives payments for this billing plan.
	Price float64 `json:"price,omitempty"` // Price of the Billing Plan
	Service string `json:"service,omitempty"` // Name of the service that the plan applies to.
	Version string `json:"version,omitempty"` // Version of the Billing Plan schema
	Attributes map[string]interface{} `json:"attributes,omitempty"` // Collection of attribute values.
	Id string `json:"id,omitempty"` // The Billing Plan ID
	Limits map[string]interface{} `json:"limits,omitempty"` // A collection of named numeric values
}

// GitHubAccount represents the GitHubAccount schema from the OpenAPI specification
type GitHubAccount struct {
	Accounttype string `json:"accountType,omitempty"` // Type of GitHub account
	Id int `json:"id,omitempty"` // Id of GitHub account
}

// TestCloudFileHash represents the TestCloudFileHash schema from the OpenAPI specification
type TestCloudFileHash struct {
	Filetype string `json:"fileType"` // Type of the file
	Relativepath string `json:"relativePath"` // Relative path of the file
	Checksum string `json:"checksum"` // SHA256 hash of the file
}

// ApiTokenGetUserResponse represents the ApiTokenGetUserResponse schema from the OpenAPI specification
type ApiTokenGetUserResponse struct {
	Token_id string `json:"token_id"` // The token's unique id (UUID)
	Token_scope []string `json:"token_scope"` // The token's scope. A list of allowed roles.
	User_email string `json:"user_email"` // The user email
	User_id string `json:"user_id"` // The unique id (UUID) of the user
	User_origin string `json:"user_origin"` // The creation origin of the user who created this api token
}

// ErrorLocation represents the ErrorLocation schema from the OpenAPI specification
type ErrorLocation struct {
	Uri string `json:"uri,omitempty"`
}

// HandledErrors represents the HandledErrors schema from the OpenAPI specification
type HandledErrors struct {
	Errors []map[string]interface{} `json:"errors,omitempty"` // Errors list.
	Nextlink string `json:"nextLink,omitempty"`
}

// DistributionGroup represents the DistributionGroup schema from the OpenAPI specification
type DistributionGroup struct {
	Id string `json:"id"` // ID identifying a unique distribution group.
	Name string `json:"name,omitempty"` // A name identifying a unique distribution group.
	Is_latest bool `json:"is_latest,omitempty"` // Is the containing release the latest one in this distribution group.
}

// LegacyCodePushReleaseInfo represents the LegacyCodePushReleaseInfo schema from the OpenAPI specification
type LegacyCodePushReleaseInfo struct {
	Appversion string `json:"appVersion,omitempty"`
	Description string `json:"description,omitempty"`
	Isdisabled bool `json:"isDisabled,omitempty"`
	Ismandatory bool `json:"isMandatory,omitempty"`
	Rollout int `json:"rollout,omitempty"`
}

// GrantAdminRoleRequest represents the GrantAdminRoleRequest schema from the OpenAPI specification
type GrantAdminRoleRequest struct {
	Admin_role string `json:"admin_role"` // The new admin_role
}

// Symbol represents the Symbol schema from the OpenAPI specification
type Symbol struct {
	Status string `json:"status"` // Whether the symbol is ignored.
	Alternate_symbol_ids []string `json:"alternate_symbol_ids"` // The other symbols in the same file
	App_id string `json:"app_id"` // The application that this symbol belongs to
	TypeField string `json:"type"` // The type of the symbol for the current symbol upload
	Url string `json:"url"` // The path name of the symbol file in blob storage
	Version string `json:"version,omitempty"` // The version number. Optional for Apple. Required for Android.
	Platform string `json:"platform"` // The platform that this symbol is associated with
	Symbol_upload_id string `json:"symbol_upload_id"` // The id of the symbol upload this symbol belongs to.
	Symbol_id string `json:"symbol_id"` // The unique id for this symbol (uuid)
	Build string `json:"build,omitempty"` // The build number. Optional for Apple. Required for Android.
	Origin string `json:"origin"` // The origin of the symbol file
}

// V2MissingSymbol represents the V2MissingSymbol schema from the OpenAPI specification
type V2MissingSymbol struct {
	Platform string `json:"platform,omitempty"` // symbol plarform
	Status string `json:"status"` // symbol status
	Symbol_id string `json:"symbol_id"` // symbol id
	Name string `json:"name"` // symbol name
}

// ApiTokenResponsev2 represents the ApiTokenResponsev2 schema from the OpenAPI specification
type ApiTokenResponsev2 struct {
	Created_at string `json:"created_at"` // The creation time
	Description string `json:"description,omitempty"` // The description of the token
	Encrypted_token string `json:"encrypted_token,omitempty"` // The encrypted value of a token. This value will only be returned for token of type in_app_update.
	Id string `json:"id"` // The unique id (UUID) of the api token
	Scope []string `json:"scope,omitempty"` // The token's scope. A list of allowed roles.
}

// ReleasePublishErrorResponse represents the ReleasePublishErrorResponse schema from the OpenAPI specification
type ReleasePublishErrorResponse struct {
	Is_log_available bool `json:"is_log_available,omitempty"` // boolean property to tell if logs are available for download
	Message string `json:"message,omitempty"` // error Details
}

// GDPRPurgeVerifyResponse represents the GDPRPurgeVerifyResponse schema from the OpenAPI specification
type GDPRPurgeVerifyResponse struct {
	Success bool `json:"success"` // indicate whether GDPR purge operation succeeds or not
}

// DiagnosticsThread represents the DiagnosticsThread schema from the OpenAPI specification
type DiagnosticsThread struct {
	Exception DiagnosticsException `json:"exception,omitempty"` // a exception
	Frames []map[string]interface{} `json:"frames"` // frames of that thread
	Platform string `json:"platform,omitempty"` // SDK/Platform this thread is beeing generated from
	Relevant bool `json:"relevant,omitempty"` // Shows if a thread is relevant or not. Is false if all frames are non relevant, otherwise true
	Title string `json:"title"` // name of the thread
	Crashed bool `json:"crashed,omitempty"` // True if this thread crashed
}

// AddAppTesterRequest represents the AddAppTesterRequest schema from the OpenAPI specification
type AddAppTesterRequest struct {
	Release_id int `json:"release_id"` // The ID of the release the user was added to
	User_id string `json:"user_id"` // The user ID of the tester that needs to be added
}

// DeviceModel represents the DeviceModel schema from the OpenAPI specification
type DeviceModel struct {
	Availabilitycount float64 `json:"availabilityCount,omitempty"`
	Releasedate string `json:"releaseDate,omitempty"`
	Screensize interface{} `json:"screenSize,omitempty"` // Physical device screen dimensions
	Cpu interface{} `json:"cpu,omitempty"` // CPU data for device
	Deviceframe interface{} `json:"deviceFrame,omitempty"`
	Platform string `json:"platform,omitempty"`
	Resolution interface{} `json:"resolution,omitempty"` // Device screen resolution
	Screenrotation float64 `json:"screenRotation,omitempty"`
	Formfactor string `json:"formFactor,omitempty"`
	Manufacturer string `json:"manufacturer,omitempty"`
	Model string `json:"model,omitempty"`
	Name string `json:"name,omitempty"`
	Dimensions interface{} `json:"dimensions,omitempty"` // Physical device dimensions
	Memory interface{} `json:"memory,omitempty"` // Memory data for device
}

// AppleConnectionSecretResponse represents the AppleConnectionSecretResponse schema from the OpenAPI specification
type AppleConnectionSecretResponse struct {
	Displayname string `json:"displayName"` // display name of shared connection
	Servicetype string `json:"serviceType"` // service type of shared connection can be apple|gitlab|googleplay|jira
	Credentialtype string `json:"credentialType,omitempty"` // the type of the credential
}

// DeviceSetOwner represents the DeviceSetOwner schema from the OpenAPI specification
type DeviceSetOwner struct {
	Name string `json:"name"` // Name of the account
	TypeField string `json:"type"` // Type of account
	Displayname string `json:"displayName,omitempty"` // Display name of the account
	Id string `json:"id"` // Account ID
}

// TestCloudToolset represents the TestCloudToolset schema from the OpenAPI specification
type TestCloudToolset struct {
	Projects []interface{} `json:"projects"` // The TestCloud projects detected
}

// Audience represents the Audience schema from the OpenAPI specification
type Audience struct {
	Estimated_count int64 `json:"estimated_count,omitempty"` // Estimated audience size.
	Name string `json:"name,omitempty"` // Audience name.
	State string `json:"state,omitempty"` // Audience state.
	Definition string `json:"definition,omitempty"` // Audience definition in OData format.
	Description string `json:"description,omitempty"` // Audience description.
}

// AccessKey represents the AccessKey schema from the OpenAPI specification
type AccessKey struct {
	Issession bool `json:"isSession,omitempty"` // Legacy property which indicate if accessKey was created from session
	Name string `json:"name,omitempty"` // Key of access key
	Createdby string `json:"createdBy,omitempty"` // Account name of creator.
	Createdtime float64 `json:"createdTime,omitempty"` // Created time of access key
	Description string `json:"description,omitempty"` // Description of access key
	Expires float64 `json:"expires,omitempty"` // Time of expiry of access key
	Friendlyname string `json:"friendlyName,omitempty"` // Friendly name of access key
	Id string `json:"id,omitempty"` // Id of accessKey
}

// LegacyAccount represents the LegacyAccount schema from the OpenAPI specification
type LegacyAccount struct {
	Email string `json:"email,omitempty"` // The calling user's email.
	Linkedproviders []string `json:"linkedProviders,omitempty"` // Array of linked authentication providers associated with the account.
	Name string `json:"name,omitempty"` // The account name of the calling user.
}

// MetricsValues represents the MetricsValues schema from the OpenAPI specification
type MetricsValues struct {
}

// ResetPasswordUsingTokenRequest represents the ResetPasswordUsingTokenRequest schema from the OpenAPI specification
type ResetPasswordUsingTokenRequest struct {
	New_password string `json:"new_password"` // The new password. Needs to be at least 8 characters long and contain at least one lower- and one uppercase letter.
	Token string `json:"token"` // The reset password token that was sent to the user
}

// BranchConfigurationWithId represents the BranchConfigurationWithId schema from the OpenAPI specification
type BranchConfigurationWithId struct {
	Testsenabled bool `json:"testsEnabled,omitempty"`
	Toolsets interface{} `json:"toolsets,omitempty"` // The branch build configuration for each toolset
	Trigger string `json:"trigger,omitempty"`
	Artifactversioning map[string]interface{} `json:"artifactVersioning,omitempty"` // The versioning configuration for artifacts built for this branch
	Badgeisenabled bool `json:"badgeIsEnabled,omitempty"`
	Clonefrombranch string `json:"cloneFromBranch,omitempty"` // A configured branch name to clone from. If provided, all other parameters will be ignored. Only supported in POST requests.
	Signed bool `json:"signed,omitempty"`
	Id int `json:"id"`
}

// AppResponse represents the AppResponse schema from the OpenAPI specification
type AppResponse struct {
	Release_type string `json:"release_type,omitempty"` // A one-word descriptive release-type value that starts with a capital letter but is otherwise lowercase
	Description string `json:"description,omitempty"` // The description of the app
	Display_name string `json:"display_name"` // The display name of the app
	Icon_source string `json:"icon_source,omitempty"` // The string representation of the source of the app's icon
	Id string `json:"id"` // The unique ID (UUID) of the app
	Os string `json:"os"` // The OS the app will be running on
	Owner interface{} `json:"owner"` // The information about the app's owner
	Icon_url string `json:"icon_url,omitempty"` // The string representation of the URL pointing to the app's icon
	Name string `json:"name"` // The name of the app used in URLs
	Platform string `json:"platform,omitempty"` // The platform of the app
	Updated_at string `json:"updated_at,omitempty"` // The last updated date of this app
	App_secret string `json:"app_secret,omitempty"` // A unique and secret key used to identify the app in communication with the ingestion endpoint for crash reporting and analytics
	Azure_subscription interface{} `json:"azure_subscription,omitempty"`
	Created_at string `json:"created_at,omitempty"` // The created date of this app
	Member_permissions []string `json:"member_permissions,omitempty"` // The permissions of the calling user
	Origin string `json:"origin,omitempty"` // The creation origin of this app
}

// CrashDeleteCounter represents the CrashDeleteCounter schema from the OpenAPI specification
type CrashDeleteCounter struct {
	Crash_group_id string `json:"crash_group_id,omitempty"`
	Crash_id string `json:"crash_id,omitempty"`
	Crashes_deleted int `json:"crashes_deleted,omitempty"`
	App_id string `json:"app_id,omitempty"`
	Attachments_deleted int `json:"attachments_deleted,omitempty"`
	Blobs_failed int `json:"blobs_failed,omitempty"`
	Blobs_succeeded int `json:"blobs_succeeded,omitempty"`
}

// XcodeToolset represents the XcodeToolset schema from the OpenAPI specification
type XcodeToolset struct {
	Xcodeschemecontainers []interface{} `json:"xcodeSchemeContainers"` // The Xcode scheme containers
}

// JavaScriptSolution represents the JavaScriptSolution schema from the OpenAPI specification
type JavaScriptSolution struct {
	Packagejsonpath string `json:"packageJsonPath"` // The path to the detected package.json
	Reactnativeversion string `json:"reactNativeVersion,omitempty"` // Version of React Native from package.json files
}

// ExportBlobConfiguration represents the ExportBlobConfiguration schema from the OpenAPI specification
type ExportBlobConfiguration struct {
	TypeField string `json:"type"` // Type of export configuration
	Backfill bool `json:"backfill,omitempty"` // Field to determine if backfilling should occur. The default value is true. If set to false export starts from date and time of config creation.
	Export_entities []string `json:"export_entities,omitempty"`
	Resource_group string `json:"resource_group,omitempty"` // The resource group name on azure
	Resource_name string `json:"resource_name,omitempty"` // The resource name on azure
}

// AlertingVstsBugtrackerSettings represents the AlertingVstsBugtrackerSettings schema from the OpenAPI specification
type AlertingVstsBugtrackerSettings struct {
	Owner_name string `json:"owner_name"`
	TypeField string `json:"type"` // type of bugtracker
	Callback_url string `json:"callback_url,omitempty"`
}

// PerformanceReport represents the PerformanceReport schema from the OpenAPI specification
type PerformanceReport struct {
	Video interface{} `json:"video,omitempty"`
	Device_snapshot_id string `json:"device_snapshot_id,omitempty"`
	Performance_data map[string]interface{} `json:"performance_data,omitempty"`
}

// BlobInfo represents the BlobInfo schema from the OpenAPI specification
type BlobInfo struct {
	Url string `json:"url"`
	Size float64 `json:"size"`
}

// AADTenantAddRequest represents the AADTenantAddRequest schema from the OpenAPI specification
type AADTenantAddRequest struct {
	Aad_tenant_id string `json:"aad_tenant_id"` // The AAD tenant id
	Display_name string `json:"display_name"` // The name of the AAD Tenant
	User_id string `json:"user_id"` // The user wanting to add this tenant to the organization, must be an admin of the organization
}

// CrashGroupPlace represents the CrashGroupPlace schema from the OpenAPI specification
type CrashGroupPlace struct {
	Place_name string `json:"place_name,omitempty"` // Place name.
	Crash_count int64 `json:"crash_count,omitempty"` // Count of places.
}

// DeviceFrame represents the DeviceFrame schema from the OpenAPI specification
type DeviceFrame struct {
	Full interface{} `json:"full,omitempty"`
	Grid interface{} `json:"grid,omitempty"`
}

// V2StatusResponse represents the V2StatusResponse schema from the OpenAPI specification
type V2StatusResponse struct {
	Status string `json:"status"`
}

// ItunesTeamsRequest represents the ItunesTeamsRequest schema from the OpenAPI specification
type ItunesTeamsRequest struct {
	Cookie string `json:"cookie,omitempty"` // The 30-day session cookie for multi-factor authentication backed accounts.
	Password string `json:"password,omitempty"` // The password for the Apple Developer account.
	Service_connection_id string `json:"service_connection_id,omitempty"` // The service_connection_id of the stored Apple credentials instead of username, password.
	Username string `json:"username,omitempty"` // The username for the Apple Developer account.
}

// ArchIdentifier represents the ArchIdentifier schema from the OpenAPI specification
type ArchIdentifier struct {
	Uuid string `json:"uuid"` // The unique identifier.
	Architecture string `json:"architecture"` // The architecture that the UUID belongs to, i.e. armv7 or arm64.
}

// UserProfileResponse represents the UserProfileResponse schema from the OpenAPI specification
type UserProfileResponse struct {
	Avatar_url string `json:"avatar_url,omitempty"` // The avatar URL of the user
	Can_change_password bool `json:"can_change_password,omitempty"` // User is required to send an old password in order to change the password.
	Display_name string `json:"display_name"` // The full name of the user. Might for example be first and last name
	Email string `json:"email"` // The email address of the user
	Id string `json:"id"` // The unique id (UUID) of the user
	Name string `json:"name"` // The unique name that is used to identify the user.
	Origin string `json:"origin"` // The creation origin of this user
	Permissions []string `json:"permissions,omitempty"` // The permissions the user has for the app
}

// CodePushReleaseInfo represents the CodePushReleaseInfo schema from the OpenAPI specification
type CodePushReleaseInfo struct {
	Target_binary_range string `json:"target_binary_range,omitempty"`
	Description string `json:"description,omitempty"`
	Is_disabled bool `json:"is_disabled,omitempty"`
	Is_mandatory bool `json:"is_mandatory,omitempty"`
	Rollout int `json:"rollout,omitempty"`
}

// CodePushReleaseMetric represents the CodePushReleaseMetric schema from the OpenAPI specification
type CodePushReleaseMetric struct {
	Installed int `json:"installed,omitempty"`
	Label string `json:"label"`
	Active int `json:"active"`
	Downloaded int `json:"downloaded,omitempty"`
	Failed int `json:"failed,omitempty"`
}

// LegacyUpdateInfoData represents the LegacyUpdateInfoData schema from the OpenAPI specification
type LegacyUpdateInfoData struct {
	Rollout int `json:"rollout,omitempty"`
	Appversion string `json:"appVersion,omitempty"`
	Description string `json:"description,omitempty"`
	Isdisabled bool `json:"isDisabled,omitempty"`
	Ismandatory bool `json:"isMandatory,omitempty"`
	Updateappversion bool `json:"updateAppVersion,omitempty"`
	Downloadurl string `json:"downloadURL,omitempty"`
	Isavailable bool `json:"isAvailable,omitempty"`
	Label string `json:"label,omitempty"`
	Packagehash string `json:"packageHash,omitempty"`
	Packagesize float64 `json:"packageSize,omitempty"`
	Shouldrunbinaryversion bool `json:"shouldRunBinaryVersion,omitempty"`
}

// DateTimeCounts represents the DateTimeCounts schema from the OpenAPI specification
type DateTimeCounts struct {
	Count int64 `json:"count,omitempty"` // Count of the object.
	Datetime string `json:"datetime,omitempty"` // The ISO 8601 datetime.
}

// HockeyAppCompatibilityReleaseResponse represents the HockeyAppCompatibilityReleaseResponse schema from the OpenAPI specification
type HockeyAppCompatibilityReleaseResponse struct {
	Install_url string `json:"install_url,omitempty"`
	Notes string `json:"notes,omitempty"`
	Device_family string `json:"device_family,omitempty"`
	Mandatory bool `json:"mandatory,omitempty"`
	Version string `json:"version,omitempty"`
	Bundle_identifier string `json:"bundle_identifier,omitempty"`
	Uploaded_at string `json:"uploaded_at,omitempty"`
	Appsize int `json:"appsize,omitempty"`
	External bool `json:"external,omitempty"`
	Minimum_os_version string `json:"minimum_os_version,omitempty"`
	Shortversion string `json:"shortversion,omitempty"`
	Id int `json:"id,omitempty"`
}

// AddOrganizationAdminRequest represents the AddOrganizationAdminRequest schema from the OpenAPI specification
type AddOrganizationAdminRequest struct {
	Assigning_reason string `json:"assigning_reason"` // The explanation for adding new org admin.
	Issue_id string `json:"issue_id"` // The id of the related Intercom issue.
	New_org_admin_id string `json:"new_org_admin_id"` // The internal unique id (UUID) of the account.
	Responsible_admin_id string `json:"responsible_admin_id"` // The id of the user who started adding new org admin.
}

// PrivateAppleSecretResponse represents the PrivateAppleSecretResponse schema from the OpenAPI specification
type PrivateAppleSecretResponse struct {
	Displayname string `json:"displayName,omitempty"` // display name of shared connection
	Id string `json:"id"` // id of the shared connection
	Is2fa bool `json:"is2FA,omitempty"` // if the account is a 2FA account or not
	Isvalid bool `json:"isValid,omitempty"` // whether the credentials are valid or not
	Servicetype string `json:"serviceType"` // service type of shared connection can be apple|gitlab|googleplay|jira|applecertificate
}

// Subscription represents the Subscription schema from the OpenAPI specification
type Subscription struct {
	Active bool `json:"active,omitempty"` // Is the subscription currently active?
	Concurrentdeviceslimit int `json:"concurrentDevicesLimit,omitempty"` // Customer limits on concurrent devices
	Daysleft float64 `json:"daysLeft,omitempty"` // The number of days left in the subscription
	Endsat string `json:"endsAt,omitempty"` // The date the subscription will end or ended
	Id string `json:"id,omitempty"` // Id of the subscription
	Runningdevices int `json:"runningDevices,omitempty"` // Current device concurrency utilization
	Startsat string `json:"startsAt,omitempty"` // The date the subscription began
	Tier map[string]interface{} `json:"tier,omitempty"`
}

// SharedConnectionPatchRequest represents the SharedConnectionPatchRequest schema from the OpenAPI specification
type SharedConnectionPatchRequest struct {
	Data map[string]interface{} `json:"data"` // Represents the data for connecting to service
	Displayname string `json:"displayName,omitempty"` // Display name of the shared connection
}

// LogFlowBooleanProperty represents the LogFlowBooleanProperty schema from the OpenAPI specification
type LogFlowBooleanProperty struct {
	TypeField string `json:"type"`
	Name string `json:"name"`
}

// GitHubMarketplacePurchase represents the GitHubMarketplacePurchase schema from the OpenAPI specification
type GitHubMarketplacePurchase struct {
	Account map[string]interface{} `json:"account,omitempty"` // GitHub account information
	Plan map[string]interface{} `json:"plan,omitempty"` // GitHub Marketplace plan
}

// DeviceAvailability represents the DeviceAvailability schema from the OpenAPI specification
type DeviceAvailability struct {
	Available float64 `json:"available"`
	Maximum float64 `json:"maximum"`
	Registered float64 `json:"registered"`
}

// OrgUserPermissionResponse represents the OrgUserPermissionResponse schema from the OpenAPI specification
type OrgUserPermissionResponse struct {
	Orgid string `json:"orgId"` // The unique id (UUID) of the org
	Userrole string `json:"userRole"` // The user role for the org
}

// TeamResponse represents the TeamResponse schema from the OpenAPI specification
type TeamResponse struct {
	Description string `json:"description,omitempty"` // The description of the team
	Display_name string `json:"display_name"` // The display name of the team
	Id string `json:"id"` // The internal unique id (UUID) of the team.
	Name string `json:"name"` // The name of the team
}

// DeviceConfiguration represents the DeviceConfiguration schema from the OpenAPI specification
type DeviceConfiguration struct {
	Id string `json:"id,omitempty"` // The unique id of the device configuration
	Image interface{} `json:"image,omitempty"`
	Marketshare float64 `json:"marketShare,omitempty"`
	Model interface{} `json:"model,omitempty"`
	Name string `json:"name,omitempty"` // The name of the device model and OS version
	Os string `json:"os,omitempty"`
	Osname string `json:"osName,omitempty"`
	Tier float64 `json:"tier,omitempty"` // The tier
}

// DeviceResolution represents the DeviceResolution schema from the OpenAPI specification
type DeviceResolution struct {
	Height string `json:"height,omitempty"`
	Ppi string `json:"ppi,omitempty"`
	Width string `json:"width,omitempty"`
}

// DeviceMemory represents the DeviceMemory schema from the OpenAPI specification
type DeviceMemory struct {
	Formattedsize string `json:"formattedSize,omitempty"`
}

// DiagnosticIdResponse represents the DiagnosticIdResponse schema from the OpenAPI specification
type DiagnosticIdResponse struct {
	Diagnostic_id string `json:"diagnostic_id,omitempty"` // diagnostic id
}

// ReleaseCreateRequest represents the ReleaseCreateRequest schema from the OpenAPI specification
type ReleaseCreateRequest struct {
	Ipa_uuids []map[string]interface{} `json:"ipa_uuids,omitempty"` // A list of UUIDs for architectures for an iOS app.
	Proxy_flow bool `json:"proxy_flow,omitempty"` // If true this release was uploaded to the AKS upload proxy
	Fingerprint string `json:"fingerprint"` // MD5 checksum of the release binary.
	Icon_asset_id string `json:"icon_asset_id,omitempty"` // The assetId associated with the icon uploaded to app center file upload service.
	Languages []string `json:"languages,omitempty"` // The languages supported by the release. Limited to 510 characters in a serialized array.
	Package_url string `json:"package_url,omitempty"` // The URL to the release's binary.
	Size int `json:"size"` // The release's size in bytes.
	Unique_identifier string `json:"unique_identifier,omitempty"` // The identifier of the app's bundle.
	Appex_provisioning_profiles []map[string]interface{} `json:"appex_provisioning_profiles,omitempty"` // iOS app extension provisioning profiles included in the release.
	Provision map[string]interface{} `json:"provision,omitempty"` // An object containing information about an iOS provisioning profile.
	Device_family string `json:"device_family,omitempty"` // The release's device family.
	Version string `json:"version,omitempty"` // The release's version.<br> For iOS: CFBundleVersion from info.plist.<br> For Android: android:versionCode from AppManifest.xml.
	File_extension string `json:"file_extension,omitempty"` // The file extension of the asset. Does not include the initial period.
	Minimum_os_version string `json:"minimum_os_version,omitempty"` // The release's minimum required operating system.
	Upload_id string `json:"upload_id"` // The upload id associated with the release, to map to the releases upload table.
	Build_version string `json:"build_version,omitempty"` // The release's short version.<br> For iOS: CFBundleShortVersionString from info.plist.<br> For Android: android:versionName from AppManifest.xml.
}

// DistributionGroupDetailsResponse represents the DistributionGroupDetailsResponse schema from the OpenAPI specification
type DistributionGroupDetailsResponse struct {
	Origin string `json:"origin"` // The creation origin of this distribution group
	Display_name string `json:"display_name,omitempty"` // The name of the distribution group
	Id string `json:"id"` // The unique ID of the distribution group
	Is_public bool `json:"is_public"` // Whether the distribution group is public
	Name string `json:"name"` // The name of the distribution group used in URLs
	Is_shared bool `json:"is_shared,omitempty"` // Whether the distribution group is shared group or not
	Notified_user_count float64 `json:"notified_user_count,omitempty"` // The count of non-pending users in the distribution group who will be notified by new releases
	Total_apps_count float64 `json:"total_apps_count,omitempty"` // The count of apps associated with this distribution group
	Total_user_count float64 `json:"total_user_count,omitempty"` // The count of users in the distribution group
	Users []interface{} `json:"users,omitempty"` // The distribution group users
	Group_type string `json:"group_type,omitempty"` // Type of group (Default, HockeyAppDefault or MicrosoftDogfooding)
}

// DeleteReleasesContainer represents the DeleteReleasesContainer schema from the OpenAPI specification
type DeleteReleasesContainer struct {
	Releases []map[string]interface{} `json:"releases"`
}

// BasicAppResponse represents the BasicAppResponse schema from the OpenAPI specification
type BasicAppResponse struct {
	Id string `json:"id"` // The unique ID (UUID) of the app
	Name string `json:"name"` // The name of the app used in URLs
	Release_type string `json:"release_type,omitempty"` // A one-word descriptive release-type value that starts with a capital letter but is otherwise lowercase
	Os string `json:"os"` // The OS the app will be running on
	Owner interface{} `json:"owner"` // The information about the app's owner
	Icon_url string `json:"icon_url,omitempty"` // The string representation of the URL pointing to the app's icon
	Icon_source string `json:"icon_source,omitempty"` // The string representation of the source of the app's icon
	Description string `json:"description,omitempty"` // The description of the app
	Display_name string `json:"display_name"` // The display name of the app
}

// AppleLoginResponse represents the AppleLoginResponse schema from the OpenAPI specification
type AppleLoginResponse struct {
	Successful bool `json:"successful,omitempty"` // True when login was successful.
}

// ApiTokensCreateRequest represents the ApiTokensCreateRequest schema from the OpenAPI specification
type ApiTokensCreateRequest struct {
	Description string `json:"description,omitempty"` // The description of the token
	Scope []string `json:"scope,omitempty"` // The scope for this token.
}

// BuildAgentQueueResponse represents the BuildAgentQueueResponse schema from the OpenAPI specification
type BuildAgentQueueResponse struct {
	Builddefinition string `json:"buildDefinition,omitempty"` // Name of the build definition
	Name string `json:"name,omitempty"` // Name of the queue
}

// UserEmailRequest represents the UserEmailRequest schema from the OpenAPI specification
type UserEmailRequest struct {
	User_email string `json:"user_email"` // The user's email address
}

// DistributionGroupPrivateResponse represents the DistributionGroupPrivateResponse schema from the OpenAPI specification
type DistributionGroupPrivateResponse struct {
	Id string `json:"id"` // The unique ID of the distribution group
	Is_public bool `json:"is_public"` // Whether the distribution group is public
	Name string `json:"name"` // The name of the distribution group used in URLs
	Origin string `json:"origin"` // The creation origin of this distribution group
	Display_name string `json:"display_name,omitempty"` // The name of the distribution group
	Group_type string `json:"group_type,omitempty"` // Type of group
}

// PatchReleaseUploadStatusResponse represents the PatchReleaseUploadStatusResponse schema from the OpenAPI specification
type PatchReleaseUploadStatusResponse struct {
	Upload_status string `json:"upload_status"` // The current upload status.
	Id string `json:"id"` // The ID for the upload.
}

// AlertBugTrackerRepo represents the AlertBugTrackerRepo schema from the OpenAPI specification
type AlertBugTrackerRepo struct {
	Private bool `json:"private,omitempty"`
	Url string `json:"url"`
	Description string `json:"description,omitempty"`
	Id string `json:"id"`
	Name string `json:"name"`
	Owner map[string]interface{} `json:"owner,omitempty"` // Repository owner object
}

// ReleaseDestinationRequest represents the ReleaseDestinationRequest schema from the OpenAPI specification
type ReleaseDestinationRequest struct {
	Id string `json:"id"` // Unique id of the release destination
	Mandatory_update bool `json:"mandatory_update,omitempty"` // Flag to mark the release for the provided destinations as mandatory
	Notify_testers bool `json:"notify_testers,omitempty"` // Flag to enable or disable notifications to testers
}

// AppleCertificateSecretDetailsResponse represents the AppleCertificateSecretDetailsResponse schema from the OpenAPI specification
type AppleCertificateSecretDetailsResponse struct {
	Servicetype string `json:"serviceType"` // service type of shared connection can be apple|gitlab|googleplay|jira
	Credentialtype string `json:"credentialType,omitempty"` // the type of the credential
	Displayname string `json:"displayName"` // display name of shared connection
}

// GetInAppUpdateTokenResponse represents the GetInAppUpdateTokenResponse schema from the OpenAPI specification
type GetInAppUpdateTokenResponse struct {
	Token string `json:"token"` // The api token generated will not be accessible again
}

// ErrorLogDiagnostics represents the ErrorLogDiagnostics schema from the OpenAPI specification
type ErrorLogDiagnostics struct {
	Timestamp string `json:"timestamp"` // Log creation timestamp.
	TypeField string `json:"type"` // Log type.
	Device map[string]interface{} `json:"device"` // Device characteristics.
	Install_id string `json:"install_id"` // Install ID.
}

// ReleaseDetailsUpdateResponse represents the ReleaseDetailsUpdateResponse schema from the OpenAPI specification
type ReleaseDetailsUpdateResponse struct {
	Release_notes string `json:"release_notes,omitempty"`
}

// BillingResourceUsage represents the BillingResourceUsage schema from the OpenAPI specification
type BillingResourceUsage struct {
	Buildservice map[string]interface{} `json:"buildService,omitempty"` // Resource usage for a single Mobile Center service
	Testservice map[string]interface{} `json:"testService,omitempty"` // Resource usage for a single Mobile Center service
}

// CrashingAppDetail represents the CrashingAppDetail schema from the OpenAPI specification
type CrashingAppDetail struct {
	Appid string `json:"appId,omitempty"` // application identifier
	Appversion string `json:"appVersion,omitempty"` // application version
	Crashgroupid string `json:"crashGroupId,omitempty"` // crash group identifier
}

// CreateStoreSecretRequest represents the CreateStoreSecretRequest schema from the OpenAPI specification
type CreateStoreSecretRequest struct {
	Secret_json interface{} `json:"secret_json,omitempty"`
	Tenant_id string `json:"tenant_id,omitempty"` // the tenant id for user
}

// NewAppReleaseAlertingEvent represents the NewAppReleaseAlertingEvent schema from the OpenAPI specification
type NewAppReleaseAlertingEvent struct {
	Event_id string `json:"event_id"` // A unique identifier for this event instance. Useful for deduplication
	Event_timestamp string `json:"event_timestamp"` // ISO 8601 date time when event was generated
	Properties map[string]interface{} `json:"properties,omitempty"` // Obsolete. Use emailProperties.
	App_release_properties map[string]interface{} `json:"app_release_properties,omitempty"` // Properties of new application release
	Disable_webhook bool `json:"disable_webhook,omitempty"` // indicate whether notify via webhook or not
	User_ids []string `json:"user_ids,omitempty"` // List of users who need to receive an email notification. If this is not null, then only sending emails will be triggered even if the event requires calling webhooks or doing other actions.
}

// PublishDevicesResponse represents the PublishDevicesResponse schema from the OpenAPI specification
type PublishDevicesResponse struct {
	Profile_file_name string `json:"profile_file_name,omitempty"` // The file name for the provisioning profile.
	Profiles_zip_base64 string `json:"profiles_zip_base64"` // The updated provisioning profiles zip base64 encoded.
}

// DistributionGroupRequest represents the DistributionGroupRequest schema from the OpenAPI specification
type DistributionGroupRequest struct {
	Display_name string `json:"display_name,omitempty"` // The display name of the distribution group. If not specified, the name will be used.
	Name string `json:"name"` // The name of the distribution group
}

// ApiTokenDeleteResponse represents the ApiTokenDeleteResponse schema from the OpenAPI specification
type ApiTokenDeleteResponse struct {
	Id string `json:"id"` // The unique id (UUID) of the api token
	Token_hash string `json:"token_hash"` // The hashed value of api token
}

// BugTrackerIssueResult represents the BugTrackerIssueResult schema from the OpenAPI specification
type BugTrackerIssueResult struct {
	Url string `json:"url,omitempty"`
	Bug_tracker_type string `json:"bug_tracker_type,omitempty"`
	Event_type string `json:"event_type,omitempty"`
	Id string `json:"id,omitempty"`
	Mobile_center_id string `json:"mobile_center_id,omitempty"`
	Repo_name string `json:"repo_name,omitempty"`
	Title string `json:"title,omitempty"`
}

// LogFlowCustomPropertyLog represents the LogFlowCustomPropertyLog schema from the OpenAPI specification
type LogFlowCustomPropertyLog struct {
	Device map[string]interface{} `json:"device"` // Device characteristics.
	Install_id string `json:"install_id"` // Install ID.
	Timestamp string `json:"timestamp"` // Log creation timestamp.
	TypeField string `json:"type"` // Log type.
}

// OrgNameAvailabilityResponse represents the OrgNameAvailabilityResponse schema from the OpenAPI specification
type OrgNameAvailabilityResponse struct {
	Name string `json:"name"` // The generated org name
	Available bool `json:"available"` // The availability status of the requested org name
}

// SystemVersionNameGroup represents the SystemVersionNameGroup schema from the OpenAPI specification
type SystemVersionNameGroup struct {
	Name string `json:"name,omitempty"` // Name of version group
	Versions []string `json:"versions,omitempty"`
}

// LogFlowErrorLog represents the LogFlowErrorLog schema from the OpenAPI specification
type LogFlowErrorLog struct {
	Device map[string]interface{} `json:"device"` // Device characteristics.
	Install_id string `json:"install_id"` // Install ID.
	Timestamp string `json:"timestamp"` // Log creation timestamp.
	TypeField string `json:"type"` // Log type.
}

// ReleaseUpdateResponse represents the ReleaseUpdateResponse schema from the OpenAPI specification
type ReleaseUpdateResponse struct {
	Release_notes string `json:"release_notes,omitempty"`
	Destinations []map[string]interface{} `json:"destinations,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Mandatory_update bool `json:"mandatory_update,omitempty"`
	Provisioning_status_url string `json:"provisioning_status_url,omitempty"`
}

// TransferRepoOwnerRequest represents the TransferRepoOwnerRequest schema from the OpenAPI specification
type TransferRepoOwnerRequest struct {
	Repository_url string `json:"repository_url"` // The url of repository to have its ownership transferred.
	Responsible_admin_id string `json:"responsible_admin_id"` // The id of the user who started transfer process.
	Transfer_reason string `json:"transfer_reason"` // The explanation for starting transfer process.
	Issue_id string `json:"issue_id"` // The id of the related Intercom issue.
	New_owner_id string `json:"new_owner_id"` // The internal unique id (UUID) of the user.
}

// AccountResponse represents the AccountResponse schema from the OpenAPI specification
type AccountResponse struct {
	Name string `json:"name"` // The slug name of the account
	Origin string `json:"origin"` // The creation origin of this account
	TypeField string `json:"type"` // The type of this account
	Display_name string `json:"display_name"` // The display name of the account
	Email string `json:"email,omitempty"` // The account's email. For org that value might be empty.
	Id string `json:"id"` // The internal unique id (UUID) of the account.
}

// CrashGroupModel represents the CrashGroupModel schema from the OpenAPI specification
type CrashGroupModel struct {
	Crash_count int64 `json:"crash_count,omitempty"` // Count of model.
	Model_name string `json:"model_name,omitempty"` // Model's name.
}

// EventLog represents the EventLog schema from the OpenAPI specification
type EventLog struct {
	TypeField string `json:"type"` // Log type.
	Device map[string]interface{} `json:"device"` // Device characteristics.
	Install_id string `json:"install_id"` // Install ID.
	Timestamp string `json:"timestamp"` // Log creation timestamp.
	Properties map[string]interface{} `json:"properties,omitempty"` // Additional key/value pair parameters.
}

// ServiceBusStatusResponse represents the ServiceBusStatusResponse schema from the OpenAPI specification
type ServiceBusStatusResponse struct {
	Subscriptions []interface{} `json:"subscriptions,omitempty"`
	Status string `json:"status"`
}

// ExternalUserRequest represents the ExternalUserRequest schema from the OpenAPI specification
type ExternalUserRequest struct {
	Avatar_url string `json:"avatar_url,omitempty"` // The avatar URL of the user
	Display_name string `json:"display_name,omitempty"` // The full name of the user. Might for example be first and last name
	Email string `json:"email"` // The email address of the user
	Name string `json:"name,omitempty"` // The unique name that is used to identify the user. If no explicit value is given, a default will be auto-generated from the `display_name` value
	Organization_invitation string `json:"organization_invitation,omitempty"` // The token of the organization invitation which lead to signup
	Tester_invitation string `json:"tester_invitation,omitempty"` // The token of the test invitation which lead to signup
	App_invitation string `json:"app_invitation,omitempty"` // The token of the app invitation which lead to signup
}

// UpdateResignStatusResponse represents the UpdateResignStatusResponse schema from the OpenAPI specification
type UpdateResignStatusResponse struct {
	Profiles_zip_base64 string `json:"profiles_zip_base64,omitempty"` // A zip of the updated provisioning profiles. Base64 encoded.
	Status string `json:"status"` // The status.
}

// DeviceInfoRequest represents the DeviceInfoRequest schema from the OpenAPI specification
type DeviceInfoRequest struct {
	Os_build string `json:"os_build,omitempty"` // The build number of the last known OS version running on the device
	Os_version string `json:"os_version,omitempty"` // The last known OS version running on the device
	Owner_id string `json:"owner_id,omitempty"` // The user ID of the device owner.
	Serial string `json:"serial,omitempty"` // The device's serial number. Always empty or undefined at present.
	Udid string `json:"udid"` // The Unique Device IDentifier of the device
	Imei string `json:"imei,omitempty"` // The device's International Mobile Equipment Identity number. Always empty or undefined at present.
	Model string `json:"model"` // The model identifier of the device, in the format iDeviceM,N
}

// PrivateBasicReleaseDetailsResponse represents the PrivateBasicReleaseDetailsResponse schema from the OpenAPI specification
type PrivateBasicReleaseDetailsResponse struct {
	Is_external_build bool `json:"is_external_build,omitempty"` // This value determines if a release is external or not.
	Publishing_status string `json:"publishing_status,omitempty"` // the publishing status of the distributed release
	Short_version string `json:"short_version,omitempty"` // The release's short version.<br> For iOS: CFBundleShortVersionString from info.plist.<br> For Android: android:versionName from AppManifest.xml.
	Id int `json:"id,omitempty"` // ID identifying this unique release.
	Mandatory_update bool `json:"mandatory_update,omitempty"` // A boolean which determines whether the release is a mandatory update or not.
	Origin string `json:"origin,omitempty"` // The release's origin
	Distribution_group_id string `json:"distribution_group_id,omitempty"` // the destination id of release where it is distributed.
	Is_latest bool `json:"is_latest,omitempty"` // Indicates if this is the latest release in the group.
	Uploaded_at string `json:"uploaded_at,omitempty"` // UTC time in ISO 8601 format of the uploaded time.
	Version string `json:"version,omitempty"` // The release's version.<br> For iOS: CFBundleVersion from info.plist.<br> For Android: android:versionCode from AppManifest.xml.
	Destination_type string `json:"destination_type,omitempty"` // The destination type.<br> <b>group</b>: The release distributed to internal groups and distribution_groups details will be returned.<br> <b>store</b>: The release distributed to external stores and distribution_stores details will be returned. <br>
}

// AppleCredentialsSecretRequest represents the AppleCredentialsSecretRequest schema from the OpenAPI specification
type AppleCredentialsSecretRequest struct {
	Displayname string `json:"displayName,omitempty"` // display name of shared connection
	Servicetype string `json:"serviceType"` // service type of shared connection can be apple|gitlab|googleplay|jira|applecertificate
	Credentialtype string `json:"credentialType,omitempty"` // credential type of the shared connection. Values can be credentials|certificate
	Data map[string]interface{} `json:"data,omitempty"` // shared connection details
}

// UpdateCheckResponse represents the UpdateCheckResponse schema from the OpenAPI specification
type UpdateCheckResponse struct {
	Update_info interface{} `json:"update_info"`
}

// AppName represents the AppName schema from the OpenAPI specification
type AppName struct {
	Name string `json:"name"`
}

// VersionDiagnostics represents the VersionDiagnostics schema from the OpenAPI specification
type VersionDiagnostics struct {
	Count int64 `json:"count,omitempty"` // version count
	Previous_count int64 `json:"previous_count,omitempty"` // the count of previous time range of the version
	Version string `json:"version,omitempty"` // version
}

// AppCrashesInfo represents the AppCrashesInfo schema from the OpenAPI specification
type AppCrashesInfo struct {
	Features map[string]interface{} `json:"features"`
	Has_crashes bool `json:"has_crashes"`
}

// Model represents the Model schema from the OpenAPI specification
type Model struct {
	Count int64 `json:"count,omitempty"` // Count current of model.
	Model_name string `json:"model_name,omitempty"` // Model's name.
	Previous_count int64 `json:"previous_count,omitempty"` // Count of previous model.
}

// DeviceSetUpdate represents the DeviceSetUpdate schema from the OpenAPI specification
type DeviceSetUpdate struct {
	Devices []string `json:"devices"` // List of device IDs
	Name string `json:"name"` // The name of the device set
}

// GitHubAccountLite represents the GitHubAccountLite schema from the OpenAPI specification
type GitHubAccountLite struct {
	Id string `json:"id,omitempty"` // GitHub Account Id
	Login string `json:"login,omitempty"` // GitHub Account Login Name
	TypeField string `json:"type,omitempty"` // GitHub Account Type
	Url string `json:"url,omitempty"` // GitHub Account Url
}

// LogWithProperties represents the LogWithProperties schema from the OpenAPI specification
type LogWithProperties struct {
	TypeField string `json:"type"` // Log type.
	Device map[string]interface{} `json:"device"` // Device characteristics.
	Install_id string `json:"install_id"` // Install ID.
	Timestamp string `json:"timestamp"` // Log creation timestamp.
}

// PasswordUpdateRequest represents the PasswordUpdateRequest schema from the OpenAPI specification
type PasswordUpdateRequest struct {
	Old_password string `json:"old_password,omitempty"` // The old password, if needed.
	New_password string `json:"new_password"` // The new password that will be set for the user. Needs to be at least 8 characters long and contain at least one lower- and one uppercase letter.
}

// LogFlowNumberProperty represents the LogFlowNumberProperty schema from the OpenAPI specification
type LogFlowNumberProperty struct {
	Name string `json:"name"`
	TypeField string `json:"type"`
}

// WebSocketContainer represents the WebSocketContainer schema from the OpenAPI specification
type WebSocketContainer struct {
	Url string `json:"url"` // WebSocket URL
}

// AlertingEvent represents the AlertingEvent schema from the OpenAPI specification
type AlertingEvent struct {
	Event_timestamp string `json:"event_timestamp"` // ISO 8601 date time when event was generated
	Properties map[string]interface{} `json:"properties,omitempty"` // Obsolete. Use emailProperties.
	Event_id string `json:"event_id"` // A unique identifier for this event instance. Useful for deduplication
}

// DistributionGroupAADGroupRequest represents the DistributionGroupAADGroupRequest schema from the OpenAPI specification
type DistributionGroupAADGroupRequest struct {
	Aad_groups []interface{} `json:"aad_groups,omitempty"` // The list of aad group ids and names to add
}

// OrganizationPatchRequest represents the OrganizationPatchRequest schema from the OpenAPI specification
type OrganizationPatchRequest struct {
	Display_name string `json:"display_name,omitempty"` // The full (friendly) name of the organization.
	Name string `json:"name,omitempty"` // The name of the organization used in URLs
}

// TestRunStatistics represents the TestRunStatistics schema from the OpenAPI specification
type TestRunStatistics struct {
	Devicesfinished float64 `json:"devicesFinished,omitempty"` // Number of finished devices
	Skipped float64 `json:"skipped,omitempty"` // Number of skipped tests
	Total float64 `json:"total,omitempty"` // Number of tests in total
	Devicesfailed float64 `json:"devicesFailed,omitempty"` // Number of failed devices
	Passed float64 `json:"passed,omitempty"` // Number of passed tests
	Devices float64 `json:"devices,omitempty"` // Number of devices running the test
	Failed float64 `json:"failed,omitempty"` // Number of failed tests
	Peakmemory float64 `json:"peakMemory,omitempty"` // The max amount of MB used during the test run
	Totaldeviceminutes float64 `json:"totalDeviceMinutes,omitempty"` // The number of minutes of device time the test has been runnign
}

// DistributionGroupAADGroupResponse represents the DistributionGroupAADGroupResponse schema from the OpenAPI specification
type DistributionGroupAADGroupResponse struct {
	Id string `json:"id,omitempty"` // The internal unique id (UUID) of the AAD group.
	Is_aad_group bool `json:"is_aad_group,omitempty"`
	Aad_group_id string `json:"aad_group_id,omitempty"` // The AAD unique id (UUID) of the AAD group.
	Display_name string `json:"display_name,omitempty"` // The display name of the AAD group
	Distribution_group_name string `json:"distribution_group_name,omitempty"` // The distribution group of the AAD group
}

// MetricsValuesByApp represents the MetricsValuesByApp schema from the OpenAPI specification
type MetricsValuesByApp struct {
}

// ErrorDateTimeCounts represents the ErrorDateTimeCounts schema from the OpenAPI specification
type ErrorDateTimeCounts struct {
	Datetime string `json:"datetime,omitempty"` // the ISO 8601 datetime
	Count int64 `json:"count,omitempty"` // count of the object
}

// ErrorRetentionSettings represents the ErrorRetentionSettings schema from the OpenAPI specification
type ErrorRetentionSettings struct {
	Retention_in_days int `json:"retention_in_days"`
}

// FilterReleasesContainer represents the FilterReleasesContainer schema from the OpenAPI specification
type FilterReleasesContainer struct {
	Releases []map[string]interface{} `json:"releases,omitempty"`
}

// AppleMultifactorSecretDetails represents the AppleMultifactorSecretDetails schema from the OpenAPI specification
type AppleMultifactorSecretDetails struct {
	Username string `json:"username"` // username to connect to apple store.
	Appspecificpassword string `json:"appSpecificPassword,omitempty"` // The app specific password required for app publishing for 2FA accounts
	Authcode string `json:"authCode"` // The 6 digit Apple OTP for Multifactor accounts
	Password string `json:"password"` // password to connect to apple store.
}

// SharedConnectionRequest represents the SharedConnectionRequest schema from the OpenAPI specification
type SharedConnectionRequest struct {
	Credentialtype string `json:"credentialType,omitempty"` // credential type of the shared connection. Values can be credentials|certificate
	Data map[string]interface{} `json:"data,omitempty"` // shared connection details
	Displayname string `json:"displayName,omitempty"` // display name of shared connection
	Servicetype string `json:"serviceType"` // service type of shared connection can be apple|gitlab|googleplay|jira|applecertificate
}

// DeviceSetConfiguration represents the DeviceSetConfiguration schema from the OpenAPI specification
type DeviceSetConfiguration struct {
	Id string `json:"id,omitempty"` // The unique id of the device configuration
	Image interface{} `json:"image,omitempty"`
	Model interface{} `json:"model,omitempty"`
	Os string `json:"os,omitempty"`
	Osname string `json:"osName,omitempty"`
}

// VersionedBillingPlan represents the VersionedBillingPlan schema from the OpenAPI specification
type VersionedBillingPlan struct {
	Document map[string]interface{} `json:"document,omitempty"` // Billing Plan
	Etag string `json:"etag,omitempty"` // The version of the object
}

// AppleTestFlightGroupResponse represents the AppleTestFlightGroupResponse schema from the OpenAPI specification
type AppleTestFlightGroupResponse struct {
	Active bool `json:"active,omitempty"` // true if group is in active state.
	Appadamid float64 `json:"appAdamId,omitempty"` // apple id of the group.
	Id string `json:"id,omitempty"` // id of the group.
	Isinternalgroup bool `json:"isInternalGroup,omitempty"` // true if the group is an internal group.
	Name string `json:"name,omitempty"` // name of the group.
	Providerid float64 `json:"providerId,omitempty"` // provider id of the group.
}

// IntuneGroup represents the IntuneGroup schema from the OpenAPI specification
type IntuneGroup struct {
	Displayname string `json:"displayName,omitempty"` // the display name of the group
	Id string `json:"id,omitempty"` // the id of the Group
	Securityenabled bool `json:"securityEnabled,omitempty"` // species if it is a security group
}

// UserSettingRequest represents the UserSettingRequest schema from the OpenAPI specification
type UserSettingRequest struct {
	Value string `json:"value"` // The setting value
}

// CrashGroupCarrier represents the CrashGroupCarrier schema from the OpenAPI specification
type CrashGroupCarrier struct {
	Carrier_name string `json:"carrier_name,omitempty"` // Carrier name.
	Crash_count int64 `json:"crash_count,omitempty"` // Crash count of carrier.
}

// DateTimePercentages represents the DateTimePercentages schema from the OpenAPI specification
type DateTimePercentages struct {
	Datetime string `json:"datetime,omitempty"` // The ISO 8601 datetime.
	Percentage float64 `json:"percentage,omitempty"` // Percentage of the object.
}

// SharedConnectionResponse represents the SharedConnectionResponse schema from the OpenAPI specification
type SharedConnectionResponse struct {
	Credentialtype string `json:"credentialType,omitempty"` // the type of the credential
	Displayname string `json:"displayName"` // display name of shared connection
	Servicetype string `json:"serviceType"` // service type of shared connection can be apple|gitlab|googleplay|jira
}

// Device represents the Device schema from the OpenAPI specification
type Device struct {
	Os_api_level int `json:"os_api_level,omitempty"` // API level when applicable like in Android (example: 15).
	App_namespace string `json:"app_namespace,omitempty"` // The bundle identifier, package identifier, or namespace, depending on what the individual plattforms use, .e.g com.microsoft.example.
	App_version string `json:"app_version"` // Application version name, e.g. 1.1.0
	Wrapper_runtime_version string `json:"wrapper_runtime_version,omitempty"` // Version of the wrapper technology framework (Xamarin runtime version or ReactNative or Cordova etc...). See wrapper_sdk_name to see if this version refers to Xamarin or ReactNative or other.
	Wrapper_sdk_name string `json:"wrapper_sdk_name,omitempty"` // Name of the wrapper SDK. Consists of the name of the SDK and the wrapper platform, e.g. "appcenter.xamarin", "hockeysdk.cordova".
	Live_update_deployment_key string `json:"live_update_deployment_key,omitempty"` // Identifier of environment that current application release belongs to, deployment key then maps to environment like Production, Staging.
	Os_name string `json:"os_name"` // OS name (example: iOS). The following OS names are standardized (non-exclusive): Android, iOS, macOS, tvOS, Windows.
	Time_zone_offset int `json:"time_zone_offset"` // The offset in minutes from UTC for the device time zone, including daylight savings time.
	Sdk_version string `json:"sdk_version"` // Version of the SDK in semver format, e.g. "1.2.0" or "0.12.3-alpha.1".
	Carrier_name string `json:"carrier_name,omitempty"` // Carrier name (for mobile devices).
	Live_update_release_label string `json:"live_update_release_label,omitempty"` // Label that is used to identify application code 'version' released via Live Update beacon running on device
	Oem_name string `json:"oem_name,omitempty"` // Device manufacturer (example: HTC).
	Os_build string `json:"os_build,omitempty"` // OS build code (example: LMY47X).
	Os_version string `json:"os_version"` // OS version (example: 9.3.0).
	Wrapper_sdk_version string `json:"wrapper_sdk_version,omitempty"` // Version of the wrapper SDK in semver format. When the SDK is embedding another base SDK (for example Xamarin.Android wraps Android), the Xamarin specific version is populated into this field while sdkVersion refers to the original Android SDK.
	Live_update_package_hash string `json:"live_update_package_hash,omitempty"` // Hash of all files (ReactNative or Cordova) deployed to device via LiveUpdate beacon. Helps identify the Release version on device or need to download updates in future.
	Carrier_country string `json:"carrier_country,omitempty"` // Carrier country.
	App_build string `json:"app_build"` // The app's build number, e.g. 42.
	Model string `json:"model,omitempty"` // Device model (example: iPad2,3).
	Screen_size string `json:"screen_size,omitempty"` // Screen size of the device in pixels (example: 640x480).
	Carrier_code string `json:"carrier_code,omitempty"` // Carrier country code (for mobile devices).
	Sdk_name string `json:"sdk_name"` // Name of the SDK. Consists of the name of the SDK and the platform, e.g. "appcenter.ios", "hockeysdk.android".
	Locale string `json:"locale"` // Language code (example: en_US).
}

// TestCloudProjectFrameworkProperties represents the TestCloudProjectFrameworkProperties schema from the OpenAPI specification
type TestCloudProjectFrameworkProperties struct {
	Configurations []string `json:"configurations,omitempty"`
}

// OS represents the OS schema from the OpenAPI specification
type OS struct {
	Count int64 `json:"count,omitempty"` // Count current of OS.
	Os_name string `json:"os_name,omitempty"` // OS name.
	Previous_count int64 `json:"previous_count,omitempty"` // Count of previous OS.
}

// SubmitDsrOperationResponse represents the SubmitDsrOperationResponse schema from the OpenAPI specification
type SubmitDsrOperationResponse struct {
	Status string `json:"status,omitempty"` // The status of the DSR operation
}

// AzureSubscriptionAddToAppRequest represents the AzureSubscriptionAddToAppRequest schema from the OpenAPI specification
type AzureSubscriptionAddToAppRequest struct {
	Subscription_id string `json:"subscription_id"` // The azure subscription id
}

// AudienceQueryDefinition represents the AudienceQueryDefinition schema from the OpenAPI specification
type AudienceQueryDefinition struct {
	Identifiers []string `json:"identifiers,omitempty"` // List of identifiers.
	Querytype string `json:"queryType"` // Audience Query Type
}

// DataSubjectRightCustomerIdRequest represents the DataSubjectRightCustomerIdRequest schema from the OpenAPI specification
type DataSubjectRightCustomerIdRequest struct {
	Datasubjectidentifier string `json:"dataSubjectIdentifier,omitempty"` // customer account id (b2c identifier) / customer user id (free form text) depending on the value of the fied `type`
	TypeField string `json:"type,omitempty"` // type of the customer dataSubjectIdentifier
}

// VSTSProject represents the VSTSProject schema from the OpenAPI specification
type VSTSProject struct {
	Description string `json:"description,omitempty"` // Project description
	Id string `json:"id,omitempty"` // Project id
	Name string `json:"name,omitempty"` // Project name
	State string `json:"state,omitempty"` // Project state
	Url string `json:"url,omitempty"` // Project URL
	Visibility string `json:"visibility,omitempty"` // Project visibility
}

// EventCountPerDevice represents the EventCountPerDevice schema from the OpenAPI specification
type EventCountPerDevice struct {
	Previous_avg_count_per_device float64 `json:"previous_avg_count_per_device,omitempty"`
	Avg_count_per_device float64 `json:"avg_count_per_device,omitempty"`
	Count_per_device []interface{} `json:"count_per_device,omitempty"`
}

// BuildTimelineRecord represents the BuildTimelineRecord schema from the OpenAPI specification
type BuildTimelineRecord struct {
	Result string `json:"result,omitempty"`
	Starttime string `json:"startTime,omitempty"`
	Warningcount float64 `json:"warningCount,omitempty"`
	Order float64 `json:"order,omitempty"`
	TypeField string `json:"type,omitempty"`
	Currentoperation string `json:"currentOperation,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Percentcomplete float64 `json:"percentComplete,omitempty"`
	Errorcount float64 `json:"errorCount,omitempty"`
	Issues []interface{} `json:"issues,omitempty"`
	State string `json:"state,omitempty"`
	Finishtime string `json:"finishTime,omitempty"`
}

// SymbolUpload represents the SymbolUpload schema from the OpenAPI specification
type SymbolUpload struct {
	Status string `json:"status"` // The current status for the symbol upload
	Symbol_upload_id string `json:"symbol_upload_id"` // The id for the current symbol upload
	App_id string `json:"app_id"` // The application that this symbol upload belongs to
	File_size float64 `json:"file_size,omitempty"` // The size of the file in Mebibytes. This may be 0 until the status is indexed
	Origin string `json:"origin,omitempty"` // The origin of the symbol upload
	Symbols_uploaded []interface{} `json:"symbols_uploaded,omitempty"` // The symbols found in the upload. This may be empty until the status is indexed
	File_name string `json:"file_name,omitempty"` // The file name for the symbol upload
	Symbol_type string `json:"symbol_type"` // The type of the symbol for the current symbol upload
	Timestamp string `json:"timestamp,omitempty"` // When the symbol upload was committed, or last transaction time if not committed
	User interface{} `json:"user,omitempty"` // User information of the one who intitiated the symbol upload
}

// AutoProvisioningConfigRequest represents the AutoProvisioningConfigRequest schema from the OpenAPI specification
type AutoProvisioningConfigRequest struct {
	Allow_auto_provisioning bool `json:"allow_auto_provisioning,omitempty"` // When *true* enables auto provisioning
	Apple_developer_account_key string `json:"apple_developer_account_key,omitempty"` // A key to a secret in customer-credential-store. apple_developer_account refers to the user's developer account that is used to log into https://developer.apple.com. Normally the user's email.
	Apple_distribution_certificate_key string `json:"apple_distribution_certificate_key,omitempty"` // A key to a secret in customer-credential-store. distribution_certificate refers to the customer's certificate (that holds the private key) that will be used to sign the app.
}

// GitHubRepositoryLite represents the GitHubRepositoryLite schema from the OpenAPI specification
type GitHubRepositoryLite struct {
	Id float64 `json:"id,omitempty"` // The repository id
}

// ReleaseDetailsResponse represents the ReleaseDetailsResponse schema from the OpenAPI specification
type ReleaseDetailsResponse struct {
	Is_external_build bool `json:"is_external_build,omitempty"` // This value determines if a release is external or not.
	Min_os string `json:"min_os,omitempty"` // The release's minimum required operating system.
	Bundle_identifier string `json:"bundle_identifier,omitempty"` // The identifier of the apps bundle.
	Id int `json:"id"` // ID identifying this unique release.
	Is_provisioning_profile_syncing bool `json:"is_provisioning_profile_syncing,omitempty"` // A flag that determines whether the release's provisioning profile is still extracted or not.
	App_name string `json:"app_name"` // The app's name (extracted from the uploaded release).
	Distribution_groups []interface{} `json:"distribution_groups,omitempty"` // OBSOLETE. Will be removed in next version. A list of distribution groups that are associated with this release.
	Short_version string `json:"short_version"` // The release's short version.<br> For iOS: CFBundleShortVersionString from info.plist. For Android: android:versionName from AppManifest.xml.
	App_icon_url string `json:"app_icon_url"` // A URL to the app's icon.
	Destination_type string `json:"destination_type,omitempty"` // OBSOLETE. Will be removed in next version. The destination type.<br> <b>group</b>: The release distributed to internal groups and distribution_groups details will be returned.<br> <b>store</b>: The release distributed to external stores and distribution_stores details will be returned.<br> <b>tester</b>: The release distributed testers details will be returned.<br>
	Provisioning_profile_expiry_date string `json:"provisioning_profile_expiry_date,omitempty"` // expiration date of provisioning profile in UTC format.
	Status string `json:"status,omitempty"` // Status of the release.
	Uploaded_at string `json:"uploaded_at"` // UTC time in ISO 8601 format of the uploaded time.
	Can_resign bool `json:"can_resign,omitempty"` // In calls that allow passing `udid` in the query string, this value determines if a release can be re-signed. When true, after a re-sign, the tester will be able to install the release from his registered devices. Will not be returned for non-iOS platforms.
	Device_family string `json:"device_family,omitempty"` // The release's device family.
	Is_udid_provisioned bool `json:"is_udid_provisioned,omitempty"` // In calls that allow passing `udid` in the query string, this value will hold the provisioning status of that UDID in this release. Will be ignored for non-iOS platforms.
	Distribution_stores []interface{} `json:"distribution_stores,omitempty"` // OBSOLETE. Will be removed in next version. A list of distribution stores that are associated with this release.
	Fingerprint string `json:"fingerprint,omitempty"` // MD5 checksum of the release binary.
	Secondary_download_url string `json:"secondary_download_url,omitempty"` // The URL that hosts the secondary binary for this release, such as the apk file for aab releases.
	Install_url string `json:"install_url,omitempty"` // The href required to install a release on a mobile device. On iOS devices will be prefixed with `itms-services://?action=download-manifest&url=`
	Origin string `json:"origin,omitempty"` // The release's origin
	Provisioning_profile_type string `json:"provisioning_profile_type,omitempty"` // The type of the provisioning profile for the requested app version.
	Destinations []interface{} `json:"destinations,omitempty"` // A list of distribution groups or stores.
	Provisioning_profile_name string `json:"provisioning_profile_name,omitempty"` // The release's provisioning profile name.
	Release_notes string `json:"release_notes,omitempty"` // The release's release notes.
	App_display_name string `json:"app_display_name"` // The app's display name.
	Version string `json:"version"` // The release's version.<br> For iOS: CFBundleVersion from info.plist. For Android: android:versionCode from AppManifest.xml.
	Android_min_api_level string `json:"android_min_api_level,omitempty"` // The release's minimum required Android API level.
	App_os string `json:"app_os,omitempty"` // The app's OS.
	Download_url string `json:"download_url,omitempty"` // The URL that hosts the binary for this release.
	Enabled bool `json:"enabled"` // This value determines the whether a release currently is enabled or disabled.
	Build map[string]interface{} `json:"build,omitempty"` // Contains metadata about the build that produced the release being uploaded
	Package_hashes []string `json:"package_hashes,omitempty"` // Hashes for the packages.
	Size int `json:"size,omitempty"` // The release's size in bytes.
}

// DateTimePropertyDiagnostics represents the DateTimePropertyDiagnostics schema from the OpenAPI specification
type DateTimePropertyDiagnostics struct {
	Name string `json:"name"`
	TypeField string `json:"type"`
}

// EventCountPerSession represents the EventCountPerSession schema from the OpenAPI specification
type EventCountPerSession struct {
	Avg_count_per_session float64 `json:"avg_count_per_session,omitempty"`
	Count_per_session []interface{} `json:"count_per_session,omitempty"`
	Previous_avg_count_per_session float64 `json:"previous_avg_count_per_session,omitempty"`
}

// AgentQueueResponse represents the AgentQueueResponse schema from the OpenAPI specification
type AgentQueueResponse struct {
	Id int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// SigningConfig represents the SigningConfig schema from the OpenAPI specification
type SigningConfig struct {
	Hasstorefile bool `json:"hasStoreFile,omitempty"` // Indicates if storeFile is specified in the signing configuration
}

// AudienceDevicePropertiesListResult represents the AudienceDevicePropertiesListResult schema from the OpenAPI specification
type AudienceDevicePropertiesListResult struct {
	Values map[string]interface{} `json:"values"` // List of device properties.
}

// AndroidProject represents the AndroidProject schema from the OpenAPI specification
type AndroidProject struct {
	Androidmodules []interface{} `json:"androidModules"` // Android Gradle modules
	Gradlewrapperpath string `json:"gradleWrapperPath,omitempty"` // The path of the Gradle wrapper
}

// FeatureFlagsResponse represents the FeatureFlagsResponse schema from the OpenAPI specification
type FeatureFlagsResponse struct {
	Feature_flags []string `json:"feature_flags"`
}

// ValidationResponse represents the ValidationResponse schema from the OpenAPI specification
type ValidationResponse struct {
	App_id string `json:"app_id,omitempty"` // app id
	Status interface{} `json:"status,omitempty"` // Status Data from store
}

// LegacyCollaborators represents the LegacyCollaborators schema from the OpenAPI specification
type LegacyCollaborators struct {
	Collaborators map[string]interface{} `json:"collaborators,omitempty"`
}

// Status represents the Status schema from the OpenAPI specification
type Status struct {
	Status string `json:"status"`
}

// NumberProperty represents the NumberProperty schema from the OpenAPI specification
type NumberProperty struct {
	TypeField string `json:"type"`
	Name string `json:"name"`
}

// PrivateSharedConnectionResponse represents the PrivateSharedConnectionResponse schema from the OpenAPI specification
type PrivateSharedConnectionResponse struct {
	Servicetype string `json:"serviceType"` // service type of shared connection can be apple|gitlab|googleplay|jira|applecertificate
	Displayname string `json:"displayName,omitempty"` // display name of shared connection
	Id string `json:"id"` // id of the shared connection
	Is2fa bool `json:"is2FA,omitempty"` // if the account is a 2FA account or not
	Isvalid bool `json:"isValid,omitempty"` // whether the credentials are valid or not
}

// RetentionSettings represents the RetentionSettings schema from the OpenAPI specification
type RetentionSettings struct {
	Retention_in_days int `json:"retention_in_days"`
}

// GdprVerificationResult represents the GdprVerificationResult schema from the OpenAPI specification
type GdprVerificationResult struct {
	Status bool `json:"status,omitempty"` // Verification status. True means that the verification was successfull.
	Message string `json:"message,omitempty"` // Optional error message if the verification failed.
}

// DeviceRegistrationUrl represents the DeviceRegistrationUrl schema from the OpenAPI specification
type DeviceRegistrationUrl struct {
	Registration_url string `json:"registration_url"` // The url that can be navigated to in order to start the device registration process.
}

// BooleanPropertyDiagnostics represents the BooleanPropertyDiagnostics schema from the OpenAPI specification
type BooleanPropertyDiagnostics struct {
	TypeField string `json:"type"`
	Name string `json:"name"`
}

// Log represents the Log schema from the OpenAPI specification
type Log struct {
	Device map[string]interface{} `json:"device"` // Device characteristics.
	Install_id string `json:"install_id"` // Install ID.
	Timestamp string `json:"timestamp"` // Log creation timestamp.
	TypeField string `json:"type"` // Log type.
}

// ReleaseMetadata represents the ReleaseMetadata schema from the OpenAPI specification
type ReleaseMetadata struct {
	Dsa_signature string `json:"dsa_signature,omitempty"` // dsa signature of the release for the sparkle feed.
	Ed_signature string `json:"ed_signature,omitempty"` // edDSA signature of the release for the sparkle feed.
}

// JiraConnectionSecretResponse represents the JiraConnectionSecretResponse schema from the OpenAPI specification
type JiraConnectionSecretResponse struct {
	Credentialtype string `json:"credentialType,omitempty"` // the type of the credential
	Displayname string `json:"displayName"` // display name of shared connection
	Servicetype string `json:"serviceType"` // service type of shared connection can be apple|gitlab|googleplay|jira
}

// DistributionStartSessionLog represents the DistributionStartSessionLog schema from the OpenAPI specification
type DistributionStartSessionLog struct {
	Timestamp string `json:"timestamp"` // Log creation timestamp.
	TypeField string `json:"type"` // Log type.
	Device map[string]interface{} `json:"device"` // Device characteristics.
	Install_id string `json:"install_id"` // Install ID.
}

// ErrorResponse represents the ErrorResponse schema from the OpenAPI specification
type ErrorResponse struct {
	ErrorField interface{} `json:"error"`
}

// StepReport represents the StepReport schema from the OpenAPI specification
type StepReport struct {
	Finishedsnapshots []string `json:"finishedSnapshots"`
	Devicescreenshots []map[string]interface{} `json:"deviceScreenshots"`
}

// InternalUserSignupResponse represents the InternalUserSignupResponse schema from the OpenAPI specification
type InternalUserSignupResponse struct {
	External_user_id string `json:"external_user_id,omitempty"` // The user ID given by the external provider
	Id string `json:"id"` // The unique id (UUID) of the user
	Name string `json:"name"` // The unique name that is used to identify the user.
	Status string `json:"status,omitempty"` // The current status of the user record after signup
	Display_name string `json:"display_name"` // The full name of the user. Might for example be first and last name
	Email string `json:"email"` // The email address of the user
	External_provider string `json:"external_provider,omitempty"` // The name of the external auth provider
}

// CodePushStatusMetricMetadata represents the CodePushStatusMetricMetadata schema from the OpenAPI specification
type CodePushStatusMetricMetadata struct {
	App_version string `json:"app_version,omitempty"`
	Client_unique_id string `json:"client_unique_id,omitempty"`
	Deployment_key string `json:"deployment_key"`
	Label string `json:"label,omitempty"`
	Previous_deployment_key string `json:"previous_deployment_key,omitempty"`
	Previous_label_or_app_version string `json:"previous_label_or_app_version,omitempty"`
	Status string `json:"status,omitempty"`
}

// GooglePlayConnectionSecretRequest represents the GooglePlayConnectionSecretRequest schema from the OpenAPI specification
type GooglePlayConnectionSecretRequest struct {
	Credentialtype string `json:"credentialType,omitempty"` // credential type of the shared connection. Values can be credentials|certificate
	Data map[string]interface{} `json:"data,omitempty"` // shared connection details
	Displayname string `json:"displayName,omitempty"` // display name of shared connection
	Servicetype string `json:"serviceType"` // service type of shared connection can be apple|gitlab|googleplay|jira|applecertificate
}

// OrgDistributionGroupAppResponse represents the OrgDistributionGroupAppResponse schema from the OpenAPI specification
type OrgDistributionGroupAppResponse struct {
	Icon_source string `json:"icon_source,omitempty"` // The string representation of the source of the app's icon
	Release_type string `json:"release_type,omitempty"` // A one-word descriptive release-type value that starts with a capital letter but is otherwise lowercase
	Display_name string `json:"display_name"` // The display name of the app
	Name string `json:"name"` // The name of the app used in URLs
	Description string `json:"description,omitempty"` // The description of the app
	Os string `json:"os"` // The OS the app will be running on
	Icon_url string `json:"icon_url,omitempty"` // The string representation of the URL pointing to the app's icon
	Id string `json:"id"` // The unique ID (UUID) of the app
	Owner interface{} `json:"owner"` // The information about the app's owner
	Platform string `json:"platform,omitempty"` // The platform of the app
	Origin string `json:"origin,omitempty"` // The creation origin of this app
}

// AlertWebhook represents the AlertWebhook schema from the OpenAPI specification
type AlertWebhook struct {
	Event_types []string `json:"event_types"` // Event types enabled for webhook
	Id string `json:"id,omitempty"` // The unique id (UUID) of the webhook
	Name string `json:"name"` // display name of the webhook
	Url string `json:"url"` // target url of the webhook
	Enabled bool `json:"enabled,omitempty"` // Allows eanble/disable webhook
}

// AudienceDevicePropertyValuesListResult represents the AudienceDevicePropertyValuesListResult schema from the OpenAPI specification
type AudienceDevicePropertyValuesListResult struct {
	Values []string `json:"values"` // List of device property values.
}

// TesterAppWithReleaseResponse represents the TesterAppWithReleaseResponse schema from the OpenAPI specification
type TesterAppWithReleaseResponse struct {
	Release interface{} `json:"release,omitempty"`
	Description string `json:"description,omitempty"` // The description of the app
	Display_name string `json:"display_name,omitempty"` // The app's display name.
	Icon_url string `json:"icon_url,omitempty"` // A URL to the app's icon.
	Id string `json:"id,omitempty"` // The unique ID (UUID) of the app
	Name string `json:"name,omitempty"` // The app's name.
	Os string `json:"os,omitempty"` // The app's os.
	Owner map[string]interface{} `json:"owner,omitempty"` // The information about the app's owner
}

// AlertingGithubBugtrackerSettings represents the AlertingGithubBugtrackerSettings schema from the OpenAPI specification
type AlertingGithubBugtrackerSettings struct {
	TypeField string `json:"type"` // type of bugtracker
	Callback_url string `json:"callback_url,omitempty"`
	Owner_name string `json:"owner_name"`
}

// BranchStatus represents the BranchStatus schema from the OpenAPI specification
type BranchStatus struct {
	Configured bool `json:"configured"`
	Lastbuild interface{} `json:"lastBuild,omitempty"`
}

// PrivateCreateStoreRequest represents the PrivateCreateStoreRequest schema from the OpenAPI specification
type PrivateCreateStoreRequest struct {
	Intune_details interface{} `json:"intune_details,omitempty"`
	Name string `json:"name,omitempty"` // name of the store.
	TypeField string `json:"type,omitempty"` // store Type
}

// ExportConfigurationAppInsightsKey represents the ExportConfigurationAppInsightsKey schema from the OpenAPI specification
type ExportConfigurationAppInsightsKey struct {
	Export_entities []string `json:"export_entities,omitempty"`
	Resource_group string `json:"resource_group,omitempty"` // The resource group name on azure
	Resource_name string `json:"resource_name,omitempty"` // The resource name on azure
	TypeField string `json:"type"` // Type of export configuration
	Backfill bool `json:"backfill,omitempty"` // Field to determine if backfilling should occur. The default value is true. If set to false export starts from date and time of config creation.
}

// AlertingError represents the AlertingError schema from the OpenAPI specification
type AlertingError struct {
	Request_id string `json:"request_id"` // Unique request identifier for tracking
	Code int `json:"code"` // The status code return by the API. It can be 400 or 404 or 409 or 500.
	Message string `json:"message,omitempty"` // The reason for the request failed
}

// ErrorAttachment represents the ErrorAttachment schema from the OpenAPI specification
type ErrorAttachment struct {
	Contenttype string `json:"contentType,omitempty"`
	Crashid string `json:"crashId,omitempty"`
	Createdtime string `json:"createdTime,omitempty"`
	Filename string `json:"fileName,omitempty"`
	Size int64 `json:"size,omitempty"`
	Appid string `json:"appId,omitempty"`
	Attachmentid string `json:"attachmentId,omitempty"`
	Bloblocation string `json:"blobLocation,omitempty"`
}

// JiraConnectionSecretRequest represents the JiraConnectionSecretRequest schema from the OpenAPI specification
type JiraConnectionSecretRequest struct {
	Servicetype string `json:"serviceType"` // service type of shared connection can be apple|gitlab|googleplay|jira|applecertificate
	Credentialtype string `json:"credentialType,omitempty"` // credential type of the shared connection. Values can be credentials|certificate
	Data map[string]interface{} `json:"data,omitempty"` // shared connection details
	Displayname string `json:"displayName,omitempty"` // display name of shared connection
}

// BugTrackerStateResult represents the BugTrackerStateResult schema from the OpenAPI specification
type BugTrackerStateResult struct {
	State string `json:"state,omitempty"` // bugtracker state
}

// LegacyCodePushReleaseResponse represents the LegacyCodePushReleaseResponse schema from the OpenAPI specification
type LegacyCodePushReleaseResponse struct {
	PackageField interface{} `json:"package"`
}

// DeviceSelection represents the DeviceSelection schema from the OpenAPI specification
type DeviceSelection struct {
	Shortid string `json:"shortId"` // Identifier of the device selection
}

// Versions represents the Versions schema from the OpenAPI specification
type Versions struct {
	Total int64 `json:"total,omitempty"` // The total count of versions.
	Versions []map[string]interface{} `json:"versions,omitempty"` // List of version count.
}

// AudienceListResult represents the AudienceListResult schema from the OpenAPI specification
type AudienceListResult struct {
	Nextlink string `json:"nextLink,omitempty"`
	Values []interface{} `json:"values"` // List of audiences.
}

// BuildInfo represents the BuildInfo schema from the OpenAPI specification
type BuildInfo struct {
	Commit_message string `json:"commit_message,omitempty"` // The commit message of the build producing the release
	Branch_name string `json:"branch_name,omitempty"` // The branch name of the build producing the release
	Commit_hash string `json:"commit_hash,omitempty"` // The commit hash of the build producing the release
}

// StatusData represents the StatusData schema from the OpenAPI specification
type StatusData struct {
	Track string `json:"track,omitempty"` // track information from store
	Version string `json:"version,omitempty"` // version of the app from store
	Status string `json:"status,omitempty"` // status from store
	Storetype string `json:"storetype,omitempty"` // store type
}

// ExportConfigurationAppInsightsLinkedSubscription represents the ExportConfigurationAppInsightsLinkedSubscription schema from the OpenAPI specification
type ExportConfigurationAppInsightsLinkedSubscription struct {
	Export_entities []string `json:"export_entities,omitempty"`
	Resource_group string `json:"resource_group,omitempty"` // The resource group name on azure
	Resource_name string `json:"resource_name,omitempty"` // The resource name on azure
	TypeField string `json:"type"` // Type of export configuration
	Backfill bool `json:"backfill,omitempty"` // Field to determine if backfilling should occur. The default value is true. If set to false export starts from date and time of config creation.
}

// ApiTokensPostRequest represents the ApiTokensPostRequest schema from the OpenAPI specification
type ApiTokensPostRequest struct {
	Encrypted_token string `json:"encrypted_token,omitempty"` // An encrypted value of the token.
	Scope []string `json:"scope,omitempty"` // The scope for this token. An array of supported roles.
	Token_hash string `json:"token_hash,omitempty"` // The hashed value of api token
	Token_type string `json:"token_type,omitempty"` // The token's type. public:managed by the user; in_app_update:special token for in-app update scenario; buid:dedicated for CI usage for now; session:for CLI session management; tester_app: used for tester mobile app; default is "public".'
	Description string `json:"description,omitempty"` // The description of the token
}

// LogFlowEventLog represents the LogFlowEventLog schema from the OpenAPI specification
type LogFlowEventLog struct {
	Device map[string]interface{} `json:"device"` // Device characteristics.
	Install_id string `json:"install_id"` // Install ID.
	Timestamp string `json:"timestamp"` // Log creation timestamp.
	TypeField string `json:"type"` // Log type.
	Properties map[string]interface{} `json:"properties,omitempty"` // Additional key/value pair parameters.
}

// AppleConnectionNonSecretResponse represents the AppleConnectionNonSecretResponse schema from the OpenAPI specification
type AppleConnectionNonSecretResponse struct {
	Credentialtype string `json:"credentialType,omitempty"` // the type of the credential
	Displayname string `json:"displayName"` // display name of shared connection
	Servicetype string `json:"serviceType"` // service type of shared connection can be apple|gitlab|googleplay|jira
}

// BranchConfigurationArtifactVersioning represents the BranchConfigurationArtifactVersioning schema from the OpenAPI specification
type BranchConfigurationArtifactVersioning struct {
	Buildnumberformat string `json:"buildNumberFormat,omitempty"`
}

// ToolsetProject represents the ToolsetProject schema from the OpenAPI specification
type ToolsetProject struct {
	Name string `json:"name,omitempty"`
	Path string `json:"path,omitempty"`
}

// CrashGroupChange represents the CrashGroupChange schema from the OpenAPI specification
type CrashGroupChange struct {
	Annotation string `json:"annotation,omitempty"`
	Status interface{} `json:"status,omitempty"`
}

// LegacyDeployment represents the LegacyDeployment schema from the OpenAPI specification
type LegacyDeployment struct {
	Key string `json:"key,omitempty"` // Deployment key (aka Deployment Id)
	Name string `json:"name"` // Updated deployment name
	PackageField interface{} `json:"package,omitempty"`
	Createdtime int `json:"createdTime,omitempty"` // Time at which the deployment was created as a Unix timestamp.
	Id string `json:"id,omitempty"` // The ID of the deployment (internal use only).
}

// DistributionRequest represents the DistributionRequest schema from the OpenAPI specification
type DistributionRequest struct {
	Notifytesters bool `json:"notifyTesters,omitempty"`
	Releasenotes string `json:"releaseNotes,omitempty"` // The release notes
	Destinations []interface{} `json:"destinations,omitempty"` // Array of objects {id:string, type:string} with "id" being the distribution group ID, store ID, or tester email, and "type" being "group", "store", or "tester"
	Mandatoryupdate bool `json:"mandatoryUpdate,omitempty"`
}

// UserAuthResponse represents the UserAuthResponse schema from the OpenAPI specification
type UserAuthResponse struct {
	External_provider string `json:"external_provider,omitempty"` // The name of the external auth provider
	External_user_id string `json:"external_user_id,omitempty"` // The user ID given by the external provider
	Id string `json:"id"` // The unique id (UUID) of the user
	Name string `json:"name"` // The unique name that is used to identify the user
	Origin string `json:"origin"` // The creation origin of this user
	Display_name string `json:"display_name"` // The full name of the user. Might for example be first and last name
	Email string `json:"email"` // The email address of the user
}

// ReleaseRealTimeStatusResponse represents the ReleaseRealTimeStatusResponse schema from the OpenAPI specification
type ReleaseRealTimeStatusResponse struct {
	App_id string `json:"app_id,omitempty"` // app id
	Release_id string `json:"release_id,omitempty"` // release id
	Status interface{} `json:"status,omitempty"` // Status Data from store
}

// CrashGroupAndVersion represents the CrashGroupAndVersion schema from the OpenAPI specification
type CrashGroupAndVersion struct {
	App_version string `json:"app_version,omitempty"`
	Crash_group_id string `json:"crash_group_id,omitempty"`
}

// CrashFreeDevicePercentages represents the CrashFreeDevicePercentages schema from the OpenAPI specification
type CrashFreeDevicePercentages struct {
	Average_percentage float64 `json:"average_percentage,omitempty"` // Average percentage.
	Daily_percentages []interface{} `json:"daily_percentages,omitempty"` // The crash-free percentage per day.
}

// TestGDPRFileSetFile represents the TestGDPRFileSetFile schema from the OpenAPI specification
type TestGDPRFileSetFile struct {
	Hash_file_url string `json:"hash_file_url,omitempty"`
	Path string `json:"path,omitempty"`
	App_upload_id string `json:"app_upload_id,omitempty"`
	Hash_file_id string `json:"hash_file_id,omitempty"`
}

// ApplicationStatusResponse represents the ApplicationStatusResponse schema from the OpenAPI specification
type ApplicationStatusResponse struct {
	Version string `json:"version,omitempty"` // The version of the application
	Version_type string `json:"version_type"` // The type of version being returned (production/edit/test flight).
}

// SendNotificationRequest represents the SendNotificationRequest schema from the OpenAPI specification
type SendNotificationRequest struct {
	Userids []string `json:"userIds"` // user list to send email notification
}

// ValidationErrorResponse represents the ValidationErrorResponse schema from the OpenAPI specification
type ValidationErrorResponse struct {
	Id string `json:"id"`
	Message string `json:"message"`
	Code string `json:"code"`
}

// CreateGdprTestData represents the CreateGdprTestData schema from the OpenAPI specification
type CreateGdprTestData struct {
	Appname string `json:"appName,omitempty"`
	Useremail string `json:"userEmail,omitempty"`
	Username string `json:"userName,omitempty"`
	Accountid string `json:"accountId,omitempty"`
	Appid string `json:"appId,omitempty"`
}

// HasBuildUploadedResponse represents the HasBuildUploadedResponse schema from the OpenAPI specification
type HasBuildUploadedResponse struct {
	Has_build_uploaded bool `json:"has_build_uploaded,omitempty"` // true if a build has been uploaded, false otherwise
}

// ErrorGroupOperatingSystems represents the ErrorGroupOperatingSystems schema from the OpenAPI specification
type ErrorGroupOperatingSystems struct {
	Errorcount int64 `json:"errorCount,omitempty"`
	Operatingsystems []map[string]interface{} `json:"operatingSystems,omitempty"`
}

// TesterAppRelease represents the TesterAppRelease schema from the OpenAPI specification
type TesterAppRelease struct {
	Enabled bool `json:"enabled"` // This value determines the whether a release currently is enabled or disabled.
	Id int `json:"id"` // ID identifying this unique release.
	Is_external_build bool `json:"is_external_build,omitempty"` // This value determines if a release is external or not.
	Mandatory_update bool `json:"mandatory_update"` // A boolean which determines whether the release is a mandatory update or not.
	Origin string `json:"origin,omitempty"` // The release's origin
	Short_version string `json:"short_version"` // The release's short version.<br> For iOS: CFBundleShortVersionString from info.plist.<br> For Android: android:versionName from AppManifest.xml.
	Uploaded_at string `json:"uploaded_at"` // UTC time in ISO 8601 format of the uploaded time.
	Version string `json:"version"` // The release's version.<br> For iOS: CFBundleVersion from info.plist.<br> For Android: android:versionCode from AppManifest.xml.
	Install_url string `json:"install_url,omitempty"` // The href required to install a release on a mobile device. On iOS devices will be prefixed with `itms-services://?action=download-manifest&url=`
	Release_notes string `json:"release_notes,omitempty"` // The release's release notes.
	Size int `json:"size"` // The release's size in bytes.
}

// UserNameAvailabilityResponse represents the UserNameAvailabilityResponse schema from the OpenAPI specification
type UserNameAvailabilityResponse struct {
	Available bool `json:"available"` // The availability status of the requested user name
	Name string `json:"name"` // The requested user name
}

// ErrorGroupModel represents the ErrorGroupModel schema from the OpenAPI specification
type ErrorGroupModel struct {
	Errorcount int64 `json:"errorCount,omitempty"` // count of errors in a model
	Modelcode string `json:"modelCode,omitempty"` // model code
	Modelname string `json:"modelName,omitempty"` // model name
}

// UserProfileAdminResponse represents the UserProfileAdminResponse schema from the OpenAPI specification
type UserProfileAdminResponse struct {
	Avatar_url string `json:"avatar_url,omitempty"` // The avatar URL of the user
	Can_change_password bool `json:"can_change_password,omitempty"` // User is required to send an old password in order to change the password.
	Display_name string `json:"display_name"` // The full name of the user. Might for example be first and last name
	Email string `json:"email"` // The email address of the user
	Id string `json:"id"` // The unique id (UUID) of the user
	Name string `json:"name"` // The unique name that is used to identify the user.
	Origin string `json:"origin"` // The creation origin of this user
	Permissions []string `json:"permissions,omitempty"` // The permissions the user has for the app
	Role string `json:"role,omitempty"` // The user's role in the organization
}

// LegacyDeploymentMetric represents the LegacyDeploymentMetric schema from the OpenAPI specification
type LegacyDeploymentMetric struct {
	Failed int `json:"failed,omitempty"` // The number of times this release has failed to be installed on a device
	Installed int `json:"installed,omitempty"` // The number of times this release has been installed on a device
	Active int `json:"active"` // The number of devices that have this release installed currently
	Downloaded int `json:"downloaded,omitempty"` // The number of times this release has been downloaded
}

// UsagePeriod represents the UsagePeriod schema from the OpenAPI specification
type UsagePeriod struct {
	Endtime string `json:"endTime,omitempty"` // Exclusive end time of the usage period.
	Starttime string `json:"startTime,omitempty"` // Inclusive start time of the usage period
	Byaccount map[string]interface{} `json:"byAccount,omitempty"` // A collection of named numeric values
	Byapp map[string]interface{} `json:"byApp,omitempty"` // A collection of named numeric values grouped by app
}

// ResignInfo represents the ResignInfo schema from the OpenAPI specification
type ResignInfo struct {
	Certificate_expiration string `json:"certificate_expiration,omitempty"` // The expiration date of the certificate used for the resign attempt
	Certificate_name string `json:"certificate_name,omitempty"` // The name of the certificate used for the resign attempt
	Group_name string `json:"group_name,omitempty"` // The group name of the resign attempt
	Profile_name string `json:"profile_name,omitempty"` // The provisioning profile name of group for the given resign attempt
	Profile_type string `json:"profile_type,omitempty"` // The provisioning profile type of group for the given resign attempt
}

// ErrorGroupOperatingSystem represents the ErrorGroupOperatingSystem schema from the OpenAPI specification
type ErrorGroupOperatingSystem struct {
	Operatingsystemname string `json:"operatingSystemName,omitempty"` // OS name
	Errorcount int64 `json:"errorCount,omitempty"` // count of OS
}

// SymbolUploadBeginRequest represents the SymbolUploadBeginRequest schema from the OpenAPI specification
type SymbolUploadBeginRequest struct {
	Client_callback string `json:"client_callback,omitempty"` // The callback URL that the client can optionally provide to get status updates for the current symbol upload
	File_name string `json:"file_name,omitempty"` // The file name for the symbol upload
	Symbol_type string `json:"symbol_type"` // The type of the symbol for the current symbol upload
	Version string `json:"version,omitempty"` // The version number. Optional for Apple. Required for Android.
	Build string `json:"build,omitempty"` // The build number. Optional for Apple. Required for Android.
}

// StoreNotification represents the StoreNotification schema from the OpenAPI specification
type StoreNotification struct {
	Service string `json:"service,omitempty"`
	Status string `json:"status,omitempty"`
	Valid_until int `json:"valid_until,omitempty"`
}

// AppUserPermissionResponse represents the AppUserPermissionResponse schema from the OpenAPI specification
type AppUserPermissionResponse struct {
	User_id string `json:"user_id"` // The unique id (UUID) of the user
	App_id string `json:"app_id"` // The unique id (UUID) of the app
	App_origin string `json:"app_origin"` // The creation origin of this app
	App_secret string `json:"app_secret"` // A unique and secret key used to identify the app in communication with the ingestion endpoint for crash reporting and analytics
	Permissions []string `json:"permissions"` // The permissions the user has for the app
	User_email string `json:"user_email"` // The email of the user
}

// GDPRInvitationDetailResponse represents the GDPRInvitationDetailResponse schema from the OpenAPI specification
type GDPRInvitationDetailResponse struct {
	App interface{} `json:"app,omitempty"`
	Invitation_id string `json:"invitation_id"` // The id of the invitation
	Organization interface{} `json:"organization,omitempty"`
}

// SymbolUploadBeginResponse represents the SymbolUploadBeginResponse schema from the OpenAPI specification
type SymbolUploadBeginResponse struct {
	Expiration_date string `json:"expiration_date"` // Describes how long the upload_url is valid
	Symbol_upload_id string `json:"symbol_upload_id"` // The id for the current upload
	Upload_url string `json:"upload_url"` // The URL where the client needs to upload the symbol blob to
}

// LegacyApp represents the LegacyApp schema from the OpenAPI specification
type LegacyApp struct {
	Collaborators map[string]interface{} `json:"collaborators,omitempty"`
	Deployments []string `json:"deployments,omitempty"`
	Name string `json:"name,omitempty"` // The app name.
}

// ClearProperty represents the ClearProperty schema from the OpenAPI specification
type ClearProperty struct {
	Name string `json:"name"`
	TypeField string `json:"type"`
}

// PostCreateReleaseUploadRequest represents the PostCreateReleaseUploadRequest schema from the OpenAPI specification
type PostCreateReleaseUploadRequest struct {
	Build_number string `json:"build_number,omitempty"` // User defined build number
	Build_version string `json:"build_version,omitempty"` // User defined build version
}

// OptimizelyUserMetaDataResponse represents the OptimizelyUserMetaDataResponse schema from the OpenAPI specification
type OptimizelyUserMetaDataResponse struct {
	Userid string `json:"userId,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// ReleaseCount represents the ReleaseCount schema from the OpenAPI specification
type ReleaseCount struct {
	Distribution_group string `json:"distribution_group,omitempty"` // Distribution group queried.
	Release_id string `json:"release_id"`
	Total_count int64 `json:"total_count"` // Total count of downloads.
	Unique_count int64 `json:"unique_count"` // Count of unique downloads against user id.
}

// OrganizationInvitationSimpleDetailResponse represents the OrganizationInvitationSimpleDetailResponse schema from the OpenAPI specification
type OrganizationInvitationSimpleDetailResponse struct {
	Email string `json:"email"` // The email address of the invited user
	Id string `json:"id"` // The unique ID (UUID) of the invitation
	Role string `json:"role"` // The role assigned to the invited user
}

// BranchConfigurationToolsets represents the BranchConfigurationToolsets schema from the OpenAPI specification
type BranchConfigurationToolsets struct {
	Xamarin interface{} `json:"xamarin,omitempty"` // Build configuration for Xamarin projects
	Xcode interface{} `json:"xcode,omitempty"` // Build configuration when Xcode is part of the build steps
	Android interface{} `json:"android,omitempty"` // Build configuration for Android projects
	Javascript interface{} `json:"javascript,omitempty"` // Build configuration when React Native, or other JavaScript tech, is part of the build steps
}

// LogFlowStartServiceLog represents the LogFlowStartServiceLog schema from the OpenAPI specification
type LogFlowStartServiceLog struct {
	Device map[string]interface{} `json:"device"` // Device characteristics.
	Install_id string `json:"install_id"` // Install ID.
	Timestamp string `json:"timestamp"` // Log creation timestamp.
	TypeField string `json:"type"` // Log type.
}

// DeviceDimensions represents the DeviceDimensions schema from the OpenAPI specification
type DeviceDimensions struct {
	Height map[string]interface{} `json:"height,omitempty"`
	Width map[string]interface{} `json:"width,omitempty"`
	Depth map[string]interface{} `json:"depth,omitempty"`
}

// OrganizationInvitationDetailResponse represents the OrganizationInvitationDetailResponse schema from the OpenAPI specification
type OrganizationInvitationDetailResponse struct {
	Id string `json:"id"` // The unique ID (UUID) of the invitation
	Invited_by interface{} `json:"invited_by"`
	Is_existing_user bool `json:"is_existing_user"` // Indicates whether the invited user already exists
	Organization interface{} `json:"organization"`
	Role string `json:"role,omitempty"` // The role assigned to the invited user
	Email string `json:"email"` // The email address of the invited user
}

// PatchReleaseRequest represents the PatchReleaseRequest schema from the OpenAPI specification
type PatchReleaseRequest struct {
	Is_wrapper_request bool `json:"is_wrapper_request,omitempty"` // request is for wrapping or not
	Status string `json:"status,omitempty"` // updated status of release
	Wrap_package_url string `json:"wrap_package_url,omitempty"` // package url for wrapping request
	Dest_publish_id string `json:"dest_publish_id,omitempty"` // Destination Publish Id
	Error_contextid string `json:"error_contextId,omitempty"` // contextId for failed error message
	Error_details string `json:"error_details,omitempty"` // failure error details from store
}

// AcquisitionStatusSuccessResponse represents the AcquisitionStatusSuccessResponse schema from the OpenAPI specification
type AcquisitionStatusSuccessResponse struct {
	Code string `json:"code"` // The code indicating the status
	Message string `json:"message"` // The message indicating the status
}

// StringPropertyDiagnostics represents the StringPropertyDiagnostics schema from the OpenAPI specification
type StringPropertyDiagnostics struct {
	Name string `json:"name"`
	TypeField string `json:"type"`
}

// UserUpdateRequest represents the UserUpdateRequest schema from the OpenAPI specification
type UserUpdateRequest struct {
	Display_name string `json:"display_name,omitempty"` // The full name of the user. Might for example be first and last name
}

// DistributionGroupTesterGetResponse represents the DistributionGroupTesterGetResponse schema from the OpenAPI specification
type DistributionGroupTesterGetResponse struct {
	Display_name string `json:"display_name,omitempty"` // The full name of the tester. Might for example be first and last name
	Email string `json:"email"` // The email address of the tester
	Name string `json:"name"` // The unique name that is used to identify the tester.
}

// IntuneAppsRequest represents the IntuneAppsRequest schema from the OpenAPI specification
type IntuneAppsRequest struct {
	Created_month string `json:"created_month,omitempty"` // PartitionKey year-month
}

// DistributionGroupRelease represents the DistributionGroupRelease schema from the OpenAPI specification
type DistributionGroupRelease struct {
	Uploaded_at string `json:"uploaded_at"` // UTC time in ISO 8601 format of the uploaded time.
	Version string `json:"version"` // The release's version.<br> For iOS: CFBundleVersion from info.plist.<br> For Android: android:versionCode from AppManifest.xml.
	Enabled bool `json:"enabled"` // This value determines the whether a release currently is enabled or disabled.
	Id int `json:"id"` // ID identifying this unique release.
	Is_external_build bool `json:"is_external_build,omitempty"` // This value determines if a release is external or not.
	Mandatory_update bool `json:"mandatory_update"` // A boolean which determines whether the release is a mandatory update or not.
	Origin string `json:"origin,omitempty"` // The release's origin
	Short_version string `json:"short_version"` // The release's short version.<br> For iOS: CFBundleShortVersionString from info.plist.<br> For Android: android:versionName from AppManifest.xml.
}

// AppMembershipsValidationResponse represents the AppMembershipsValidationResponse schema from the OpenAPI specification
type AppMembershipsValidationResponse struct {
	Distribution_group_users_with_missing_memberships []interface{} `json:"distribution_group_users_with_missing_memberships,omitempty"`
	Excess_app_memberships interface{} `json:"excess_app_memberships,omitempty"`
	Organization_admins_with_missing_memberships []interface{} `json:"organization_admins_with_missing_memberships,omitempty"`
	App_users_with_missing_memberships []interface{} `json:"app_users_with_missing_memberships,omitempty"`
}

// AppleLoginRequest represents the AppleLoginRequest schema from the OpenAPI specification
type AppleLoginRequest struct {
	Cookie string `json:"cookie,omitempty"` // The 30-day session cookie for multi-factor authentication backed accounts.
	Password string `json:"password"` // The password for the Apple Developer account.
	Team_identifier string `json:"team_identifier,omitempty"` // Identifier of the team to use when logged in.
	Username string `json:"username"` // The username for the Apple Developer account.
}

// ExportConfigurationResult represents the ExportConfigurationResult schema from the OpenAPI specification
type ExportConfigurationResult struct {
	Creation_time string `json:"creation_time"` // Creation time in ISO 8601 format
	Export_configuration map[string]interface{} `json:"export_configuration,omitempty"` // Export configuration
	Export_entities []string `json:"export_entities,omitempty"`
	Id string `json:"id"` // Export configuration id
	State string `json:"state"` // State of the export job
	Export_type string `json:"export_type"` // Target resource type of export configuration
	Resource_group string `json:"resource_group,omitempty"` // resource group for the storage account/App Insights resource
	State_info string `json:"state_info,omitempty"` // Additional information about export configuration state
	Last_run_time string `json:"last_run_time,omitempty"` // Latest time in ISO 8601 format when export completed successfully
	Resource_name string `json:"resource_name,omitempty"` // Storage accout or Appinsights resource name
}

// UWPToolset represents the UWPToolset schema from the OpenAPI specification
type UWPToolset struct {
	Uwpsolutions []interface{} `json:"uwpSolutions"` // The UWP solutions detected
}

// PublishDevicesRequest represents the PublishDevicesRequest schema from the OpenAPI specification
type PublishDevicesRequest struct {
	Account_service_connection_id string `json:"account_service_connection_id,omitempty"` // The service_connection_id of the stored Apple credentials instead of username, password.
	Devices []string `json:"devices,omitempty"` // Array of device UDID's to be published to the Apple Developer account.
	Password string `json:"password,omitempty"` // The password for the Apple Developer account to publish the devices to.
	Publish_all_devices bool `json:"publish_all_devices,omitempty"` // When set to true, all unprovisioned devices will be published to the Apple Developer account. When false, only the provided devices will be published to the Apple Developer account.
	Username string `json:"username,omitempty"` // The username for the Apple Developer account to publish the devices to.
}

// AppleTestFlightGroupRequest represents the AppleTestFlightGroupRequest schema from the OpenAPI specification
type AppleTestFlightGroupRequest struct {
	Service_connection_id string `json:"service_connection_id,omitempty"` // The service_connection_id of the stored Apple credentials instead of username, password.
	Team_identifier string `json:"team_identifier,omitempty"` // Identifier of the team to use when logged in.
	Username string `json:"username,omitempty"` // The username for the Apple Developer account.
	Apple_id string `json:"apple_id,omitempty"` // apple_id of the app for which test flight groups need to be fetched.
	Bundle_identifier string `json:"bundle_identifier,omitempty"` // apple_id of the app for which test flight groups need to be fetched.
	Cookie string `json:"cookie,omitempty"` // The 30-day session cookie for multi-factor authentication backed accounts.
	Password string `json:"password,omitempty"` // The password for the Apple Developer account.
}

// AppleMutifactorLoginRequest represents the AppleMutifactorLoginRequest schema from the OpenAPI specification
type AppleMutifactorLoginRequest struct {
	Authcode string `json:"authcode"` // This is the six digit OTP used for completing the multi-factor authentication
	Username string `json:"username"` // The username for the Apple Developer account.
}

// Build represents the Build schema from the OpenAPI specification
type Build struct {
	Queuetime string `json:"queueTime"` // The time the build was queued
	Result string `json:"result"` // The build result
	Sourceversion string `json:"sourceVersion"` // The source SHA
	Starttime string `json:"startTime,omitempty"` // The time the build was started
	Sourcebranch string `json:"sourceBranch"` // The source branch name
	Buildnumber string `json:"buildNumber"` // The build number
	Lastchangeddate string `json:"lastChangedDate,omitempty"` // The time the build status was last changed
	Status string `json:"status"` // The build status
	Finishtime string `json:"finishTime,omitempty"` // The time the build was finished
	Id int `json:"id"` // The build ID
}

// AlertingBugtracker represents the AlertingBugtracker schema from the OpenAPI specification
type AlertingBugtracker struct {
	Token_id string `json:"token_id,omitempty"` // ID of OAuth token
	TypeField string `json:"type,omitempty"` // type of bugtracker
	Event_types []string `json:"event_types,omitempty"` // Event types enabled for bugtracker
	Settings map[string]interface{} `json:"settings,omitempty"` // Bugtracker specific settings
	State string `json:"state,omitempty"` // bugtracker state
}

// LegacyAuthenticationResponse represents the LegacyAuthenticationResponse schema from the OpenAPI specification
type LegacyAuthenticationResponse struct {
	Authenticated bool `json:"authenticated,omitempty"` // The authentication status of the user.
}

// LogDiagnostics represents the LogDiagnostics schema from the OpenAPI specification
type LogDiagnostics struct {
	Device map[string]interface{} `json:"device"` // Device characteristics.
	Install_id string `json:"install_id"` // Install ID.
	Timestamp string `json:"timestamp"` // Log creation timestamp.
	TypeField string `json:"type"` // Log type.
}

// UsageRecordStatus represents the UsageRecordStatus schema from the OpenAPI specification
type UsageRecordStatus struct {
	Latestbuildusagerecordtime string `json:"latestBuildUsageRecordTime,omitempty"` // The time of the most recent Build service usage record
	Latesttestusagerecordtime string `json:"latestTestUsageRecordTime,omitempty"` // The time of the most recent Test service usage record
	Expectedlatestbuildexists bool `json:"expectedLatestBuildExists,omitempty"` // Is the age of the most recent Build service usage record within expected limits
	Expectedlatesttestexists bool `json:"expectedLatestTestExists,omitempty"` // Is the age of the most recent Test service usage record within expected limits
}

// AppId represents the AppId schema from the OpenAPI specification
type AppId struct {
	Id string `json:"id,omitempty"` // the id of the app
}

// LegacyDeploymentHistoryResponse represents the LegacyDeploymentHistoryResponse schema from the OpenAPI specification
type LegacyDeploymentHistoryResponse struct {
	History []interface{} `json:"history,omitempty"` // Array containing the deployment's package history.
}

// AlertUserAppEmailSettingsResult represents the AlertUserAppEmailSettingsResult schema from the OpenAPI specification
type AlertUserAppEmailSettingsResult struct {
	Request_id string `json:"request_id"` // Unique request identifier for tracking
	Etag string `json:"eTag,omitempty"` // The ETag of the entity
	Enabled bool `json:"enabled"` // Allows to forcefully disable emails on app or user level
	Settings []map[string]interface{} `json:"settings"` // The settings the user has for the app
	Userid string `json:"userId,omitempty"` // The unique id (UUID) of the user
	Appid string `json:"appId,omitempty"` // Application ID
	Request_id string `json:"request_id,omitempty"` // Unique request
	User_enabled bool `json:"user_enabled,omitempty"` // A flag indicating if settings are enabled at user/global level
}

// AudienceSummary represents the AudienceSummary schema from the OpenAPI specification
type AudienceSummary struct {
	Name string `json:"name,omitempty"` // Audience name.
	State string `json:"state,omitempty"` // Audience state.
	Definition string `json:"definition,omitempty"` // Audience definition in OData format.
	Description string `json:"description,omitempty"` // Audience description.
	Estimated_count int64 `json:"estimated_count,omitempty"` // Estimated audience size.
}

// IntuneAppCategoryResponse represents the IntuneAppCategoryResponse schema from the OpenAPI specification
type IntuneAppCategoryResponse struct {
	Name string `json:"name,omitempty"` // display name for the app category
	Id string `json:"id,omitempty"` // ID for the category.
}

// TestGDPRTestRun represents the TestGDPRTestRun schema from the OpenAPI specification
type TestGDPRTestRun struct {
	Dsym_hash_file_url string `json:"dsym_hash_file_url,omitempty"`
	Id string `json:"id,omitempty"`
	Locale string `json:"locale,omitempty"`
	App_hash_file_id string `json:"app_hash_file_id,omitempty"`
	App_hash_file_url string `json:"app_hash_file_url,omitempty"`
	App_icon_url string `json:"app_icon_url,omitempty"`
	Dsym_hash_file_id string `json:"dsym_hash_file_id,omitempty"`
}

// ProvisioningProfileResponse represents the ProvisioningProfileResponse schema from the OpenAPI specification
type ProvisioningProfileResponse struct {
	Provisioning_bundle_id string `json:"provisioning_bundle_id,omitempty"` // The bundle identifier associated with the profile.
	Provisioning_profile_name string `json:"provisioning_profile_name,omitempty"` // The name of the provisioning profile.
	Provisioning_profile_type string `json:"provisioning_profile_type"`
	Team_identifier string `json:"team_identifier,omitempty"` // The team identifier.
	Udids []string `json:"udids,omitempty"`
	Appex_profiles []ProvisioningProfileResponse `json:"appex_profiles,omitempty"` // Array of provisioning profiles for any app extensions
}

// CreateStoreSecretResponse represents the CreateStoreSecretResponse schema from the OpenAPI specification
type CreateStoreSecretResponse struct {
	Secret_id string `json:"secret_id,omitempty"` // the secret id for store secret
}

// DataSubjectRightUpdateStatusOperation represents the DataSubjectRightUpdateStatusOperation schema from the OpenAPI specification
type DataSubjectRightUpdateStatusOperation struct {
	Status string `json:"status"` // Operation status
	Participantdata string `json:"participantData,omitempty"` // String field to be used by participant for any intermediate statuses or data they need to save
	Requestid string `json:"requestId"` // Request identifier of the operation
}

// CreateAccessKeyResponse represents the CreateAccessKeyResponse schema from the OpenAPI specification
type CreateAccessKeyResponse struct {
	Accesskey interface{} `json:"accessKey,omitempty"`
}

// Release represents the Release schema from the OpenAPI specification
type Release struct {
	Release string `json:"release"` // Release Id.
}

// GeneralDistributionGroupDetailsResponse represents the GeneralDistributionGroupDetailsResponse schema from the OpenAPI specification
type GeneralDistributionGroupDetailsResponse struct {
	Display_name string `json:"display_name,omitempty"` // The name of the distribution group
	Id string `json:"id"` // The unique ID of the distribution group
	Is_public bool `json:"is_public"` // Whether the distribution group is public
	Name string `json:"name"` // The name of the distribution group used in URLs
	Origin string `json:"origin"` // The creation origin of this distribution group
	Owner_org_id string `json:"owner_org_id,omitempty"` // If distribution group is owned by an org, this is the unique org ID
	Is_shared bool `json:"is_shared,omitempty"` // Whether the distribution group is shared group or not
	Owner_app_id string `json:"owner_app_id,omitempty"` // If distribution group is owned by an app, this is the unique app ID
}

// BugTrackerTokenId represents the BugTrackerTokenId schema from the OpenAPI specification
type BugTrackerTokenId struct {
	Token_id string `json:"token_id,omitempty"`
}

// DeviceDiagnostics represents the DeviceDiagnostics schema from the OpenAPI specification
type DeviceDiagnostics struct {
	Live_update_deployment_key string `json:"live_update_deployment_key,omitempty"` // Identifier of environment that current application release belongs to, deployment key then maps to environment like Production, Staging.
	Os_name string `json:"os_name"` // OS name (example: iOS). The following OS names are standardized (non-exclusive): Android, iOS, macOS, tvOS, Windows.
	Time_zone_offset int `json:"time_zone_offset"` // The offset in minutes from UTC for the device time zone, including daylight savings time.
	Screen_size string `json:"screen_size,omitempty"` // Screen size of the device in pixels (example: 640x480).
	Sdk_name string `json:"sdk_name"` // Name of the SDK. Consists of the name of the SDK and the platform, e.g. "appcenter.ios", "hockeysdk.android".
	Os_api_level int `json:"os_api_level,omitempty"` // API level when applicable like in Android (example: 15).
	Carrier_code string `json:"carrier_code,omitempty"` // Carrier country code (for mobile devices).
	Model string `json:"model,omitempty"` // Device model (example: iPad2,3).
	App_namespace string `json:"app_namespace,omitempty"` // The bundle identifier, package identifier, or namespace, depending on what the individual plattforms use, .e.g com.microsoft.example.
	Sdk_version string `json:"sdk_version"` // Version of the SDK in semver format, e.g. "1.2.0" or "0.12.3-alpha.1".
	Wrapper_sdk_name string `json:"wrapper_sdk_name,omitempty"` // Name of the wrapper SDK. Consists of the name of the SDK and the wrapper platform, e.g. "appcenter.xamarin", "hockeysdk.cordova".
	Wrapper_sdk_version string `json:"wrapper_sdk_version,omitempty"` // Version of the wrapper SDK in semver format. When the SDK is embedding another base SDK (for example Xamarin.Android wraps Android), the Xamarin specific version is populated into this field while sdkVersion refers to the original Android SDK.
	Carrier_country string `json:"carrier_country,omitempty"` // Carrier country.
	Carrier_name string `json:"carrier_name,omitempty"` // Carrier name (for mobile devices).
	Os_version string `json:"os_version"` // OS version (example: 9.3.0).
	Wrapper_runtime_version string `json:"wrapper_runtime_version,omitempty"` // Version of the wrapper technology framework (Xamarin runtime version or ReactNative or Cordova etc...). See wrapper_sdk_name to see if this version refers to Xamarin or ReactNative or other.
	App_build string `json:"app_build"` // The app's build number, e.g. 42.
	Live_update_package_hash string `json:"live_update_package_hash,omitempty"` // Hash of all files (ReactNative or Cordova) deployed to device via LiveUpdate beacon. Helps identify the Release version on device or need to download updates in future.
	Live_update_release_label string `json:"live_update_release_label,omitempty"` // Label that is used to identify application code 'version' released via Live Update beacon running on device
	Os_build string `json:"os_build,omitempty"` // OS build code (example: LMY47X).
	App_version string `json:"app_version"` // Application version name, e.g. 1.1.0
	Locale string `json:"locale"` // Language code (example: en_US).
	Oem_name string `json:"oem_name,omitempty"` // Device manufacturer (example: HTC).
}

// BillingError represents the BillingError schema from the OpenAPI specification
type BillingError struct {
	ErrorField map[string]interface{} `json:"error,omitempty"`
}

// LogTraceDefinition represents the LogTraceDefinition schema from the OpenAPI specification
type LogTraceDefinition struct {
	Expiration string `json:"expiration,omitempty"`
	Install_id string `json:"install_id,omitempty"`
	App_secret string `json:"app_secret"`
}

// CrashGroupOperatingSystem represents the CrashGroupOperatingSystem schema from the OpenAPI specification
type CrashGroupOperatingSystem struct {
	Crash_count int64 `json:"crash_count,omitempty"` // Count of OS.
	Operating_system_name string `json:"operating_system_name,omitempty"` // OS name.
}

// SymbolUploadLocation represents the SymbolUploadLocation schema from the OpenAPI specification
type SymbolUploadLocation struct {
	Uri string `json:"uri"`
}

// ErrorDeleteCounter represents the ErrorDeleteCounter schema from the OpenAPI specification
type ErrorDeleteCounter struct {
	Blobssucceeded int `json:"blobsSucceeded,omitempty"`
	Errorgroupid string `json:"errorGroupId,omitempty"`
	Errorid string `json:"errorId,omitempty"`
	Errorsdeleted int `json:"errorsDeleted,omitempty"`
	Appid string `json:"appId,omitempty"`
	Attachmentsdeleted int `json:"attachmentsDeleted,omitempty"`
	Blobsfailed int `json:"blobsFailed,omitempty"`
}

// DataSubjectRightResponse represents the DataSubjectRightResponse schema from the OpenAPI specification
type DataSubjectRightResponse struct {
	Createdat string `json:"createdAt"` // ISO 8601 format timestamp of when request was created.
	Token string `json:"token"` // Unique request identifier
}

// UpdateDevicesResponse represents the UpdateDevicesResponse schema from the OpenAPI specification
type UpdateDevicesResponse struct {
	Status_url string `json:"status_url"` // URL that can be used to check the status of the update devices operation.
}

// DistributionStore represents the DistributionStore schema from the OpenAPI specification
type DistributionStore struct {
	Publishing_status string `json:"publishing_status,omitempty"` // publishing status of the release in the store.
	TypeField string `json:"type,omitempty"` // type of the distribution store currently stores type can be intune, googleplay or windows.
	Id string `json:"id"` // ID identifying a unique distribution store.
	Name string `json:"name,omitempty"` // A name identifying a unique distribution store.
	Is_latest bool `json:"is_latest,omitempty"` // Is the containing release the latest one in this distribution store.
}

// EventResponseResult represents the EventResponseResult schema from the OpenAPI specification
type EventResponseResult struct {
	Request_id string `json:"request_id"` // Unique request identifier for tracking
}

// ExportConfigurationBlobStorageLinkedSubscription represents the ExportConfigurationBlobStorageLinkedSubscription schema from the OpenAPI specification
type ExportConfigurationBlobStorageLinkedSubscription struct {
	Backfill bool `json:"backfill,omitempty"` // Field to determine if backfilling should occur. The default value is true. If set to false export starts from date and time of config creation.
	Export_entities []string `json:"export_entities,omitempty"`
	Resource_group string `json:"resource_group,omitempty"` // The resource group name on azure
	Resource_name string `json:"resource_name,omitempty"` // The resource name on azure
	TypeField string `json:"type"` // Type of export configuration
	Blob_path_format_kind string `json:"blob_path_format_kind,omitempty"` // The path to the blob when enum set to 'WithoutAppId' is 'year/month/day/hour/minute' and when set to 'WithAppId' is 'appId/year/month/day/hour/minute'
}

// UserDataResponse represents the UserDataResponse schema from the OpenAPI specification
type UserDataResponse struct {
	Display_name string `json:"display_name,omitempty"` // The display name of the user
	Id string `json:"id,omitempty"` // The unique id (UUID) of the user
	Name string `json:"name,omitempty"` // The name of the user
	Avatar_url string `json:"avatar_url,omitempty"` // The url at which the user's avatar can be reached
}

// StatusResponse represents the StatusResponse schema from the OpenAPI specification
type StatusResponse struct {
	Status string `json:"status"`
}

// VersionsDiagnostics represents the VersionsDiagnostics schema from the OpenAPI specification
type VersionsDiagnostics struct {
	Total int64 `json:"total,omitempty"` // the total count of versions
	Versions []map[string]interface{} `json:"versions,omitempty"` // list of version count
}

// PrivateIntuneStoreRequest represents the PrivateIntuneStoreRequest schema from the OpenAPI specification
type PrivateIntuneStoreRequest struct {
	App_category interface{} `json:"app_category,omitempty"`
	Target_audience interface{} `json:"target_audience,omitempty"`
	Tenant_id string `json:"tenant_id,omitempty"` // tenant id of the intune store
}

// StoresBasicDetails represents the StoresBasicDetails schema from the OpenAPI specification
type StoresBasicDetails struct {
	Id string `json:"id,omitempty"` // ID identifying a unique distribution store.
	Name string `json:"name,omitempty"` // A name identifying a unique distribution store.
	Publishing_status string `json:"publishing_status,omitempty"` // publishing status of the release in the store.
	TypeField string `json:"type,omitempty"` // type of the distribution store currently stores type can be intune or googleplay.
}

// TeamUpdateRequest represents the TeamUpdateRequest schema from the OpenAPI specification
type TeamUpdateRequest struct {
	Display_name string `json:"display_name"` // The new display name of the team
}

// InternalUserRequest represents the InternalUserRequest schema from the OpenAPI specification
type InternalUserRequest struct {
	Tester_invitation string `json:"tester_invitation,omitempty"` // The token of the test invitation which lead to signup
	App_invitation string `json:"app_invitation,omitempty"` // The token of the app invitation which lead to signup
	Avatar_url string `json:"avatar_url,omitempty"` // The avatar URL of the user
	Name string `json:"name"` // The unique name that is used to identify the user.
	Organization_invitation string `json:"organization_invitation,omitempty"` // The token of the organization invitation which lead to signup
	Email string `json:"email"` // The email address of the user
	Display_name string `json:"display_name,omitempty"` // The full name of the user. Might for example be first and last name
	Portal_subdomain string `json:"portal_subdomain,omitempty"` // The sub-domain of the portal from which this request was made. Will be used to build the invitation link.
	Password string `json:"password"` // The password of the user. Needs to be at least 8 characters long and contain at least one lower- and one uppercase letter.
}

// ErrorAttachmentText represents the ErrorAttachmentText schema from the OpenAPI specification
type ErrorAttachmentText struct {
	Content string `json:"content,omitempty"`
}

// ProvisioningProfile represents the ProvisioningProfile schema from the OpenAPI specification
type ProvisioningProfile struct {
	Udids []string `json:"udids,omitempty"`
	Application_identifier string `json:"application_identifier"` // The application identifier.
	Expired_at string `json:"expired_at"` // The profile's expiration date in RFC 3339 format, i.e. 2017-07-21T17:32:28Z
	Name string `json:"name"` // The name of the provisioning profile.
	Profile_type string `json:"profile_type"`
	Team_identifier string `json:"team_identifier"` // The team identifier.
}

// GitHubBillingAccount represents the GitHubBillingAccount schema from the OpenAPI specification
type GitHubBillingAccount struct {
	Name string `json:"name"` // The unique name that used to identify the owner
	TypeField string `json:"type"` // The owner type. Can either be 'org' or 'user'
	Display_name string `json:"display_name"` // The account's display name
	Id string `json:"id"` // The unique id (UUID) of the account
}

// AppRepoPostRequest represents the AppRepoPostRequest schema from the OpenAPI specification
type AppRepoPostRequest struct {
	Repo_url string `json:"repo_url"` // The absolute URL of the repository
	Service_connection_id string `json:"service_connection_id,omitempty"` // The id of the service connection stored in customer credential store
	User_id string `json:"user_id"` // The unique id (UUID) of the user who configured the repository
	External_user_id string `json:"external_user_id,omitempty"` // The external user id from the provider
	Installation_id string `json:"installation_id,omitempty"` // Installation id from the provider
	Repo_id string `json:"repo_id,omitempty"` // Repository id from the provider
	Repo_provider string `json:"repo_provider,omitempty"` // The provider of the repository
}

// AudienceTestResult represents the AudienceTestResult schema from the OpenAPI specification
type AudienceTestResult struct {
	Estimated_total_count int64 `json:"estimated_total_count,omitempty"` // Estimated total audience size.
	Custom_properties map[string]interface{} `json:"custom_properties,omitempty"` // Custom properties used in the definition.
	Definition string `json:"definition,omitempty"` // Audience definition in OData format.
	Estimated_count int64 `json:"estimated_count,omitempty"` // Estimated audience size.
}

// JiraSecretDetailsResponse represents the JiraSecretDetailsResponse schema from the OpenAPI specification
type JiraSecretDetailsResponse struct {
	Baseurl string `json:"baseUrl"` // baseUrl to connect to jira instance
	Username string `json:"username"` // username to connect to jira instance
}

// AppleCertificateSecretRequest represents the AppleCertificateSecretRequest schema from the OpenAPI specification
type AppleCertificateSecretRequest struct {
	Credentialtype string `json:"credentialType,omitempty"` // credential type of the shared connection. Values can be credentials|certificate
	Data map[string]interface{} `json:"data,omitempty"` // shared connection details
	Displayname string `json:"displayName,omitempty"` // display name of shared connection
	Servicetype string `json:"serviceType"` // service type of shared connection can be apple|gitlab|googleplay|jira|applecertificate
}

// MemberDevicesInfoResponse represents the MemberDevicesInfoResponse schema from the OpenAPI specification
type MemberDevicesInfoResponse struct {
	Registered_at string `json:"registered_at,omitempty"` // Timestamp of when the device was registered in ISO format.
	Display_name string `json:"display_name,omitempty"` // The full name of the user. Might for example be first and last name
	Status string `json:"status"` // The provisioning status of the device.
	Can_change_password bool `json:"can_change_password,omitempty"` // User is required to send an old password in order to change the password.
	Device_name string `json:"device_name"` // The device description, in the format "iPhone 7 Plus (A1784)"
	Full_device_name string `json:"full_device_name,omitempty"` // A combination of the device model name and the owner name.
	Model string `json:"model"` // The model identifier of the device, in the format iDeviceM,N
	Serial string `json:"serial,omitempty"` // The device's serial number. Always empty or undefined at present.
	Avatar_url string `json:"avatar_url,omitempty"` // The avatar URL of the user
	Email string `json:"email"` // The email address of the user
	Imei string `json:"imei,omitempty"` // The device's International Mobile Equipment Identity number. Always empty or undefined at present.
	Owner_id string `json:"owner_id,omitempty"` // The user ID of the device owner.
	Os_version string `json:"os_version"` // The last known OS version running on the device
	Id string `json:"id"` // The unique id (UUID) of the user
	Name string `json:"name,omitempty"` // The unique name that is used to identify the user.
	Udid string `json:"udid"` // The Unique Device IDentifier of the device
	Os_build string `json:"os_build"` // The last known OS version running on the device
	Invite_pending bool `json:"invite_pending,omitempty"` // Whether the has accepted the invite. Available when an invite is pending, and the value will be "true".
}

// AlertWebhookListResult represents the AlertWebhookListResult schema from the OpenAPI specification
type AlertWebhookListResult struct {
	Values []map[string]interface{} `json:"values"`
}

// XcodeBranchConfigurationProperties represents the XcodeBranchConfigurationProperties schema from the OpenAPI specification
type XcodeBranchConfigurationProperties struct {
	Targettoarchive string `json:"targetToArchive,omitempty"` // The target id of the selected scheme to archive
	Certificatefilename string `json:"certificateFilename,omitempty"`
	Scheme string `json:"scheme,omitempty"`
	Certificateuploadid string `json:"certificateUploadId,omitempty"`
	Forcelegacybuildsystem bool `json:"forceLegacyBuildSystem,omitempty"` // Setting this to true forces the build to use Xcode legacy build system. Otherwise, the setting from workspace settings is used. By default new build system is used if workspace setting is not committed to the repository. Only used for iOS React Native app, with Xcode 10.
	Xcodeprojectsha string `json:"xcodeProjectSha,omitempty"` // The selected pbxproject hash to the repositroy
	Podfilepath string `json:"podfilePath,omitempty"` // Path to CococaPods file, if present
	Provisioningprofilefileid string `json:"provisioningProfileFileId,omitempty"`
	Provisioningprofileuploadid string `json:"provisioningProfileUploadId,omitempty"`
	Xcodeversion string `json:"xcodeVersion,omitempty"` // Xcode version used to build. Available versions can be found in "/xcode_versions" API. Default is latest stable version, at the time when the configuration is set.
	Certificateencoded string `json:"certificateEncoded,omitempty"`
	Archiveconfiguration string `json:"archiveConfiguration,omitempty"` // The build configuration of the target to archive
	Provisioningprofilefilename string `json:"provisioningProfileFilename,omitempty"`
	Appextensionprovisioningprofilefiles []interface{} `json:"appExtensionProvisioningProfileFiles,omitempty"`
	Certificatepassword string `json:"certificatePassword,omitempty"`
	Automaticsigning bool `json:"automaticSigning,omitempty"`
	Cartfilepath string `json:"cartfilePath,omitempty"` // Path to Carthage file, if present
	Provisioningprofileencoded string `json:"provisioningProfileEncoded,omitempty"`
	Teamid string `json:"teamId,omitempty"`
	Projectorworkspacepath string `json:"projectOrWorkspacePath,omitempty"` // Xcode project/workspace path
	Certificatefileid string `json:"certificateFileId,omitempty"`
}

// CrashAttachmentLocation represents the CrashAttachmentLocation schema from the OpenAPI specification
type CrashAttachmentLocation struct {
	Uri string `json:"uri"`
}

// AddAADGroupResponse represents the AddAADGroupResponse schema from the OpenAPI specification
type AddAADGroupResponse struct {
	Id string `json:"id,omitempty"` // The unique ID (UUID) of the aad group
}

// ErrorDownloadLink represents the ErrorDownloadLink schema from the OpenAPI specification
type ErrorDownloadLink struct {
	Link string `json:"link"`
}

// BranchYamlConfiguration represents the BranchYamlConfiguration schema from the OpenAPI specification
type BranchYamlConfiguration struct {
	Yaml string `json:"yaml,omitempty"` // Azure Pipelines YAML file
}

// ErrorGroups represents the ErrorGroups schema from the OpenAPI specification
type ErrorGroups struct {
	Errorgroups []map[string]interface{} `json:"errorGroups,omitempty"`
	Nextlink string `json:"nextLink,omitempty"`
}

// OSes represents the OSes schema from the OpenAPI specification
type OSes struct {
	Oses []interface{} `json:"oses,omitempty"`
	Total int64 `json:"total,omitempty"`
}

// PrivateAppleCredentialsSecretResponse represents the PrivateAppleCredentialsSecretResponse schema from the OpenAPI specification
type PrivateAppleCredentialsSecretResponse struct {
	Servicetype string `json:"serviceType"` // service type of shared connection can be apple|gitlab|googleplay|jira|applecertificate
	Displayname string `json:"displayName,omitempty"` // display name of shared connection
	Id string `json:"id"` // id of the shared connection
	Is2fa bool `json:"is2FA,omitempty"` // if the account is a 2FA account or not
	Isvalid bool `json:"isValid,omitempty"` // whether the credentials are valid or not
}

// TestRun represents the TestRun schema from the OpenAPI specification
type TestRun struct {
	Id string `json:"id,omitempty"` // The unique id of the test upload
	Stats map[string]interface{} `json:"stats,omitempty"` // Summary single test run on Xamarin Test Cloud
	Status string `json:"status,omitempty"` // Deprecated. Use resultStatus instead.
	Testtype string `json:"testType,omitempty"` // The name of the test framework used to run this test
	Appversion string `json:"appVersion,omitempty"` // The compiled version of the app binary
	Description string `json:"description,omitempty"` // Human readable explanation of the current test status
	Testseries string `json:"testSeries,omitempty"` // The name of the test series with which this test upload is associated
	Platform string `json:"platform,omitempty"` // The device platform targeted by the test. Possible values are 'ios' or 'android'
	Date string `json:"date,omitempty"` // The date and time the test was uploaded
	Resultstatus string `json:"resultStatus,omitempty"` // The passed/failed state
	Runstatus string `json:"runStatus,omitempty"` // The current status of the test run, in relation to the various phases
	State string `json:"state,omitempty"` // Deprecated. Use runStatus instead.
}

// BuildConcurrencyResponse represents the BuildConcurrencyResponse schema from the OpenAPI specification
type BuildConcurrencyResponse struct {
	Committed_quantity float64 `json:"committed_quantity,omitempty"` // The number of pipelines committed, which can be equal or greater than the number from the billing plan
	Quantity float64 `json:"quantity,omitempty"` // The number of pipelines set by the billing plan
}

// V2SymbolUpdateInfo represents the V2SymbolUpdateInfo schema from the OpenAPI specification
type V2SymbolUpdateInfo struct {
	App_id string `json:"app_id"` // application id
	Status string `json:"status"` // symbol upload status
	Symbol_id string `json:"symbol_id"` // UUID of the symbol
}

// SecretDetails represents the SecretDetails schema from the OpenAPI specification
type SecretDetails struct {
	Refresh_token_expiry string `json:"refresh_token_expiry,omitempty"` // the expiry of refresh token
	Id_token string `json:"id_token,omitempty"` // the id token of user
	Refresh_token string `json:"refresh_token,omitempty"` // the refresh token for user
}

// UserAppPermissionsUpdateRequest represents the UserAppPermissionsUpdateRequest schema from the OpenAPI specification
type UserAppPermissionsUpdateRequest struct {
	Permissions []string `json:"permissions"` // The permissions the user has for the app
}

// LegacyCollaborator represents the LegacyCollaborator schema from the OpenAPI specification
type LegacyCollaborator struct {
	Iscurrentaccount bool `json:"isCurrentAccount,omitempty"` // Is current collaborator the same as current account.
	Permission string `json:"permission,omitempty"` // Which permission does current account has.
}

// BuildConfiguration represents the BuildConfiguration schema from the OpenAPI specification
type BuildConfiguration struct {
	Name string `json:"name"` // Name of build configuration (the same as a build type name)
	Signingconfig interface{} `json:"signingConfig,omitempty"` // Android signing config. Null if not specified
}

// ApiTokensPrivateCreateResponse represents the ApiTokensPrivateCreateResponse schema from the OpenAPI specification
type ApiTokensPrivateCreateResponse struct {
	Description string `json:"description,omitempty"` // The description of the token
	Encrypted_token string `json:"encrypted_token,omitempty"` // The encrypted value of a token. This value will only be returned for token of type in_app_update.
	Id string `json:"id"` // The unique id (UUID) of the api token
	Scope []string `json:"scope,omitempty"` // The scope for this token.
	Api_token string `json:"api_token"` // The api token generated will not be accessible again
	Created_at string `json:"created_at"` // The creation time
}

// StringProperty represents the StringProperty schema from the OpenAPI specification
type StringProperty struct {
	Name string `json:"name"`
	TypeField string `json:"type"`
}

// PatchReleaseUploadStatusRequest represents the PatchReleaseUploadStatusRequest schema from the OpenAPI specification
type PatchReleaseUploadStatusRequest struct {
	Upload_status string `json:"upload_status"` // The new status of the release upload
}

// PrivateAppleCertificateSecretResponse represents the PrivateAppleCertificateSecretResponse schema from the OpenAPI specification
type PrivateAppleCertificateSecretResponse struct {
	Is2fa bool `json:"is2FA,omitempty"` // if the account is a 2FA account or not
	Isvalid bool `json:"isValid,omitempty"` // whether the credentials are valid or not
	Servicetype string `json:"serviceType"` // service type of shared connection can be apple|gitlab|googleplay|jira|applecertificate
	Displayname string `json:"displayName,omitempty"` // display name of shared connection
	Id string `json:"id"` // id of the shared connection
}

// BuildTimeline represents the BuildTimeline schema from the OpenAPI specification
type BuildTimeline struct {
	Id string `json:"id,omitempty"`
	Records []interface{} `json:"records,omitempty"`
}

// TestRunState represents the TestRunState schema from the OpenAPI specification
type TestRunState struct {
	Message []string `json:"message,omitempty"` // Multi-line message that describes the status
	Wait_time int `json:"wait_time,omitempty"` // Time (in seconds) that the client should wait for before checking the status again
	Exit_code int `json:"exit_code,omitempty"` // The exit code that the client should use when exiting. Used for indicating status to the caller of the client. 0: test run completes with no failing tests 1: test run completes with at least one failing test 2: test run failed to complete. Status for test run is unknown
}

// CrashGroupsContainer represents the CrashGroupsContainer schema from the OpenAPI specification
type CrashGroupsContainer struct {
	Crash_groups []map[string]interface{} `json:"crash_groups"`
	Limited_result_set bool `json:"limited_result_set"`
	Continuation_token string `json:"continuation_token,omitempty"` // Cassandra request continuation token. The token is used for pagination.
}

// LogFlowGenericLogContainer represents the LogFlowGenericLogContainer schema from the OpenAPI specification
type LogFlowGenericLogContainer struct {
	Exceeded_max_limit bool `json:"exceeded_max_limit,omitempty"` // indicates if the number of available logs are more than the max allowed return limit(100).
	Last_received_log_timestamp string `json:"last_received_log_timestamp,omitempty"` // the timestamp of the last log received. This value can be used as the start time parameter in the consecutive API call.
	Logs []map[string]interface{} `json:"logs"` // the list of logs
}

// StoreDestinationDetails represents the StoreDestinationDetails schema from the OpenAPI specification
type StoreDestinationDetails struct {
	Appid string `json:"appId,omitempty"` // app id of application.
	Dest_publish_id string `json:"dest_publish_id,omitempty"` // destination ID identifying a unique id in distribution store.
	Store_type string `json:"store_type,omitempty"` // type of store.
}

// Deployment represents the Deployment schema from the OpenAPI specification
type Deployment struct {
	Key string `json:"key,omitempty"`
	Latest_release interface{} `json:"latest_release,omitempty"`
	Name string `json:"name"`
}

// StartSessionLog represents the StartSessionLog schema from the OpenAPI specification
type StartSessionLog struct {
	Timestamp string `json:"timestamp"` // Log creation timestamp.
	TypeField string `json:"type"` // Log type.
	Device map[string]interface{} `json:"device"` // Device characteristics.
	Install_id string `json:"install_id"` // Install ID.
}

// AutoProvisioningConfigResponse represents the AutoProvisioningConfigResponse schema from the OpenAPI specification
type AutoProvisioningConfigResponse struct {
	App_id string `json:"app_id,omitempty"` // The identifier of the App.
	Apple_developer_account_key string `json:"apple_developer_account_key,omitempty"` // A key to a secret in customer-credential-store. apple_developer_account refers to the user's developer account that is used to log into https://developer.apple.com. Normally the user's email.
	Apple_distribution_certificate_key string `json:"apple_distribution_certificate_key,omitempty"` // A key to a secret in customer-credential-store. distribution_certificate refers to the cusomer's certificate (that holds the private key) that will be used to sign the app.
	Destination_id string `json:"destination_id,omitempty"` // The identifier of the destination.
	Id float64 `json:"id,omitempty"` // The identifier of the config.
	Allow_auto_provisioning bool `json:"allow_auto_provisioning,omitempty"` // When *true* enables auto provisioning
}

// UserSettingResponse represents the UserSettingResponse schema from the OpenAPI specification
type UserSettingResponse struct {
	Marketing_opt_in string `json:"marketing_opt_in,omitempty"` // The marketing opt-in setting
}

// BuildIssue represents the BuildIssue schema from the OpenAPI specification
type BuildIssue struct {
	TypeField string `json:"type,omitempty"`
	Category string `json:"category,omitempty"`
	Message string `json:"message,omitempty"`
}

// DistributionGroupResponse represents the DistributionGroupResponse schema from the OpenAPI specification
type DistributionGroupResponse struct {
	Is_public bool `json:"is_public"` // Whether the distribution group is public
	Name string `json:"name"` // The name of the distribution group used in URLs
	Origin string `json:"origin"` // The creation origin of this distribution group
	Display_name string `json:"display_name,omitempty"` // The name of the distribution group
	Id string `json:"id"` // The unique ID of the distribution group
}

// ManagementReleaseDetailsResponse represents the ManagementReleaseDetailsResponse schema from the OpenAPI specification
type ManagementReleaseDetailsResponse struct {
	Deletedat string `json:"deletedAt,omitempty"` // UTC time the release was created in ISO 8601 format.
	Distinctid int `json:"distinctId,omitempty"` // ID identifying this unique release.
	Enabled bool `json:"enabled,omitempty"` // This value determines the whether a release currently is enabled or disabled.
	Origin string `json:"origin,omitempty"` // The release's origin
	Sortversion string `json:"sortVersion,omitempty"` // The release's sortVersion.
	Version string `json:"version,omitempty"` // The release's short version.<br> For iOS: CFBundleShortVersionString from info.plist.<br> For Android: android:versionName from AppManifest.xml.
	Buildversion string `json:"buildVersion,omitempty"` // The release's buildVersion.<br> For iOS: CFBundleVersion from info.plist.<br> For Android: android:versionCode from AppManifest.xml.
	Createdat string `json:"createdAt,omitempty"` // UTC time the release was created in ISO 8601 format.
}

// UpdateIdentityProvidersRequest represents the UpdateIdentityProvidersRequest schema from the OpenAPI specification
type UpdateIdentityProvidersRequest struct {
}

// ApiTokensPrivateCreateRequest represents the ApiTokensPrivateCreateRequest schema from the OpenAPI specification
type ApiTokensPrivateCreateRequest struct {
	Token_type string `json:"token_type,omitempty"` // The token's type (default "public") public: managed by the user in_app_update: special token for in-app update scenario buid: dedicated for CI usage for now session: for CLI session management tester_app: used for tester mobile app
	Description string `json:"description,omitempty"` // The description of the token
	Principal_id string `json:"principal_id"` // The principal ID assigned to the API token
	Principal_type string `json:"principal_type"` // The principal type assigned to the API token
	Scope []string `json:"scope,omitempty"` // The scope for this token (default "all").
}

// AppBuildFeature represents the AppBuildFeature schema from the OpenAPI specification
type AppBuildFeature struct {
	Name string `json:"name,omitempty"`
	Value bool `json:"value,omitempty"`
}

// LogFlowGenericLog represents the LogFlowGenericLog schema from the OpenAPI specification
type LogFlowGenericLog struct {
	Account_id string `json:"account_id,omitempty"` // Account ID of the authenticated user.
	Properties map[string]interface{} `json:"properties,omitempty"` // event specific properties.
	TypeField string `json:"type"` // Log type.
	Device map[string]interface{} `json:"device"` // Device characteristics.
	Session_id string `json:"session_id,omitempty"` // Session ID.
	Auth_provider string `json:"auth_provider,omitempty"` // Auth service provider.
	Event_name string `json:"event_name,omitempty"` // Event name.
	Install_id string `json:"install_id"` // Install ID.
	Event_id string `json:"event_id,omitempty"` // Event ID.
	Message_id string `json:"message_id,omitempty"` // Message ID.
	Timestamp string `json:"timestamp"` // Log creation timestamp.
}

// CrashCounts represents the CrashCounts schema from the OpenAPI specification
type CrashCounts struct {
	Count int64 `json:"count,omitempty"` // Total crash count.
	Crashes []interface{} `json:"crashes,omitempty"` // The total crash count for day.
}

// IosAppExtensionInfo represents the IosAppExtensionInfo schema from the OpenAPI specification
type IosAppExtensionInfo struct {
	Name string `json:"name"` // App extension name
	Targetbundleidentifier string `json:"targetBundleIdentifier"` // App extension bundle identifier
}

// SymbolStatusResponse represents the SymbolStatusResponse schema from the OpenAPI specification
type SymbolStatusResponse struct {
	Status string `json:"status"` // Whether the symbol is ignored.
	Symbol_id string `json:"symbol_id"` // The unique id for this symbol (uuid)
	App_id string `json:"app_id"` // The application that this symbol belongs to
}

// DistributionGroupWithoutIsLatest represents the DistributionGroupWithoutIsLatest schema from the OpenAPI specification
type DistributionGroupWithoutIsLatest struct {
	Id string `json:"id"` // ID identifying a unique distribution group.
	Name string `json:"name,omitempty"` // A name identifying a unique distribution group.
}

// TestCloudFileHashResponse represents the TestCloudFileHashResponse schema from the OpenAPI specification
type TestCloudFileHashResponse struct {
	Filetype string `json:"fileType"` // Type of the file
	Relativepath string `json:"relativePath,omitempty"` // Relative path of the file
	Uploadstatus map[string]interface{} `json:"uploadStatus"` // Status of the upload
	Checksum string `json:"checksum"` // SHA256 hash of the file
}

// SuccessResponse represents the SuccessResponse schema from the OpenAPI specification
type SuccessResponse struct {
	Message string `json:"message"`
}

// PageLogDiagnostics represents the PageLogDiagnostics schema from the OpenAPI specification
type PageLogDiagnostics struct {
	Install_id string `json:"install_id"` // Install ID.
	Timestamp string `json:"timestamp"` // Log creation timestamp.
	TypeField string `json:"type"` // Log type.
	Device map[string]interface{} `json:"device"` // Device characteristics.
	Properties map[string]interface{} `json:"properties,omitempty"` // Additional key/value pair parameters.
}

// UpdateDevicesRequest represents the UpdateDevicesRequest schema from the OpenAPI specification
type UpdateDevicesRequest struct {
	Devices []string `json:"devices,omitempty"` // Array of device UDID's to be published to the Apple Developer account.
	P12_base64 string `json:"p12_base64,omitempty"` // The certificate to use for resigning the application with the updated provisioning profiles.
	P12_password string `json:"p12_password,omitempty"` // The password certificate if one is needed.
	Publish_all_devices bool `json:"publish_all_devices,omitempty"` // When set to true, all unprovisioned devices will be published to the Apple Developer account. When false, only the provided devices will be published to the Apple Developer account.
	Destinations []map[string]interface{} `json:"destinations,omitempty"` // Array of distribution groups that the devices should be provisioned from.
	Account_service_connection_id string `json:"account_service_connection_id,omitempty"` // The service_connection_id of the stored Apple credentials instead of username, password.
	Password string `json:"password,omitempty"` // The password for the Apple Developer account to publish the devices to.
	P12_service_connection_id string `json:"p12_service_connection_id,omitempty"` // The service_connection_id of the stored Apple certificate instead of p12_base64 value.
	Release_id float64 `json:"release_id,omitempty"` // When provided, will update the provided release with the new set of devices. By default the latest release of the distribution group is used when this property is omitted. If `release_id` is passed in the path, there is no need to pass in the body as well.
	Username string `json:"username,omitempty"` // The username for the Apple Developer account to publish the devices to.
}

// MSAUser represents the MSAUser schema from the OpenAPI specification
type MSAUser struct {
	Hexcid string `json:"hexcid"` // the user's CID
}

// LogFlowStringProperty represents the LogFlowStringProperty schema from the OpenAPI specification
type LogFlowStringProperty struct {
	TypeField string `json:"type"`
	Name string `json:"name"`
}

// TesterAppResponse represents the TesterAppResponse schema from the OpenAPI specification
type TesterAppResponse struct {
	Display_name string `json:"display_name"` // The display name of the app
	Icon_url string `json:"icon_url,omitempty"` // The string representation of the URL pointing to the app's icon
	Id string `json:"id"` // The unique ID (UUID) of the app
	Name string `json:"name"` // The name of the app used in URLs
	Owner interface{} `json:"owner"` // The information about the app's owner
	Icon_source string `json:"icon_source,omitempty"` // The string representation of the source of the app's icon
	Release_type string `json:"release_type,omitempty"` // A one-word descriptive release-type value that starts with a capital letter but is otherwise lowercase
	Description string `json:"description,omitempty"` // The description of the app
	Os string `json:"os"` // The OS the app will be running on
	Permissions []string `json:"permissions,omitempty"` // The permissions associated with the app
	Microsoft_internal bool `json:"microsoft_internal,omitempty"` // it indicates if the app is microsoft internal
}

// ErrorAppBuilds represents the ErrorAppBuilds schema from the OpenAPI specification
type ErrorAppBuilds struct {
	Appbuilds []string `json:"appBuilds,omitempty"`
}

// ResignAttemptResponse represents the ResignAttemptResponse schema from the OpenAPI specification
type ResignAttemptResponse struct {
	Appid string `json:"appId"` // App ID that the resign operation is being performed against.
	Contextid string `json:"contextId"` // Context ID for the resigning operation.
	Destinations []map[string]interface{} `json:"destinations,omitempty"` // List of destinations that the resign operation is being performed against.
	Resignid string `json:"resignId"` // ID of the resign operation.
	Userid string `json:"userId"` // ID of the user performing the resign operaiton.
	Starttime float64 `json:"startTime"` // The time that the resign operation was started.
	Errorcode string `json:"errorCode,omitempty"` // Error code associated with the exception.
	Errormessage string `json:"errorMessage,omitempty"` // Error message associated with the exception.
	Status string `json:"status"` // The status of the resigning operation.
	Originalreleaseid float64 `json:"originalReleaseId"` // ID of the release which is being resigned.
}

// AppDistributionGroupUsersRequest represents the AppDistributionGroupUsersRequest schema from the OpenAPI specification
type AppDistributionGroupUsersRequest struct {
	Member_ids []string `json:"member_ids,omitempty"`
}

// ServiceBillingPlans represents the ServiceBillingPlans schema from the OpenAPI specification
type ServiceBillingPlans struct {
	Canselecttrialplan bool `json:"canSelectTrialPlan,omitempty"` // Can customer select trial plan for that service (if it exists)?
	Currentbillingperiod interface{} `json:"currentBillingPeriod,omitempty"` // Billing plans for a given period
	Lasttrialplanexpirationtime string `json:"lastTrialPlanExpirationTime,omitempty"` // Expiration time of the last selected trial plan. Will be null if trial plan was not used.
}

// BuildPatch represents the BuildPatch schema from the OpenAPI specification
type BuildPatch struct {
	Status string `json:"status,omitempty"` // The build status; used to cancel builds
}

// TeamUserResponse represents the TeamUserResponse schema from the OpenAPI specification
type TeamUserResponse struct {
	Display_name string `json:"display_name"` // The full name of the user. Might for example be first and last name
	Email string `json:"email"` // The email address of the user
	Name string `json:"name"` // The unique name that is used to identify the user.
	Role interface{} `json:"role"` // The role of the user has within the team
}

// ReleaseWithDistributionGroup represents the ReleaseWithDistributionGroup schema from the OpenAPI specification
type ReleaseWithDistributionGroup struct {
	Release string `json:"release"` // Release Id.
}

// AzureSubscriptionAddRequest represents the AzureSubscriptionAddRequest schema from the OpenAPI specification
type AzureSubscriptionAddRequest struct {
	Is_billing bool `json:"is_billing,omitempty"` // If the subscription is used for billing
	Subscription_id string `json:"subscription_id"` // The azure subscription id
	Subscription_name string `json:"subscription_name"` // The name of the azure subscription
	Tenant_id string `json:"tenant_id"` // The tenant id of the azure subscription belongs to
}

// EventDeviceCount represents the EventDeviceCount schema from the OpenAPI specification
type EventDeviceCount struct {
	Devices_count []interface{} `json:"devices_count,omitempty"`
	Previous_total_devices_with_event int64 `json:"previous_total_devices_with_event,omitempty"`
	Total_devices int64 `json:"total_devices,omitempty"`
	Total_devices_with_event int64 `json:"total_devices_with_event,omitempty"`
}

// AppleMappingRequest represents the AppleMappingRequest schema from the OpenAPI specification
type AppleMappingRequest struct {
	Service_connection_id string `json:"service_connection_id"` // Id for the shared service connection. In case of Apple AppStore, this connection will be used to create and connect to the Apple AppStore in Mobile Center.
	Team_identifier string `json:"team_identifier"` // ID of the Team associated with the app in apple store
	Apple_id string `json:"apple_id,omitempty"` // ID of the apple application in apple store, takes precedence over bundle_identifier when both are provided
	Bundle_identifier string `json:"bundle_identifier,omitempty"` // Bundle Identifier of the apple package
}

// AlertingBugtrackerSettings represents the AlertingBugtrackerSettings schema from the OpenAPI specification
type AlertingBugtrackerSettings struct {
	TypeField string `json:"type"` // type of bugtracker
	Callback_url string `json:"callback_url,omitempty"`
	Owner_name string `json:"owner_name"`
}

// GetOrCreateRepositoryProviderMappingRequest represents the GetOrCreateRepositoryProviderMappingRequest schema from the OpenAPI specification
type GetOrCreateRepositoryProviderMappingRequest struct {
	External_account_name string `json:"external_account_name,omitempty"` // The account name given by the external provider. If provided, create an organization and the mapping. If not, create mapping with user.
}

// LinkAADtoUserRequest represents the LinkAADtoUserRequest schema from the OpenAPI specification
type LinkAADtoUserRequest struct {
	Aad_tenant_ids []string `json:"aad_tenant_ids"` // An array of AAD tenant data needed to link the user to the tenants
	Role string `json:"role,omitempty"` // The role of the user to be added
}

// AppleMappingResponse represents the AppleMappingResponse schema from the OpenAPI specification
type AppleMappingResponse struct {
	App_id string `json:"app_id,omitempty"` // ID of the apple application in Mobile Center
	Apple_id string `json:"apple_id,omitempty"` // ID of the apple application in apple store
	Service_connection_id string `json:"service_connection_id,omitempty"` // Id for the shared service connection. In case of Apple AppStore, this connection will be used to create and connect to the Apple AppStore in Mobile Center.
	Team_identifier string `json:"team_identifier,omitempty"` // ID of the Team associated with the app in apple store
}

// CustomProperty represents the CustomProperty schema from the OpenAPI specification
type CustomProperty struct {
	Name string `json:"name"`
	TypeField string `json:"type"`
}

// DeviceCpu represents the DeviceCpu schema from the OpenAPI specification
type DeviceCpu struct {
	Core string `json:"core,omitempty"`
	Frequency string `json:"frequency,omitempty"`
	Text string `json:"text,omitempty"`
}

// SymbolUploadEndRequest represents the SymbolUploadEndRequest schema from the OpenAPI specification
type SymbolUploadEndRequest struct {
	Status string `json:"status"` // The desired operation for the symbol upload
}

// DistributionGroupUserGetResponse represents the DistributionGroupUserGetResponse schema from the OpenAPI specification
type DistributionGroupUserGetResponse struct {
	Name string `json:"name,omitempty"` // The unique name that is used to identify the user.
	Avatar_url string `json:"avatar_url,omitempty"` // The avatar URL of the user
	Can_change_password bool `json:"can_change_password,omitempty"` // User is required to send an old password in order to change the password.
	Display_name string `json:"display_name,omitempty"` // The full name of the user. Might for example be first and last name
	Email string `json:"email"` // The email address of the user
	Id string `json:"id,omitempty"` // The unique id (UUID) of the user
	Invite_pending bool `json:"invite_pending,omitempty"` // Whether the has accepted the invite. Available when an invite is pending, and the value will be "true".
}

// AlertOperationResult represents the AlertOperationResult schema from the OpenAPI specification
type AlertOperationResult struct {
	Request_id string `json:"request_id"` // Unique request identifier for tracking
}

// IntuneStoreRequest represents the IntuneStoreRequest schema from the OpenAPI specification
type IntuneStoreRequest struct {
	App_category interface{} `json:"app_category,omitempty"`
	Secret_json interface{} `json:"secret_json,omitempty"`
	Target_audience interface{} `json:"target_audience,omitempty"`
	Tenant_id string `json:"tenant_id,omitempty"` // tenant id of the intune store
}

// HandledError represents the HandledError schema from the OpenAPI specification
type HandledError struct {
	Country string `json:"country,omitempty"`
	Errorid string `json:"errorId,omitempty"`
	Hasbreadcrumbs bool `json:"hasBreadcrumbs,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Hasattachments bool `json:"hasAttachments,omitempty"`
	Osversion string `json:"osVersion,omitempty"`
	Userid string `json:"userId,omitempty"`
	Language string `json:"language,omitempty"`
	Devicename string `json:"deviceName,omitempty"`
	Ostype string `json:"osType,omitempty"`
}

// IntuneCategories represents the IntuneCategories schema from the OpenAPI specification
type IntuneCategories struct {
	Value []interface{} `json:"value,omitempty"` // categories for intune app
	Odata_context string `json:"odata.context,omitempty"` // context
}

// CrashGroupCarriers represents the CrashGroupCarriers schema from the OpenAPI specification
type CrashGroupCarriers struct {
	Crash_count int64 `json:"crash_count,omitempty"`
	Carriers []map[string]interface{} `json:"carriers,omitempty"`
}

// NotifyReleasesContainer represents the NotifyReleasesContainer schema from the OpenAPI specification
type NotifyReleasesContainer struct {
	Releases []interface{} `json:"releases"`
}

// AddUserAsRoleRequest represents the AddUserAsRoleRequest schema from the OpenAPI specification
type AddUserAsRoleRequest struct {
	Role string `json:"role,omitempty"` // The role of the user to be added
}

// DistributionGroupUserRequest represents the DistributionGroupUserRequest schema from the OpenAPI specification
type DistributionGroupUserRequest struct {
	User_emails []string `json:"user_emails,omitempty"` // The list of emails of the users
}

// UserProfileResponseManagement represents the UserProfileResponseManagement schema from the OpenAPI specification
type UserProfileResponseManagement struct {
	Permissions []string `json:"permissions,omitempty"` // The permissions the user has for the app
	Avatar_url string `json:"avatar_url,omitempty"` // The avatar URL of the user
	Can_change_password bool `json:"can_change_password,omitempty"` // User is required to send an old password in order to change the password.
	Display_name string `json:"display_name"` // The full name of the user. Might for example be first and last name
	Email string `json:"email"` // The email address of the user
	Id string `json:"id"` // The unique id (UUID) of the user
	Name string `json:"name"` // The unique name that is used to identify the user.
	Origin string `json:"origin"` // The creation origin of this user
	Admin_role string `json:"admin_role,omitempty"` // The new admin_role
	Feature_flags []string `json:"feature_flags,omitempty"` // The feature flags that are enabled for this app
	Settings interface{} `json:"settings,omitempty"` // The user's settings
	Identity_providers []interface{} `json:"identity_providers,omitempty"` // The identity providers associated with the user's account
	Updated_at string `json:"updated_at,omitempty"` // The date when the app was last updated
	Verified bool `json:"verified,omitempty"` // A boolean flag that indicates if the user is already verified
}

// AppMembership represents the AppMembership schema from the OpenAPI specification
type AppMembership struct {
	Origin string `json:"origin,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
	Source_id string `json:"source_id,omitempty"`
	Source_type string `json:"source_type,omitempty"`
	User_id string `json:"user_id,omitempty"`
	App_id string `json:"app_id,omitempty"`
}

// AppRepoResponse represents the AppRepoResponse schema from the OpenAPI specification
type AppRepoResponse struct {
	Repo_provider string `json:"repo_provider,omitempty"` // The provider of the repository
	Repo_url string `json:"repo_url"` // The absolute URL of the repository
	Service_connection_id string `json:"service_connection_id,omitempty"` // The id of the service connection stored in customer credential store
	Id string `json:"id"` // The unique id (UUID) of the repository integration
	User_id string `json:"user_id"` // The unique id (UUID) of the user who configured the repository
	External_user_id string `json:"external_user_id,omitempty"` // User id from the provider
	App_id string `json:"app_id"` // The unique id (UUID) of the app that this repository integration belongs to
	Repo_id string `json:"repo_id,omitempty"` // Repository id from the provider
	Installation_id string `json:"installation_id,omitempty"` // Installation id from the provider
}

// AppPatchRequest represents the AppPatchRequest schema from the OpenAPI specification
type AppPatchRequest struct {
	Description string `json:"description,omitempty"` // A short text describing the app
	Display_name string `json:"display_name,omitempty"` // The display name of the app
	Icon_asset_id string `json:"icon_asset_id,omitempty"` // The uuid for the icon's asset id from ACFUS
	Icon_url string `json:"icon_url,omitempty"` // The string representation of the URL pointing to the app's icon
	Name string `json:"name,omitempty"` // The name of the app used in URLs
	Release_type string `json:"release_type,omitempty"` // A one-word descriptive release type value that starts with a capital letter but is otherwise lowercase
}

// AppVersion represents the AppVersion schema from the OpenAPI specification
type AppVersion struct {
	App_version_id string `json:"app_version_id"`
	Build_number string `json:"build_number,omitempty"`
	Display_name string `json:"display_name"`
	App_id string `json:"app_id"`
	App_version string `json:"app_version"`
}

// HanledErrorLogDiagnostics represents the HanledErrorLogDiagnostics schema from the OpenAPI specification
type HanledErrorLogDiagnostics struct {
	Install_id string `json:"install_id"` // Install ID.
	Timestamp string `json:"timestamp"` // Log creation timestamp.
	TypeField string `json:"type"` // Log type.
	Device map[string]interface{} `json:"device"` // Device characteristics.
}

// AppAddRequest represents the AppAddRequest schema from the OpenAPI specification
type AppAddRequest struct {
	Name string `json:"name"` // The name of the app to be added to the distribution group
}

// DateTimeDownloadReleaseCounts represents the DateTimeDownloadReleaseCounts schema from the OpenAPI specification
type DateTimeDownloadReleaseCounts struct {
	Counts []map[string]interface{} `json:"counts,omitempty"` // Release counts per day.
	Total int64 `json:"total,omitempty"`
	Unique int64 `json:"unique,omitempty"`
}

// GenericLogContainerDiagnostics represents the GenericLogContainerDiagnostics schema from the OpenAPI specification
type GenericLogContainerDiagnostics struct {
	Exceeded_max_limit bool `json:"exceeded_max_limit,omitempty"` // indicates if the number of available logs are more than the max allowed return limit(100).
	Last_received_log_timestamp string `json:"last_received_log_timestamp,omitempty"` // the timestamp of the last log received. This value can be used as the start time parameter in the consecutive API call.
	Logs []map[string]interface{} `json:"logs"` // the list of logs
}

// BasicReleaseDetailsResponse represents the BasicReleaseDetailsResponse schema from the OpenAPI specification
type BasicReleaseDetailsResponse struct {
	Id int `json:"id"` // ID identifying this unique release.
	Is_external_build bool `json:"is_external_build,omitempty"` // This value determines if a release is external or not.
	Origin string `json:"origin,omitempty"` // The release's origin
	Uploaded_at string `json:"uploaded_at"` // UTC time in ISO 8601 format of the uploaded time.
	Build map[string]interface{} `json:"build,omitempty"` // Build information for the release
	Short_version string `json:"short_version"` // The release's short version.<br> For iOS: CFBundleShortVersionString from info.plist.<br> For Android: android:versionName from AppManifest.xml.
	Distribution_groups []interface{} `json:"distribution_groups,omitempty"` // OBSOLETE. Will be removed in next version. A list of distribution groups that are associated with this release.
	Distribution_stores []interface{} `json:"distribution_stores,omitempty"` // OBSOLETE. Will be removed in next version. A list of distribution stores that are associated with this release.
	Destination_type string `json:"destination_type,omitempty"` // OBSOLETE. Will be removed in next version. The destination type.<br> <b>group</b>: The release distributed to internal groups and distribution_groups details will be returned.<br> <b>store</b>: The release distributed to external stores and distribution_stores details will be returned. <br>
	Destinations []interface{} `json:"destinations,omitempty"` // A list of distribution groups or stores.
	Version string `json:"version"` // The release's version.<br> For iOS: CFBundleVersion from info.plist.<br> For Android: android:versionCode from AppManifest.xml.
	Enabled bool `json:"enabled"` // This value determines the whether a release currently is enabled or disabled.
	File_extension string `json:"file_extension,omitempty"` // The file extension of the main (user-uploaded) package file.
}

// AlertBugTrackerRepoOwner represents the AlertBugTrackerRepoOwner schema from the OpenAPI specification
type AlertBugTrackerRepoOwner struct {
	Login string `json:"login,omitempty"`
	Name string `json:"name,omitempty"`
	Id string `json:"id,omitempty"`
}

// ApiTokensGetResponse represents the ApiTokensGetResponse schema from the OpenAPI specification
type ApiTokensGetResponse struct {
	Created_at string `json:"created_at"` // The creation time
	Description string `json:"description,omitempty"` // The description of the token
	Id string `json:"id"` // The unique id (UUID) of the api token
	Scope []string `json:"scope,omitempty"` // The scope for this token.
}

// GitHubInstallationRepository represents the GitHubInstallationRepository schema from the OpenAPI specification
type GitHubInstallationRepository struct {
	Language string `json:"language,omitempty"` // The language in the repository
	Name string `json:"name,omitempty"` // The repository name
	Updated_at string `json:"updated_at,omitempty"` // The repository updated time
	Clone_url string `json:"clone_url,omitempty"` // URL used to clone the repository
	Default_branch string `json:"default_branch,omitempty"` // The default branch for the repo
	Description string `json:"description,omitempty"` // The discription of repository
	Fork bool `json:"fork,omitempty"` // Whether the repository is a fork
}

// Events represents the Events schema from the OpenAPI specification
type Events struct {
	Events []map[string]interface{} `json:"events,omitempty"`
	Total int64 `json:"total,omitempty"` // The total count of events.
	Total_devices int64 `json:"total_devices,omitempty"` // The active device over this period.
}

// AlertingAccessTokenResponse represents the AlertingAccessTokenResponse schema from the OpenAPI specification
type AlertingAccessTokenResponse struct {
	External_user_email string `json:"external_user_email"` // The email of external user that used to authenticate aginst the external oauth provider
	Access_token_id string `json:"access_token_id"` // ID of the access token
	External_account_name string `json:"external_account_name"` // The account name of external user that used to authenticate against the external oauth provider or basic auth
	External_provider_name string `json:"external_provider_name"` // External provider name
}

// CrashGroup represents the CrashGroup schema from the OpenAPI specification
type CrashGroup struct {
	Impacted_users int `json:"impacted_users,omitempty"`
	New_crash_group_id string `json:"new_crash_group_id"`
	Last_occurrence string `json:"last_occurrence"`
	Exception string `json:"exception,omitempty"`
	Count int `json:"count"`
	Crash_group_id string `json:"crash_group_id"`
	Status string `json:"status"`
	Fatal bool `json:"fatal"` // Crash or handled exception
	Annotation string `json:"annotation"`
	First_occurrence string `json:"first_occurrence"`
	Reason_frame interface{} `json:"reason_frame,omitempty"` // frame belonging to the reason of the crash
	App_version string `json:"app_version"`
	Build string `json:"build"`
	Crash_reason string `json:"crash_reason"`
	Display_id string `json:"display_id"`
}

// XcodeVersion represents the XcodeVersion schema from the OpenAPI specification
type XcodeVersion struct {
	Current bool `json:"current,omitempty"` // If the Xcode is latest stable
	Name string `json:"name,omitempty"` // The version name
}

// Version represents the Version schema from the OpenAPI specification
type Version struct {
	Previous_count int64 `json:"previous_count,omitempty"` // The count of previous time range of the version.
	Version string `json:"version,omitempty"` // Version.
	Count int64 `json:"count,omitempty"` // Version count.
}

// ActiveCrashingAppDetails represents the ActiveCrashingAppDetails schema from the OpenAPI specification
type ActiveCrashingAppDetails struct {
	Appswithcrashes []map[string]interface{} `json:"appsWithCrashes,omitempty"` // details of the apps with crashes
	Nextlink string `json:"nextLink,omitempty"`
}

// ExternalStoreResponse represents the ExternalStoreResponse schema from the OpenAPI specification
type ExternalStoreResponse struct {
	TypeField string `json:"type,omitempty"` // Store Type
	Created_by string `json:"created_by,omitempty"` // The ID of the principal that created the store.
	Created_by_principal_type string `json:"created_by_principal_type,omitempty"` // The type of the principal that created the store.
	Id string `json:"id,omitempty"` // Store id
	Intune_details interface{} `json:"intune_details,omitempty"` // Store details for intune
	Name string `json:"name,omitempty"` // Store Name
	Service_connection_id string `json:"service_connection_id,omitempty"` // Id for the shared service connection. In case of Apple / GooglePlay stores, this connection will be used to connect to the Apple / Google stores in App Center.
	Track string `json:"track,omitempty"` // Store track
}

// TeamRequest represents the TeamRequest schema from the OpenAPI specification
type TeamRequest struct {
	Display_name string `json:"display_name"` // The display name of the team
	Name string `json:"name,omitempty"` // The name of the team
	Description string `json:"description,omitempty"` // The description of the team
}

// SessionDurationsDistribution represents the SessionDurationsDistribution schema from the OpenAPI specification
type SessionDurationsDistribution struct {
	Average_duration string `json:"average_duration,omitempty"` // The average session duration for current time range.
	Distribution []map[string]interface{} `json:"distribution,omitempty"` // The count of sessions in these buckets.
	Previous_average_duration string `json:"previous_average_duration,omitempty"` // The previous average session duration for previous time range.
}

// UserInvitationPermissionsUpdateRequest represents the UserInvitationPermissionsUpdateRequest schema from the OpenAPI specification
type UserInvitationPermissionsUpdateRequest struct {
	Permissions []string `json:"permissions"` // The permissions the user has for the app in the invitation
}

// ErrorGroupsSearchResult represents the ErrorGroupsSearchResult schema from the OpenAPI specification
type ErrorGroupsSearchResult struct {
	Errorgroups []map[string]interface{} `json:"errorGroups,omitempty"`
	Hasmoreresults bool `json:"hasMoreResults,omitempty"`
}

// GenericLogDiagnostics represents the GenericLogDiagnostics schema from the OpenAPI specification
type GenericLogDiagnostics struct {
	Message_id string `json:"message_id,omitempty"` // Message ID.
	Timestamp string `json:"timestamp"` // Log creation timestamp.
	Device map[string]interface{} `json:"device"` // Device characteristics.
	Install_id string `json:"install_id"` // Install ID.
	TypeField string `json:"type"` // Log type.
	Event_id string `json:"event_id,omitempty"` // Event ID.
	Properties map[string]interface{} `json:"properties,omitempty"` // event specific properties.
	Session_id string `json:"session_id,omitempty"` // Session ID.
	Event_name string `json:"event_name,omitempty"` // Event name.
}

// LogFlowClearProperty represents the LogFlowClearProperty schema from the OpenAPI specification
type LogFlowClearProperty struct {
	Name string `json:"name"`
	TypeField string `json:"type"`
}

// Crash represents the Crash schema from the OpenAPI specification
type Crash struct {
	Details interface{} `json:"details,omitempty"`
	Device_name string `json:"device_name,omitempty"`
	New_crash_group_id string `json:"new_crash_group_id"`
	User_email string `json:"user_email,omitempty"`
	Crash_id string `json:"crash_id"`
	Stacktrace Stacktrace `json:"stacktrace,omitempty"` // a stacktrace in a processed and prettyfied way
	User_name string `json:"user_name"`
	Build string `json:"build"`
	Device string `json:"device"`
	New_crash_id string `json:"new_crash_id"`
	Timestamp string `json:"timestamp"`
	Display_id string `json:"display_id,omitempty"`
	Os_type string `json:"os_type,omitempty"`
	Version string `json:"version"`
	Os_version string `json:"os_version"`
}

// ResignStatus represents the ResignStatus schema from the OpenAPI specification
type ResignStatus struct {
	Error_code string `json:"error_code,omitempty"` // Error code for any error that occured during the resigning operation.
	Error_message string `json:"error_message,omitempty"` // Error message for any error that occured during the resigning operation.
	Status string `json:"status"` // The status of the resign
}

// StartServiceLogDiagnostics represents the StartServiceLogDiagnostics schema from the OpenAPI specification
type StartServiceLogDiagnostics struct {
	Install_id string `json:"install_id"` // Install ID.
	Timestamp string `json:"timestamp"` // Log creation timestamp.
	TypeField string `json:"type"` // Log type.
	Device map[string]interface{} `json:"device"` // Device characteristics.
}

// TestGDPRAccount represents the TestGDPRAccount schema from the OpenAPI specification
type TestGDPRAccount struct {
	Id string `json:"id,omitempty"`
}

// AndroidBranchConfigurationProperties represents the AndroidBranchConfigurationProperties schema from the OpenAPI specification
type AndroidBranchConfigurationProperties struct {
	Keystorefilename string `json:"keystoreFilename,omitempty"` // The name of the keystore file
	Keystorepassword string `json:"keystorePassword,omitempty"` // The password of the keystore
	Keyalias string `json:"keyAlias,omitempty"` // The key alias
	Runtests bool `json:"runTests,omitempty"` // Whether to run unit tests during the build (default)
	Automaticsigning bool `json:"automaticSigning,omitempty"` // Whether to apply automatic signing or not
	Module string `json:"module,omitempty"` // The Gradle module to build
	Buildvariant string `json:"buildVariant,omitempty"` // The Android build variant to build
	Keypassword string `json:"keyPassword,omitempty"` // The key password
	Keystoreencoded string `json:"keystoreEncoded,omitempty"` // The keystore encoded value
	Runlint bool `json:"runLint,omitempty"` // Whether to run lint checks during the build (default)
	Gradlewrapperpath string `json:"gradleWrapperPath,omitempty"` // Path to the Gradle wrapper script
	Isroot bool `json:"isRoot,omitempty"` // Whether it is the root module or not
}

// GlobalFeatureFlagsResponse represents the GlobalFeatureFlagsResponse schema from the OpenAPI specification
type GlobalFeatureFlagsResponse struct {
	Value map[string]interface{} `json:"value"` // The dictionary of global state values indexed by feature flag names
}

// PrivateGooglePlayConnectionSecretResponse represents the PrivateGooglePlayConnectionSecretResponse schema from the OpenAPI specification
type PrivateGooglePlayConnectionSecretResponse struct {
	Isvalid bool `json:"isValid,omitempty"` // whether the credentials are valid or not
	Servicetype string `json:"serviceType"` // service type of shared connection can be apple|gitlab|googleplay|jira|applecertificate
	Displayname string `json:"displayName,omitempty"` // display name of shared connection
	Id string `json:"id"` // id of the shared connection
	Is2fa bool `json:"is2FA,omitempty"` // if the account is a 2FA account or not
}

// DistributionResponse represents the DistributionResponse schema from the OpenAPI specification
type DistributionResponse struct {
	Status string `json:"status,omitempty"` // Status of the Request
	Upload_id string `json:"upload_id,omitempty"` // A unique ID of the upload
}

// SharedConnectionDetails represents the SharedConnectionDetails schema from the OpenAPI specification
type SharedConnectionDetails struct {
	Password string `json:"password,omitempty"` // password to connect to shared connection.
	Username string `json:"username,omitempty"` // username to connect to shared connection.
}

// ExportConfiguration represents the ExportConfiguration schema from the OpenAPI specification
type ExportConfiguration struct {
	Resource_name string `json:"resource_name,omitempty"` // The resource name on azure
	TypeField string `json:"type"` // Type of export configuration
	Backfill bool `json:"backfill,omitempty"` // Field to determine if backfilling should occur. The default value is true. If set to false export starts from date and time of config creation.
	Export_entities []string `json:"export_entities,omitempty"`
	Resource_group string `json:"resource_group,omitempty"` // The resource group name on azure
}

// DestinationDetails represents the DestinationDetails schema from the OpenAPI specification
type DestinationDetails struct {
	TypeField string `json:"type"`
	Id string `json:"id"`
}

// CodePushReleasePromote represents the CodePushReleasePromote schema from the OpenAPI specification
type CodePushReleasePromote struct {
	Is_disabled bool `json:"is_disabled,omitempty"`
	Is_mandatory bool `json:"is_mandatory,omitempty"`
	Rollout int `json:"rollout,omitempty"`
	Target_binary_range string `json:"target_binary_range,omitempty"`
	Description string `json:"description,omitempty"`
	Label string `json:"label,omitempty"`
}

// AudienceDefinition represents the AudienceDefinition schema from the OpenAPI specification
type AudienceDefinition struct {
	Custom_properties map[string]interface{} `json:"custom_properties,omitempty"` // Custom properties used in the definition.
	Definition string `json:"definition"` // Audience definition in OData format.
	Description string `json:"description,omitempty"` // Audience description.
	Enabled bool `json:"enabled,omitempty"`
}

// LogFlowStartSessionLog represents the LogFlowStartSessionLog schema from the OpenAPI specification
type LogFlowStartSessionLog struct {
	Device map[string]interface{} `json:"device"` // Device characteristics.
	Install_id string `json:"install_id"` // Install ID.
	Timestamp string `json:"timestamp"` // Log creation timestamp.
	TypeField string `json:"type"` // Log type.
}

// UpdateResignStatusRequest represents the UpdateResignStatusRequest schema from the OpenAPI specification
type UpdateResignStatusRequest struct {
	Status string `json:"status"` // The updated status for the resigning request.
	Error_code string `json:"error_code,omitempty"` // Error code if an error occured in the resigning operation.
	Error_message string `json:"error_message,omitempty"` // Error message if an error occured in the resigning operation.
	Releasemetadata map[string]interface{} `json:"releaseMetadata,omitempty"` // releaseMetadata from ios resigner extractor
}

// ExternallyHostedReleaseCreateRequest represents the ExternallyHostedReleaseCreateRequest schema from the OpenAPI specification
type ExternallyHostedReleaseCreateRequest struct {
	Build_number string `json:"build_number,omitempty"` // The build number of the uploaded binary
	Build_version string `json:"build_version"` // The build version of the uploaded binary
	External_download_url string `json:"external_download_url"` // The external URL to the release's binary.
}

// ClearPropertyDiagnostics represents the ClearPropertyDiagnostics schema from the OpenAPI specification
type ClearPropertyDiagnostics struct {
	Name string `json:"name"`
	TypeField string `json:"type"`
}

// CrashGroupPlaces represents the CrashGroupPlaces schema from the OpenAPI specification
type CrashGroupPlaces struct {
	Places []map[string]interface{} `json:"places,omitempty"`
	Crash_count int64 `json:"crash_count,omitempty"`
}

// UserLiteProfileResponse represents the UserLiteProfileResponse schema from the OpenAPI specification
type UserLiteProfileResponse struct {
	Display_name string `json:"display_name"` // The full name of the user. Might for example be first and last name
	Email string `json:"email"` // The email address of the user
	Id string `json:"id"` // The unique id (UUID) of the user
}

// ReleaseStoreDestinationRequest represents the ReleaseStoreDestinationRequest schema from the OpenAPI specification
type ReleaseStoreDestinationRequest struct {
	Id string `json:"id"` // Unique id of the release destination
}

// NewCrashGroupAlertingEvent represents the NewCrashGroupAlertingEvent schema from the OpenAPI specification
type NewCrashGroupAlertingEvent struct {
	Event_timestamp string `json:"event_timestamp"` // ISO 8601 date time when event was generated
	Properties map[string]interface{} `json:"properties,omitempty"` // Obsolete. Use emailProperties.
	Event_id string `json:"event_id"` // A unique identifier for this event instance. Useful for deduplication
	Crash_group_properties map[string]interface{} `json:"crash_group_properties,omitempty"` // Properties of new crash group
}

// OrgDistributionGroupDetailsResponse represents the OrgDistributionGroupDetailsResponse schema from the OpenAPI specification
type OrgDistributionGroupDetailsResponse struct {
	Name string `json:"name"` // The name of the distribution group used in URLs
	Origin string `json:"origin"` // The creation origin of this distribution group
	Display_name string `json:"display_name,omitempty"` // The name of the distribution group
	Id string `json:"id"` // The unique ID of the distribution group
	Is_public bool `json:"is_public"` // Whether the distribution group is public
	Apps []interface{} `json:"apps,omitempty"` // The apps associated with the distribution group
	Total_apps_count float64 `json:"total_apps_count,omitempty"` // The count of apps associated with this distribution group
	Total_users_count float64 `json:"total_users_count,omitempty"` // The count of users in the distribution group
}

// CustomPropertyLog represents the CustomPropertyLog schema from the OpenAPI specification
type CustomPropertyLog struct {
	Install_id string `json:"install_id"` // Install ID.
	Timestamp string `json:"timestamp"` // Log creation timestamp.
	TypeField string `json:"type"` // Log type.
	Device map[string]interface{} `json:"device"` // Device characteristics.
}

// UploadedSymbolInfo represents the UploadedSymbolInfo schema from the OpenAPI specification
type UploadedSymbolInfo struct {
	Platform string `json:"platform"` // The platform the symbol is associated with
	Symbol_id string `json:"symbol_id"` // The symbol id of the symbol binary
}

// EventPropertyValue represents the EventPropertyValue schema from the OpenAPI specification
type EventPropertyValue struct {
	Count int64 `json:"count,omitempty"` // The count of the the event property value.
	Name string `json:"name,omitempty"` // The event property value name.
	Previous_count int64 `json:"previous_count,omitempty"` // The count of previous time range of the event property value.
}

// ServiceResourceUsage represents the ServiceResourceUsage schema from the OpenAPI specification
type ServiceResourceUsage struct {
	Currentusageperiod map[string]interface{} `json:"currentUsagePeriod,omitempty"` // Usage for a single period
}

// CrashOverall represents the CrashOverall schema from the OpenAPI specification
type CrashOverall struct {
	Crash_count int64 `json:"crash_count,omitempty"`
	Device_count int64 `json:"device_count,omitempty"`
}

// AllAccountsAggregatedBillingInformation represents the AllAccountsAggregatedBillingInformation schema from the OpenAPI specification
type AllAccountsAggregatedBillingInformation struct {
	Aggregatedbillings map[string]interface{} `json:"aggregatedBillings,omitempty"` // Aggregated Billing Information for a user or an organization
}

// GenericLogContainer represents the GenericLogContainer schema from the OpenAPI specification
type GenericLogContainer struct {
	Exceeded_max_limit bool `json:"exceeded_max_limit,omitempty"` // indicates if the number of available logs are more than the max allowed return limit(100).
	Last_received_log_timestamp string `json:"last_received_log_timestamp,omitempty"` // the timestamp of the last log received. This value can be used as the start time parameter in the consecutive API call.
	Logs []map[string]interface{} `json:"logs"` // the list of logs
}

// AnalyticsReleasesParameter represents the AnalyticsReleasesParameter schema from the OpenAPI specification
type AnalyticsReleasesParameter struct {
	User_id string `json:"user_id"` // user id
	Distribution_group_id string `json:"distribution_group_id"` // distribution group id
	Release_id int `json:"release_id"` // release id
}

// GooglePlayCredentialNonSecretDetailsResponse represents the GooglePlayCredentialNonSecretDetailsResponse schema from the OpenAPI specification
type GooglePlayCredentialNonSecretDetailsResponse struct {
	Credentialtype string `json:"credentialType,omitempty"` // the type of the credential
	Displayname string `json:"displayName"` // display name of shared connection
	Servicetype string `json:"serviceType"` // service type of shared connection can be apple|gitlab|googleplay|jira
}

// TestReport represents the TestReport schema from the OpenAPI specification
type TestReport struct {
	Platform string `json:"platform"`
	Features []map[string]interface{} `json:"features"`
	Revision float64 `json:"revision"`
	Stats map[string]interface{} `json:"stats"`
	Testtype string `json:"testType"`
	App_upload_id string `json:"app_upload_id"`
	Date string `json:"date"`
	Id string `json:"id"`
	Schema_version float64 `json:"schema_version"`
	Snapshot_fatal_errors []map[string]interface{} `json:"snapshot_fatal_errors,omitempty"`
	Date_finished string `json:"date_finished"`
	Errormessage string `json:"errorMessage,omitempty"`
	Device_logs []map[string]interface{} `json:"device_logs"`
	Finished_device_snapshots []string `json:"finished_device_snapshots"`
}

// JiraSecretDetails represents the JiraSecretDetails schema from the OpenAPI specification
type JiraSecretDetails struct {
	Baseurl string `json:"baseUrl"` // baseUrl to connect to jira instance
	Password string `json:"password"` // password to connect to jira instance
	Username string `json:"username"` // username to connect to jira instance
}

// StoreSecretResponse represents the StoreSecretResponse schema from the OpenAPI specification
type StoreSecretResponse struct {
	Tenant_id string `json:"tenant_id,omitempty"` // Tenant Id for Intune
	TypeField string `json:"type,omitempty"` // Store Type
	Id string `json:"id,omitempty"` // Store id
	Name string `json:"name,omitempty"` // Store Name
	Secret string `json:"secret,omitempty"` // Secret Json
}

// SymbolLocation represents the SymbolLocation schema from the OpenAPI specification
type SymbolLocation struct {
	Uri string `json:"uri"`
}

// DataSubjectRightStatusResponse represents the DataSubjectRightStatusResponse schema from the OpenAPI specification
type DataSubjectRightStatusResponse struct {
	Sasurlexpired bool `json:"sasUrlExpired,omitempty"` // Whether Azure Storage shared access signature (SAS) URL has expired or not.
	Status string `json:"status"` // Status of data subject right request
	Message string `json:"message"` // explanation message of the status
	Sasurl string `json:"sasUrl,omitempty"` // Azure Storage shared access signature (SAS) URL for exported user data.
}

// Destination represents the Destination schema from the OpenAPI specification
type Destination struct {
	Name string `json:"name,omitempty"` // A name identifying a unique distribution group.
	Id string `json:"id"` // ID identifying a unique distribution group.
	Is_latest bool `json:"is_latest,omitempty"` // Is the containing release the latest one in this distribution group.
	Id string `json:"id"` // ID identifying a unique distribution store.
	Name string `json:"name,omitempty"` // A name identifying a unique distribution store.
	Publishing_status string `json:"publishing_status,omitempty"` // publishing status of the release in the store.
	TypeField string `json:"type,omitempty"` // type of the distribution store currently stores type can be intune, googleplay or windows.
	Is_latest bool `json:"is_latest,omitempty"` // Is the containing release the latest one in this distribution store.
	Id string `json:"id,omitempty"` // Id of a distribution group / store. The release will be associated with this distribution group / store. If the distribution group / store doesn't exist a 400 is returned. If both distribution group / store name and id are passed, the id is taking precedence.
	Name string `json:"name,omitempty"` // Name of a distribution group / distribution store. The release will be associated with this distribution group or store. If the distribution group / store doesn't exist a 400 is returned. If both distribution group / store name and id are passed, the id is taking precedence.
	Display_name string `json:"display_name,omitempty"` // Display name for the group or tester
	Destination_type string `json:"destination_type,omitempty"` // Destination can be either store or group.
}

// AvailabilityOfDevicesResponse represents the AvailabilityOfDevicesResponse schema from the OpenAPI specification
type AvailabilityOfDevicesResponse struct {
	Ipods interface{} `json:"ipods"` // ...
	Watches interface{} `json:"watches"` // ...
	Ipads interface{} `json:"ipads"` // ...
	Iphones interface{} `json:"iphones"` // ...
}

// UpdateExternalUrlRequest represents the UpdateExternalUrlRequest schema from the OpenAPI specification
type UpdateExternalUrlRequest struct {
	External_download_url string `json:"external_download_url"` // The external URL to the release's binary.
}

// DistributionGroupUserPostResponse represents the DistributionGroupUserPostResponse schema from the OpenAPI specification
type DistributionGroupUserPostResponse struct {
	Status int `json:"status"` // The status code of the result
	User_email string `json:"user_email,omitempty"` // The email of the user
	Code string `json:"code,omitempty"` // The code of the result
	Invite_pending bool `json:"invite_pending,omitempty"` // Whether the has accepted the invite. Available when an invite is pending, and the value will be "true".
	Message string `json:"message,omitempty"` // The message of the result
}

// AppWithTeamPermissionsResponse represents the AppWithTeamPermissionsResponse schema from the OpenAPI specification
type AppWithTeamPermissionsResponse struct {
	Display_name string `json:"display_name"` // The display name of the app
	Id string `json:"id"` // The unique ID (UUID) of the app
	Os string `json:"os"` // The OS the app will be running on
	Icon_url string `json:"icon_url,omitempty"` // The string representation of the URL pointing to the app's icon
	Description string `json:"description,omitempty"` // The description of the app
	Icon_source string `json:"icon_source,omitempty"` // The string representation of the source of the app's icon
	Owner interface{} `json:"owner"` // The information about the app's owner
	Name string `json:"name"` // The name of the app used in URLs
	Release_type string `json:"release_type,omitempty"` // A one-word descriptive release-type value that starts with a capital letter but is otherwise lowercase
	Updated_at string `json:"updated_at,omitempty"` // The last updated date of this app
	App_secret string `json:"app_secret,omitempty"` // A unique and secret key used to identify the app in communication with the ingestion endpoint for crash reporting and analytics
	Azure_subscription interface{} `json:"azure_subscription,omitempty"`
	Created_at string `json:"created_at,omitempty"` // The created date of this app
	Member_permissions []string `json:"member_permissions,omitempty"` // The permissions of the calling user
	Origin string `json:"origin,omitempty"` // The creation origin of this app
	Platform string `json:"platform,omitempty"` // The platform of the app
	Team_permissions []string `json:"team_permissions,omitempty"` // The permissions the team has for the app
}

// JiraCredentialNonSecretDetailsResponse represents the JiraCredentialNonSecretDetailsResponse schema from the OpenAPI specification
type JiraCredentialNonSecretDetailsResponse struct {
	Credentialtype string `json:"credentialType,omitempty"` // the type of the credential
	Displayname string `json:"displayName"` // display name of shared connection
	Servicetype string `json:"serviceType"` // service type of shared connection can be apple|gitlab|googleplay|jira
}

// Language represents the Language schema from the OpenAPI specification
type Language struct {
	Language_name string `json:"language_name,omitempty"` // Language's name.
	Previous_count int64 `json:"previous_count,omitempty"` // Count of previous lanugage.
	Count int64 `json:"count,omitempty"` // Count current of language.
}

// AudienceBlobResult represents the AudienceBlobResult schema from the OpenAPI specification
type AudienceBlobResult struct {
	Url string `json:"url,omitempty"` // Location of the audience blob.
}

// DeviceSetModel represents the DeviceSetModel schema from the OpenAPI specification
type DeviceSetModel struct {
	Formfactor string `json:"formFactor,omitempty"`
	Manufacturer string `json:"manufacturer,omitempty"`
	Name string `json:"name,omitempty"`
	Releasedate string `json:"releaseDate,omitempty"`
}

// IntuneSecretDetails represents the IntuneSecretDetails schema from the OpenAPI specification
type IntuneSecretDetails struct {
	Id_token string `json:"id_token,omitempty"` // the id token of user
	Refresh_token string `json:"refresh_token,omitempty"` // the refresh token for user
	Refresh_token_expiry string `json:"refresh_token_expiry,omitempty"` // the expiry of refresh token
}

// DistributionGroupWithUsersResponse represents the DistributionGroupWithUsersResponse schema from the OpenAPI specification
type DistributionGroupWithUsersResponse struct {
	Notified_user_count float64 `json:"notified_user_count"` // The count of non-pending users in the distribution group who will be notified by new releases
	Total_groups_count float64 `json:"total_groups_count,omitempty"` // The count of aad groups in the distribution group
	Total_user_count float64 `json:"total_user_count"` // The count of users in the distribution group
	Users []interface{} `json:"users"` // The distribution group users
	Aad_groups []interface{} `json:"aad_groups,omitempty"` // The distribution group aad groups
	Id string `json:"id"` // The unique ID of the distribution group
	Is_public bool `json:"is_public,omitempty"` // Whether the distribution group is public
	Name string `json:"name"` // The name of the distribution group used in URLs
}

// DistributionSettingsResponse represents the DistributionSettingsResponse schema from the OpenAPI specification
type DistributionSettingsResponse struct {
	Default_public bool `json:"default_public"` // Is this app using the "Default to Public Distribution Groups" logic
}

// OrganizationResponse represents the OrganizationResponse schema from the OpenAPI specification
type OrganizationResponse struct {
	Origin string `json:"origin"` // The creation origin of this organization
	Updated_at string `json:"updated_at"` // The date the organization was last updated at
	Avatar_url string `json:"avatar_url,omitempty"` // The URL to a user-uploaded Avatar image
	Created_at string `json:"created_at"` // The creation date of this organization
	Display_name string `json:"display_name"` // The display name of the organization
	Id string `json:"id"` // The internal unique id (UUID) of the organization.
	Name string `json:"name"` // The slug name of the organization
}

// DeviceScreenSize represents the DeviceScreenSize schema from the OpenAPI specification
type DeviceScreenSize struct {
	Cm string `json:"cm,omitempty"`
	In string `json:"in,omitempty"`
}

// ErrorFreeDevicePercentages represents the ErrorFreeDevicePercentages schema from the OpenAPI specification
type ErrorFreeDevicePercentages struct {
	Averagepercentage float64 `json:"averagePercentage,omitempty"` // Average percentage
	Dailypercentages []interface{} `json:"dailyPercentages,omitempty"` // The error-free percentage per day.
}

// DistributionGroupUserDeleteResponse represents the DistributionGroupUserDeleteResponse schema from the OpenAPI specification
type DistributionGroupUserDeleteResponse struct {
	Code string `json:"code,omitempty"` // The code of the result
	Message int `json:"message,omitempty"` // The message of the result
	Status int `json:"status"` // The status code of the result
	User_email string `json:"user_email,omitempty"` // The email of the user
}

// AppleCertificateNonSecretDetails represents the AppleCertificateNonSecretDetails schema from the OpenAPI specification
type AppleCertificateNonSecretDetails struct {
	Displayname string `json:"displayName"` // The display name (CN) of the certificate
	Certificatevalidityenddate string `json:"certificateValidityEndDate"` // The date-time till which the certificate is valid
	Certificatevaliditystartdate string `json:"certificateValidityStartDate"` // The date-time from which the certificate is valid
}

// XamarinToolset represents the XamarinToolset schema from the OpenAPI specification
type XamarinToolset struct {
	Xamarinsolutions []interface{} `json:"xamarinSolutions"` // Xamarin solutions for the toolset
}

// DistributionGroupAADGroupBase represents the DistributionGroupAADGroupBase schema from the OpenAPI specification
type DistributionGroupAADGroupBase struct {
	Display_name string `json:"display_name,omitempty"` // The display name of the aad group
	Aad_group_id string `json:"aad_group_id,omitempty"` // The id of the aad group
}

// StoresBasicReleaseDetails represents the StoresBasicReleaseDetails schema from the OpenAPI specification
type StoresBasicReleaseDetails struct {
	Destination_type string `json:"destination_type,omitempty"` // Destination for this release.
	Distribution_stores []interface{} `json:"distribution_stores,omitempty"` // a list of distribution stores that are associated with this release.
	Id float64 `json:"id,omitempty"` // ID identifying this unique release.
	Short_version string `json:"short_version,omitempty"` // The release's short version. For iOS: CFBundleShortVersionString from info.plist. For Android: android:versionName from AppManifest.xml.
	Uploaded_at string `json:"uploaded_at,omitempty"` // UTC time in ISO 8601 format of the uploaded time.
	Version string `json:"version,omitempty"` // The release's version. For iOS: CFBundleVersion from info.plist. For Android: android:versionCode from AppManifest.xml.
}

// AnalyticsModels represents the AnalyticsModels schema from the OpenAPI specification
type AnalyticsModels struct {
	Models []interface{} `json:"models,omitempty"`
	Total int64 `json:"total,omitempty"`
}

// Failure represents the Failure schema from the OpenAPI specification
type Failure struct {
	Message string `json:"message"`
}

// AppMembershipsResponse represents the AppMembershipsResponse schema from the OpenAPI specification
type AppMembershipsResponse struct {
	Memberships []interface{} `json:"memberships,omitempty"` // An array of all ways a user has access to the app, based on the app_memberships table.
	App_origin string `json:"app_origin,omitempty"` // The app's origin
}

// ReleaseWithDistributionGroupAndUserId represents the ReleaseWithDistributionGroupAndUserId schema from the OpenAPI specification
type ReleaseWithDistributionGroupAndUserId struct {
	Release string `json:"release"` // Release Id.
	Distribution_group string `json:"distribution_group,omitempty"` // Distribution group Id.
}

// LogFlowLog represents the LogFlowLog schema from the OpenAPI specification
type LogFlowLog struct {
	Device map[string]interface{} `json:"device"` // Device characteristics.
	Install_id string `json:"install_id"` // Install ID.
	Timestamp string `json:"timestamp"` // Log creation timestamp.
	TypeField string `json:"type"` // Log type.
}

// UpdateInfoData represents the UpdateInfoData schema from the OpenAPI specification
type UpdateInfoData struct {
	Description string `json:"description,omitempty"`
	Is_disabled bool `json:"is_disabled,omitempty"`
	Is_mandatory bool `json:"is_mandatory,omitempty"`
	Rollout int `json:"rollout,omitempty"`
	Target_binary_range string `json:"target_binary_range,omitempty"`
	Package_size float64 `json:"package_size,omitempty"`
	Should_run_binary_version bool `json:"should_run_binary_version,omitempty"`
	Update_app_version bool `json:"update_app_version,omitempty"`
	Download_url string `json:"download_url,omitempty"`
	Is_available bool `json:"is_available,omitempty"`
	Label string `json:"label,omitempty"`
	Package_hash string `json:"package_hash,omitempty"`
}

// LegacyCodePushApp represents the LegacyCodePushApp schema from the OpenAPI specification
type LegacyCodePushApp struct {
	Manuallyprovisiondeployments bool `json:"manuallyProvisionDeployments,omitempty"` // Whether the user provided their own deployments. Not currently in use.
	Name string `json:"name"` // The app name.
	Os string `json:"os,omitempty"` // The app os.
	Platform string `json:"platform,omitempty"` // The app platform.
}

// PostPlaceholderReleaseUploadForDeleteRequest represents the PostPlaceholderReleaseUploadForDeleteRequest schema from the OpenAPI specification
type PostPlaceholderReleaseUploadForDeleteRequest struct {
	Error_details string `json:"error_details,omitempty"` // Additional details to identify why the placeholder ReleaseUpload is being created.
}

// JavaScriptBranchConfigurationProperties represents the JavaScriptBranchConfigurationProperties schema from the OpenAPI specification
type JavaScriptBranchConfigurationProperties struct {
	Packagejsonpath string `json:"packageJsonPath,omitempty"` // Path to package.json file for the main project, e.g. "package.json" or "myapp/package.json"
	Reactnativeversion string `json:"reactNativeVersion,omitempty"` // Version of React Native from package.json files
	Runtests bool `json:"runTests,omitempty"` // Whether to run Jest unit tests, via npm test, during the build
}

// TokenQueryResult represents the TokenQueryResult schema from the OpenAPI specification
type TokenQueryResult struct {
	Tokens []string `json:"tokens,omitempty"` // List of tokens.
}

// CrashGroupLanguages represents the CrashGroupLanguages schema from the OpenAPI specification
type CrashGroupLanguages struct {
	Crash_count int64 `json:"crash_count,omitempty"`
	Languages []map[string]interface{} `json:"languages,omitempty"`
}

// LegacyAppListResponse represents the LegacyAppListResponse schema from the OpenAPI specification
type LegacyAppListResponse struct {
	Apps []interface{} `json:"apps,omitempty"`
}

// PrivateReleaseDetailsResponse represents the PrivateReleaseDetailsResponse schema from the OpenAPI specification
type PrivateReleaseDetailsResponse struct {
	App_display_name string `json:"app_display_name,omitempty"` // The app's display name.
	App_name string `json:"app_name,omitempty"` // The app's name (extracted from the uploaded release).
	Provisioning_profile_name string `json:"provisioning_profile_name,omitempty"` // The release's provisioning profile name.
	Secondary_download_url string `json:"secondary_download_url,omitempty"` // The URL that hosts the secondary binary for this release, such as the apk file for aab releases.
	Origin string `json:"origin,omitempty"` // The release's origin
	Android_min_api_level string `json:"android_min_api_level,omitempty"` // The release's minimum required Android API level.
	Bundle_identifier string `json:"bundle_identifier,omitempty"` // The identifier of the apps bundle.
	Version string `json:"version,omitempty"` // The release's version.<br> For iOS: CFBundleVersion from info.plist. For Android: android:versionCode from AppManifest.xml.
	Publishing_status string `json:"publishing_status,omitempty"` // the publishing status of the distributed release
	App_icon_url string `json:"app_icon_url,omitempty"` // A URL to the app's icon.
	Provisioning_profile_type string `json:"provisioning_profile_type,omitempty"` // The type of the provisioning profile for the requested app version.
	Status string `json:"status,omitempty"` // OBSOLETE. Will be removed in next version. The availability concept is now replaced with distributed. Any 'available' release will be associated with the default distribution group of an app.</br> The release state.<br> <b>available</b>: The uploaded release has been distributed.<br> <b>unavailable</b>: The uploaded release is not visible to the user. <br>
	Size int `json:"size,omitempty"` // The release's size in bytes.
	Destination_type string `json:"destination_type,omitempty"` // The destination type.<br> <b>group</b>: The release distributed to internal groups and distribution_groups details will be returned.<br> <b>store</b>: The release distributed to external stores and distribution_stores details will be returned. <br>
	Is_provisioning_profile_syncing bool `json:"is_provisioning_profile_syncing,omitempty"` // A flag that determines whether the release's provisioning profile is still extracted or not.
	Device_family string `json:"device_family,omitempty"` // The release's device family.
	Download_url string `json:"download_url,omitempty"` // The URL that hosts the binary for this release.
	Fingerprint string `json:"fingerprint,omitempty"` // MD5 checksum of the release binary.
	Release_notes string `json:"release_notes,omitempty"` // The release's release notes.
	Min_os string `json:"min_os,omitempty"` // The release's minimum required operating system.
	Install_url string `json:"install_url,omitempty"` // The href required to install a release on a mobile device. On iOS devices will be prefixed with `itms-services://?action=download-manifest&url=`
	Distribution_group_id string `json:"distribution_group_id,omitempty"` // the destination where release is distributed
	Id int `json:"id,omitempty"` // ID identifying this unique release.
	Is_external_build bool `json:"is_external_build,omitempty"` // This value determines if a release is external or not.
	Short_version string `json:"short_version,omitempty"` // The release's short version.<br> For iOS: CFBundleShortVersionString from info.plist. For Android: android:versionName from AppManifest.xml.
	Uploaded_at string `json:"uploaded_at,omitempty"` // UTC time in ISO 8601 format of the uploaded time.
}

// BillingInformationPlans represents the BillingInformationPlans schema from the OpenAPI specification
type BillingInformationPlans struct {
	Buildservice map[string]interface{} `json:"buildService,omitempty"` // Billing Plans for a single service
	Testservice map[string]interface{} `json:"testService,omitempty"` // Billing Plans for a single service
}

// HandledErrorReasonFrame represents the HandledErrorReasonFrame schema from the OpenAPI specification
type HandledErrorReasonFrame struct {
	Methodparams string `json:"methodParams,omitempty"` // parameters of the frames method
	Osexceptiontype string `json:"osExceptionType,omitempty"` // OS exception type. (aka. SIGNAL)
	Exceptiontype string `json:"exceptionType,omitempty"` // Exception type.
	Line int `json:"line,omitempty"` // line number
	Method string `json:"method,omitempty"` // name of the method
	Appcode bool `json:"appCode,omitempty"` // this line isn't from any framework
	Codeformatted string `json:"codeFormatted,omitempty"` // Formatted frame string
	Coderaw string `json:"codeRaw,omitempty"` // Unformatted Frame string
	Frameworkname string `json:"frameworkName,omitempty"` // Name of the framework
	Classname string `json:"className,omitempty"` // name of the class
	File string `json:"file,omitempty"` // name of the file
	Classmethod bool `json:"classMethod,omitempty"` // is a class method
	Language string `json:"language,omitempty"` // programming language of the frame
}

// TransferAppAdminRequest represents the TransferAppAdminRequest schema from the OpenAPI specification
type TransferAppAdminRequest struct {
	Issue_id string `json:"issue_id"` // The id of the related Intercom issue.
	New_owner_id string `json:"new_owner_id"` // The internal unique id (UUID) of the user/org.
	Responsible_admin_id string `json:"responsible_admin_id"` // The id of the user who started transfer process.
	Transfer_reason string `json:"transfer_reason"` // The explanation for starting transfer process.
}

// StorePatchRequest represents the StorePatchRequest schema from the OpenAPI specification
type StorePatchRequest struct {
	Service_connection_id string `json:"service_connection_id"` // Service connection id to updated.
}

// ErrorDownload represents the ErrorDownload schema from the OpenAPI specification
type ErrorDownload struct {
}

// DiagnosticsStackTrace represents the DiagnosticsStackTrace schema from the OpenAPI specification
type DiagnosticsStackTrace struct {
	Exception DiagnosticsException `json:"exception,omitempty"` // a exception
	Reason string `json:"reason,omitempty"`
	Threads []DiagnosticsThread `json:"threads,omitempty"`
	Title string `json:"title,omitempty"`
}

// LegacyDeploymentsResponse represents the LegacyDeploymentsResponse schema from the OpenAPI specification
type LegacyDeploymentsResponse struct {
	Deployments []interface{} `json:"deployments,omitempty"`
}

// LogFlowLogContainer represents the LogFlowLogContainer schema from the OpenAPI specification
type LogFlowLogContainer struct {
	Exceeded_max_limit bool `json:"exceeded_max_limit,omitempty"` // indicates if the number of available logs are more than the max allowed return limit(100).
	Last_received_log_timestamp string `json:"last_received_log_timestamp,omitempty"` // the timestamp of the last log received. This value can be used as the start time parameter in the consecutive API call.
	Logs []map[string]interface{} `json:"logs"` // the list of logs
}

// EventSetting represents the EventSetting schema from the OpenAPI specification
type EventSetting struct {
	Default_value string `json:"default_value,omitempty"` // Default frequency of event
	Event_type string `json:"event_type"` // Event Name
	Value string `json:"value"` // Frequency of event
}

// UWPSolution represents the UWPSolution schema from the OpenAPI specification
type UWPSolution struct {
	Path string `json:"path"` // The path to the UWP solution
	Configurations []string `json:"configurations"` // The possible configurations detected for the UWP solution
}

// UserProfileResponseInternal represents the UserProfileResponseInternal schema from the OpenAPI specification
type UserProfileResponseInternal struct {
	Avatar_url string `json:"avatar_url,omitempty"` // The avatar URL of the user
	Can_change_password bool `json:"can_change_password,omitempty"` // User is required to send an old password in order to change the password.
	Display_name string `json:"display_name"` // The full name of the user. Might for example be first and last name
	Email string `json:"email"` // The email address of the user
	Id string `json:"id"` // The unique id (UUID) of the user
	Name string `json:"name"` // The unique name that is used to identify the user.
	Origin string `json:"origin"` // The creation origin of this user
	Permissions []string `json:"permissions,omitempty"` // The permissions the user has for the app
	Admin_role string `json:"admin_role,omitempty"` // The new admin_role
	Feature_flags []string `json:"feature_flags,omitempty"` // The feature flags that are enabled for this app
	Settings interface{} `json:"settings,omitempty"` // The user's settings
}

// FeatureCreateRequest represents the FeatureCreateRequest schema from the OpenAPI specification
type FeatureCreateRequest struct {
	Description string `json:"description,omitempty"` // The friendly name of the feature
	Display_name string `json:"display_name"` // The full (friendly) name of the feature.
	Name string `json:"name"` // The unique name of the feature
	State int `json:"state,omitempty"` // The state of the feature
}

// LogFlowPageLog represents the LogFlowPageLog schema from the OpenAPI specification
type LogFlowPageLog struct {
	Timestamp string `json:"timestamp"` // Log creation timestamp.
	TypeField string `json:"type"` // Log type.
	Device map[string]interface{} `json:"device"` // Device characteristics.
	Install_id string `json:"install_id"` // Install ID.
	Properties map[string]interface{} `json:"properties,omitempty"` // Additional key/value pair parameters.
}

// SessionsPerDevice represents the SessionsPerDevice schema from the OpenAPI specification
type SessionsPerDevice struct {
	Sessions_per_user []map[string]interface{} `json:"sessions_per_user,omitempty"` // The session count for each interval per device.
	Total_count int64 `json:"total_count,omitempty"` // Total session per device count.
	Average_sessions_per_user float64 `json:"average_sessions_per_user,omitempty"` // Average seesion per user.
	Previous_average_sessions_per_user float64 `json:"previous_average_sessions_per_user,omitempty"` // Previous average session per user.
	Previous_total_count int64 `json:"previous_total_count,omitempty"` // Previous total count.
}

// SourceRepository represents the SourceRepository schema from the OpenAPI specification
type SourceRepository struct {
	Name string `json:"name,omitempty"` // The repository name
	Clone_url string `json:"clone_url,omitempty"` // URL used to clone the repository
}

// NumberPropertyDiagnostics represents the NumberPropertyDiagnostics schema from the OpenAPI specification
type NumberPropertyDiagnostics struct {
	Name string `json:"name"`
	TypeField string `json:"type"`
}

// BillingPlanSelection represents the BillingPlanSelection schema from the OpenAPI specification
type BillingPlanSelection struct {
	Count int `json:"count,omitempty"` // Number of instances of the billing plan.
	Plan map[string]interface{} `json:"plan,omitempty"` // Billing Plan
}

// BugTrackerIssuesResult represents the BugTrackerIssuesResult schema from the OpenAPI specification
type BugTrackerIssuesResult struct {
	Issues []interface{} `json:"issues,omitempty"`
}

// AppleCredentialNonSecretDetailsResponse represents the AppleCredentialNonSecretDetailsResponse schema from the OpenAPI specification
type AppleCredentialNonSecretDetailsResponse struct {
	Credentialtype string `json:"credentialType,omitempty"` // the type of the credential
	Displayname string `json:"displayName"` // display name of shared connection
	Servicetype string `json:"serviceType"` // service type of shared connection can be apple|gitlab|googleplay|jira
}

// BranchConfigurationRevisionAuthor represents the BranchConfigurationRevisionAuthor schema from the OpenAPI specification
type BranchConfigurationRevisionAuthor struct {
	Displayname string `json:"displayName,omitempty"`
	Url string `json:"url,omitempty"`
}

// TestCloudProject represents the TestCloudProject schema from the OpenAPI specification
type TestCloudProject struct {
	Path string `json:"path"` // The path to the TestCloud project
	Frameworkproperties interface{} `json:"frameworkProperties,omitempty"`
	Frameworktype string `json:"frameworkType"`
}

// ErrorDetailsv2 represents the ErrorDetailsv2 schema from the OpenAPI specification
type ErrorDetailsv2 struct {
	Code string `json:"code"`
	Message string `json:"message"`
}

// OrganizationAadGroupPatchRequest represents the OrganizationAadGroupPatchRequest schema from the OpenAPI specification
type OrganizationAadGroupPatchRequest struct {
	Role string `json:"role,omitempty"` // The user's role in the organizatiion
}

// IntuneTargetAudience represents the IntuneTargetAudience schema from the OpenAPI specification
type IntuneTargetAudience struct {
	Name string `json:"name,omitempty"` // display name for the target audience/group
}

// SubscriptionMetrics represents the SubscriptionMetrics schema from the OpenAPI specification
type SubscriptionMetrics struct {
	Message_count float64 `json:"message_count"` // The number of messages in the subscription
	Name string `json:"name"` // The name of the subsciption (prefixed with the topic name)
}

// CodePushReleaseUpload represents the CodePushReleaseUpload schema from the OpenAPI specification
type CodePushReleaseUpload struct {
	Id string `json:"id"` // The ID for the newly created upload. It is going to be required later in the process.
	Token string `json:"token"` // The URL encoded token used for upload permissions.
	Upload_domain string `json:"upload_domain"` // The URL domain used to upload the release.
}

// LegacyCodePushStatusMetricMetadata represents the LegacyCodePushStatusMetricMetadata schema from the OpenAPI specification
type LegacyCodePushStatusMetricMetadata struct {
	Previouslabelorappversion string `json:"previousLabelOrAppVersion,omitempty"`
	Status string `json:"status,omitempty"`
	Appversion string `json:"appVersion,omitempty"`
	Clientuniqueid string `json:"clientUniqueId,omitempty"`
	Deploymentkey string `json:"deploymentKey,omitempty"`
	Label string `json:"label,omitempty"`
	Previousdeploymentkey string `json:"previousDeploymentKey,omitempty"`
}

// AlertEmailSettings represents the AlertEmailSettings schema from the OpenAPI specification
type AlertEmailSettings struct {
	Settings []map[string]interface{} `json:"settings"` // The settings the user has for the app
}

// StartSessionLogDiagnostics represents the StartSessionLogDiagnostics schema from the OpenAPI specification
type StartSessionLogDiagnostics struct {
	TypeField string `json:"type"` // Log type.
	Device map[string]interface{} `json:"device"` // Device characteristics.
	Install_id string `json:"install_id"` // Install ID.
	Timestamp string `json:"timestamp"` // Log creation timestamp.
}

// AADGroup represents the AADGroup schema from the OpenAPI specification
type AADGroup struct {
	Aad_group_id string `json:"aad_group_id"` // The id of the aad group
	Display_name string `json:"display_name"` // The display name of the aad group
	Tenant_id string `json:"tenant_id"` // The id of the aad tenant
}

// Commit represents the Commit schema from the OpenAPI specification
type Commit struct {
	Sha string `json:"sha,omitempty"` // The commit SHA
	Url string `json:"url,omitempty"` // The URL to the commit
}

// CrashRawLocation represents the CrashRawLocation schema from the OpenAPI specification
type CrashRawLocation struct {
	Uri string `json:"uri"`
}

// MessageEnvelope represents the MessageEnvelope schema from the OpenAPI specification
type MessageEnvelope struct {
	Messagetype string `json:"messageType,omitempty"` // Type of the message
	Message map[string]interface{} `json:"message,omitempty"` // Body of the message
	Messageid string `json:"messageId,omitempty"` // Unique id of the message
}

// OptimizelyUserMetaDataRequest represents the OptimizelyUserMetaDataRequest schema from the OpenAPI specification
type OptimizelyUserMetaDataRequest struct {
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// DeploymentInternal represents the DeploymentInternal schema from the OpenAPI specification
type DeploymentInternal struct {
	Name string `json:"name"`
	Key string `json:"key,omitempty"`
	Latest_release interface{} `json:"latest_release,omitempty"`
}

// AppResponseInternal represents the AppResponseInternal schema from the OpenAPI specification
type AppResponseInternal struct {
	Icon_url string `json:"icon_url,omitempty"` // The string representation of the URL pointing to the app's icon
	Owner interface{} `json:"owner"` // The information about the app's owner
	Name string `json:"name"` // The name of the app used in URLs
	Display_name string `json:"display_name"` // The display name of the app
	Os string `json:"os"` // The OS the app will be running on
	Release_type string `json:"release_type,omitempty"` // A one-word descriptive release-type value that starts with a capital letter but is otherwise lowercase
	Id string `json:"id"` // The unique ID (UUID) of the app
	Description string `json:"description,omitempty"` // The description of the app
	Icon_source string `json:"icon_source,omitempty"` // The string representation of the source of the app's icon
	App_secret string `json:"app_secret,omitempty"` // A unique and secret key used to identify the app in communication with the ingestion endpoint for crash reporting and analytics
	Azure_subscription interface{} `json:"azure_subscription,omitempty"`
	Created_at string `json:"created_at,omitempty"` // The created date of this app
	Member_permissions []string `json:"member_permissions,omitempty"` // The permissions of the calling user
	Origin string `json:"origin,omitempty"` // The creation origin of this app
	Platform string `json:"platform,omitempty"` // The platform of the app
	Updated_at string `json:"updated_at,omitempty"` // The last updated date of this app
	Cutover_from_hockeyapp_at string `json:"cutover_from_hockeyapp_at,omitempty"` // The cutover date of this app
	Feature_flags []string `json:"feature_flags,omitempty"` // The feature flags that are enabled for this app
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // The repositories associated with this app
	User_permissions map[string]interface{} `json:"user_permissions,omitempty"` // the permissions for the specified app user
}

// BooleanProperty represents the BooleanProperty schema from the OpenAPI specification
type BooleanProperty struct {
	Name string `json:"name"`
	TypeField string `json:"type"`
}

// CommitDetails represents the CommitDetails schema from the OpenAPI specification
type CommitDetails struct {
	Url string `json:"url,omitempty"` // The URL to the commit
	Sha string `json:"sha,omitempty"` // The commit SHA
	Commit map[string]interface{} `json:"commit,omitempty"`
}

// BuildLog represents the BuildLog schema from the OpenAPI specification
type BuildLog struct {
	Value []string `json:"value,omitempty"`
}

// TestCloudFileHashDeprecated represents the TestCloudFileHashDeprecated schema from the OpenAPI specification
type TestCloudFileHashDeprecated struct {
	Relative_path string `json:"relative_path"` // Relative path of the file
	Byte_range string `json:"byte_range,omitempty"` // Range of bytes required to verify ownership of the file
	Checksum string `json:"checksum"` // SHA256 hash of the file
	File_type string `json:"file_type"` // Type of the file
}

// FileAsset represents the FileAsset schema from the OpenAPI specification
type FileAsset struct {
	Token string `json:"token,omitempty"`
	Uploaddomain string `json:"uploadDomain,omitempty"`
	Uploadwindowlocation string `json:"uploadWindowLocation,omitempty"`
	Urlencodedtoken string `json:"urlEncodedToken,omitempty"`
	Id string `json:"id,omitempty"`
	Location string `json:"location,omitempty"`
}

// LegacyDeploymentHistory represents the LegacyDeploymentHistory schema from the OpenAPI specification
type LegacyDeploymentHistory struct {
	Allof interface{} `json:"allOf,omitempty"`
	Description string `json:"description,omitempty"` // The description of the release.
	Originaldeployment string `json:"originalDeployment,omitempty"` // The original deployment of the release, if it's ever been promoted.
	Originallabel string `json:"originalLabel,omitempty"` // The original label of the release, if it's ever been updated.
	Packagehash string `json:"packageHash,omitempty"` // The package's hash value (internal use).
}

// Place represents the Place schema from the OpenAPI specification
type Place struct {
	Previous_count int64 `json:"previous_count,omitempty"` // The count of previous time range of the place.
	Code string `json:"code,omitempty"` // The place code.
	Count int64 `json:"count,omitempty"` // The count of the this place.
}

// PostRepositoryProviderMappingRequest represents the PostRepositoryProviderMappingRequest schema from the OpenAPI specification
type PostRepositoryProviderMappingRequest struct {
	Account_id string `json:"account_id"` // App Center account id to link to this provider and external id
	External_account_id string `json:"external_account_id"` // Id of user in the external provider service
	Provider string `json:"provider"` // Supported external providers of source control repositories
}

// LogFlowDateTimeProperty represents the LogFlowDateTimeProperty schema from the OpenAPI specification
type LogFlowDateTimeProperty struct {
	Name string `json:"name"`
	TypeField string `json:"type"`
}

// LegacyCodePushReleaseModification represents the LegacyCodePushReleaseModification schema from the OpenAPI specification
type LegacyCodePushReleaseModification struct {
	Packageinfo map[string]interface{} `json:"packageInfo"` // The release package information
}

// PublicDistributionGroupsResponse represents the PublicDistributionGroupsResponse schema from the OpenAPI specification
type PublicDistributionGroupsResponse struct {
	Id string `json:"id"` // The id of the distribution group
}

// AppleMultifactorLoginResponse represents the AppleMultifactorLoginResponse schema from the OpenAPI specification
type AppleMultifactorLoginResponse struct {
	Expires string `json:"expires,omitempty"` // The expiry date of the cookie generated by apple
	Username string `json:"username,omitempty"` // the apple developer account.
	Cookie string `json:"cookie,omitempty"` // The 30-day session Token generated by apple after successfully logging in with Multifactor authentication.
}

// DailySession represents the DailySession schema from the OpenAPI specification
type DailySession struct {
	Datetime string `json:"datetime,omitempty"` // The ISO 8601 datetime.
	Count int64 `json:"count,omitempty"`
}

// XamarinSDKBundle represents the XamarinSDKBundle schema from the OpenAPI specification
type XamarinSDKBundle struct {
	Current bool `json:"current,omitempty"` // If the SDK is latest stable
	Monoversion string `json:"monoVersion,omitempty"` // The Mono version
	Sdkbundle string `json:"sdkBundle,omitempty"` // The Xamarin SDK version
	Stable bool `json:"stable,omitempty"` // If the SDK is stable
	Xcodeversions []string `json:"xcodeVersions,omitempty"` // Specific for iOS SDK. A list of Xcode versions supported by current SDK version
}

// BillingPeriod represents the BillingPeriod schema from the OpenAPI specification
type BillingPeriod struct {
	Byaccount interface{} `json:"byAccount,omitempty"` // Selection of a billing plan
	Endtime string `json:"endTime,omitempty"` // Exclusive end of the period.
	Starttime string `json:"startTime,omitempty"` // Inclusive start of the period
}

// ActiveDeviceCounts represents the ActiveDeviceCounts schema from the OpenAPI specification
type ActiveDeviceCounts struct {
	Daily []interface{} `json:"daily,omitempty"` // The active device count for each interval.
	Monthly []interface{} `json:"monthly,omitempty"` // The active device count for each interval with a month's retention.
	Weekly []interface{} `json:"weekly,omitempty"` // The active device count for each interval with a week's retention.
}

// LogWithPropertiesDiagnostics represents the LogWithPropertiesDiagnostics schema from the OpenAPI specification
type LogWithPropertiesDiagnostics struct {
	Install_id string `json:"install_id"` // Install ID.
	Timestamp string `json:"timestamp"` // Log creation timestamp.
	TypeField string `json:"type"` // Log type.
	Device map[string]interface{} `json:"device"` // Device characteristics.
}

// AlertCrashGroupStateChange represents the AlertCrashGroupStateChange schema from the OpenAPI specification
type AlertCrashGroupStateChange struct {
	State string `json:"state,omitempty"`
}

// IntuneCategoryValue represents the IntuneCategoryValue schema from the OpenAPI specification
type IntuneCategoryValue struct {
	Displayname string `json:"displayName,omitempty"` // the display name for the category
	Id string `json:"id,omitempty"` // the id of the category
	Lastmodifieddatetime string `json:"lastModifiedDateTime,omitempty"` // modified date for category
}

// PrivateUpdateUploadDetails represents the PrivateUpdateUploadDetails schema from the OpenAPI specification
type PrivateUpdateUploadDetails struct {
	Error_message string `json:"error_message"` // Message of the error
	Status string `json:"status"`
}

// AppleTestFlightGroupsResponse represents the AppleTestFlightGroupsResponse schema from the OpenAPI specification
type AppleTestFlightGroupsResponse struct {
	Appleid float64 `json:"appleId,omitempty"` // apple id of the group.
	Id string `json:"id,omitempty"` // id of the group.
	Name string `json:"name,omitempty"` // name of the group.
	Providerid float64 `json:"providerId,omitempty"` // provider id of the group.
}

// BranchConfigurationRevision represents the BranchConfigurationRevision schema from the OpenAPI specification
type BranchConfigurationRevision struct {
	Definitionurl string `json:"definitionUrl,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Revision float64 `json:"revision,omitempty"`
	Changetype string `json:"changeType,omitempty"`
	Changedby interface{} `json:"changedBy,omitempty"` // user who made a change in branch configuration
	Changeddate string `json:"changedDate,omitempty"`
}

// DateTimeDecimalCounts represents the DateTimeDecimalCounts schema from the OpenAPI specification
type DateTimeDecimalCounts struct {
	Count float64 `json:"count,omitempty"` // Decimal count of the object.
	Datetime string `json:"datetime,omitempty"` // The ISO 8601 datetime.
}

// JavaScriptToolset represents the JavaScriptToolset schema from the OpenAPI specification
type JavaScriptToolset struct {
	Javascriptsolutions []interface{} `json:"javascriptSolutions,omitempty"` // The React Native solutions detected
	Packagejsonpaths []string `json:"packageJsonPaths"` // Paths for detected package.json files
}

// ReleaseDestinationResponse represents the ReleaseDestinationResponse schema from the OpenAPI specification
type ReleaseDestinationResponse struct {
	Id string `json:"id"` // Unique id for the release destination
	Mandatory_update bool `json:"mandatory_update"` // Flag to mark the release for the provided destinations as mandatory
	Provisioning_status_url string `json:"provisioning_status_url,omitempty"` // The url to check provisioning status.
}

// DeviceList represents the DeviceList schema from the OpenAPI specification
type DeviceList struct {
	Devices []string `json:"devices"`
}

// ProvisioningProfileMigration represents the ProvisioningProfileMigration schema from the OpenAPI specification
type ProvisioningProfileMigration struct {
	Url string `json:"url,omitempty"` // A provisioning profile URL that indicates where to download it from.
	Bundle_id string `json:"bundle_id"` // The bundle/application identifier.
	Expired_at string `json:"expired_at,omitempty"` // The provisioning profile's expiration date in RFC 3339 format, i.e. 2017-07-21T17:32:28Z.
	Is_appex bool `json:"is_appex"` // A boolean value that indicates whether the provisioning profile represents an app extension.
	Name string `json:"name"` // The name of the provisioning profile.
	Team_identifier string `json:"team_identifier"` // The team identifier.
	TypeField int `json:"type"` // The type of provisoning profile.
	Udids []string `json:"udids,omitempty"` // A list of UDIDs of provisioned devices.
}

// PostRestoreAppRequest represents the PostRestoreAppRequest schema from the OpenAPI specification
type PostRestoreAppRequest struct {
	Responsibleadminid string `json:"responsibleAdminId"` // The internal unique id (UUID) of the account of the user, who makes the request.
}

// AndroidModule represents the AndroidModule schema from the OpenAPI specification
type AndroidModule struct {
	Buildtypes []string `json:"buildTypes,omitempty"` // The detected build types of the Android module
	Buildvariants []string `json:"buildVariants,omitempty"` // The detected build variants of the Android module (matrix of product flavor + build type (debug|release))
	Hasbundle bool `json:"hasBundle,omitempty"` // Module contains bundle settings
	Isroot bool `json:"isRoot,omitempty"` // Whether the module is at the root level of the project
	Name string `json:"name"` // Name of the Android module
	Productflavors []string `json:"productFlavors,omitempty"` // The product flavors of the Android module
	Buildconfigurations []interface{} `json:"buildConfigurations,omitempty"` // The detected build configurations of the Android module
}

// ApiTokensCreateResponse represents the ApiTokensCreateResponse schema from the OpenAPI specification
type ApiTokensCreateResponse struct {
	Id string `json:"id"` // The unique id (UUID) of the api token
	Scope []string `json:"scope,omitempty"` // The scope for this token.
	Api_token string `json:"api_token"` // The api token generated will not be accessible again
	Created_at string `json:"created_at"` // The creation time
	Description string `json:"description,omitempty"` // The description of the token
}

// RepoInfo represents the RepoInfo schema from the OpenAPI specification
type RepoInfo struct {
	External_user_id string `json:"external_user_id,omitempty"` // The external user ID
	Repo_url string `json:"repo_url"` // The repository url
}

// IntuneGroups represents the IntuneGroups schema from the OpenAPI specification
type IntuneGroups struct {
	Odata_context string `json:"odata.context,omitempty"` // context
	Value []interface{} `json:"value,omitempty"` // categories for intune app
}

// CloseAccountOrganizationResponse represents the CloseAccountOrganizationResponse schema from the OpenAPI specification
type CloseAccountOrganizationResponse struct {
	Origin string `json:"origin"` // The creation origin of this organization
	Updated_at string `json:"updated_at"` // The date the organization was last updated at
	Avatar_url string `json:"avatar_url,omitempty"` // The URL to a user-uploaded Avatar image
	Created_at string `json:"created_at"` // The creation date of this organization
	Display_name string `json:"display_name"` // The display name of the organization
	Id string `json:"id"` // The internal unique id (UUID) of the organization.
	Name string `json:"name"` // The slug name of the organization
	Collaborators_count float64 `json:"collaborators_count,omitempty"` // The number of collaborators from the organization
}

// GetReleaseStatusResponse represents the GetReleaseStatusResponse schema from the OpenAPI specification
type GetReleaseStatusResponse struct {
	Upload_status string `json:"upload_status"` // The current upload status.
	Error_details string `json:"error_details,omitempty"` // Details describing what went wrong processing the upload. Will only be set if status = 'error'.
	Id string `json:"id"` // The ID for the upload.
	Release_distinct_id float64 `json:"release_distinct_id,omitempty"` // The distinct ID of the release. Will only be set when the status = 'readyToBePublished'.
	Release_url interface{} `json:"release_url,omitempty"` // The URL of the release. Will only be set when the status = 'readyToBePublished'.
}

// CrashGroupContainer represents the CrashGroupContainer schema from the OpenAPI specification
type CrashGroupContainer struct {
	Crash_groups []map[string]interface{} `json:"crash_groups"`
}

// ReleaseTesterDestinationRequest represents the ReleaseTesterDestinationRequest schema from the OpenAPI specification
type ReleaseTesterDestinationRequest struct {
	Email string `json:"email"` // Tester's email address
	Mandatory_update bool `json:"mandatory_update,omitempty"` // Flag to mark the release for the provided destinations as mandatory
	Notify_testers bool `json:"notify_testers,omitempty"` // Flag to enable or disable notifications to testers
}

// OrgComplianceSettingsResponse represents the OrgComplianceSettingsResponse schema from the OpenAPI specification
type OrgComplianceSettingsResponse struct {
	Certificate_connection_id string `json:"certificate_connection_id"` // certificate connection id to wrap and resign the app after wrapping
	Id string `json:"id"` // The internal unique id (UUID) of the organization compliance setting
	Is_mam_enabled bool `json:"is_mam_enabled,omitempty"` // flag to tell if mam warpping is enabled on the Org
	Org_id string `json:"org_id"` // The internal unique id (UUID) of the organization.
}

// VSTSAccount represents the VSTSAccount schema from the OpenAPI specification
type VSTSAccount struct {
	Accounturi string `json:"accountUri,omitempty"` // Account uri
	Projects []interface{} `json:"projects,omitempty"` // Account projects
	User interface{} `json:"user,omitempty"` // VSTS user profile
	Accountid string `json:"accountId,omitempty"` // Account id
	Accountname string `json:"accountName,omitempty"` // Account name
	Accountstatus string `json:"accountStatus,omitempty"` // Account status
	Accounttype string `json:"accountType,omitempty"` // Account type
}

// InvitationDetailResponse represents the InvitationDetailResponse schema from the OpenAPI specification
type InvitationDetailResponse struct {
	Invitation_id string `json:"invitation_id"` // The id of the invitation
	Invited_by interface{} `json:"invited_by"`
	Organization interface{} `json:"organization,omitempty"`
	App interface{} `json:"app,omitempty"`
}

// TestCloudErrorDetails represents the TestCloudErrorDetails schema from the OpenAPI specification
type TestCloudErrorDetails struct {
	Message string `json:"message"` // Human-readable message that describes the error
	Status string `json:"status"` // Status of the operation
}

// OrganizationUserResponse represents the OrganizationUserResponse schema from the OpenAPI specification
type OrganizationUserResponse struct {
	Display_name string `json:"display_name"` // The full name of the user. Might for example be first and last name
	Email string `json:"email"` // The email address of the user
	Joined_at string `json:"joined_at"` // The date when the user joined the organization
	Name string `json:"name"` // The unique name that is used to identify the user.
	Role string `json:"role"` // The role the user has within the organization
}

// DateTimeProperty represents the DateTimeProperty schema from the OpenAPI specification
type DateTimeProperty struct {
	Name string `json:"name"`
	TypeField string `json:"type"`
}

// ExportConfigurationBlobStorageConnectionString represents the ExportConfigurationBlobStorageConnectionString schema from the OpenAPI specification
type ExportConfigurationBlobStorageConnectionString struct {
	Resource_name string `json:"resource_name,omitempty"` // The resource name on azure
	TypeField string `json:"type"` // Type of export configuration
	Backfill bool `json:"backfill,omitempty"` // Field to determine if backfilling should occur. The default value is true. If set to false export starts from date and time of config creation.
	Export_entities []string `json:"export_entities,omitempty"`
	Resource_group string `json:"resource_group,omitempty"` // The resource group name on azure
	Blob_path_format_kind string `json:"blob_path_format_kind,omitempty"` // The path to the blob when enum set to 'WithoutAppId' is 'year/month/day/hour/minute' and when set to 'WithAppId' is 'appId/year/month/day/hour/minute'
}

// TestCloudStartTestRunResult represents the TestCloudStartTestRunResult schema from the OpenAPI specification
type TestCloudStartTestRunResult struct {
	Accepted_devices []string `json:"accepted_devices,omitempty"` // List with names of accepted devices
	Rejected_devices []string `json:"rejected_devices,omitempty"` // List with names and descriptions of rejected devices
}

// SymbolUploadUserInfo represents the SymbolUploadUserInfo schema from the OpenAPI specification
type SymbolUploadUserInfo struct {
	Display_name string `json:"display_name,omitempty"` // The full name of the user. Might for example be first and last name
	Email string `json:"email,omitempty"` // The email of the user
}

// LegacyCodePushRelease represents the LegacyCodePushRelease schema from the OpenAPI specification
type LegacyCodePushRelease struct {
	Diffpackagemap map[string]interface{} `json:"diffPackageMap,omitempty"` // Object containing URL and size of changed package hashes contained in the release
	Manifestbloburl string `json:"manifestBlobUrl,omitempty"` // The URL location of the package's manifest file.
	Releasedbyuserid string `json:"releasedByUserId,omitempty"` // User ID that triggered most recent release
	Ismandatory bool `json:"isMandatory,omitempty"` // Flag used to determine if release is mandatory
	Releasemethod string `json:"releaseMethod,omitempty"` // Method used to deploy release
	Uploadtime int `json:"uploadTime,omitempty"` // Release upload time as epoch Unix timestamp
	Appversion string `json:"appVersion,omitempty"` // The version of the release
	Bloburl string `json:"blobUrl,omitempty"` // Location (URL) of release package
	Isdisabled bool `json:"isDisabled,omitempty"` // Flag used to determine if release is disabled
	Rollout int `json:"rollout,omitempty"` // Percentage (out of 100) that release is deployed to
	Label string `json:"label,omitempty"` // Release label (aka release name)
	Size int `json:"size,omitempty"` // Size of release package
}

// IntuneStoreResponse represents the IntuneStoreResponse schema from the OpenAPI specification
type IntuneStoreResponse struct {
	Target_audience interface{} `json:"target_audience,omitempty"`
	App_category interface{} `json:"app_category,omitempty"`
}

// ErrorAttachmentLocation represents the ErrorAttachmentLocation schema from the OpenAPI specification
type ErrorAttachmentLocation struct {
	Uri string `json:"uri,omitempty"`
}

// TestSeries represents the TestSeries schema from the OpenAPI specification
type TestSeries struct {
	Testruns []map[string]interface{} `json:"testRuns,omitempty"` // Most recent test runs
	Mostrecentactivity string `json:"mostRecentActivity,omitempty"` // Date of the latest test run that used this test series
	Name string `json:"name"` // Name of the test series
	Slug string `json:"slug"` // Unique, human-readable identifier of the test series
}

// DownloadContainer represents the DownloadContainer schema from the OpenAPI specification
type DownloadContainer struct {
	Uri string `json:"uri"` // Download URI
}

// LegacyAccountResponse represents the LegacyAccountResponse schema from the OpenAPI specification
type LegacyAccountResponse struct {
	Account map[string]interface{} `json:"account,omitempty"` // Object containing the account information.
}

// DistributionGroupAADGroupPostResponse represents the DistributionGroupAADGroupPostResponse schema from the OpenAPI specification
type DistributionGroupAADGroupPostResponse struct {
	Display_name string `json:"display_name,omitempty"` // The display name of the AAD group
	Id string `json:"id,omitempty"` // The internal unique id (UUID) of the AAD group.
	Aad_group_id string `json:"aad_group_id,omitempty"` // The AAD unique id (UUID) of the AAD group.
}

// AppInvitationDetailResponse represents the AppInvitationDetailResponse schema from the OpenAPI specification
type AppInvitationDetailResponse struct {
	Distribution_group map[string]interface{} `json:"distribution_group,omitempty"` // The organization that owns the distribution group, if it exists
	App_count float64 `json:"app_count,omitempty"` // The number of apps in the group
	Id string `json:"id"` // The unique ID (UUID) of the invitation
	Is_existing_user bool `json:"is_existing_user"` // Indicates whether the invited user already exists
	App interface{} `json:"app"`
	Invited_by interface{} `json:"invited_by"`
	Permissions []string `json:"permissions,omitempty"` // The permissions the user has for the app
	Email string `json:"email"` // The email address of the invited user
	Invite_type string `json:"invite_type"` // The invitation type
}

// LogContainer represents the LogContainer schema from the OpenAPI specification
type LogContainer struct {
	Logs []map[string]interface{} `json:"logs"` // the list of logs
	Exceeded_max_limit bool `json:"exceeded_max_limit,omitempty"` // indicates if the number of available logs are more than the max allowed return limit(100).
	Last_received_log_timestamp string `json:"last_received_log_timestamp,omitempty"` // the timestamp of the last log received. This value can be used as the start time parameter in the consecutive API call.
}

// DiagnosticsStackFrame represents the DiagnosticsStackFrame schema from the OpenAPI specification
type DiagnosticsStackFrame struct {
	App_code bool `json:"app_code"` // this line isn't from any framework
	Class_method bool `json:"class_method,omitempty"` // is a class method
	Class_name string `json:"class_name,omitempty"` // name of the class
	Code_formatted string `json:"code_formatted"` // Formatted frame string
	Language string `json:"language,omitempty"` // programming language of the frame
	Line int `json:"line,omitempty"` // line number
	Code_raw string `json:"code_raw"` // Raw frame string
	Method string `json:"method,omitempty"` // name of the method
	File string `json:"file,omitempty"` // name of the file
	Framework_name string `json:"framework_name,omitempty"` // Name of the framework
	Address string `json:"address,omitempty"` // address of the frame
	Method_params string `json:"method_params,omitempty"` // parameters of the frames method
	Relevant bool `json:"relevant,omitempty"` // frame should be shown always
}

// DistributionGroupAadGroupsDeleteResponse represents the DistributionGroupAadGroupsDeleteResponse schema from the OpenAPI specification
type DistributionGroupAadGroupsDeleteResponse struct {
	Aad_group_id string `json:"aad_group_id,omitempty"` // The aad id of the group
	Code string `json:"code,omitempty"` // The code of the result
	Message int `json:"message,omitempty"` // The message of the result
	Status int `json:"status"` // The status code of the result
}

// DestinationError represents the DestinationError schema from the OpenAPI specification
type DestinationError struct {
	Id string `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"` // Error Codes:<br> <b>invalid_store_secrets</b>: While distributing to store, secrets provided for store are not valid.<br> <b>store_release_bad_request</b>: Proper package release details for the store is not provided.<br> <b>store_release_unauthorized</b>: User is not authorized to publish to store due to invalid developer credentials.<br> <b>store_release_forbidden</b>: Publish to store is forbidden due to conflicts/errors in the release version and already existing version in the store.<br> <b>store_release_promotion</b>: Release already distributed, promoting a release is not supported.<br> <b>store_track_deactivated</b>: One or more tracks would be deactivated with this release. This is not supported yet.<br> <b>store_release_not_found</b>: App with the given package name is not found in the store.<br> <b>store_release_not_available</b>: The release is not available.<br> <b>internal_server_error</b>: Failed to distribute to a destination due to an internal server error.
}

// XcodeScheme represents the XcodeScheme schema from the OpenAPI specification
type XcodeScheme struct {
	Hastestaction bool `json:"hasTestAction"` // Does scheme have a test action?
	Name string `json:"name"` // Scheme name
	Archiveconfiguration string `json:"archiveConfiguration,omitempty"` // Build configuration set in Archive action
	Archiveproject interface{} `json:"archiveProject,omitempty"`
}

// AppFeatures represents the AppFeatures schema from the OpenAPI specification
type AppFeatures struct {
	Crashgroup_analytics_impactedusers bool `json:"crashgroup_analytics_impactedusers,omitempty"` // App supports the 'impacted users' metric
	Crashgroup_modify_annotation bool `json:"crashgroup_modify_annotation,omitempty"` // App supports modification of crashgroup annotation
	Crashgroup_modify_status bool `json:"crashgroup_modify_status,omitempty"` // App supports modification of crashgroup status
	Search bool `json:"search,omitempty"` // App supports search API
	Crash_download_raw bool `json:"crash_download_raw,omitempty"` // App supports download of raw crashes
	Crashgroup_analytics_crashfreeusers bool `json:"crashgroup_analytics_crashfreeusers,omitempty"` // App supports the 'crash free user' metric
}

// PrivateReleaseUpdateRequest represents the PrivateReleaseUpdateRequest schema from the OpenAPI specification
type PrivateReleaseUpdateRequest struct {
	Publishing_status string `json:"publishing_status,omitempty"` // The store publishing status.
}

// VSTSProfile represents the VSTSProfile schema from the OpenAPI specification
type VSTSProfile struct {
	Id string `json:"id,omitempty"` // Profile id
	Publicalias string `json:"publicAlias,omitempty"` // Profile alias
	Displayname string `json:"displayName,omitempty"` // Profile display name
	Emailaddress string `json:"emailAddress,omitempty"` // Profile email
}

// ResendVerificationRequest represents the ResendVerificationRequest schema from the OpenAPI specification
type ResendVerificationRequest struct {
	Name string `json:"name"` // The email or name of the user to resend verification
}

// XamarinSolution represents the XamarinSolution schema from the OpenAPI specification
type XamarinSolution struct {
	Defaultconfiguration string `json:"defaultConfiguration,omitempty"` // Solution default configuration
	Path string `json:"path"` // Path to solution
	Configurations []string `json:"configurations"` // Solution configurations
}

// ErrorGroupListItem represents the ErrorGroupListItem schema from the OpenAPI specification
type ErrorGroupListItem struct {
	State string `json:"state"`
	Annotation string `json:"annotation,omitempty"`
	Exceptionappcode bool `json:"exceptionAppCode,omitempty"`
	Exceptionmessage string `json:"exceptionMessage,omitempty"`
	Exceptionmethod string `json:"exceptionMethod,omitempty"`
	Exceptiontype string `json:"exceptionType,omitempty"`
	Reasonframes []interface{} `json:"reasonFrames,omitempty"`
	Appversion string `json:"appVersion"`
	Exceptionclassmethod bool `json:"exceptionClassMethod,omitempty"`
	Lastoccurrence string `json:"lastOccurrence"`
	Coderaw string `json:"codeRaw,omitempty"`
	Exceptionfile string `json:"exceptionFile,omitempty"`
	Hidden bool `json:"hidden,omitempty"`
	Devicecount int64 `json:"deviceCount"`
	Errorgroupid string `json:"errorGroupId"`
	Firstoccurrence string `json:"firstOccurrence"`
	Count int64 `json:"count"`
	Exceptionline string `json:"exceptionLine,omitempty"`
	Exceptionclassname string `json:"exceptionClassName,omitempty"`
	Appbuild string `json:"appBuild,omitempty"`
}

// CodePushRelease represents the CodePushRelease schema from the OpenAPI specification
type CodePushRelease struct {
	Is_mandatory bool `json:"is_mandatory,omitempty"`
	Rollout int `json:"rollout,omitempty"`
	Target_binary_range string `json:"target_binary_range,omitempty"`
	Description string `json:"description,omitempty"`
	Is_disabled bool `json:"is_disabled,omitempty"`
	Original_label string `json:"original_label,omitempty"` // Set on 'Promote' and 'Rollback'
	Package_hash string `json:"package_hash,omitempty"`
	Release_method string `json:"release_method,omitempty"` // The release method is unknown if unspecified
	Blob_url string `json:"blob_url,omitempty"`
	Diff_package_map map[string]interface{} `json:"diff_package_map,omitempty"`
	Released_by string `json:"released_by,omitempty"`
	Upload_time int `json:"upload_time,omitempty"`
	Original_deployment string `json:"original_deployment,omitempty"` // Set on 'Promote'
	Size float64 `json:"size,omitempty"`
	Label string `json:"label,omitempty"`
}

// DataSubjectRightEmailRequest represents the DataSubjectRightEmailRequest schema from the OpenAPI specification
type DataSubjectRightEmailRequest struct {
	Email string `json:"email"` // Email used for cancel delete with x-authz-bypass headers
}

// BillingPlansSelection represents the BillingPlansSelection schema from the OpenAPI specification
type BillingPlansSelection struct {
	Buildservice interface{} `json:"buildService,omitempty"` // Selection of a billing plan
	Testservice interface{} `json:"testService,omitempty"` // Selection of a billing plan
}

// ItunesAppsRequest represents the ItunesAppsRequest schema from the OpenAPI specification
type ItunesAppsRequest struct {
	Password string `json:"password,omitempty"` // The password for the Apple Developer account.
	Service_connection_id string `json:"service_connection_id,omitempty"` // The service_connection_id of the stored Apple credentials instead of username, password.
	Team_identifier string `json:"team_identifier,omitempty"` // Identifier of the team to use when logged in.
	Username string `json:"username,omitempty"` // The username for the Apple Developer account.
	Cookie string `json:"cookie,omitempty"` // The 30-day session cookie for multi-factor authentication backed accounts.
}

// ProvisioningProfileFile represents the ProvisioningProfileFile schema from the OpenAPI specification
type ProvisioningProfileFile struct {
	Filename string `json:"fileName,omitempty"` // Name of uploaded provisioning profile
	Targetbundleidentifier string `json:"targetBundleIdentifier,omitempty"` // Target the provisioning profile is used to sign
	Uploadid string `json:"uploadId,omitempty"` // Upload id to App Center File Upload Store
	Fileid string `json:"fileId,omitempty"` // File id from secure file storage
}

// TestRunSummary represents the TestRunSummary schema from the OpenAPI specification
type TestRunSummary struct {
	Date string `json:"date,omitempty"` // Date of the test run.
	Failed float64 `json:"failed,omitempty"` // Number of failed tests
	Passed float64 `json:"passed,omitempty"` // Number of passed tests
	Statusdescription string `json:"statusDescription,omitempty"` // Human-readable status of the test run.
	Completed bool `json:"completed,omitempty"` // Tells whether the test run has completed
}

// CrashAttachment represents the CrashAttachment schema from the OpenAPI specification
type CrashAttachment struct {
	Content_type string `json:"content_type"`
	Crash_id string `json:"crash_id"`
	Created_time string `json:"created_time"`
	File_name string `json:"file_name"`
	Size float64 `json:"size"`
	App_id string `json:"app_id"`
	Attachment_id string `json:"attachment_id"`
	Blob_location string `json:"blob_location"`
}

// StoresReleaseDetails represents the StoresReleaseDetails schema from the OpenAPI specification
type StoresReleaseDetails struct {
	Android_min_api_level string `json:"android_min_api_level,omitempty"` // The release's minimum required Android API level.
	Install_url string `json:"install_url,omitempty"` // The href required to install a release on a mobile device. On iOS devices will be prefixed with `itms-services://?action=download-manifest&url=`
	Min_os string `json:"min_os,omitempty"` // The release's minimum required operating system.
	Download_url string `json:"download_url,omitempty"` // The URL that hosts the binary for this release.
	Id float64 `json:"id,omitempty"` // ID identifying this unique release.
	Distribution_stores []interface{} `json:"distribution_stores,omitempty"` // a list of distribution stores that are associated with this release.
	Status string `json:"status,omitempty"` // OBSOLETE. Will be removed in next version. The availability concept is now replaced with distributed. Any 'available' release will be associated with the default distribution group of an app.</br> The release state.<br> <b>available</b>: The uploaded release has been distributed.<br> <b>unavailable</b>: The uploaded release is not visible to the user. <br>
	App_display_name string `json:"app_display_name,omitempty"` // The app's display name.
	Size float64 `json:"size,omitempty"` // The release's size in bytes.
	Fingerprint string `json:"fingerprint,omitempty"` // MD5 checksum of the release binary.
	Short_version string `json:"short_version,omitempty"` // The release's short version.<br> For iOS: CFBundleShortVersionString from info.plist. For Android: android:versionName from AppManifest.xml.
	Uploaded_at string `json:"uploaded_at,omitempty"` // UTC time in ISO 8601 format of the uploaded time.
	Version string `json:"version,omitempty"` // The release's version.<br> For iOS: CFBundleVersion from info.plist. For Android: android:versionCode from AppManifest.xml.
	App_name string `json:"app_name,omitempty"` // The app's name (extracted from the uploaded release).
	Bundle_identifier string `json:"bundle_identifier,omitempty"` // The identifier of the apps bundle.
	Release_notes string `json:"release_notes,omitempty"` // The release's release notes.
}

// DataSubjectRighBlobContainerInfo represents the DataSubjectRighBlobContainerInfo schema from the OpenAPI specification
type DataSubjectRighBlobContainerInfo struct {
	Blobpath string `json:"blobPath"`
	Sasuri string `json:"sasUri"`
}

// SkipValidationRequest represents the SkipValidationRequest schema from the OpenAPI specification
type SkipValidationRequest struct {
	Skip_validation bool `json:"skip_validation,omitempty"` // true if we want to skip the validation, false otherwise
}

// DeviceConfigurationResponse represents the DeviceConfigurationResponse schema from the OpenAPI specification
type DeviceConfigurationResponse struct {
	Data_url string `json:"data_url"` // A data URL containing a signed mobileconfig profile
}

// OrganizationResponseManagement represents the OrganizationResponseManagement schema from the OpenAPI specification
type OrganizationResponseManagement struct {
	Updated_at string `json:"updated_at"` // The date the organization was last updated at
	Avatar_url string `json:"avatar_url,omitempty"` // The URL to a user-uploaded Avatar image
	Created_at string `json:"created_at"` // The creation date of this organization
	Display_name string `json:"display_name"` // The display name of the organization
	Id string `json:"id"` // The internal unique id (UUID) of the organization.
	Name string `json:"name"` // The slug name of the organization
	Origin string `json:"origin"` // The creation origin of this organization
	Feature_flags []string `json:"feature_flags,omitempty"` // The feature flags that are enabled for this organization
	Created_at string `json:"created_at,omitempty"` // The date when the organization was created
	Email string `json:"email,omitempty"` // The organization email, if the app was synced from HockeyApp
	Updated_at string `json:"updated_at,omitempty"` // The date when the organization was updated
}

// ReleaseCounts represents the ReleaseCounts schema from the OpenAPI specification
type ReleaseCounts struct {
	Counts []map[string]interface{} `json:"counts"`
	Total int64 `json:"total,omitempty"`
}

// Languages represents the Languages schema from the OpenAPI specification
type Languages struct {
	Languages []interface{} `json:"languages,omitempty"`
	Total int64 `json:"total,omitempty"`
}

// OrganizationResponseInternal represents the OrganizationResponseInternal schema from the OpenAPI specification
type OrganizationResponseInternal struct {
	Created_at string `json:"created_at"` // The creation date of this organization
	Display_name string `json:"display_name"` // The display name of the organization
	Id string `json:"id"` // The internal unique id (UUID) of the organization.
	Name string `json:"name"` // The slug name of the organization
	Origin string `json:"origin"` // The creation origin of this organization
	Updated_at string `json:"updated_at"` // The date the organization was last updated at
	Avatar_url string `json:"avatar_url,omitempty"` // The URL to a user-uploaded Avatar image
	Feature_flags []string `json:"feature_flags,omitempty"` // The feature flags that are enabled for this organization
}

// GitHubInstallationLite represents the GitHubInstallationLite schema from the OpenAPI specification
type GitHubInstallationLite struct {
	Id float64 `json:"id,omitempty"` // GitHub Installation Id
	Account interface{} `json:"account,omitempty"` // The GitHub Installation
	App_id float64 `json:"app_id,omitempty"` // GitHub Installation App Id
}

// PostExternalDownloadUrl represents the PostExternalDownloadUrl schema from the OpenAPI specification
type PostExternalDownloadUrl struct {
	Download_url string `json:"download_url,omitempty"` // The new download URL
}

// AppRepoPatchRequest represents the AppRepoPatchRequest schema from the OpenAPI specification
type AppRepoPatchRequest struct {
	External_user_id string `json:"external_user_id,omitempty"` // The external user id from the provider
	Repo_url string `json:"repo_url,omitempty"` // The absolute URL of the repository
	Service_connection_id string `json:"service_connection_id,omitempty"` // The id of the service connection stored in customer credential store
	User_id string `json:"user_id,omitempty"` // The unique id (UUID) of the user
}

// BuildAgentQueue represents the BuildAgentQueue schema from the OpenAPI specification
type BuildAgentQueue struct {
	Queue string `json:"queue"`
}

// GooglePlayConnectionNonSecretResponse represents the GooglePlayConnectionNonSecretResponse schema from the OpenAPI specification
type GooglePlayConnectionNonSecretResponse struct {
	Credentialtype string `json:"credentialType,omitempty"` // the type of the credential
	Displayname string `json:"displayName"` // display name of shared connection
	Servicetype string `json:"serviceType"` // service type of shared connection can be apple|gitlab|googleplay|jira
}

// BranchProperties represents the BranchProperties schema from the OpenAPI specification
type BranchProperties struct {
	Branch interface{} `json:"branch,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
}

// OrganizationUserPatchRequest represents the OrganizationUserPatchRequest schema from the OpenAPI specification
type OrganizationUserPatchRequest struct {
	Role string `json:"role,omitempty"` // The user's role in the organizatiion
}

// ErrorGroupState represents the ErrorGroupState schema from the OpenAPI specification
type ErrorGroupState struct {
	Annotation string `json:"annotation,omitempty"`
	State string `json:"state"`
}

// LegacyDeploymentResponse represents the LegacyDeploymentResponse schema from the OpenAPI specification
type LegacyDeploymentResponse struct {
	Deployment map[string]interface{} `json:"deployment,omitempty"`
}

// CustomPropertyDiagnostics represents the CustomPropertyDiagnostics schema from the OpenAPI specification
type CustomPropertyDiagnostics struct {
	Name string `json:"name"`
	TypeField string `json:"type"`
}

// AddAppTesterResponse represents the AddAppTesterResponse schema from the OpenAPI specification
type AddAppTesterResponse struct {
	Release_id int `json:"release_id"` // The ID of the release the user was added to
	User_id string `json:"user_id"` // The user ID of the tester that needs to be added
}

// FeaturePatchRequest represents the FeaturePatchRequest schema from the OpenAPI specification
type FeaturePatchRequest struct {
	Description string `json:"description,omitempty"` // The friendly name of the feature
	Display_name string `json:"display_name,omitempty"` // The full (friendly) name of the feature.
	State int `json:"state,omitempty"` // The state of the feature
}

// OrgComplianceSettingsRequest represents the OrgComplianceSettingsRequest schema from the OpenAPI specification
type OrgComplianceSettingsRequest struct {
	Certificate_connection_id string `json:"certificate_connection_id"` // certificate connection id to wrap and resign the app after wrapping
}

// AlertingCrashGroup represents the AlertingCrashGroup schema from the OpenAPI specification
type AlertingCrashGroup struct {
	Reason string `json:"reason,omitempty"`
	Stack_trace []string `json:"stack_trace,omitempty"`
	Url string `json:"url,omitempty"`
	App_display_name string `json:"app_display_name,omitempty"`
	App_platform string `json:"app_platform,omitempty"` // SDK/Platform this thread is beeing generated from
	App_version string `json:"app_version,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Stacktrace represents the Stacktrace schema from the OpenAPI specification
type Stacktrace struct {
	Threads []Thread `json:"threads,omitempty"`
	Title string `json:"title,omitempty"`
	Exception Exception `json:"exception,omitempty"` // a exception
	Reason string `json:"reason,omitempty"`
}

// UnhandledErrorLog represents the UnhandledErrorLog schema from the OpenAPI specification
type UnhandledErrorLog struct {
	Device map[string]interface{} `json:"device"` // Device characteristics.
	Install_id string `json:"install_id"` // Install ID.
	Timestamp string `json:"timestamp"` // Log creation timestamp.
	TypeField string `json:"type"` // Log type.
}

// ErrorGroupModels represents the ErrorGroupModels schema from the OpenAPI specification
type ErrorGroupModels struct {
	Errorcount int64 `json:"errorCount,omitempty"`
	Models []map[string]interface{} `json:"models,omitempty"`
}

// StackFrame represents the StackFrame schema from the OpenAPI specification
type StackFrame struct {
	Address string `json:"address,omitempty"` // address of the frame
	App_code bool `json:"app_code"` // this line isn't from any framework
	Framework_name string `json:"framework_name,omitempty"` // Name of the framework
	Line int `json:"line,omitempty"` // line number
	Class_name string `json:"class_name,omitempty"` // name of the class
	Code_raw string `json:"code_raw"` // Raw frame string
	Method_params string `json:"method_params,omitempty"` // parameters of the frames method
	Code_formatted string `json:"code_formatted"` // Formatted frame string
	File string `json:"file,omitempty"` // name of the file
	Language string `json:"language,omitempty"` // programming language of the frame
	Relevant bool `json:"relevant,omitempty"` // frame should be shown always
	Class_method bool `json:"class_method,omitempty"` // is a class method
	Method string `json:"method,omitempty"` // name of the method
}

// V2MissingSymbolCrashGroupsResponse represents the V2MissingSymbolCrashGroupsResponse schema from the OpenAPI specification
type V2MissingSymbolCrashGroupsResponse struct {
	Groups []interface{} `json:"groups"` // list of crash groups formed by missing symbols combination
	Total_crash_count int `json:"total_crash_count"` // total number of crashes for all the groups
}

// CustomPropertyLogDiagnostics represents the CustomPropertyLogDiagnostics schema from the OpenAPI specification
type CustomPropertyLogDiagnostics struct {
	Device map[string]interface{} `json:"device"` // Device characteristics.
	Install_id string `json:"install_id"` // Install ID.
	Timestamp string `json:"timestamp"` // Log creation timestamp.
	TypeField string `json:"type"` // Log type.
}

// AppRequest represents the AppRequest schema from the OpenAPI specification
type AppRequest struct {
	Release_type string `json:"release_type,omitempty"` // A one-word descriptive release-type value that starts with a capital letter but is otherwise lowercase
	Description string `json:"description,omitempty"` // A short text describing the app
	Display_name string `json:"display_name"` // The descriptive name of the app. This can contain any characters
	Name string `json:"name,omitempty"` // The name of the app used in URLs
	Os string `json:"os"` // The OS the app will be running on
	Platform string `json:"platform"` // The platform of the app
}

// UserUpdateRequestInternal represents the UserUpdateRequestInternal schema from the OpenAPI specification
type UserUpdateRequestInternal struct {
	Display_name string `json:"display_name,omitempty"` // The full name of the user. Might for example be first and last name
	Email string `json:"email,omitempty"` // The email address for this user
	Name string `json:"name,omitempty"` // The new, unique name that is used to identify.
	Next_nps_survey_date string `json:"next_nps_survey_date,omitempty"` // The date in the future when the user should be checked again for NPS eligibility
}

// GenericLog represents the GenericLog schema from the OpenAPI specification
type GenericLog struct {
	Event_id string `json:"event_id,omitempty"` // Event ID.
	Account_id string `json:"account_id,omitempty"` // Account ID of the authenticated user.
	Message_id string `json:"message_id,omitempty"` // Message ID.
	Device map[string]interface{} `json:"device"` // Device characteristics.
	Event_name string `json:"event_name,omitempty"` // Event name.
	Install_id string `json:"install_id"` // Install ID.
	Properties map[string]interface{} `json:"properties,omitempty"` // event specific properties.
	Session_id string `json:"session_id,omitempty"` // Session ID.
	Auth_provider string `json:"auth_provider,omitempty"` // Auth service provider.
	Timestamp string `json:"timestamp"` // Log creation timestamp.
	TypeField string `json:"type"` // Log type.
}

// CrashGroupOperatingSystems represents the CrashGroupOperatingSystems schema from the OpenAPI specification
type CrashGroupOperatingSystems struct {
	Crash_count int64 `json:"crash_count,omitempty"`
	Operating_systems []map[string]interface{} `json:"operating_systems,omitempty"`
}

// ReleaseUpdateRequest represents the ReleaseUpdateRequest schema from the OpenAPI specification
type ReleaseUpdateRequest struct {
	Mandatory_update bool `json:"mandatory_update,omitempty"` // A boolean which determines whether this version should be a mandatory update or not.
	Notify_testers bool `json:"notify_testers,omitempty"` // A boolean which determines whether to notify testers of a new release, default to true.
	Destination_id string `json:"destination_id,omitempty"` // OBSOLETE. Will be removed in future releases - use destinations instead. Id of a destination. The release will be associated with this destination. If the destination doesn't exist a 400 is returned. If both destination name and id are passed, the id is taking precedence.
	Destination_type string `json:"destination_type,omitempty"` // Not used anymore.
	Destinations []map[string]interface{} `json:"destinations,omitempty"` // Distribute this release under the following list of destinations (store groups or distribution groups).
	Metadata map[string]interface{} `json:"metadata,omitempty"` // An object containing all the release metadata.
	Release_notes string `json:"release_notes,omitempty"` // Release notes for this release.
	Build map[string]interface{} `json:"build,omitempty"` // Contains metadata about the build that produced the release being uploaded
	Distribution_group_id string `json:"distribution_group_id,omitempty"` // OBSOLETE. Will be removed in future releases - use destinations instead. Id of a distribution group. The release will be associated with this distribution group. If the distribution group doesn't exist a 400 is returned. If both distribution group name and id are passed, the id is taking precedence.
	Distribution_group_name string `json:"distribution_group_name,omitempty"` // OBSOLETE. Will be removed in future releases - use destinations instead. Name of a distribution group. The release will be associated with this distribution group. If the distribution group doesn't exist a 400 is returned. If both distribution group name and id are passed, the id is taking precedence.
	Destination_name string `json:"destination_name,omitempty"` // OBSOLETE. Will be removed in future releases - use destinations instead. Name of a destination. The release will be associated with this destination. If the destination doesn't exist a 400 is returned. If both distribution group name and id are passed, the id is taking precedence.
}

// GDPRValidationRequest represents the GDPRValidationRequest schema from the OpenAPI specification
type GDPRValidationRequest struct {
	Release_ids []int `json:"release_ids"` // a list of release ids to validate (AC flow)
	Release_upload_ids []string `json:"release_upload_ids,omitempty"` // a list of release uploads ids to validate
	Hockeyapp_release_ids []int `json:"hockeyapp_release_ids,omitempty"` // a list of release ids to validate (HA flow)
}

// CreateAccessKeyRequest represents the CreateAccessKeyRequest schema from the OpenAPI specification
type CreateAccessKeyRequest struct {
	Createdby string `json:"createdBy,omitempty"` // Name of creator current access key
	Friendlyname string `json:"friendlyName,omitempty"` // Friendly name of the access key
	Ttl float64 `json:"ttl,omitempty"` // Time to live of the access key
}

// ExternalStoreRequest represents the ExternalStoreRequest schema from the OpenAPI specification
type ExternalStoreRequest struct {
	Track string `json:"track,omitempty"` // track of the store. Can be production, alpha & beta for googleplay. Can be production, testflight-internal & testflight-external for Apple Store.
	TypeField string `json:"type,omitempty"` // store Type
	Intune_details interface{} `json:"intune_details,omitempty"`
	Name string `json:"name,omitempty"` // name of the store. In case of googleplay, and Apple store this is fixed to Production.
	Service_connection_id string `json:"service_connection_id,omitempty"` // Id for the shared service connection. In case of Apple AppStore, this connection will be used to create and connect to the Apple AppStore in Mobile Center.
}

// DistributionGroupAppAddRequest represents the DistributionGroupAppAddRequest schema from the OpenAPI specification
type DistributionGroupAppAddRequest struct {
	Apps []interface{} `json:"apps,omitempty"` // The list of apps to add to distribution group
}

// GitHubBranchLite represents the GitHubBranchLite schema from the OpenAPI specification
type GitHubBranchLite struct {
	Sha string `json:"sha,omitempty"` // The discription of repository
	Ref string `json:"ref,omitempty"` // The repository name
	Repo interface{} `json:"repo,omitempty"` // The lite version of GitHub repository
}

// BranchConfiguration represents the BranchConfiguration schema from the OpenAPI specification
type BranchConfiguration struct {
	Badgeisenabled bool `json:"badgeIsEnabled,omitempty"`
	Clonefrombranch string `json:"cloneFromBranch,omitempty"` // A configured branch name to clone from. If provided, all other parameters will be ignored. Only supported in POST requests.
	Signed bool `json:"signed,omitempty"`
	Testsenabled bool `json:"testsEnabled,omitempty"`
	Toolsets interface{} `json:"toolsets,omitempty"` // The branch build configuration for each toolset
	Trigger string `json:"trigger,omitempty"`
	Artifactversioning map[string]interface{} `json:"artifactVersioning,omitempty"` // The versioning configuration for artifacts built for this branch
}

// ReasonStackFrame represents the ReasonStackFrame schema from the OpenAPI specification
type ReasonStackFrame struct {
	Framework_name string `json:"framework_name,omitempty"` // Name of the framework
	Class_method bool `json:"class_method,omitempty"` // is a class method
	Class_name string `json:"class_name,omitempty"` // name of the class
	Code_formatted string `json:"code_formatted,omitempty"` // Formatted frame string
	Code_raw string `json:"code_raw,omitempty"` // Unformatted Frame string
	Language string `json:"language,omitempty"` // programming language of the frame
	Line int `json:"line,omitempty"` // line number
	Method string `json:"method,omitempty"` // name of the method
	Method_params string `json:"method_params,omitempty"` // parameters of the frames method
	Os_exception_type string `json:"os_exception_type,omitempty"` // OS exception type. (aka. SIGNAL)
	App_code bool `json:"app_code,omitempty"` // this line isn't from any framework
	Exception_type string `json:"exception_type,omitempty"` // Exception type.
	File string `json:"file,omitempty"` // name of the file
}

// TestGDPRFeatureFlag represents the TestGDPRFeatureFlag schema from the OpenAPI specification
type TestGDPRFeatureFlag struct {
	Name string `json:"name,omitempty"`
	Target_id string `json:"target_id,omitempty"`
}

// ApiTokenClaim represents the ApiTokenClaim schema from the OpenAPI specification
type ApiTokenClaim struct {
	Claim_type string `json:"claim_type,omitempty"`
	Claim_value string `json:"claim_value,omitempty"`
}

// PatchRepoInfo represents the PatchRepoInfo schema from the OpenAPI specification
type PatchRepoInfo struct {
	External_user_id string `json:"external_user_id,omitempty"` // The external user ID
}

// CrashGroupModels represents the CrashGroupModels schema from the OpenAPI specification
type CrashGroupModels struct {
	Crash_count int64 `json:"crash_count,omitempty"`
	Models []map[string]interface{} `json:"models,omitempty"`
}

// IntuneGroupValue represents the IntuneGroupValue schema from the OpenAPI specification
type IntuneGroupValue struct {
	Displayname string `json:"displayName,omitempty"` // the display name of the group
	Id string `json:"id,omitempty"` // the id of the Group
}

// TestCloudHashUploadStatus represents the TestCloudHashUploadStatus schema from the OpenAPI specification
type TestCloudHashUploadStatus struct {
	Statuscode float64 `json:"statusCode"` // HTTP status code that represent result of upload
	Location string `json:"location,omitempty"` // URI that should be used to make POST request if file with given hash doesn't exist. This is set when status_code is equal to 412
}

// Places represents the Places schema from the OpenAPI specification
type Places struct {
	Places []map[string]interface{} `json:"places,omitempty"`
	Total int64 `json:"total,omitempty"`
}

// TestGDPRApp represents the TestGDPRApp schema from the OpenAPI specification
type TestGDPRApp struct {
	Hash_files_url string `json:"hash_files_url,omitempty"`
}

// ReleaseExternalUrlUpdateResponse represents the ReleaseExternalUrlUpdateResponse schema from the OpenAPI specification
type ReleaseExternalUrlUpdateResponse struct {
	External_download_url string `json:"external_download_url,omitempty"`
}

// GitHubMarketplacePlan represents the GitHubMarketplacePlan schema from the OpenAPI specification
type GitHubMarketplacePlan struct {
	Id int `json:"id,omitempty"` // Id of the GitHub plan
}

// FilterVersionsContainer represents the FilterVersionsContainer schema from the OpenAPI specification
type FilterVersionsContainer struct {
	Versions []map[string]interface{} `json:"versions,omitempty"`
}

// EventPropertyValues represents the EventPropertyValues schema from the OpenAPI specification
type EventPropertyValues struct {
	Total int64 `json:"total,omitempty"` // The total property value counts.
	Values []map[string]interface{} `json:"values,omitempty"` // The event property values.
}

// ReleaseStoreDestinationResponse represents the ReleaseStoreDestinationResponse schema from the OpenAPI specification
type ReleaseStoreDestinationResponse struct {
	Id string `json:"id"` // Unique id for the release destination
}

// AvailabilityOfDevicesRequest represents the AvailabilityOfDevicesRequest schema from the OpenAPI specification
type AvailabilityOfDevicesRequest struct {
	Username string `json:"username,omitempty"` // The username for the Apple Developer account.
	Password string `json:"password,omitempty"` // The password for the Apple Developer account.
	Service_connection_id string `json:"service_connection_id,omitempty"` // The service_connection_id of the stored Apple credentials instad of username, password.
}

// ErrorGroup represents the ErrorGroup schema from the OpenAPI specification
type ErrorGroup struct {
	State string `json:"state"`
	Annotation string `json:"annotation,omitempty"`
}

// AlertWebhookPingResult represents the AlertWebhookPingResult schema from the OpenAPI specification
type AlertWebhookPingResult struct {
	Request_id string `json:"request_id"` // Unique request identifier for tracking
	Response_reason string `json:"response_reason,omitempty"` // Reason returned in response from calling webhook
	Response_status_code int `json:"response_status_code"` // HTTP status code returned in response from calling webhook
}

// XcodeArchiveProject represents the XcodeArchiveProject schema from the OpenAPI specification
type XcodeArchiveProject struct {
	Archivetargetid string `json:"archiveTargetId"` // The Id of the target to archive
	Projectname string `json:"projectName"` // The project to archive container name
	Projectpath string `json:"projectPath,omitempty"` // Full path of the target project
}

// V2MissingSymbolCrashGroup represents the V2MissingSymbolCrashGroup schema from the OpenAPI specification
type V2MissingSymbolCrashGroup struct {
	App_build string `json:"app_build"` // application build
	App_id string `json:"app_id"` // application id
	App_ver string `json:"app_ver"` // application version
	Error_count int `json:"error_count,omitempty"` // number of errors that belong to this group
	Crash_count int `json:"crash_count,omitempty"` // number of crashes that belong to this group
	Missing_symbols []interface{} `json:"missing_symbols"` // list of missing symbols
	Status string `json:"status"` // group status
	Last_modified string `json:"last_modified"` // last update date for the group
	Symbol_group_id string `json:"symbol_group_id"` // id of the symbol group
}

// AppleSecretDetails represents the AppleSecretDetails schema from the OpenAPI specification
type AppleSecretDetails struct {
	Authcode string `json:"authCode,omitempty"` // 6 digit auth code
	Password string `json:"password,omitempty"` // password to connect to apple store.
	Username string `json:"username,omitempty"` // username to connect to apple store.
}

// JiraConnectionNonSecretResponse represents the JiraConnectionNonSecretResponse schema from the OpenAPI specification
type JiraConnectionNonSecretResponse struct {
	Displayname string `json:"displayName"` // display name of shared connection
	Servicetype string `json:"serviceType"` // service type of shared connection can be apple|gitlab|googleplay|jira
	Credentialtype string `json:"credentialType,omitempty"` // the type of the credential
}

// AzureSubscriptionResponse represents the AzureSubscriptionResponse schema from the OpenAPI specification
type AzureSubscriptionResponse struct {
	Is_billing bool `json:"is_billing,omitempty"` // If the subscription is used for billing
	Is_microsoft_internal bool `json:"is_microsoft_internal,omitempty"` // If the subscription is internal Microsoft subscription
	Subscription_id string `json:"subscription_id"` // The azure subscription id
	Subscription_name string `json:"subscription_name"` // The name of the azure subscription
	Tenant_id string `json:"tenant_id"` // The tenant id of the azure subscription belongs to
	Is_billable bool `json:"is_billable,omitempty"` // If the subscription can be used for billing
}

// CrashDetails represents the CrashDetails schema from the OpenAPI specification
type CrashDetails struct {
	App_start_timestamp string `json:"app_start_timestamp,omitempty"` // Application launch timestamp (example: 1985-04-12T23:20:50.52Z).
	Carrier_country string `json:"carrier_country,omitempty"` // Carrier country code (for mobile devices).
	Carrier_name string `json:"carrier_name,omitempty"` // Carrier name (for mobile devices).
	Locale string `json:"locale"` // Language code (example: en_US).
	Os_build string `json:"os_build,omitempty"` // OS build code (example: LMY47X).
	Rooted bool `json:"rooted"` // Whether the device where the crash occurred is rooted or jailbroken
	Screen_size string `json:"screen_size"` // Screen size of the device in pixels (example: 640x480).
}

// Tier represents the Tier schema from the OpenAPI specification
type Tier struct {
	Name string `json:"name,omitempty"` // The name of the tier
}

// DeviceSet represents the DeviceSet schema from the OpenAPI specification
type DeviceSet struct {
	Owner map[string]interface{} `json:"owner"` // The owner of a device set
	Slug string `json:"slug,omitempty"` // Slug of the device set
	Deviceconfigurations []interface{} `json:"deviceConfigurations"`
	Id string `json:"id"` // Identifier of the device set
	Manufacturercount float64 `json:"manufacturerCount,omitempty"` // The number of manufacturers in the device set's device selection
	Name string `json:"name"` // Name of the device set
	Osversioncount float64 `json:"osVersionCount,omitempty"` // The number of os versions in the device set's device selection
}

// AppleConnectionSecretRequest represents the AppleConnectionSecretRequest schema from the OpenAPI specification
type AppleConnectionSecretRequest struct {
	Servicetype string `json:"serviceType"` // service type of shared connection can be apple|gitlab|googleplay|jira|applecertificate
	Credentialtype string `json:"credentialType,omitempty"` // credential type of the shared connection. Values can be credentials|certificate
	Data map[string]interface{} `json:"data,omitempty"` // shared connection details
	Displayname string `json:"displayName,omitempty"` // display name of shared connection
}

// HasTestflightMetadataResponse represents the HasTestflightMetadataResponse schema from the OpenAPI specification
type HasTestflightMetadataResponse struct {
	Has_testflight_metadata bool `json:"has_testflight_metadata,omitempty"` // true if the app has the testflight metadata, false otherwise
}

// ExportConfigurationListResult represents the ExportConfigurationListResult schema from the OpenAPI specification
type ExportConfigurationListResult struct {
	Total int64 `json:"total,omitempty"` // the total count of exports
	Values []map[string]interface{} `json:"values"`
	Nextlink string `json:"nextLink,omitempty"`
}

// ReleaseUpdateError represents the ReleaseUpdateError schema from the OpenAPI specification
type ReleaseUpdateError struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Destinations []interface{} `json:"destinations,omitempty"`
	Mandatory_update bool `json:"mandatory_update,omitempty"`
	Release_notes string `json:"release_notes,omitempty"`
}

// CodePushUploadedRelease represents the CodePushUploadedRelease schema from the OpenAPI specification
type CodePushUploadedRelease struct {
	Rollout int `json:"rollout,omitempty"` // This specifies the percentage of users (as an integer between 1 and 100) that should be eligible to receive this update.
	Target_binary_version string `json:"target_binary_version"` // the binary version of the application
	Deployment_name string `json:"deployment_name,omitempty"` // This specifies which deployment you want to release the update to. Default is Staging.
	Description string `json:"description,omitempty"` // This provides an optional "change log" for the deployment.
	Disabled bool `json:"disabled,omitempty"` // This specifies whether an update should be downloadable by end users or not.
	Mandatory bool `json:"mandatory,omitempty"` // This specifies whether the update should be considered mandatory or not (e.g. it includes a critical security fix).
	No_duplicate_release_error bool `json:"no_duplicate_release_error,omitempty"` // This specifies that if the update is identical to the latest release on the deployment, the CLI should generate a warning instead of an error.
	Release_upload interface{} `json:"release_upload"` // The upload metadata from the release initialization step.
}

// V2FailureResponse represents the V2FailureResponse schema from the OpenAPI specification
type V2FailureResponse struct {
	Message string `json:"message"`
	Code string `json:"code"`
}

// XcodeSchemeContainer represents the XcodeSchemeContainer schema from the OpenAPI specification
type XcodeSchemeContainer struct {
	Workspaceprojectpaths string `json:"workspaceProjectPaths,omitempty"` // Related projects paths for xcworkspace
	Xcodeprojectsha string `json:"xcodeProjectSha,omitempty"` // repo object Id of the pbxproject
	Appextensiontargets []interface{} `json:"appExtensionTargets,omitempty"` // Information regarding project app extensions, if present
	Cartfilepath string `json:"cartfilePath,omitempty"` // Path to Carthage file, if present
	Path string `json:"path"` // Path to project
	Podfilepath string `json:"podfilePath,omitempty"` // Path to CocoaPods file, if present
	Sharedschemes []interface{} `json:"sharedSchemes"` // Project schemes
}

// AllItunesAppsResponse represents the AllItunesAppsResponse schema from the OpenAPI specification
type AllItunesAppsResponse struct {
	Iconurl string `json:"iconUrl,omitempty"` // url for the app icon from app store
	Name string `json:"name,omitempty"` // App Name
	Apple_id string `json:"apple_id,omitempty"` // apple id for app team id.
	Bundle_id string `json:"bundle_id,omitempty"` // bundle identifier of app
}

// AttributesValues represents the AttributesValues schema from the OpenAPI specification
type AttributesValues struct {
}

// RepoConfigCommon represents the RepoConfigCommon schema from the OpenAPI specification
type RepoConfigCommon struct {
	External_user_id string `json:"external_user_id,omitempty"` // The external user id from the repository provider. Required for GitLab.com repositories
	Repo_id string `json:"repo_id,omitempty"` // The repository id from the repository provider. Required for repositories connected from GitHub App and GitLab.com
	Repo_url string `json:"repo_url,omitempty"` // The repository's git url, must be a HTTPS URL
	Service_connection_id string `json:"service_connection_id,omitempty"` // The id of the service connection (private). Required for GitLab self-hosted repositories
}
