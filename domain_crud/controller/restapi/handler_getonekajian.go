package restapi

import (
	"context"
	"jadwalkajiansalaf/domain_crud/model/vo"
	"jadwalkajiansalaf/domain_crud/usecase/getonekajian"
	"jadwalkajiansalaf/shared/gogen"
	"jadwalkajiansalaf/shared/infrastructure/logger"
	"jadwalkajiansalaf/shared/model/payload"
	"jadwalkajiansalaf/shared/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *controller) getOneKajianHandler() gin.HandlerFunc {

	type InportRequest = getonekajian.InportRequest
	type InportResponse = getonekajian.InportResponse

	inport := gogen.GetInport[InportRequest, InportResponse](r.GetUsecase(InportRequest{}))

	type request struct {
	}

	type response struct {
		*InportResponse
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		//var jsonReq request
		//if err := c.BindJSON(&jsonReq); err != nil {
		//	r.log.Error(ctx, err.Error())
		//	c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
		//	return
		//}

		var req InportRequest
		req.KajianID = vo.KajianID(c.Param("kajian_id"))

		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		jsonRes.InportResponse = res

		r.log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
