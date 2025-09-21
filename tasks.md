## TODO

- [x] Change the single expression parsing nature to list of statements parsing nature
  - [x] Change to list of statements
  - [x] Add print statement
- [x] Add seperate compiler from the parser
- add string support

  - in lexer to define string literals#
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

  - Will add functions like Object.Add(a, b) instead of a.Add(b)

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
- [x] Add break and continue
- [ ] Fix break and continue bug which can occur if some other part of the AST emits bytecode of OpBreak 2 times as operands.
- [x] Fix operator precedence
- [x] Add functions
- [x] Add return statements
- [ ] Add closures
- [ ] Add Maps
  - [ ] Test with closures
- [ ] Add Arrays
  - [ ] Test returning an array of closures closing over a loop variable
- [ ] Add Prototypal inheritance
- [ ] Add File Modules with imports and exports

### better errors

- Add good error showing in lexer
- Add good error showing in parser
- Add good error showing in VM

### Dev quality of life.

- [x] add test.lox to watch
- [ ] add individual tests for features.

### 22 sept tasks

- [ ] add more `fs` functions
  - MVP (8 funcs): readFile, writeFile, appendFile, stat, exists, readDir, mkdir, remove
  - Streaming (4–5 funcs): open, close, read, write, seek
- [ ] add `fetch` under http
- [ ] add `TCP` support
  - write `psql` driver?
- [ ] add `setInterval` and `setTimeout`
- [ ] runtime (2–3 funcs: print, exit, stacktrace)

### Speed Test

- Api forwarding with caching and cache dumping
  - Tests Http (serving, requesting), Json parsing, and Map writes.
  - Cache dumping tests FS and async I/O
- Dump Cache to PG?
- Cache x 100
- Dump the Cache x 100 to FS
- File streaming
  - Tests FS + Heep serving
    All steps except 1 will have the serve answering constant calls to /please-say POST calls which will reply with the bdoy of the request to measure RPS and concurrency.
