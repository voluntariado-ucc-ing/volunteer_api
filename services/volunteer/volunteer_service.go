package volunteer_service

import (
	"github.com/voluntariado-ucc-ing/volunteer_api/clients"
	"github.com/voluntariado-ucc-ing/volunteer_api/domain/apierrors"
	"github.com/voluntariado-ucc-ing/volunteer_api/domain/volunteer"
	"math/rand"
	"time"
)

type volunteerService struct{}

type volunteerServiceInterface interface {
	CreateVolunteer(volunteer *volunteer.Volunteer) (*volunteer.Volunteer, apierrors.ApiError)
	GetVolunteer(id int64) (*volunteer.Volunteer, apierrors.ApiError)
	UpdateVolunteer(volunteer *volunteer.Volunteer) (*volunteer.Volunteer, apierrors.ApiError)
	DeleteVolunteer(id int64) apierrors.ApiError
	GetAllVolunteers() ([]volunteer.Volunteer, apierrors.ApiError)
	GetVolunteerByUsername(username string) (*volunteer.Volunteer, apierrors.ApiError)
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
	/* TODO async
	password := rand.Uint64()
	err = providers.SendMail(volunteer.Username, fmt.Sprintf("%d", password))
	if err != nil {
		return nil, err
	}
	*/
	return volunteer, nil
}

func (v volunteerService) GetVolunteer(id int64) (*volunteer.Volunteer, apierrors.ApiError) {
	res, err := clients.GetVolunteerById(id)
	if err != nil {
		return nil, err
	}
	if res.VolunteerProfileId.Int64 == 0 { // Case doesnt have full details yet.
		return res, nil
	}
	return clients.GetVolunteerFullDetailsById(id)
}

func (v volunteerService) UpdateVolunteer(updateRequest *volunteer.Volunteer) (*volunteer.Volunteer, apierrors.ApiError) {
	current, err := clients.GetVolunteerById(updateRequest.Id)
	if err != nil {
		return nil, err
	}
	current.UpdateFields(*updateRequest)
	if current.VolunteerProfileId.Int64 == 0 {
		// User doesnt have direction in details, so create it
		dirId, err := clients.InsertDirection(current.VolunteerDetails.Direction)
		if err != nil {
			return nil, err
		}
		// Assign direction id to details and insert
		current.VolunteerDetails.Direction.DirectionId = dirId
		detId, err := clients.InsertVolunteerDetails(current.VolunteerDetails)
		if err != nil {
			return nil, err
		}
		// Assign details id to user and insert after
		current.VolunteerDetails.DetailsId = detId
		if err := clients.UpdateVolunteerTable(current); err != nil {
			return nil, err
		}
	} else {
		current.VolunteerDetails.DetailsId = current.VolunteerProfileId.Int64

		if err := clients.UpdateVolunteerTableHavingProfileId(current); err != nil {
			return nil, err
		}

		if err := clients.UpdateVolunteerDetailsTable(&current.VolunteerDetails); err != nil {
			return nil, err
		}

		directionId, err := clients.GetDirectionIdByProfileId(current.VolunteerDetails.DetailsId)
		if err != nil {
			return nil, err
		}

		current.VolunteerDetails.Direction.DirectionId = directionId
		if err := clients.UpdateDirectionTable(&current.VolunteerDetails.Direction); err != nil {
			return nil, err
		}
	}
	return current, nil
}

func (v volunteerService) GetAllVolunteers() ([]volunteer.Volunteer, apierrors.ApiError) {
	res := make([]volunteer.Volunteer, 0)
	ids, err := clients.GetAllVolunteerIds()
	if err != nil {
		return nil, err
	}

	for _, id := range ids {
		v, err := v.GetVolunteer(id)
		if err != nil {
			return nil, err
		}
		res = append(res, *v)
	}

	return res, nil
}

func (v volunteerService) GetVolunteerByUsername(username string) (*volunteer.Volunteer, apierrors.ApiError) {
	id, err := clients.GetIdByMail(username)
	if err != nil {
		return nil, err
	}
	return v.GetVolunteer(id)
}

func (v volunteerService) DeleteVolunteer(id int64) apierrors.ApiError {
	if err := clients.DeleteVolunteer(id); err != nil {
		return err
	}
	return nil
}
