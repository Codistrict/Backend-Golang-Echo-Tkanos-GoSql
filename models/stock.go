package models

import (
	"net/http"
	"project-1/db"
	str "project-1/struct"
)

func Input_Inventory(nama_barang string, jumlah_barang int, harga_barang int) (Response, error) {
	var res Response
	var invent str.Insert_Stock

	con := db.CreateCon()

	sqlStatement := "SELECT nama_barang FROM stock WHERE nama_barang=?"

	_ = con.QueryRow(sqlStatement, nama_barang).Scan(&invent.Nama_barang)

	if invent.Nama_barang == "" {

		sqlStatement := "INSERT INTO stock (nama_barang,jumlah_barang,harga_barang) values(?,?,?)"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(nama_barang, jumlah_barang, harga_barang)

		invent.Nama_barang = nama_barang
		invent.Jumlah_barang = jumlah_barang
		invent.Harga_barang = harga_barang

		stmt.Close()

		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = invent

	} else {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		invent.Nama_barang = ""
		invent.Harga_barang = 0
		invent.Jumlah_barang = 0
		res.Data = invent
	}

	return res, nil
}

func Read_Stock() (Response, error) {
	var res Response
	var arr_invent []str.Read_Stock
	var invent str.Read_Stock

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM stock"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&invent.Kode_stock, &invent.Nama_barang, &invent.Jumlah_barang, &invent.Harga_barang)
		if err != nil {
			return res, err
		}
		arr_invent = append(arr_invent, invent)
	}

	if arr_invent == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr_invent
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr_invent
	}

	return res, nil
}

func Update_Stock(kode_inventory string, nama_barang string, jumlah_barang int, harga_barang int) (Response, error) {
	var res Response
	con := db.CreateCon()

	sqlstatement := "UPDATE stock SET nama_barang=?,jumlah_barang=?,harga_barang=? WHERE kode_stock=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama_barang, jumlah_barang, harga_barang, kode_inventory)

	if err != nil {
		return res, err
	}

	rowschanged, err := result.RowsAffected()

	if err != nil {
		return res, err
	}

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Suksess"
	res.Data = map[string]int64{
		"rows": rowschanged,
	}

	return res, nil
}

func Check_Nama_Stcok(kode_inventory string, nama_barang string) (Response, error) {
	var res Response
	var check str.Check_Nama_Stock

	con := db.CreateCon()

	sqlstatement := "SELECT kode_stock FROM stock WHERE kode_stock!=? && nama_barang==?"

	_ = con.QueryRow(sqlstatement, kode_inventory, nama_barang).Scan(&check.Kode_inventory)

	if check.Kode_inventory == "" {
		check.Kode_inventory = kode_inventory
		check.Status = "true"
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = check
	} else {
		check.Status = "false"
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = check
	}

	return res, nil
}
