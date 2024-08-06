package plugin

import (
	"fmt"
	"log/slog"
)

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

func (receiver *SDKVersion) String() string {
	if receiver == nil {
		return ""
	}

	if receiver.ID == "" {
		return ""
	}

	var out = receiver.ID

	switch receiver.Type {
	case TypeStable:
	case TypeUnstable, TypeArchived:
		out += fmt.Sprintf(" (%s)", receiver.Type)
	default:
		slog.Error(fmt.Sprintf("unknown SDK version type: %s", receiver.Type))
	}

	if receiver.Installed {
		out += " [installed]"
	}

	return out
}
