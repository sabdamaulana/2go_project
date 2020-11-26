package models

import (
	"net/http"
	"project_2go/db"
)

type Mahasiswa struct {
	Id    int    `json:"id"`
	Nama  string `json:"nama"`
	Nim   string `json:"nim"`
	Email string `json:"email"`
}

func FetchAllMahasiswa() (Response, error) {
	var obj Mahasiswa
	var arrobj []Mahasiswa
	var res Response

	con := db.CreateCon()
	sqlStatement := "SELECT * FROM mahasiswa"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Nama, &obj.Nim, &obj.Email)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

func StoreMahasiswa(nama string, nim string, email string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT mahasiswa (nama, nim, email) VALUES (?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama, nim, email)
	if err != nil {
		return res, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Succes"
	res.Data = map[string]int64{
		"last_inserted_id" : lastInsertId,
	}

	return res, nil

}

func UpdateMahasiswa(id int, nama string, nim string, email string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE mahasiswa SET nama = ?, nim = ?, email = ? Where id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama, nim, email, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func DeleteMahasiswa(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM mahasiswa WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}
	return res, nil
}
