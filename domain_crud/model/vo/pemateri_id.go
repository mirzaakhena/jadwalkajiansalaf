package vo

import (
	"fmt"
	"time"
)

type PemateriID string

func NewPemateriID(randomStringID string, now time.Time) (PemateriID, error) {
	var obj = PemateriID(fmt.Sprintf("OBJ-%s-%s", now.Format("060102"), randomStringID))

	// you may change it as necessary ...

	return obj, nil
}

func (r PemateriID) String() string {
	return string(r)
}
