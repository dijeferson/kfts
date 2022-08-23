package provider

import (
	"strconv"
	"strings"
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

// entrySeparator defines the string that is used by the Kindle as the entry separator in
// clippings.
const entrySeparator string = "==========\r\n"

// Processor is responsible for processing the data.
type Processor struct{}

// Process processes the clipping payload, and parses it into a slice of Clipping.
func (p Processor) Process(payload string) (*[]Clipping, error) {
	var result []Clipping

	entries := strings.Split(payload, entrySeparator)

	for _, entry := range entries {
		// Skip empty entries
		if entry != "" {
			clip, err := p.parseClipping(entry)

			if err != nil {
				return nil, err
			}

			result = append(result, *clip)
		}
	}

	return &result, nil
}

func (p Processor) parseClipping(clip string) (*Clipping, error) {
	var result Clipping

	lines := strings.Split(clip, "\r\n")
	book, err := p.parseBookInfo(lines[0])
	if err != nil {
		return nil, &ErrFailedToProcessBookInfo{err: err, data: clip}
	}
	result.SetBook(*book)

	position, err := p.parsePosition(lines[1])
	if err != nil {
		return nil, &ErrFailedToProcessPosition{err: err, data: clip}
	}

	result.SetPosition(*position)
	createdAt, err := p.parseCreationDate(lines[1])

	if err != nil {
		return nil, &ErrFailedToProcessCreationDate{err: err, data: clip}
	}

	if createdAt != nil {
		result.SetCreatedAt(*createdAt)
	}

	op, err := p.parseOperation(lines[1])

	if err != nil {
		return nil, err
	}

	if op == note || op == highlight {
		for i := 2; i < len(lines); i++ {
			result.Note += lines[i]
		}
	}

	return &result, nil
}

func (p Processor) parseOperation(input string) (operation, error) {
	opSection := strings.ToLower(strings.TrimSpace(strings.Split(input, "|")[0]))
	switch {
	case strings.Contains(opSection, "destaque") || strings.Contains(opSection, "highlight"):
		return highlight, nil
	case strings.Contains(opSection, "nota") || strings.Contains(opSection, "note"):
		return note, nil
	case strings.Contains(opSection, "marcador") || strings.Contains(opSection, "bookmark"):
		return bookmark, nil
	}

	return unknown, &ErrInvalidOperation{data: opSection}
}

func (p Processor) parseBookInfo(info string) (*Book, error) {
	var result Book

	bookData := strings.Split(info, "(")

	if len(bookData) > 1 {
		result.SetAuthor(bookData[1][:len(bookData[1])-1])
	}

	result.SetTitle(bookData[0])

	return &result, nil
}

func (p Processor) parsePosition(input string) (*Position, error) {
	positionSection := strings.TrimSpace(strings.Split(input, "|")[0])
	positionSplits := strings.Split(positionSection, " ")[len(strings.Split(positionSection, " "))-1:]
	positions := strings.Split(positionSplits[0], "-")

	var result Position
	for _, pos := range positions {
		position, err := strconv.Atoi(pos)

		if err != nil {
			return nil, &ErrFailedToParsePosition{err: err, data: input}
		}

		result = append(result, position)
	}

	return &result, nil
}

func (p Processor) parseCreationDate(input string) (*DateTime, error) {
	dateSection := strings.TrimSpace(strings.Split(input, "|")[1])

	rawDate := strings.Replace(dateSection, "Added on ", "", 1)
	rawDate = strings.Replace(rawDate, "Adicionado: ", "", 1)

	dt, err := time.Parse("Monday, January _2, 2006 3:04:05 PM", rawDate)

	if err != nil {
		return nil, nil //err --> swallow error for now
	}

	result := DateTime(dt)

	return &result, nil
}
