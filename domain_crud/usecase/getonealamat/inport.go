package getonealamat

import (
	"jadwalkajiansalaf/domain_crud/model/entity"
	"jadwalkajiansalaf/domain_crud/model/vo"
	"jadwalkajiansalaf/shared/gogen"
)

type Inport = gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	AlamatID vo.LokasiID
}

type InportResponse struct {
	Alamat *entity.Lokasi
}
