package direction

type Direction struct {
	DirectionId int64  `json:"direction_id"`
	Street      string `json:"street"`
	Number      int64  `json:"number"`
	Details     string `json:"details"`
	City        string `json:"city"`
	PostalCode  int64  `json:"postal_code"`
}

func (d *Direction) UpdateDirection(newD Direction) {
	if newD.Street != "" {
		d.Street = newD.Street
	}
	if newD.Number != 0 {
		d.Number = newD.Number
	}
	if newD.Details != "" {
		d.Details = newD.Details
	}
	if newD.City != "" {
		d.City = newD.City
	}
	if newD.PostalCode != 0 {
		d.PostalCode = newD.PostalCode
	}
}
