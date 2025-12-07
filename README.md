# Smart Motor Monitoring System

**Smart Motor Monitor** adalah sistem pemantauan kendaraan *real-time* yang dibangun dengan pendekatan **Functional Programming (FP)**. Proyek ini mendemonstrasikan bagaimana Golang dapat digunakan untuk membangun *backend* yang tangguh dan modular tanpa menggunakan *framework* eksternal, dipadukan dengan *frontend* modern yang responsif.

---

## ✨ Fitur Utama

### 🧠 Backend (Golang & FP)
* **Arsitektur Functional Programming**: Menggunakan *Pure Functions*, *First-Class Functions*, dan *Higher-Order Functions* untuk logika bisnis.
* **Logika Status Komposit**: Mampu mendeteksi berbagai masalah sekaligus (contoh: *"Mesin Overheat + Tegangan Kritis"*) secara dinamis.
* **Auto-Timestamp & ID**: Setiap data sensor yang masuk otomatis diberi ID unik dan penanda waktu server.
* **RESTful API**: Mendukung operasi `GET`, `POST`, dan `DELETE`.

### 💻 Dashboard Web (HTML5 + JS)
* **Monitoring Real-time**: Data diperbarui otomatis setiap 2 detik.
* **Panel Simulasi Sensor**: Form input bawaan untuk menguji sistem tanpa perangkat keras fisik.
* **Visualisasi Status**: Indikator warna otomatis (Hijau/Merah) berdasarkan kondisi mesin.
* **Manajemen Log**: Tabel riwayat lengkap dengan fitur *Reset Log*.

---

## 🛠️ Implementasi Functional Programming

Sistem ini bukan hanya sekadar CRUD biasa, tetapi menerapkan konsep FP secara ketat pada `main.go`:

1.  **Pure Functions**: Fungsi `updateSensorData` menjamin *immutability* data (data lama tidak diubah, tapi digantikan data baru).
2.  **First-Class Functions**: Aturan logika status (`checkOverheat`, `checkLowVoltage`) disimpan sebagai variabel fungsi yang independen.
3.  **Higher-Order Functions**: Fungsi `resolveStatus` menerima kumpulan fungsi aturan lain sebagai argumen untuk menentukan status akhir secara fleksibel.

---

## 🚀 Cara Menjalankan

### 1. Persiapan
Pastikan [Go (Golang)](https://go.dev/dl/) sudah terinstall di komputer Anda.

### 2. Jalankan Server
Buka terminal di folder proyek, lalu ketik:

go run main.go
Output terminal akan terlihat seperti ini:

Server running at http://localhost:8080
3. Akses Dashboard
Buka browser dan kunjungi: 👉 http://localhost:8080/dashboard.html

📡 Dokumentasi API
Berikut adalah daftar endpoint yang tersedia untuk integrasi:

1. Kirim Data Sensor Baru
URL: /api/sensor

Method: POST

Body (JSON):

JSON

{
  "suhu": 85.5,
  "tegangan": 12.1,
  "rpm": 4000,
  "kecepatan": 110
}
Response: { "msg": "Data updated", "time": "10:30:15" }

2. Ambil Data Terkini
URL: /api/sensor

Method: GET

Response:

JSON

{
  "id": 5,
  "waktu": "10:30:15",
  "suhu": 85.5,
  "tegangan": 12.1,
  "rpm": 4000,
  "kecepatan": 110
}
3. Cek Status Kesehatan Mesin
URL: /api/sensor/status

Method: GET

Response (Contoh jika ada masalah ganda):

JSON

{
  "status": "Mesin Overheat + Kecepatan Berbahaya"
}
4. Kelola Riwayat (History)
Ambil Data: GET /api/sensor/history

Hapus Data: DELETE /api/sensor/history

📂 Struktur Proyek
/
├── main.go           # Logic Backend (Server, FP Logic, API)
├── dashboard.html    # Frontend (UI, Fetch API, CSS)
└── README.md         # Dokumentasi Proyek
