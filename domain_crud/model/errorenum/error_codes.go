package errorenum

import "jadwalkajiansalaf/shared/model/apperror"

const (
	SomethingError            apperror.ErrorType = "ER0000 something error"
	NamaAdminTidakBolehKosong apperror.ErrorType = "ER0001 nama admin tidak boleh kosong"
)
