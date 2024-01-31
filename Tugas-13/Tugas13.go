package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type NilaiMahasiswa struct {
	Nama, MataKuliah, IndeksNilai string
	Nilai, ID                     uint
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}

func main() {
	http.HandleFunc("/nilai-mahasiswa", HandleNilaiMahasiswa)
	http.HandleFunc("/get-nilai-mahasiswa", HandleGetNilaiMahasiswa)

	fmt.Println("Server listening on :3301...")
	http.ListenAndServe(":3301", nil)
}

func HandleNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handlePostNilaiMahasiswa(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handlePostNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
	// Authentication login
	username, password, ok := r.BasicAuth()
	if !ok || !checkCredentials(username, password) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Menerima data dalam bentuk JSON atau formData
	var data NilaiMahasiswa
	if strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}
	} else {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}
		data.Nama = r.FormValue("Nama")
		data.MataKuliah = r.FormValue("MataKuliah")
		data.Nilai = parseUint(r.FormValue("Nilai"))
	}

	// Mengolah ID dan IndeksNilai
	data.ID = uint(len(nilaiNilaiMahasiswa) + 1)
	data.IndeksNilai = calculateIndeksNilai(data.Nilai)

	// Menambahkan data ke nilaiNilaiMahasiswa
	nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, data)

	w.WriteHeader(http.StatusCreated)
}

func HandleGetNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Basic Authentication
	username, password, ok := r.BasicAuth()
	if !ok || !checkCredentials(username, password) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Menampilkan semua data NilaiMahasiswa
	json.NewEncoder(w).Encode(nilaiNilaiMahasiswa)
}

func checkCredentials(username, password string) bool {
	return username == "user" && password == "password"
}

func parseUint(s string) uint {
	value, err := strconv.ParseUint(s, 10, 0)
	if err != nil {
		return 0
	}
	return uint(value)
}

func calculateIndeksNilai(nilai uint) string {
	switch {
	case nilai >= 80:
		return "A"
	case nilai >= 70:
		return "B"
	case nilai >= 60:
		return "C"
	case nilai >= 50:
		return "D"
	default:
		return "E"
	}
}
