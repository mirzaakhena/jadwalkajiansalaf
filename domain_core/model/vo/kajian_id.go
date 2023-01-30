package vo

import (
	"fmt"
	"time"
)

type KajianID string

func NewKajianID(randomStringID string, now time.Time) (KajianID, error) {
	var obj = KajianID(fmt.Sprintf("JKS-%s-%s", now.Format("060102"), randomStringID))

	// you may change it as necessary ...

	return obj, nil
}

func (r KajianID) String() string {
	return string(r)
}
