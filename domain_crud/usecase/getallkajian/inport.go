package getallkajian

import (
	"jadwalkajiansalaf/domain_crud/model/repository"
	"jadwalkajiansalaf/shared/gogen"
)

type Inport = gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	repository.FindAllKajianFilterRequest
}

type InportResponse struct {
	Count int64
	Items []any
}
