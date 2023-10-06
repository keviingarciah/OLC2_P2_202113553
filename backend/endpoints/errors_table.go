package endpoints

import (
	"backend/structures"
	"fmt"
	"strings"
)

func CreateErrorsTable(lexical []LexicalError, syntax []SyntaxError, semantic []structures.SemanticError) string {
	// Contar el número de errores
	var n = 0
	// Crear el string de salida
	var sb strings.Builder
	sb.WriteString("\ndigraph G {\n")
	sb.WriteString("node [shape=none, fontname=\"Arial\", fontsize=12];\n")
	sb.WriteString("edge [color=transparent];\n")
	sb.WriteString("ET [label=<\n")

	// Agregar los nodos
	sb.WriteString("<TABLE BORDER=\"1\" CELLBORDER=\"0\" CELLSPACING=\"0\" BGCOLOR=\"white\">\n")
	sb.WriteString("<TR>\n")
	sb.WriteString("<TD COLSPAN=\"6\" BGCOLOR=\"BLACK\"><FONT COLOR=\"white\">Tabla de Errores</FONT></TD>\n")
	sb.WriteString("</TR>\n")
	sb.WriteString("<TR>\n")
	sb.WriteString("<TD PORT=\"N\" BGCOLOR=\"#F0F0F0\">No.</TD>\n")
	sb.WriteString("<TD PORT=\"Type\" BGCOLOR=\"#F0F0F0\">Tipo</TD>\n")
	sb.WriteString("<TD PORT=\"Description\" BGCOLOR=\"#F0F0F0\">Descripción</TD>\n")
	sb.WriteString("<TD PORT=\"Linea\" BGCOLOR=\"#F0F0F0\">Línea</TD>\n")
	sb.WriteString("<TD PORT=\"Columna\" BGCOLOR=\"#F0F0F0\">Columna</TD>\n")
	sb.WriteString("</TR>\n")

	// Lexical Errors
	if len(lexical) != 0 {
		for _, error := range lexical {
			n++
			sb.WriteString("<TR>\n")
			sb.WriteString(fmt.Sprintf("<TD PORT=\"N_%d\">%d</TD>\n", n, n))
			sb.WriteString(fmt.Sprintf("<TD PORT=\"Type_%d\">%s</TD>\n", n, "Léxico"))
			sb.WriteString(fmt.Sprintf("<TD PORT=\"Description_%d\">%s</TD>\n", n, error.Message))
			sb.WriteString(fmt.Sprintf("<TD PORT=\"Linea_%d\">%d</TD>\n", n, error.Line))
			sb.WriteString(fmt.Sprintf("<TD PORT=\"Columna_%d\">%d</TD>\n", n, error.Column))
			sb.WriteString("</TR>\n")

		}

	}

	// Syntax Errors
	if len(syntax) != 0 {
		for _, error := range syntax {
			n++
			sb.WriteString("<TR>\n")
			sb.WriteString(fmt.Sprintf("<TD PORT=\"N_%d\">%d</TD>\n", n, n))
			sb.WriteString(fmt.Sprintf("<TD PORT=\"Type_%d\">%s</TD>\n", n, "Léxico"))
			sb.WriteString(fmt.Sprintf("<TD PORT=\"Description_%d\">%s</TD>\n", n, error.Message))
			sb.WriteString(fmt.Sprintf("<TD PORT=\"Linea_%d\">%d</TD>\n", n, error.Line))
			sb.WriteString(fmt.Sprintf("<TD PORT=\"Columna_%d\">%d</TD>\n", n, error.Column))
			sb.WriteString("</TR>\n")

		}

	}

	// Semantic Errors
	if len(semantic) != 0 {
		for _, error := range semantic {
			n++
			sb.WriteString("<TR>\n")
			sb.WriteString(fmt.Sprintf("<TD PORT=\"N_%d\">%d</TD>\n", n, n))
			sb.WriteString(fmt.Sprintf("<TD PORT=\"Type_%d\">%s</TD>\n", n, "Semantico"))
			sb.WriteString(fmt.Sprintf("<TD PORT=\"Description_%d\">%s</TD>\n", n, error.Message))
			sb.WriteString(fmt.Sprintf("<TD PORT=\"Linea_%d\">%d</TD>\n", n, error.Line))
			sb.WriteString(fmt.Sprintf("<TD PORT=\"Columna_%d\">%d</TD>\n", n, error.Column))
			sb.WriteString("</TR>\n")
		}
	}
	sb.WriteString("</TABLE>\n")
	sb.WriteString(">];\n")
	sb.WriteString("}")

	return sb.String()
}
