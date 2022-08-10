package statements

import (
	// "github.com/gopackt/jinjia2/exec"
	"fmt"

	"github.com/gopackt/jinjia2/exec"
	"github.com/gopackt/jinjia2/nodes"
	"github.com/gopackt/jinjia2/parser"
	"github.com/gopackt/jinjia2/tokens"
)

type FirstofStmt struct {
	Location *tokens.Token
	Args     []nodes.Expression
}

func (stmt *FirstofStmt) Position() *tokens.Token { return stmt.Location }
func (stmt *FirstofStmt) String() string {
	t := stmt.Position()
	return fmt.Sprintf("FirstofStmt(Args=%s, Line=%d Col=%d)", stmt.Args, t.Line, t.Col)
}

func (stmt *FirstofStmt) Execute(r *exec.Renderer, tag *nodes.StatementBlock) error {
	for _, arg := range stmt.Args {
		val := r.Eval(arg)
		if val.IsError() {
			return val
		}

		if val.IsTrue() {
			r.RenderValue(val)
			return nil
		}
	}

	return nil
}

func firstofParser(p *parser.Parser, args *parser.Parser) (nodes.Statement, error) {
	stmt := &FirstofStmt{
		Location: p.Current(),
	}

	for !args.End() {
		node, err := args.ParseExpression()
		if err != nil {
			return nil, err
		}
		stmt.Args = append(stmt.Args, node)
	}

	return stmt, nil
}

func init() {
	All.Register("firstof", firstofParser)
}
