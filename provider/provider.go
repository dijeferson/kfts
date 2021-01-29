package provider

import "time"

// Crônicas de Nárnia (C.S.Lewis)
// - Sua nota ou posição 334 | Adicionado: sábado, 16 de maio de 2015 13:16:48

// CSLewis ja conhecia a teoria dos multiversos? Qdo foi proposta?
// ==========
// Crônicas de Nárnia (C.S.Lewis)
// - Seu destaque ou posição 333-334 | Adicionado: sábado, 16 de maio de 2015 13:16:48

// Lembre-se do túnel; não pertence a nenhuma das casas, mas você pode andar por ele e entrar em qualquer uma delas. Não será este bosque uma coisa parecida?... Um lugar que não pertence a nenhum dos mundos, mas que dá acesso a todos os mundos?
// ==========

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
