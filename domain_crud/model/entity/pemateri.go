package entity

import (
	"jadwalkajiansalaf/domain_crud/model/vo"
	"time"
)

type Pemateri struct {
	ID      vo.PemateriID `bson:"_id" json:"id"`
	Created time.Time     `bson:"created" json:"created"`
	Updated time.Time     `bson:"updated" json:"updated"`
	Nama    string
}

type PemateriCreateRequest struct {
	RandomString string    `json:"-"`
	Now          time.Time `json:"-"`

	// edit or add new necessary field for create request here ...

}

func (r PemateriCreateRequest) Validate() error {

	// validate the create request here ...

	return nil
}

func NewPemateri(req PemateriCreateRequest) (*Pemateri, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	id, err := vo.NewPemateriID(req.RandomString, req.Now)
	if err != nil {
		return nil, err
	}

	var obj Pemateri
	obj.ID = id
	obj.Created = req.Now
	obj.Updated = req.Now

	// another field input here ...

	return &obj, nil
}

type PemateriUpdateRequest struct {
	Now time.Time `json:"-"`

	// add new necessary field to update request here ...

}

func (r PemateriUpdateRequest) Validate() error {

	// validate the update request here ...

	return nil
}

func (r *Pemateri) Update(req PemateriUpdateRequest) error {

	err := req.Validate()
	if err != nil {
		return err
	}

	r.Updated = req.Now

	// update field here ...

	return nil
}
