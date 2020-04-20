package volunteer_service

import (
	"volutarios_api/domain/apierrors"
	"volutarios_api/domain/volunteer"
)

type volunteerService struct{}

type volunteerServiceInterface interface {
	CreateVolunteer(volunteer volunteer.Volunteer) (*volunteer.Volunteer, apierrors.ApiError)
	GetVolunteer(id int64) (*volunteer.Volunteer, apierrors.ApiError)
	UpdateVolunteer(volunteer volunteer.Volunteer) (*volunteer.Volunteer, apierrors.ApiError)
	DeleteVolunteer(volunteer volunteer.Volunteer) (*volunteer.Volunteer, apierrors.ApiError)
}

var (
	VolunteerService volunteerServiceInterface
)

func init() {
	VolunteerService = &volunteerService{}
}

func (v volunteerService) CreateVolunteer(volunteer volunteer.Volunteer) (*volunteer.Volunteer, apierrors.ApiError) {
	// TODO: Implement!
	return nil, nil
}

func (v volunteerService) GetVolunteer(id int64) (*volunteer.Volunteer, apierrors.ApiError) {
	// TODO: Implement!
	return nil, nil
}

func (v volunteerService) UpdateVolunteer(volunteer volunteer.Volunteer) (*volunteer.Volunteer, apierrors.ApiError) {
	// TODO: Implement!
	return nil, nil
}

func (v volunteerService) DeleteVolunteer(volunteer volunteer.Volunteer) (*volunteer.Volunteer, apierrors.ApiError) {
	// TODO: Implement!
	return nil, nil
}


