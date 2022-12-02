// ====================
// Data Storage
// ====================
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
    startDate: "02/14/2021",
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
// ====================
// End of Data Storage
// ====================

// ====================
// Handle Add Project
// ====================

// Fungsi menghitung selisih waktu dengan 2 parameter
const selisihWaktu = (start, end) => {
  let startDate = new Date(start);
  let endDate = new Date(end);
  // menghitung selisih waktu dalam bentuk milisecond
  let durasi = endDate - startDate;
  // konversi milisecond ke hari
  let satuHari = 1000 * 60 * 60 * 24;
  let durasiHari = Math.floor(durasi / satuHari);

  // setiap 30 hari menjadi 1 bulan
  let durasiBulan = 0;
  while (durasiHari >= 30) {
    durasiBulan += 1;
    durasiHari -= 30;
  }

  // setiap 12 bulan menjadi 1 tahun
  let durasiTahun = 0;
  while (durasiBulan >= 12) {
    durasiTahun += 1;
    durasiBulan -= 12;
  }

  let output;
  if (durasiTahun != 0 && durasiBulan != 0) {
    output = `${durasiTahun} Tahun, ${durasiBulan} Bulan, ${durasiHari} Hari`;
  } else if (durasiBulan != 0) {
    output = `${durasiBulan} Bulan, ${durasiHari} Hari`;
  } else {
    output = `${durasiHari} Hari`;
  }

  return output;
};

// Mengambil form node
let addProjectForm = document.getElementById("projectForm");
// console.log(addProjectForm);

// Menambahkan event ke form node untuk handle form submit (apabila terdapat addProjectForm, karena addProjectForm hanya berada di page project, tidak ada di page project-detail)
addProjectForm != null &&
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
// ====================
// End of Handle Add Project
// ====================

// ====================
// Project List
// ====================
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

// var untuk menampung node delete button, disimpan diluar untuk menghindari re-deklarasi di perulangan
let deleteButton;

