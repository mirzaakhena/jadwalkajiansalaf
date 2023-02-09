package application

import (
	"jadwalkajiansalaf/domain_crud/controller/restapi"
	"jadwalkajiansalaf/domain_crud/gateway/withgormdb2"
	"jadwalkajiansalaf/domain_crud/usecase/getalladmin"
	"jadwalkajiansalaf/domain_crud/usecase/getallalamat"
	"jadwalkajiansalaf/domain_crud/usecase/getallkajian"
	"jadwalkajiansalaf/domain_crud/usecase/getallkategori"
	"jadwalkajiansalaf/domain_crud/usecase/getallpemateri"
	"jadwalkajiansalaf/domain_crud/usecase/getallpenyelenggara"
	"jadwalkajiansalaf/domain_crud/usecase/getoneadmin"
	"jadwalkajiansalaf/domain_crud/usecase/getonealamat"
	"jadwalkajiansalaf/domain_crud/usecase/getonekajian"
	"jadwalkajiansalaf/domain_crud/usecase/getonekategori"
	"jadwalkajiansalaf/domain_crud/usecase/getonepemateri"
	"jadwalkajiansalaf/domain_crud/usecase/getonepenyelenggara"
	"jadwalkajiansalaf/domain_crud/usecase/runadmincreate"
	"jadwalkajiansalaf/domain_crud/usecase/runadmindelete"
	"jadwalkajiansalaf/domain_crud/usecase/runadminupdate"
	"jadwalkajiansalaf/domain_crud/usecase/runalamatcreate"
	"jadwalkajiansalaf/domain_crud/usecase/runalamatdelete"
	"jadwalkajiansalaf/domain_crud/usecase/runalamatupdate"
	"jadwalkajiansalaf/domain_crud/usecase/runkajiancreate"
	"jadwalkajiansalaf/domain_crud/usecase/runkajiandelete"
	"jadwalkajiansalaf/domain_crud/usecase/runkajianupdate"
	"jadwalkajiansalaf/domain_crud/usecase/runkategoricreate"
	"jadwalkajiansalaf/domain_crud/usecase/runkategoridelete"
	"jadwalkajiansalaf/domain_crud/usecase/runkategoriupdate"
	"jadwalkajiansalaf/domain_crud/usecase/runpematericreate"
	"jadwalkajiansalaf/domain_crud/usecase/runpemateridelete"
	"jadwalkajiansalaf/domain_crud/usecase/runpemateriupdate"
	"jadwalkajiansalaf/domain_crud/usecase/runpenyelenggaracreate"
	"jadwalkajiansalaf/domain_crud/usecase/runpenyelenggaradelete"
	"jadwalkajiansalaf/domain_crud/usecase/runpenyelenggaraupdate"
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

	const appName = "myapp"

	cfg := config.ReadConfig()

	appData := gogen.NewApplicationData(appName)

	log := logger.NewSimpleJSONLogger(appData)

	jwtToken := token.NewJWTToken(cfg.JWTSecretKey)

	datasource := withgormdb2.NewGateway(log, appData, cfg)

	primaryDriver := restapi.NewController(appData, log, cfg, jwtToken)

	primaryDriver.AddUsecase(
		//
		runadmincreate.NewUsecase(datasource),
		getalladmin.NewUsecase(datasource),
		getoneadmin.NewUsecase(datasource),
		runadminupdate.NewUsecase(datasource),
		runadmindelete.NewUsecase(datasource),
		runalamatcreate.NewUsecase(datasource),
		getallalamat.NewUsecase(datasource),
		getonealamat.NewUsecase(datasource),
		runalamatupdate.NewUsecase(datasource),
		runalamatdelete.NewUsecase(datasource),
		runkajiancreate.NewUsecase(datasource),
		getallkajian.NewUsecase(datasource),
		getonekajian.NewUsecase(datasource),
		runkajianupdate.NewUsecase(datasource),
		runkajiandelete.NewUsecase(datasource),
		runkategoricreate.NewUsecase(datasource),
		getallkategori.NewUsecase(datasource),
		getonekategori.NewUsecase(datasource),
		runkategoriupdate.NewUsecase(datasource),
		runkategoridelete.NewUsecase(datasource),
		runpematericreate.NewUsecase(datasource),
		getallpemateri.NewUsecase(datasource),
		getonepemateri.NewUsecase(datasource),
		runpemateriupdate.NewUsecase(datasource),
		runpemateridelete.NewUsecase(datasource),
		runpenyelenggaracreate.NewUsecase(datasource),
		getallpenyelenggara.NewUsecase(datasource),
		getonepenyelenggara.NewUsecase(datasource),
		runpenyelenggaraupdate.NewUsecase(datasource),
		runpenyelenggaradelete.NewUsecase(datasource),
	)

	primaryDriver.RegisterRouter()

	primaryDriver.Start()

	return nil
}
