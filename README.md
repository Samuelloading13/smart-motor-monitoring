📡 Smart Motor Monitoring System

Sistem ini merupakan aplikasi pemantauan motor pintar (Smart Motor) yang dibangun menggunakan Golang (Backend API) dan HTML/JS (Dashboard Web).
Proyek ini menampilkan data sensor seperti suhu, tegangan, rpm, dan kecepatan dalam bentuk dashboard yang diperbarui secara real-time.

🚀 Fitur Utama
🟦 API Backend (Golang)

Endpoint POST /api/sensor
Mengirimkan data sensor baru.

Endpoint GET /api/sensor
Mengambil data sensor terbaru.

Endpoint GET /api/sensor/status
Menentukan status motor berdasarkan nilai sensor.

Endpoint GET /api/sensor/history
Menampilkan riwayat pembacaan sensor.

🟩 Dashboard Web

Menampilkan data suhu, tegangan, rpm, dan kecepatan.

Menampilkan status motor (baik, waspada, atau bahaya).

Menampilkan history sensor dalam bentuk tabel.

Update otomatis (auto refresh setiap 2 detik).

Menggunakan HTML, CSS, dan JavaScript murni — tanpa framework.

🟧 Filosofi FP (Functional Programming)

Tidak ada data global yang dimodifikasi langsung.

Update sensor menggunakan pure function:

func updateSensorData(old SensorData, new SensorData) SensorData {
    return newData
}


State hanya berubah melalui return value, bukan pengubahan langsung.

📁 Struktur Folder
/
├── main.go               → Backend API
├── index.html            → Dashboard utama
├── /history/             → Penyimpanan history (opsional)
└── README.md

▶️ Cara Menjalankan
1. Clone repo
git clone https://github.com/<username>/smart-motor-monitoring.git
cd smart-motor-monitoring

2. Jalankan backend
go run main.go


Server berjalan di:

http://localhost:8080

3. Buka dashboard

Cukup buka:

http://localhost:8080/index.html

📦 Endpoint API
POST /api/sensor

Kirim data sensor:

{
  "suhu": 36.5,
  "tegangan": 12.3,
  "rpm": 4200,
  "kecepatan": 60
}

GET /api/sensor

Ambil data terbaru:

{
  "suhu": 36.5,
  "tegangan": 12.3,
  "rpm": 4200,
  "kecepatan": 60
}

GET /api/sensor/status

Contoh respons:

{
  "status": "Motor dalam kondisi baik"
}

GET /api/sensor/history

Contoh:

[
  { "suhu": 36.1, "tegangan": 12.2, "rpm": 4100, "kecepatan": 58 },
  { "suhu": 36.5, "tegangan": 12.3, "rpm": 4200, "kecepatan": 60 }
]

🛠 Teknologi yang Digunakan

Go (net/http)

HTML + CSS + JavaScript

Postman (testing API)

📌 Pengembangan Selanjutnya

Menambahkan grafik real-time (Chart.js)

Menambah autentikasi user

Menghubungkan dengan sensor nyata melalui ESP32/Arduino

Penyimpanan history menggunakan database (PostgreSQL / SQLite)
