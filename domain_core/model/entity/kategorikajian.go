package entity

import (
	"jadwalkajiansalaf/domain_core/model/vo"
	"time"
)

type KategoriKajian struct {
	ID      vo.KategoriKajianID `bson:"_id" json:"id"`
	Created time.Time           `bson:"created" json:"created"`
	Updated time.Time           `bson:"updated" json:"updated"`
	Nama    string
}

type KategoriKajianCreateRequest struct {
	RandomString string    `json:"-"`
	Now          time.Time `json:"-"`

	// edit or add new necessary field for create request here ...

}

func (r KategoriKajianCreateRequest) Validate() error {

	// validate the create request here ...

	return nil
}

func NewKategoriKajian(req KategoriKajianCreateRequest) (*KategoriKajian, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	id, err := vo.NewKategoriKajianID(req.RandomString, req.Now)
	if err != nil {
		return nil, err
	}

	var obj KategoriKajian
	obj.ID = id
	obj.Created = req.Now
	obj.Updated = req.Now

	// another field input here ...

	return &obj, nil
}

type KategoriKajianUpdateRequest struct {
	Now time.Time `json:"-"`

	// add new necessary field to update request here ...

}

func (r KategoriKajianUpdateRequest) Validate() error {

	// validate the update request here ...

	return nil
}

func (r *KategoriKajian) Update(req KategoriKajianUpdateRequest) error {

	err := req.Validate()
	if err != nil {
		return err
	}

	r.Updated = req.Now

	// update field here ...

	return nil
}
