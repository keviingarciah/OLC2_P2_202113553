package generator

import (
	"fmt"
)

type Generator struct {
	// Variables
	Temporal int
	Label    int

	// Lists
	TempList []interface{}

	// Special
	DefaultLabel string

	// Code
	Code      []interface{}
	FinalCode []interface{}
	Natives   []interface{}
	FuncCode  []interface{}

	// Native Flags
	PrintStringFlag    bool
	PrintTrueFlag      bool
	PrintFalseFlag     bool
	PrintNilFlag       bool
	PrintMathErrorFlag bool
	ConcatStringFlag   bool
	CompareStringFlag  bool
	StringToFloatFlag  bool

	// Transfers
	BreakStack    []string
	ContinueStack []string
	ReturnStack   []string

	// Main Code Flag
	MainCode bool

	// Counters
	StackCounter int
}

func NewGenerator() Generator {
	generator := Generator{
		Temporal: 0,
		Label:    0,

		PrintStringFlag:    true,
		PrintTrueFlag:      true,
		PrintFalseFlag:     true,
		PrintNilFlag:       true,
		PrintMathErrorFlag: true,
		ConcatStringFlag:   true,
		CompareStringFlag:  true,
		StringToFloatFlag:  true,

		MainCode: true,

		StackCounter: 0,
	}
	return generator
}

func (g Generator) GetCode() []interface{} {
	return g.Code
}

func (g Generator) GetFinalCode() []interface{} {
	return g.FinalCode
}

func (g Generator) GetTemps() []interface{} {
	return g.TempList
}

// Generar un nuevo temporal
func (g *Generator) NewTemp() string {
	temp := "t" + fmt.Sprintf("%v", g.Temporal)
	g.Temporal = g.Temporal + 1
	//Lo guardamos para declararlo
	g.TempList = append(g.TempList, temp)
	return temp
}

// Generador de Label
func (g *Generator) NewLabel() string {
	temp := g.Label
	g.Label = g.Label + 1
	return "L" + fmt.Sprintf("%v", temp)
}

// Añade Label al codigo
func (g *Generator) AddLabel(Label string) {
	if g.MainCode {
		g.Code = append(g.Code, Label+":\n")
	} else {
		g.FuncCode = append(g.FuncCode, Label+":\n")
	}
}

func (g *Generator) AddIf(left string, right string, operator string, Label string) {
	if g.MainCode {
		g.Code = append(g.Code, "if("+left+" "+operator+" "+right+") goto "+Label+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, "if("+left+" "+operator+" "+right+") goto "+Label+";\n")
	}
}

func (g *Generator) AddGoto(Label string) {
	if g.MainCode {
		g.Code = append(g.Code, "goto "+Label+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, "goto "+Label+";\n")
	}
}

func (g *Generator) AddExpression(target string, left string, right string, operator string) {
	if g.MainCode {
		g.Code = append(g.Code, target+" = "+left+" "+operator+" "+right+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, target+" = "+left+" "+operator+" "+right+";\n")
	}
}

func (g *Generator) AddAssign(target, val string) {
	if g.MainCode {
		g.Code = append(g.Code, target+" = "+val+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, target+" = "+val+";\n")
	}
}

func (g *Generator) AddPrintf(typePrint string, value string) {
	if g.MainCode {
		g.Code = append(g.Code, "printf(\"%"+typePrint+"\", "+value+");\n")
	} else {
		g.FuncCode = append(g.FuncCode, "printf(\"%"+typePrint+"\", "+value+");\n")
	}
}

func (g *Generator) AddSetHeap(index string, value string) {
	if g.MainCode {
		g.Code = append(g.Code, "heap["+index+"] = "+value+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, "heap["+index+"] = "+value+";\n")
	}
}

func (g *Generator) AddGetHeap(target string, index string) {
	if g.MainCode {
		g.Code = append(g.Code, target+" = heap["+index+"];\n")
	} else {
		g.FuncCode = append(g.FuncCode, target+" = heap["+index+"];\n")
	}
}

func (g *Generator) AddSetStack(index string, value string) {
	if g.MainCode {
		g.Code = append(g.Code, "stack["+index+"] = "+value+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, "stack["+index+"] = "+value+";\n")
	}
}

func (g *Generator) AddGetStack(target string, index string) {
	if g.MainCode {
		g.Code = append(g.Code, target+" = stack["+index+"];\n")
	} else {
		g.FuncCode = append(g.FuncCode, target+" = stack["+index+"];\n")
	}
}

