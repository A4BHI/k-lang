package main

type Lexer struct {
	Source     string
	currentPos int
	nextPos    int
	ch         byte
}

var Out []string

const (
	MAKE  = "MAKE"
	EQUAL = "EQUAL"
	IF    = "IF"
	PLUS  = "PLUS"
)

var Tokens = map[string]string{
	"make": MAKE,
	"if":   IF,
	"plus": PLUS,
}
var s string

func (l *Lexer) OutPut() {
	for i := 0; i < len(l.Source); i++ {
		if l.Source[i] == ' ' {
			l.currentPos++
		} else {
			l.ch = l.Source[i]
			b := handleOperators(l.ch)

			if !b {
				s += string(l.ch)

			}

		}
	}

}

func handleOperators(op byte) bool {
	switch op {
	case '+':
		Out = append(Out, Tokens[string(op)])
		return true
		break
	default:
		return false
	}
	return true
}

func main() {
	// OutPut("make")

}
