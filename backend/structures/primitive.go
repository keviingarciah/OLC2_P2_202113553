package structures

// ------------------ PRIMITIVE ------------------
type Primitive struct {
	Value      string
	Type       string
	DataType   string
	IsTemporal bool
}

func (p Primitive) GetValue() string {
	return p.Value
}

func (p Primitive) GetType() string {
	return p.Type
}

func (p Primitive) GetDataType() string {
	return p.DataType
}

func (p Primitive) GetIsTemporal() bool {
	return p.IsTemporal
}
