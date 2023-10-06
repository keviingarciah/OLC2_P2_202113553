package reports

import (
	"backend/structures"
	"fmt"
	"strings"
)

func CreateTS(ts map[string]structures.Symbol) string {
	// Crear el string de salida
	var sb strings.Builder
	sb.WriteString("\ndigraph G {\n")
	sb.WriteString("node [shape=none, fontname=\"Arial\", fontsize=12];\n")
	sb.WriteString("edge [color=transparent];\n")
	sb.WriteString("TS [label=<\n")

	// Agregar los nodos
	sb.WriteString("<TABLE BORDER=\"1\" CELLBORDER=\"0\" CELLSPACING=\"0\" BGCOLOR=\"white\">\n")
	sb.WriteString("<TR>\n")
	sb.WriteString("<TD COLSPAN=\"6\" BGCOLOR=\"BLACK\"><FONT COLOR=\"white\">Tabla de Símbolos</FONT></TD>\n")
	sb.WriteString("</TR>\n")
	sb.WriteString("<TR>\n")
	sb.WriteString("<TD PORT=\"ID\" BGCOLOR=\"#F0F0F0\">ID</TD>\n")
	sb.WriteString("<TD PORT=\"SymbolType\" BGCOLOR=\"#F0F0F0\">Tipo de Símbolo</TD>\n")
	sb.WriteString("<TD PORT=\"DataType\" BGCOLOR=\"#F0F0F0\">Tipo de Dato</TD>\n")
	sb.WriteString("<TD PORT=\"Enviroment\" BGCOLOR=\"#F0F0F0\">Ámbito</TD>\n")
	sb.WriteString("<TD PORT=\"Line\" BGCOLOR=\"#F0F0F0\">Línea</TD>\n")
	sb.WriteString("<TD PORT=\"Column\" BGCOLOR=\"#F0F0F0\">Columna</TD>\n")
	sb.WriteString("</TR>\n")

	// Iterar el mapa y mostrar clave y atributos de Symbol
	for key, symbol := range ts {
		sb.WriteString("<TR>\n")
		sb.WriteString(fmt.Sprintf("<TD PORT=\"%s\">%s</TD>\n", key, key))
		sb.WriteString(fmt.Sprintf("<TD PORT=\"SymbolType_%s\">%s</TD>\n", key, symbol.SymbolType))
		sb.WriteString(fmt.Sprintf("<TD PORT=\"DataType_%s\">%s</TD>\n", key, symbol.DataType))
		sb.WriteString(fmt.Sprintf("<TD PORT=\"Enviroment_%s\">%s</TD>\n", key, symbol.Enviroment))
		sb.WriteString(fmt.Sprintf("<TD PORT=\"Line_%s\">%d</TD>\n", key, symbol.Line))
		sb.WriteString(fmt.Sprintf("<TD PORT=\"Column_%s\">%d</TD>\n", key, symbol.Column))
		sb.WriteString("</TR>\n")
	}

	sb.WriteString("</TABLE>\n")
	sb.WriteString(">];\n")
	sb.WriteString("}")

	return sb.String()
}
