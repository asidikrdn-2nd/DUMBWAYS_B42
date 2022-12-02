package controller

import (
	"dumbways-task_9/models"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

// fungsi untuk menghandle route home page
func HandleHome(w http.ResponseWriter, r *http.Request) {
	// Men-set tipe kontennya sebagai html
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// mengambil file index.hml sebagai template
	homePage, err := template.ParseFiles("views/index.html")
	// apabila file index.hml tidak ditemukan, tampilkan pesan error
	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	// menyiapkan content yang akan di sajikan di halaman html
	responseContent := map[string]interface{}{
		"DataProject": models.GetAllDataProject(),
	}

	// eksekusi template, kirim sebagai response. Dan menyisipkan isi variabel responseContent sebagai datanya.
	homePage.Execute(w, responseContent)
}

// fungsi untuk menghandle route contact page
func HandleContact(w http.ResponseWriter, r *http.Request) {
	// Men-set tipe kontennya sebagai html
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// mengambil file index.hml sebagai template
	contactPage, err := template.ParseFiles("views/contact.html")
	// apabila file index.hml tidak ditemukan, tampilkan pesan error
	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	// eksekusi template, kirim sebagai response
	contactPage.Execute(w, nil)
}

// fungsi untuk menghandle route project page
func HandleProjectForm(w http.ResponseWriter, r *http.Request) {
	// Men-set tipe kontennya sebagai html
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// mengambil file index.hml sebagai template
	addProjectPage, err := template.ParseFiles("views/project.html")
	// apabila file index.hml tidak ditemukan, tampilkan pesan error
	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	// eksekusi template, kirim sebagai response
	addProjectPage.Execute(w, nil)
}

// fungsi untuk menghandle route add-project page
func HandleAddProject(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Maaf ya gaes, sementara ini fitur belum tersedia"))
}

// fungsi untuk menghandle route edit-project page
func HandleEditProject(w http.ResponseWriter, r *http.Request) {
	// Apabila methodnya GET, maka tampilkan form edit project
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		// mengambil nilai id yang ada di url
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		// fmt.Println(id)

		// menyiapkan template html
		formEditProjectPage, err := template.ParseFiles("views/edit-project.html")
		if err != nil {
			w.Write([]byte("Message :" + err.Error()))
			return
		}

		// mengambil data project dengan id yang sesuai dari database
		dataProject := models.GetDataProject(id)

		// menyiapkan content yang akan di sajikan di halaman html
		responseContent := map[string]interface{}{
			"DataProject": map[string]interface{}{
				"ID":          dataProject.Id,
				"ProjectName": dataProject.ProjectName,
				"StartDate":   dataProject.StartDate.Format("2006-01-02"),
				"EndDate":     dataProject.EndDate.Format("2006-01-02"),
				"ProjectDesc": dataProject.ProjectDesc,
				"Tech":        dataProject.Tech,
				"Img":         dataProject.Img,
			},
		}

		// mengeksekusi template sebagai response dan mengirim responseContent untuk disajikan di halaman html
		formEditProjectPage.Execute(w, responseContent)
	} else if r.Method == "POST" { // Apabila methodnya POST, artinya form sudah di submit, maka handle data yang didapat dari form tersebut
		w.Write([]byte("Maaf ya gaes, sementara ini fitur belum tersedia"))
	}
}

// fungsi untuk menghandle route delete-project page
func HandleDeleteProject(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Maaf ya gaes, sementara ini fitur belum tersedia"))
}

// fungsi untuk menghandle route project-detail page
func HandleProjectDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// mengambil nilai id yang ada di url dan langsung di konversi ke bentuk integer
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	// fmt.Println(id)

	// menyiapkan template html
	projectDetailPage, err := template.ParseFiles("views/project-detail.html")
	if err != nil {
		w.Write([]byte("Message :" + err.Error()))
		return
	}

	// menyiapkan content yang akan di sajikan di halaman html
	responseContent := map[string]interface{}{
		"DataProject": models.GetDataProject(id),
	}

	// mengeksekusi template sebagai response dan mengirim responseContent untuk disajikan di halaman html
	projectDetailPage.Execute(w, responseContent)
}
