package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// type pegawai
type tPegawai struct {
	username, password string
}

// type menu
const MAXMENU int = 200
type tMenu struct {
	kategori                     string
	nama                         string
	harga, stok, idMenu, terjual int
}
type tabMenu [MAXMENU]tMenu

//type transaksi
const MAXTRANS int = 100
type transaksi struct {
	namaMenu    [MAXTRANS]string
	banyakBeli  int
	idTrans     int
	namaPembeli string
	hargaMenu   [MAXTRANS]int
}
type tabTransaksi [MAXTRANS]transaksi

func main() {
	// var terjual tabbeli
	var menu tabMenu
	var nMenu int
	var index int
	var nTrans int
	var transaksi tabTransaksi
	var pegawai tPegawai
	var pilih int
	login(&pegawai)
	MenuAwal(&menu, &nMenu)
	TransaksiAwal(&transaksi, &menu, &nTrans)
	for {
		beranda(&pilih)
		switch pilih {
		case 1:
			DataMenu(&menu, &nMenu)
		case 2:
			CariMenu(menu, nMenu, index)
		case 3:
			MenuBestSeller(&menu, nMenu)
		case 4:
			TambahMenu(&menu, &nMenu)
		case 5:
			UbahMenu(&menu, nMenu)
		case 6:
			HapusMenu(&menu, &nMenu)
		case 7:
			dataTransaksi(&transaksi, &nTrans, &menu, &nMenu)
		case 8:
			modal(&transaksi, &nTrans, menu, nMenu)
		}
	}
}
func clearScreen() {
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func login(pgw *tPegawai){
/* Fungsi login pengguna */
clearScreen()
fmt.Println("\x1b[35m====================================\x1b[0m")
fmt.Println("\x1b[35m||\x1b[0m          SWEET n'Sour          \x1b[35m||\x1b[0m")
fmt.Println("\x1b[35m====================================\x1b[0m")
fmt.Println("Silakan input username dan password")
fmt.Printf("            Username:\n             ")
fmt.Scan(&pgw.username)
fmt.Printf("            Password:\n             ")
fmt.Scan(&pgw.password)
clearScreen()

}

func beranda(p *int) {
/* Fungsi untuk memilih menu awal bagi pengguna */
	fmt.Println("\x1b[35m====================================\x1b[0m")
	fmt.Println("\x1b[35m||\x1b[0m          SWEET n'Sour          \x1b[35m||\x1b[0m")
	fmt.Println("\x1b[35m====================================\x1b[0m")
	fmt.Println("\x1b[35m|\x1b[0m1. Data Menu                      \x1b[35m|\x1b[0m")
	fmt.Println("\x1b[35m|\x1b[0m2. Cari Menu                      \x1b[35m|\x1b[0m")
	fmt.Println("\x1b[35m|\x1b[0m3. Menu Best Seller               \x1b[35m|\x1b[0m")
	fmt.Println("\x1b[35m|\x1b[0m4. Tambah Menu                    \x1b[35m|\x1b[0m")
	fmt.Println("\x1b[35m|\x1b[0m5. Ubah Menu                      \x1b[35m|\x1b[0m")
	fmt.Println("\x1b[35m|\x1b[0m6. Hapus menu                     \x1b[35m|\x1b[0m")
	fmt.Println("\x1b[35m|\x1b[0m7. Transaksi                      \x1b[35m|\x1b[0m")
	fmt.Println("\x1b[35m|\x1b[0m8. Modal dan pendapatan           \x1b[35m|\x1b[0m")
	fmt.Println("\x1b[35m|\x1b[0m9. Exit                           \x1b[35m|\x1b[0m")
	fmt.Println("\x1b[35m====================================\x1b[0m")
	fmt.Print("\x1b[33mPilih (1/2/3/4/5/6/7/8/9): \x1b[0m")
	fmt.Print("\x1b[33m")
	fmt.Scan(p)
	fmt.Print("\x1b[0m")
	clearScreen()
	fmt.Println("------------------------------")
	clearScreen()
}
func MenuAwal(MN *tabMenu, n *int) {
/* Data dummy untuk menu awal yang sudah tersedia pada program
	dengan asumsi harga = harga awal + keuntungan
	*/
	MN[0].kategori = "Drink"
	MN[0].idMenu = 11
	MN[0].nama = "Jus"
	MN[0].harga = 12000
	MN[0].stok = 30

	MN[1].kategori = "Drink"
	MN[1].idMenu = 12
	MN[1].nama = "Milkshake"
	MN[1].harga = 20000
	MN[1].stok = 30

	MN[2].kategori = "Dessert"
	MN[2].idMenu = 21
	MN[2].nama = "Puding"
	MN[2].harga = 15000
	MN[2].stok = 25

	MN[3].kategori = "Dessert"
	MN[3].idMenu = 22
	MN[3].nama = "Pastry"
	MN[3].harga = 15000
	MN[3].stok = 22

	MN[4].kategori = "Snack"
	MN[4].idMenu = 31
	MN[4].nama = "Kentang"
	MN[4].harga = 12000
	MN[4].stok = 30

	MN[5].kategori = "Snack"
	MN[5].idMenu = 32
	MN[5].nama = "Onion_Ring"
	MN[5].harga = 120004
	MN[5].stok = 30

	*n = 6
}
func TransaksiAwal(T *tabTransaksi, MN *tabMenu, nTrans *int) {
/* Data dummy untuk transaksi awal yang sudah tersedia pada program */
	T[0].idTrans = 1
	T[0].namaPembeli = "john wick"
	T[0].banyakBeli = 2
	T[0].namaMenu[0] = "Jus"
	T[0].namaMenu[1] = "Onion_Ring"
	T[0].hargaMenu[0] = 12000
	T[0].hargaMenu[1] = 20000
	MN[0].terjual++
	MN[5].terjual++

	T[1].idTrans = 2
	T[1].namaPembeli = "Kane"
	T[1].banyakBeli = 2
	T[1].namaMenu[0] = "Pastry"
	T[1].namaMenu[1] = "Jus"
	T[1].hargaMenu[0] = 12000
	T[1].hargaMenu[1] = 12000
	MN[3].terjual++
	MN[0].terjual++

	T[2].idTrans = 3
	T[2].namaPembeli = "Johnny"
	T[2].banyakBeli = 2
	T[2].namaMenu[0] = "Milkshake"
	T[2].namaMenu[1] = "Kentang"
	T[2].hargaMenu[0] = 12000
	T[2].hargaMenu[1] = 20000
	MN[1].terjual++
	MN[4].terjual++

	*nTrans = 3
}

func DataMenu(MN *tabMenu, n *int) {
//Melihat menu yang tersedia berdasarkan kategori
	clearScreen()
	var p int
	fmt.Println("Kategori menu:")
	fmt.Println("1. Drink")
	fmt.Println("2. Dessert")
	fmt.Println("3. Snack")
	fmt.Println("------------------------------")
	fmt.Println("(Ketik \x1b[34m0\x1b[0m untuk kembali)")
	fmt.Print("Pilih kategori: ")
	fmt.Scan(&p)
	clearScreen()
	fmt.Println()
	if p == 1 {
		fmt.Println("\x1b[35m--------KATEGORI DRINK--------\x1b[0m")
		for i := 0; i < *n; i++ {
			if MN[i].kategori == "Drink" {
				fmt.Printf("ID menu: %d \nNama   : % s\nHarga  : %d\nstok   : %d\n", MN[i].idMenu, MN[i].nama, MN[i].harga, MN[i].stok)
				fmt.Println(("\x1b[35m------------------------------\x1b[0m"))
			}
		}
		var back string
		fmt.Print("Back? (\x1b[34my\x1b[0m/\x1b[34mn\x1b[0m): ")
		fmt.Print("\x1b[34m")
		fmt.Scan(&back)
		fmt.Print("\x1b[0m")
		fmt.Println("------------------------------")
		clearScreen()
		if back == "y" {
			DataMenu(MN, n)
		}
	} else if p == 2 {
		fmt.Println("\x1b[35m-------KATEGORI DESSERT-------\x1b[0m")
		for i := 0; i < *n; i++ {
			if MN[i].kategori == "Dessert" {
				fmt.Printf("ID menu: %d \nNama   : % s\nHarga  : %d\nstok   : %d\n", MN[i].idMenu, MN[i].nama, MN[i].harga, MN[i].stok)
				fmt.Println(("\x1b[35m------------------------------\x1b[0m"))
			}
		}
		var back string
		fmt.Print("Back? (\x1b[34my\x1b[0m/\x1b[34mn\x1b[0m): ")
		fmt.Print("\x1b[34m")
		fmt.Scan(&back)
		fmt.Print("\x1b[0m")
		fmt.Println("------------------------------")
		clearScreen()
		if back == "y" {
			DataMenu(MN, n)
		}
	} else if p == 3 {
		fmt.Println("\x1b[35m--------KATEGORI SNACK--------\x1b[0m")
		for i := 0; i < *n; i++ {
			if MN[i].kategori == "Snack" {
				fmt.Printf("ID menu: %d \nNama   : % s\nHarga  : %d\nstok   : %d\n", MN[i].idMenu, MN[i].nama, MN[i].harga, MN[i].stok)
				fmt.Println(("\x1b[35m------------------------------\x1b[0m"))
			}
		}
		var back string
		fmt.Print("Back? (\x1b[34my\x1b[0m/\x1b[34mn\x1b[0m): ")
		fmt.Print("\x1b[34m")
		fmt.Scan(&back)
		fmt.Print("\x1b[0m")
		fmt.Println("------------------------------")
		clearScreen()
		if back == "y" {
			DataMenu(MN, n)
		}
	} else {

	}
	clearScreen()
}
func cetakDataMenu(MN tabMenu, n int) {
	//Mencetak data menu yang tersedia di prosedur menuAWal 
	clearScreen()
	for i := 0; i < n; i++ {
		fmt.Printf("Kategori  : %s\nID Menu   : %d\nNama menu : %s\nStok      : %d\nHarga menu: %d\n", MN[i].kategori, MN[i].idMenu, MN[i].nama, MN[i].stok, MN[i].harga)
		fmt.Println("\x1b[35m------------------------------\x1b[0m")
	}
}
func CariMenu(MN tabMenu, n int, idx int) {
// Mencari menu berdasarkan ID, Nama, dan harga termurah
	clearScreen()
	var x, p int
	fmt.Println("Cari berdasarkan?")
	fmt.Println("1. ID menu")
	fmt.Println("2. Nama menu")
	fmt.Println("3. Harga Termurah")
	fmt.Print("Pilih kategori: ")
	fmt.Scan(&x)
	fmt.Println("------------------------------")
	clearScreen()
	if x == 2 {
		idx = sequential_idxnama(MN, n, p)
		if idx != -1 {
			fmt.Printf("Berikut menu dengan nama %s:\n", MN[idx].nama)
			fmt.Printf("kategori: %s\nID      : %d\nNama    : %s\nHarga   : %d\nstok    : %d\n", MN[idx].kategori, MN[idx].idMenu, MN[idx].nama, MN[idx].harga, MN[idx].stok)
		} else if idx == -1 {
			fmt.Println("\x1b[31m*****Menu tidak ditemukan*****\x1b[0m")
		}
	} else if x == 1 {
		idx = binary_idxid(MN, n, p)
		if idx != -1 {
			fmt.Printf("Berikut menu dengan ID %d:\n", MN[idx].idMenu)
			fmt.Printf("kategori: %s\nID      : %d\nNama    : %s\nHarga   : %d\nstok    : %d\n", MN[idx].kategori, MN[idx].idMenu, MN[idx].nama, MN[idx].harga, MN[idx].stok)	
		} else if idx == -1 {
			fmt.Println("\x1b[31m*****Menu tidak ditemukan*****\x1b[0m")
		}
	}else if x == 3{
		MenuTermurah(&MN, n)

	}
	var back string
	fmt.Println("------------------------------")
	fmt.Print("Back? (\x1b[34my\x1b[0m/\x1b[34mn\x1b[0m): ")
	fmt.Print("\x1b[34m")
		fmt.Scan(&back)
		fmt.Print("\x1b[0m")
	fmt.Println("------------------------------")
	if back == "y" {
		CariMenu(MN, n, idx)
	} else {

	}
	clearScreen()
}
func binary_idxid(MN tabMenu, n int, idx int) int {
// Pencarian index ID menu yang dicari menggunakan binary search
	var left, right, mid, cari int
	idx = -1
	left = 0
	right = n
	fmt.Print("Masukan ID: ")
	fmt.Scan(&cari)
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if MN[mid].idMenu < cari {
			left = mid + 1
		} else if MN[mid].idMenu > cari {
			right = mid - 1
		} else {
			idx = mid
		}
	}
	return idx
}
func sequential_idxnama(MN tabMenu, n int, idx int) int {
//Pecarian index nama menu yang dicari menggunakan sequential search 
	var cari string
	var i int
	idx = -1
	i = 0
	fmt.Print("Masukkan nama menu: ")
	fmt.Scan(&cari)
	for i < n && idx == -1 {
		if cari == MN[i].nama {
			idx = i
		}
		i = i + 1
	}
	return idx
}

