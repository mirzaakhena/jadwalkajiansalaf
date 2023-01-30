package entity

import (
	"jadwalkajiansalaf/domain_core/model/vo"
	"time"
)

type StatusKajian string

const (
	StatusKajianDraft    = StatusKajian("DRAFT")
	StatusKajianAktif    = StatusKajian("AKTIF")
	StatusKajianNonAktif = StatusKajian("NON_AKTIF")
	StatusKajianBatal    = StatusKajian("BATAL")
)

type JenisJamaah string

const (
	JenisJamaahIkhwan = JenisJamaah("IKHWAN")
	JenisJamaahAkhwat = JenisJamaah("AKHWAT")
	JenisJamaahUmum   = JenisJamaah("UMUM")
)

type NamaPlatform string

const (
	NamaPlatformYoutube    = NamaPlatform("Youtube")
	NamaPlatformInstagram  = NamaPlatform("Instagram")
	NamaPlatformTwitter    = NamaPlatform("Twitter")
	NamaPlatformTiktok     = NamaPlatform("Tiktok")
	NamaPlatformFacebook   = NamaPlatform("Facebook")
	NamaPlatformGoogleMeet = NamaPlatform("GoogleMeet")
	NamaPlatformZoom       = NamaPlatform("Zoom")
)

type Platform struct {
	Nama      NamaPlatform
	AlamatURL string
}

type Kajian struct {
	ID            vo.KajianID    `json:"id" bson:"_id"`
	Created       time.Time      `json:"created" bson:"created"`
	Updated       time.Time      `json:"updated" bson:"updated"`
	Judul         string         `json:"judul,omitempty" bson:"judul"`
	Deskripsi     string         `json:"deskripsi" bson:"deskripsi"`
	CatatanKhusus []string       `json:"catatan_khusus" bson:"catatan_khusus"`
	Pemateri      string         `json:"pemateri,omitempty" bson:"pemateri"`
	Tanggal       time.Time      `json:"tanggal" bson:"tanggal"`
	Hari          []time.Weekday `json:"hari,omitempty" bson:"hari"`
	WaktuMulai    time.Time      `json:"waktu_mulai" bson:"waktu_mulai"`
	WaktuSelesai  time.Time      `json:"waktu_selesai" bson:"waktu_selesai"`
	Lokasi        AlamatKajian   `json:"lokasi" bson:"lokasi"`
	Penyelenggara []string       `json:"penyelenggara,omitempty" bson:"penyelenggara"`
	LinkStreaming []Platform     `json:"link_streaming,omitempty" bson:"link_streaming"`
	PosterURL     string         `json:"poster_url,omitempty" bson:"poster_url"`
	JenisJamaah   JenisJamaah    `json:"jenis_jamaah" bson:"jenis_jamaah"`
	Kategori      string         `json:"kategori,omitempty" bson:"kategori"`
	MediaSosial   []Platform     `json:"media_sosial" bson:"media_sosial"`
	StatusKajian  StatusKajian   `json:"status_kajian" bson:"status_kajian"`
}

type KajianCreateRequest struct {
	RandomString string    `json:"-"`
	Now          time.Time `json:"-"`
}

func (r KajianCreateRequest) Validate() error {

	// validate the create request here ...

	return nil
}

func NewKajian(req KajianCreateRequest) (*Kajian, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	id, err := vo.NewKajianID(req.RandomString, req.Now)
	if err != nil {
		return nil, err
	}

	var obj Kajian
	obj.ID = id
	obj.Created = req.Now
	obj.Updated = req.Now

	// another field input here ...

	return &obj, nil
}

type KajianUpdateRequest struct {
	Now time.Time `json:"-"`

	// add new necessary field to update request here ...

}

func (r KajianUpdateRequest) Validate() error {

	// validate the update request here ...

	return nil
}

func (r *Kajian) Update(req KajianUpdateRequest) error {

	err := req.Validate()
	if err != nil {
		return err
	}

	r.Updated = req.Now

	// update field here ...

	return nil
}
