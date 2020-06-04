package volunteer

type MedicalInfo struct {
	MedicalInfoId int64  `json:"medical_info_id"`
	VolunteerId   int64  `json:"volunteer_id"`
	MedicalData   []byte `json:"medical_data"`
}
