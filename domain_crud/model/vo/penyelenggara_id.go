package vo

import (
	"fmt"
	"time"
)

type PenyelenggaraID string

func NewPenyelenggaraID(randomStringID string, now time.Time) (PenyelenggaraID, error) {
	var obj = PenyelenggaraID(fmt.Sprintf("OBJ-%s-%s", now.Format("060102"), randomStringID))

	// you may change it as necessary ...

	return obj, nil
}

func (r PenyelenggaraID) String() string {
	return string(r)
}
