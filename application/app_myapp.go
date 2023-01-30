package application

import (
	"jadwalkajiansalaf/domain_core/controller/restapi"
	"jadwalkajiansalaf/domain_core/gateway/withgormdb"
	"jadwalkajiansalaf/domain_core/usecase/getalladmin"
	"jadwalkajiansalaf/domain_core/usecase/getallalamatkajian"
	"jadwalkajiansalaf/domain_core/usecase/getallkajian"
	"jadwalkajiansalaf/domain_core/usecase/getallkategorikajian"
	"jadwalkajiansalaf/domain_core/usecase/getallpemateri"
	"jadwalkajiansalaf/domain_core/usecase/getallpenyelenggarakajian"
	"jadwalkajiansalaf/domain_core/usecase/getoneadmin"
	"jadwalkajiansalaf/domain_core/usecase/getonealamatkajian"
	"jadwalkajiansalaf/domain_core/usecase/getonekajian"
	"jadwalkajiansalaf/domain_core/usecase/getonekategorikajian"
	"jadwalkajiansalaf/domain_core/usecase/getonepemateri"
	"jadwalkajiansalaf/domain_core/usecase/getonepenyelenggarakajian"
	"jadwalkajiansalaf/domain_core/usecase/runadmincreate"
	"jadwalkajiansalaf/domain_core/usecase/runadmindelete"
	"jadwalkajiansalaf/domain_core/usecase/runadminupdate"
	"jadwalkajiansalaf/domain_core/usecase/runalamatkajiancreate"
	"jadwalkajiansalaf/domain_core/usecase/runalamatkajiandelete"
	"jadwalkajiansalaf/domain_core/usecase/runalamatkajianupdate"
	"jadwalkajiansalaf/domain_core/usecase/runkajiancreate"
	"jadwalkajiansalaf/domain_core/usecase/runkajiandelete"
	"jadwalkajiansalaf/domain_core/usecase/runkajianupdate"
	"jadwalkajiansalaf/domain_core/usecase/runkategorikajiancreate"
	"jadwalkajiansalaf/domain_core/usecase/runkategorikajiandelete"
	"jadwalkajiansalaf/domain_core/usecase/runkategorikajianupdate"
	"jadwalkajiansalaf/domain_core/usecase/runpematericreate"
	"jadwalkajiansalaf/domain_core/usecase/runpemateridelete"
	"jadwalkajiansalaf/domain_core/usecase/runpemateriupdate"
	"jadwalkajiansalaf/domain_core/usecase/runpenyelenggarakajiancreate"
	"jadwalkajiansalaf/domain_core/usecase/runpenyelenggarakajiandelete"
	"jadwalkajiansalaf/domain_core/usecase/runpenyelenggarakajianupdate"
	"jadwalkajiansalaf/shared/config"
	"jadwalkajiansalaf/shared/gogen"
	"jadwalkajiansalaf/shared/infrastructure/logger"
	"jadwalkajiansalaf/shared/infrastructure/token"
)

type myApp struct{}

func NewMyApp() gogen.Runner {
	return &myApp{}
}

func (myApp) Run() error {

	const appName = "myApp"

	cfg := config.ReadConfig()

	appData := gogen.NewApplicationData(appName)

	log := logger.NewSimpleJSONLogger(appData)

	jwtToken := token.NewJWTToken(cfg.JWTSecretKey)

	datasource := withgormdb.NewGateway(log, appData, cfg)

	primaryDriver := restapi.NewController(appData, log, cfg, jwtToken)

	primaryDriver.AddUsecase(
		//
		getalladmin.NewUsecase(datasource),
		getallalamatkajian.NewUsecase(datasource),
		getallkajian.NewUsecase(datasource),
		getallkategorikajian.NewUsecase(datasource),
		getallpemateri.NewUsecase(datasource),
		getallpenyelenggarakajian.NewUsecase(datasource),
		getoneadmin.NewUsecase(datasource),
		getonealamatkajian.NewUsecase(datasource),
		getonekajian.NewUsecase(datasource),
		getonekategorikajian.NewUsecase(datasource),
		getonepemateri.NewUsecase(datasource),
		getonepenyelenggarakajian.NewUsecase(datasource),
		runadmincreate.NewUsecase(datasource),
		runadmindelete.NewUsecase(datasource),
		runadminupdate.NewUsecase(datasource),
		runalamatkajiancreate.NewUsecase(datasource),
		runalamatkajiandelete.NewUsecase(datasource),
		runalamatkajianupdate.NewUsecase(datasource),
		runkajiancreate.NewUsecase(datasource),
		runkajiandelete.NewUsecase(datasource),
		runkajianupdate.NewUsecase(datasource),
		runkategorikajiancreate.NewUsecase(datasource),
		runkategorikajiandelete.NewUsecase(datasource),
		runkategorikajianupdate.NewUsecase(datasource),
		runpematericreate.NewUsecase(datasource),
		runpemateridelete.NewUsecase(datasource),
		runpemateriupdate.NewUsecase(datasource),
		runpenyelenggarakajiancreate.NewUsecase(datasource),
		runpenyelenggarakajiandelete.NewUsecase(datasource),
		runpenyelenggarakajianupdate.NewUsecase(datasource),
	)

	primaryDriver.RegisterRouter()

	primaryDriver.Start()

	return nil
}
