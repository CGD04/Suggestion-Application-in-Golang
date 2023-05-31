package routes

import (
	"fmt"
	"log"
	controller "myapp/controllers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func InitializeRoutes() {
	router := mux.NewRouter()
	//routes for students
	router.HandleFunc("/student", controller.AddStudent).Methods("POST")
	router.HandleFunc("/student/{stdID}", controller.GetStud).Methods("GET")
	router.HandleFunc("/student/{stdid}", controller.UpdateHandler).Methods("PUT")
	router.HandleFunc("/student/{stdid}", controller.DeleteStudent).Methods("DELETE")
	router.HandleFunc("/students", controller.GetAllStuds).Methods("GET")

	//routes for courses
	router.HandleFunc("/course", controller.AddCourse).Methods("POST")
	router.HandleFunc("/course/{id}", controller.GetCourse).Methods("GET")
	router.HandleFunc("/course/{id}", controller.UpdateCourse).Methods("PUT")
	router.HandleFunc("/course/{id}", controller.DeleteCourse).Methods("DELETE")
	router.HandleFunc("/courses", controller.GetAllCourses).Methods("GET")

	//route for signup
	router.HandleFunc("/signup", controller.AddUserHandler).Methods("POST")
	router.HandleFunc("/login", controller.LoginHandler).Methods("POST")

	router.HandleFunc("/logout", controller.LogoutHandler).Methods("GET")

	router.HandleFunc("/enroll", controller.EnrollStudent).Methods("POST")
	router.HandleFunc("/showEnrolled/{stdid}/{cid}", controller.ShowEnrolled).Methods("GET")
	router.HandleFunc("/deleteEnrolled/{stdid}/{cid}", controller.DeleteEnrolled).Methods("DELETE")
	router.HandleFunc("/showEnrolledAll", controller.ShowAllEnrolled).Methods("GET")

	// to serve static file
	fhandler := http.FileServer(http.Dir("./view"))
	router.PathPrefix("/").Handler(fhandler)

	err := http.ListenAndServe(":8080", router)
	log.Println("Application is running on port 8080........")
	if err != nil {
		fmt.Println("problem with listenandserve")
		os.Exit(1)
	}
}
