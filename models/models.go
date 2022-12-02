package models

import (
	"fmt"
	"strconv"
	"time"
)

// membuat struct untuk data project
type DataProject struct {
	Id          string   `json:"id"`
	ProjectName string   `json:"projectName"`
	StartDate   string   `json:"startDate"`
	EndDate     string   `json:"endDate"`
	ProjectDesc string   `json:"projectDesc"`
	Tech        []string `json:"tech"`
	Img         string   `json:"img"`
}

// menambahkan method ShowDate pada struct DataProject, method ini digunakan untuk menampilkan tanggal dengan format yang lebih rapih untuk dibaca
func (p DataProject) ShowDate(date string) string {
	// parsing string menjadi time.Time (sama seeperti fungsi new Date() pada javascript)
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		panic(err.Error())
	}

	month := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	return strconv.Itoa(t.Day()) + " " + month[t.Month()-1] + " " + strconv.Itoa(t.Year())
}

// menambahkan method ShowDuration pada struct DataProject, method ini digunakan untuk menampilkan durasi/selisih waktu pengerjaan project
func (p DataProject) ShowDuration(startDate string, endDate string) string {
	// parsing string menjadi time.Time (sama seeperti fungsi new Date() pada javascript)
	tStart, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		panic(err.Error())
	}
	// parsing string menjadi time.Time (sama seeperti fungsi new Date() pada javascript)
	tEnd, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		panic(err.Error())
	}

	duration := tEnd.Sub(tStart).Hours()

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

// membuat storage kumpulan/array data project yang ditampung pada variabel data
var ProjectList = []DataProject{
	{"01",
		"Mini Expense Tracker",
		"2022-01-14",
		"2022-03-14",
		`"Expense Tracker" merujuk ke sistem pencatatan pengeluaran, aplikasi ini akan menyimpan pemasukan dan pengeluaran user serta menampilkan total selisih antara keduanya.`,
		[]string{"html5", "css3", "js", "react"},
		"mini-expense-tracker.png"},
	{"02",
		"CRUD Mahasiswa",
		"2022-02-14",
		"2022-03-09",
		`Dalam programming, CRUD merupakan singkatan dari Create Read Update dan Delete. Yakni aplikasi yang berisi pengolahan data. 
		Biasanya CRUD butuh database sebagai media penyimpanan. Akan tetapi untuk sementara ini app CRUD Mahasiswa lebih fokus ke kode React, CRUD ini hanya disimpan di memory saja.`,
		[]string{"html5", "js", "react"},
		"crud-mahasiswa.png"},
	{"03",
		"Landing Page",
		"2021-10-14",
		"2022-03-14",
		`Landing Page sederhana yang berisikan informasi perusahaan untuk menarik pelanggan. Untuk sementara ini data yang ada di dalam landing page masih bersifat dummy, 
		karena proyek ini merupakan hasil belajar HTML CSS dan Bootstrap. Akan tetapi proyek ini dapat dijadikan template apabila ada proyek serupa kedepannya.`,
		[]string{"html5", "css3"},
		"landingpage.png"},
	{"04",
		"YouNoob",
		"2021-02-14",
		"2022-03-14",
		`Lorem ipsum dolor sit amet consectetur adipisicing elit. Iure voluptatibus fugiat, veniam magnam eos pariatur earum illo odit,
    eius voluptas expedita. Ullam, repellendus inventore exercitationem perferendis beatae ea enim ad, nemo tempora
    ducimus sunt in, eaque illum eius quo necessitatibus non delectus nam? Inventore, aliquam.`,
		[]string{"html5", "css3", "js"},
		"younoob.png"},
}

// fungsi untuk mengambil/menampilkan semua data
func GetAllDataProject() []DataProject {
	data := ProjectList

	// mengembalikankumpulan data project
	return data
}

// fungsi untuk mengambil/menampilkan satu data
func GetDataProject(idProject string) DataProject {
	// fmt.Println(idProject)

	var data DataProject
	// mencari project yang memiliki id sesuai
	for _, project := range ProjectList {
		if project.Id == idProject {
			data = project
		}
	}

	// mengembalikankumpulan data project
	return data
}

// fungsi untuk menambahkan data
func AddProject(newDataProject DataProject) {
	ProjectList = append(ProjectList, newDataProject)

	fmt.Println("Data Berhasil ditambahkan")
}

// fungsi untuk mengubah data
func EditProject(idProject string) {
	fmt.Println("Sementara belum bisa mengubah data")
}

// fungsi untuk menghapus data
func DeleteProject(idProject string) {
	fmt.Println("Sementara belum bisa menghapus data")
}
