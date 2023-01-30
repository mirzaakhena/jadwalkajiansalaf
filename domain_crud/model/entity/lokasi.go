package entity

import (
	"jadwalkajiansalaf/domain_crud/model/vo"
	"time"
)

type Lokasi struct {
	ID         vo.LokasiID `json:"id,omitempty" bson:"_id"`
	Created    time.Time   `json:"created" bson:"created"`
	Updated    time.Time   `json:"updated" bson:"updated"`
	NamaTempat string      `json:"nama_tempat,omitempty" bson:"nama_tempat"`
	Alamat     string      `json:"alamat,omitempty" bson:"alamat"`
	Koordinat  [2]float64  `json:"koordinat,omitempty" bson:"koordinat"`
}

type LokasiCreateRequest struct {
	RandomString string    `json:"-"`
	Now          time.Time `json:"-"`

	// edit or add new necessary field for create request here ...

}

func (r LokasiCreateRequest) Validate() error {

	// validate the create request here ...

	return nil
}

func NewLokasi(req LokasiCreateRequest) (*Lokasi, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	id, err := vo.NewLokasiID(req.RandomString, req.Now)
	if err != nil {
		return nil, err
	}

	var obj Lokasi
	obj.ID = id
	obj.Created = req.Now
	obj.Updated = req.Now

	// another field input here ...

	return &obj, nil
}

type LokasiUpdateRequest struct {
	Now time.Time `json:"-"`

	// add new necessary field to update request here ...

}

func (r LokasiUpdateRequest) Validate() error {

	// validate the update request here ...

	return nil
}

func (r *Lokasi) Update(req LokasiUpdateRequest) error {

	err := req.Validate()
	if err != nil {
		return err
	}

	r.Updated = req.Now

	// update field here ...

	return nil
}
