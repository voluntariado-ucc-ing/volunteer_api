package volunteer_service

import (
	"fmt"
	"github.com/voluntariado-ucc-ing/volunteer_api/clients"
	"github.com/voluntariado-ucc-ing/volunteer_api/domain/apierrors"
	"github.com/voluntariado-ucc-ing/volunteer_api/domain/volunteer"
	"github.com/voluntariado-ucc-ing/volunteer_api/providers"
	"math/rand"
	"time"
)

type volunteerService struct{}

type volunteerServiceInterface interface {
	CreateVolunteer(volunteer *volunteer.Volunteer) (*volunteer.Volunteer, apierrors.ApiError)
	GetVolunteer(id int64) (*volunteer.Volunteer, apierrors.ApiError)
	UpdateVolunteer(volunteer *volunteer.Volunteer) (*volunteer.Volunteer, apierrors.ApiError)
	DeleteVolunteer(id int64) apierrors.ApiError
}

var (
	VolunteerService volunteerServiceInterface
)

func init() {
	rand.Seed(time.Now().UnixNano())
	VolunteerService = &volunteerService{}
}

func (v volunteerService) CreateVolunteer(volunteer *volunteer.Volunteer) (*volunteer.Volunteer, apierrors.ApiError) {
	id, err := clients.InsertVolunteer(volunteer)
	if err != nil {
		return nil, err
	}
	volunteer.Id = id
	password := rand.Uint64()
	err = providers.SendMail(volunteer.Email, fmt.Sprintf("%d", password))
	if err != nil {
		return nil, err
	}
	return volunteer, nil
}

func (v volunteerService) GetVolunteer(id int64) (*volunteer.Volunteer, apierrors.ApiError) {
	vol, err := clients.GetVolunteerById(id)
	if err != nil {
		return nil, err
	}
	return vol, nil
}

func (v volunteerService) UpdateVolunteer(updateRequest *volunteer.Volunteer) (*volunteer.Volunteer, apierrors.ApiError) {
	current, err := clients.GetVolunteerById(updateRequest.Id)
	if err != nil {
		return nil, err
	}
	current.UpdateFields(*updateRequest)
	if err := clients.UpdateVolunteer(current); err != nil {
		return nil, err
	}
	return current, nil
}

func (v volunteerService) DeleteVolunteer(id int64) apierrors.ApiError {
	if err := clients.DeleteVolunteer(id); err != nil {
		return err
	}
	return nil
}
