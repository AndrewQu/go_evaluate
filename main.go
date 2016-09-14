package main

import (
   "fmt"
   "strconv"
   "strings"
)

type Token struct {
   IsOperator bool;
   Value int
}

type TokenNode struct {
   Node Token;
   LeftNode *TokenNode;
   RightNode *TokenNode;
   Parent *TokenNode;
}

type TokenArray struct {
   NumTokens int;
   Tokens [] Token
}
// Returns the tokens priority (in increasing order): 
// operands, + -, * /, ( )
func (tk *Token) Priority() byte {
   if tk.IsOperator == false {
      return 0
   } else {
      if tk.Value == '+' || tk.Value == '-' {
         return 1
      } else if tk.Value == '*' || tk.Value == '/' {
         return 2
      } else { // ( )
         return 3
      }
   }
}

func (tn *TokenNode) Priority() byte {
    return tn.Node.Priority()
}

func (ta *TokenArray) Tokenise(exp string) {
   ops := "+-*/()"
   
   exp = strings.Replace(exp, " ", "", len(exp))
   ta.Tokens = make([] Token, len(exp))
   ta.NumTokens = 0
   
   ts := 0 // Start of the next token
   for i := 0; i < len(exp); i++ {
      if strings.IndexByte(ops, exp[i]) >= 0 {
         if ts < i {
            n, _ := strconv.Atoi(exp[ts : i])
            ta.AddToken(false, n)
         }
         ta.AddToken(true, int(exp[i]))
         ts = i + 1
      } else if '0' <= exp[i] && exp[i] <= '9' {
      // A digit, continue
      } else {
         fmt.Println("Invalid character at", exp[: i+1])
         return
      }
   }
   if ts < len(exp) {
      n, _ := strconv.Atoi(exp[ts : ])
      ta.AddToken(false, n)
   }
}

func (ta *TokenArray) AddToken(isOperator bool, value int) {
   if ta.NumTokens < len(ta.Tokens) {
      ta.Tokens[ta.NumTokens] = Token { isOperator, value }
      ta.NumTokens++
   } else {
      panic("Token array is full")
   }
}

func (ta *TokenArray) TokenAt(i int) Token {
   return ta.Tokens[i]
}

// Build the binary tree
func (ta *TokenArray) MakeBinaryTree( ) *TokenNode {
   root := &TokenNode { ta.TokenAt(0), nil, nil, nil }
   if root.Node.IsOperator == true {
      panic("Error: first token of the expression cannot be an operator!")
   }
   lastNode := root

   for i := 1; i < ta.NumTokens; i++ {
      tkn := ta.TokenAt(i)
      node := TokenNode { tkn, nil, nil, nil}
      if tkn.IsOperator == false {
         if lastNode.Node.IsOperator == false {
            panic("Error: Near token \"" + string(tkn.Value) + "\"")
         }
         lastNode.RightNode = &node
         node.Parent = lastNode
      } else {	// An operator, search for the node to replace
          p := tkn.Priority()
          replNode := lastNode
          for ; replNode.Parent != nil; {
              if replNode.Parent.Priority() <= p {
                  break
              } else {
                  replNode = replNode.Parent
              }
          }
          // Replace the node
          if replNode.Parent == nil {
              root = &node
          } else {
              replNode.Parent.RightNode = &node
              node.Parent = replNode.Parent
          }
          node.LeftNode = replNode
          replNode.Parent = &node
      }
      lastNode = &node
   }
   return root
}

func PrintTree(node *TokenNode) {
   if node.Node.IsOperator {
      fmt.Printf("%c\n", node.Node.Value)
   } else {
      fmt.Println(node.Node.Value)
   }
   if node.LeftNode != nil {
      fmt.Print("Left node: ")
      PrintTree(node.LeftNode)
   }
   if node.RightNode != nil {
      fmt.Print("Right node: ")
      PrintTree(node.RightNode)
   }   
}

func EvalTreeNode(node *TokenNode) int {
   if node.Node.IsOperator == false {
      return node.Node.Value
   } else {
      a := EvalTreeNode(node.LeftNode)
      b := EvalTreeNode(node.RightNode)
      switch node.Node.Value {
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

func MakeBinaryTree(exp string) *TokenNode {
   tokens := TokenArray { }
   tokens.Tokenise(exp)
   return tokens.MakeBinaryTree()
}
func EvalExpression(exp string) int {
    return EvalTreeNode(MakeBinaryTree(exp))
}

func test(exp string, expected_res int) {
    res := EvalExpression(exp)
    fmt.Printf("Test: %v = %v\n", exp, res)
    if res != expected_res {
        fmt.Println("Error: expected result=", expected_res)
        fmt.Print("Root: ")
        PrintTree(MakeBinaryTree(exp))
    }
}

func main( ) {
   defer func () { 
      error := recover()
      if error != nil {
         fmt.Printf("Panic error: %v \n", error)
      }
   }()

   test("2 + 5", 7)
   test("2 + 5 * 2 * 2", 22)
   test("2 * 5", 10)
   test("2 * 5 + 4", 14)
}
