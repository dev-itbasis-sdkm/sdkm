package sdkversion

import (
	"fmt"
	"log/slog"
)

type PrintFormatOptions struct {
	OutputType bool

	OutputInstalled    bool
	OutputNotInstalled bool
}

func (receiver *SDKVersion) Print() string {
	return receiver.PrintWithOptions(true, true, false)
}

//nolint:cyclop // FIXME
func (receiver *SDKVersion) PrintWithOptions(outType, outInstalled, outNotInstalled bool) string {
	if receiver == nil || receiver.ID == "" {
		return ""
	}

	var out = receiver.ID

	if outType {
		switch receiver.Type {
		case TypeStable:
		case TypeUnstable, TypeArchived:
			out += fmt.Sprintf(" (%s)", receiver.Type)
		default:
			slog.Error(fmt.Sprintf("unknown SDK version type: %s", receiver.Type))
		}
	}

	if outInstalled && receiver.Installed {
		out += " [installed]"
	} else if outNotInstalled && !receiver.Installed {
		out += " [not installed]"
	}

	return out
}
