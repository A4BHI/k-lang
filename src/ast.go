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

type ExpressionStatement struct {
	Expression Expression
}

func (ES ExpressionStatement) expressionNode() {}

type InfixExpression struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (Infix InfixExpression) expressionNode() {}

type PrefixExpression struct {
	Operator   string
	Expression Expression
}

func (Prefix PrefixExpression) expressionNode() {}

type PostfixExpression struct {
	Expression Expression
	Operator   string
}

func (Postfix PostfixExpression) expressionNode() {}

type KLANG struct {
	Statements []Statement
}

type Identifier struct {
	Value string
}

func (ident *Identifier) expressionNode() {}

type BoolLiteral struct {
	Value bool
}

func (b *BoolLiteral) expressionNode() {}

type IntegerLiteral struct {
	Value int
}

func (in *IntegerLiteral) expressionNode() {}

type StringLiteral struct {
	Value string
}

func (S *StringLiteral) expressionNode() {}

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

type AssignmentStatement struct {
	Name  *Identifier
	Value Expression
}

func (Assignment AssignmentStatement) statementNode() {}

type ReturnStatement struct {
	Expression Expression
}

func (Return ReturnStatement) statementNode() {}
