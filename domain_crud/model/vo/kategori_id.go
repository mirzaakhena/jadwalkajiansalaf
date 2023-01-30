package vo

import (
	"fmt"
	"time"
)

type KategoriID string

func NewKategoriID(randomStringID string, now time.Time) (KategoriID, error) {
	var obj = KategoriID(fmt.Sprintf("OBJ-%s-%s", now.Format("060102"), randomStringID))

	// you may change it as necessary ...

	return obj, nil
}

func (r KategoriID) String() string {
	return string(r)
}