// fungsi untuk mengupdate list project sebagai child dari rowNode
const updateProjectListDOM = () => {
  // console.log(projectList);

  projectList.forEach((el) => {
    // Menghitung durasi berdasarkan tanggal mulai dan berakhir
    durasi = selisihWaktu(el.startDate, el.endDate);

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
      <small>Durasi : ${durasi}</small>
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

// menampilkan list project saat web pertama di load (apabila terdapat rowNode, karena rowNode hanya berada di page project, tidak ada di page project-detail)
rowNode != null && updateProjectListDOM();
// ====================
// End of Project List
// ====================

// ====================
// Project Detail
// ====================
let projectDisplay = document.getElementById("project-display");

// mengambil id dari query yang disisipkan di url
let idFromQuery = window.location.search;
idFromQuery = idFromQuery.split("=");
idFromQuery = idFromQuery[1];
// console.log(idFromQuery);

// membuat fungsi format tanggal
const formatTanggal = (date) => {
  let tanggal = new Date(date);

  let arrBulan = [
    "Jan",
    "Feb",
    "Mar",
    "Apr",
    "Mei",
    "Jun",
    "Jul",
    "Agus",
    "Sep",
    "Oct",
    "Nov",
    "Des",
  ];

  return `${tanggal.getDate()} ${
    arrBulan[tanggal.getMonth()]
  } ${tanggal.getFullYear()}`;
};

// mengambil data project yang memiliki id sesuai
let projectDisplayed = projectList.filter((element) => {
  return element.id == idFromQuery;
});
// console.log(projectDisplayed);

// jika terdapat element section#project-display
if (projectDisplay != null) {
  // dan terdapat data project yang diambil dari projectList
  if (projectDisplayed.length > 0) {
    // menghitung durasi project
    let durasi = selisihWaktu(
      projectDisplayed[0].startDate,
      projectDisplayed[0].endDate
    );

    let tech = "";

    // jalankan perulangan pada value dari property tech, lalu jika ada isi yang sesuai maka tambahkan sebuah kode html ke variabel tech
    projectDisplayed[0].tech.forEach((el) => {
      if (el == "HTML 5") {
        tech += `<div class="list-item">
          <img src="img/html5.png" alt="HTML 5" />
          <p>HTML 5</p>
        </div>`;
      } else if (el == "CSS 3") {
        tech += `<div class="list-item">
          <img src="img/css3.png" alt="CSS 3" />
          <p>CSS 3</p>
        </div>`;
      } else if (el == "JavaScript") {
        tech += `<div class="list-item">
          <img src="img/js.png" alt="JavaScript" />
          <p>JavaScript</p>
        </div>`;
      } else if (el == "React JS") {
        tech += `<div class="list-item">
          <img src="img/react.png" alt="React JS" />
          <p>React JS</p>
        </div>`;
      }
    });

    // menampilkan project detail sebagai child dari element section#project-display
    projectDisplay.innerHTML = `<a href="/project.html" style="margin-top: 10px">Kembali</a>
      <h1 class="project-title">${projectDisplayed[0].projectName}</h1>
      <div class="project-img">
        <img src=${projectDisplayed[0].img} alt="Nama Project" />
      </div>
      <div class="project-detail">
        <div class="duration">
          <h2 class="title">Duration</h2>
          <p>${formatTanggal(projectDisplayed[0].startDate)} - 
          ${formatTanggal(projectDisplayed[0].endDate)}</p>
          <p>${durasi}</p>
        </div>
        <div class="technologies">
          <h2 class="title">Technologies</h2>
          ${tech /** sisipkan kode html yang tadi disimpan di var tech */}
        </div>
      </div>
      <div class="project-desc">
        <p>
        ${projectDisplayed[0].projectDesc}
        </p>
      </div>`;
  }
  // jika terdapat element section#project-display namun data project tidak ditemukan, maka tampilkan halaman default
  else {
    projectDisplay.innerHTML = `<a href="/project.html" style="margin-top: 10px">Kembali</a>
    <h1 class="project-title">Dumbways Web App</h1>
    <div class="project-img">
      <img src="img/younoob.png" alt="Nama Project" />
    </div>
    <div class="project-detail">
      <div class="duration">
        <h2 class="title">Duration</h2>
        <p>12 Nov 2022 - 11 Des 2022</p>
        <p>1 month</p>
      </div>
      <div class="technologies">
        <h2 class="title">Technologies</h2>
        <div class="list-item">
          <img src="img/html5.png" alt="HTML 5" />
          <p>HTML 5</p>
        </div>
        <div class="list-item">
          <img src="img/css3.png" alt="CSS 3" />
          <p>CSS 3</p>
        </div>
        <div class="list-item">
          <img src="img/js.png" alt="JS" />
          <p>JavaScript</p>
        </div>
        <div class="list-item">
          <img src="img/react.png" alt="React JS" />
          <p>React JS</p>
        </div>
      </div>
    </div>
    <div class="project-desc">
      <p>
        Lorem ipsum dolor, sit amet consectetur adipisicing elit. Ratione
        fugit ullam culpa quos quibusdam, eos neque voluptatum doloremque
        accusantium fuga minus, aliquam sit, quidem nulla animi aspernatur
        quisquam at asperiores iste facere veritatis enim atque eligendi.
        Quidem officia a ipsam vero id ad. Praesentium quaerat, nam harum
        quibusdam debitis placeat voluptatum quos nisi eaque voluptas
        aliquam labore quod soluta ullam enim sed provident deserunt quam
        ducimus veritatis deleniti recusandae. Amet id autem iure doloribus,
        veritatis culpa officiis impedit tempore exercitationem officia.
        Odit, suscipit deserunt. Architecto amet fugiat illo. Cupiditate
        deserunt eligendi eum! Autem maiores cum repudiandae iste aliquam
        laborum voluptates!
      </p>
      <p>
        Lorem ipsum dolor sit amet consectetur adipisicing elit. At quo ex
        sapiente! Ut harum eaque aliquid quis necessitatibus voluptatum sed
        ipsam itaque nemo reprehenderit nisi dolorum, recusandae aliquam
        similique hic, numquam molestias, repellendus eligendi quos!
        Molestiae dolores quo sunt adipisci quam at animi cumque impedit,
        vero, porro eos laboriosam non?
      </p>
      <p>
        Lorem ipsum dolor sit amet consectetur adipisicing elit. Unde, ullam
        recusandae animi eaque repellendus modi perferendis corporis sint
        consequatur itaque quas obcaecati consequuntur nisi cupiditate a
        minima debitis dolores. Possimus, voluptate? Enim laboriosam cumque
        est quas minus quae velit facilis molestias sunt et a sint
        necessitatibus, cupiditate aliquid dignissimos dolore quidem
        doloribus quo. Iusto eveniet molestias autem atque pariatur odio
        beatae, nemo voluptas facere libero expedita consequatur dolores.
        Quidem ipsum provident laboriosam quia dolor, adipisci, ea fuga
        alias maiores optio corrupti dolores minus dicta non esse officia
        necessitatibus nulla. Repudiandae distinctio labore, voluptas cumque
        eligendi sapiente eos veritatis quisquam totam quas delectus eum
        repellat deleniti illo provident eius autem! Ipsam quisquam illo
        nihil, repellendus perspiciatis at cum vitae, obcaecati odio
        consequuntur blanditiis eos quia reiciendis sunt, expedita molestias
        officia nemo tempora tempore possimus exercitationem numquam? Omnis
        harum ipsam dolore minima officia delectus similique necessitatibus,
        rem, laudantium neque repellat tempore suscipit aliquam cupiditate?
        Eaque numquam sit consequuntur natus nihil, esse dolore magnam
        soluta ab aliquam inventore officiis, reprehenderit quae eius fuga
        excepturi facere minima est totam.
      </p>
    </div>`;
  }
}

// ====================
// End Project Detail
// ====================
