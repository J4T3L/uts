package main

import "fmt"

type LinkList struct {
	container Member
	next      *LinkList
}
type Link struct {
	contain Buku
	next    *Link
}
type Member struct {
	Nomor  int
	Nama   string
	Alamat string
	No_Hp  string
	buku   string
}
type Buku struct {
	Nomor    int
	NamaB    string
	Penulis  string
	Penerbit string
}

// FUNCTION MENU CLI
func menu(l string) {
	fmt.Println("menu pilihan manipulasi data ", l, " :")
	fmt.Println("1. Insert")
	fmt.Println("2. Update")
	fmt.Println("3. Delete")
	fmt.Println("4. View All")
	fmt.Println("5. BACK <-")
}

// FUNCTION TAMBAH DATA PENGUNJUNG
func insertP(linkedList *LinkList, member Member) {
	//function tambah data pengunjung
	tempList := LinkList{}
	tempList.container = member

	temp := linkedList
	if temp.next == nil {
		temp.next = &tempList
	} else {
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = &tempList
	}
}

// FUNCTION TAMBAH DATA BUKU
func insertB(linkedList *Link, buku Buku) {
	//function tambah data BUku
	tempList := Link{}
	tempList.contain = buku

	temp := linkedList
	if temp.next == nil {
		temp.next = &tempList
	} else {
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = &tempList
	}
}

// FUNCTION TAMPIL DATA PENGUNJUNG
func ViewP(linkedlist *LinkList) {
	temp := *linkedlist
	fmt.Println("________________________________________________________")
	fmt.Println("| NO  |		Nama     |	Alamat	|	NO_hp	|	BUKU ")
	fmt.Println("===============================================================")
	for temp.next != nil {

		fmt.Println(
			"|", temp.next.container.Nomor, ".",
			"|", temp.next.container.Nama,
			"		 |", temp.next.container.Alamat,
			"	|", temp.next.container.No_Hp, "	|",
			"|	", temp.next.container.buku, "	|",
			"\n==================================================================",
		)
		temp.next = temp.next.next

	}
}

// FUNCTION TAMPIL DATA BUKU
func ViewB(link *Link) {
	temp := *link
	fmt.Println("________________________________________________________")
	fmt.Println("| NO  |		Nama     |	Penulis	|	Penerbit	|")
	fmt.Println("=========================================================")
	for temp.next != nil {

		fmt.Println(
			"|", temp.next.contain.Nomor, ".",
			"|", temp.next.contain.NamaB,
			"		 |", temp.next.contain.Penulis,
			"	|", temp.next.contain.Penerbit, "	|",
			"\n=========================================================",
		)
		temp.next = temp.next.next

	}
}

// FUNCTION UPDATE DATA PENGUNJUNG
func UpdateP(list *LinkList, no int, member Member) {
	temp := list
	temp.container = member
	var fisrt int = 0
	for temp.next != nil {
		fisrt++
		if temp.next.container.Nomor == no {
			temp.next.container = member
		}
		temp = temp.next
	}
}

// FUNCTION UPDATE DATA BUKU
func UpdateB(list *Link, no int, member Buku) {
	temp := list
	temp.contain = member
	var fisrt int = 0
	for temp.next != nil {
		fisrt++
		if temp.next.contain.Nomor == no {
			temp.next.contain = member
		}
		temp = temp.next
	}
}

// FUNCTION HAPUS DATA PENGUNJUNG
func DeleteP(list *LinkList, no int) {
	temp := list
	var fisrt int = 0
	for temp.next != nil {
		fisrt++
		if temp.next.container.Nomor == no {
			if fisrt == 1 {
				list.next = temp.next.next
				break
			} else if list.next != nil {
				temp.next = temp.next.next
				break
			} else {
				temp.next = nil
			}
		}
		temp = temp.next
	}
}

// FUNCTION HAPUS DATA BUKU
func DeleteB(list *Link, NO int) {
	temp := list
	var fisrt int = 0
	for temp.next != nil {
		fisrt++
		if temp.next.contain.Nomor == NO {
			if fisrt == 1 {
				list.next = temp.next.next
				break
			} else if list.next != nil {
				temp.next = temp.next.next
				break
			} else {
				temp.next = nil
			}
		}
		temp = temp.next
	}
}

