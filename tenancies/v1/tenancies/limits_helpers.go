package tenancies

// Please keep in alphabetical order
const (
	AccessPoliciesLimitsLabel     = "AccessPolicies"
	AppRegistrationsLimitsLabel   = "AppRegistrations"
	AssetsKhipuLimitsLabel        = "AssetsKhipu"
	AssetsPublicLimitsLabel       = "AssetsPublic"
	BlobsLimitsLabel              = "Blobs"
	BlobsSizeLimitsLabel          = "BlobsSize"
	CompliancePoliciesLimitsLabel = "CompliancePolicies"
	EnterpriseSSOLimitsLabel      = "EnterpriseSSO"
	InvitesLimitsLabel            = "Invites"
	LocationsLimitsLabel          = "Locations"
	TenantUsersLimitsLabel        = "TenantUsers"
)

var (
	// Please keep in alphabetical order
	ValidLimitsKeys = [...]string{
		AccessPoliciesLimitsLabel,
		AppRegistrationsLimitsLabel,
		AssetsKhipuLimitsLabel,
		AssetsPublicLimitsLabel,
		BlobsLimitsLabel,
		BlobsSizeLimitsLabel,
		CompliancePoliciesLimitsLabel,
		EnterpriseSSOLimitsLabel,
		InvitesLimitsLabel,
		LocationsLimitsLabel,
		TenantUsersLimitsLabel,
	}
)

type Limits map[string]int32

func NewLimits() Limits {
	return make(map[string]int32)
}
