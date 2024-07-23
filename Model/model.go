package Model

import (
	"Sistem_Perpustakaan/Database"
	"Sistem_Perpustakaan/Node"
)

func InsertBook(id int, judul string, pencipta string, genre string) {
	newDataBook := Node.TableBook{
		Data: Node.Book{Id: id, Judul: judul, Pencipta: pencipta, Genre: genre},
		Next: nil,
	}
	var tempLL *Node.TableBook
	tempLL = &Database.DBBook
	if Database.DBBook.Next == nil {
		Database.DBBook.Next = &newDataBook
	} else {
		for tempLL.Next != nil {
			tempLL = tempLL.Next
		}
		tempLL.Next = &newDataBook
	}
}

func ReadAllBook() *Node.TableBook {
	var tempLL *Node.TableBook
	tempLL = &Database.DBBook
	if Database.DBBook.Next == nil {
		return nil
	} else {
		return tempLL
	}
}

func DeleteBook(id int) bool {
	var tempLL *Node.TableBook
	tempLL = &Database.DBBook
	if tempLL.Next == nil {
		return false
	} else {
		for tempLL.Next != nil {
			if tempLL.Next.Data.Id == id {
				tempLL.Next = tempLL.Next.Next
				return true
			}
			tempLL = tempLL.Next
		}
		return false
	}
}

func SearchBook(id int) *Node.TableBook {
	var tempLL *Node.TableBook
	tempLL = Database.DBBook.Next
	cek := false
	if Database.DBBook.Next == nil {
		return nil
	} else {
		for tempLL != nil {
			if id == tempLL.Data.Id {
				cek = true
				return tempLL
			}
			tempLL = tempLL.Next
		}
		if cek != true {
			return nil
		}
	}
	return nil
}

func UpdateBook(Book Node.Book) bool {
	addr := SearchBook(Book.Id)
	if addr == nil {
		return false
	} else {
		addr.Data.Judul = Book.Judul
		addr.Data.Pencipta = Book.Pencipta
		addr.Data.Genre = Book.Genre
		return true
	}
}
