package models

import (
	"Tugas-15/config"
	"Tugas-15/utils"
)

type Matakuliah struct {
	ID   int
	Nama string
}

func GetAllMatakuliah() ([]*Matakuliah, error) {
	rows, err := config.DB.Query("SELECT id, nama FROM mata_kuliah")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var matakuliahs []*Matakuliah
	for rows.Next() {
		var mk Matakuliah
		err := rows.Scan(&mk.ID, &mk.Nama)
		if err != nil {
			return nil, err
		}
		matakuliahs = append(matakuliahs, &mk)
	}

	return matakuliahs, nil
}

func (mk *Matakuliah) CreateMatakuliah() (*Matakuliah, error) {
	if mk.Nama == "" {
		return nil, utils.NewBadRequestError("Nama matakuliah tidak boleh kosong")
	}

	result, err := config.DB.Exec("INSERT INTO mata_kuliah (nama) VALUES (?)", mk.Nama)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	mk.ID = int(id)

	return mk, nil
}

func GetMatakuliahByID(id int) (*Matakuliah, error) {
	mk := &Matakuliah{}
	err := config.DB.QueryRow("SELECT id, nama FROM mata_kuliah WHERE id=?", id).Scan(&mk.ID, &mk.Nama)
	if err != nil {
		return nil, err
	}
	return mk, nil
}

func (mk *Matakuliah) UpdateMatakuliah() (*Matakuliah, error) {
	// Validasi input sebelum melakukan update
	if mk.Nama == "" {
		return nil, utils.NewBadRequestError("Nama matakuliah tidak boleh kosong")
	}

	// Query untuk melakukan update pada matakuliah di database
	_, err := config.DB.Exec("UPDATE matakuliah SET nama=? WHERE id=?", mk.Nama, mk.ID)
	if err != nil {
		return nil, err
	}

	return mk, nil
}

func (mk *Matakuliah) DeleteMatakuliah() error {
	// Query untuk menghapus matakuliah dari database berdasarkan ID
	_, err := config.DB.Exec("DELETE FROM mata_kuliah WHERE id=?", mk.ID)
	if err != nil {
		return err
	}

	return nil
}