func TambahMenu(MN *tabMenu, nMenu *int) {
// Untuk menambah menu yang ingin ditambahkan kedalam prosedur menuAwal
	var lanjut string
	var cek int
	fmt.Println("Masukkan menu yang ingin ditambahkan")
	fmt.Print("Kategori menu    : ")
	fmt.Scan(&MN[*nMenu].kategori)
	for i := 0; i < *nMenu && cek == 0; i++ {
		if MN[*nMenu].kategori == MN[i].kategori {
			cek = 1
		}
	}
	if cek == 1 {
		fmt.Print("ID menu          : ")
		fmt.Scan(&MN[*nMenu].idMenu)
		fmt.Print("Nama menu        : ")
		fmt.Scan(&MN[*nMenu].nama)
		fmt.Print("Harga menu       : ")
		fmt.Scan(&MN[*nMenu].harga)
		fmt.Print("Stok menu        : ")
		fmt.Scan(&MN[*nMenu].stok)
		for i := 0; i < *nMenu; i++ {
			if MN[i].idMenu == MN[*nMenu].idMenu || MN[i].nama == MN[*nMenu].nama {
				fmt.Println("Nama atau ID sudah terdaftar mohon ganti salah satu atau keduanya")
				fmt.Print("Kategori menu    : ")
				fmt.Scan(&MN[*nMenu].kategori)
				fmt.Print("ID menu          : ")
				fmt.Scan(&MN[*nMenu].idMenu)
				fmt.Print("Nama menu        : ")
				fmt.Scan(&MN[*nMenu].nama)
				fmt.Print("Harga menu       : ")
				fmt.Scan(&MN[*nMenu].harga)
				fmt.Print("Stok menu        : ")
				fmt.Scan(&MN[*nMenu].stok)
				fmt.Println("\x1b[32m*****Menu berhasil ditambahkan*****\x1b[0m")
			}
		}
		*nMenu++
		cek = 0
		fmt.Println("------------------------------")
		fmt.Println("1. lanjut")
		fmt.Println("2. stop")
		fmt.Print("\x1b[33mPilih(1/2): \x1b[0m")
		fmt.Print("\x1b[33m")
		fmt.Scan(&lanjut)
		fmt.Print("\x1b[0m")
		fmt.Println("------------------------------")
		if lanjut == "1" {
			fmt.Println("Silahkan masukkan kembali menu yang ingin ditambahkan")
			fmt.Print("Kategori menu    : ")
			fmt.Scan(&MN[*nMenu].kategori)
			for i := 0; i < *nMenu && cek == 0; i++ {
				if MN[*nMenu].kategori == MN[i].kategori {
					cek = 1
				}
			}
			if cek == 1 {
				fmt.Print("ID menu          : ")
				fmt.Scan(&MN[*nMenu].idMenu)
				fmt.Print("Nama menu        : ")
				fmt.Scan(&MN[*nMenu].nama)
				fmt.Print("Harga menu     : ")
				fmt.Scan(&MN[*nMenu].harga)
				fmt.Print("Stok menu        : ")
				fmt.Scan(&MN[*nMenu].stok)
				fmt.Println("\x1b[32m*****Menu berhasil ditambahkan*****\x1b[0m")
				*nMenu++
			} else {
				fmt.Println("------------------------")
				fmt.Println("\x1b[31m*****Maaf, kategori tidak tersdia*****\x1b[0m")
				fmt.Printf("\x1b[31mSaat ini, kami hanya bisa menambahkan kategori\nDrink, Dessert, dan Snack\n\x1b[0m")
			}
		} else {
			fmt.Println("\x1b[32m*****Menu berhasil ditambahkan*****\x1b[0m")
		}
	} else {
		fmt.Println("-----------------------------")
		fmt.Println("\x1b[31m*****Maaf, kategori tidak tersdia*****\x1b[0m")
		fmt.Printf("\x1b[31mSaat ini, kami hanya bisa menambahkan kategori\nDrink, Dessert, dan Snack\n\x1b[0m")
	}
}
func UbahMenu(MN *tabMenu, n int) {
// Mengubah atau mengedit menu menggunakan seqquential seacrh
	var cari int
	cetakDataMenu(*MN, n)
	cari = sequential_idxnama(*MN, n, cari)
	if cari == -1 {
		fmt.Println("\x1b[31m*****Tidak terdapat menu seperti itu******\x1b[0m")
	} else if cari != -1 {
		fmt.Print("Kategori          : ")
		fmt.Scan(&MN[cari].kategori)
		fmt.Print("ID menu           : ")
		fmt.Scan(&MN[cari].idMenu)
		fmt.Print("Nama Menu         : ")
		fmt.Scan(&MN[cari].nama)
		fmt.Print("Harga Menu        : ")
		fmt.Scan(&MN[cari].harga)
		fmt.Print("Stok Menu         : ")
		fmt.Scan(&MN[cari].stok)
		fmt.Println("\x1b[32m*****Data menu berhasil di edit*****\x1b[0m")
	}
}
func HapusMenu(MN *tabMenu, n *int) {
// Menghapus menu yang sudah ada atau yang sudah ditambahkan menggunakan sequential search
	clearScreen()
	var i, cari int
	cetakDataMenu(*MN, *n)
	fmt.Println("Masukkan menu yang akan dihapus")
	cari = sequential_idxnama(*MN, *n, cari)
	if cari == -1 {
		fmt.Println("\x1b[31m*****Tidak terdapat menu seperti itu******\x1b[0m")
	} else if cari != -1 {
		for i = cari; i < *n; i++ {
			MN[i] = MN[i+1]
		}
		fmt.Println("\x1b[32m*****Menu berhasil di hapus*****\x1b[0m")
	}
	*n = *n - 1
	
}

