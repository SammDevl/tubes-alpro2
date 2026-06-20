package main
import "fmt"

const maksimal = 1000

type Jadwal struct {
	namamk, dosen, ruang, hari string
	jamMulai, jamSelesai int
}

type TabJadwal [maksimal]Jadwal

func cekKonflikJadwal(A TabJadwal, n int, hari string, mulai, selesai int) bool {
	var i int
	var konflik bool
	
	i = 0
	konflik = false
	
	for i < n && konflik == false {
		if A[i].hari == hari {
			if mulai < A[i].jamSelesai && selesai > A[i].jamMulai {
				konflik = true
				fmt.Printf("\nKonflik! Sudah ada kelas %s dengan Dosen %s Ruangan %s Pada Hari %s Jam %d-%d\n", A[i].namamk, A[i].dosen, A[i].ruang, A[i].hari, A[i].jamMulai, A[i].jamSelesai)
			}
		}
		i = i + 1
	}

	return konflik
}

func tambahJadwal(A *TabJadwal, n *int, mk, dosen, ruangan, hari string, mulai, selesai int) {
	if *n < maksimal {
		if cekKonflikJadwal(*A, *n, hari, mulai, selesai) == true {
			fmt.Print("Data gagal ditambahkan karena terjadi konflik jadwal.\n")
		} else {
			A[*n].namamk = mk
			A[*n].dosen = dosen
			A[*n].ruang = ruangan
			A[*n].hari = hari
			A[*n].jamMulai = mulai
			A[*n].jamSelesai = selesai
			*n = *n + 1
			fmt.Print("\nData berhasil ditambahkan.\n")
		}
	} else {
		fmt.Print("\nKapasitas penyimpanan penuh.\n")
	}
}

func ubahJadwal(A *TabJadwal, n *int, mkLama string, hrLama string) {
	var found bool
	var i int 
	var MKbaru, DosenBaru, RuanganBaru, HariBaru string
	var jmBaru, jsBaru int
	
	i = 0
	found = false

	for i < *n && found == false {
		if A[i].namamk == mkLama && A[i].hari == hrLama {
			fmt.Print("Masukkan Nama MK Baru: ")
			fmt.Scan(&MKbaru)
			fmt.Print("Masukkan Dosen Baru: ")
			fmt.Scan(&DosenBaru)
			fmt.Print("Masukkan Ruangan Baru: ")
			fmt.Scan(&RuanganBaru)
			fmt.Print("Masukkan Hari Baru: ")
			fmt.Scan(&HariBaru)
			fmt.Print("Masukkan Jam Mulai Baru (HHMM): ")
			fmt.Scan(&jmBaru)
			fmt.Print("Masukkan Jam Selesai Baru (HHMM): ")
			fmt.Scan(&jsBaru)
			found = true
			if cekKonflikJadwal(*A, *n, HariBaru, jmBaru, jsBaru) == true {
			fmt.Print("Data gagal ditambahkan karena terjadi konflik jadwal.\n")
			} else {
			A[*n].namamk = MKbaru
			A[*n].dosen = DosenBaru
			A[*n].ruang = RuanganBaru
			A[*n].hari = HariBaru
			A[*n].jamMulai = jmBaru
			A[*n].jamSelesai = jsBaru
			fmt.Print("\nData berhasil ditambahkan.\n")
			}
		}
		i = i + 1
	}

	if found == false {
		fmt.Print("\nData mata kuliah tidak ditemukan.\n")
	}
}

func hapusJadwal(A *TabJadwal, n *int, mk string, hrLama string) {
	var found bool
	var idx int
	var i int

	found = false
	i = 0
	idx = -1
	for i < *n && found == false {
		if A[i].namamk == mk && A[i].hari == hrLama {
			idx = i
			found = true
		}
		i = i + 1
	}

	if found == true {
		var j int
		j = idx
		for j < *n-1 {
			A[j] = A[j+1]
			j = j + 1
		}
		*n = *n - 1
		fmt.Print("Data berhasil dihapus.\n")
	} else {
		fmt.Print("Data tidak ditemukan.\n")
	}
}

