package main

import "fmt"

type dataTim struct {
	namaTim    string
	negara     string
	cabangGame string
	win        int
	draw       int
	loss       int
	poin       int
}

type tabTim [1000]dataTim

var timArray tabTim
var jumlahTim int = 0

func main() {
	menu()
}

func menu() {
	var userInput int
	var exit = false

	for !exit {
		fmt.Println("\n=== E-SPORTS TOURNAMENT SYSTEM ===")
		fmt.Println("1. Input Data Tim")
		fmt.Println("2. Edit Data Tim / Update Skor")
		fmt.Println("3. Hapus Data Tim")
		fmt.Println("4. Tampilkan Klasemen")
		fmt.Println("0. Exit")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&userInput)

		switch userInput {
		case 1:
			inputDataTim()
		case 2:
			editDataTim()
		case 3:
			hapusDataTim()
		case 4:
			tampilkanKlasemen()
		case 0:
			exit = true
		default:
			fmt.Println("Pilihan tidak valid. Mohon masukkan pilihan yang valid!")
		}
	}
}

func inputDataTim() {
	if jumlahTim >= 1000 {
		fmt.Println("Kapasitas turnamen penuh!")
		return
	}

	var t dataTim
	fmt.Print("Masukkan Nama Tim: ")
	fmt.Scan(&t.namaTim)
	fmt.Print("Masukkan Negara Asal: ")
	fmt.Scan(&t.negara)
	fmt.Print("Masukkan Cabang Game: ")
	fmt.Scan(&t.cabangGame)
	fmt.Print("Masukkan Jumlah Menang (Win): ")
	fmt.Scan(&t.win)
	fmt.Print("Masukkan Jumlah Seri (Draw): ")
	fmt.Scan(&t.draw)
	fmt.Print("Masukkan Jumlah Kalah (Loss): ")
	fmt.Scan(&t.loss)

	t.poin = (t.win * 3) + (t.draw * 1)

	timArray[jumlahTim] = t
	jumlahTim++

	fmt.Println("✅ Tim berhasil ditambahkan!")
}

func editDataTim() {
	if jumlahTim == 0 {
		fmt.Println("Belum ada data tim yang terdaftar.")
		return
	}

	var namaDicari string
	fmt.Print("Masukkan Nama Tim yang ingin diubah: ")
	fmt.Scan(&namaDicari)

	found := false
	i := 0
	for !found && i < jumlahTim {
		if timArray[i].namaTim == namaDicari {
			found = true
			fmt.Print("Masukkan Nama Tim Baru: ")
			fmt.Scan(&timArray[i].namaTim)
			fmt.Print("Masukkan Negara Asal Baru: ")
			fmt.Scan(&timArray[i].negara)
			fmt.Print("Masukkan Cabang Game Baru: ")
			fmt.Scan(&timArray[i].cabangGame)
			fmt.Print("Masukkan Jumlah Menang (Win) Baru: ")
			fmt.Scan(&timArray[i].win)
			fmt.Print("Masukkan Jumlah Seri (Draw) Baru: ")
			fmt.Scan(&timArray[i].draw)
			fmt.Print("Masukkan Jumlah Kalah (Loss) Baru: ")
			fmt.Scan(&timArray[i].loss)

			timArray[i].poin = (timArray[i].win * 3) + (timArray[i].draw * 1)
			
			fmt.Println("Data tim dan skor berhasil diperbarui!")
		}
		i++
	}

	if !found {
		fmt.Println("Tim tidak ditemukan.")
	}
}

func hapusDataTim() {
	if jumlahTim == 0 {
		fmt.Println("Belum ada data tim yang terdaftar.")
		return
	}

	var namaDicari string
	fmt.Print("Masukkan Nama Tim yang ingin dihapus: ")
	fmt.Scan(&namaDicari)

	found := false
	i := 0
	for !found && i < jumlahTim {
		if timArray[i].namaTim == namaDicari {
			found = true
			for j := i; j < jumlahTim-1; j++ {
				timArray[j] = timArray[j+1]
			}
			jumlahTim--
			fmt.Println("Data tim berhasil dihapus.")
		}
		i++
	}

	if !found {
		fmt.Println("Tim tidak ditemukan.")
	}
}

func tampilkanKlasemen() {
	if jumlahTim == 0 {
		fmt.Println("\n--- KLASEMEN SAAT INI ---")
		fmt.Println("Belum ada data tim yang terdaftar.")
		return
	}
	var gameDicari string
	fmt.Print("\nMasukkan Cabang Game yang ingin dilihat klasemennya : ")
	fmt.Scan(&gameDicari)

	for i := 0; i < jumlahTim-1; i++ {
		maxIdx := i
		for j := i + 1; j < jumlahTim; j++ {
			if timArray[j].poin > timArray[maxIdx].poin {
				maxIdx = j
			}
		}
		timArray[i], timArray[maxIdx] = timArray[maxIdx], timArray[i]
	}

	fmt.Printf("\n--- KLASEMEN GAME: %s ---\n", gameDicari)
	fmt.Printf("%-3s | %-15s | %-12s | %-10s | %-3s | %-3s | %-4s | %-4s\n", "No", "Nama Tim", "Negara", "Game", "W", "D", "L", "Poin")
	fmt.Println("---------------------------------------------------------------------------------")
	
	nomorUrut := 1
	adaTim := false
	
	for i := 0; i < jumlahTim; i++ {
		if timArray[i].cabangGame == gameDicari {
			fmt.Printf("%-3d | %-15s | %-12s | %-10s | %-3d | %-3d | %-4d | %-4d\n", 
				nomorUrut, timArray[i].namaTim, timArray[i].negara, timArray[i].cabangGame, timArray[i].win, timArray[i].draw, timArray[i].loss, timArray[i].poin)
			nomorUrut++
			adaTim = true
		}
	}

	if !adaTim {
		fmt.Printf("Belum ada data tim untuk cabang game %s.\n", gameDicari)
	}
}