package model

import "myapp/datastore/postgres"


type Admin struct {
	FirstName string `json:"fname"`
	LastName string `json:"lname"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func (a *Admin) Create() error {
	const queryCreateUser = "INSERT INTO admin (FirstName,LastName,Email,Password) VALUES($1,$2,$3,$4)"
	_,err := postgres.Db.Exec(queryCreateUser,a.FirstName,a.LastName,a.Email,a.Password)
	return err
}

func (a *Admin) Check(email string) error {
	const queryCheck = "Select * from admin where Email = $1;"
	err := postgres.Db.QueryRow(queryCheck,email).Scan(&a.FirstName,&a.LastName,&a.Email,&a.Password)
	return err
}