package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Mahasiswa struct {
	NIM    string
	Name   string
	Course string
	Origin string
}

type Buku struct {
	Author              string
	Judul               string
	Tahun               int
	TanggalPeminjaman   string
	TanggalPengembalian string
}

type Pegawai struct {
	NIP     string
	Nama    string
	Jabatan string
}

var studentData []Mahasiswa
var bookData []Buku
var employeeData []Pegawai

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n╔════════════════════════════╗")
		fmt.Println("║         MAIN MENU          ║")
		fmt.Println("╠════════════════════════════╣")
		fmt.Println("║ 1. Add Data                ║")
		fmt.Println("║ 2. View Data               ║")
		fmt.Println("║ 3. Update Data             ║")
		fmt.Println("║ 4. Delete Data             ║")
		fmt.Println("║ 5. Search Data             ║")
		fmt.Println("║ 6. Exit                    ║")
		fmt.Println("╚════════════════════════════╝")
		fmt.Print("Choose menu: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			addData(scanner)
		case "2":
			viewData()
		case "3":
			updateData(scanner)
		case "4":
			deleteData(scanner)
		case "5":
			searchData(scanner)
		case "6":
			fmt.Println("Exiting program.")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func addData(scanner *bufio.Scanner) {
	for {
		fmt.Println("\n╔════════════════════════════╗")
		fmt.Println("║         ADD DATA           ║")
		fmt.Println("╠════════════════════════════╣")
		fmt.Println("║ 1. Mahasiswa               ║")
		fmt.Println("║ 2. Buku                    ║")
		fmt.Println("║ 3. Pegawai                 ║")
		fmt.Println("║ 4. Back                    ║")
		fmt.Println("╚════════════════════════════╝")
		fmt.Print("Choose: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			var m Mahasiswa
			fmt.Println("\n╔════════════════════════════╗")
			fmt.Println("║     ADD MAHASISWA DATA     ║")
			fmt.Println("╚════════════════════════════╝")
			for {
				fmt.Print("NIM: ")
				scanner.Scan()
				m.NIM = scanner.Text()
				if isNumeric(m.NIM) {
					break
				}
				fmt.Println("Error: NIM must be numeric.")
			}
			fmt.Print("Name: ")
			scanner.Scan()
			m.Name = scanner.Text()
			fmt.Print("Course: ")
			scanner.Scan()
			m.Course = scanner.Text()
			fmt.Print("Place of Origin (Country): ")
			scanner.Scan()
			m.Origin = scanner.Text()

			studentData = append(studentData, m)
			fmt.Println("\n✅ Mahasiswa data successfully added.")
		case "2":
			var b Buku
			fmt.Println("\n╔════════════════════════════╗")
			fmt.Println("║        ADD BUKU DATA       ║")
			fmt.Println("╚════════════════════════════╝")
			fmt.Print("Author: ")
			scanner.Scan()
			b.Author = scanner.Text()
			fmt.Print("Judul: ")
			scanner.Scan()
			b.Judul = scanner.Text()
			fmt.Print("Tahun: ")
			scanner.Scan()
			fmt.Sscanf(scanner.Text(), "%d", &b.Tahun)
			fmt.Print("Tanggal Peminjaman: ")
			scanner.Scan()
			b.TanggalPeminjaman = scanner.Text()
			fmt.Print("Tanggal Pengembalian: ")
			scanner.Scan()
			b.TanggalPengembalian = scanner.Text()

			bookData = append(bookData, b)
			fmt.Println("\n✅ Buku data successfully added.")
		case "3":
			var p Pegawai
			fmt.Println("\n╔════════════════════════════╗")
			fmt.Println("║      ADD PEGAWAI DATA      ║")
			fmt.Println("╚════════════════════════════╝")
			fmt.Print("NIP: ")
			scanner.Scan()
			p.NIP = scanner.Text()
			fmt.Print("Nama: ")
			scanner.Scan()
			p.Nama = scanner.Text()
			fmt.Print("Jabatan: ")
			scanner.Scan()
			p.Jabatan = scanner.Text()

			employeeData = append(employeeData, p)
			fmt.Println("\n✅ Pegawai data successfully added.")
		case "4":
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func isNumeric(input string) bool {
	match, _ := regexp.MatchString(`^\d+$`, input)
	return match
}

func viewData() {
	for {
		fmt.Println("\n╔════════════════════════════╗")
		fmt.Println("║        VIEW DATA MENU      ║")
		fmt.Println("╠════════════════════════════╣")
		fmt.Println("║ 1. Mahasiswa               ║")
		fmt.Println("║ 2. Buku                    ║")
		fmt.Println("║ 3. Pegawai                 ║")
		fmt.Println("║ 4. Back                    ║")
		fmt.Println("╚════════════════════════════╝")
		fmt.Print("Choose: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			if len(studentData) == 0 {
				fmt.Println("No Mahasiswa data available.")
				continue
			}
			for {
				fmt.Println("\n╔════════════════════════════╗")
				fmt.Println("║      SORT MAHASISWA        ║")
				fmt.Println("╠════════════════════════════╣")
				fmt.Println("║ 1. Ascending (Selection)   ║")
				fmt.Println("║ 2. Descending (Insertion)  ║")
				fmt.Println("║ 3. Back                    ║")
				fmt.Println("╚════════════════════════════╝")
				fmt.Print("Choose: ")
				scanner.Scan()
				sortChoice := scanner.Text()
				if sortChoice == "3" {
					break
				} else if sortChoice == "1" {
					selectionSortMahasiswa(studentData, true)
				} else if sortChoice == "2" {
					insertionSortMahasiswa(studentData, false)
				} else {
					fmt.Println("Invalid sort choice.")
					continue
				}

				
				fmt.Println("\n╔═══════╦═══════════════╦═══════════════════════╦═══════════════════╦═══════════════════╗")
				fmt.Println("║  NO   ║      NIM       ║         NAME          ║      COURSE       ║      ORIGIN       ║")
				fmt.Println("╠═══════╬═══════════════╬═══════════════════════╬═══════════════════╬═══════════════════╣")

				
				for i, m := range studentData {
					fmt.Printf("║ %-5d ║ %-13s ║ %-21s ║ %-17s ║ %-17s ║\n",
						i+1, m.NIM, truncateString(m.Name, 21), truncateString(m.Course, 17), truncateString(m.Origin, 17))
				}

				
				fmt.Println("╚═══════╩═══════════════╩═══════════════════════╩═══════════════════╩═══════════════════╝")
				break
			}
		case "2":
			if len(bookData) == 0 {
				fmt.Println("No Buku data available.")
				continue
			}
			for {
				fmt.Println("\n╔════════════════════════════╗")
				fmt.Println("║         SORT BUKU          ║")
				fmt.Println("╠════════════════════════════╣")
				fmt.Println("║ 1. Ascending (Selection)   ║")
				fmt.Println("║ 2. Descending (Insertion)  ║")
				fmt.Println("║ 3. Back                    ║")
				fmt.Println("╚════════════════════════════╝")
				fmt.Print("Choose: ")
				scanner.Scan()
				sortChoice := scanner.Text()
				if sortChoice == "3" {
					break
				} else if sortChoice == "1" {
					selectionSortBuku(bookData, true)
				} else if sortChoice == "2" {
					insertionSortBuku(bookData, false)
				} else {
					fmt.Println("Invalid sort choice.")
					continue
				}

				
				fmt.Println("\n╔═══════╦═══════════════════╦═══════════════════════╦═══════╦═══════════════════╦═══════════════════╗")
				fmt.Println("║  NO   ║       AUTHOR       ║         JUDUL         ║ TAHUN ║  TGL PEMINJAMAN   ║ TGL PENGEMBALIAN  ║")
				fmt.Println("╠═══════╬═══════════════════╬═══════════════════════╬═══════╬═══════════════════╬═══════════════════╣")

			
				for i, b := range bookData {
					fmt.Printf("║ %-5d ║ %-17s ║ %-21s ║ %-5d ║ %-17s ║ %-17s ║\n",
						i+1, truncateString(b.Author, 17), truncateString(b.Judul, 21),
						b.Tahun, truncateString(b.TanggalPeminjaman, 17), truncateString(b.TanggalPengembalian, 17))
				}

				fmt.Println("╚═══════╩═══════════════════╩═══════════════════════╩═══════╩═══════════════════╩═══════════════════╝")
				break
			}
		case "3":
			if len(employeeData) == 0 {
				fmt.Println("No Pegawai data available.")
				continue
			}
			for {
				fmt.Println("\n╔════════════════════════════╗")
				fmt.Println("║        SORT PEGAWAI        ║")
				fmt.Println("╠════════════════════════════╣")
				fmt.Println("║ 1. Ascending (Selection)   ║")
				fmt.Println("║ 2. Descending (Insertion)  ║")
				fmt.Println("║ 3. Back                    ║")
				fmt.Println("╚════════════════════════════╝")
				fmt.Print("Choose: ")
				scanner.Scan()
				sortChoice := scanner.Text()
				if sortChoice == "3" {
					break
				} else if sortChoice == "1" {
					selectionSortPegawai(employeeData, true)
				} else if sortChoice == "2" {
					insertionSortPegawai(employeeData, false)
				} else {
					fmt.Println("Invalid sort choice.")
					continue
				}

			
				fmt.Println("\n╔═══════╦═══════════════════╦═══════════════════════╦═══════════════════╗")
				fmt.Println("║  NO   ║        NIP        ║         NAMA          ║      JABATAN      ║")
				fmt.Println("╠═══════╬═══════════════════╬═══════════════════════╬═══════════════════╣")

			
				for i, p := range employeeData {
					fmt.Printf("║ %-5d ║ %-17s ║ %-21s ║ %-17s ║\n",
						i+1, p.NIP, truncateString(p.Nama, 21), truncateString(p.Jabatan, 17))
				}

				fmt.Println("╚═══════╩═══════════════════╩═══════════════════════╩═══════════════════╝")
				break
			}
		case "4":
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func selectionSortMahasiswa(data []Mahasiswa, ascending bool) {
	for i := 0; i < len(data)-1; i++ {
		idx := i
		for j := i + 1; j < len(data); j++ {
			if ascending && data[j].NIM < data[idx].NIM {
				idx = j
			}
		}
		data[i], data[idx] = data[idx], data[i]
	}
}

func insertionSortMahasiswa(data []Mahasiswa, descending bool) {
	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && descending && data[j].NIM < key.NIM {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

func selectionSortBuku(data []Buku, ascending bool) {
	for i := 0; i < len(data)-1; i++ {
		idx := i
		for j := i + 1; j < len(data); j++ {
			if ascending && data[j].Judul < data[idx].Judul {
				idx = j
			}
		}
		data[i], data[idx] = data[idx], data[i]
	}
}

func insertionSortBuku(data []Buku, descending bool) {
	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && descending && data[j].Judul < key.Judul {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

func selectionSortPegawai(data []Pegawai, ascending bool) {
	for i := 0; i < len(data)-1; i++ {
		idx := i
		for j := i + 1; j < len(data); j++ {
			if ascending && data[j].NIP < data[idx].NIP {
				idx = j
			}
		}
		data[i], data[idx] = data[idx], data[i]
	}
}

func insertionSortPegawai(data []Pegawai, descending bool) {
	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && descending && data[j].NIP < key.NIP {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

func updateData(scanner *bufio.Scanner) {
	for {
		fmt.Println("\n╔════════════════════════════╗")
		fmt.Println("║        UPDATE DATA         ║")
		fmt.Println("╠════════════════════════════╣")
		fmt.Println("║ 1. Mahasiswa               ║")
		fmt.Println("║ 2. Buku                    ║")
		fmt.Println("║ 3. Pegawai                 ║")
		fmt.Println("║ 4. Back                    ║")
		fmt.Println("╚════════════════════════════╝")
		fmt.Print("Choose: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Println("\n╔════════════════════════════╗")
			fmt.Println("║    UPDATE MAHASISWA DATA   ║")
			fmt.Println("╚════════════════════════════╝")
			fmt.Print("Enter NIM of the Mahasiswa to update: ")
			scanner.Scan()
			nim := scanner.Text()
			for i, m := range studentData {
				if m.NIM == nim {
					
					fmt.Println("\n📋 Current Data:")
					printStudent(m)

					fmt.Println("\n╔════════════════════════════╗")
					fmt.Println("║      FIELD TO UPDATE       ║")
					fmt.Println("╠════════════════════════════╣")
					fmt.Println("║ 1. Name                    ║")
					fmt.Println("║ 2. NIM                     ║")
					fmt.Println("║ 3. Course                  ║")
					fmt.Println("║ 4. Origin                  ║")
					fmt.Println("╚════════════════════════════╝")
					fmt.Print("Choose: ")
					scanner.Scan()
					field := scanner.Text()
					switch field {
					case "1":
						fmt.Print("Enter new Name: ")
						scanner.Scan()
						studentData[i].Name = scanner.Text()
					case "2":
						fmt.Print("Enter new NIM: ")
						scanner.Scan()
						studentData[i].NIM = scanner.Text()
					case "3":
						fmt.Print("Enter new Course: ")
						scanner.Scan()
						studentData[i].Course = scanner.Text()
					case "4":
						fmt.Print("Enter new Origin: ")
						scanner.Scan()
						studentData[i].Origin = scanner.Text()
					default:
						fmt.Println("Invalid choice.")
					}
					fmt.Println("\n✅ Mahasiswa data successfully updated.")
					return
				}
			}
			fmt.Println("\n❌ Mahasiswa with the specified NIM not found.")
		case "2":
			fmt.Println("\n╔════════════════════════════╗")
			fmt.Println("║      UPDATE BUKU DATA      ║")
			fmt.Println("╚════════════════════════════╝")
			fmt.Print("Enter Judul of the Buku to update: ")
			scanner.Scan()
			judul := scanner.Text()
			for i, b := range bookData {
				if strings.ToLower(b.Judul) == strings.ToLower(judul) {
					
					fmt.Println("\n📋 Current Data:")
					printBuku(b)

					fmt.Println("\n╔════════════════════════════╗")
					fmt.Println("║      FIELD TO UPDATE       ║")
					fmt.Println("╠════════════════════════════╣")
					fmt.Println("║ 1. Author                  ║")
					fmt.Println("║ 2. Judul                   ║")
					fmt.Println("║ 3. Tahun                   ║")
					fmt.Println("║ 4. Tanggal Peminjaman      ║")
					fmt.Println("║ 5. Tanggal Pengembalian    ║")
					fmt.Println("╚════════════════════════════╝")
					fmt.Print("Choose: ")
					scanner.Scan()
					field := scanner.Text()
					switch field {
					case "1":
						fmt.Print("Enter new Author: ")
						scanner.Scan()
						bookData[i].Author = scanner.Text()
					case "2":
						fmt.Print("Enter new Judul: ")
						scanner.Scan()
						bookData[i].Judul = scanner.Text()
					case "3":
						fmt.Print("Enter new Tahun: ")
						scanner.Scan()
						fmt.Sscanf(scanner.Text(), "%d", &bookData[i].Tahun)
					case "4":
						fmt.Print("Enter new Tanggal Peminjaman: ")
						scanner.Scan()
						bookData[i].TanggalPeminjaman = scanner.Text()
					case "5":
						fmt.Print("Enter new Tanggal Pengembalian: ")
						scanner.Scan()
						bookData[i].TanggalPengembalian = scanner.Text()
					default:
						fmt.Println("Invalid choice.")
					}
					fmt.Println("\n✅ Buku data successfully updated.")
					return
				}
			}
			fmt.Println("\n❌ Buku with the specified Judul not found.")
		case "3":
			fmt.Println("\n╔════════════════════════════╗")
			fmt.Println("║     UPDATE PEGAWAI DATA    ║")
			fmt.Println("╚════════════════════════════╝")
			fmt.Print("Enter NIP of the Pegawai to update: ")
			scanner.Scan()
			nip := scanner.Text()
			for i, p := range employeeData {
				if p.NIP == nip {
					
					fmt.Println("\n📋 Current Data:")
					printPegawai(p)

					fmt.Println("\n╔════════════════════════════╗")
					fmt.Println("║      FIELD TO UPDATE       ║")
					fmt.Println("╠════════════════════════════╣")
					fmt.Println("║ 1. Nama                    ║")
					fmt.Println("║ 2. NIP                     ║")
					fmt.Println("║ 3. Jabatan                 ║")
					fmt.Println("╚════════════════════════════╝")
					fmt.Print("Choose: ")
					scanner.Scan()
					field := scanner.Text()
					switch field {
					case "1":
						fmt.Print("Enter new Nama: ")
						scanner.Scan()
						employeeData[i].Nama = scanner.Text()
					case "2":
						fmt.Print("Enter new NIP: ")
						scanner.Scan()
						employeeData[i].NIP = scanner.Text()
					case "3":
						fmt.Print("Enter new Jabatan: ")
						scanner.Scan()
						employeeData[i].Jabatan = scanner.Text()
					default:
						fmt.Println("Invalid choice.")
					}
					fmt.Println("\n✅ Pegawai data successfully updated.")
					return
				}
			}
			fmt.Println("\n❌ Pegawai with the specified NIP not found.")
		case "4":
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func deleteData(scanner *bufio.Scanner) {
	for {
		fmt.Println("\n╔════════════════════════════╗")
		fmt.Println("║        DELETE DATA         ║")
		fmt.Println("╠════════════════════════════╣")
		fmt.Println("║ 1. Mahasiswa               ║")
		fmt.Println("║ 2. Buku                    ║")
		fmt.Println("║ 3. Pegawai                 ║")
		fmt.Println("║ 4. Back                    ║")
		fmt.Println("╚════════════════════════════╝")
		fmt.Print("Choose: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Println("\n╔════════════════════════════╗")
			fmt.Println("║    DELETE MAHASISWA DATA   ║")
			fmt.Println("╚════════════════════════════╝")
			fmt.Print("Enter NIM of the Mahasiswa to delete: ")
			scanner.Scan()
			nim := scanner.Text()
			for i, m := range studentData {
				if m.NIM == nim {
					fmt.Println("\n📋 Data to delete:")
					printStudent(m)

					fmt.Print("\nDelete this data? (yes/no): ")
					scanner.Scan()
					confirm := strings.ToLower(scanner.Text())
					if confirm == "yes" {
						studentData = append(studentData[:i], studentData[i+1:]...)
						fmt.Println("\n✅ Mahasiswa data successfully deleted.")
					} else {
						fmt.Println("\n❌ Deletion canceled.")
					}
					return
				}
			}
			fmt.Println("\n❌ Mahasiswa with the specified NIM not found.")
		case "2":
			fmt.Println("\n╔════════════════════════════╗")
			fmt.Println("║      DELETE BUKU DATA      ║")
			fmt.Println("╚════════════════════════════╝")
			fmt.Print("Enter Judul of the Buku to delete: ")
			scanner.Scan()
			judul := scanner.Text()
			for i, b := range bookData {
				if strings.ToLower(b.Judul) == strings.ToLower(judul) {
					fmt.Println("\n📋 Data to delete:")
					printBuku(b)

					fmt.Print("\nDelete this data? (yes/no): ")
					scanner.Scan()
					confirm := strings.ToLower(scanner.Text())
					if confirm == "yes" {
						bookData = append(bookData[:i], bookData[i+1:]...)
						fmt.Println("\n✅ Book data successfully deleted.")
					} else {
						fmt.Println("\n❌ Deletion canceled.")
					}
					return
				}
			}
			fmt.Println("\n❌ Buku with the specified Judul not found.")
		case "3":
			fmt.Println("\n╔════════════════════════════╗")
			fmt.Println("║     DELETE PEGAWAI DATA    ║")
			fmt.Println("╚════════════════════════════╝")
			fmt.Print("Enter NIP of the Pegawai to delete: ")
			scanner.Scan()
			nip := scanner.Text()
			for i, p := range employeeData {
				if p.NIP == nip {
					fmt.Println("\n📋 Data to delete:")
					printPegawai(p)

					fmt.Print("\nDelete this data? (yes/no): ")
					scanner.Scan()
					confirm := strings.ToLower(scanner.Text())
					if confirm == "yes" {
						employeeData = append(employeeData[:i], employeeData[i+1:]...)
						fmt.Println("\n✅ Pegawai data successfully deleted.")
					} else {
						fmt.Println("\n❌ Deletion canceled.")
					}
					return
				}
			}
			fmt.Println("\n❌ Pegawai with the specified NIP not found.")
		case "4":
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func searchData(scanner *bufio.Scanner) {
	for {
		fmt.Println("\n╔════════════════════════════╗")
		fmt.Println("║        SEARCH MENU         ║")
		fmt.Println("╠════════════════════════════╣")
		fmt.Println("║ 1. Buku yang tersedia      ║")
		fmt.Println("║ 2. Buku yang dipinjam      ║")
		fmt.Println("║ 3. Keanggotaan perpustakaan║")
		fmt.Println("║ 4. Back                    ║")
		fmt.Println("╚════════════════════════════╝")
		fmt.Print("Choose: ")
		scanner.Scan()
		bukuChoice := scanner.Text()
		switch bukuChoice {
		case "1":
			availableBooks := []Buku{}
			for _, b := range bookData {
				if b.TanggalPeminjaman == "" {
					availableBooks = append(availableBooks, b)
				}
			}
			if len(availableBooks) == 0 {
				fmt.Println("No available books.")
				break
			}
			fmt.Println("\n╔════════════════════════════╗")
			fmt.Println("║       SEARCH METHOD        ║")
			fmt.Println("╠════════════════════════════╣")
			fmt.Println("║ 1. Sequential Search       ║")
			fmt.Println("║ 2. Binary Search           ║")
			fmt.Println("╚════════════════════════════╝")
			fmt.Print("Choose: ")
			scanner.Scan()
			method := scanner.Text()
			fmt.Print("Enter keyword for Buku: ")
			scanner.Scan()
			keyword := strings.ToLower(scanner.Text())
			if method == "1" {
				sequentialSearchBukuFiltered(availableBooks, keyword)
			} else if method == "2" {
				selectionSortBuku(availableBooks, true)
				binarySearchBukuFiltered(availableBooks, keyword)
			} else {
				fmt.Println("Invalid search method.")
			}
		case "2":
			borrowedBooks := []Buku{}
			for _, b := range bookData {
				if b.TanggalPeminjaman != "" {
					borrowedBooks = append(borrowedBooks, b)
				}
			}
			if len(borrowedBooks) == 0 {
				fmt.Println("No borrowed books.")
				break
			}
			fmt.Println("\n╔════════════════════════════╗")
			fmt.Println("║       SEARCH METHOD        ║")
			fmt.Println("╠════════════════════════════╣")
			fmt.Println("║ 1. Sequential Search       ║")
			fmt.Println("║ 2. Binary Search           ║")
			fmt.Println("╚════════════════════════════╝")
			fmt.Print("Choose: ")
			scanner.Scan()
			method := scanner.Text()
			fmt.Print("Enter keyword for Buku: ")
			scanner.Scan()
			keyword := strings.ToLower(scanner.Text())
			if method == "1" {
				sequentialSearchBukuFiltered(borrowedBooks, keyword)
			} else if method == "2" {
				selectionSortBuku(borrowedBooks, true)
				binarySearchBukuFiltered(borrowedBooks, keyword)
			} else {
				fmt.Println("Invalid search method.")
			}
		case "3":
			if len(studentData) == 0 {
				fmt.Println("No Mahasiswa data available.")
				break
			}
			fmt.Print("Enter NIM or Name for Mahasiswa: ")
			scanner.Scan()
			keyword := strings.ToLower(scanner.Text())
			found := false
			for _, m := range studentData {
				if strings.Contains(strings.ToLower(m.NIM), keyword) ||
					strings.Contains(strings.ToLower(m.Name), keyword) {
					printStudent(m)
					found = true
				}
			}
			if !found {
				fmt.Println("No Mahasiswa found as library member.")
			}
		case "4":
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func sequentialSearchMahasiswa(keyword string) {
	found := false
	for _, m := range studentData {
		if strings.Contains(strings.ToLower(m.NIM), keyword) ||
			strings.Contains(strings.ToLower(m.Name), keyword) ||
			strings.Contains(strings.ToLower(m.Course), keyword) ||
			strings.Contains(strings.ToLower(m.Origin), keyword) {
			printStudent(m)
			found = true
		}
	}
	if !found {
		fmt.Println("No Mahasiswa data found.")
	}
}

func binarySearchMahasiswa(keyword string) {
	low, high := 0, len(studentData)-1
	for low <= high {
		mid := (low + high) / 2
		if strings.Contains(strings.ToLower(studentData[mid].NIM), keyword) {
			printStudent(studentData[mid])
			return
		} else if keyword < strings.ToLower(studentData[mid].NIM) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	fmt.Println("No Mahasiswa data found.")
}

func sequentialSearchBuku(keyword string) {
	found := false
	for _, b := range bookData {
		if strings.Contains(strings.ToLower(b.Author), keyword) ||
			strings.Contains(strings.ToLower(b.Judul), keyword) {
			fmt.Printf("Author: %s, Judul: %s, Tahun: %d, Tanggal Peminjaman: %s, Tanggal Pengembalian: %s\n",
				b.Author, b.Judul, b.Tahun, b.TanggalPeminjaman, b.TanggalPengembalian)
			found = true
		}
	}
	if !found {
		fmt.Println("No Buku data found.")
	}
}

func binarySearchBuku(keyword string) {
	low, high := 0, len(bookData)-1
	for low <= high {
		mid := (low + high) / 2
		if strings.Contains(strings.ToLower(bookData[mid].Judul), keyword) {
			fmt.Printf("Author: %s, Judul: %s, Tahun: %d, Tanggal Peminjaman: %s, Tanggal Pengembalian: %s\n",
				bookData[mid].Author, bookData[mid].Judul, bookData[mid].Tahun, bookData[mid].TanggalPeminjaman, bookData[mid].TanggalPengembalian)
			return
		} else if keyword < strings.ToLower(bookData[mid].Judul) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	fmt.Println("No Buku data found.")
}

func sequentialSearchPegawai(keyword string) {
	found := false
	for _, p := range employeeData {
		if strings.Contains(strings.ToLower(p.NIP), keyword) ||
			strings.Contains(strings.ToLower(p.Nama), keyword) ||
			strings.Contains(strings.ToLower(p.Jabatan), keyword) {
			fmt.Printf("NIP: %s, Nama: %s, Jabatan: %s\n", p.NIP, p.Nama, p.Jabatan)
			found = true
		}
	}
	if !found {
		fmt.Println("No Pegawai data found.")
	}
}

func binarySearchPegawai(keyword string) {
	low, high := 0, len(employeeData)-1
	for low <= high {
		mid := (low + high) / 2
		if strings.Contains(strings.ToLower(employeeData[mid].NIP), keyword) {
			fmt.Printf("NIP: %s, Nama: %s, Jabatan: %s\n", employeeData[mid].NIP, employeeData[mid].Nama, employeeData[mid].Jabatan)
			return
		} else if keyword < strings.ToLower(employeeData[mid].NIP) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	fmt.Println("No Pegawai data found.")
}

func printStudent(m Mahasiswa) {
	fmt.Println("\n╔═══════════════╦═══════════════════════╦═══════════════════╦═══════════════════╗")
	fmt.Println("║      NIM      ║         NAME          ║      COURSE       ║      ORIGIN       ║")
	fmt.Println("╠═══════════════╬═══════════════════════╬═══════════════════╬═══════════════════╣")
	fmt.Printf("║ %-13s ║ %-21s ║ %-17s ║ %-17s ║\n",
		m.NIM, truncateString(m.Name, 21), truncateString(m.Course, 17), truncateString(m.Origin, 17))
	fmt.Println("╚═══════════════╩═══════════════════════╩═══════════════════╩═══════════════════╝")
}


func sequentialSearchBukuFiltered(data []Buku, keyword string) {
	found := false
	foundBooks := []Buku{}

	for _, b := range data {
		if strings.Contains(strings.ToLower(b.Author), keyword) ||
			strings.Contains(strings.ToLower(b.Judul), keyword) {
			foundBooks = append(foundBooks, b)
			found = true
		}
	}

	if !found {
		fmt.Println("No Buku data found.")
		return
	}

	
	fmt.Println("\n╔═══════╦═══════════════════╦═══════════════════════╦═══════╦═══════════════════╦═══════════════════╗")
	fmt.Println("║  NO   ║       AUTHOR       ║         JUDUL         ║ TAHUN ║  TGL PEMINJAMAN   ║ TGL PENGEMBALIAN  ║")
	fmt.Println("╠═══════╬═══════════════════╬═══════════════════════╬═══════╬═══════════════════╬═══════════════════╣")

	
	for i, b := range foundBooks {
		fmt.Printf("║ %-5d ║ %-17s ║ %-21s ║ %-5d ║ %-17s ║ %-17s ║\n",
			i+1, truncateString(b.Author, 17), truncateString(b.Judul, 21),
			b.Tahun, truncateString(b.TanggalPeminjaman, 17), truncateString(b.TanggalPengembalian, 17))
	}

	
	fmt.Println("╚═══════╩═══════════════════╩═══════════════════════╩═══════╩═══════════════════╩═══════════════════╝")
}

func binarySearchBukuFiltered(data []Buku, keyword string) {
	low, high := 0, len(data)-1
	found := false

	for low <= high {
		mid := (low + high) / 2
		if strings.Contains(strings.ToLower(data[mid].Judul), keyword) {
			
			fmt.Println("\n╔═══════════════════╦═══════════════════════╦═══════╦═══════════════════╦═══════════════════╗")
			fmt.Println("║       AUTHOR       ║         JUDUL         ║ TAHUN ║  TGL PEMINJAMAN   ║ TGL PENGEMBALIAN  ║")
			fmt.Println("╠═══════════════════╬═══════════════════════╬═══════╬═══════════════════╬═══════════════════╣")

			fmt.Printf("║ %-17s ║ %-21s ║ %-5d ║ %-17s ║ %-17s ║\n",
				truncateString(data[mid].Author, 17), truncateString(data[mid].Judul, 21),
				data[mid].Tahun, truncateString(data[mid].TanggalPeminjaman, 17), truncateString(data[mid].TanggalPengembalian, 17))

			fmt.Println("╚═══════════════════╩═══════════════════════╩═══════╩═══════════════════╩═══════════════════╝")
			found = true
			return
		} else if keyword < strings.ToLower(data[mid].Judul) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if !found {
		fmt.Println("No Buku data found.")
	}
}


func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}


func printBuku(b Buku) {
	fmt.Println("\n╔═══════════════════╦═══════════════════════╦═══════╦═══════════════════╦═══════════════════╗")
	fmt.Println("║       AUTHOR       ║         JUDUL         ║ TAHUN ║  TGL PEMINJAMAN   ║ TGL PENGEMBALIAN  ║")
	fmt.Println("╠═══════════════════╬═══════════════════════╬═══════╬═══════════════════╬═══════════════════╣")
	fmt.Printf("║ %-17s ║ %-21s ║ %-5d ║ %-17s ║ %-17s ║\n",
		truncateString(b.Author, 17), truncateString(b.Judul, 21),
		b.Tahun, truncateString(b.TanggalPeminjaman, 17), truncateString(b.TanggalPengembalian, 17))
	fmt.Println("╚═══════════════════╩═══════════════════════╩═══════╩═══════════════════╩═══════════════════╝")
}

func printPegawai(p Pegawai) {
	fmt.Println("\n╔═══════════════════╦═══════════════════════╦═══════════════════╗")
	fmt.Println("║        NIP        ║         NAMA          ║      JABATAN      ║")
	fmt.Println("╠═══════════════════╬═══════════════════════╬═══════════════════╣")
	fmt.Printf("║ %-17s ║ %-21s ║ %-17s ║\n",
		p.NIP, truncateString(p.Nama, 21), truncateString(p.Jabatan, 17))
	fmt.Println("╚═══════════════════╩═══════════════════════╩═══════════════════╝")
}
