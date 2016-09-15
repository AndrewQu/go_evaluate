package main

// Represents a token in expressions
type Token struct {
   IsOperator bool;
   Value int
}
// Return the token's operator (+ - * / ( ))
func (tk *Token) Operator() int {
    if !tk.IsOperator {
        panic("Try to get operator, but it is an operand.")
    }
    return tk.Value
}
// Return the token's operand
func (tk *Token) Operand() int {
    if tk.IsOperator {
        panic("Try to get operanad, but it is an operator.")
    }
    return tk.Value
}
// Returns the tokens priority (in increasing order): numbers, + -, * /, ( )
func (tk *Token) Priority() byte {
   if tk.IsOperator == false {
      return 0
   } else {
       op := tk.Operator()
       if op == '+' || op == '-' {
         return 1
      } else if op == '*' || op == '/' {
         return 2
      } else { // ( )
         return 3
      }
   }
}
// A tree node in the binary tree representing an expression
type TokenNode struct {
   TheToken Token;
   LeftNode *TokenNode;
   RightNode *TokenNode;
   Parent *TokenNode;
}
// Returns the tree node's priority
func (tn *TokenNode) Priority() byte {
    return tn.TheToken.Priority()
}
func (tn *TokenNode) IsLeftBracket() bool {
    return tn.TheToken.Value == '('
}
// Recursively evaluate tree node
func EvalTreeNode(node *TokenNode) int {
   if node.TheToken.IsOperator == false {
      return node.TheToken.Operand()
   } else {
      a := EvalTreeNode(node.LeftNode)
      b := EvalTreeNode(node.RightNode)
      switch node.TheToken.Operator() {
      case '+' :
         return a + b
      case '-' :
         return a - b
      case '*' :
         return a * b
      case '/' :
         return a / b
      }
      return 0;
   }
}