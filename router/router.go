package router

import (
	"dumbways-task_12/controller"
	"dumbways-task_12/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	route := mux.NewRouter()

	// membuat routing file server (berguna untuk mengakses file-file static/aset yang dibutuhkan, misalnya file css, gambar, js, dll)
	route.PathPrefix("/static/").Handler(http.StripPrefix("/static/", // semua request yang masuk ke endpoint `/static` akan diarahkan ke actual handler
		http.FileServer(http.Dir("./assets")))) // actual handler
	/**
	StripPrefix berguna untuk menghapus prefix pada endpoint yang di-req agar path file menjadi valid, contoh (sesuai dengan kondisi diatas) :
	Tanpa StripPrefix ketika mengakses endpoint `/static/` maka path yang diakses adalah => `assets/static/...` atau => `./static/...` relatif dari folder `assets`
	Dengan StripPrefix `/static` ketika mengakses endpoint `/static/` maka path yang diakses adalah => `assets/...` atau => `./...` relatif dari folder `assets`
	*/

	// routing home page
	route.HandleFunc("/", controller.HandleHome)

	// routing contact page
	route.HandleFunc("/contact", controller.HandleContact)

	// routing project form page
	route.HandleFunc("/project", controller.HandleProjectForm)

	// routing project detail page
	route.HandleFunc("/project-detail/{id}", controller.HandleProjectDetail)

	// routing add project page
	route.HandleFunc("/add-project", middleware.UploadFile(controller.HandleAddProject)).Methods("POST") // menginclude kan middleware ke controller HandleAddProject

	// routing edit project page
	route.HandleFunc("/edit-project/{id}", middleware.UpdateFile(controller.HandleEditProject)).Methods("POST", "GET")

	// routing delete project page
	route.HandleFunc("/delete-project/{id}", controller.HandleDeleteProject)

	// routing register page
	route.HandleFunc("/register", controller.HandleRegisterForm).Methods("POST", "GET")

	// routing login page
	route.HandleFunc("/login", controller.HandleLoginForm)

	// routing logout
	route.HandleFunc("/logout", controller.HandleLogout)

	return route
}
