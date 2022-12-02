-- membuat tabel tb_projects
CREATE TABLE IF NOT EXISTS tb_projects (
	id serial NOT NULL PRIMARY KEY,
	name varchar NOT NULL,
	start_date date NOT NULL,
	end_date date NOT NULL,
	description text NOT NULL,
	technologies varchar[] NOT NULL,
	image varchar NOT NULL
);

-- menghapus semua data di tb_projects
DELETE FROM tb_projects;

-- insert data
INSERT INTO tb_projects (name, start_date, end_date, description, technologies, image) VALUES
('Mini Expense Tracker', '2022-01-14', '2022-03-14', 'Expense Tracker" merujuk ke sistem pencatatan pengeluaran, aplikasi ini akan menyimpan pemasukan dan pengeluaran user serta menampilkan total selisih antara keduanya.',
'{"html5", "css3", "js", "react"}', 'mini-expense-tracker.png'),
('CRUD Mahasiswa', '2022-02-14', '2022-03-9', 'Dalam programming, CRUD merupakan singkatan dari Create Read Update dan Delete. Yakni aplikasi yang berisi pengolahan data. Biasanya CRUD butuh database sebagai media penyimpanan. Akan tetapi untuk sementara ini app CRUD Mahasiswa lebih fokus ke kode React, CRUD ini hanya disimpan di memory saja.',
'{"html5", "js", "react"}', 'crud-mahasiswa.png'),
('Landing Page', '2021-10-14', '2022-03-14', 'Landing Page sederhana yang berisikan informasi perusahaan untuk menarik pelanggan. Untuk sementara ini data yang ada di dalam landing page masih bersifat dummy, karena proyek ini merupakan hasil belajar HTML CSS dan Bootstrap. Akan tetapi proyek ini dapat dijadikan template apabila ada proyek serupa kedepannya.',
'{"html5", "css3"}', 'landingpage.png'),
('YouNoob', '2021-02-14', '2022-03-14', 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Iure voluptatibus fugiat, veniam magnam eos pariatur earum illo odit, eius voluptas expedita. Ullam, repellendus inventore exercitationem perferendis beatae ea enim ad, nemo tempora ducimus sunt in, eaque illum eius quo necessitatibus non delectus nam? Inventore, aliquam.',
'{"html5", "css3", "js"}', 'younoob.png');