package model

import (
	"myapp/datastore/postgres"
)

type Student struct {
	StdID     int64
	FirstName string
	LastName  string
	Email     string
}

const queryInsertUser = "INSERT INTO student (StdID, FirstName,LastName,Email) VALUES($1,$2,$3,$4)"

func (s *Student) Create() error {
	_, err := postgres.Db.Exec(queryInsertUser, s.StdID, s.FirstName, s.LastName, s.Email)
	// fmt.Println(fff)
	return err
}

func (s *Student) Read() error {
	const queryGetUser = "SELECT StdID, FirstName,LastName,Email FROM student WHERE StdID =$1;"
	return postgres.Db.QueryRow(queryGetUser, s.StdID).Scan(&s.StdID, &s.FirstName, &s.LastName, &s.Email)
}

func (s *Student) Update(a int64) error {
	const queryUpdateUser = "UPDATE student SET StdID = $1,FirstName = $2,LastName= $3, Email=$4 WHERE StdID = $5 RETURNING StdID;"
	return postgres.Db.QueryRow(queryUpdateUser, s.StdID, s.FirstName, s.LastName, s.Email, a).Scan(&s.StdID)
}

func (s *Student) Delete(a int64) error {
	const queryDeleteUser = "DELETE  FROM student WHERE StdID = $1 RETURNING StdID;"
	return postgres.Db.QueryRow(queryDeleteUser, a).Scan(&s.StdID)
}

func GetAllStudents() ([]Student, error) {
	const queryGetAll = "SELECT * FROM student;"
	table, err := postgres.Db.Query(queryGetAll)
	if err != nil {
		return nil, err
	}
	students := []Student{}
	for table.Next() {
		var s Student
		dbErr := table.Scan(&s.StdID, &s.FirstName, &s.LastName, &s.Email)
		if dbErr != nil {
			return nil, dbErr
		}
		students = append(students, s)

	}
	table.Close()
	return students, nil
}
