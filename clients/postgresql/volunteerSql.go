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
	queryGet    = "SELECT v.volunteer_id, v.first_name, v.last_name, v.dni, v.email, v.volunteer_type, v.status FROM test_volunteer.volunteer v WHERE v.volunteer_id=$1 "
	queryUpdate = "UPDATE test_volunteer.volunteer v SET first_name=$1, last_name=$2, email=$3, dni=$4, volunteer_type=$5, status=$6 WHERE v.volunteer_id=$7"
	queryDelete = "UPDATE test_volunteer.volunteer v SET status=$1 WHERE v.volunteer_id=$2"
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

func InsertVolunteer(vol *volunteer.Volunteer) (int64, apierrors.ApiError) {
	var id int64
	q, err := dbClient.Prepare(queryInsert)
	if err != nil {
		return 0, apierrors.NewInternalServerApiError("Error preparing insert statement", err)
	}
	res := q.QueryRow(vol.FirstName, vol.LastName, vol.Email, vol.Dni, vol.Type)
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
	err = res.Scan(&vol.Id, &vol.FirstName, &vol.LastName, &vol.Dni, &vol.Email, &vol.Type, &vol.Status)
	if err != nil {
		return nil, apierrors.NewNotFoundApiError("Database query error for get")
	}
	return &vol, nil
}

func UpdateVolunteer(vol *volunteer.Volunteer) apierrors.ApiError {
	if _, err := dbClient.Exec(queryUpdate, vol.FirstName, vol.LastName, vol.Email, vol.Dni, vol.Type, vol.Status, vol.Id); err != nil{
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
