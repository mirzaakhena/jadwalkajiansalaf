package restapi

func (r *controller) RegisterRouter() {

	resource := r.Router.Group("/api/v1", r.authentication())

	resource.GET("/getalladmin", r.authorization(), r.getAllAdminHandler())
	resource.GET("/getallalamatkajian", r.authorization(), r.getAllAlamatKajianHandler())
	resource.GET("/getallkajian", r.authorization(), r.getAllKajianHandler())
	resource.GET("/getallkategorikajian", r.authorization(), r.getAllKategoriKajianHandler())
	resource.GET("/getallpemateri", r.authorization(), r.getAllPemateriHandler())
	resource.GET("/getallpenyelenggarakajian", r.authorization(), r.getAllPenyelenggaraKajianHandler())
	resource.GET("/getoneadmin", r.authorization(), r.getOneAdminHandler())
	resource.GET("/getonealamatkajian", r.authorization(), r.getOneAlamatKajianHandler())
	resource.GET("/getonekajian", r.authorization(), r.getOneKajianHandler())
	resource.GET("/getonekategorikajian", r.authorization(), r.getOneKategoriKajianHandler())
	resource.GET("/getonepemateri", r.authorization(), r.getOnePemateriHandler())
	resource.GET("/getonepenyelenggarakajian", r.authorization(), r.getOnePenyelenggaraKajianHandler())
	resource.POST("/runadmincreate", r.authorization(), r.runAdminCreateHandler())
	resource.POST("/runadmindelete", r.authorization(), r.runAdminDeleteHandler())
	resource.POST("/runadminupdate", r.authorization(), r.runAdminUpdateHandler())
	resource.POST("/runalamatkajiancreate", r.authorization(), r.runAlamatKajianCreateHandler())
	resource.POST("/runalamatkajiandelete", r.authorization(), r.runAlamatKajianDeleteHandler())
	resource.POST("/runalamatkajianupdate", r.authorization(), r.runAlamatKajianUpdateHandler())
	resource.POST("/runkajiancreate", r.authorization(), r.runKajianCreateHandler())
	resource.POST("/runkajiandelete", r.authorization(), r.runKajianDeleteHandler())
	resource.POST("/runkajianupdate", r.authorization(), r.runKajianUpdateHandler())
	resource.POST("/runkategorikajiancreate", r.authorization(), r.runKategoriKajianCreateHandler())
	resource.POST("/runkategorikajiandelete", r.authorization(), r.runKategoriKajianDeleteHandler())
	resource.POST("/runkategorikajianupdate", r.authorization(), r.runKategoriKajianUpdateHandler())
	resource.POST("/runpematericreate", r.authorization(), r.runPemateriCreateHandler())
	resource.POST("/runpemateridelete", r.authorization(), r.runPemateriDeleteHandler())
	resource.POST("/runpemateriupdate", r.authorization(), r.runPemateriUpdateHandler())
	resource.POST("/runpenyelenggarakajiancreate", r.authorization(), r.runPenyelenggaraKajianCreateHandler())
	resource.POST("/runpenyelenggarakajiandelete", r.authorization(), r.runPenyelenggaraKajianDeleteHandler())
	resource.POST("/runpenyelenggarakajianupdate", r.authorization(), r.runPenyelenggaraKajianUpdateHandler())
}
