package entity

import (
	"jadwalkajiansalaf/domain_crud/model/vo"
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

type Platform string

const (
	NamaPlatformYoutube    = Platform("Youtube")
	NamaPlatformInstagram  = Platform("Instagram")
	NamaPlatformTwitter    = Platform("Twitter")
	NamaPlatformTiktok     = Platform("Tiktok")
	NamaPlatformFacebook   = Platform("Facebook")
	NamaPlatformGoogleMeet = Platform("GoogleMeet")
	NamaPlatformZoom       = Platform("Zoom")
)

type MediaSosial struct {
	Platform Platform `json:"platform"`
	Alamat   string   `json:"alamat"`
}

type Kajian struct {
	ID            vo.KajianID    `bson:"_id" json:"id"`
	Created       time.Time      `bson:"created" json:"created"`
	Updated       time.Time      `bson:"updated" json:"updated"`
	Judul         string         `json:"judul" bson:"judul"`
	Deskripsi     string         `json:"deskripsi,omitempty" bson:"deskripsi"`
	CatatanKhusus []string       `json:"catatan_khusus" bson:"catatan_khusus"`
	Pemateri      string         `json:"pemateri" bson:"pemateri"`
	Tanggal       time.Time      `json:"tanggal,omitempty" bson:"tanggal"`
	Hari          []time.Weekday `json:"hari,omitempty" bson:"hari"`
	WaktuMulai    time.Time      `json:"waktu_mulai" bson:"waktu_mulai"`
	WaktuSelesai  time.Time      `json:"waktu_selesai" bson:"waktu_selesai"`
	LokasiID      vo.LokasiID    `json:"lokasi_id" bson:"lokasi_id"`
	Lokasi        *Lokasi        `json:"lokasi" bson:"lokasi"`
	Penyelenggara []string       `json:"penyelenggara,omitempty" bson:"penyelenggara"`
	LinkStreaming []MediaSosial  `json:"link_streaming,omitempty" bson:"link_streaming"`
	PosterURL     string         `json:"poster_url,omitempty" bson:"poster_url"`
	JenisJamaah   JenisJamaah    `json:"jenis_jamaah" bson:"jenis_jamaah"`
	Kategori      string         `json:"kategori,omitempty" bson:"kategori"`
	MediaSosial   []MediaSosial  `json:"media_sosial" bson:"media_sosial"`
	StatusKajian  StatusKajian   `json:"status_kajian" bson:"status_kajian"`
}

type KajianCreateRequest struct {
	RandomString string    `json:"-"`
	Now          time.Time `json:"-"`

	// edit or add new necessary field for create request here ...

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
