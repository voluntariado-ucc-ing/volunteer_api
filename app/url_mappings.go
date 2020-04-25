package app

import "volutarios_api/controllers"

func mapUrls(){
	router.GET("/ping", controllers.PingController.Ping)
	router.GET("/volunteer/get/:id", controllers.VolunteerController.Get)

	router.POST("/volunteer/create", controllers.VolunteerController.Create)
	router.POST("/volunteer/import", controllers.VolunteerController.ImportCsv)

	router.PUT("/volunteer/update/:id", controllers.VolunteerController.Update)

	router.DELETE("/volunteer/delete/:id", controllers.VolunteerController.Delete)
}