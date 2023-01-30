package vo

import (
	"fmt"
	"time"
)

type AlamatKajianID string

func NewAlamatKajianID(randomStringID string, now time.Time) (AlamatKajianID, error) {
	var obj = AlamatKajianID(fmt.Sprintf("OBJ-%s-%s", now.Format("060102"), randomStringID))

	// you may change it as necessary ...

	return obj, nil
}

func (r AlamatKajianID) String() string {
	return string(r)
}
