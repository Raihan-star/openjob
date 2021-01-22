package models

import (
	"net/http"
	"openjob/db"
)

type Openjob struct {
	Code       string `json:"code"`
	Nama       string `json:"nama"`
	Keterangan string `json:"keterangan"`
	Status     string `json:"status"`
	Tanggal    string `json:"tanggal"`
}

func FetchAllOpenjob() (Response, error) {
	var obj Openjob
	var arrobj []Openjob
	var res Response

	con := db.CreateCon()
	sqlStatement := "SELECT * FROM data"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Code, &obj.Nama, &obj.Keterangan, &obj.Status, &obj.Tanggal)
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

func StoreOpenjob(code string, nama string, keterangan string, status string, tanggal string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT data (code, nama, keterangan, status, tanggal) VALUES (?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(code, nama, keterangan, status, tanggal)
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
		"last_inserted_id": lastInsertId,
	}

	return res, nil

}

func UpdateOpenjob(code int, nama string, keterangan string, status string, tanggal string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE data SET nama = ?, keterangan = ?, status = ?, tanggal = ? Where code = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama, keterangan, status, tanggal, code)
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

func DeleteOpenjob(code int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM data WHERE code = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(code)
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
