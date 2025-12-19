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

type FnStatement struct {
	Name      *Identifier
	Arguments []*Identifier
	Body      []Statement
}

func (fn *FnStatement) statementNode() {}

type FunctionCall struct {
	Function  Expression
	Arguments []Expression
}

func (c *FunctionCall) expressionNode() {}

type InfixExpression struct {
	Left     Expression
	Operator string
	Right    Expression
}

type KLANG struct {
	Statements []Statement
}

type Identifier struct {
	Value string
}

func (ident *Identifier) expressionNode() {}

type IntegerLiteral struct {
	Value int
}

func (in *IntegerLiteral) expressionNode() {}

type MakeStatement struct {
	Name  *Identifier
	Value Expression
}

func (m *MakeStatement) statementNode() {}

type IfStatement struct {
	Condition      Expression
	TrueStatement  []Statement
	FalseStatement []Statement
}

func (ifs *IfStatement) statementNode() {}

type WhileStatement struct {
	Condition Expression
	Body      []Statement
}

func (while *WhileStatement) statementNode() {}

type ForStatement struct {
	Initialization Statement
	Condition      Expression
	Updation       Statement
	Body           []Statement
}

func (For *ForStatement) statementNode() {}
