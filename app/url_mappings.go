package app

import "github.com/voluntariado-ucc-ing/volunteer_api/controllers"

func mapUrls(){
	router.GET("/ping", controllers.PingController.Ping)
	router.GET("/volunteer/get/:id", controllers.VolunteerController.Get)
	router.GET("/volunteer/get", controllers.VolunteerController.GetByUsername)
	router.GET("/volunteer/all", controllers.VolunteerController.GetAllVolunteers)



	router.POST("/volunteer/create", controllers.VolunteerController.Create)
	router.POST("/volunteer/import", controllers.VolunteerController.ImportCsv)

	router.PUT("/volunteer/update/:id", controllers.VolunteerController.Update)

	router.DELETE("/volunteer/delete/:id", controllers.VolunteerController.Delete)
}