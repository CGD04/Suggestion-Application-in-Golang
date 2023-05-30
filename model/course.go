package model

import (
	"myapp/datastore/postgres"
)


type Course struct {
	CourseId string `json:"id"`
	CourseName string `json:"courseName"`
}


func (c *Course) InsertData() error {
	const queryInsertCourse = "INSERT INTO course (CourseId,CourseName) VALUES ($1,$2);"
	_,err := postgres.Db.Exec(queryInsertCourse,c.CourseId,c.CourseName)
	return err
}

func (c *Course) GetInfo(a string) error {
	const queryGetCourse = "SELECT * FROM course WHERE CourseId = $1;"
	return postgres.Db.QueryRow(queryGetCourse,a).Scan(&c.CourseId,&c.CourseName)
}
func (c *Course) Update(a string) error {
	const queryUpdate = "UPDATE course SET CourseId = $1, CourseName = $2 WHERE CourseId = $3 RETURNING CourseId;"
	err := postgres.Db.QueryRow(queryUpdate,c.CourseId,c.CourseName,a).Scan(&c.CourseId)
	return err
}

func (c *Course) Delete(a string) error {
	const queryDeleteCourse = "DELETE FROM course WHERE CourseId = $1 RETURNING CourseId;"
	err := postgres.Db.QueryRow(queryDeleteCourse,a).Scan(&c.CourseId)
	return err
}

func GetCourses() ([]Course,error) {
	const queryGetCourses = "SELECT * FROM course "
	table,err := postgres.Db.Query(queryGetCourses)
	
	if err != nil {
		return nil, err
	}
	students := []Course{}

	for table.Next() {
		var c Course
		dbErr := table.Scan(&c.CourseId,&c.CourseName)
		if dbErr != nil {
			return nil, dbErr
		}
		students = append(students,c)
	}
	table.Close()
	return students, nil
}

func  Check(a string) (string){
	const query = "SELECT CourseId FROM course WHERE CourseId = $1;"
	var row string
	postgres.Db.QueryRow(query,a).Scan(&row)
	return row 
}