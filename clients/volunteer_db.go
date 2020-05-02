package clients

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/voluntariado-ucc-ing/volunteer_api/config"
	"github.com/voluntariado-ucc-ing/volunteer_api/domain/apierrors"
	"github.com/voluntariado-ucc-ing/volunteer_api/domain/direction"
	"github.com/voluntariado-ucc-ing/volunteer_api/domain/volunteer"
	"log"
)

var dbClient *sql.DB

const (
	queryGetAllVolunteers        = "SELECT v.volunteer_id FROM voluntariado_ing.volunteers v WHERE v.volunteer_id IS NOT NULL"
	queryGetFullVolunteerDetails = "SELECT v.volunteer_id, v.username, v.first_name, v.last_name, v.document_id, v.status, vd.volunteer_details_id, vd.contact_mail,vd.phone_number, vd.photo_url, vd.birth_date, vd.has_car, vd.university, vd.career, vd.career_year, vd.works, vd.career_condition,d.direction_id, d.street, d.number, d.details, d.city, d.postal_code FROM voluntariado_ing.volunteers v INNER JOIN voluntariado_ing.volunteer_details vd ON v.profile_id=vd.volunteer_details_id INNER JOIN  voluntariado_ing.directions d ON vd.direction_id = d.direction_id WHERE v.volunteer_id=$1"
	queryInsertDirection         = "INSERT INTO voluntariado_ing.directions (street, number, details, city, postal_code) VALUES ($1,$2,$3,$4,$5) RETURNING direction_id"
	queryInsertVolunteer         = "INSERT INTO voluntariado_ing.volunteers (first_name, last_name, username, document_id) VALUES ($1,$2,$3,$4) RETURNING volunteer_id"
	queryInsertDetails           = "INSERT INTO voluntariado_ing.volunteer_details (contact_mail, phone_number, photo_url, birth_date, has_car, direction_id, university, career, career_year, works, career_condition) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING volunteer_details_id"
	queryGetById                 = "SELECT v.volunteer_id, v.first_name, v.last_name, v.document_id, v.username, v.status, v.profile_id FROM voluntariado_ing.volunteers v WHERE v.volunteer_id=$1 "
	queryGetByUsername           = "SELECT v.volunteer_id FROM voluntariado_ing.volunteers v WHERE v.username=$1 "
	queryGetDirectionId          = "SELECT v.direction_id FROM voluntariado_ing.volunteer_details v WHERE v.volunteer_details_id=$1"
	queryUpdate                  = "UPDATE voluntariado_ing.volunteers v SET first_name=$1, last_name=$2, username=$3, document_id=$4,status=$5,profile_id=$6 WHERE v.volunteer_id=$7"
	queryUpdateHavingProfile     = "UPDATE voluntariado_ing.volunteers v SET first_name=$1, last_name=$2, username=$3, document_id=$4,status=$5 WHERE v.volunteer_id=$6"
	queryUpdateDetails           = "UPDATE voluntariado_ing.volunteer_details v SET contact_mail=$1, phone_number=$2, photo_url=$3, has_car=$4, university=$5, career=$6, career_year=$7, career_condition=$8, works=$9 WHERE v.volunteer_details_id=$10"
	queryUpdateDirections        = "UPDATE voluntariado_ing.directions d SET street=$1, number=$2, details=$3, city=$4, postal_code=$5 WHERE d.direction_id=$6"
	queryDelete                  = "UPDATE voluntariado_ing.volunteers v SET status=$1 WHERE v.volunteer_id=$2"
)

func init() {
	var err error
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s",
		config.GetDatabaseHost(),
		config.GetDatabaseUser(),
		config.GetDatabasePassword(),
		config.GetDatabaseName())

	dbClient, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := dbClient.Ping(); err != nil {
		log.Fatal(err)
	}
}

func InsertVolunteer(vol *volunteer.Volunteer) (int64, apierrors.ApiError) {
	var id int64
	q, err := dbClient.Prepare(queryInsertVolunteer)
	if err != nil {
		return 0, apierrors.NewInternalServerApiError("Error preparing insert statement", err)
	}
	res := q.QueryRow(vol.FirstName, vol.LastName, vol.Username, vol.DocumentId)
	err = res.Scan(&id)
	if err != nil {
		return 0, apierrors.NewInternalServerApiError("Error scaning last insert id for create", err)
	}
	return id, nil
}

func InsertVolunteerDetails(det volunteer.Details) (int64, apierrors.ApiError) {
	var id int64
	q, err := dbClient.Prepare(queryInsertDetails)
	if err != nil {
		return 0, apierrors.NewInternalServerApiError("Error preparing insert statement", err)
	}
	res := q.QueryRow(det.ContactMail, det.PhoneNumber, det.PhotoUrl, det.BirthDate, det.HasCar, det.Direction.DirectionId, det.University, det.Career, det.CareerYear, det.Works, det.CareerCondition)
	err = res.Scan(&id)
	if err != nil {
		return 0, apierrors.NewInternalServerApiError("Error scaning last insert id for create details", err)
	}
	return id, nil
}

func GetDirectionIdByProfileId(detailsId int64) (int64, apierrors.ApiError) {
	var id int64
	q, err := dbClient.Prepare(queryGetDirectionId)
	if err != nil {
		return 0, apierrors.NewInternalServerApiError("error preparing get direction id statement", err)
	}
	res := q.QueryRow(detailsId)
	err = res.Scan(&id)
	if err != nil {
		return 0, apierrors.NewNotFoundApiError("direction not found")
	}
	return id, nil
}

