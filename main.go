package main

import (
   "fmt"
)
func PrintTree(node *TokenNode) {
   if node.TheToken.IsOperator {
      fmt.Printf("%c\n", node.TheToken.Operator())
   } else {
      fmt.Println(node.TheToken.Operand())
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

func test(exp string, expected_res int) {
    res, root := EvalExpression0(exp)
    fmt.Printf("Test: %v = %v\n", exp, res)
    if res != expected_res {
        fmt.Println("Error: expected result=", expected_res)
        fmt.Print("Root: ")
        PrintTree(root)
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
