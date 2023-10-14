package structures

// ------------------ Symbol Table ------------------
type Symbol struct {
	// Symbol
	Line       int
	Column     int
	Type       string
	DataType   string
	Value      string
	Enviroment string
	// C3D
	Address string
}
