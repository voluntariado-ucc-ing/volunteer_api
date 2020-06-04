package app

import "github.com/voluntariado-ucc-ing/volunteer_api/controllers"

func mapUrls(){
	router.GET("/ping", controllers.PingController.Ping)
	router.GET("/volunteer/get/:id", controllers.VolunteerController.Get)
	router.GET("/volunteer/get", controllers.VolunteerController.GetByUsername)
	router.GET("/volunteer/all", controllers.VolunteerController.GetAllVolunteers)
	router.GET("/volunteer/medical_info/:volunteer_id", controllers.VolunteerController.GetMedicalInfo)

	router.POST("/volunteer/create", controllers.VolunteerController.Create)
	router.POST("/volunteer/import", controllers.VolunteerController.ImportCsv)
	router.POST("/volunteer/medical_info/:volunteer_id", controllers.VolunteerController.SetMedicalInfo)
	router.POST("/volunteer/auth", controllers.VolunteerController.AuthVolunteer)

	router.PUT("/volunteer/update/:id", controllers.VolunteerController.Update)
	router.PUT("/volunteer/auth/update", controllers.VolunteerController.UpdatePassword)

	router.DELETE("/volunteer/delete/:id", controllers.VolunteerController.Delete)
}