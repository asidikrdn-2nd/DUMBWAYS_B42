package models

import (
	"context"
	"dumbways-task_9/config"
	"fmt"
	"strconv"
	"time"
)

// membuat struct untuk data project
type DataProject struct {
	Id          int
	ProjectName string
	StartDate   time.Time
	EndDate     time.Time
	ProjectDesc string
	Tech        []string
	Img         string
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

// fungsi untuk mengambil/menampilkan semua data
func GetAllDataProject() []DataProject {
	// koneksi ke database
	db, err := config.CreateConnection()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close(context.Background()) // menutup database di akhir, setelah tidak digunakan

	// mengambil data dari database
	dataQuery, errQuery := db.Query(context.Background(), "SELECT id, name, start_date, end_date, description, technologies, image FROM tb_projects")
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
		err := dataQuery.Scan(&eachData.Id, &eachData.ProjectName, &eachData.StartDate, &eachData.EndDate, &eachData.ProjectDesc, &eachData.Tech, &eachData.Img)
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
	errScanRow := db.QueryRow(context.Background(), "SELECT id, name, start_date, end_date, description, technologies, image FROM tb_projects WHERE id = $1", idProject).Scan(&ProjectList.Id, &ProjectList.ProjectName, &ProjectList.StartDate, &ProjectList.EndDate, &ProjectList.ProjectDesc, &ProjectList.Tech, &ProjectList.Img)
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
	fmt.Println("Feature is Not Available")
}

// fungsi untuk mengubah data
func EditProject(newDataProject DataProject) {
	fmt.Println("Feature is Not Available")
}

// fungsi untuk menghapus data
func DeleteProject(idProject int) {
	fmt.Println("Feature is Not Available")
}
