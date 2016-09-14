# go_evaluate
A numerical expression parser and evaluator using GO

1. Expression with integers, + - * / operators. For example, 2 * 5 + 20
2. Tokenize the expression.
3. Tokens (operators and operands) have priorities.
   a number has priority 0
   + - has priority 1
   * / has priroity 2
4. Build a binary tree. For example, 2 * 5 + 20, the tree is
      +
     / \
    *   20
   / \
  2   5
5. Evaluate
   