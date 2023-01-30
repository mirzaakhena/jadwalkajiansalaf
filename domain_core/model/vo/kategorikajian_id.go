package vo

import (
	"fmt"
	"time"
)

type KategoriKajianID string

func NewKategoriKajianID(randomStringID string, now time.Time) (KategoriKajianID, error) {
	var obj = KategoriKajianID(fmt.Sprintf("OBJ-%s-%s", now.Format("060102"), randomStringID))

	// you may change it as necessary ...

	return obj, nil
}

func (r KategoriKajianID) String() string {
	return string(r)
}
