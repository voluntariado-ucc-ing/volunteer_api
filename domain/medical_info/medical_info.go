package medical_info

type MedicalInfo struct {
	EmergencyNumber             string   `json:"emergency_number"`
	InsuranceName               string   `json:"insurance_name"`
	InsuranceNumber             string   `json:"insurance_number"`
	HospitalCenters             []string `json:"hospital_centers"`
	BloodType                   string   `json:"blood_type"`
	Antitetanic                 bool     `json:"antitetanic"`
	AntitetanicYear             string   `json:"antitetanic_year"`
	InMedicalTreatment          bool     `json:"in_medical_treatment"`
	MedicalTreatment            string   `json:"medical_treatment"`
	UsesMedicines               bool     `json:"uses_medicines"`
	MedicineTypes               string   `json:"medicine_type"`
	MedicineFrecuency           string   `json:"medicine_frecuency"`
	HasAlergies                 bool     `json:"has_alergies"`
	Alergies                    []string `json:"alergies"`
	AlergicToMedicine           bool     `json:"has_alergie_to_medicine"`
	AlergicMedicineType         string   `json:"alergic_medicine_type"`
	HeartDesease                bool     `json:"has_heart_desease"`
	HeartDeseaseType            string   `json:"heart_desease"`
	Hypertension                bool     `json:"has_hypertension"`
	Lipothymy                   bool     `json:"has_lipothymy"`
	NeuroDesease                bool     `json:"has_neuro_desease"`
	NeuroDeseaseType            string   `json:"neuro_desease"`
	Diabetes                    bool     `json:"has_diabetes"`
	Asthma                      bool     `json:"has_asthma"`
	DigestiveDesease            bool     `json:"has_digestive_desease"`
	DigestiveDeseaseType        string   `json:"digestive_desease"`
	EndocrinologicalDesease     bool     `json:"has_endocrinological_desease"`
	EndocrinologicalDeseaseType string   `json:"endocrinological_desease"`
	SpineAlterations            bool     `json:"has_spine_alterations"`
	SpineAlterationsType        string   `json:"spine_alterations"`
	CoagulationDisorder         bool     `json:"has_coagulation_desorder"`
	Smokes                      bool     `json:"smokes"`
	FoodIntolerances            bool     `json:"has_food_intolerances"`
	FoodIntolerancesType        string   `json:"food_intolerances"`
	Comentaries                 string   `json:"others"`
}

type VolunteerMedicalInfo struct {
	VolunteerId int64 `json:"volunteer_id"`
	MedicalInfo
}
