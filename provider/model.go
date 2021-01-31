package provider

import (
	"fmt"
	"time"
)

// Position is the highlight position within the ebook.
type Position []int

// Book is the ebook information.
type Book struct {
	// Title is the ebook title.
	Title string

	// Author is the ebook author.
	Author string
}

// DateTime represents an instant in time (date and time).
type DateTime time.Time

// Clipping is the structure that holds a single clipping information.
type Clipping struct {
	// Book is the book information
	Book Book

	// Note is the note(s) added to the highlight.
	Note string

	// CreatedAt is the time when the clipping was created.
	CreatedAt DateTime

	// Position represents the list of positions (abstract of pages) of the highlight within the ebook.
	Position Position
}

// SetNote sets the note with a *string*.
func (c *Clipping) SetNote(note string) {
	c.Note = note
}

// SetBook sets the book with an instance of *Book*.
func (c *Clipping) SetBook(book Book) {
	c.Book = book
}

// SetCreatedAt sets the clipping creating date.
func (c *Clipping) SetCreatedAt(date DateTime) {
	c.CreatedAt = date
}

// SetPosition sets the position array.
func (c *Clipping) SetPosition(position Position) {
	c.Position = position
}

// SetTitle sets the book title.
func (b *Book) SetTitle(title string) {
	b.Title = title
}

// SetAuthor sets the book author.
func (b *Book) SetAuthor(author string) {
	b.Author = author
}

// TODO: Implement Clipping Stringer
// func (c Clipping) String() string {
// 	return c.Book.String()
// }

func (b Book) String() string {
	return fmt.Sprintf("%s - %s", b.Title, b.Author)
}
