package getonealamatkajian

import (
	"jadwalkajiansalaf/domain_core/model/entity"
	"jadwalkajiansalaf/domain_core/model/vo"
	"jadwalkajiansalaf/shared/gogen"
)

type Inport = gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	AlamatKajianID vo.AlamatKajianID
}

type InportResponse struct {
	AlamatKajian *entity.AlamatKajian
}
