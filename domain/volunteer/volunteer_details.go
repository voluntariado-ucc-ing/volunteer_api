package volunteer

import "github.com/voluntariado-ucc-ing/volunteer_api/domain/direction"

type Details struct {
	DetailsId       int64               `json:"volunteer_details_id"`
	ContactMail     string              `json:"contact_email"`
	PhoneNumber     string              `json:"phone_number"`
	PhotoUrl        string              `json:"photo_url"`
	BirthDate       string              `json:"birth_date"`
	HasCar          bool                `json:"has_car"`
	University      string              `json:"university"`
	Career          string              `json:"career"`
	CareerYear      int                 `json:"career_year"`
	CareerCondition string              `json:"career_condition"`
	Works           int                 `json:"works"`
	Direction       direction.Direction `json:"direction"`
}

func (d *Details) UpdateDetails(newDet Details) {
	if newDet.ContactMail != "" {
		d.ContactMail = newDet.ContactMail
	}
	if newDet.PhoneNumber != "" {
		d.PhoneNumber = newDet.PhoneNumber
	}
	if newDet.PhotoUrl != "" {
		d.PhotoUrl = newDet.PhotoUrl
	}
	if newDet.BirthDate != "" {
		d.BirthDate = newDet.BirthDate
	}
	d.HasCar = newDet.HasCar
	if newDet.University != "" {
		d.University = newDet.University
	}
	if newDet.Career != "" {
		d.Career = newDet.Career
	}
	if newDet.CareerYear != 0 {
		d.CareerYear = newDet.CareerYear
	}
	if newDet.CareerCondition != "" {
		d.CareerCondition = newDet.CareerCondition
	}
	d.Works = newDet.Works
	d.Direction.UpdateDirection(newDet.Direction)
}
