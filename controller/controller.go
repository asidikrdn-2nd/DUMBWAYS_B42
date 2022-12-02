package controller

import (
	"dumbways-task_12/models"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
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

	var LoginState models.Login

	// mengambil nilai session dengan nama "LOGIN_SESSION"
	var store = sessions.NewCookieStore([]byte("LOGIN_SESSION"))
	session, _ := store.Get(r, "LOGIN_SESSION")

	// mengisikan nilai ke var LoginState sesuai kondisi login yang didapatkan dari session
	if session.Values["IsLogin"] != true {
		LoginState.IsLogin = false
	} else {
		LoginState.IsLogin = session.Values["IsLogin"].(bool)
		LoginState.UserName = session.Values["Name"].(string)
		LoginState.UserId = session.Values["Id"].(int)
	}

	// mengambil nilai flashes dengan key "message"
	fm := session.Flashes("message")
	// fmt.Println("session.Flashes :", fm)

	// mengambil tiap string dari slice yang dikirim flashes lalu menyimpannya ke dalam satu slice baru
	var flashes []string
	if len(fm) > 0 {
		session.Save(r, w)
		for _, fl := range fm {
			// fmt.Println("isi tiap element flashes :", fl)
			flashes = append(flashes, fl.(string))
		}
	}

	// menggabungkan semua string pada pesan flashes yang disimpan di slice tadi, untuk diparsing ke html
	LoginState.FlashData = strings.Join(flashes, "")

	// mengambil list project dari database
	dataProject := models.GetAllDataProject()

	// menyiapkan content yang akan di sajikan di halaman html
	responseContent := map[string]interface{}{
		"LoginState":  LoginState,
		"DataProject": dataProject,
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

	var LoginState models.Login

	// mengambil nilai session dengan nama "LOGIN_SESSION"
	var store = sessions.NewCookieStore([]byte("LOGIN_SESSION"))
	session, _ := store.Get(r, "LOGIN_SESSION")

	// mengisikan nilai ke var LoginState sesuai kondisi login yang didapatkan dari session
	if session.Values["IsLogin"] != true {
		LoginState.IsLogin = false
	} else {
		LoginState.IsLogin = session.Values["IsLogin"].(bool)
		LoginState.UserName = session.Values["Name"].(string)
		LoginState.UserId = session.Values["Id"].(int)
	}

	// menyiapkan content yang akan di sajikan di halaman html
	responseContent := map[string]interface{}{
		"LoginState": LoginState,
	}

	// eksekusi template, kirim sebagai response
	contactPage.Execute(w, responseContent)
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

	var LoginState models.Login

	// mengambil nilai session dengan nama "LOGIN_SESSION"
	var store = sessions.NewCookieStore([]byte("LOGIN_SESSION"))
	session, _ := store.Get(r, "LOGIN_SESSION")

	// mengisikan nilai ke var LoginState sesuai kondisi login yang didapatkan dari session
	if session.Values["IsLogin"] != true {
		LoginState.IsLogin = false
	} else {
		LoginState.IsLogin = session.Values["IsLogin"].(bool)
		LoginState.UserName = session.Values["Name"].(string)
		LoginState.UserId = session.Values["Id"].(int)
	}

	// menyiapkan content yang akan di sajikan di halaman html
	responseContent := map[string]interface{}{
		"LoginState": LoginState,
	}

	// eksekusi template, kirim sebagai response
	addProjectPage.Execute(w, responseContent)
}

// fungsi untuk menghandle route add-project page
func HandleAddProject(w http.ResponseWriter, r *http.Request) {
	// Handling dan parsing data dari form data yang ada data file nya. Argumen 1024 pada method tersebut adalah maxMemory
	if err := r.ParseMultipartForm(1024); err != nil {
		panic(err.Error())
	}

	// membuat variabel untuk menampung data project
	var project models.DataProject

	// mengambil data dari form untuk mengisi nilai milik variabel project
	project.ProjectName = r.FormValue("projectName")
	project.StartDate, _ = time.Parse("2006-01-02", r.FormValue("startDate")) // data tanggal yang didapat dari form langsung di parsing ke tipe data time.Time
	project.EndDate, _ = time.Parse("2006-01-02", r.FormValue("endDate"))     // data tanggal yang didapat dari form langsung di parsing ke tipe data time.Time
	project.ProjectDesc = r.FormValue("projectDesc")

	// periksa masing-masing nilai chechbox, apabila tidak kosong (terceklis) maka tambahkan nilainya ke slice project.Tech
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

	// validasi checkbox, apabila tidak ada yang di-ceklis maka akan menampilkan pesan error
	if len(project.Tech) <= 0 {
		http.Error(w, "Please choose one or more technologies checkbox", http.StatusBadRequest)
		return
	}

	// mengambil nilai session dengan nama "LOGIN_SESSION"
	var store = sessions.NewCookieStore([]byte("LOGIN_SESSION"))
	session, _ := store.Get(r, "LOGIN_SESSION")

	// men-set UserId pada project sesuai dengan id_user yang sedang login saat ini
	project.UserId = session.Values["Id"].(int)

	// membuat sebuah var untuk digunakan sebagai key pada context (untuk mengatasi warning should not use built-in type string as key)
	var UploadFileID models.ContextKey = "FileName"

	// mengambil data dari context yang dikirim oleh middleware.UploadFile
	dataContex := r.Context().Value(UploadFileID)

	// menggunakan filename yang didapat untuk mengisi property Img milik object project
	project.Img = dataContex.(string)

	// menambahkan project baru ke local storage
	models.AddProject(project)

	// Redirect ke halaman index
	http.Redirect(w, r, "/", http.StatusSeeOther)
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

		var LoginState models.Login

		// mengambil nilai session dengan nama "LOGIN_SESSION"
		var store = sessions.NewCookieStore([]byte("LOGIN_SESSION"))
		session, _ := store.Get(r, "LOGIN_SESSION")

		// mengisikan nilai ke var LoginState sesuai kondisi login yang didapatkan dari session
		if session.Values["IsLogin"] != true {
			LoginState.IsLogin = false
		} else {
			LoginState.IsLogin = session.Values["IsLogin"].(bool)
			LoginState.UserName = session.Values["Name"].(string)
			LoginState.UserId = session.Values["Id"].(int)
		}

		// mengambil data project dengan id yang sesuai dari database
		dataProject := models.GetDataProject(id)

		// menyiapkan content yang akan di sajikan di halaman html
		responseContent := map[string]interface{}{
			"LoginState": LoginState,
			"DataProject": map[string]interface{}{
				"Id":          dataProject.Id,
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
		// Handling dan parsing data dari form data yang ada data file nya. Argumen 1024 pada method tersebut adalah maxMemory
		if err := r.ParseMultipartForm(1024); err != nil {
			panic(err.Error())
		}

		// membuat variabel untuk menampung data project
		var project models.DataProject

		// mengambil nilai id yang ada di url
		project.Id, _ = strconv.Atoi(mux.Vars(r)["id"])
		// fmt.Println(project.Id)

		// mengambil data dari form untuk mengisi nilai milik variabel project
		project.ProjectName = r.FormValue("projectName")
		project.StartDate, _ = time.Parse("2006-01-02", r.FormValue("startDate")) // data tanggal yang didapat dari form langsung di parsing ke tipe data time.Time
		project.EndDate, _ = time.Parse("2006-01-02", r.FormValue("endDate"))     // data tanggal yang didapat dari form langsung di parsing ke tipe data time.Time
		project.ProjectDesc = r.FormValue("projectDesc")

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

		// validasi checkbox, apabila tidak ada yang di-ceklis maka akan menampilkan pesan error
		if len(project.Tech) <= 0 {
			http.Error(w, "Please choose one or more technologies checkbox", http.StatusBadRequest)
			return
		}

		// membuat sebuah var untuk digunakan sebagai key pada context (untuk mengatasi warning should not use built-in type string as key)
		var UpdateFileID models.ContextKey = "FileName"

		// mengambil data dari context yang dikirim oleh middleware.UploadFile
		dataContex := r.Context().Value(UpdateFileID)

		// menggunakan filename yang didapat untuk mengisi property Img milik object project
		project.Img = dataContex.(string)

		// menambahkan project baru ke local storage
		models.EditProject(project)

		// Redirect ke halaman index
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// fungsi untuk menghandle route delete-project page
func HandleDeleteProject(w http.ResponseWriter, r *http.Request) {
	// mengambil nilai id yang ada di url dan langsung di konversi ke bentuk integer
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	// fmt.Println(id)

	// menjalankan fungsi DeleteProject yang ada di package models
	models.DeleteProject(id)

	// redirect ke halaman index
	http.Redirect(w, r, "/", http.StatusSeeOther)
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

	var LoginState models.Login

	// mengambil nilai session dengan nama "LOGIN_SESSION"
	var store = sessions.NewCookieStore([]byte("LOGIN_SESSION"))
	session, _ := store.Get(r, "LOGIN_SESSION")

	// mengisikan nilai ke var LoginState sesuai kondisi login yang didapatkan dari session
	if session.Values["IsLogin"] != true {
		LoginState.IsLogin = false
	} else {
		LoginState.IsLogin = session.Values["IsLogin"].(bool)
		LoginState.UserName = session.Values["Name"].(string)
		LoginState.UserId = session.Values["Id"].(int)
	}

	// menyiapkan content yang akan di sajikan di halaman html
	responseContent := map[string]interface{}{
		"LoginState":  LoginState,
		"DataProject": models.GetDataProject(id),
	}

	// mengeksekusi template sebagai response dan mengirim responseContent untuk disajikan di halaman html
	projectDetailPage.Execute(w, responseContent)
}

// fungsi untuk menghandle register (tampilan dan fungsi register)
func HandleRegisterForm(w http.ResponseWriter, r *http.Request) {
	// Apabila methodnya GET, maka tampilkan form edit project
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		// menyiapkan template html
		formRegisterPage, err := template.ParseFiles("views/register.html")
		if err != nil {
			w.Write([]byte("Message :" + err.Error()))
			return
		}

		var LoginState models.Login

		// mengambil nilai session dengan nama "LOGIN_SESSION"
		var store = sessions.NewCookieStore([]byte("LOGIN_SESSION"))
		session, _ := store.Get(r, "LOGIN_SESSION")

		// mengisikan nilai ke var LoginState sesuai kondisi login yang didapatkan dari session
		if session.Values["IsLogin"] != true {
			LoginState.IsLogin = false
		} else {
			LoginState.IsLogin = session.Values["IsLogin"].(bool)
			LoginState.UserName = session.Values["Name"].(string)
		}

		// menyiapkan content yang akan di sajikan di halaman html
		responseContent := map[string]interface{}{
			"LoginState": LoginState,
		}

		// mengeksekusi template sebagai response untuk disajikan di halaman html
		formRegisterPage.Execute(w, responseContent)
	} else if r.Method == "POST" { // Apabila methodnya POST, artinya form sudah di submit, maka handle data yang didapat dari form tersebut
		// parsing data dari form
		if err := r.ParseForm(); err != nil {
			log.Fatal(err)
		}

		// membuat variabel dengan struct User
		var newUserAccount models.User

		// ambil data dari form
		newUserAccount.Name = r.FormValue("name")
		newUserAccount.Email = r.FormValue("email")
		password := r.FormValue("password")

		// enkripsi password dengan bcrypt
		newUserAccount.Password, _ = bcrypt.GenerateFromPassword([]byte(password), 10)

		// menyimpan data ke UserAccount di package models
		models.AddUserAccount(newUserAccount)

		// coba liat hasil data nya di postman
		// json.NewEncoder(w).Encode(newUserAccount)

		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

// fungsi untuk menghandle login (tampilan dan fungsi login)
func HandleLoginForm(w http.ResponseWriter, r *http.Request) {
	// Apabila methodnya GET, maka tampilkan form edit project
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		// menyiapkan template html
		formLoginPage, err := template.ParseFiles("views/login.html")
		if err != nil {
			w.Write([]byte("Message :" + err.Error()))
			return
		}

		var LoginState models.Login

		// mengambil nilai session dengan nama "LOGIN_SESSION"
		var store = sessions.NewCookieStore([]byte("LOGIN_SESSION"))
		session, _ := store.Get(r, "LOGIN_SESSION")

		// mengisikan nilai ke var LoginState sesuai kondisi login yang didapatkan dari session
		if session.Values["IsLogin"] != true {
			LoginState.IsLogin = false
		} else {
			LoginState.IsLogin = session.Values["IsLogin"].(bool)
			LoginState.UserName = session.Values["Name"].(string)
		}

		// menyiapkan content yang akan di sajikan di halaman html
		responseContent := map[string]interface{}{
			"LoginState": LoginState,
		}

		// mengeksekusi template sebagai response untuk disajikan di halaman html
		formLoginPage.Execute(w, responseContent)
	} else if r.Method == "POST" { // Apabila methodnya POST, artinya form sudah di submit, maka handle data yang didapat dari form tersebut
		// parsing data dari form
		if err := r.ParseForm(); err != nil {
			log.Fatal(err)
		}

		// ambil data dari form
		email := r.FormValue("email")
		password := r.FormValue("password")

		// mengambil detail akun dari database
		userAccountLogin, errGetUser := models.GetUserDataByEmail(email)
		if errGetUser != nil {
			http.Error(w, "User Not Found", http.StatusNotFound)
			return
		}

		// compare password yang terdaftar dengan password yang diinputkan saat login
		err := bcrypt.CompareHashAndPassword([]byte(userAccountLogin.Password), []byte(password))
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		// mengambil nilai session dengan ID "LOGIN_SESSION"
		var store = sessions.NewCookieStore([]byte("LOGIN_SESSION"))
		session, _ := store.Get(r, "LOGIN_SESSION")

		// men-set nilai pada session LOGIN_SESSION
		session.Values["IsLogin"] = true
		session.Values["Name"] = userAccountLogin.Name
		session.Values["Id"] = userAccountLogin.Id
		session.Options.MaxAge = 10800 // dalam satuan detik, 10800 detik = 3 jam

		// menambahkan flash, arg1 = isi pesan, arg2 = key/identifier dari pesan yang dikirim
		session.AddFlash("Login success", "message")

		// menyimpan session tersebut
		session.Save(r, w)

		// redirect ke halaman home
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

// fungsi untuk menghandle route logout untuk menghapus session
func HandleLogout(w http.ResponseWriter, r *http.Request) {
	// mengambil session "LOGIN_SESSION" dari cookies
	var store = sessions.NewCookieStore([]byte("LOGIN_SESSION"))
	session, _ := store.Get(r, "LOGIN_SESSION")
	// men-set umur cookies tersebut menjadi -1 (expired)
	session.Options.MaxAge = -1
	// menyimpan session tersebut
	session.Save(r, w)

	// redirect ke halaman home
	http.Redirect(w, r, "/", http.StatusFound)
}
