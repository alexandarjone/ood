/*
    requirement:
        1. online reading platform

    clarification:
        1. user: only reader
        2. each page is fixed length?
        3. functionality:
            i. get book content based on title
            ii. view a list of book title
            iii. keep track of reading progress
    
    workflow:
        i. user get a list of book title from booklist
        ii. user get the content of a book with the title
        iii. uesr reads the book. and if user read the book again, they start reading from the page they left
    
    core objects:
        i. book
        ii. book list
        iii. user
        iv. reading record

    
    feedback
        1. communication is good
        2. question answering is good
        3. 
*/
type Book interface {
    GetID() int64
    GetTitle() string
    GetContent() string
}

type BookList interface {
    GetBooks() []Book
    // GetBookByTitle takes title of the book, and return the book
    GetBookByID(int64) Book
    // since reader is the only user, we don't need functions such as AddBook(Book)
}

type User interface {
    GetID() int64
}

// type ReadingRecord interface {
//     // GetPage returns the page where user left based on userID and bookID
//     GetPage(int64, int64) int
//     // GetNextPage returns next page
//     GetNextPage(int64, int64) int
// }

type ReadingRecord interface {
    // GetPage returns the page where user left based on userID and bookID
    GetStartTextIndex(int64, int64) int
    // GetNextPage returns next page
    GetNextPageStartTextIndex(int64, int64) int
}

// Book implementation
type book struct {
    id int64
    title string
    content []string
}

func (b *book) GetID() int64 {
    return b.id
}
func (b *book) GetTitle() string {
    return b.title
}
func (b *book) GetContent(pageIndex int) string {
    return b.content[pageIndex]
}

func NewBook(id, title string, content []string) Book {
    return &book{
        id: id,
        title: title,
        content: content,
    }
}

// ReadingRecord implementation
type readingRecord struct {
    booklist BookList
    // userID -> bookID -> pageindex
    readingRecord map[int64]map[int64]int

}

func (r *readingRecord) GetStartTextIndex(uyserID int64, bookID int64) int {
    if _, exists := readingRecord[uyserID]; !exists {
        readingRecord[uyserID] = make(map[int64]int)
    }
    return r.readingRecord[uyserID][bookID]
}

func (r *readingRecord) GetNextPageStartTextIndex(uyserID int64, bookID int64) int {
    if _, exists := readingRecord[uyserID]; !exists {
        readingRecord[uyserID] = make(map[int64]int)
    }
    startTextIndex := readingRecord[uyserID][bookID]
    pageTextLength := GetPageTextLength(userID)
    
    startTextIndex += pageTextLength
    return startTextIndex
}

// func (r *readingRecord) GetPage(uyserID int64, bookID int64) int {
//     if _, exists := readingRecord[uyserID]; !exists {
//         readingRecord[uyserID] = make(map[int64]int)
//     }
//     return r.readingRecord[uyserID][bookID]
// }

// func (r *readingRecord) GetNextPage(uyserID int64, bookID int64) int {
//     if _, exists := readingRecord[uyserID]; !exists {
//         readingRecord[uyserID] = make(map[int64]int)
//     }
//     pageIndex := r.readingRecord[uyserID][bookID]
//     pageIndex = (pageIndex+1) % len(booklist.GetBookByID(bookID))
//     r.readingRecord[uyserID][bookID] = pageIndex
//     return pageIndex
// }

func NewReadingRecord() ReadingRecord {
    return &readingRecord{
        readingRecord: make(map[int64]map[int64]int)
    }
}


func main() {
    fmt.Printf("Hello LeetCoder")
}