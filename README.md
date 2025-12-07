# 🏍️ Smart Motor Monitor Pro

![Go Version](https://img.shields.io/badge/Go-1.20%2B-00ADD8?style=for-the-badge&logo=go)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)
![Status](https://img.shields.io/badge/Status-Active-brightgreen?style=for-the-badge)

**Smart Motor Monitor Pro** adalah sistem pemantauan kondisi kendaraan berbasis IoT (*Internet of Things*) yang dirancang dengan pendekatan **Functional Programming (FP)**.

Proyek ini mendemonstrasikan bagaimana membangun *backend* yang tangguh, modular, dan *stateless* menggunakan Golang, dipadukan dengan *dashboard real-time* yang responsif tanpa ketergantungan pada *framework* frontend yang berat.

---

## 🌟 Fitur Unggulan

### 🧠 Backend Cerdas (Golang)
* **Arsitektur Functional Programming**: Meminimalisir *side-effects* dan mutasi data yang tidak perlu.
* **Analisis Status Multi-Dimensi**: Mampu mendeteksi komplikasi masalah (misal: *"Mesin Overheat"* **DAN** *"Tegangan Kritis"* secara bersamaan).
* **Auto-Timestamp & ID Generation**: Otomatisasi metadata untuk setiap paket data yang masuk.
* **RESTful API Ringan**: endpoint JSON standar untuk integrasi mudah.

### 💻 Dashboard Interaktif
* **Real-time Monitoring**: Tampilan data diperbarui otomatis setiap 2 detik (Auto-fetch).
* **Simulator Sensor Bawaan**: Panel input untuk menguji sistem tanpa perlu perangkat keras sensor fisik.
* **Visualisasi Peringatan Dini**: Indikator warna berubah (Hijau/Merah) sesuai kondisi kesehatan mesin.
* **Audit Log**: Tabel riwayat data lengkap dengan fitur reset.

---

## 🛠️ Teknologi & Konsep

| Komponen | Teknologi | Deskripsi |
| :--- | :--- | :--- |
| **Language** | ![Go](https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white) | Bahasa utama backend server. |
| **Paradigm** | **Functional Programming** | Pendekatan logika utama (Pure Functions, HOF). |
| **Frontend** | HTML5, CSS3, ES6+ | Dashboard UI tanpa framework (Vanilla JS). |
| **Protocol** | HTTP / REST / JSON | Komunikasi data antara Client dan Server. |

---

## 📐 Implementasi Functional Programming

Sistem ini menerapkan prinsip FP secara ketat pada `main.go` untuk logika bisnis yang bersih:

1.  **Pure Functions (`updateSensorData`)**
    * Memastikan **Immutability**. Data lama tidak diubah di tempat (*in-place*), melainkan digantikan oleh struktur data baru.
    
2.  **First-Class Functions (`checkOverheat`, dll)**
    * Aturan status (seperti batas suhu) disimpan dalam variabel fungsi. Ini membuat logika sangat modular dan mudah diganti.

3.  **Higher-Order Functions (`resolveStatus`)**
    * Fungsi cerdas yang menerima **kumpulan fungsi lain** sebagai argumen untuk menentukan status akhir secara dinamis.

---

## ⚙️ Logika Deteksi Masalah

Sistem diprogram dengan ambang batas (*threshold*) keamanan sebagai berikut:

| Parameter | Kondisi Aman | Kondisi Bahaya | Output Status |
| :--- | :--- | :--- | :--- |
| **🌡️ Suhu** | ≤ 80°C | **> 80°C** | `Mesin Overheat` |
| **⚡ Tegangan** | ≥ 11.0 V | **< 11.0 V** | `Tegangan Kritis` |
| **🚀 Kecepatan** | ≤ 100 km/h | **> 100 km/h** | `Kecepatan Berbahaya` |

> 💡 **Info:** Jika lebih dari satu kondisi bahaya terpenuhi, sistem akan menggabungkan statusnya (Contoh: `"Mesin Overheat + Tegangan Kritis"`).

---

## 🚀 Cara Menjalankan

### 1. Prasyarat
Pastikan Anda telah menginstal [Go (Golang)](https://go.dev/dl/).

### 2. Clone & Run
Buka terminal pada direktori proyek:

```bash
# Jalankan server
go run main.go
````

Jika berhasil, terminal akan menampilkan:

```
Server running at http://localhost:8080
```

### 3\. Buka Dashboard

Buka browser favorit Anda dan akses:
👉 **http://localhost:8080/dashboard.html**

-----

## 📡 Dokumentasi API

Gunakan endpoint ini untuk mengintegrasikan alat lain (seperti Postman atau ESP32).

### 1\. 📤 Kirim Data Sensor (POST)

  * **Endpoint:** `/api/sensor`
  * **Body (JSON):**
    ```json
    {
      "suhu": 90.5,
      "tegangan": 10.5,
      "rpm": 4500,
      "kecepatan": 60
    }
    ```
  * **Respon Sukses:**
    ```json
    { "msg": "Data updated", "time": "14:30:05" }
    ```

### 2\. 📥 Ambil Data Terkini (GET)

  * **Endpoint:** `/api/sensor`
  * **Respon:**
    ```json
    {
      "id": 12,
      "waktu": "14:30:05",
      "suhu": 90.5,
      "tegangan": 10.5,
      "rpm": 4500,
      "kecepatan": 60
    }
    ```

### 3\. ⚠️ Cek Status Diagnosa (GET)

  * **Endpoint:** `/api/sensor/status`
  * **Respon (Contoh Komplikasi):**
    ```json
    { "status": "Mesin Overheat + Tegangan Kritis" }
    ```

-----

## 📂 Struktur Direktori

```
/smart-motor-monitoring
├── 📄 main.go           # Jantung aplikasi (Server & FP Logic)
├── 📄 dashboard.html    # Antarmuka Pengguna (UI)
└── 📄 README.md         # Dokumentasi Proyek
```

-----

\<p align="center"\>
Dibuat dengan ❤️ menggunakan Golang & Functional Programming.
\</p\>

```
```
