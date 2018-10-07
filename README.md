# Salesforce Formula Parser

Parse Salesforce Formula and, Build AST in Golang

## Usage

```bash
$ go get github.com/tzmfreedom/go-salesforce-formula-parser
```

```golang
package main

import (
	"github.com/tzmfreedom/go-salesforce-formula-parser"
)

func main() {
	ast := parser.ParseString('IF(TRUE, "foo", IF(FALSE, "bar", "baz"))')
	// ast := parser.ParseFile(os.Args[1])
	
	// =>
	//   &parser.FunctionCallNode{
	//     Name: "IF",
	//     Args: []interface {}{
	//       &parser.BooleanNode{
	//         Value: true,
	//       },
	//       &parser.StringNode{
	//         Value: "foo",
	//       },
	//       &parser.FunctionCallNode{
	//         Name: "IF",
	//         Args: []interface {}{
	//           &parser.BooleanNode{
	//             Value: false,
	//           },
	//           &parser.StringNode{
	//             Value: "bar",
	//           },
	//           &parser.StringNode{
	//             Value: "baz",
	//           },
	//         },
	//       },
	//     },
	//   }
}
```