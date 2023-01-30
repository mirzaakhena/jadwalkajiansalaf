package restapi

import (
	"context"
	"jadwalkajiansalaf/domain_core/model/entity"
	"jadwalkajiansalaf/domain_core/model/vo"
	"jadwalkajiansalaf/domain_core/usecase/getoneadmin"
	"jadwalkajiansalaf/shared/gogen"
	"jadwalkajiansalaf/shared/infrastructure/logger"
	"jadwalkajiansalaf/shared/model/payload"
	"jadwalkajiansalaf/shared/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *controller) getOneAdminHandler() gin.HandlerFunc {

	type InportRequest = getoneadmin.InportRequest
	type InportResponse = getoneadmin.InportResponse

	inport := gogen.GetInport[InportRequest, InportResponse](r.GetUsecase(InportRequest{}))

	type request struct {
		AdminID vo.AdminID `form:"admin_id,omitempty,default=0"`
	}

	type response struct {
		Admin *entity.Admin `json:"admin"`
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
		req.AdminID = jsonReq.AdminID

		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		jsonRes.Admin = res.Admin

		r.log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