func (g *Generator) AddCall(target string) {
	if g.MainCode {
		g.Code = append(g.Code, target+"();\n")
	} else {
		g.FuncCode = append(g.FuncCode, target+"();\n")
	}
}

func (g *Generator) AddBr() {
	if g.MainCode {
		g.Code = append(g.Code, "\n")
	} else {
		g.FuncCode = append(g.FuncCode, "\n")
	}
}

func (g *Generator) AddComment(target string) {
	if g.MainCode {
		g.Code = append(g.Code, "//"+target+"\n")
	} else {
		g.FuncCode = append(g.FuncCode, "//"+target+"\n")
	}
}

// agregar headers
func (g *Generator) GenerateFinalCode() {
	//****************** add head
	g.FinalCode = append(g.FinalCode, "/*------HEADER------*/\n")
	g.FinalCode = append(g.FinalCode, "#include <stdio.h>\n")
	g.FinalCode = append(g.FinalCode, "float heap[10000];\n")
	g.FinalCode = append(g.FinalCode, "float stack[10000];\n")
	g.FinalCode = append(g.FinalCode, "float P;\n")
	g.FinalCode = append(g.FinalCode, "float H;\n")

	//****************** add temporal declaration
	tempArr := g.GetTemps()
	if len(tempArr) > 0 {
		g.FinalCode = append(g.FinalCode, "float ")

		tmpDec := fmt.Sprintf("%v", tempArr[0])
		tempArr = tempArr[1:]
		for _, s := range tempArr {
			tmpDec += ", "
			tmpDec += fmt.Sprintf("%v", s)
		}
		tmpDec += ";\n\n"
		g.FinalCode = append(g.FinalCode, tmpDec)
	}
	//****************** add natives functions
	if len(g.Natives) > 0 {
		g.FinalCode = append(g.FinalCode, "/*------NATIVES------*/\n")
		g.FinalCode = append(g.FinalCode, g.Natives...)
	}
	//****************** add functions
	if len(g.FuncCode) > 0 {
		g.FinalCode = append(g.FinalCode, "/*------FUNCTIONS------*/\n")
		g.FinalCode = append(g.FinalCode, g.FuncCode...)
	}
	//****************** add main
	g.FinalCode = append(g.FinalCode, "/*------MAIN------*/\n")
	g.FinalCode = append(g.FinalCode, "int main() {\n")

	for _, s := range g.Code {
		print(s.(string))
		g.FinalCode = append(g.FinalCode, "\t"+s.(string))
	}
	g.FinalCode = append(g.FinalCode, "\n\treturn 0;\n}\n")
}

func (g *Generator) PrintString() {
	if g.PrintStringFlag {
		//generando temporales y etiquetas
		newTemp1 := g.NewTemp()
		newTemp2 := g.NewTemp()
		newTemp3 := g.NewTemp()
		newLvl1 := g.NewLabel()
		newLvl2 := g.NewLabel()
		//se genera la funcion printstring
		g.Natives = append(g.Natives, "void _print_string_() {\n")
		g.Natives = append(g.Natives, "\t"+newTemp1+" = P + 1;\n")
		g.Natives = append(g.Natives, "\t"+newTemp2+" = stack[(int)"+newTemp1+"];\n")
		g.Natives = append(g.Natives, "\t"+newLvl2+":\n")
		g.Natives = append(g.Natives, "\t"+newTemp3+" = heap[(int)"+newTemp2+"];\n")
		g.Natives = append(g.Natives, "\tif("+newTemp3+" == -1) goto "+newLvl1+";\n")
		g.Natives = append(g.Natives, "\tprintf(\"%c\", (char)"+newTemp3+");\n")
		g.Natives = append(g.Natives, "\t"+newTemp2+" = "+newTemp2+" + 1;\n")
		g.Natives = append(g.Natives, "\tgoto "+newLvl2+";\n")
		g.Natives = append(g.Natives, "\t"+newLvl1+":\n")
		g.Natives = append(g.Natives, "\treturn;\n")
		g.Natives = append(g.Natives, "}\n\n")

		g.PrintStringFlag = false
	}
}

