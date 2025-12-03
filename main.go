package main

import "fmt"

type Buku struct {
	ID      int
	Judul   string
	Penulis string
	Harga   float64
	Genre   string
}

var daftarBuku = []Buku{
	{1, "Seporsi Mie Ayam Sebelum Mati", "Brian Khrisna", 89000, "Fiksi"},
	{2, "0 MDPL", "Nurwina Sari", 95000, "Travel/Romance"},
	{3, "Dilan 1991", "Pidi Baiq", 79000, "Teen Romance"},
	{4, "Goresan Seorang Berandal", "Mohan Hazian", 85000, "Biografi/Memoar"},
	{5, "Laut Bercerita", "Leila S. Chudori", 110000, "Historical Fiction"},
	{6, "Slow Living", "Sabrina Ara", 75000, "Self Help"},
	{7, "Kios Pasar Sore", "Reda Gaudiamo", 68000, "Fiksi"},
	{8, "Aku Titip Dia, Ya!", "Rafi Ibadi", 82000, "Romance"},
	{9, "Untuk Satu Nama", "Rafi Ibadi", 88000, "Romance"},
	{10, "Nanti Juga Sembuh Sendiri", "HeloBagas", 90000, "Self Help"},
	{11, "Jika Lukamu Sedalam Laut", "patahan.ranting", 92000, "Motivasi"},
	{12, "Tentang Luka", "kopioppi", 70000, "Puisi"},
	{13, "Yang Katanya Cemara", "Vania Winola", 95000, "Family"},
	{14, "Hai Nak!", "Reda Gaudiamo", 65000, "Parenting"},
	{15, "Kamu Tak Harus Sempurna", "Anastasia Satriyo", 99000, "Psychology"},
}

var keranjang []Buku

func main() {
	var pilihan int
	isRunning := true

	for isRunning {
		fmt.Println("\n========================================")
		fmt.Println("    BOOKAVY CLI - SYSTEM BACKEND        ")
		fmt.Println("========================================")
		fmt.Println("1. Featured Book")
		fmt.Println("2. Lihat Semua Koleksi")
		fmt.Println("3. Cek Keranjang & Checkout")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu [0-3]: ")

		_, err := fmt.Scan(&pilihan)
		if err != nil {
			var discard string
			fmt.Scan(&discard)
			fmt.Println("Error: Masukkan angka saja.")
			continue
		}

		switch pilihan {
		case 1:
			tampilkanFeatured()
		case 2:
			tampilkanKatalog()
		case 3:
			lihatKeranjang()
		case 0:
			fmt.Println("Terima kasih telah berkunjung.")
			isRunning = false
		default:
			fmt.Println("Pilihan tidak tersedia.")
		}
	}
}

func tampilkanFeatured() {
	b := daftarBuku[0]
	fmt.Println("\n--- FEATURED BOOK OF THE MONTH ---")
	fmt.Printf("Judul   : %s\n", b.Judul)
	fmt.Printf("Penulis : %s\n", b.Penulis)
	fmt.Printf("Harga   : Rp %.0f\n", b.Harga)
	fmt.Printf("Genre   : %s\n", b.Genre)
	fmt.Println("Sinopsis: Kisah tentang rasa lapar, kehilangan, dan mie ayam.")

	fmt.Print("\n[1] Beli Buku Ini  [0] Kembali: ")
	var aksi int
	fmt.Scan(&aksi)
	if aksi == 1 {
		keranjang = append(keranjang, b)
		fmt.Println("Berhasil masuk keranjang.")
	}
}

func tampilkanKatalog() {
	fmt.Println("\n--- DAFTAR KOLEKSI BUKU ---")
	for i, b := range daftarBuku {
		fmt.Printf("%d. %s (Oleh: %s) - Rp %.0f\n", i+1, b.Judul, b.Penulis, b.Harga)
	}

	var id int
	fmt.Print("\nMasukkan Nomor Buku untuk dibeli (0 untuk batal): ")
	_, err := fmt.Scan(&id)
	if err != nil {
		var discard string
		fmt.Scan(&discard)
		fmt.Println("Input harus angka.")
		return
	}

	if id > 0 && id <= len(daftarBuku) {
		bukuTerpilih := daftarBuku[id-1]
		keranjang = append(keranjang, bukuTerpilih)
		fmt.Printf("Buku '%s' berhasil ditambahkan ke keranjang.\n", bukuTerpilih.Judul)
	} else if id != 0 {
		fmt.Println("Nomor buku tidak valid.")
	}
}

func lihatKeranjang() {
	fmt.Println("\n--- KERANJANG BELANJA ANDA ---")
	if len(keranjang) == 0 {
		fmt.Println("(Keranjang masih kosong)")
	} else {
		total := 0.0
		for i, b := range keranjang {
			fmt.Printf("%d. %s - Rp %.0f\n", i+1, b.Judul, b.Harga)
			total += b.Harga
		}
		fmt.Println("----------------------------------------")
		fmt.Printf("TOTAL TAGIHAN: Rp %.0f\n", total)
		fmt.Println("----------------------------------------")

		var aksi int
		fmt.Println("\n[1] Checkout (Bayar)")
		fmt.Println("[0] Kembali ke Menu Utama")
		fmt.Print("Pilih aksi: ")
		_, err := fmt.Scan(&aksi)

		if err == nil && aksi == 1 {
			prosesCheckout(total)
		}
	}
}

func prosesCheckout(total float64) {
	fmt.Println("\n--- PROSES PEMBAYARAN ---")
	fmt.Printf("Total yang harus dibayar: Rp %.0f\n", total)

	var bayar float64
	for {
		fmt.Print("Masukkan jumlah uang: Rp ")
		_, err := fmt.Scan(&bayar)

		if err != nil {
			var discard string
			fmt.Scan(&discard)
			fmt.Println("Masukkan nominal angka.")
			continue
		}

		if bayar >= total {
			kembalian := bayar - total
			fmt.Printf("Pembayaran Berhasil. Kembalian: Rp %.0f\n", kembalian)
			fmt.Println("Buku akan segera dikirim. Terima kasih!")

			keranjang = nil
			break
		} else {
			fmt.Printf("Uang kurang sebesar Rp %.0f. Silakan input ulang.\n", total-bayar)
		}
	}
}
