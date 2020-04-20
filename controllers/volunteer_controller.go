package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"volutarios_api/domain/apierrors"
	"volutarios_api/domain/volunteer"
	volunteer_service "volutarios_api/services/volunteer"
)

var (
	VolunteerController volunteerControllerInterface = &volunteerController{}
)

type volunteerControllerInterface interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type volunteerController struct{}

func (v *volunteerController) Create(c *gin.Context) {
	var volunteerRequest volunteer.Volunteer
	if err := c.ShouldBindJSON(&volunteerRequest); err != nil {
		apiErr := apierrors.NewBadRequestApiError("Error invalid JSON")
		c.JSON(http.StatusBadRequest, apiErr)
		return
	}

	res, err := volunteer_service.VolunteerService.CreateVolunteer(volunteerRequest)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (v *volunteerController) Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		badR := apierrors.NewBadRequestApiError("Error parsing parameter")
		c.JSON(badR.Status(), badR)
		return
	}
	res, serErr := volunteer_service.VolunteerService.GetVolunteer(id)
	if serErr != nil {
		c.JSON(serErr.Status(), serErr)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (v *volunteerController) Update(c *gin.Context) {
	// TODO: Implement!
	c.JSON(http.StatusOK, map[string]string{"status": "updated"})
}

func (v *volunteerController) Delete(c *gin.Context) {
	// TODO: Implement!
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
