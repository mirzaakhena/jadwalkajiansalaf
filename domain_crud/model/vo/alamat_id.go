package vo

import (
	"fmt"
	"time"
)

type LokasiID string

func NewLokasiID(randomStringID string, now time.Time) (LokasiID, error) {
	var obj = LokasiID(fmt.Sprintf("OBJ-%s-%s", now.Format("060102"), randomStringID))

	// you may change it as necessary ...

	return obj, nil
}

func (r LokasiID) String() string {
	return string(r)
}
