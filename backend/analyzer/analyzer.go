package analyzer

import (
	"backend/compiler"
	"backend/parser"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gofiber/fiber/v2"
)

// ------------------ Response ------------------
type Response struct {
	Output      string
	Cst         string
	SymbolTable string
	ErrorTable  string
	Message     string
}

type Request struct {
	Content string `json:"content"`
}

// ------------------ Analyzer ------------------
func Analyzer() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Limpiar la consola
		//compiler.ClearPrint()

		// Aqui se obtiene del body el codigo
		code := string(c.Body())

		// Initialize the lexer & parser
		input := antlr.NewInputStream(code)
		lexer := parser.NewGrammarLexer(input)
		tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		p := parser.NewGrammarParser(tokens)

		// Error handling
		errorListener := NewMyErrorListener()
		lexer.RemoveErrorListeners()
		lexer.AddErrorListener(errorListener)
		p.RemoveErrorListeners()
		p.AddErrorListener(errorListener)

		// Build and walk parse tree
		p.BuildParseTrees = true // tell the parser to build parse trees
		tree := p.Start_()       // parse the input

		// Create the visitor
		visitor := compiler.NewVisitor()
		// Visit the tree
		visitor.Visit(tree)

		// Generate the code
		visitor.Generator.GenerateFinalCode()

		// Get the output
		output := ""
		for _, item := range visitor.Generator.GetFinalCode() {
			output += item.(string)
		}

		// Create the output
		response := Response{
			Output:  output,
			Message: "Compilado correctamente!",
		}
		return c.Status(fiber.StatusOK).JSON(response)
	}
}
