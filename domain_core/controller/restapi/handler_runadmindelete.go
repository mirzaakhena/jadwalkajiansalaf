package restapi

import (
	"context"
	"jadwalkajiansalaf/domain_core/model/vo"
	"jadwalkajiansalaf/domain_core/usecase/runadmindelete"
	"jadwalkajiansalaf/shared/gogen"
	"jadwalkajiansalaf/shared/infrastructure/logger"
	"jadwalkajiansalaf/shared/model/payload"
	"jadwalkajiansalaf/shared/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *controller) runAdminDeleteHandler() gin.HandlerFunc {

	type InportRequest = runadmindelete.InportRequest
	type InportResponse = runadmindelete.InportResponse

	inport := gogen.GetInport[InportRequest, InportResponse](r.GetUsecase(InportRequest{}))

	type request struct {
		AdminID vo.AdminID `json:"admin_id"`
	}

	type response struct {
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		var jsonReq request
		err := c.BindJSON(&jsonReq)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var req InportRequest
		req.AdminID = jsonReq.AdminID

		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		_ = res

		r.log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
