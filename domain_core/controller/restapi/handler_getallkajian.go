package restapi

import (
	"context"
	"jadwalkajiansalaf/domain_core/usecase/getallkajian"
	"jadwalkajiansalaf/shared/gogen"
	"jadwalkajiansalaf/shared/infrastructure/logger"
	"jadwalkajiansalaf/shared/model/payload"
	"jadwalkajiansalaf/shared/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *controller) getAllKajianHandler() gin.HandlerFunc {

	type InportRequest = getallkajian.InportRequest
	type InportResponse = getallkajian.InportResponse

	inport := gogen.GetInport[InportRequest, InportResponse](r.GetUsecase(InportRequest{}))

	type request struct {
	}

	type response struct {
		Count int64 `json:"count"`
		Items []any `json:"items"`
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

		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		jsonRes.Count = res.Count
		jsonRes.Items = res.Items

		r.log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
