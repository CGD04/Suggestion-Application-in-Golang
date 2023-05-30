package model

import "myapp/datastore/postgres"

type Enroll struct {
	StdId    int64  `json:"stdid"`
	CourseId string `json:"courseid"`
	Date     string `json:"date"`
}

func (e *Enroll) EnrollStud() error {
	const queryInsertData = "insert into enroll (std_id,course_id,date_enrolled) values($1,$2,$3) returning std_id;"
	err := postgres.Db.QueryRow(queryInsertData, e.StdId, e.CourseId, e.Date).Scan(&e.StdId)
	return err
}

func (e *Enroll) GetEnrolledStudent(stdid int64, cid string) error {
	const queryGet = "select * from enroll where std_id = $1 and course_id = $2;"
	return postgres.Db.QueryRow(queryGet, stdid, cid).Scan(&e.StdId, &e.CourseId, &e.Date)
}

func GetAll() ([]Enroll, error) {
	const query = "select * from enroll;"
	table, err := postgres.Db.Query(query)
	if err != nil {
		return nil, err
	}
	enrolled := []Enroll{}
	for table.Next() {
		var e Enroll
		dbErr := table.Scan(&e.StdId, &e.CourseId, &e.Date)
		if dbErr != nil {
			return nil, dbErr
		}
		enrolled = append(enrolled, e)
	}
	table.Close()
	return enrolled, nil
}

func (e *Enroll) DeleteEnrolledStudent(stdid int64, cid string) error {
	const queryDelete = "delete from enroll where std_id = $1 and course_id = $2 returning std_id;"
	err := postgres.Db.QueryRow(queryDelete, stdid, cid).Scan(&e.StdId)
	return err
}