func (g *Generator) PrintTrue() {
	if g.PrintTrueFlag {
		//se genera el print true
		g.Natives = append(g.Natives, "void _print_true_() {\n")
		g.Natives = append(g.Natives, "\tprintf(\"%c\", 116);\n")
		g.Natives = append(g.Natives, "\tprintf(\"%c\", 114);\n")
		g.Natives = append(g.Natives, "\tprintf(\"%c\", 117);\n")
		g.Natives = append(g.Natives, "\tprintf(\"%c\", 101);\n")
		g.Natives = append(g.Natives, "\treturn;\n")
		g.Natives = append(g.Natives, "}\n\n")

		g.PrintTrueFlag = false
	}
}

func (g *Generator) PrintFalse() {
	if g.PrintFalseFlag {
		//se genera el print true
		g.Natives = append(g.Natives, "void _print_false_() {\n")
		g.Natives = append(g.Natives, "\tprintf(\"%c\", 102);\n")
		g.Natives = append(g.Natives, "\tprintf(\"%c\", 97);\n")
		g.Natives = append(g.Natives, "\tprintf(\"%c\", 108);\n")
		g.Natives = append(g.Natives, "\tprintf(\"%c\", 115);\n")
		g.Natives = append(g.Natives, "\tprintf(\"%c\", 101);\n")
		g.Natives = append(g.Natives, "\treturn;\n")
		g.Natives = append(g.Natives, "}\n\n")

		g.PrintFalseFlag = false
	}
}

func (g *Generator) PrintNil() {
	if g.PrintNilFlag {
		//se genera el print true
		g.Natives = append(g.Natives, "void _print_nil_() {\n")
		g.Natives = append(g.Natives, "\tprintf(\"%c\", 110);\n")
		g.Natives = append(g.Natives, "\tprintf(\"%c\", 105);\n")
		g.Natives = append(g.Natives, "\tprintf(\"%c\", 108);\n")
		g.Natives = append(g.Natives, "\treturn;\n")
		g.Natives = append(g.Natives, "}\n\n")

		g.PrintNilFlag = false
	}
}

func (g *Generator) PrintMathError() {
	if g.PrintMathErrorFlag {
		//se genera el print math error
		g.Natives = append(g.Natives, "void _print_math_error_() {\n")
		g.Natives = append(g.Natives, "\tprintf(\"%c\", 77);\n")
		g.Natives = append(g.Natives, "\tprintf(\"%c\", 97);\n")
		g.Natives = append(g.Natives, "\tprintf(\"%c\", 116);\n")
		g.Natives = append(g.Natives, "\tprintf(\"%c\", 104);\n")
		g.Natives = append(g.Natives, "\tprintf(\"%c\", 69);\n")
		g.Natives = append(g.Natives, "\tprintf(\"%c\", 114);\n")
		g.Natives = append(g.Natives, "\tprintf(\"%c\", 114);\n")
		g.Natives = append(g.Natives, "\tprintf(\"%c\", 111);\n")
		g.Natives = append(g.Natives, "\tprintf(\"%c\", 114);\n")
		g.Natives = append(g.Natives, "\tprintf(\"%c\", 10);\n")
		g.Natives = append(g.Natives, "\treturn;\n")
		g.Natives = append(g.Natives, "}\n\n")

		g.PrintMathErrorFlag = false
	}
}

