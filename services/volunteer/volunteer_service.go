package volunteer_service

import (
	"fmt"
	"github.com/voluntariado-ucc-ing/volunteer_api/clients"
	"github.com/voluntariado-ucc-ing/volunteer_api/domain/apierrors"
	"github.com/voluntariado-ucc-ing/volunteer_api/domain/auth"
	"github.com/voluntariado-ucc-ing/volunteer_api/domain/medical_info"
	"github.com/voluntariado-ucc-ing/volunteer_api/domain/volunteer"
	"github.com/voluntariado-ucc-ing/volunteer_api/providers"
	"github.com/voluntariado-ucc-ing/volunteer_api/utils"
	"math/rand"
	"time"
)

type volunteerService struct{}

type volunteerServiceInterface interface {
	ValidateAuth(authRequest auth.Credentials) (*volunteer.Volunteer, apierrors.ApiError)
	CreateVolunteer(volunteer *volunteer.Volunteer) (*volunteer.Volunteer, apierrors.ApiError)
	GetVolunteer(id int64) (*volunteer.Volunteer, apierrors.ApiError)
	UpdateVolunteer(volunteer *volunteer.Volunteer) (*volunteer.Volunteer, apierrors.ApiError)
	DeleteVolunteer(id int64) apierrors.ApiError
	GetAllVolunteers() ([]volunteer.Volunteer, apierrors.ApiError)
	GetVolunteerByUsername(username string) (*volunteer.Volunteer, apierrors.ApiError)
	UpdatePassword(credentials auth.Credentials) apierrors.ApiError
	SetMedicalInfo(volunteerId int64, info medical_info.MedicalInfo) apierrors.ApiError
	GetMedicalInfo(volunteerId int64) ([]byte, apierrors.ApiError)
}

var (
	VolunteerService volunteerServiceInterface
)

func init() {
	rand.Seed(time.Now().UnixNano())
	VolunteerService = &volunteerService{}
}

func (v volunteerService) CreateVolunteer(vol *volunteer.Volunteer) (*volunteer.Volunteer, apierrors.ApiError) {
	generatedPassword := fmt.Sprintf("%d", rand.Uint32())
	fmt.Println(generatedPassword)
	hashedPassword, hashErr := utils.HashPassword(generatedPassword)
	if hashErr != nil {
		return nil, apierrors.NewInternalServerApiError("Error creating password", hashErr)
	}

	vol.Password = hashedPassword
	id, err := clients.InsertVolunteer(vol)
	if err != nil {
		return nil, err
	}

	vol.Id = id

	go providers.SendMail(vol.Username, generatedPassword)

	return vol, nil
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

	if err := clients.GetVolunteerAlreadyLoggedIn(current.Id); err != nil {
		_ = clients.SetVolunteerAlreadyLoggedIn(current.Id)
	}

	updated, err := v.GetVolunteer(current.Id)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (v volunteerService) GetAllVolunteers() ([]volunteer.Volunteer, apierrors.ApiError) {
	res := make([]volunteer.Volunteer, 0)
	ids, err := clients.GetAllVolunteerIds()
	if err != nil {
		return nil, err
	}

	input := make(chan volunteer.VolunteerConcurrent)
	defer close(input)
	for _, id := range ids {
		go v.getConcurrentVolunteer(id, input)
	}

	for i := 0; i < len(ids); i++ {
		result := <-input
		if result.Error != nil {
			return nil, result.Error
		}
		res = append(res, *result.Volunteer)
	}

	return res, nil
}

func (v volunteerService) getConcurrentVolunteer(id int64, output chan volunteer.VolunteerConcurrent) {
	vol, err := v.GetVolunteer(id)
	output <- volunteer.VolunteerConcurrent{Volunteer: vol, Error: err}
	return
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

func (v volunteerService) ValidateAuth(authRequest auth.Credentials) (*volunteer.Volunteer, apierrors.ApiError) {
	hashedPass, err := clients.GetHashedPasswordByUsername(authRequest.Username)
	if err != nil {
		return nil, err
	}
	if utils.CheckPasswordHash(authRequest.Password, hashedPass) {
		res, err := v.GetVolunteerByUsername(authRequest.Username)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		if err := clients.GetVolunteerAlreadyLoggedIn(res.Id); err == nil {
			t := true
			res.HasLoggedIn = &t
		} else {
			f := false
			res.HasLoggedIn = &f
		}

		return res, nil
	}
	return nil, apierrors.NewForbiddenApiError("Forbidden. Incorrect username or password")
}

func (v volunteerService) UpdatePassword(credentials auth.Credentials) apierrors.ApiError {
	if _, err := v.ValidateAuth(credentials); err != nil { // Validate if old password is right
		return err
	}

	hashedPass, err := utils.HashPassword(credentials.NewPassword)
	if err != nil {
		return apierrors.NewInternalServerApiError("Error hashing password", err)
	}
	return clients.UpdatePassword(hashedPass, credentials.Username)
}

func (v volunteerService) SetMedicalInfo(volunteerId int64, info medical_info.MedicalInfo) apierrors.ApiError {
	return clients.InsertMedicalInfo(volunteerId, info)
}

func (v volunteerService) GetMedicalInfo(volunteerId int64) ([]byte, apierrors.ApiError) {
	data, err := clients.GetVolunteerMedicalInfoById(volunteerId)
	if err != nil {
		return nil, err
	}
	// Returning []byte instead of medical info struct to avoid signature changes related errors
	return []byte(*data), nil
}
