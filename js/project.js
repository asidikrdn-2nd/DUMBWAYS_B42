let projectList = [
  {
    id: "01",
    projectName: "Mini Expense Tracker",
    startDate: "01/14/2022",
    endDate: "03/14/2022",
    projectDesc: `"Expense Tracker" merujuk ke sistem pencatatan pengeluaran,
    aplikasi ini akan menyimpan pemasukan dan pengeluaran user serta
    menampilkan total selisih antara keduanya.`,
    tech: ["HTML 5", "CSS 3", "JavaScript", "React JS"],
    img: "img/mini-expense-tracker.png",
  },
  {
    id: "02",
    projectName: "CRUD Mahasiswa",
    startDate: "02/14/2022",
    endDate: "03/14/2022",
    projectDesc: `Dalam programming, CRUD merupakan singkatan dari Create Read
    Update dan Delete. Yakni aplikasi yang berisi pengolahan data.
    Biasanya CRUD butuh database sebagai media penyimpanan. Akan
    tetapi untuk sementara ini app CRUD Mahasiswa lebih fokus ke
    kode React, CRUD ini hanya disimpan di memory saja.`,
    tech: ["HTML 5", "JavaScript", "React JS"],
    img: "img/crud-mahasiswa.png",
  },
  {
    id: "03",
    projectName: "Landing Page",
    startDate: "10/14/2021",
    endDate: "03/14/2022",
    projectDesc: `Landing Page sederhana yang berisikan informasi perusahaan untuk
    menarik pelanggan. Untuk sementara ini data yang ada di dalam
    landing page masih bersifat dummy, karena proyek ini merupakan
    hasil belajar HTML CSS dan Bootstrap. Akan tetapi proyek ini
    dapat dijadikan template apabila ada proyek serupa kedepannya.`,
    tech: ["HTML 5", "CSS 3"],
    img: "img/landingpage.png",
  },
  {
    id: "04",
    projectName: "YouNoob",
    startDate: "12/14/2021",
    endDate: "03/14/2022",
    projectDesc: `Lorem ipsum dolor sit amet consectetur adipisicing elit. Iure
    voluptatibus fugiat, veniam magnam eos pariatur earum illo odit,
    eius voluptas expedita. Ullam, repellendus inventore
    exercitationem perferendis beatae ea enim ad, nemo tempora
    ducimus sunt in, eaque illum eius quo necessitatibus non
    delectus nam? Inventore, aliquam.`,
    tech: ["HTML 5", "CSS 3", "JavaScript"],
    img: "img/younoob.png",
  },
];

// membuat fungsi untuk hapus satu item list project
const handleDeleteClick = (e) => {
  // console.log(deleteButton);
  // console.log(e.target.id);
  projectList = projectList.filter((el) => {
    return el.id != e.target.id;
  });
  // console.log(projectList);

  // mengosongkan list project
  rowNode.innerHTML = "";
  // menampilkan kembali list project satu persatu setelah diupdate
  updateProjectListDOM();
};

// mengambil node baris sebagai parent dari list project
let rowNode = document.getElementById("projectListRow");
let deleteButton;

