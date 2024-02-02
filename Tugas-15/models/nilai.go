package models

import (
	"Tugas-15/utils"
	"Tugas-15/config"
)

type Nilai struct {
	ID           int
	Indeks       string
	Skor         int
	MahasiswaID  int
	MatakuliahID int
}

func (n *Nilai) CreateNilai() (*Nilai, error) {
	// Validasi input sebelum menyimpan ke database
	if n.Skor < 0 || n.Skor > 100 {
		return nil, utils.NewBadRequestError("Skor harus berada di antara 0 dan 100")
	}

	// Hitung Indeks berdasarkan Skor
	if n.Skor >= 90 {
		n.Indeks = "A"
	} else if n.Skor >= 80 {
		n.Indeks = "B"
	} else if n.Skor >= 70 {
		n.Indeks = "C"
	} else if n.Skor >= 60 {
		n.Indeks = "D"
	} else {
		n.Indeks = "E"
	}

	// Query untuk menyimpan nilai ke database
	result, err := config.DB.Exec("INSERT INTO nilai (indeks, skor, mahasiswa_id, mata_kuliah_id) VALUES (?, ?, ?, ?)",
		n.Indeks, n.Skor, n.MahasiswaID, n.MatakuliahID)
	if err != nil {
		return nil, err
	}

	// Ambil ID yang baru saja dibuat
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	n.ID = int(id)

	return n, nil
}

func GetAllNilai() ([]*Nilai, error) {
	rows, err := config.DB.Query("SELECT id, indeks, skor, mahasiswa_id, mata_kuliah_id FROM nilai")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nilais []*Nilai
	for rows.Next() {
		var n Nilai
		err := rows.Scan(&n.ID, &n.Indeks, &n.Skor, &n.MahasiswaID, &n.MatakuliahID)
		if err != nil {
			return nil, err
		}
		nilais = append(nilais, &n)
	}

	return nilais, nil
}

// func (n *Nilai) CreateNilai() (*Nilai, error) {
// 	// Validasi input sebelum menyimpan ke database
// 	if n.Indeks == "" || n.Skor < 0 || n.Skor > 100 {
// 		return nil, utils.NewBadRequestError("Indeks dan skor harus diisi dengan benar")
// 	}

// 	// Query untuk menyimpan nilai ke database
// 	result, err := config.DB.Exec("INSERT INTO nilai (indeks, skor, mahasiswa_id, mata_kuliah_id) VALUES (?, ?, ?, ?)",
// 		n.Indeks, n.Skor, n.MahasiswaID, n.MatakuliahID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Ambil ID yang baru saja dibuat
// 	id, err := result.LastInsertId()
// 	if err != nil {
// 		return nil, err
// 	}
// 	n.ID = int(id)

// 	return n, nil
// }

func GetNilaiByID(id int) (*Nilai, error) {
	n := &Nilai{}
	err := config.DB.QueryRow("SELECT id, indeks, skor, mahasiswa_id, mata_kuliah_id FROM nilai WHERE id=?", id).
		Scan(&n.ID, &n.Indeks, &n.Skor, &n.MahasiswaID, &n.MatakuliahID)
	if err != nil {
		return nil, err
	}
	return n, nil
}

func (n *Nilai) UpdateNilai() (*Nilai, error) {
	// Validasi input sebelum melakukan update
	if n.Indeks == "" || n.Skor < 0 || n.Skor > 100 {
		return nil, utils.NewBadRequestError("Indeks dan skor harus diisi dengan benar")
	}

	// Query untuk melakukan update pada nilai di database
	_, err := config.DB.Exec("UPDATE nilai SET indeks=?, skor=?, mahasiswa_id=?, mata_kuliah_id=? WHERE id=?",
		n.Indeks, n.Skor, n.MahasiswaID, n.MatakuliahID, n.ID)
	if err != nil {
		return nil, err
	}

	return n, nil
}

func (n *Nilai) DeleteNilai() error {
	// Query untuk menghapus nilai dari database berdasarkan ID
	_, err := config.DB.Exec("DELETE FROM nilai WHERE id=?", n.ID)
	if err != nil {
		return err
	}

	return nil
}
