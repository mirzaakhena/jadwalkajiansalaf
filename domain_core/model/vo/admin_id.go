package vo

import (
	"fmt"
)

type AdminID string

func NewAdminID(randomStringID string) (AdminID, error) {
	var obj = AdminID(fmt.Sprintf("ADM-%s", randomStringID))

	// you may change it as necessary ...

	return obj, nil
}

func (r AdminID) String() string {
	return string(r)
}
