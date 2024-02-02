package models

import (
	"Tugas-15/config"
	"Tugas-15/utils"
)

type Mahasiswa struct {
	ID   int
	Nama string
}

func GetAllMahasiswa() ([]*Mahasiswa, error) {
	rows, err := config.DB.Query("SELECT id, nama FROM mahasiswa")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mahasiswas []*Mahasiswa
	for rows.Next() {
		var m Mahasiswa
		err := rows.Scan(&m.ID, &m.Nama)
		if err != nil {
			return nil, err
		}
		mahasiswas = append(mahasiswas, &m)
	}

	return mahasiswas, nil
}

func (m *Mahasiswa) CreateMahasiswa() (*Mahasiswa, error) {
	// Validasi input sebelum menyimpan ke database
	if m.Nama == "" {
		return nil, utils.NewBadRequestError("Nama mahasiswa tidak boleh kosong")
	}

	// Query untuk menyimpan mahasiswa ke database
	result, err := config.DB.Exec("INSERT INTO mahasiswa (nama) VALUES (?)", m.Nama)
	if err != nil {
		return nil, err
	}

	// Ambil ID yang baru saja dibuat
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	m.ID = int(id)

	return m, nil
}

func GetMahasiswaByID(id int) (*Mahasiswa, error) {
	m := &Mahasiswa{}
	err := config.DB.QueryRow("SELECT id, nama FROM mahasiswa WHERE id=?", id).Scan(&m.ID, &m.Nama)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (m *Mahasiswa) UpdateMahasiswa() (*Mahasiswa, error) {
	// Validasi input sebelum melakukan update
	if m.Nama == "" {
		return nil, utils.NewBadRequestError("Nama mahasiswa tidak boleh kosong")
	}

	// Query untuk melakukan update pada mahasiswa di database
	_, err := config.DB.Exec("UPDATE mahasiswa SET nama=? WHERE id=?", m.Nama, m.ID)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (m *Mahasiswa) DeleteMahasiswa() error {
	// Query untuk menghapus mahasiswa dari database berdasarkan ID
	_, err := config.DB.Exec("DELETE FROM mahasiswa WHERE id=?", m.ID)
	if err != nil {
		return err
	}

	return nil
}


