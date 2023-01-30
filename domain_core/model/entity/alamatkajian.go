package entity

import (
	"jadwalkajiansalaf/domain_core/model/vo"
	"time"
)

type AlamatKajian struct {
	ID         vo.AlamatKajianID `bson:"_id" json:"id"`
	Created    time.Time         `bson:"created" json:"created"`
	Updated    time.Time         `bson:"updated" json:"updated"`
	NamaTempat string
	Alamat     string
	Koordinat  [2]float64
}

type AlamatKajianCreateRequest struct {
	RandomString string    `json:"-"`
	Now          time.Time `json:"-"`

	// edit or add new necessary field for create request here ...

}

func (r AlamatKajianCreateRequest) Validate() error {

	// validate the create request here ...

	return nil
}

func NewAlamatKajian(req AlamatKajianCreateRequest) (*AlamatKajian, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	id, err := vo.NewAlamatKajianID(req.RandomString, req.Now)
	if err != nil {
		return nil, err
	}

	var obj AlamatKajian
	obj.ID = id
	obj.Created = req.Now
	obj.Updated = req.Now

	// another field input here ...

	return &obj, nil
}

type AlamatKajianUpdateRequest struct {
	Now time.Time `json:"-"`

	// add new necessary field to update request here ...

}

func (r AlamatKajianUpdateRequest) Validate() error {

	// validate the update request here ...

	return nil
}

func (r *AlamatKajian) Update(req AlamatKajianUpdateRequest) error {

	err := req.Validate()
	if err != nil {
		return err
	}

	r.Updated = req.Now

	// update field here ...

	return nil
}
