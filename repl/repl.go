package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/alfiehiscox/monkey-go/evaluator"
	"github.com/alfiehiscox/monkey-go/lexer"
	"github.com/alfiehiscox/monkey-go/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errs []string) {
	for _, msg := range errs {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
