package entity

import (
	"jadwalkajiansalaf/domain_crud/model/errorenum"
	"jadwalkajiansalaf/domain_crud/model/vo"
	"time"
)

type Admin struct {
	ID       vo.AdminID `json:"id,omitempty" bson:"_id"`
	Created  time.Time  `json:"created" bson:"created"`
	Updated  time.Time  `json:"updated" bson:"updated"`
	Nama     string     `json:"nama,omitempty" bson:"nama"`
	Password []byte     `json:"password,omitempty" bson:"password"`
	Alamat   string     `json:"alamat,omitempty" bson:"alamat"`
	Telepon  string     `json:"telepon,omitempty" bson:"telepon"`
	Email    string     `json:"email,omitempty" bson:"email"`
}

type AdminCreateRequest struct {
	RandomString string    `json:"-"`
	Now          time.Time `json:"-"`
	Nama         string    `json:"nama"`
}

func (r AdminCreateRequest) Validate() error {

	// validate the create request here ...
	if r.Nama == "" {
		return errorenum.NamaAdminTidakBolehKosong
	}

	return nil
}

func NewAdmin(req AdminCreateRequest) (*Admin, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	id, err := vo.NewAdminID(req.RandomString, req.Now)
	if err != nil {
		return nil, err
	}

	var obj Admin
	obj.ID = id
	obj.Created = req.Now
	obj.Updated = req.Now
	obj.Nama = req.Nama

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
