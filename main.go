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
// Run test without using channels method
func test0(exp string, expected_res int) {
    res, root := EvalExpression0(exp)
    fmt.Printf("Test: %v = %v\n", exp, res)
    if res != expected_res {
        fmt.Println("Error: expected result =", expected_res)
        fmt.Print("Root: ")
        PrintTree(root)
    }
}
// Run test using channels method
func test(exp string, expected_res int) {
    res, root := EvalExpression(exp)
    fmt.Printf("Test: %v = %v\n", exp, res)
    if res != expected_res {
        fmt.Println("Error: expected result =", expected_res)
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

   fmt.Println("**** Tests using channels")
   test("2 + 5", 7)
   test("2 + 5 * 2 * 2", 22)
   test("2 * 5", 10)
   test("2 * 5 + 4", 14)
   test("2 * (5 + 4)", 18)
   test("2 * (5 + 4) + 3 * 4", 30)
   test("(2 + 3) * 5", 25)
   test("((3 + 5) * 2 + 3) * 5", 95)

   fmt.Println("\n**** Tests without channels")
   test0("2 + 5", 7)
   test0("2 + 5 * 2 * 2", 22)
   test0("2 * 5", 10)
   test0("2 * 5 + 4 + 4", 18)
}