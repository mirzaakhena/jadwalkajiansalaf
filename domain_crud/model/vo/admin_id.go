package vo

import (
	"fmt"
	"time"
)

type AdminID string

func NewAdminID(randomStringID string, now time.Time) (AdminID, error) {
	var obj = AdminID(fmt.Sprintf("OBJ-%s-%s", now.Format("060102"), randomStringID))

	// you may change it as necessary ...

	return obj, nil
}

func (r AdminID) String() string {
	return string(r)
}
