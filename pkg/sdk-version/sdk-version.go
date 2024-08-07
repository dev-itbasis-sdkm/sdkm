package sdkversion

type VersionType string

const (
	TypeStable   VersionType = "stable"
	TypeUnstable VersionType = "unstable"
	TypeArchived VersionType = "archived"
)

type SDKVersion struct {
	ID        string
	Type      VersionType
	Installed bool `json:"-"`
}
