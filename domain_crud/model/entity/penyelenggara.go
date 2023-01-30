package entity

import (
	"jadwalkajiansalaf/domain_crud/model/vo"
	"time"
)

type Penyelenggara struct {
	ID      vo.PenyelenggaraID `bson:"_id" json:"id"`
	Created time.Time          `bson:"created" json:"created"`
	Updated time.Time          `bson:"updated" json:"updated"`
	Nama    string
	Telepon string
	Alamat  string
}

type PenyelenggaraCreateRequest struct {
	RandomString string    `json:"-"`
	Now          time.Time `json:"-"`

	// edit or add new necessary field for create request here ...

}

func (r PenyelenggaraCreateRequest) Validate() error {

	// validate the create request here ...

	return nil
}

func NewPenyelenggara(req PenyelenggaraCreateRequest) (*Penyelenggara, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	id, err := vo.NewPenyelenggaraID(req.RandomString, req.Now)
	if err != nil {
		return nil, err
	}

	var obj Penyelenggara
	obj.ID = id
	obj.Created = req.Now
	obj.Updated = req.Now

	// another field input here ...

	return &obj, nil
}

type PenyelenggaraUpdateRequest struct {
	Now time.Time `json:"-"`

	// add new necessary field to update request here ...

}

func (r PenyelenggaraUpdateRequest) Validate() error {

	// validate the update request here ...

	return nil
}

func (r *Penyelenggara) Update(req PenyelenggaraUpdateRequest) error {

	err := req.Validate()
	if err != nil {
		return err
	}

	r.Updated = req.Now

	// update field here ...

	return nil
}
