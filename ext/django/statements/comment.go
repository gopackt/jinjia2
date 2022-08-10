package statements

import (
	"fmt"

	"github.com/gopackt/jinjia2/nodes"
	"github.com/gopackt/jinjia2/parser"
	"github.com/gopackt/jinjia2/tokens"
)

type CommentStmt struct {
	Location *tokens.Token
}

func (stmt *CommentStmt) Position() *tokens.Token { return stmt.Location }
func (stmt *CommentStmt) String() string {
	t := stmt.Position()
	return fmt.Sprintf("Block(Line=%d Col=%d)", t.Line, t.Col)
}

func commentParser(p *parser.Parser, args *parser.Parser) (nodes.Statement, error) {
	commentNode := &CommentStmt{p.Current()}

	err := p.SkipUntil("endcomment")
	if err != nil {
		return nil, err
	}

	if !args.End() {
		return nil, args.Error("Tag 'comment' does not take any argument.", nil)
	}

	return commentNode, nil
}

func init() {
	All.Register("comment", commentParser)
}
