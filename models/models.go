package models

import (
	"context"
	"dumbways-task_12/config"
	"fmt"
	"strconv"
	"time"
)

// ================================
// DATA STRUCT
// ================================

// membuat tipe data untuk context key
type ContextKey string

// membuat struct untuk data project
type DataProject struct {
	Id          int
	ProjectName string
	StartDate   time.Time
	EndDate     time.Time
	ProjectDesc string
	Tech        []string
	Img         string
	UserId      int
	UserName    string
}

// menambahkan method ShowDate pada struct DataProject, method ini digunakan untuk menampilkan tanggal dengan format yang lebih rapih untuk dibaca
func (p DataProject) ShowDate(date time.Time) string {
	month := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	return strconv.Itoa(date.Day()) + " " + month[date.Month()-1] + " " + strconv.Itoa(date.Year())
}

// menambahkan method ShowDuration pada struct DataProject, method ini digunakan untuk menampilkan durasi/selisih waktu pengerjaan project
func (p DataProject) ShowDuration(startDate time.Time, endDate time.Time) string {
	// menghitung durasi menggunakan method .Sub milik object time.Time
	duration := endDate.Sub(startDate).Hours()

	var day int = 0
	var month int = 0
	var year int = 0

	// setiap 24 jam menjadi 1 hari
	for duration >= 24 {
		day += 1
		duration -= 24
	}
	// setiap 30 hari menjadi 1 bulan
	for day >= 30 {
		month += 1
		day -= 30
	}
	// setiap 12 bulan menjadi 1 tahun
	for month >= 12 {
		year += 1
		month -= 12
	}

	if year != 0 && month != 0 {
		return strconv.Itoa(year) + " Year, " + strconv.Itoa(month) + " Month, " + strconv.Itoa(day) + " Day"
	} else if month != 0 {
		return strconv.Itoa(month) + " Month, " + strconv.Itoa(day) + " Day"
	} else {
		return strconv.Itoa(day) + " Day"
	}
}

// membuat struct untuk user account
type User struct {
	Id       int
	Name     string
	Email    string
	Password []byte
}

// membuat struct untuk login status
type Login struct {
	IsLogin   bool
	UserName  string
	UserId    int
	FlashData string
}

// ================================
// END OF DATA STRUCT
// ================================

// ================================
// FUNCTION
// ================================

// fungsi untuk mengambil/menampilkan semua data
func GetAllDataProject() []DataProject {
	// koneksi ke database
	db, err := config.CreateConnection()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close(context.Background()) // menutup database di akhir, setelah tidak digunakan

	// mengambil data dari database
	dataQuery, errQuery := db.Query(context.Background(), `SELECT tb_projects.id_project, tb_projects.name, tb_projects.start_date, tb_projects.end_date, 
		tb_projects.description, tb_projects.technologies, tb_projects.image, tb_projects.id_user, tb_users.name 
	FROM tb_projects LEFT JOIN tb_users
	ON tb_projects.id_user = tb_users.id_user
	ORDER BY id_project ASC`)
	if errQuery != nil {
		panic(errQuery.Error())
	}

	// membuat storage kumpulan/array data project yang ditampung pada variabel data
	var ProjectList []DataProject

	// melakukan perulangan pada hasil query sesuai jumlah baris data yang didapat
	for dataQuery.Next() {
		// menyiapkan variabel untuk menampung tiap baris data hasil query
		var eachData DataProject

		// mengambil tiap baris data hasil query
		err := dataQuery.Scan(&eachData.Id, &eachData.ProjectName, &eachData.StartDate, &eachData.EndDate, &eachData.ProjectDesc, &eachData.Tech, &eachData.Img, &eachData.UserId, &eachData.UserName)
		if err != nil {
			panic(err.Error())
		}

		// menambahkan tiap baris data ke projectList
		ProjectList = append(ProjectList, eachData)
	}

	// mengembalikan kumpulan data project
	return ProjectList
}

