package getoneadmin

import (
	"jadwalkajiansalaf/domain_core/model/entity"
	"jadwalkajiansalaf/domain_core/model/vo"
	"jadwalkajiansalaf/shared/gogen"
)

type Inport = gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	AdminID vo.AdminID
}

type InportResponse struct {
	Admin *entity.Admin
}
