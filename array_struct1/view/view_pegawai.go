package view

import (
	"array_struct1/model"
	"array_struct1/node"
	"bufio"
	"fmt"
	"os"
)

func Insert () {
	var kota, nama, notelp, email string
	var id, nomer int
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukan ID Pegawai : ")
	fmt.Scan(&id)

	fmt.Print("Masukan Nama Pegawai : ")
	nama, _ = reader.ReadString('\n')
	nama = nama[:len(nama)-1]

	fmt.Print("Masukan Kota : ")
	fmt.Scan(&kota)

	fmt.Print("Masukan Jalan : ")
	jalan, _ := reader.ReadString('\n')
	jalan = jalan[:len(jalan)-1]

	fmt.Print("Masukan Nomer Rumah : ")
	fmt.Scan(&nomer)

	fmt.Print("Masukan Nomer Telepon : ")
	fmt.Scan(&notelp)

	fmt.Print("Masukan Email : ")
	fmt.Scan(&email)

	pegawai := node.Pegawai {
		ID: id,
		Nama: nama,
		Alamat: node.Address{jalan, kota, nomer},
		NoTelp: notelp,
		Email: email,
	}

	cek := model.CreatePegawai(pegawai)
	if cek {
		fmt.Println("== Pegawai Berhasil Ditambahkan ==")
	} else {
		fmt.Println("Pegawai Gagal Ditambahkan")
	} 
	fmt.Println()
}

func Views () {
	fmt.Println("Daftar Pegawai")
	for i, emp := range model.ReadPegawai() {
		fmt.Println("--- Pegawai ke - ", i + 1, " ---")
		fmt.Println("ID Pegawai\t : ", emp.ID)
		fmt.Println("ID Pegawai\t : ", emp.Nama)
		fmt.Println("ID Pegawai\t : ", emp.Alamat.Jalan, emp.Alamat.Nomer, emp.Alamat.Kota)
		fmt.Println("ID Pegawai\t : ", emp.NoTelp)
		fmt.Println("ID Pegawai\t : ", emp.Email)
		fmt.Println()
	}
}

func Update ()  {
	var id, nomer int
	var kota, nama, notelp, email string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukan ID Pegawai : ")
	fmt.Scan(&id)

	fmt.Print("Masukan Nama Pegawai : ")
	nama, _ = reader.ReadString('\n')
	nama = nama[:len(nama)-1]

	fmt.Print("Masukan Kota : ")
	fmt.Scan(&kota)

	fmt.Print("Masukan Jalan : ")
	jalan, _ := reader.ReadString('\n')
	jalan = jalan[:len(jalan)-1]

	fmt.Print("Masukan Nomer Rumah : ")
	fmt.Scan(&nomer)

	fmt.Print("Masukan Nomer Telepon : ")
	fmt.Scan(&notelp)

	fmt.Print("Masukan Email : ")
	fmt.Scan(&email)

	pegawai := node.Pegawai {
		ID: id,
		Nama: nama,
		Alamat: node.Address{jalan, kota, nomer},
		NoTelp: notelp,
		Email: email,
	}

	cek := model.CreatePegawai(pegawai)
	if cek {
		fmt.Println("== Pegawai Berhasil Ditambahkan ==")
	} else {
		fmt.Println("Pegawai Gagal Ditambahkan")
	} 
	fmt.Println()
}

func Delete() {
	var id int
	fmt.Print("Masukan ID Pegawai Yang Akan Dihapus : ")
	fmt.Scan(&id)

	cek := model.DeletePegawai(id)
	if cek {
		fmt.Println("== Pegawai Berhasil Dihapus ==")
	} else {
		fmt.Println("Pegawai Gagal Dihapus")
	}
	fmt.Println()
}
