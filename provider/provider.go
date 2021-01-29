package provider

import (
	"fmt"
	"time"
)

/**
This is a sample of the clipping file.
==========
Crônicas de Nárnia (C.S.Lewis)
- Sua nota ou posição 334 | Adicionado: sábado, 16 de maio de 2015 13:16:48

CSLewis ja conhecia a teoria dos multiversos? Qdo foi proposta?
==========
Crônicas de Nárnia (C.S.Lewis)
- Seu destaque ou posição 333-334 | Adicionado: sábado, 16 de maio de 2015 13:16:48

Lembre-se do túnel; não pertence a nenhuma das casas, mas você pode andar por ele e ...
==========
Talmidim52: O passo a passo de Jesus em meditações semanais (Kivtz, Ed René)
- Your Highlight on Location 466-470 | Added on Monday, January 21, 2019 1:36:43 PM

as palavras de Jesus foram “o Reino de Deus está próximo”. Mas esse “próximo” não qu...
==========
Talmidim52: O passo a passo de Jesus em meditações semanais (Kivtz, Ed René)
- Your Highlight on Location 494-498 | Added on Monday, January 21, 2019 1:40:10 PM

O que é passar por uma experiência de arrependimento? O que é experimentar a metanoia? ...
==========
*/

// Position is the highlight position within the ebook
type Position []uint

// Book is the ebook information
type Book struct {
	// Title is the ebook title
	Title string

	// Author is the ebook author
	Author string
}

// DateTime represents an instant in time (date and time)
type DateTime time.Time

// Clipping is the structure that holds a single clipping information
type Clipping struct {
	// Book is the book information
	Book Book

	// Note is the note(s) added to the highlight
	Note string

	// CreatedAt is the time when the clipping was created
	CreatedAt DateTime

	// Position represents the list of positions (abstract of pages) of the highlight within the ebook
	Position Position
}

func (c *Clipping) SetNote(note string) {
	c.Note = note
}

func (c *Clipping) SetBook(book Book) {
	c.Book = book
}

func (c *Clipping) SetBookInfo(title, author string) {
	c.Book.Author = author
	c.Book.Title = title
}

func (c *Clipping) SetBookTitle(title string) {
	c.Book.Title = title
}

func (c *Clipping) SetBookAuthor(author string) {
	c.Book.Author = author
}

func (c *Clipping) SetCreatedAtFromString(createdAt string) {
	// TODO: Convert from string to DateTime
}

func (c *Clipping) SetCreatedAt(date DateTime) {
	c.CreatedAt = date
}

func (c *Clipping) SetPosition(position Position) {
	c.Position = position
}

func (c *Clipping) AddPosition(position uint) {
	for _, v := range c.Position {
		if v == position {
			return
		}
	}

	c.Position = append(c.Position, position)
}

func (c *Clipping) RemovePosition(position uint) {
	// TODO: Implement RemovePosition
	panic("Not Implemented.")
}

func (b *Book) SetInfo(title, author string) {
	b.Author = author
	b.Title = title
}

func (b *Book) SetTitle(title string) {
	b.Title = title
}

func (b *Book) SetAuthor(author string) {
	b.Author = author
}

func (c *Clipping) String() string {
	// TODO: Implement Clipping Stringer
	return c.Book.String()
}

func (b *Book) String() string {
	return fmt.Sprintf("%s - %s", b.Title, b.Author)
}
