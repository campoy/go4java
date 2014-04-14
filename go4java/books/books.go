package books

import "strings"

// A Book has a title and an author.
type Book struct {
	Title  string
	Author *Author
}

func (b Book) String() string { return b.Title + " by " + b.Author.String() }

// An Author has a First and Last name.
type Author struct {
	First, Last string
}

// If the Author is a nil pointer, "unknown author" is returned.
func (p *Author) String() string {
	if p == nil {
		return "unknown author"
	}
	return p.First + " " + strings.ToUpper(p.Last)
}
