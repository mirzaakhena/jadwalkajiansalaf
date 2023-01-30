package restapi

import (
	"context"
	"jadwalkajiansalaf/domain_core/model/entity"
	"jadwalkajiansalaf/domain_core/model/vo"
	"jadwalkajiansalaf/domain_core/usecase/getonekategorikajian"
	"jadwalkajiansalaf/shared/gogen"
	"jadwalkajiansalaf/shared/infrastructure/logger"
	"jadwalkajiansalaf/shared/model/payload"
	"jadwalkajiansalaf/shared/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *controller) getOneKategoriKajianHandler() gin.HandlerFunc {

	type InportRequest = getonekategorikajian.InportRequest
	type InportResponse = getonekategorikajian.InportResponse

	inport := gogen.GetInport[InportRequest, InportResponse](r.GetUsecase(InportRequest{}))

	type request struct {
		KategoriKajianID vo.KategoriKajianID `form:"kategori_kajian_id,omitempty,default=0"`
	}

	type response struct {
		KategoriKajian *entity.KategoriKajian `json:"kategori_kajian"`
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		var jsonReq request
		err := c.Bind(&jsonReq)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var req InportRequest
		req.KategoriKajianID = jsonReq.KategoriKajianID

		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		jsonRes.KategoriKajian = res.KategoriKajian

		r.log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
