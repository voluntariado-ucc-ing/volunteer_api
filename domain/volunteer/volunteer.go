package volunteer

import "github.com/voluntariado-ucc-ing/volunteer_api/domain/apierrors"

const (
	StatusDeleted = "deleted"
)

type Volunteer struct {
	Id        int64  `json:"volunteer_id,omitempty"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Dni       int64  `json:"dni"`
	Status    string `json:"status,omitempty"`
}

func (v *Volunteer) ValidateMail() apierrors.ApiError {
	// TODO: Implement!
	return nil
}

func (v *Volunteer) UpdateFields(newVol Volunteer) {
	if newVol.Status != "" {
		v.Status = newVol.Status
	}
	if newVol.Email != "" {
		v.Email = newVol.Email
	}
	if newVol.FirstName != "" {
		v.FirstName = newVol.FirstName
	}
	if newVol.LastName != "" {
		v.LastName = newVol.LastName
	}
	if newVol.Dni != 0 {
		v.Dni = newVol.Dni
	}
}
