package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Siswa struct {
	ID     int
	Nama   string
	Umur   int
	Gender string
	Kota   string
	Kelas  int
}

var siswaData []Siswa
var currentID int

func main() {
	for {
		MenuSiswa()
	}
}

func containsIgnoreCase(a, b string) bool {
	return strings.Contains(strings.ToLower(a), strings.ToLower(b))
}

func MenuSiswa() {
	fmt.Println("\n=== MENU DATA SISWA ===")
	fmt.Println("1. Tambah Data Siswa")
	fmt.Println("2. Lihat Data Siswa")
	fmt.Println("3. Hapus Data Siswa")
	fmt.Println("4. Ubah Data Siswa")
	fmt.Println("5. Cari Data Siswa")
	fmt.Println("6. Keluar")
	fmt.Println("")
	fmt.Print("Pilih Menu 1-6: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	choice := scanner.Text()

	switch choice {
	case "1":
		TambahSiswa()
	case "2":
		LihatSiswa()
	case "3":
		HapusSiswa()
	case "4":
		EditSiswa()
	case "5":
		CariSiswa()
	case "6":
		fmt.Println("Menu Data Siswa Ditutup")
		os.Exit(0)
	default:
		fmt.Println("Pilihan Tidak Ada. Coba lagi")
	}
}

func TambahSiswa() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Nama Siswa: ")
	scanner.Scan()
	nama := scanner.Text()

	fmt.Print("Umur Siswa: ")
	scanner.Scan()
	umur, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Umur harus berupa angka")
		return
	}

	fmt.Print("Kelas: ")
	scanner.Scan()
	kelas, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Kelas harus berapa angka")
		return
	}

	fmt.Print("Jenis Kelamin (L/P): ")
	scanner.Scan()
	Gender := scanner.Text()
	if Gender != "L" && Gender != "P" {
		fmt.Println("Jenis kelamin harus 'L' atau 'P' ")
		return
	}

	fmt.Print("Asal Kota: ")
	scanner.Scan()
	Kota := scanner.Text()

	currentID++
	newSiswa := Siswa{ID: currentID, Nama: nama, Umur: umur, Kelas: kelas, Gender: Gender, Kota: Kota}
	siswaData = append(siswaData, newSiswa)

	fmt.Println("Data siswa berhasil ditambahkan!")
}

func LihatSiswa() {
	if len(siswaData) == 0 {
		fmt.Println("Belum ada data siswa.")
		return
	}

	// Mengubah Data Siswa ke JSON
	jsonData, err := json.MarshalIndent(siswaData, "", "  ")
	if err != nil {
		fmt.Println("Terjadi kesalahan saat mengubah ke JSON:", err)
		return
	}

	fmt.Println("\n=== DATA SISWA ===")
	fmt.Println(string(jsonData)) // Menampilkan JSON sebagai string
}

func HapusSiswa() {
	if len(siswaData) == 0 {
		fmt.Println("Belum ada data siswa")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan ID siswa yang ingin dihapus: ")
	scanner.Scan()
	idStr := scanner.Text()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID harus berupa angka.")
		return
	}

	// Menghapus Siswa Dengan ID!
	for i, siswa := range siswaData {
		if siswa.ID == id {
			fmt.Printf("Apakah Anda yakin untuk menghapus siswa dengan ID %d? (y/n): ", id)
			scanner.Scan()
			konfirmasi := scanner.Text()
			if konfirmasi == "y" {
				siswaData = append(siswaData[:i], siswaData[i+1:]...)
				fmt.Printf("Siswa dengan ID %d berhasil dihapus.\n", id)
			} else {
				fmt.Println("Penghapusan dibatalkan.")
			}
			return
		}
	}

	fmt.Printf("Siswa dengan ID %d tidak ditemukan.\n", id)
}
func EditSiswa() {
	if len(siswaData) == 0 {
		fmt.Println("Belum ada data siswa.")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan ID siswa yang ingin diubah: ")
	scanner.Scan()
	idStr := scanner.Text()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID harus berupa angka. Coba lagi.")
		return
	}

	// Cari siswa berdasarkan Nama
	for i, siswa := range siswaData {
		if siswa.ID == id {
			fmt.Printf("Mengubah data untuk siswa dengan Nama %d.\n", id)

			// Edit Nama
			fmt.Printf("Nama Lama: %s\nMasukkan Nama Baru (tekan Enter untuk melewati): ", siswa.Nama)
			scanner.Scan()
			namaBaru := scanner.Text()
			if namaBaru != "" {
				siswaData[i].Nama = namaBaru
			}

			// Edit Umur
			fmt.Printf("Umur Lama: %d\nMasukkan Umur Baru (tekan Enter untuk melewati): ", siswa.Umur)
			scanner.Scan()
			umurStr := scanner.Text()
			if umurStr != "" {
				umurBaru, err := strconv.Atoi(umurStr)
				if err != nil {
					fmt.Println("Umur harus berupa angka. Coba lagi.")
					return
				}
				siswaData[i].Umur = umurBaru
			}

			// Edit Gender
			fmt.Printf("Jenis Kelamin Lama: %s\nMasukkan Jenis Kelamin Baru (L/P, tekan Enter untuk melewati): ", siswa.Gender)
			scanner.Scan()
			genderBaru := scanner.Text()
			if genderBaru != "" && (genderBaru == "L" || genderBaru == "P") {
				siswaData[i].Gender = genderBaru
			}

			// Edit Kota
			fmt.Printf("Kota Lama: %s\nMasukkan Kota Baru (tekan Enter untuk melewati): ", siswa.Kota)
			scanner.Scan()
			kotaBaru := scanner.Text()
			if kotaBaru != "" {
				siswaData[i].Kota = kotaBaru
			}

			fmt.Printf("Data siswa dengan Nama %d berhasil diperbarui.\n", id)
			return
		}
	}

	fmt.Printf("Siswa dengan Nama %d tidak ditemukan.\n", id)
}

func CariSiswa() {
	if len(siswaData) == 0 {
		fmt.Println("Belum ada data siswa.")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan Nama siswa yang ingin dicari: ")
	scanner.Scan()
	nama := scanner.Text()

	fmt.Println("\n=== HASIL PENCARIAN ===")
	found := false
	for _, siswa := range siswaData {
		if containsIgnoreCase(siswa.Nama, nama) {
			fmt.Printf("ID: %d, Nama: %s, Umur: %d, Jenis Kelamin: %s, Kota: %s\n",
				siswa.ID, siswa.Nama, siswa.Umur, siswa.Gender, siswa.Kota)
			found = true
		}
	}

	if !found {
		fmt.Printf("Tidak ada siswa dengan nama yang mengandung '%s'.\n", nama)
	}
}
