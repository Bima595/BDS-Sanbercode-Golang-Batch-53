package main

import (
	"Tugas-15/config"
	"Tugas-15/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"fmt"
)

func main() {
	// Inisialisasi koneksi ke database
	config.InitDB()

	// Membuat instance Gin
	r := gin.Default()

	// CRUD untuk Mahasiswa
	r.GET("/mahasiswa", GetAllMahasiswa)
	r.POST("/mahasiswa", CreateMahasiswa)
	r.GET("/mahasiswa/:id", GetMahasiswaByID)
	r.PUT("/mahasiswa/:id", UpdateMahasiswa)
	r.DELETE("/mahasiswa/:id", DeleteMahasiswaHandler)

	// CRUD untuk Matakuliah
	r.GET("/nilai", GetAllNilai)
	r.POST("/matakuliah", CreateMatakuliah)
	r.GET("/matakuliah/:id", GetMatakuliahByID)
	r.PUT("/matakuliah/:id", UpdateMatakuliah)
	r.DELETE("/matakuliah/:id", DeleteMatakuliahHandler)

	// CRUD untuk Nilai
	r.GET("/matakuliah", GetAllMatakuliah)
	r.POST("/nilai", CreateNilai)
	r.GET("/nilai/:id", GetNilaiByID)
	r.PUT("/nilai/:id", UpdateNilai)
	r.DELETE("/nilai/:id", DeleteNilaiHandler)

	// Menjalankan server
	fmt.Println("Server Running at Port 3301")
	r.Run(":3301")
}

func GetAllMahasiswa(c *gin.Context) {
	mahasiswas, err := models.GetAllMahasiswa()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mahasiswas)
}
// CRUD untuk Mahasiswa
func CreateMahasiswa(c *gin.Context) {
	var input models.Mahasiswa
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdMahasiswa, err := input.CreateMahasiswa()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdMahasiswa)
}

func GetMahasiswaByID(c *gin.Context) {
	mahasiswaID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mahasiswa ID"})
		return
	}

	mahasiswa, err := models.GetMahasiswaByID(mahasiswaID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mahasiswa not found"})
		return
	}

	c.JSON(http.StatusOK, mahasiswa)
}

func UpdateMahasiswa(c *gin.Context) {
	mahasiswaID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mahasiswa ID"})
		return
	}

	var updatedMahasiswa models.Mahasiswa
	if err := c.ShouldBindJSON(&updatedMahasiswa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedMahasiswa.ID = mahasiswaID
	mahasiswa, err := updatedMahasiswa.UpdateMahasiswa()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mahasiswa)
}

func DeleteMahasiswaHandler(c *gin.Context) {
	mahasiswaID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mahasiswa ID"})
		return
	}

	m := &models.Mahasiswa{ID: mahasiswaID}

	err = m.DeleteMahasiswa()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Mahasiswa deleted"})
}

// CRUD untuk Matakuliah
func GetAllMatakuliah(c *gin.Context) {
	matakuliahs, err := models.GetAllMatakuliah()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, matakuliahs)
}

func CreateMatakuliah(c *gin.Context) {
	var input models.Matakuliah
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdMatakuliah, err := input.CreateMatakuliah()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdMatakuliah)
}

func GetMatakuliahByID(c *gin.Context) {
	matakuliahID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid matakuliah ID"})
		return
	}

	matakuliah, err := models.GetMatakuliahByID(matakuliahID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Matakuliah not found"})
		return
	}

	c.JSON(http.StatusOK, matakuliah)
}

func UpdateMatakuliah(c *gin.Context) {
	matakuliahID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid matakuliah ID"})
		return
	}

	var updatedMatakuliah models.Matakuliah
	if err := c.ShouldBindJSON(&updatedMatakuliah); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedMatakuliah.ID = matakuliahID
	matakuliah, err := updatedMatakuliah.UpdateMatakuliah()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, matakuliah)
}

func DeleteMatakuliahHandler(c *gin.Context) {
	matakuliahID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid matakuliah ID"})
		return
	}

	// Create an instance of Matakuliah
	mk := &models.Matakuliah{ID: matakuliahID}

	// Call the DeleteMatakuliah function
	err = mk.DeleteMatakuliah()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Matakuliah deleted"})
}

// CRUD untuk Nilai
func GetAllNilai(c *gin.Context) {
	nilais, err := models.GetAllNilai()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nilais)
}

func CreateNilai(c *gin.Context) {
	var input models.Nilai
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdNilai, err := input.CreateNilai()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdNilai)
}

func GetNilaiByID(c *gin.Context) {
	nilaiID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid nilai ID"})
		return
	}

	nilai, err := models.GetNilaiByID(nilaiID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Nilai not found"})
		return
	}

	c.JSON(http.StatusOK, nilai)
}

func UpdateNilai(c *gin.Context) {
	nilaiID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid nilai ID"})
		return
	}

	var updatedNilai models.Nilai
	if err := c.ShouldBindJSON(&updatedNilai); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedNilai.ID = nilaiID
	nilai, err := updatedNilai.UpdateNilai()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nilai)
}

func DeleteNilaiHandler(c *gin.Context) {
	nilaiID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid nilai ID"})
		return
	}

	// Create an instance of Nilai
	n := &models.Nilai{ID: nilaiID}

	// Call the DeleteNilai function
	err = n.DeleteNilai()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Nilai deleted"})
}

