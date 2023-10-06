package structures

// ------------------ Symbol Table ------------------
type Symbol struct {
	Line       int
	Column     int
	SymbolType string
	DataType   string
	Value      interface{}
	Enviroment string
}

// ------------------ Errors ------------------
type SemanticError struct {
	Line    int
	Column  int
	Message string
}
