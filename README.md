# Task 9 - Data Modelling & Database

Setelah mempelajari dan mempraktikan terkait materi data modeling, maka tugas kalian adalah membuat database dan table beserta isi dari pada tablenya. Serta buatlah fungsi fetch data project pada page home.

## Cara Menggunakan Aplikasi

Berikut cara menjalankan aplikasi ini di komputer lokal :

- Pastikan sudah menginstall `Go Compiler`, `PostgreeSQL` dan `Git SCM` di komputer anda
- Clone repository ini dengan menjalankan `git clone https://github.com/asidikrdn-2nd/dumbways-task_9` di terminal/cmd
- Buka aplikasi PGAdmin4, lalu buat database baru di PostgreeSQL dengan nama `db_personal_web`
- Masuk ke database `db_personal_web`, buka query tool, kemudian jalankan query yang ada pada file `dumbways-task_9/query.sql` di query tool
- Cek database dan isinya dengan menjalankan perintah `SELECT * FROM tb_projects`, jika muncul dan tidak error, bisa lanjut ke tahap berikutnya
- Masuk ke folder `dumbways-task_9`
- Buat file `.env` lalu isikan `DATABASE_URL="postgres://user:password@localhost/db_personal_web"`. Pada bagian `user` dan `password` silahkan disesuaikan dengan user dan password pada PostgreeSQL masing-masing
- Selanjutnya buka terminal pada folder `dumbways-task_9` dan jalankan perintah `go run main.go`

## Fitur Yang Tersedia

Berikut daftar fitur yang bisa digunakan pada aplikasi ini :

- Akses halaman `Home`, `Add Project`, dan `Contact Me`
- Tampilan setiap halaman sudah responsive, bisa digunakan di smartphone (*jika ada bug pada tampilan, silahkan beri tahu ya*)
- Pada halaman `Home`, pengguna dapat melihat daftar project yang tersimpan di database
- Pada halaman `Home`, pengguna dapat mengakses halaman `Project Detail` dengan meng-klik judul atau gambar salah satu project yang ditampilkan di section My Project
<!-- - Pada halaman `Home`, pengguna dapat menghapus salah satu project dengan meng-klik tombol `Delete` pada project tersebut -->
<!-- - Pada halaman `Home`, pengguna dapat edit/update detail salah satu project dengan meng-klik tombol `Edit` pada project tersebut. Penggguna akan diarahkan ke form update project yang sudah berisi detail project tadi, silahkan ubah bagian yang ingin diupdate. Semua inputan pada form tersebut harus diisi, kecuali pada bagian upload image, pengguna tidak perlu mengupload image apabila tidak ingin mengganti gambar dari project tersebut. Preview gambar dari project tersebut juga ditampilkan untuk memudahkan pengguna. -->
- Apabila pengguna mengunjungi halaman `Contact Me`, pengguna dapat mengirim email ke saya dengan subject dan body yang sesuai dengan form yang pengguna isi
<!-- - Apabila pengguna mengunjungi halaman `Add Project`, pengguna dapat menambahkan data project baru dengan mengisi form yang tersedia lalu men-submitnya. Detail data project yang baru ditambahkan akan otomatis terupdate di daftar My Project, selain itu, detail project yang ditambahkan juga akan muncul di console/terminal saat pengguna mensubmit form tersebut. -->
- Beberapa fitur lainnya masih dalam tahap pengerjaan, stay tune hehe
