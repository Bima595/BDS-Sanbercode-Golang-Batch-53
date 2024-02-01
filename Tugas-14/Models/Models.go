// Mahasiswa.go
package Mahasiswa

import (
	"time"
	"errors"
)

// Mahasiswa model
type Mahasiswa struct {
	ID        int       `json:"id"`
	Nama      string    `json:"nama"`
	NIM       string    `json:"nim"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Nilai model
type Nilai struct {
	ID          int       `json:"id"`
	MahasiswaID int       `json:"mahasiswa_id"`
	MataKuliah  string    `json:"mata_kuliah"`
	Nilai       float64   `json:"nilai"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Database simulation (replace with your actual database operations)
var mahasiswaList []Mahasiswa

// HasData checks if Mahasiswa table has data
func HasData() bool {
	return len(mahasiswaList) > 0
}

// InsertMahasiswa inserts a new Mahasiswa to the database
func InsertMahasiswa(newMahasiswa Mahasiswa) error {
	// Simulate database insert operation
	newMahasiswa.ID = len(mahasiswaList) + 1
	newMahasiswa.CreatedAt = time.Now()
	newMahasiswa.UpdatedAt = time.Now()
	mahasiswaList = append(mahasiswaList, newMahasiswa)
	return nil
}

// GetAllMahasiswa retrieves all Mahasiswa data
func GetAllMahasiswa() ([]Mahasiswa, error) {
	// Simulate database retrieval operation
	return mahasiswaList, nil
}

func GetMahasiswaByID(id int) (*Mahasiswa, error) {
	for _, m := range mahasiswaList {
		if m.ID == id {
			return &m, nil
		}
	}
	return nil, errors.New("Mahasiswa not found")
}

func UpdateMahasiswa(id int, updatedMahasiswa *Mahasiswa) error {
    for i, m := range mahasiswaList {
        if m.ID == id {
            // Update the Mahasiswa data
            updatedMahasiswa.ID = m.ID
            updatedMahasiswa.CreatedAt = m.CreatedAt
            updatedMahasiswa.UpdatedAt = time.Now()
            mahasiswaList[i] = *updatedMahasiswa // Dereference the pointer
            return nil
        }
    }
    return errors.New("Mahasiswa not found")
}

// DeleteMahasiswa deletes a Mahasiswa from the database
func DeleteMahasiswa(id int) error {
	for i, m := range mahasiswaList {
		if m.ID == id {
			// Remove the Mahasiswa from the list
			mahasiswaList = append(mahasiswaList[:i], mahasiswaList[i+1:]...)
			return nil
		}
	}
	return errors.New("Mahasiswa not found")
}


// GetGrade returns the grade based on the nilai
func GetGrade(nilai float64) string {
	// Simulate grading logic (replace with your actual grading logic)
	switch {
	case nilai >= 80:
		return "A"
	case nilai >= 70:
		return "B"
	case nilai >= 60:
		return "C"
	default:
		return "F"
	}
}

// Additional functions can be added as needed for your application logic.
