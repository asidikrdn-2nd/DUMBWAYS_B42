package middleware

import (
	"context"
	"dumbways-task_12/models"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// membuat middleware untuk menghandle upload file
func UploadFile(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// mengambil nama project dari inputan form
		projectName := r.FormValue("projectName")

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

		// menyusun string untuk digunakan sebagai nama file gambar
		filenameStr := projectName                     // mengambil nama project
		filenameStr = strings.ToLower(filenameStr)     // mem-format nama project menjadi huruf kecil
		filenameArr := strings.Split(filenameStr, " ") // memisahkan tiap kata pada nama project
		filenameStr = strings.Join(filenameArr, "-")   // menggabungkan nama project yang tadi dipisah (spasi berganti menjadi tanda '-')

		// memberi nama pada file gambar
		filename := fmt.Sprintf("%s%s", filenameStr, filepath.Ext(handler.Filename))
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

		// membuat sebuah var untuk digunakan sebagai key pada context (untuk mengatasi warning should not use built-in type string as key)
		var UploadFileID models.ContextKey = "FileName"

		// membuat sebuah context baru dengan menyisipkan value di dalamnya, valuenya adalah filename dan keynya adalah "NameOfUploadedFile"
		ctx := context.WithValue(r.Context(), UploadFileID, filename)

		// mengirim nilai context ke object http.HandlerFunc yang menjadi parameter saat fungsi middleware ini dipanggil
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// membuat middleware untuk menghandle update/edit file
func UpdateFile(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// mengambil id project dari url parameter
		projectId, _ := strconv.Atoi(mux.Vars(r)["id"])

		// mengambil nama project dari inputan form
		projectName := r.FormValue("projectName")

		// menyiapkan variabel penampung nama file
		var filename string

		// mengambil file dari form
		uploadedFile, handler, err := r.FormFile("projectImg")
		if err != nil {
			// apabila err tidak bernilai nil (artinya file tidak berhasil diambil dari form alias tidak ada file yang diupload, maka set filename dengan value projectImg yang lama)
			filename = models.GetDataProject(projectId).Img // men-set filename sebagai nilai Img dari variabel/object project
		} else {
			// apabila err bernilai nil (artinya file berhasil diambil dari form, maka handle file yang diupload)

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

			// menyusun string untuk digunakan sebagai nama file gambar
			filenameStr := projectName                     // mengambil nama project
			filenameStr = strings.ToLower(filenameStr)     // mem-format nama project menjadi huruf kecil
			filenameArr := strings.Split(filenameStr, " ") // memisahkan tiap kata pada nama project
			filenameStr = strings.Join(filenameArr, "-")   // menggabungkan nama project yang tadi dipisah (spasi berganti menjadi tanda '-')

			// memberi nama pada file gambar
			filename = fmt.Sprintf("%s%s", filenameStr, filepath.Ext(handler.Filename))
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
		}

		// membuat sebuah var untuk digunakan sebagai key pada context (untuk mengatasi warning should not use built-in type string as key)
		var UpdateFileID models.ContextKey = "FileName"

		// membuat sebuah context baru dengan menyisipkan value di dalamnya, valuenya adalah filename dan keynya adalah "NameOfUploadedFile"
		ctx := context.WithValue(r.Context(), UpdateFileID, filename)

		// mengirim nilai context ke object http.HandlerFunc yang menjadi parameter saat fungsi middleware ini dipanggil
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
