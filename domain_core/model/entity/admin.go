package entity

import (
	"jadwalkajiansalaf/domain_core/model/vo"
	"time"
)

type Admin struct {
	ID       vo.AdminID `bson:"_id" json:"id"`
	Created  time.Time  `bson:"created" json:"created"`
	Updated  time.Time  `bson:"updated" json:"updated"`
	Nama     string
	Password []byte
	Alamat   string
	Telepon  string
	Email    string
}

type AdminCreateRequest struct {
	RandomString string    `json:"-"`
	Now          time.Time `json:"-"`
	Nama         string
	Password     []byte
	Alamat       string
	Telepon      string
	Email        string
}

func (r AdminCreateRequest) Validate() error {

	// validate the create request here ...

	return nil
}

func NewAdmin(req AdminCreateRequest) (*Admin, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	id, err := vo.NewAdminID(req.RandomString)
	if err != nil {
		return nil, err
	}

	var obj Admin
	obj.ID = id
	obj.Created = req.Now
	obj.Updated = req.Now

	// another field input here ...

	return &obj, nil
}

type AdminUpdateRequest struct {
	Now time.Time `json:"-"`

	// add new necessary field to update request here ...

}

func (r AdminUpdateRequest) Validate() error {

	// validate the update request here ...

	return nil
}

func (r *Admin) Update(req AdminUpdateRequest) error {

	err := req.Validate()
	if err != nil {
		return err
	}

	r.Updated = req.Now

	// update field here ...

	return nil
}
