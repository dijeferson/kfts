package provider

type operation string

const (
	highlight operation = "highlight" // Destaque
	note      operation = "note"      // Nota
	bookmark  operation = "bookmark"  // Marcador
	unknown   operation = "unknown"
)
