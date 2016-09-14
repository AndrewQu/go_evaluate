package main

import (
   "strconv"
   "strings"
)
// Represents a tokens array
type TokenArray struct {
   NumTokens int;
   Tokens [] Token
}
// Add a token into the tokens array
func (ta *TokenArray) AddToken(isOperator bool, value int) {
   if ta.NumTokens < len(ta.Tokens) {
      ta.Tokens[ta.NumTokens] = Token { isOperator, value }
      ta.NumTokens++
   } else {
      panic("Token array is full")
   }
}
// Tokenise an expression into an tokens array
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
         panic(strings.Join([] string { "Invalid character at", exp[: i+1]}, ""))
      }
   }
   if ts < len(exp) {
      n, _ := strconv.Atoi(exp[ts : ])
      ta.AddToken(false, n)
   }
}
// Build the binary tree from the tokens array
func (ta *TokenArray) MakeBinaryTree( ) *TokenNode {
   root := &TokenNode { ta.Tokens[0], nil, nil, nil }
   if root.TheToken.IsOperator == true {
      panic("Error: first token of the expression cannot be an operator!")
   }
   lastNode := root

   for i := 1; i < ta.NumTokens; i++ {
      tkn := ta.Tokens[i]
      node := TokenNode { tkn, nil, nil, nil}
      if tkn.IsOperator == false {
         if lastNode.TheToken.IsOperator == false {
            panic("Error: Near token \"" + string(tkn.Operand()) + "\"")
         }
         lastNode.RightNode = &node
         node.Parent = lastNode
      } else {	// An operator, search for the node to replace
          p := node.Priority()
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
// Evaluate an expression. This is the UI function
func EvalExpression0(exp string) (int, *TokenNode) {
    tokens := TokenArray { }
    tokens.Tokenise(exp)
    root := tokens.MakeBinaryTree()
    return EvalTreeNode(root), root
}