package compiler

import (
	"backend/structures"
	"strconv"
)

type Environment struct {
	Symbols map[string]structures.Symbol
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
