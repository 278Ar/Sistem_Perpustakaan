package Node

type Book struct {
	Id       int
	Judul    string
	Pencipta string
	Genre    string
}

type TableBook struct {
	Data Book
	Next *TableBook
}
