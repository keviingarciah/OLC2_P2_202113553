package structures

// ------------------ Structs------------------
type Struct struct {
	Attributes map[string]StructAttribute
}

type StructAttribute struct {
	Index         int
	PrimitiveType string
	ExprType      string
	StructType    string
	Value         string
}
