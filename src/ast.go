package klang

type Node interface {
	tokLiteral() string
}

type Statement interface {
	Node
	statementNode()
}
type Expression interface {
	Node
	expressionNode()
}

type Identifier struct {
	Value string
}

func (ident *Identifier) expressionNode() {}

type IntegerLiteral struct {
	Value int
}

func (in *IntegerLiteral) expressionNode() {}

type CallExpression struct {
	Function   Expression
	Parameters []Expression
}

type InfixExpression struct {
	Left     Expression
	Operator string
	Right    Expression
}

type KLANG struct {
	Statements []Statement
}

type MakeStatement struct {
	Name  *Identifier
	Value Expression
}

func (m MakeStatement) statementNode() {}

type IfStatement struct {
	Condition      Expression
	TrueStatement  []Statement
	FalseStatement []Statement
}

func (ifs IfStatement) statementNode() {}

type WhileStatement struct {
	Condition Expression
	Body      []Statement
}

func (while *WhileStatement) statementNode() {}
