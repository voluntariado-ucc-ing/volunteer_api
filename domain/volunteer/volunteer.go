package volunteer

import "volutarios_api/domain/apierrors"

type Volunteer struct {
	Id        int64  `json:"volunteer_id,omitempty"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Dni       int64  `json:"dni"`
	Type      string `json:"volunteer_type"`
}

func (v *Volunteer) ValidateMail() apierrors.ApiError {
	// TODO: Implement!
	return nil
}
