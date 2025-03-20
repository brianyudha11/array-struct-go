package main

import (
	"array_struct1/model"
	"array_struct1/node"
	"fmt"
)

func main() {
	pegawai1 := node.Pegawai{
		ID:     1,
		Nama:   "John Doe",
		Alamat: node.Address{"Jalan Sudirman", "Jakarta", 123},
		NoTelp: "08123456789",
		Email:  "gmail.com",
	}

	pegawai2 := node.Pegawai{
		ID:     2,
		Nama:   "John Doe",
		Alamat: node.Address{"Jalan Sudirman", "Jakarta", 123},
		NoTelp: "08123456789",
		Email:  "yahoo.com",
	}

	model.CreatePegawai(pegawai1)
	model.CreatePegawai(pegawai2)
	listpegawai := model.ReadPegawai()
	
	pegawai3 := node.Pegawai {
		ID:     3,
		Nama:   "John Doe",
		Alamat: node.Address{"Jalan Ahmad Jais", "Surabaya", 123},
		NoTelp: "08123456789",
		Email:  "gmail.com",
	}
	
	fmt.Println(listpegawai)
	model.UpdatePegawai(pegawai3, 2)
	fmt.Println("After Update")
	for _, emp := range model.ReadPegawai() {
		fmt.Println(emp)
	}

	model.DeletePegawai(3)
	fmt.Println("After Delete")
	for _, emp := range model.ReadPegawai() {
		fmt.Println(emp)
	}
}