func InsertDirection(dir direction.Direction) (int64, apierrors.ApiError) {
	var id int64
	q, err := dbClient.Prepare(queryInsertDirection)
	if err != nil {
		return 0, apierrors.NewInternalServerApiError("Error preparing insert statement", err)
	}
	res := q.QueryRow(dir.Street, dir.Number, dir.Details, dir.City, dir.PostalCode)
	err = res.Scan(&id)
	if err != nil {
		return 0, apierrors.NewInternalServerApiError("Error scaning last insert id for create details", err)
	}
	return id, nil
}

func GetVolunteerById(id int64) (*volunteer.Volunteer, apierrors.ApiError) {
	var vol volunteer.Volunteer
	q, err := dbClient.Prepare(queryGetById)
	if err != nil {
		return nil, apierrors.NewInternalServerApiError("Error preparing get statement", err)
	}
	res := q.QueryRow(id)
	err = res.Scan(&vol.Id, &vol.FirstName, &vol.LastName, &vol.DocumentId, &vol.Username, &vol.StatusId, &vol.VolunteerProfileId)
	if err != nil {
		fmt.Println(err)
		return nil, apierrors.NewNotFoundApiError("volunteer not found")
	}
	return &vol, nil
}

func GetVolunteerFullDetailsById(id int64) (*volunteer.Volunteer, apierrors.ApiError) {
	var vol volunteer.Volunteer
	q, err := dbClient.Prepare(queryGetFullVolunteerDetails)
	if err != nil {
		return nil, apierrors.NewInternalServerApiError("Error preparing get full details statement", err)
	}
	res := q.QueryRow(id)
	err = res.Scan(&vol.Id, &vol.Username, &vol.FirstName, &vol.LastName, &vol.DocumentId, &vol.StatusId,
		&vol.VolunteerDetails.DetailsId, &vol.VolunteerDetails.ContactMail, &vol.VolunteerDetails.PhoneNumber,
		&vol.VolunteerDetails.PhotoUrl, &vol.VolunteerDetails.BirthDate, &vol.VolunteerDetails.HasCar,
		&vol.VolunteerDetails.University, &vol.VolunteerDetails.Career, &vol.VolunteerDetails.CareerYear,
		&vol.VolunteerDetails.Works, &vol.VolunteerDetails.CareerCondition, &vol.VolunteerDetails.Direction.DirectionId,
		&vol.VolunteerDetails.Direction.Street, &vol.VolunteerDetails.Direction.Number, &vol.VolunteerDetails.Direction.Details,
		&vol.VolunteerDetails.Direction.City, &vol.VolunteerDetails.Direction.PostalCode)
	if err != nil {
		fmt.Println(err)
		return nil, apierrors.NewNotFoundApiError("volunteer full details not found (no profile_id)")
	}
	vol.Status = vol.StatusId.Int32
	return &vol, nil
}

func GetAllVolunteerIds() ([]int64, apierrors.ApiError) {
	ids := make([]int64, 0)
	q, err := dbClient.Prepare(queryGetAllVolunteers)
	if err != nil {
		fmt.Println(err)
		return nil, apierrors.NewInternalServerApiError("Error preparing get all volunteers statement", err)
	}

	res, err := q.Query()
	if err != nil {
		fmt.Println(err)
		return nil, apierrors.NewNotFoundApiError("no volunteers found")
	}

	defer res.Close()

	for res.Next() {
		var id int64
		err := res.Scan(&id)
		if err != nil {
			return nil, apierrors.NewNotFoundApiError("id not found in get all")
		}
		ids = append(ids, id)
	}

	return ids, nil
}

func UpdateVolunteerTable(vol *volunteer.Volunteer) apierrors.ApiError {
	if _, err := dbClient.Exec(queryUpdate, vol.FirstName, vol.LastName, vol.Username, vol.DocumentId, vol.Status, vol.VolunteerDetails.DetailsId, vol.Id); err != nil {
		return apierrors.NewInternalServerApiError("Error database query response for update", err)
	}
	return nil
}

func UpdateVolunteerTableHavingProfileId(vol *volunteer.Volunteer) apierrors.ApiError {
	if _, err := dbClient.Exec(queryUpdateHavingProfile, vol.FirstName, vol.LastName, vol.Username, vol.DocumentId, vol.Status, vol.Id); err != nil {
		return apierrors.NewInternalServerApiError("Error database query response for update", err)
	}
	return nil
}

func UpdateVolunteerDetailsTable(det *volunteer.Details) apierrors.ApiError {
	if _, err := dbClient.Exec(queryUpdateDetails, det.ContactMail, det.PhoneNumber, det.PhotoUrl, det.HasCar, det.University, det.Career, det.CareerYear, det.CareerCondition, det.Works, det.DetailsId); err != nil {
		return apierrors.NewInternalServerApiError("Error database query response for update details", err)
	}
	return nil
}

func UpdateDirectionTable(dir *direction.Direction) apierrors.ApiError {
	if _, err := dbClient.Exec(queryUpdateDirections, dir.Street, dir.Number, dir.Details, dir.City, dir.PostalCode, dir.DirectionId); err != nil {
		return apierrors.NewInternalServerApiError("Error database query response for update details", err)
	}
	return nil
}

func GetIdByMail(username string) (int64, apierrors.ApiError) {
	var volId int64
	q, err := dbClient.Prepare(queryGetByUsername)
	if err != nil {
		return 0, apierrors.NewInternalServerApiError("Error preparing get full details statement", err)
	}
	res := q.QueryRow(username)
	err = res.Scan(&volId)
	if err != nil {
		fmt.Println(err)
		return 0, apierrors.NewNotFoundApiError("username not found")
	}

	return volId, nil
}

func DeleteVolunteer(id int64) apierrors.ApiError {
	if _, err := dbClient.Exec(queryDelete, volunteer.StatusDeleted, id); err != nil {
		return apierrors.NewInternalServerApiError("Error database query response for logical delete", err)
	}
	return nil
}
