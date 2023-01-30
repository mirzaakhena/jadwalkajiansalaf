package entity

import (
	"jadwalkajiansalaf/domain_core/model/vo"
	"time"
)

type PenyelenggaraKajian struct {
	ID      vo.PenyelenggaraKajianID `bson:"_id" json:"id"`
	Created time.Time                `bson:"created" json:"created"`
	Updated time.Time                `bson:"updated" json:"updated"`
	Nama    string
	Telepon string
	Alamat  string
}

type PenyelenggaraKajianCreateRequest struct {
	RandomString string    `json:"-"`
	Now          time.Time `json:"-"`

	// edit or add new necessary field for create request here ...

}

func (r PenyelenggaraKajianCreateRequest) Validate() error {

	// validate the create request here ...

	return nil
}

func NewPenyelenggaraKajian(req PenyelenggaraKajianCreateRequest) (*PenyelenggaraKajian, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	id, err := vo.NewPenyelenggaraKajianID(req.RandomString, req.Now)
	if err != nil {
		return nil, err
	}

	var obj PenyelenggaraKajian
	obj.ID = id
	obj.Created = req.Now
	obj.Updated = req.Now

	// another field input here ...

	return &obj, nil
}

type PenyelenggaraKajianUpdateRequest struct {
	Now time.Time `json:"-"`

	// add new necessary field to update request here ...

}

func (r PenyelenggaraKajianUpdateRequest) Validate() error {

	// validate the update request here ...

	return nil
}

func (r *PenyelenggaraKajian) Update(req PenyelenggaraKajianUpdateRequest) error {

	err := req.Validate()
	if err != nil {
		return err
	}

	r.Updated = req.Now

	// update field here ...

	return nil
}
