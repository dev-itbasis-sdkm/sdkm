package filestorage

import (
	"fmt"
	"strings"
	"time"
)

const updatedFormat = time.DateOnly

type updated time.Time

//goland:noinspection GoMixedReceiverTypes
func (receiver *updated) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)

	t, err := time.Parse(updatedFormat, s)
	if err != nil {
		return err //nolint:wrapcheck // TODO
	}

	*receiver = updated(t)

	return nil
}

//goland:noinspection GoMixedReceiverTypes
func (receiver updated) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", receiver.String())), nil
}

//goland:noinspection GoMixedReceiverTypes
func (receiver updated) String() string {
	t := time.Time(receiver)

	return t.Format(updatedFormat)
}
