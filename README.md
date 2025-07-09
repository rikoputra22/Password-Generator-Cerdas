# 🚀 Password Toolkit Cerdas

Sebuah aplikasi web modern untuk membuat dan menganalisis kata sandi, dibangun dengan backend Go yang super cepat dan frontend yang interaktif. Ini bukan sekadar generator biasa, melainkan sebuah toolkit untuk memahami dan menciptakan kata sandi yang benar-benar aman.

!(Screenshot.png)
-----------------

## ✨ Fitur Utama

* **Dua Mode Generator:**
    * **Acak:** Membuat kata sandi yang kompleks dengan kombinasi huruf, angka, dan simbol.
    * **Mudah Diingat:** Membuat kata sandi aman menggunakan metode Diceware dengan kata-kata bahasa Indonesia yang digabung dengan angka.
* **Analisis Kekuatan Real-time:** Setiap kata sandi yang dibuat langsung dianalisis menggunakan algoritma **zxcvbn** untuk memberikan umpan balik yang konkret.
* **Visualisasi Cerdas:** Menampilkan skor kekuatan dan estimasi waktu yang dibutuhkan untuk meretas kata sandi (misal: "5 abad", "10 hari").
* **Tampilan Modern & Responsif:** Dibangun dengan **Tailwind CSS** dan **Alpine.js** untuk pengalaman pengguna yang bersih dan interaktif di semua perangkat.
* **Animasi Halus:** Menggunakan **AutoAnimate** untuk transisi antarmuka yang mulus dan memuaskan.

---

## 🛠️ Teknologi yang Digunakan

* **Backend:** Go (Golang)
* **Frontend:** HTML, Tailwind CSS, Alpine.js
* **Analisis:** zxcvbn-go
* **Animasi:** AutoAnimate
* **Ikon:** Lucide Icons

---

## ⚙️ Cara Menjalankan Secara Lokal

Untuk menjalankan proyek ini di komputer Anda, ikuti langkah-langkah berikut:

1.  **Prasyarat:**
    * Pastikan Anda sudah menginstal **Go (versi 1.18 atau lebih baru)**.

2.  **Kloning Repositori:**
    ```bash
    git clone [https://github.com/rikoputra22/Password-Generator-Cerdas.git](https://github.com/rikoputra22/Password-Generator-Cerdas.git)
    cd Password-Generator-Cerdas
    ```

3.  **Siapkan Wordlist:**
    * Buat sebuah file bernama `wordlist-id.txt` di dalam folder utama.
    * Isi file tersebut dengan daftar kata-kata bahasa Indonesia (satu kata per baris).

4.  **Instalasi Dependensi:**
    * Buka terminal di folder proyek dan jalankan perintah ini untuk mengunduh semua kebutuhan proyek secara otomatis.
    ```bash
    go mod tidy
    ```

5.  **Jalankan Server:**
    ```bash
    go run main.go
    ```

6.  Buka browser Anda dan kunjungi **`http://localhost:8888`**.

---

## 📄 Lisensi

Proyek ini dilisensikan di bawah Lisensi MIT.