func dataTransaksi(T *tabTransaksi, nTrans *int, MN *tabMenu, n *int) {
/* Fungsi untuk melihat data transaksi, tambah transaksi, edit transaksi, dan hapus transaksi */
	clearScreen()
	var pilih string
	fmt.Println("\x1b[35m====================================\x1b[0m")
	fmt.Println("\x1b[35m||\x1b[0m          SWEET n'Sour          \x1b[35m||\x1b[0m")
	fmt.Println("\x1b[35m====================================\x1b[0m")
	fmt.Println("\x1b[35m|\x1b[0m1. Cetak Data Transaksi           \x1b[35m|\x1b[0m")
	fmt.Println("\x1b[35m|\x1b[0m2. Tambah Data Transaksi          \x1b[35m|\x1b[0m")
	fmt.Println("\x1b[35m|\x1b[0m3. Edit Data Transaksi            \x1b[35m|\x1b[0m")
	fmt.Println("\x1b[35m|\x1b[0m4. Hapus Data Transaksi           \x1b[35m|\x1b[0m")
	fmt.Println("\x1b[35m|\x1b[0m5. Exit                           \x1b[35m|\x1b[0m")
	fmt.Println("\x1b[35m====================================\x1b[0m")
	fmt.Print("\x1b[33mPilih (1/2/3/4/5/6/7/8): \x1b[0m")
	fmt.Print("\x1b[33m")
	fmt.Scan(&pilih)
	fmt.Print("\x1b[0m")
	clearScreen()
	fmt.Println("------------------------------")
	for pilih != "5" {
		if pilih == "1" {
			cetakTransaksi(*T, *nTrans)
		} else if pilih == "2" {
			tambahDataTransaksi(T, nTrans, MN, n)
		} else if pilih == "3" {
			editDataTransaksi(T, nTrans, MN, n)
		} else if pilih == "4" {
			hapusDataTransaksi(T, nTrans, MN, n)
		} else {
		}
	fmt.Println("\x1b[35m====================================\x1b[0m")
	fmt.Println("\x1b[35m||\x1b[0m          SWEET n'Sour          \x1b[35m||\x1b[0m")
	fmt.Println("\x1b[35m====================================\x1b[0m")
	fmt.Println("\x1b[35m|\x1b[0m1. Cetak Data Transaksi           \x1b[35m|\x1b[0m")
	fmt.Println("\x1b[35m|\x1b[0m2. Tambah Data Transaksi          \x1b[35m|\x1b[0m")
	fmt.Println("\x1b[35m|\x1b[0m3. Edit Data Transaksi            \x1b[35m|\x1b[0m")
	fmt.Println("\x1b[35m|\x1b[0m4. Hapus Data Transaksi           \x1b[35m|\x1b[0m")
	fmt.Println("\x1b[35m|\x1b[0m5. Exit                           \x1b[35m|\x1b[0m")
	fmt.Println("\x1b[35m====================================\x1b[0m")
	fmt.Print("\x1b[33mPilih (1/2/3/4/5/6/7/8): \x1b[0m")
		fmt.Scan(&pilih)
		fmt.Println("------------------------------")
	}
	clearScreen()
}
func cetakTransaksi(T tabTransaksi, nTrans int) {
// Mencetak data dummy transaksi yang ada di prosedur transaksiAwal
	clearScreen()
	fmt.Println("\x1b[35m********Data Transaksi********\x1b[0m")
	for i := 0; i < nTrans; i++ {
		for j := 0; j < T[i].banyakBeli; j++ {
			fmt.Printf("ID Trans    : %d\nNama Pembeli: %s\nNama Menu   : %s\nHarga Menu  : %d\n",T[i].idTrans, T[i].namaPembeli, T[i].namaMenu[j], T[i].hargaMenu[j])
			fmt.Println("\x1b[35m------------------------------\x1b[0m")
		}
	}
}
func tambahDataTransaksi(T *tabTransaksi, nTrans *int, MN *tabMenu, n *int) {
// Menambah data transaksi ketika ada pembelian
	var nama, namaMenu string
	var simpan transaksi
	clearScreen()
	var banyakBeli int
	simpan = T[*nTrans+1]
	fmt.Println("ID Transaksi	:", *nTrans+1)
	T[*nTrans].idTrans = *nTrans + 1
	fmt.Println("------------------------------")
	fmt.Println("(Ketik \x1b[34m0\x1b[0m untuk kembali)")
	fmt.Print("Masukan nama pembeli: ")
	fmt.Scan(&nama)
	if nama == "0" {
		dataTransaksi(T, nTrans, MN, n)
	} else {
		clearScreen()
		var cek1 int
		T[*nTrans].namaPembeli = nama
		fmt.Print("Masukan banyak menu : ")
		fmt.Scan(&banyakBeli)
		T[*nTrans].banyakBeli = banyakBeli
		for i := 0; i < banyakBeli; i++ {
			fmt.Print("Masukan nama menu   : ")
			fmt.Scan(&namaMenu)
			//idx = sequential_idxnama(*MN, *n, namaMenu)
			cek1 = 0
			for j := 0; j < *n; j++ {
				if namaMenu == MN[j].nama {
					T[*nTrans].namaMenu[i] = MN[j].nama
					T[*nTrans].hargaMenu[i] = MN[j].harga
					MN[j].stok--
					MN[j].terjual++
					fmt.Printf("Harga menu  	    : %d\n", T[*nTrans].hargaMenu[i])
					cek1 = 1
				}
			}
		}
		if cek1 == 0 {
			fmt.Println("\x1b[31m*****Data transaksi gagal ditambahkan*****\x1b[0m")
			T[*nTrans+1] = simpan
		} else {
			cek1 = 0
			fmt.Println("\x1b[32m*****Data transaksi berhasil ditambahkan*****\x1b[0m")
			*nTrans += 1
		}
	}
}
func editDataTransaksi(T *tabTransaksi, nTrans *int, MN *tabMenu, n *int) {
// Mengubah atau mengedit data transaksi dengan mencari ID transaksi yang ingin diedit menggunakan sequential search
// Memasukan kembali banyak menu yang ingin dibeli, nama pembeli, dan nama menu 
	cetakTransaksi(*T, *nTrans)
	fmt.Println("------------------------------")
	var idt, idx, i int
	var simpan transaksi
	idx = -1
	i = 0
	fmt.Println("(Ketik \x1b[34m0\x1b[0m untuk kembali)")
	fmt.Print("Masukan ID transaksi yang akan di edit: ")
	fmt.Scan(&idt)
	if idt == 0 {
		dataTransaksi(T, nTrans, MN, n)
	} else {
		for i < *nTrans && idx == -1 {
			if idt == T[i].idTrans {
				idx = i
			}
			i = i + 1
		}
		clearScreen()
		if idx != -1 {
			var cek, idx2, cek2 int
			simpan = T[idx]
			fmt.Print("Banyak menu yang akan dibeli: ")
			fmt.Scan(&T[idx].banyakBeli)
			fmt.Print("Masukan nama pembeli        : ")
			fmt.Scan(&T[idx].namaPembeli)
			for j := 0; j < T[idx].banyakBeli; j++ {
				fmt.Print("Masukan nama menu           : ")
				fmt.Scan(&T[idx].namaMenu[j])
				cek2 = 0
				for k := 0; k < *n && cek == 0; k++ {
					if T[idx].namaMenu[j] == MN[k].nama {
						cek = 1
						idx2 = k
					}
				}
				if cek == 1 {
					cek = 0
					fmt.Println("Harga menu                  :", MN[idx2].harga)
					T[idx].hargaMenu[j] = MN[idx2].harga
					cek2 = 1
				}
			}
			if cek2 == 1 {
				fmt.Println("\x1b[32m*****Data transaksi berhasil di edit.*****\x1b[0m")
			} else {
				fmt.Println("\x1b[31m*****Data transaksi gagal di edit.*****\x1b[0m")
				T[idx] = simpan
			}
		} else {
			fmt.Println("\x1b[31m*****Data transaksi gagal di edit.*****\x1b[0m")
		}
	}

}
func hapusDataTransaksi(T *tabTransaksi, nTrans *int, MN *tabMenu, n *int) {
// Menghapus data transaksi dengan mencari ID transaksi yang ingin dihapus menggunakan sequential search
	clearScreen()
	cetakTransaksi(*T, *nTrans)
	fmt.Println("---------------------")
	var idt, idx, i int
	idx = -1
	i = 0
	fmt.Println("(Ketik \x1b[34m0\x1b[0m untuk kembali)")
	fmt.Print("Masukan ID taransaksi: ")
	fmt.Scan(&idt)
	if idt == 0 {
		dataTransaksi(T, nTrans, MN, n)
	} else {
		for i < *nTrans && idx == -1 {
			if idt == T[i].idTrans {
				idx = i
			}
			i = i + 1
		}
		if idx != -1 {
			for i := idx; i < *nTrans; i++ {
				T[i].namaPembeli = T[i+1].namaPembeli
				T[i].namaMenu = T[i+1].namaMenu
				T[i].banyakBeli = T[i+1].banyakBeli
			}
			fmt.Println("\x1b[32m*****Data transaksi berhasil dihapus*****\x1b[0m")
			*nTrans = *nTrans - 1
		} else {
			fmt.Println("\x1b[31m*****Data transaksi tidak ditemukan****\x1b[0m")
		}
	}
	
}

