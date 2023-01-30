package vo

import (
	"fmt"
	"time"
)

type PenyelenggaraKajianID string

func NewPenyelenggaraKajianID(randomStringID string, now time.Time) (PenyelenggaraKajianID, error) {
	var obj = PenyelenggaraKajianID(fmt.Sprintf("OBJ-%s-%s", now.Format("060102"), randomStringID))

	// you may change it as necessary ...

	return obj, nil
}

func (r PenyelenggaraKajianID) String() string {
	return string(r)
}
