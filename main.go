package main

import (
	"Sistem_Perpustakaan/Model"
	"Sistem_Perpustakaan/Node"
	"fmt"
)

func main() {
	// Model Insert Book
	Model.InsertBook(1, "The Great Gatsby", "F. Scott Fitzgerald", "Fiction")
	Model.InsertBook(2, "1984", "George Orwell", "Dystopian")
	Model.InsertBook(3, "To Kill a Mockingbird", "Harper Lee", "Classic")

	// Model Read all Book
	fmt.Println("Daftar semua Buku:")
	printAllBook(Model.ReadAllBook())

	// Model Search Book
	Book := Model.SearchBook(2)
	if Book != nil {
		fmt.Println("Buku yang ditemukan:", Book.Data)
	} else {
		fmt.Println("Buku tidak ditemukan")
	}

	// Model Update Book
	updatedBook := Node.Book{Id: 2, Judul: "1984", Pencipta: "George Orwell", Genre: "Political Fiction"}
	success := Model.UpdateBook(updatedBook)
	if success {
		fmt.Println("Book berhasil diperbarui")
	} else {
		fmt.Println("Pembaruan Buku gagal")
	}

	// Model Delete Book
	success = Model.DeleteBook(3)
	if success {
		fmt.Println("Penghapusan Buku Berhasil")
	} else {
		fmt.Println("Penghapusan Buku gagal")
	}

	// Read all Book setelah delete Book
	fmt.Println("Daftar semua Buku setelah dihapus:")
	printAllBook(Model.ReadAllBook())
}

func printAllBook(head *Node.TableBook) {
	if head == nil {
		fmt.Println("Tidak ada Buku yang tersedia")
		return
	}
	terkini := head.Next
	for terkini != nil {
		fmt.Printf("ID: %d, Judul: %s, Pencipta: %s, Genre: %s\n", terkini.Data.Id, terkini.Data.Judul, terkini.Data.Pencipta, terkini.Data.Genre)
		terkini = terkini.Next
	}
}
