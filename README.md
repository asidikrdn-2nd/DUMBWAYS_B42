# Task 12 - Table Relation & File Upload (Final Project Chapter 2 Stage 1)

Setelah mempelajari terkait **relation** dan **file upload**, maka:

- Silakan kalian membuat **relasi** antara tabel **tb_users** ke **tb_projects** agar saat user tersebut login, maka pada page **home** akan menampilkan project yang user tersebut **buat**.

- Implementasikan **file upload** pada Task.

## Cara Menggunakan Aplikasi

Berikut cara menjalankan aplikasi ini di komputer lokal :

- Pastikan sudah menginstall `Go Compiler`, `PostgreeSQL` dan `Git SCM` di komputer anda
- Clone repository ini dengan menjalankan `git clone https://github.com/asidikrdn-2nd/dumbways-task_12` di terminal/cmd
- Buka aplikasi pgAdmin 4, lalu buat database baru di PostgreeSQL dengan nama `db_personal_web`
- Klik kanan pada database `db_personal_web`, pilih `Restore`, pada bagian `Filename` silahkan klik dan pilih file `dumbways-task_12/backup-db_personal_web.sql` untuk me-restore database.
- Cek database dan isinya dengan menjalankan perintah `SELECT * FROM tb_projects` dan `SELECT * FROM tb_users` pada query tool di pgAdmin 4, jika muncul dan tidak error, bisa lanjut ke tahap berikutnya
- Masuk ke folder `dumbways-task_12`
- Buat file `.env` lalu isikan `DATABASE_URL="postgres://user:password@localhost/db_personal_web"`. Pada bagian `user` dan `password` silahkan disesuaikan dengan user dan password pada PostgreeSQL masing-masing
- Selanjutnya buka terminal pada folder `dumbways-task_12` dan jalankan perintah `go run main.go`

## Fitur Yang Tersedia

Berikut daftar fitur yang bisa digunakan pada aplikasi ini :

- Fitur `Register` untuk mendaftarkan pengguna baru
- Fitur `Login` untuk masuk sebagai pengguna
- Akses halaman `Home`, `Add Project`, dan `Contact Me`\
  **Halaman `AddProject` hanya dapat diakses jika masuk sebagai pengguna*
- Tampilan setiap halaman sudah responsive, bisa digunakan di smartphone\
  **jika ada bug pada tampilan, silahkan beri tahu ya*
- Pada halaman `Home`, pengguna dapat melihat daftar project yang tersimpan di database
- Pada halaman `Home`, pengguna dapat mengakses halaman `Project Detail` dengan meng-klik judul atau gambar salah satu project yang ditampilkan di section My Project
- Pada halaman `Home`, pengguna dapat menghapus salah satu project dengan meng-klik tombol `Delete` pada project tersebut\
  **Tombol delete hanya muncul jika masuk sebagai pengguna yang sesuai dengan yang menambahkan project tersebut*
- Pada halaman `Home`, pengguna dapat edit/update detail salah satu project dengan meng-klik tombol `Edit` pada project tersebut. Penggguna akan diarahkan ke form update project yang sudah berisi detail project tadi, silahkan ubah bagian yang ingin diupdate. Semua inputan pada form tersebut harus diisi, kecuali pada bagian upload image, pengguna tidak perlu mengupload image apabila tidak ingin mengganti gambar dari project tersebut. Preview gambar dari project tersebut juga ditampilkan untuk memudahkan pengguna\
  **Tombol edit hanya muncul jika masuk sebagai pengguna yang sesuai dengan yang menambahkan project tersebut*
- Apabila pengguna mengunjungi halaman `Contact Me`, pengguna dapat mengirim email ke saya dengan subject dan body yang sesuai dengan form yang pengguna isi
- Apabila pengguna mengunjungi halaman `Add Project`, pengguna dapat menambahkan data project baru dengan mengisi form yang tersedia lalu men-submitnya. Detail data project yang baru ditambahkan akan otomatis terupdate di daftar My Project\
  **Halaman `AddProject` hanya dapat diakses jika masuk sebagai pengguna*
- Beberapa fitur lainnya masih dalam tahap pengerjaan, stay tune hehe

## Daftar User dan Project Default

Berikut adalah daftar user dan project yang sudah tersedia di dalam aplikasi dan siap digunakan :

### USER ACCOUNT

User 1

```Text
id    : 1
nama  : Admin
email : admin@mail.com
pass  : 123
```

User 2

```Text
id    : 2
nama  : User
email : user@mail.com
pass  : 456
```

### PROJECT DATA

Project 1

```Text
name    : Mini Expense Tracker
id_user : 1
```

Project 2

```Text
name    : CRUD Mahasiswa
id_user : 2
```

Project 3

```Text
name    : Landing Page
id_user : 1
```

Project 4

```Text
name    : YouNoob
id_user : 2
```
