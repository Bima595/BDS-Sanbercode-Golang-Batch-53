package main

import (
	"Tugas-14/Config"
	Mahasiswa "Tugas-14/Models"
	"Tugas-14/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	db, e := Config.MySQL()

	if e != nil {
		log.Fatal(e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	fmt.Println("Success")

	router := httprouter.New()

	router.GET("/mahasiswa", GetAllMahasiswa)
    router.POST("/mahasiswa", CreateMahasiswa)
    router.PUT("/mahasiswa/:id", UpdateMahasiswa)
    router.DELETE("/mahasiswa/:id", DeleteMahasiswa)

    fmt.Println("Server Running at Port 3301")
    log.Fatal(http.ListenAndServe(":3301", router))
}

// GetAllMahasiswa retrieves all Mahasiswa data
func GetAllMahasiswa(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Check if Mahasiswa table is empty
	if !Mahasiswa.HasData() {
		// Auto-insert sample data
		// Replace this with your actual data source
		newMahasiswa := Mahasiswa.Mahasiswa{
			ID:        1,
			Nama:      "John Doe",
			NIM:       "123456",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		Mahasiswa.InsertMahasiswa(newMahasiswa)
	}

	// Retrieve all Mahasiswa data
	mahasiswaList, err := Mahasiswa.GetAllMahasiswa()
	if err != nil {
		utils.ResponseJSON(w, "Error retrieving Mahasiswa data", http.StatusInternalServerError)
		return
	}

	utils.ResponseJSON(w, mahasiswaList, http.StatusOK)
}

// CreateMahasiswa adds a new Mahasiswa to the database
func CreateMahasiswa(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var newMahasiswa Mahasiswa.Mahasiswa
	err := json.NewDecoder(r.Body).Decode(&newMahasiswa)
	if err != nil {
		utils.ResponseJSON(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Simulate auto input logic based on Nilai
	// Replace this with your actual logic
	nilai := 85.0
	nilaiGrade := Mahasiswa.GetGrade(nilai)

	// Add the Mahasiswa to the database
	newMahasiswa.CreatedAt = time.Now()
	newMahasiswa.UpdatedAt = time.Now()
	err = Mahasiswa.InsertMahasiswa(newMahasiswa)
	if err != nil {
		utils.ResponseJSON(w, "Error creating Mahasiswa", http.StatusInternalServerError)
		return
	}

	utils.ResponseJSON(w, fmt.Sprintf("Mahasiswa added with grade: %s", nilaiGrade), http.StatusCreated)
}

// UpdateMahasiswa updates an existing Mahasiswa in the database
func UpdateMahasiswa(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    // Extract Mahasiswa ID from URL parameters
    mahasiswaID := ps.ByName("id")

    // Parse Mahasiswa ID to integer
    id, err := strconv.Atoi(mahasiswaID)
    if err != nil {
        utils.ResponseJSON(w, "Invalid Mahasiswa ID", http.StatusBadRequest)
        return
    }

    // Retrieve Mahasiswa from the database by ID
    existingMahasiswa, err := Mahasiswa.GetMahasiswaByID(id)
    if err != nil {
        utils.ResponseJSON(w, "Mahasiswa not found", http.StatusNotFound)
        return
    }

    // Decode the updated Mahasiswa data from the request body
    var updatedMahasiswa Mahasiswa.Mahasiswa
    err = json.NewDecoder(r.Body).Decode(&updatedMahasiswa)
    if err != nil {
        utils.ResponseJSON(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Update Mahasiswa fields
    existingMahasiswa.Nama = updatedMahasiswa.Nama
    existingMahasiswa.NIM = updatedMahasiswa.NIM
    // Update other fields as needed

    // Set the update time
    existingMahasiswa.UpdatedAt = time.Now()

    // Update the Mahasiswa in the database
    err = Mahasiswa.UpdateMahasiswa(id, existingMahasiswa)
    if err != nil {
        utils.ResponseJSON(w, "Error updating Mahasiswa", http.StatusInternalServerError)
        return
    }

    utils.ResponseJSON(w, "Mahasiswa updated successfully", http.StatusOK)
}

// DeleteMahasiswa deletes a Mahasiswa from the database
func DeleteMahasiswa(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    // Extract Mahasiswa ID from URL parameters
    mahasiswaID := ps.ByName("id")

    // Parse Mahasiswa ID to integer
    id, err := strconv.Atoi(mahasiswaID)
    if err != nil {
        utils.ResponseJSON(w, "Invalid Mahasiswa ID", http.StatusBadRequest)
        return
    }

    // Delete the Mahasiswa from the database
    err = Mahasiswa.DeleteMahasiswa(id)
    if err != nil {
        utils.ResponseJSON(w, "Error deleting Mahasiswa", http.StatusInternalServerError)
        return
    }

    utils.ResponseJSON(w, "Mahasiswa deleted successfully", http.StatusOK)
}