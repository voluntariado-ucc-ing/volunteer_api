package volunteerSql

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"volutarios_api/domain/apierrors"
	"volutarios_api/domain/volunteer"
)

var dbClient *sql.DB

const (
	queryInsert = "INSERT INTO test_volunteer.volunteer (first_name, last_name, email, dni, volunteer_type) VALUES ($1,$2,$3,$4,$5) RETURNING volunteer_id"
	queryGet    = "SELECT v.volunteer_id, v.first_name, v.last_name, v.dni, v.email, v.volunteer_type FROM test_volunteer.volunteer v WHERE v.volunteer_id=$1 "
)

func init() {
	var err error
	connStr := "dbname=test_volunteer sslmode=disable"
	dbClient, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err := dbClient.Ping(); err != nil {
		log.Fatal(err)
	}
}

func InsertVolunteer(vol volunteer.Volunteer) (int64, apierrors.ApiError) {
	var id int64
	q, err := dbClient.Prepare(queryInsert)
	if err != nil {
		return 0, apierrors.NewInternalServerApiError("Error preparing statement", err)
	}
	res := q.QueryRow(vol.FirstName, vol.LastName, vol.Email, vol.Dni, vol.Type)
	err = res.Scan(&id)
	if err != nil {
		return 0, apierrors.NewInternalServerApiError("Error scaning last insert id", err)
	}
	return id, nil
}

func GetVolunteerById(id int64) (*volunteer.Volunteer, apierrors.ApiError) {
	var vol volunteer.Volunteer
	q, err := dbClient.Prepare(queryGet)
	if err != nil {
		return nil, apierrors.NewInternalServerApiError("Error preparing statement", err)
	}
	res := q.QueryRow(id)
	err = res.Scan(&vol.Id, &vol.FirstName, &vol.LastName, &vol.Dni, &vol.Email, &vol.Type)
	if err != nil {
		return nil, apierrors.NewNotFoundApiError("Database query error")
	}
	return &vol, nil
}

func UpdateVolunteer(vol volunteer.Volunteer) apierrors.ApiError {
	panic("implement me")
}

func DeleteVolunteer(vol volunteer.Volunteer) apierrors.ApiError {
	panic("implement me")
}
