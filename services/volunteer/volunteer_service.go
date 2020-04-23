package volunteer_service

import (
	volunteerSql "volutarios_api/clients/postgresql"
	"volutarios_api/domain/apierrors"
	"volutarios_api/domain/volunteer"
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
	VolunteerService = &volunteerService{}
}

func (v volunteerService) CreateVolunteer(volunteer *volunteer.Volunteer) (*volunteer.Volunteer, apierrors.ApiError) {
	id, err := volunteerSql.InsertVolunteer(volunteer)
	if err != nil {
		return nil, err
	}
	volunteer.Id = id
	return volunteer, nil
}

func (v volunteerService) GetVolunteer(id int64) (*volunteer.Volunteer, apierrors.ApiError) {
	vol, err := volunteerSql.GetVolunteerById(id)
	if err != nil {
		return nil, err
	}
	return vol, nil
}

func (v volunteerService) UpdateVolunteer(updateRequest *volunteer.Volunteer) (*volunteer.Volunteer, apierrors.ApiError) {
	current, err := volunteerSql.GetVolunteerById(updateRequest.Id)
	if err != nil {
		return nil, err
	}
	current.UpdateFields(*updateRequest)
	if err := volunteerSql.UpdateVolunteer(current); err != nil {
		return nil, err
	}
	return current, nil
}

func (v volunteerService) DeleteVolunteer(id int64) apierrors.ApiError {
	if err := volunteerSql.DeleteVolunteer(id); err != nil {
		return err
	}
	return nil
}
