package repl

import (
	"bufio"
	"compiler/eval"
	"compiler/lexer"
	"compiler/parser"
	"fmt"
	"io"
)

const PROMPT = "QWQ >> "

func printParserErrors(out io.Writer, errs []error) {
	for _, err := range errs {
		io.WriteString(out, err.Error())
		io.WriteString(out, "\n")
	}
}

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)

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

		e := eval.Eval(program)
		if e != nil {
			io.WriteString(out, e.Inspect())
			io.WriteString(out, "\n")
		}

		// for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		// 	fmt.Fprintf(out, "%+v\n", tok)
		// }
	}
}
