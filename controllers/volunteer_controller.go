package controllers

import (
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"github.com/voluntariado-ucc-ing/volunteer_api/domain/apierrors"
	"github.com/voluntariado-ucc-ing/volunteer_api/domain/volunteer"
	volunteerservice "github.com/voluntariado-ucc-ing/volunteer_api/services/volunteer"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var (
	VolunteerController volunteerControllerInterface = &volunteerController{}
)

type volunteerControllerInterface interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	ImportCsv(c *gin.Context)
	GetByUsername(c *gin.Context)
	GetAllVolunteers(c *gin.Context)
}

type volunteerController struct{}

func (v *volunteerController) Create(c *gin.Context) {
	var volunteerRequest volunteer.Volunteer
	if err := c.ShouldBindJSON(&volunteerRequest); err != nil {
		apiErr := apierrors.NewBadRequestApiError("Error invalid JSON")
		c.JSON(http.StatusBadRequest, apiErr)
		return
	}

	res, err := volunteerservice.VolunteerService.CreateVolunteer(&volunteerRequest)
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
	res, serErr := volunteerservice.VolunteerService.GetVolunteer(id)
	if serErr != nil {
		c.JSON(serErr.Status(), serErr)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (v *volunteerController) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		badR := apierrors.NewBadRequestApiError("Error parsing parameter")
		c.JSON(badR.Status(), badR)
		return
	}
	var updateRequest volunteer.Volunteer
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		apiErr := apierrors.NewBadRequestApiError("Error invalid JSON")
		c.JSON(http.StatusBadRequest, apiErr)
		return
	}
	updateRequest.Id = id
	res, serErr := volunteerservice.VolunteerService.UpdateVolunteer(&updateRequest)
	if serErr != nil {
		c.JSON(serErr.Status(), serErr)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (v *volunteerController) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		badR := apierrors.NewBadRequestApiError("Error parsing parameter")
		c.JSON(badR.Status(), badR)
		return
	}
	delErr := volunteerservice.VolunteerService.DeleteVolunteer(id)
	if delErr != nil {
		c.JSON(delErr.Status(), delErr)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func (v *volunteerController) ImportCsv(c *gin.Context) {
	f, err := c.FormFile("file")
	if err != nil {
		badR := apierrors.NewBadRequestApiError("Error parsing file parameter")
		c.JSON(badR.Status(), badR)
		return
	}

	file, err := f.Open()
	if err != nil {
		internal := apierrors.NewInternalServerApiError("Error opening input file", err)
		c.JSON(internal.Status(), internal)
		return
	}

	r := csv.NewReader(file)

	var newVolunteers []volunteer.Volunteer

	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		//[0] Mail
		//[1] DNI
		dni, err := strconv.ParseInt(strings.TrimSpace(record[1]), 10, 64)
		if err != nil {
			badR := apierrors.NewBadRequestApiError("Error parsing file data")
			c.JSON(badR.Status(), badR)
			return
		}
		v := volunteer.Volunteer{
			Username:   strings.TrimSpace(record[0]),
			DocumentId: dni,
		}
		newVolunteers = append(newVolunteers, v)
	}

	for _, newVolunteer := range newVolunteers {
		_, err := volunteerservice.VolunteerService.CreateVolunteer(&newVolunteer)
		if err != nil {
			internal := apierrors.NewInternalServerApiError("Error creating user from file", err)
			c.JSON(internal.Status(), internal)
			return
		}
	}
	c.JSON(http.StatusOK, map[string]string{"status": "created"})
}

func (v *volunteerController) GetByUsername(c *gin.Context) {
	username := c.Query("username")
	res, err := volunteerservice.VolunteerService.GetVolunteerByUsername(username)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, res)
	return
}

func (v *volunteerController) GetAllVolunteers(c *gin.Context) {
	res, err := volunteerservice.VolunteerService.GetAllVolunteers()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, res)
	return
}