// fungsi untuk mengupdate list project sebagai child dari rowNode
const updateProjectListDOM = () => {
  // console.log(projectList);

  projectList.forEach((el) => {
    // Menghitung durasi berdasarkan tanggal mulai dan berakhir
    let startDate = new Date(el.startDate);
    let endDate = new Date(el.endDate);
    let durasi = endDate - startDate;
    let satuBulan = 1000 * 60 * 60 * 24 * 30;
    durasi = Math.round(durasi / satuBulan);
    durasi < 1 ? (durasi = "kurang dari 1") : (durasi = durasi);

    // menambahkan tiap card list project sebagai child element dari rowNode
    rowNode.innerHTML += `<div class="card">
    <div class="card-img">
      <a href="/project-detail.html?id=${el.id}">
      <img src="${el.img}" alt="#" />
      </a>
    </div>
    <div class="card-title">
    <a href="/project-detail.html?id=${el.id}">
      <h3>${el.projectName}</h3>
      <small>Durasi : ${durasi} Bulan</small>
      </a>
    </div>
    <div class="card-body">
      <p>
        ${el.projectDesc}
      </p>
    </div>
    <div class="tech-icon">
                ${
                  el.tech.includes("HTML 5")
                    ? '<img src="img/html5.png" alt="HTML 5" />'
                    : ""
                }
                ${
                  el.tech.includes("CSS 3")
                    ? '<img src="img/css3.png" alt="CSS 3" />'
                    : ""
                }
                ${
                  el.tech.includes("JavaScript")
                    ? '<img src="img/js.png" alt="JavaScript" />'
                    : ""
                }
                ${
                  el.tech.includes("React JS")
                    ? '<img src="img/react.png" alt="React JS" />'
                    : ""
                }
                
                
                
              </div>
              <div class="button-group">
                <button>edit</button>
                <button id=${el.id} class="delete-button">delete</button>
              </div>
  </div>`;
  });

  // mengambil seluruh node dengan class .delete-button
  deleteButton = document.querySelectorAll(".delete-button");

  // menambahkan fungsi handleDeleteClick saat event onclick pada masing-masing node deleteButton
  deleteButton.forEach((el) => {
    el.addEventListener("click", handleDeleteClick);
  });
};

// menampilkan list project saat web pertama di load
updateProjectListDOM();

// Mengambil form node
let addProjectForm = document.getElementById("projectForm");

// Menambahkan event ke form node untuk handle form submit
addProjectForm.addEventListener("submit", (e) => {
  e.preventDefault();

  // membuat object untuk project baru
  let newProject = {
    id: new Date().getTime().toString(),
    projectName: document.getElementById("projectName").value.trim(),
    startDate: document.getElementById("startDate").value.trim(),
    endDate: document.getElementById("endDate").value.trim(),
    projectDesc: document.getElementById("projectDesc").value.trim(),
    tech: [],
    img: "",
  };
  let error = "";

  // mengambil tiap checkbox node
  html = document.getElementById("html");
  css = document.getElementById("css");
  javascript = document.getElementById("javascript");
  reactjs = document.getElementById("reactjs");

  // periksa kondisi checkbox, apabila terceklis, maka push valuenya ke arrah tech
  html.checked && newProject.tech.push(html.value);
  css.checked && newProject.tech.push(css.value);
  javascript.checked && newProject.tech.push(javascript.value);
  reactjs.checked && newProject.tech.push(reactjs.value);

  // mengambil files dari node input file
  projectImg = document.getElementById("projectImg").files;
  // console.log(projectImg[0]);

  // Cek file upload apakah ada ? apakah formatnya sesuai (jpeg/png) ?
  if (projectImg.length > 0) {
    if (
      projectImg[0].type == "image/png" ||
      projectImg[0].type == "image/jpeg"
    ) {
      // jika semua syarat terpenuhi, buatlah urlnya lalu simpan di object dengan key img
      newProject.img = URL.createObjectURL(projectImg[0]);
    } else {
      newProject.img = "";
      error = "File yang anda upload harus berupa gambar jpg/png";
    }
  } else {
    newProject.img = "";
  }

  newProject.startDate = formatWaktu(newProject.startDate);
  newProject.endDate = formatWaktu(newProject.endDate);

  // cek inputan kosong
  for (data in newProject) {
    if (newProject[data] == "" && error == "") {
      error = "Inputan tidak boleh kosong";
    }
  }

  // jika tidak ada error, push data ke array
  error != "" ? alert(error) : projectList.push(newProject);
  // console.log(projectList);

  // mengosongkan list project
  rowNode.innerHTML = "";
  // menampilkan kembali list project satu persatu setelah diupdate
  updateProjectListDOM();

  // jika tidak ada error (form berhasil disubmit), reset form (agar kembali kosong)
  error == "" && addProjectForm.reset();
});

// membuat fungsi untuk menyesuaikan format waktu agar dapat digunakan pada pembuatan object Date
const formatWaktu = (time) => {
  // memecah string waktu menjadi array dengan 3 element [tahun, bulan, tanggal]
  time = time.split("-");
  // menyesuaikan format 'bulan/tanggal/tahun'
  time = `${time[1]}/${time[2]}/${time[0]}`;
  return time;
};
