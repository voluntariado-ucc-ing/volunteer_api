package volunteer

import "database/sql"

const (
	StatusDeleted = "deleted"
)

type Volunteer struct {
	Id                 int64         `json:"volunteer_id,omitempty"`
	FirstName          string        `json:"first_name"`
	LastName           string        `json:"last_name"`
	Username           string        `json:"username"`
	DocumentId         int64         `json:"document_id"`
	Status             int32         `json:"status"`
	StatusId           sql.NullInt32 `json:"-"`
	VolunteerDetails   Details       `json:"details"`
	VolunteerProfileId sql.NullInt64 `json:"-"`
}

func (v *Volunteer) UpdateFields(newVol Volunteer) {
	if newVol.Status != 0 {
		v.Status = newVol.Status
	}
	if newVol.Username != "" {
		v.Username = newVol.Username
	}
	if newVol.FirstName != "" {
		v.FirstName = newVol.FirstName
	}
	if newVol.LastName != "" {
		v.LastName = newVol.LastName
	}
	if newVol.DocumentId != 0 {
		v.DocumentId = newVol.DocumentId
	}
	v.VolunteerDetails.UpdateDetails(newVol.VolunteerDetails)
}
