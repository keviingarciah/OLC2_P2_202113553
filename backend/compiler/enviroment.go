package compiler

import (
	"backend/structures"
	"strconv"
)

// Environment
type Environment struct {
	Symbols map[string]structures.Symbol
	Structs map[string]structures.Struct
	Parent  *Environment
	Name    string
}

// Environment global name
var environmentName = "Global"

// Push Environment
func (v *Visitor) PushEnvironment(envName string) {
	newEnv := &Environment{
		Symbols: make(map[string]structures.Symbol),
		Parent:  v.currentEnv,
		Name:    envName,
	}
	v.Stack = append(v.Stack, newEnv)
	v.currentEnv = newEnv
}

// Pop Environment
func (v *Visitor) PopEnvironment() {
	if len(v.Stack) > 1 {
		v.Stack = v.Stack[:len(v.Stack)-1]
		v.currentEnv = v.Stack[len(v.Stack)-1]
	}
}

// Find Symbol Environment
func (v *Visitor) FindSymbol(name string) (structures.Symbol, bool) {
	env := v.currentEnv
	for env != nil {
		if symbol, ok := env.Symbols[name]; ok {
			return symbol, true
		}
		env = env.Parent
	}
	return structures.Symbol{}, false
}

// Find Environment
func (v *Visitor) FindSymbolEnvironment(name string) (*Environment, bool) {
	env := v.currentEnv.Parent // Comenzamos desde el entorno padre del entorno actual
	for env != nil {
		if _, ok := env.Symbols[name]; ok {
			return env, true
		}
		env = env.Parent
	}
	return nil, false
}

// Save Symbol
func (v *Visitor) SaveSymbol(line int, column int, symbolType string, symbolId string, dataType string, symbolValue string, symbolEnv string, address int) {
	// save the symbol in the current enviroment
	v.currentEnv.Symbols[symbolId] = structures.Symbol{
		Line:       line,
		Column:     column,
		Type:       symbolType,
		DataType:   dataType,
		Value:      symbolValue,
		Enviroment: symbolEnv,
		Address:    strconv.Itoa(address),
	}

	// save the value table symbol
	v.SymbolTable[symbolId] = structures.Symbol{
		Line:       line,
		Column:     column,
		Type:       symbolType,
		DataType:   dataType,
		Value:      symbolValue,
		Enviroment: symbolEnv,
		Address:    strconv.Itoa(address),
	}
}

// Find Struct Environment
func (v *Visitor) FindStruct(name string) (structures.Struct, bool) {
	env := v.currentEnv
	for env != nil {
		if strct, ok := env.Structs[name]; ok {
			return strct, true
		}
		env = env.Parent
	}
	return structures.Struct{}, false
}

// Find Environment for Struct
func (v *Visitor) FindStructEnvironment(name string) (*Environment, bool) {
	env := v.currentEnv.Parent // Comenzamos desde el entorno padre del entorno actual
	for env != nil {
		if _, ok := env.Structs[name]; ok {
			return env, true
		}
		env = env.Parent
	}
	return nil, false
}

// Save Struct
func (v *Visitor) SaveStruct(name string, strct structures.Struct) {
	v.currentEnv.Structs[name] = strct
}

/*
// Save StructAttribute
func (v *Visitor) SaveStructAttribute(structName string, attributeName string, attribute structures.StructAttribute) {
	if strct, ok := v.FindStruct(structName); ok {
		strct.Atributes[attributeName] = attribute
		v.SaveStruct(structName, strct)
	}
}
*/
