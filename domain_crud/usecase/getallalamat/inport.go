package getallalamat

import (
	"jadwalkajiansalaf/domain_crud/model/repository"
	"jadwalkajiansalaf/shared/gogen"
)

type Inport = gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	repository.FindAllAlamatFilterRequest
}

type InportResponse struct {
	Count int64
	Items []any
}