func main() {
	dataPerpus := LinkList{}
	dataBuku := Link{}
	var a int
	var menunya int = 0
	for menunya != 6 && a != 3 {
		fmt.Println("Insert   : \n 1.Data Pengunjung \n 2.Data Buku \n 3.EXIT")
		fmt.Print("data yang di proses : ")
		fmt.Scan(&a)

		var ret int = 0
		//ret = menuAwal()
		fmt.Println("ret : ", ret)
		if a == 1 {
			var P string = "Pengunjung"
			var menunya int = 0
			var NO int = 0
			for menunya != 5 {
				menu(P)

				fmt.Print("masukkan pilihan menu :")
				fmt.Scan(&menunya)
				switch menunya {
				case 1:
					fmt.Println("anda masuk pilihan 1: INSERT")
					var nama, alamat, pil, book string
					var no_hp string
					NO++
					fmt.Print("masukkan nama  pengunjung : ")
					fmt.Scan(&nama)
					fmt.Print("masukkan Alamat pengunjung : ")
					fmt.Scan(&alamat)
					fmt.Print("masukkan NO.Hp pengunjung : ")
					fmt.Scan(&no_hp)
					fmt.Printf("apakah Pengunjung meminjam buku (y/n)?")
					fmt.Scan(&pil)
					if pil == "y" || pil == "Y" {

						if dataBuku.next == nil {
							fmt.Printf("DATA KOSONG<HARAP INPUTKAN DATA DAHULU!!")
							break
						} else {
							ViewB(&dataBuku)

							fmt.Printf("Masukkan nama buku yang dipinjam : ")
							fmt.Scan(&book)
							if book == dataBuku.contain.NamaB {
								book = dataBuku.contain.NamaB

							}
							fmt.Printf("Tanggal pinjam ")
						}
					} else if pil == "n" || pil == "N" {
						book = "-"
					}
					data := Member{
						Nomor:  NO,
						Nama:   nama,
						Alamat: alamat,
						No_Hp:  no_hp,
						buku:   book,
					}
					insertP(&dataPerpus, data)
					fmt.Println("Data BErhasil di tambah !!")
				case 2:
					var no_upadate int
					fmt.Println("anda masuk pilihan 2:UPDATE :")
					var nama, alamat, pil, book string
					var no_hp string
					NO++
					fmt.Print("masukkan nama  pengunjung : ")
					fmt.Scan(&nama)
					fmt.Print("masukkan Alamat pengunjung : ")
					fmt.Scan(&alamat)
					fmt.Print("masukkan NO.Hp pengunjung : ")
					fmt.Scan(&no_hp)
					fmt.Printf("apakah Pengunjung meminjam buku (y/n)?")
					fmt.Scan(&pil)
					if pil == "y" || pil == "Y" {

						if dataBuku.next == nil {
							fmt.Println("DATA KOSONG<HARAP INPUTKAN DATA DAHULU!!")
							book = "-"
						} else {
							ViewB(&dataBuku)

							fmt.Printf("Masukkan nama buku yang dipinjam : ")
							fmt.Scan(&book)
							if book == dataBuku.contain.NamaB {
								book = dataBuku.contain.NamaB
							}
						}
					} else if pil == "n" || pil == "N" {
						book = "-"
					}
					dataNew := Member{
						Nomor:  NO,
						Nama:   nama,
						Alamat: alamat,
						No_Hp:  no_hp,
						buku:   book,
					}
					fmt.Scan(&no_upadate)
					UpdateP(&dataPerpus, no_upadate, dataNew)
				case 3:
					var no_delete int
					NO--
					fmt.Println("anda masuk pilihan 3:DELETE")
					fmt.Println("List Data Pengunjung :")
					ViewP(&dataPerpus)
					fmt.Println("nomor data yang di hapus:")
					fmt.Scan(&no_delete)
					DeleteP(&dataPerpus, no_delete)
					fmt.Println("Data nomor ", no_delete, "Dihapus!")
				case 4:
					fmt.Println("anda masuk pilihan 4. VIEW ALL")
					ViewP(&dataPerpus)
				}
			}

		} else if a == 2 {
			var menunya int = 0
			B := "Buku"
			var NO int = 0
			for menunya != 5 {
				menu(B)

				fmt.Print("masukkan pilihan menu :")
				fmt.Scan(&menunya)
				switch menunya {
				case 1:
					fmt.Println("anda masuk pilihan 1: INSERT")
					var namaB, penulis, penerbit string
					NO++
					fmt.Print("masukkan nama  Buku : ")
					fmt.Scan(&namaB)
					fmt.Print("masukkan Nama Penulis : ")
					fmt.Scan(&penulis)
					fmt.Print("masukkan Nama Penerbit : ")
					fmt.Scan(&penerbit)
					buku := Buku{
						Nomor:    NO,
						NamaB:    namaB,
						Penulis:  penulis,
						Penerbit: penerbit,
					}
					insertB(&dataBuku, buku)
					fmt.Println("Data BErhasil di tambah !!")
				case 2:
					var no_upadate int
					fmt.Println("anda masuk pilihan 2:UPDATE :")
					var namaB, penulis, penerbit string
					NO++
					fmt.Print("masukkan nama  Buku : ")
					fmt.Scan(&namaB)
					fmt.Print("masukkan Nama Penulis : ")
					fmt.Scan(&penulis)
					fmt.Print("masukkan Nama Penerbit : ")
					fmt.Scan(&penerbit)
					bukuNew := Buku{
						Nomor:    NO,
						NamaB:    namaB,
						Penulis:  penulis,
						Penerbit: penerbit,
					}
					fmt.Scan(&no_upadate)
					UpdateB(&dataBuku, no_upadate, bukuNew)
				case 3:
					var no_delete int
					fmt.Println("anda masuk pilihan 3:DELETE")
					fmt.Println("List Data Pengunjung :")
					ViewB(&dataBuku)
					fmt.Println("nomor data yang di hapus:")
					fmt.Scan(&no_delete)
					DeleteB(&dataBuku, no_delete)
					fmt.Println("Data nomor ", no_delete, "Dihapus!")
				case 4:
					fmt.Println("anda masuk pilihan 4. VIEW ALL")
					ViewB(&dataBuku)
				}
			}
		}

	}

	fmt.Println("\nKELUAR DARI PROGRAM!!")
}
