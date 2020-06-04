package medical_info

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMedicalInfoStructure(t *testing.T) {
	m := MedicalInfo{
		EmergencyNumber:             "3512146199",
		InsuranceName:               "OSDE",
		InsuranceNumber:             "123123",
		HospitalCenters:             []string{"Reina Fabiola", "Hospital Italiano"},
		BloodType:                   "0-",
		Antitetanic:                 true,
		AntitetanicYear:             "2019",
		InMedicalTreatment:          true,
		MedicalTreatment:            "tratamiento de rodilla",
		UsesMedicines:               true,
		MedicineTypes:               "tafirol 1g",
		MedicineFrecuency:           "1 vez al dia",
		HasAlergies:                 true,
		Alergies:                    []string{"mani", "canela", "chocolate"},
		AlergicToMedicine:           true,
		AlergicMedicineType:         "actron 600",
		HeartDesease:                true,
		HeartDeseaseType:            "soplo en corazon",
		Hypertension:                true,
		Lipothymy:                   true,
		NeuroDesease:                true,
		NeuroDeseaseType:            "lobulo frontal accidentado",
		Diabetes:                    true,
		Asthma:                      true,
		DigestiveDesease:            true,
		DigestiveDeseaseType:        "celiaco",
		EndocrinologicalDesease:     true,
		EndocrinologicalDeseaseType: "tiroides",
		SpineAlterations:            true,
		SpineAlterationsType:        "escoliosis",
		CoagulationDisorder:         true,
		Smokes:                      true,
		FoodIntolerances:            true,
		FoodIntolerancesType:        "lactosa",
		Comentaries:                 "No tengo comentarios",
	}

	b, _ := json.Marshal(m)

	fmt.Println(string(b))
}
