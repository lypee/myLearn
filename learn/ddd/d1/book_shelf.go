package d1

type BookShelf struct{
	Book []Book
}
func (b *BookShelf)Add(book Book){
	if b.Book == nil{
		b.Book = []Book{}
	}
	b.Book = append(b.Book , book)
}
