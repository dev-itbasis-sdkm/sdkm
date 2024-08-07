package cache

import (
	"fmt"
	"strings"
	"time"
)

const fileCacheUpdatedFormat = time.DateOnly

type fileCacheUpdated time.Time

//goland:noinspection GoMixedReceiverTypes
func (receiver *fileCacheUpdated) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)

	t, err := time.Parse(fileCacheUpdatedFormat, s)
	if err != nil {
		return err //nolint:wrapcheck // TODO
	}

	*receiver = fileCacheUpdated(t)

	return nil
}

//goland:noinspection GoMixedReceiverTypes
func (receiver fileCacheUpdated) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", receiver.String())), nil
}

//goland:noinspection GoMixedReceiverTypes
func (receiver fileCacheUpdated) String() string {
	t := time.Time(receiver)

	return t.Format(fileCacheUpdatedFormat)
}
