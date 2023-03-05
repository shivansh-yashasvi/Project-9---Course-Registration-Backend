package routers

import (
	middleware "VIT/Middleware"
	"VIT/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func InitRouter() {
	///////////////////////Admin Route//////////////////////////////////////////////
	router.HandleFunc("/admin/faculty", middleware.IsAdminAuthorized(controllers.CreateFaculty)).Methods("POST")
	router.HandleFunc("/admin/course", middleware.IsAdminAuthorized(controllers.CreateCourse)).Methods("POST")
	router.HandleFunc("/admin/student", middleware.IsAdminAuthorized(controllers.CreateStudent)).Methods("POST")
	router.HandleFunc("/admin/slot", middleware.IsAdminAuthorized(controllers.CreateSlot)).Methods("POST")

	///////////////////////Student Route//////////////////////////////////////////////
	router.HandleFunc("/faculty/{faculty_id}", middleware.IsStudentAuthorized(controllers.GetFacultyByID)).Methods("GET")
	router.HandleFunc("/course/{course_id}", middleware.IsStudentAuthorized(controllers.GetCourseByID)).Methods("GET")
	router.HandleFunc("/timetable", middleware.IsStudentAuthorized(controllers.GetTimeTable)).Methods("GET")
	router.HandleFunc("/register", middleware.IsStudentAuthorized(controllers.CreateRegister)).Methods("POST")
	///////////////////////Student-Signin/////////////////////////////////
	router.HandleFunc("/student-signin", controllers.StudentSignIn).Methods("POST")
	//////////////////////signup,signin,logout///////////////////////////////////////
	router.HandleFunc("/signup", controllers.SignUp).Methods("POST")
	router.HandleFunc("/signin", controllers.SignIn).Methods("POST")
	router.HandleFunc("/logout", controllers.Logout).Methods("POST")

	http.ListenAndServe(":8080", router)

}
