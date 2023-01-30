package entity

import (
	"jadwalkajiansalaf/domain_crud/model/vo"
	"time"
)

type Kategori struct {
	ID      vo.KategoriID `bson:"_id" json:"id"`
	Created time.Time     `bson:"created" json:"created"`
	Updated time.Time     `bson:"updated" json:"updated"`
	Nama    string
}

type KategoriCreateRequest struct {
	RandomString string    `json:"-"`
	Now          time.Time `json:"-"`

	// edit or add new necessary field for create request here ...

}

func (r KategoriCreateRequest) Validate() error {

	// validate the create request here ...

	return nil
}

func NewKategori(req KategoriCreateRequest) (*Kategori, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	id, err := vo.NewKategoriID(req.RandomString, req.Now)
	if err != nil {
		return nil, err
	}

	var obj Kategori
	obj.ID = id
	obj.Created = req.Now
	obj.Updated = req.Now

	// another field input here ...

	return &obj, nil
}

type KategoriUpdateRequest struct {
	Now time.Time `json:"-"`

	// add new necessary field to update request here ...

}

func (r KategoriUpdateRequest) Validate() error {

	// validate the update request here ...

	return nil
}

func (r *Kategori) Update(req KategoriUpdateRequest) error {

	err := req.Validate()
	if err != nil {
		return err
	}

	r.Updated = req.Now

	// update field here ...

	return nil
}
