package main

import (
	"bufio"
	"crypto/rand"
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/nbutton23/zxcvbn-go"
)

//go:embed all:frontend
var frontendFS embed.FS

// Meng-embed wordlist bahasa Indonesia
//
//go:embed wordlist-id.txt
var wordlistIdFile embed.FS

const (
	charsetLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charsetNumbers = "0123456789"
	charsetSymbols = "!@#$%^&*()_+-=[]{}|;':,.<>?"
)

// --- Fungsi untuk Menerjemahkan ---
func translateCrackTime(englishTime string) string {
	r := strings.NewReplacer(
		"centuries", "abadi", "century", "abadi",
		"years", "tahun", "year", "tahun",
		"months", "bulan", "month", "bulan",
		"days", "hari", "day", "hari",
		"hours", "jam", "hour", "jam",
		"minutes", "menit", "minute", "menit",
		"seconds", "detik", "second", "detik",
	)
	return r.Replace(englishTime)
}

func generateRandomPassword(length int, useNumbers bool, useSymbols bool) (string, error) {
	var charset = charsetLetters
	if useNumbers {
		charset += charsetNumbers
	}
	if useSymbols {
		charset += charsetSymbols
	}
	password := make([]byte, length)
	charsetLen := big.NewInt(int64(len(charset)))
	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			return "", err
		}
		password[i] = charset[n.Int64()]
	}
	return string(password), nil
}

// --- Revisi Fungsi Mnemonic ---
func generateIndoMnemonicPassword() (string, error) {
	file, err := wordlistIdFile.Open("wordlist-id.txt")
	if err != nil {
		return "", err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	if len(words) == 0 {
		return "", fmt.Errorf("wordlist kosong")
	}

	var result []string
	numWords := 2
	wordlistLen := big.NewInt(int64(len(words)))

	for i := 0; i < numWords; i++ {
		n, err := rand.Int(rand.Reader, wordlistLen)
		if err != nil {
			return "", err
		}
		result = append(result, words[n.Int64()])
	}

	// Tambahkan angka acak 2 digit (10-99)
	num, err := rand.Int(rand.Reader, big.NewInt(90)) // Angka dari 0-89
	if err != nil {
		return "", err
	}
	randomNumber := num.Int64() + 10 // Geser menjadi 10-99

	// Gabungkan menjadi format: kata-kata-angka
	return fmt.Sprintf("%s-%s-%d", result[0], result[1], randomNumber), nil
}

func handleGenerate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	mode := r.URL.Query().Get("mode")
	var password string
	var err error

	if mode == "mnemonic" {
		password, err = generateIndoMnemonicPassword()
	} else {
		length, _ := strconv.Atoi(r.URL.Query().Get("length"))
		if length < 8 || length > 128 {
			length = 16
		}
		useNumbers, _ := strconv.ParseBool(r.URL.Query().Get("numbers"))
		useSymbols, _ := strconv.ParseBool(r.URL.Query().Get("symbols"))
		password, err = generateRandomPassword(length, useNumbers, useSymbols)
	}

	if err != nil {
		http.Error(w, "Gagal membuat kata sandi", http.StatusInternalServerError)
		log.Println("Error generating password:", err)
		return
	}

	// Lakukan analisis seperti biasa
	analysis := zxcvbn.PasswordStrength(password, nil)

	// Langsung terjemahkan field di dalam objek hasil analisis
	analysis.CrackTimeDisplay = translateCrackTime(analysis.CrackTimeDisplay)

	// Buat response JSON dengan langsung menyertakan seluruh objek analisis
	response := map[string]interface{}{
		"password": password,
		"analysis": analysis,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	contentRoot, err := fs.Sub(frontendFS, "frontend")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.FileServer(http.FS(contentRoot)))
	http.HandleFunc("/generate", handleGenerate)

	log.Println("Password Toolkit Cerdas (ID) berjalan di http://localhost:8888")
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}
}