// fungsi untuk mengambil/menampilkan satu data
func GetDataProject(idProject int) DataProject {
	// fmt.Println(idProject)

	// koneksi ke database
	db, err := config.CreateConnection()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close(context.Background()) // menutup database di akhir, setelah tidak digunakan

	// membuat storage kumpulan/array data project yang ditampung pada variabel data
	var ProjectList DataProject

	// mengambil 1 baris data dari database
	errScanRow := db.QueryRow(context.Background(), "SELECT id_project, name, start_date, end_date, description, technologies, image FROM tb_projects WHERE id_project = $1", idProject).Scan(&ProjectList.Id, &ProjectList.ProjectName, &ProjectList.StartDate, &ProjectList.EndDate, &ProjectList.ProjectDesc, &ProjectList.Tech, &ProjectList.Img)
	if errScanRow != nil {
		panic(errScanRow.Error())
	}
	/**
	Karena hanya 1 baris, kita bisa menggunakan QueryRow. QueryRow hanya mengembalikan row data (tanpa mengembalikan error), sehingga data tsb bisa langsung kita Scan.
	*/

	// mengembalikan 1 baris data project
	return ProjectList
}

// fungsi untuk menambahkan data
func AddProject(newDataProject DataProject) {
	// koneksi ke database
	db, err := config.CreateConnection()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close(context.Background()) // menutup database di akhir, setelah tidak digunakan

	// Menyimpan data project baru ke database
	_, err = db.Exec(context.Background(),
		"INSERT INTO tb_projects (name, start_date, end_date, description, technologies, image, id_user) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		newDataProject.ProjectName, newDataProject.StartDate, newDataProject.EndDate, newDataProject.ProjectDesc, newDataProject.Tech, newDataProject.Img, newDataProject.UserId)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("New project has been added")
}

// fungsi untuk mengubah data
func EditProject(newDataProject DataProject) {
	// koneksi ke database
	db, err := config.CreateConnection()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close(context.Background()) // menutup database di akhir, setelah tidak digunakan

	// Mengupdate data project yang memiliki id sesuai
	_, err = db.Exec(context.Background(),
		"UPDATE tb_projects SET name=$1, start_date=$2, end_date=$3, description=$4, technologies=$5, image=$6 WHERE id_project=$7",
		newDataProject.ProjectName, newDataProject.StartDate, newDataProject.EndDate, newDataProject.ProjectDesc, newDataProject.Tech, newDataProject.Img, newDataProject.Id)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Project with id", newDataProject.Id, "has been edited")
}

// fungsi untuk menghapus data
func DeleteProject(idProject int) {
	// fmt.Println(idProject)

	// koneksi ke database
	db, err := config.CreateConnection()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close(context.Background()) // menutup database di akhir, setelah tidak digunakan

	// Menghapus data pada database yang memiliki id sesuai
	_, err = db.Exec(context.Background(), "DELETE FROM tb_projects WHERE id_project=$1", idProject)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Project with id", idProject, "has been deleted")
}

// fungsi untuk menambahkan data user baru ke database
func AddUserAccount(newUser User) {
	// koneksi ke database
	db, err := config.CreateConnection()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close(context.Background()) // menutup database di akhir, setelah tidak digunakan

	// Menghapus data pada database yang memiliki id sesuai
	_, err = db.Exec(context.Background(), "INSERT INTO tb_users (name, email, password) VALUES ($1, $2, $3)", newUser.Name, newUser.Email, newUser.Password)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("New user has been added")

	// fmt.Println(UserAccount)
}

// fungsi untuk mengambil/menampilkan data user berdasarkan email tertentu
func GetUserDataByEmail(email string) (User, error) {
	// koneksi ke database
	db, err := config.CreateConnection()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close(context.Background()) // menutup database di akhir, setelah tidak digunakan

	// membuat storage kumpulan/array data project yang ditampung pada variabel data
	var UserAccount User

	// mengambil 1 baris data dari database
	errScanRow := db.QueryRow(context.Background(), "SELECT id_user, name, email, password FROM tb_users WHERE email = $1", email).Scan(&UserAccount.Id, &UserAccount.Name, &UserAccount.Email, &UserAccount.Password)
	if errScanRow != nil {
		// panic(errScanRow.Error())
		return UserAccount, errScanRow
	}
	/**
	Karena hanya 1 baris, kita bisa menggunakan QueryRow. QueryRow hanya mengembalikan row data (tanpa mengembalikan error), sehingga data tsb bisa langsung kita Scan.
	*/

	// mengembalikan 1 baris data project
	return UserAccount, nil
}

// ================================
// END OF FUNCTION
// ================================