func cariSequentialMK(A TabJadwal, n int, mk string) {
	var i int
	var found bool
	
	found = false
	i = 0

	fmt.Print("\nHasil Pencarian Sequential (MK: ", mk, ")\n")
	for i < n {
		if A[i].namamk == mk {
			fmt.Printf("MK: %-10s | Dosen: %-7s | Ruang: %-10s | %-7s %d-%d\n", A[i].namamk, A[i].dosen, A[i].ruang, A[i].hari, A[i].jamMulai, A[i].jamSelesai)
			found = true
		}
		i = i + 1
	}

	if found == false {
		fmt.Print("Mata kuliah tidak ditemukan.\n")
	}
}

func urutBerdasarkanDosen(A *TabJadwal, n int) {
	var i, j int
	var temp Jadwal
	
	i = 1
	for i < n {
		temp = A[i]
		j = i
		for j > 0 && A[j-1].dosen > temp.dosen {
			A[j] = A[j-1]
			j = j - 1
		}
		A[j] = temp
		i = i + 1
	}
}

func cariBinaryDosen(A TabJadwal, n int, dosen string) {
	urutBerdasarkanDosen(&A, n)
	tampilkanSemua(A,n)
	var left, right, mid int
	var found bool 
	var idx int
	
	left = 0
	right = n - 1
	found = false
	idx = -1

	for left <= right && found == false {
		mid = (left + right) / 2
		if A[mid].dosen == dosen {
			found = true
			idx = mid
		} else if A[mid].dosen < dosen {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	fmt.Print("\n--- Hasil Pencarian Binary (Dosen: ", dosen, ") ---\n")
	if found == true {
		fmt.Printf("Ditemukan - MK: %-10s | Dosen: %-7s | Ruang: %-10s | %-7s %d-%d\n", A[idx].namamk, A[idx].dosen, A[idx].ruang, A[idx].hari, A[idx].jamMulai, A[idx].jamSelesai)
	} else {
		fmt.Print("Dosen tidak ditemukan.\n")
	}
}

func tampilkanStatistik(A TabJadwal, n int) {
	var hariDicari string
	var jamSekarang int
	
	fmt.Print("Masukkan Hari ini (Contoh: Senin): ")
	fmt.Scan(&hariDicari)
	fmt.Print("Masukkan Jam saat ini (Contoh: 1200): ")
	fmt.Scan(&jamSekarang)

	var totalMenit int
	var sisaKelas int
	var i int
	
	totalMenit = 0
	sisaKelas = 0
	i = 0

	for i < n {
		var jamM int = A[i].jamMulai / 100
		var mntM int = A[i].jamMulai % 100
		var jamS int = A[i].jamSelesai / 100
		var mntS int = A[i].jamSelesai % 100

		var durasi int = ((jamS * 60) + mntS) - ((jamM * 60) + mntM)
		totalMenit = totalMenit + durasi

		if A[i].hari == hariDicari && A[i].jamMulai >= jamSekarang {
			sisaKelas = sisaKelas + 1
		}
		i = i + 1
	}

	var totalJam int = totalMenit / 60
	var sisaMnt int = totalMenit % 60

	fmt.Print("\n--- Statistik Akademik ---\n")
	fmt.Print("1. Total jam kuliah minggu ini: ", totalJam, " Jam ", sisaMnt, " Menit\n")
	fmt.Print("2. Sisa kelas pada hari ", hariDicari, " setelah jam ", jamSekarang, " : ", sisaKelas, " kelas\n")
}

func tampilkanSemua(A TabJadwal, n int) {
	var i int
	
	i = 0
	
	fmt.Print("\nNo | Mata Kuliah |   Dosen   |   Ruang   |   Hari   | Jam\n")
	if n == 0 {
		fmt.Print("Belum ada data jadwal.\n")
	}
	for i < n {
		fmt.Printf("%d  | %-11s | %-9s | %-9s | %-8s | %d-%d\n", i+1, A[i].namamk, A[i].dosen, A[i].ruang, A[i].hari, A[i].jamMulai, A[i].jamSelesai)
		i = i + 1
	}
}

func bobotHari(hari string) int {
	if hari == "Senin" || hari == "senin" {
		return 1
	} else if hari == "Selasa" || hari == "selasa" {
		return 2
	} else if hari == "Rabu" || hari == "rabu" {
		return 3
	} else if hari == "Kamis" || hari == "kamis" {
		return 4
	} else if hari == "Jumat" || hari == "jumat" {
		return 5
	} else if hari == "Sabtu" || hari == "sabtu" {
		return 6
	} else if hari == "Minggu" || hari == "minggu" {
		return 7
	}
	return 8
}

func urutBerdasarkanHari(A *TabJadwal, n int, pilihan string) {
	var i, j int
	var temp Jadwal
	
	if pilihan == "a" {
		i = 1
		for i < n {
			temp = A[i]
			j = i
			
			for j > 0 && (bobotHari(A[j-1].hari) > bobotHari(temp.hari) || (bobotHari(A[j-1].hari) == bobotHari(temp.hari) && A[j-1].jamMulai > temp.jamMulai)) {
				A[j] = A[j-1]
				j = j - 1
			}
			A[j] = temp
			i = i + 1
		}
		fmt.Print("Data berhasil diurutkan Ascending (Insertion Sort).\n")
	} else if pilihan == "d" {
		i = 1
		for i < n {
			temp = A[i]
			j = i
			
			for j > 0 && (bobotHari(A[j-1].hari) < bobotHari(temp.hari) || (bobotHari(A[j-1].hari) == bobotHari(temp.hari) && A[j-1].jamMulai < temp.jamMulai)) {
				A[j] = A[j-1]
				j = j - 1
			}
			A[j] = temp
			i = i + 1
		}
		fmt.Print("Data berhasil diurutkan Descending (Insertion Sort).\n")
	}
}

func urutBerdasarkanHariSelection(A *TabJadwal, n int, pilihan string) {
	var i, j, idxTarget int
	var temp Jadwal

	if pilihan == "a" {
		i = 0
		for i < n-1 {
			idxTarget = i
			j = i + 1
			for j < n {
				if bobotHari(A[j].hari) < bobotHari(A[idxTarget].hari) || (bobotHari(A[j].hari) == bobotHari(A[idxTarget].hari) && A[j].jamMulai < A[idxTarget].jamMulai) {
					idxTarget = j
				}
				j = j + 1
			}
			temp = A[i]
			A[i] = A[idxTarget]
			A[idxTarget] = temp

			i = i + 1
		}
		fmt.Print("Data berhasil diurutkan Ascending (Selection Sort).\n")
	
	} else if pilihan == "d" {
		i = 0
		for i < n-1 {
			idxTarget = i
			j = i + 1
			for j < n {
				if bobotHari(A[j].hari) > bobotHari(A[idxTarget].hari) || (bobotHari(A[j].hari) == bobotHari(A[idxTarget].hari) && A[j].jamMulai > A[idxTarget].jamMulai) {
					idxTarget = j
				}
				j = j + 1
			}
			temp = A[i]
			A[i] = A[idxTarget]
			A[idxTarget] = temp
			
			i = i + 1
		}
		fmt.Print("Data berhasil diurutkan Descending (Selection Sort).\n")
	}
}

func main() {
	var data TabJadwal
	var nData int
	var pilihan int
	var berjalan bool
	var mk, dos, ru, hr, mkLama, mkHapus, dosCari, mkCari, pilihanUrut string
	var jm, js int
	
	nData = 0
	berjalan = true
	
	tambahJadwal(&data, &nData, "Alpro2", "Lds", "E301", "Selasa", 830, 1030)
	tambahJadwal(&data, &nData, "Matvek", "Mahmud", "TULT0714", "Senin", 1030, 1330)
	tambahJadwal(&data, &nData, "Alpro2", "Lds", "KU3.03.03", "Rabu", 1330, 1530)
	tambahJadwal(&data, &nData, "Kalkulus", "Adit", "TULT0707", "Senin", 1330, 1530)
	tambahJadwal(&data, &nData, "PBD", "Cahyo", "TULT0707", "Rabu", 1530, 1830)
	tambahJadwal(&data, &nData, "Etika_AI", "Gamma", "TULT0714", "Selasa", 1230, 1430)
	tambahJadwal(&data, &nData, "COA", "Ghifari", "KU1.03.14", "Rabu", 830, 1130)
	
	for berjalan == true {
		fmt.Print("\n==================================\n")
		fmt.Print(" SISTEM MANAJEMEN JADWAL (JadwalKu)\n")
		fmt.Print("==================================\n")
		fmt.Print("1. Tambah Jadwal\n")
		fmt.Print("2. Ubah Jadwal\n")
		fmt.Print("3. Hapus Jadwal\n")
		fmt.Print("4. Tampilkan Semua Jadwal\n")
		fmt.Print("5. Urutkan Jadwal (Hari & Jam - Insertion Sort)\n")
		fmt.Print("6. Urutkan Jadwal (Hari & Jam - Selection Sort)\n")
		fmt.Print("7. Cari berdasarkan Dosen (Binary)\n")
		fmt.Print("8. Cari berdasarkan MK (Sequential)\n")
		fmt.Print("9. Lihat Statistik\n")
		fmt.Print("0. Keluar\n")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 0 {
			berjalan = false
			fmt.Print("Terima kasih.\n")
		} else if pilihan == 1 {
			fmt.Print("Mata Kuliah: ")
			fmt.Scan(&mk)
			fmt.Print("Dosen: ")
			fmt.Scan(&dos)
			fmt.Print("Ruangan: ")
			fmt.Scan(&ru)
			fmt.Print("Hari: ")
			fmt.Scan(&hr)
			fmt.Print("Jam Mulai (HHMM): ")
			fmt.Scan(&jm)
			fmt.Print("Jam Selesai (HHMM): ")
			fmt.Scan(&js)
			tambahJadwal(&data, &nData, mk, dos, ru, hr, jm, js)
		} else if pilihan == 2 {
			fmt.Print("Masukkan Nama MK yang ingin diubah: ")
			fmt.Scan(&mkLama)
			fmt.Print("Masukkan Hari Kelas Matkul yang ingin diubah: ")
			fmt.Scan(&hr)
			ubahJadwal(&data, &nData, mkLama, hr)
		} else if pilihan == 3 {
			fmt.Print("Masukkan Nama MK yang ingin dihapus: ")
			fmt.Scan(&mkHapus)
			fmt.Print("Masukkan Hari Kelas Matkul yang ingin dihapus: ")
			fmt.Scan(&hr)
			hapusJadwal(&data, &nData, mkHapus, hr)
		} else if pilihan == 4 {
			tampilkanSemua(data, nData)
		} else if pilihan == 5 {
			fmt.Print("Pilih urutan (a/ascending, d/descending): ")
			fmt.Scan(&pilihanUrut)
			urutBerdasarkanHari(&data, nData, pilihanUrut)
		} else if pilihan == 6 {
			fmt.Print("Pilih urutan (a/ascending, d/descending): ")
			fmt.Scan(&pilihanUrut)
			urutBerdasarkanHariSelection(&data, nData, pilihanUrut)
		} else if pilihan == 7 {
			fmt.Print("Masukkan Nama Dosen: ")
			fmt.Scan(&dosCari)
			cariBinaryDosen(data, nData, dosCari) 
		} else if pilihan == 8 {
			fmt.Print("Masukkan Nama MK: ")
			fmt.Scan(&mkCari)
			cariSequentialMK(data, nData, mkCari)
		} else if pilihan == 9 {
			tampilkanStatistik(data, nData)
		} else {
			fmt.Print("Pilihan tidak valid.\n")
		}
	}
}