func (g *Generator) GenerateConcatString() {
	if g.ConcatStringFlag {
		conca := ""
		conca1 := ""
		temp := g.NewTemp()    // t2
		auxTemp := g.NewTemp() // t3
		EV := g.NewLabel()     // L0
		EF := g.NewLabel()     // L1

		g.Natives = append(g.Natives, "void _concat_string_() {\n")

		g.Natives = append(g.Natives, "\t"+temp+" = H + 0;\n")    // t2
		g.Natives = append(g.Natives, "\t"+auxTemp+" = P + 1;\n") // t3
		conca += "stack[(int)P] = " + temp + ";\n"                // 26

		temp = g.NewTemp()        // t4
		newAuxTemp := g.NewTemp() // t5
		g.Natives = append(g.Natives, "\t"+newAuxTemp+" = stack[(int) "+auxTemp+"];\n")
		g.Natives = append(g.Natives, "\t"+temp+" = P + 2;\n")
		conca1 += newAuxTemp + " = stack[(int)" + temp + "];\n" // 15
		g.Natives = append(g.Natives, "\t"+EV+":\n")            // E0
		temp = g.NewTemp()                                      // t6
		g.Natives = append(g.Natives, "\t"+temp+" = heap[(int)"+newAuxTemp+"];\n")
		g.Natives = append(g.Natives, "\tif ("+temp+" == -1) goto "+EF+";\n")
		g.Natives = append(g.Natives, "\theap[(int)H] = "+temp+";\n")
		g.Natives = append(g.Natives, "\tH = H + 1;\n")
		g.Natives = append(g.Natives, "\t"+newAuxTemp+" = "+newAuxTemp+"+ 1;\n")
		g.Natives = append(g.Natives, "\tgoto "+EV+";\n")
		g.Natives = append(g.Natives, "\t"+EF+":\n")
		g.Natives = append(g.Natives, "\t"+conca1+"\n")
		newLabel := g.NewLabel() // L2
		EV = g.NewLabel()        // L3
		g.Natives = append(g.Natives, "\t"+EV+":")
		g.Natives = append(g.Natives, "\t"+temp+" = heap[(int)"+newAuxTemp+"];\n")
		g.Natives = append(g.Natives, "\tif ("+temp+" == -1) goto "+newLabel+";\n")
		g.Natives = append(g.Natives, "\theap[(int)H] = "+temp+";\n")
		g.Natives = append(g.Natives, "\tH = H + 1;\n")
		g.Natives = append(g.Natives, "\t"+newAuxTemp+" = "+newAuxTemp+"+ 1;\n")
		g.Natives = append(g.Natives, "\tgoto "+EV+";\n")
		g.Natives = append(g.Natives, "\t"+newLabel+":\n")
		g.Natives = append(g.Natives, "\theap[(int)H] = -1;\n")
		g.Natives = append(g.Natives, "\tH = H + 1;\n")
		g.Natives = append(g.Natives, "\t"+conca+"\n")
		g.Natives = append(g.Natives, "\treturn;\n")
		g.Natives = append(g.Natives, "}\n\n")

		g.ConcatStringFlag = false
	}
}

func (g *Generator) GenerateCompareString() {
	if g.CompareStringFlag {
		conca := ""
		// conca1 := "" // conca1
		temp := g.NewTemp()       // t2
		auxTemp := g.NewTemp()    // t3
		newTemp := g.NewTemp()    // t4
		newLabel0 := g.NewLabel() // L0
		newLabel1 := g.NewLabel() // L1
		newLabel2 := g.NewLabel() // L2
		newLabel3 := g.NewLabel() // L3

		g.Natives = append(g.Natives, "void _compare_string_() {\n")

		g.Natives = append(g.Natives, "\t"+temp+" = P + 1;\n")
		g.Natives = append(g.Natives, "\t"+auxTemp+" = stack[(int)"+temp+"];\n")
		g.Natives = append(g.Natives, "\t"+temp+" = "+temp+" + 1;\n")
		g.Natives = append(g.Natives, "\t"+newTemp+" = stack[(int)"+temp+"];\n")
		g.Natives = append(g.Natives, "\t"+newLabel1+":\n")

		temp = g.NewTemp() // t5

		g.Natives = append(g.Natives, "\t"+temp+" = heap[(int)"+auxTemp+"];\n")
		conca += auxTemp + " = " + auxTemp + " + 1;\n"

		auxTemp = g.NewTemp() // t6

		g.Natives = append(g.Natives, "\t"+auxTemp+" = heap[(int)"+newTemp+"];\n")
		g.Natives = append(g.Natives, "\tif ("+temp+" != "+auxTemp+") goto "+newLabel3+";\n")
		g.Natives = append(g.Natives, "\tif ("+temp+" == -1) goto "+newLabel2+";\n")
		g.Natives = append(g.Natives, "\t"+conca+"\n")

		g.Natives = append(g.Natives, "\t"+newTemp+" = "+newTemp+" + 1;\n")
		g.Natives = append(g.Natives, "\tgoto "+newLabel1+";\n")
		g.Natives = append(g.Natives, "\t"+newLabel2+":\n")
		g.Natives = append(g.Natives, "\tstack[(int)P] = 1;\n")
		g.Natives = append(g.Natives, "\tgoto "+newLabel0+";\n")
		g.Natives = append(g.Natives, "\t"+newLabel3+":\n")
		g.Natives = append(g.Natives, "\tstack[(int)P] = 0;\n")
		g.Natives = append(g.Natives, "\t"+newLabel0+":\n") //
		g.Natives = append(g.Natives, "\treturn;\n")
		g.Natives = append(g.Natives, "}\n\n")

		g.CompareStringFlag = false
	}
}
