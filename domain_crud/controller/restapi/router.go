package restapi

func (r *controller) RegisterRouter() {

	resource := r.Router.Group("/api/v1", r.authentication())

	resource.POST("/admin", r.authorization(), r.runAdminCreateHandler())
	resource.GET("/admin", r.authorization(), r.getAllAdminHandler())
	resource.GET("/admin/:admin_id", r.authorization(), r.getOneAdminHandler())
	resource.PUT("/admin/:admin_id", r.authorization(), r.runAdminUpdateHandler())
	resource.DELETE("/admin/:admin_id", r.authorization(), r.runAdminDeleteHandler())

	resource.POST("/alamat", r.authorization(), r.runAlamatCreateHandler())
	resource.GET("/alamat", r.authorization(), r.getAllAlamatHandler())
	resource.GET("/alamat/:alamat_id", r.authorization(), r.getOneAlamatHandler())
	resource.PUT("/alamat/:alamat_id", r.authorization(), r.runAlamatUpdateHandler())
	resource.DELETE("/alamat/:alamat_id", r.authorization(), r.runAlamatDeleteHandler())

	resource.POST("/kajian", r.authorization(), r.runKajianCreateHandler())
	resource.GET("/kajian", r.authorization(), r.getAllKajianHandler())
	resource.GET("/kajian/:kajian_id", r.authorization(), r.getOneKajianHandler())
	resource.PUT("/kajian/:kajian_id", r.authorization(), r.runKajianUpdateHandler())
	resource.DELETE("/kajian/:kajian_id", r.authorization(), r.runKajianDeleteHandler())

	resource.POST("/kategori", r.authorization(), r.runKategoriCreateHandler())
	resource.GET("/kategori", r.authorization(), r.getAllKategoriHandler())
	resource.GET("/kategori/:kategori_id", r.authorization(), r.getOneKategoriHandler())
	resource.PUT("/kategori/:kategori_id", r.authorization(), r.runKategoriUpdateHandler())
	resource.DELETE("/kategori/:kategori_id", r.authorization(), r.runKategoriDeleteHandler())

	resource.POST("/pemateri", r.authorization(), r.runPemateriCreateHandler())
	resource.GET("/pemateri", r.authorization(), r.getAllPemateriHandler())
	resource.GET("/pemateri/:pemateri_id", r.authorization(), r.getOnePemateriHandler())
	resource.PUT("/pemateri/:pemateri_id", r.authorization(), r.runPemateriUpdateHandler())
	resource.DELETE("/pemateri/:pemateri_id", r.authorization(), r.runPemateriDeleteHandler())

	resource.POST("/penyelenggara", r.authorization(), r.runPenyelenggaraCreateHandler())
	resource.GET("/penyelenggara", r.authorization(), r.getAllPenyelenggaraHandler())
	resource.GET("/penyelenggara/:penyelenggara_id", r.authorization(), r.getOnePenyelenggaraHandler())
	resource.PUT("/penyelenggara/:penyelenggara_id", r.authorization(), r.runPenyelenggaraUpdateHandler())
	resource.DELETE("/penyelenggara/:penyelenggara_id", r.authorization(), r.runPenyelenggaraDeleteHandler())
}
