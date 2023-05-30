package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"myapp/model"
	"myapp/utils/date"
	"myapp/utils/httpResponse"
	"net/http"

	"github.com/gorilla/mux"
)

func EnrollStudent(w http.ResponseWriter, r *http.Request) {
	var e model.Enroll
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		httpResponse.ResponseWithError(w, http.StatusBadRequest, "invalid json")
		fmt.Println("invalid json")
		return
	}
	e.Date = date.GetDate()
	defer r.Body.Close()
	if enrollErr := e.EnrollStud(); enrollErr != nil {
		httpResponse.ResponseWithError(w, http.StatusBadRequest, "invalid query")
		fmt.Println("cannot enroll", enrollErr)
		return
	}
	httpResponse.ResponseWithJson(w, http.StatusCreated, map[string]string{"message": "successfully inserted the data"})
	fmt.Println("last")
}

func ShowEnrolled(w http.ResponseWriter, r *http.Request) {
	para := mux.Vars(r)
	stdid := para["stdid"]
	cid := para["cid"]

	stdId, err := getUserID(stdid)
	if err != nil {
		fmt.Println("error in converting the stdid to integer", err)
		return
	}
	var enroll model.Enroll
	if getErr := enroll.GetEnrolledStudent(stdId, cid); getErr != nil {
		httpResponse.ResponseWithError(w, http.StatusBadRequest, "error in getting the database")
		fmt.Println("kkkk", getErr)
		return
	}
	httpResponse.ResponseWithJson(w, http.StatusOK, enroll)
}

func ShowAllEnrolled(w http.ResponseWriter, r *http.Request) {
	if !VerifyCookie(w, r) {
		return
	}
	enroll, getErr := model.GetAll()
	if getErr != nil {
		switch getErr {
		case sql.ErrNoRows:
			httpResponse.ResponseWithError(w, http.StatusNotFound, "no data ")
			fmt.Println("no registerd or enrolled")
		default:
			fmt.Println("error in getting from database")
		}
		return
	}
	httpResponse.ResponseWithJson(w, http.StatusOK, enroll)
}

func DeleteEnrolled(w http.ResponseWriter, r *http.Request) {
	para := mux.Vars(r)
	stdid := para["stdid"]
	cid := para["cid"]

	stdId, err := getUserID(stdid)
	if err != nil {
		fmt.Println("error in converting the stdid to integer", err)
		return
	}
	var enroll model.Enroll
	fmt.Println("here")
	if getErr := enroll.DeleteEnrolledStudent(stdId, cid); getErr != nil {
		httpResponse.ResponseWithError(w, http.StatusBadRequest, "error in deleting the database")
		fmt.Println("There is an error in", getErr)
		return
	}
	httpResponse.ResponseWithJson(w, http.StatusOK, map[string]string{"message": "deleted"})
}
