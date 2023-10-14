package structures

type SemanticError struct {
	Line    int
	Column  int
	Message string
}
