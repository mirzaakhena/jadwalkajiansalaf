package getonepenyelenggarakajian

import (
	"jadwalkajiansalaf/domain_core/model/entity"
	"jadwalkajiansalaf/domain_core/model/vo"
	"jadwalkajiansalaf/shared/gogen"
)

type Inport = gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	PenyelenggaraKajianID vo.PenyelenggaraKajianID
}

type InportResponse struct {
	PenyelenggaraKajian *entity.PenyelenggaraKajian
}
