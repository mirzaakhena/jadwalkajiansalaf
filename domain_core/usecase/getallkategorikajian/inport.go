package getallkategorikajian

import (
	"jadwalkajiansalaf/domain_core/model/repository"
	"jadwalkajiansalaf/shared/gogen"
)

type Inport = gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	repository.FindAllKategoriKajianFilterRequest
}

type InportResponse struct {
	Count int64
	Items []any
}
