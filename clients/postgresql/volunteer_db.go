package volunteerSql

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/voluntariado-ucc-ing/volunteer_api/domain/apierrors"
	"github.com/voluntariado-ucc-ing/volunteer_api/domain/volunteer"
	"log"
)

var dbClient *sql.DB

const (
	queryInsert = "INSERT INTO voluntariado_ing.volunteers (first_name, last_name, email, dni) VALUES ($1,$2,$3,$4) RETURNING volunteer_id"
	queryGet    = "SELECT v.volunteer_id, v.first_name, v.last_name, v.dni, v.email, v.status FROM test_volunteer.volunteers v WHERE v.volunteer_id=$1 "
	queryUpdate = "UPDATE voluntariado_ing.volunteers v SET first_name=$1, last_name=$2, email=$3, dni=$4,status=$5 WHERE v.volunteer_id=$6"
	queryDelete = "UPDATE voluntariado_ing.volunteers v SET status=$1 WHERE v.volunteer_id=$2"
)

func init() {
	var err error
	connStr := "dbname=voluntariado_ing sslmode=disable"
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
	q, err := dbClient.Prepare(queryInsert)
	if err != nil {
		return 0, apierrors.NewInternalServerApiError("Error preparing insert statement", err)
	}
	res := q.QueryRow(vol.FirstName, vol.LastName, vol.Email, vol.Dni)
	err = res.Scan(&id)
	if err != nil {
		return 0, apierrors.NewInternalServerApiError("Error scaning last insert id for create", err)
	}
	return id, nil
}

func GetVolunteerById(id int64) (*volunteer.Volunteer, apierrors.ApiError) {
	var vol volunteer.Volunteer
	q, err := dbClient.Prepare(queryGet)
	if err != nil {
		return nil, apierrors.NewInternalServerApiError("Error preparing get statement", err)
	}
	res := q.QueryRow(id)
	err = res.Scan(&vol.Id, &vol.FirstName, &vol.LastName, &vol.Dni, &vol.Email, &vol.Status)
	if err != nil {
		return nil, apierrors.NewNotFoundApiError("Database query error for get")
	}
	return &vol, nil
}

func UpdateVolunteer(vol *volunteer.Volunteer) apierrors.ApiError {
	if _, err := dbClient.Exec(queryUpdate, vol.FirstName, vol.LastName, vol.Email, vol.Dni, vol.Status, vol.Id); err != nil{
		return apierrors.NewInternalServerApiError("Error database query response for update", err)
	}
	return nil
}

func DeleteVolunteer(id int64) apierrors.ApiError {
	if _, err := dbClient.Exec(queryDelete, volunteer.StatusDeleted, id); err != nil {
		return apierrors.NewInternalServerApiError("Error database query response for logical delete", err)
	}
	return nil
}