func modal(T *tabTransaksi, nTrans *int, MN tabMenu, n int) {
/* Untuk melihat data modal, pendapatan kotor, pendapatan bersih */
	clearScreen()
	var pilih string
	fmt.Println("\x1b[35m====================================\x1b[0m")
	fmt.Println("\x1b[35m||\x1b[0m         SWEET n'Sour           \x1b[35m||\x1b[0m")
	fmt.Println("\x1b[35m====================================\x1b[0m")
	fmt.Println("\x1b[35m|\x1b[0m1. Data Modal                    \x1b[35m |\x1b[0m")
	fmt.Println("\x1b[35m|\x1b[0m2. Pendapatan Kotor               \x1b[35m|\x1b[0m")
	fmt.Println("\x1b[35m|\x1b[0m3. Pendapatan Bersih              \x1b[35m|\x1b[0m")
	fmt.Println("\x1b[35m|\x1b[0m4. Exit                           \x1b[35m|\x1b[0m")
	fmt.Println("\x1b[35m====================================\x1b[0m")
	fmt.Print("\x1b[33mPilih (1/2/3/4/): \x1b[0m")
	fmt.Print("\x1b[33m")
	fmt.Scan(&pilih)
	fmt.Print("\x1b[0m")
	clearScreen()
	fmt.Println("------------------------------")
	for pilih != "4" {
		if pilih == "1" {
			modalAwal(MN, n)
		} else if pilih == "2" {
			cetakPenghasilanKotor(*T, *nTrans)
		} else if pilih == "3" {
			PendapatanBersih(*T, *nTrans)
		}
		fmt.Println("\x1b[35m====================================\x1b[0m")
		fmt.Println("\x1b[35m||\x1b[0m         SWEET n'Sour           \x1b[35m||\x1b[0m")
		fmt.Println("\x1b[35m====================================\x1b[0m")
		fmt.Println("\x1b[35m|\x1b[0m1. Data Modal                    \x1b[35m |\x1b[0m")
		fmt.Println("\x1b[35m|\x1b[0m2. Pendapatan Kotor               \x1b[35m|\x1b[0m")
		fmt.Println("\x1b[35m|\x1b[0m3. Pendapatan Bersih              \x1b[35m|\x1b[0m")
		fmt.Println("\x1b[35m|\x1b[0m4. Exit                           \x1b[35m|\x1b[0m")
		fmt.Println("\x1b[35m====================================\x1b[0m")
		fmt.Print("\x1b[33mPilih (1/2/3/4/): \x1b[0m")
		fmt.Print("\x1b[33m")
		fmt.Scan(&pilih)
		fmt.Print("\x1b[0m")
		fmt.Println("------------------------------")
	}
	clearScreen()
}
func modalAwal(MN tabMenu, n int) {
// Menghitung total modal awal berdasarkan menu yang tersedia di prosedur menuAwal
clearScreen()
	var i int
	var modal int

	for i = 0; i < n; i++ {
		modal += MN[i].harga * MN[i].stok
	}
	cetakDataMenu(MN, n)
	fmt.Println("\x1b[31mModal:\x1b[0m Rp", modal)
}
func cetakPenghasilanKotor(T tabTransaksi, nTrans int) {
// Mengitung pendapatan kotor berdasarkan transaksi/pembelian yang ada
	var pendapatan int
	for i := 0; i < nTrans; i++ {
		for j := 0; j < T[j].banyakBeli; j++ {
			pendapatan += T[i].hargaMenu[j]
		}
	}
	cetakTransaksi(T, nTrans)
	fmt.Println("\x1b[31mPendapatan kotor Rp:\x1b[0m",pendapatan)

}
func PendapatanBersih(T tabTransaksi, nTrans int) {
clearScreen()
//Mengitung pendapatan bersih/keuntungan yang didapat dari transaksi yang ada
	/*1. harga menu 12.000 - 15.000, keuntungan = 2000 untuk setiap menu
	  2. harga menu 16.000 - 20.000, keuntungan = 3000 untuk setiap menu
	  3. harga menu 21.000 - 25.000, keuntungan = 4500 untuk setiap menu
	  4. harga menu 26.000 - 30.000, keuntungan = 5000 untuk setiap  menu */
	var harga, pendapatanBersih int
	for i := 0; i < nTrans; i++ {
		for j := 0; j < T[i].banyakBeli; j++ {
			harga = T[i].hargaMenu[j]
			if harga >= 12000 && harga <= 15000 {
				harga = 2000
			} else if harga >= 16000 && harga <= 20000 {
				harga = 3000
			} else if harga >= 21000 && harga <= 25000 {
				harga = 4500
			} else if harga >= 26000 && harga <= 30000 {
				harga = 5000
			}
			pendapatanBersih += harga
		}
	}
	fmt.Println("\x1b[31mPendapatan bersih Rp:\x1b[0m",pendapatanBersih)

}
func MenuBestSeller(dataMenu *tabMenu, NMenu int) {
// Menampilkan 5 menu yang paling banyak terjual menggunakan selection sort
	selectionSort(*&dataMenu, NMenu)
	fmt.Println("\x1b[35m**Top 5 menu terlaris**\x1b[0m")
	for i := 0; i < 5; i++ {
		fmt.Printf("Nama menu: %s\nterjual: %d\n",dataMenu[i].nama, dataMenu[i].terjual)
		fmt.Println("\x1b[35m------------------------\x1b[0m")

	}
}
func MenuTermurah(dataMenu *tabMenu, NMenu int) {
// Menampilkan menu dari harga yang paling murah menggunakan insertion sort
	insertionSort(*&dataMenu, NMenu)
	fmt.Println("\x1b[35m**Berdasarkan harga termurah**\x1b[0m")
	for i := 0; i < NMenu; i++ {
		fmt.Printf("Nama menu : %s\nHarga menu: %d\n", dataMenu[i].nama, dataMenu[i].harga)
		fmt.Println("\x1b[35m------------------------------\x1b[0m")

	}
}
func selectionSort(MN *tabMenu, n int) {
// Mengurutkan menu secara descending (Menurun) berdasarkan banyaknya menu yang terjual
	var i, j, idx int
	var t tMenu
	i = 1
	for i < n {
		idx = i - 1
		j = i
		for j < n {
			if MN[idx].terjual < MN[j].terjual {
				idx = j
			}
			j++
		}
		t = MN[idx]
		MN[idx] = MN[i-1]
		MN[i-1] = t
		i++
	}
}

func insertionSort(MN *tabMenu, n int) {
// Mengurutkan menu secara ascending (Menaik) berdasarkan harga
	var i, j int
	var temp tMenu
	i = 1
	for i <= n-1 {
		j = i
		temp = MN[j]
		for j > 0 && temp.harga < MN[j-1].harga {
			MN[j] = MN[j-1]
			j--
		}
		MN[j] = temp
		i++
	}
}
