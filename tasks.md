## TODO

- [x] Change the single expression parsing nature to list of statements parsing nature
  - [x] Change to list of statements
  - [x] Add print statement
- [x] Add seperate compiler from the parser
- add string support

  - in lexer to define string literals
  - in objects and allow operations on strings

### Add variables

- [x] declaration statements
- [x] assignment expressions
- [x] OpLoadGlobal & OpSetGlobal
- [x] Scopes and local variables

### Add internal types

- [x] Add numbers
- [x] Add Nil
- [x] Add Strings
- [x] Add Boolean (In lox code)
- [ ] Add Arrays
- [ ] Add Maps
- [ ] Unify operation in Object module.

  - Will add functions like Object.Add(a, b Object) Object

### Add Operators and expressions

- [x] ==, !=, >=, <=, >, <
- [x] +, -
- [x] \*, /
- [x] &&, ||
- [x] Unary: +, -, !
- [x] Groupings "()"

### Overall features and fixes

- [ ] add input statement `input x` will write the user input to x
- [ ] Logical operators short curciting
  - `a = "abc" || 123` will store `"abc"` in a
  - `a = "abc" && 123` will store `123` in a
- [x] add if statements
- [x] add single line comments
- [x] add Equality comparisons
- [x] add for statements
- [ ] Add break and continue
- [x] Fix operator precedence

### better errors

- Add good error showing in lexer
- Add good error showing in parser
- Add good error showing in VM

### Dev quality of life.

- [x] add test.lox to watch
- [ ] add individual tests for features.
