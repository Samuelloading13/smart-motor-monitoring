# Smart Motor Monitoring System

Smart Motor Monitoring System adalah proyek untuk memantau kondisi motor secara real-time menggunakan **Golang sebagai backend API** dan **HTML/JavaScript sebagai dashboard web**. Sistem ini menerima data sensor (suhu, tegangan, RPM, dan kecepatan), menampilkan status motor, serta menyimpan riwayat pembacaan sensor.

---

## ✨ Fitur

### 🔵 Backend API (Golang)

* **POST /api/sensor** — menerima data sensor baru
* **GET /api/sensor** — mengambil data sensor terbaru
* **GET /api/sensor/status** — menghitung kondisi motor
* **GET /api/sensor/history** — menampilkan riwayat pembacaan

### 🟢 Dashboard Web

* Menampilkan data sensor terbaru
* Status motor (*Baik*, *Waspada*, *Bahaya*)
* Tabel riwayat pembacaan sensor
* Update otomatis setiap 2 detik
* Menggunakan HTML + CSS + JavaScript (tanpa framework)

### 🟣 Functional Programming (FP)

* Update data menggunakan pure function
* Tanpa mengubah state secara langsung
* Data baru selalu dikembalikan melalui return value

---

## 📂 Struktur Proyek

```
/
├── main.go
├── index.html
└── README.md
```

---

## 🚀 Cara Menjalankan

### 1. Clone Repository

```sh
git clone https://github.com/<username>/<repo-name>.git
cd <repo-name>
```

### 2. Jalankan Backend

```sh
go run main.go
```

Server berjalan di:

```
http://localhost:8080
```

### 3. Buka Dashboard

Akses:

```
http://localhost:8080/index.html
```

---

## 📡 Contoh Format Request/Response

### POST `/api/sensor`

Body:

```json
{
  "suhu": 36.5,
  "tegangan": 12.3,
  "rpm": 4200,
  "kecepatan": 60
}
```

Response:

```json
{ "status": "Data sensor diterima" }
```

---

### GET `/api/sensor`

```json
{
  "suhu": 36.5,
  "tegangan": 12.3,
  "rpm": 4200,
  "kecepatan": 60
}
```

---

### GET `/api/sensor/status`

```json
{
  "status": "Motor dalam kondisi baik"
}
```

---

### GET `/api/sensor/history`

```json
[
  { "suhu": 36.1, "tegangan": 12.2, "rpm": 4100, "kecepatan": 58 },
  { "suhu": 36.5, "tegangan": 12.3, "rpm": 4200, "kecepatan": 60 }
]
```

---

## 🛠 Teknologi

* Go (net/http)
* HTML
* CSS
* JavaScript

---

## 📌 Pengembangan Mendatang

* Menambah grafik realtime (Chart.js)
* Sistem login dan autentikasi
* Penyimpanan history dengan database (PostgreSQL/SQLite)
* Integrasi dengan perangkat ESP32/Arduino

---
