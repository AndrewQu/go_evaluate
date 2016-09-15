package main

import (
   "strconv"
   "strings"
)
// Tokenise an expression and pipe tokens to the given channel
func Tokenise(exp string, ch (chan Token) ) {
   ops := "+-*/()"
   exp = strings.Replace(exp, " ", "", len(exp))
   defer func() { close(ch) }()

   ts := 0 // Start of the next token
   for i := 0; i < len(exp); i++ {
      if strings.IndexByte(ops, exp[i]) >= 0 {
         if ts < i {
            n, _ := strconv.Atoi(exp[ts : i])
            ch <- Token { false, n }
         }
         ch <- Token { true, int(exp[i]) }
         ts = i + 1
      } else if '0' <= exp[i] && exp[i] <= '9' {
      // A digit, continue
      } else {
         panic(strings.Join([] string { "Invalid character at", exp[: i+1]}, ""))
      }
   }
   if ts < len(exp) {
      n, _ := strconv.Atoi(exp[ts : ])
      ch <- Token { false, n }
   }
}
// Build the binary tree from the tokens array
func MakeBinaryTree(ch (chan Token), ch_root (chan *TokenNode) ) {
   defer func() { close(ch_root) }()    // Make sure the chan is always closed

   tkn, ok := <- ch
   if !ok { return }

   root := &TokenNode { tkn, nil, nil, nil }
   if root.TheToken.IsOperator == true && root.IsLeftBracket() == false {
      panic("Error: first token of the expression cannot be an operator!")
   }
   lastNode := root

   for tkn = range ch {
      node := TokenNode { tkn, nil, nil, nil}
      if tkn.IsOperator == false {
          if lastNode.TheToken.IsOperator == false {
             panic("Error: Near token \"" + string(tkn.Operand()) + "\"")
          }
          if lastNode.IsLeftBracket() {
              lastNode.LeftNode = &node
          } else {
              lastNode.RightNode = &node
          }
          node.Parent = lastNode
          lastNode = &node
      } else {	// An operator, search for the node to replace
          op := tkn.Operator()
          if op == '(' {
              lastNode.RightNode = &node
              node.Parent = lastNode
              lastNode = &node
          } else if op == ')' {
              lbrktNode := lastNode   // Find the left bracket
              for ; lbrktNode.Parent != nil; {
                  if lbrktNode.IsLeftBracket() {
                      break
                  }
                  lbrktNode = lbrktNode.Parent
              }
              if lbrktNode.IsLeftBracket() == false {
                  panic("Inbalanced left bracket.")
              }
              v := EvalTreeNode(lbrktNode.RightNode)
              lbrktNode.TheToken = Token { false, v }
              lbrktNode.LeftNode = nil
              lbrktNode.RightNode = nil
              lastNode = lbrktNode
          } else { // + - * /
              p := node.Priority()
              replNode := lastNode
              for ; replNode.Parent != nil && !replNode.Parent.IsLeftBracket(); {
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
              lastNode = &node
          }
      }
   }
   ch_root <- root
}
// Evaluate an expression. This is the UI function
func EvalExpression(exp string) (int, *TokenNode) {
    ch := make( chan Token, 10)         // Channel for tokens
    ch_root := make(chan *TokenNode)    // Channel for the root of the binary tree

    go MakeBinaryTree(ch, ch_root)  // Start the thread that receives and process tokens
    // When all done, the root of the binary tree will be sent through ch_root

    Tokenise(exp, ch)   // Tokenise the expression and send tokens to ch

    root, ok := <- ch_root  // Wait for tree root from MakeBinaryTree thread
    if ok {
        return EvalTreeNode(root), root 
    } else {
        return 0, nil
    }
}