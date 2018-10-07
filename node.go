package parser

type BinaryNode struct {
	Right interface{}
	Left  interface{}
	Op    string
}

type FunctionCallNode struct {
	Name string
	Args []interface{}
}

type FieldReferenceNode struct {
	Path []string
}

type StringNode struct {
	Value string
}

type BooleanNode struct {
	Value bool
}

type FloatingPointNode struct {
	Value float64
}

type IntegerNode struct {
	Value int
}
