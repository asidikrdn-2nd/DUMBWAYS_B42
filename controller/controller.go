package controller

import (
	"dumbways-task_7/models"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"text/template"
	"time"

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

	content := map[string]interface{}{
		"DataProject": models.GetAllDataProject(),
	}

	// eksekusi template, kirim sebagai response. Dan menyisipkan isi variabel content sebagai datanya.
	homePage.Execute(w, content)
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
	// Handling data dari form
	if err := r.ParseMultipartForm(1024); err != nil {
		panic(err.Error())
	}

	// membuat variabel untuk menampung data project
	var project models.DataProject

	// Men-set id project dari waktu saat form disubmit dalam bentuk milisecond
	project.Id = strconv.FormatInt(time.Now().UnixNano(), 10)

	// mengambil data dari form untuk mengisi nilai milik variabel project
	project.ProjectName = r.FormValue("projectName")
	project.StartDate = r.FormValue("startDate")
	project.EndDate = r.FormValue("endDate")
	project.ProjectDesc = r.FormValue("projectDesc")

	// mengambil file dari form
	uploadedFile, handler, err := r.FormFile("projectImg")
	if err != nil {
		http.Error(w, "Please upload a JPEG or PNG image", http.StatusBadRequest)
		return
	}
	defer uploadedFile.Close()

	// Apabila format file bukan .jpg atau .png, maka tampilkan error
	if filepath.Ext(handler.Filename) != ".jpg" && filepath.Ext(handler.Filename) != ".png" {
		http.Error(w, "The provided file format is not allowed. Please upload a JPEG or PNG image", http.StatusBadRequest)
		return
	}

	// mengambil direktori aktif
	dir, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}

	// memberi nama pada file sesuai dengan Id project, agar nama filenya unik dan tidak ada yang sama
	filename := fmt.Sprintf("%s%s", project.Id, filepath.Ext(handler.Filename))
	project.Img = filename // men-set filename sebagai nilai Img dari variabel/object project
	// fmt.Println(filename)

	// menentukan lokasi file
	fileLocation := filepath.Join(dir, "assets/img", filename)
	// fmt.Println(fileLocation)

	// membuat file baru yang menjadi tempat untuk menampung hasil salinan file upload
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err.Error())
	}
	defer targetFile.Close()

	// Menyalin file hasil upload, ke file baru yang menjadi target
	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		panic(err.Error())
	}

	// periksa masing-masing nilai chechbox, apabila tidak kosong (terceklis) maka tambahka nilainya ke slice project.Tech
	if r.FormValue("html5") != "" {
		project.Tech = append(project.Tech, r.FormValue("html5"))
	}
	if r.FormValue("css3") != "" {
		project.Tech = append(project.Tech, r.FormValue("css3"))
	}
	if r.FormValue("js") != "" {
		project.Tech = append(project.Tech, r.FormValue("js"))
	}
	if r.FormValue("react") != "" {
		project.Tech = append(project.Tech, r.FormValue("react"))
	}

	// Menampilkan data project ke console terminal
	fmt.Println("========================================")
	fmt.Println("Project Name : " + project.ProjectName)
	fmt.Println("Project Start :", project.StartDate)
	fmt.Println("Project End :", project.EndDate)
	fmt.Println("Project Desc : " + project.ProjectDesc)
	fmt.Println("Project Tech :", project.Tech)
	fmt.Println("Project Image Name before Renamed : " + handler.Filename)
	fmt.Println("Project Image Name after Renamed : " + project.Img)
	fmt.Println("Project Image Location : " + fileLocation)
	fmt.Println("========================================")

	// menambahkan project baru ke local storage
	models.AddProject(project)

	// Redirect ke halaman index
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// fungsi untuk menghandle route edit-project page
func HandleEditProject(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Maap gaes, sementara fiturnya belum berfungsi"))
}

// fungsi untuk menghandle route delete-project page
func HandleDeleteProject(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Maap gaes, sementara fiturnya belum berfungsi"))
}

// fungsi untuk menghandle route project-detail page
func HandleProjectDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// mengambil nilai id yang ada di url
	id := mux.Vars(r)["id"]
	// fmt.Println(id)

	projectDetailPage, err := template.ParseFiles("views/project-detail.html")
	if err != nil {
		w.Write([]byte("Message :" + err.Error()))
		return
	}

	content := map[string]interface{}{
		"DataProject": models.GetDataProject(id),
	}

	projectDetailPage.Execute(w, content)
}
