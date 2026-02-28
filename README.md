<p align="center">
  <h1 align="center">Turtle</h1>
  <p align="center">Fast, reliable, and delightfully simple programming language for modern development.</p>
  <p align="center">
    <a href="https://turtle-interpreter.vercel.app">Website</a> &middot;
    <a href="https://turtle-interpreter.vercel.app/#playground">Online Playground</a> &middot;
    <a href="https://turtle-interpreter.vercel.app/#benchmarks">Benchmarks</a> &middot;
    <a href="https://github.com/Abbas-Askari/turtle-website">Website Source</a>
  </p>
</p>

---

## What is Turtle?

Turtle is a dynamically-typed programming language with clean syntax combining the best of Go and JavaScript. It compiles to portable bytecode that runs on an efficient Go-powered virtual machine with async I/O, a rich standard library, and a simple concurrency model.

```javascript
import "net";
import "json";

let server = new net.Server();

server.get("/hello", (req, res) => {
    res.writeHeader("Content-Type", "application/json");
    res.writeStatus(200);
    res.write(json.stringify({ message: "Hello from Turtle!" }));
    res.end();
});

server.listen(3000);
print "Server running at http://localhost:3000";
```

## Features

- **Familiar Syntax** — Clean syntax that's easy to pick up if you know Go or JavaScript
- **Bytecode Compilation** — 3-stage compiler (lexer, parser, codegen) produces portable bytecode
- **Single-Threaded Async** — Node-like async I/O powered by Go goroutines, mutexes & channels under the hood
- **Rich Standard Library** — JSON, HTTP, file system, TCP networking, OS operations, and more built-in
- **Prototypal Inheritance** — Flexible object system with prototypes
- **Closures** — First-class functions with full lexical closure support
- **Safe Mode** — Sandboxed execution that blocks access to OS, FS, and networking libraries
- **VM Safety** — Running on a virtual machine provides isolation and safety guarantees

## Quick Start

### Build from source

```bash
go build -o turtle .
```

### Run a file

```bash
./turtle main.turtle
```

### Flags

```
./turtle [-v] [-s] <filename>
```

| Flag | Description |
|------|-------------|
| `-v` | Verbose/debug mode — prints tokens, AST, and bytecode |
| `-s` | Safe mode — blocks `os`, `fs`, `net`, `http`, `tcp` imports |

### Docker

```bash
docker build -t turtle .
docker run turtle
```

## Language Tour

### Variables & Functions

```javascript
let name = "Turtle";
let version = 1;

func greet(who) {
    print "Hello, " + who + "!";
}

greet(name); // Hello, Turtle!
```

### Arrow Functions

```javascript
let double = (x) => x * 2;
let add = (a, b) => a + b;

print double(21); // 42
print add(3, 4);  // 7
```

### Closures

```javascript
func makeCounter(start) {
    let count = start;
    func increment() {
        print "Count: " + (count = count + 1);
    }
    return increment;
}

let counter = makeCounter(0);
counter(); // Count: 1
counter(); // Count: 2
counter(); // Count: 3
```

### Prototypes

```javascript
func Animal(name, sound) {
    this.name = name;
    this.sound = sound;
}

Animal.prototype.speak = () => {
    print this.name + " says " + this.sound;
};

let dog = new Animal("Dog", "Woof");
dog.speak(); // Dog says Woof
```

### Arrays

```javascript
Array.push = (x) => {
    return this + [x];
};

Array.map = (fn) => {
  print fn;
  let res = [];
  for let i = 0; i < this.length; i = i + 1 {
    res = res.push(fn(this[i], i));
  }
  return res;
};

Array.filter = (fn) => {
  let res = [];
  for let i = 0; i < this.length; i = i + 1 {
    if fn(this[i], i) {
      res = res.push(this[i]);
    }
  }
  return res;
};

let nums = [1, 2, 3, 4, 5];

let doubled = nums.map((x, _) => x * 2);
let evens = nums.filter((x, _) => x % 2 == 0);

print doubled; // [2, 4, 6, 8, 10]
print evens;   // [2, 4]
```

### Async I/O

```javascript
import "async";

async.setTimeout(() => {
    print "Delayed by 500ms";
}, 500);

async.setInterval(() => {
    print "Every 200ms";
}, 200);
```

### HTTP Server

```javascript
import "net";
import "json";
import "fs";

let server = new net.Server();

server.get("/", (req, res) => {
    res.writeHeader("Content-Type", "text/plain");
    res.writeStatus(200);
    res.write("Welcome to Turtle!");
    res.end();
});

server.post("/data", (req, res) => {
    let body = json.parse(req.body());
    res.writeHeader("Content-Type", "application/json");
    res.writeStatus(200);
    res.write(json.stringify(body));
    res.end();
});

server.listen(8080);
```

### File System

```javascript
import "fs";
import "json";

let f = fs.open("data.json", "r");
let content = f.read(f.length());
f.close();

let data = json.parse(content[0]);
print data;
```

### TCP & PostgreSQL

Turtle's standard library is powerful enough to implement a [full PostgreSQL driver](scripts/pg.turtle) over raw TCP — no package manager needed:

```javascript
import "tcp";

tcp.connect("localhost:5432", (socket, err) => {
    if err != nil {
        print "Connection failed: " + err;
        return;
    }
    print "Connected!";
    // Full PG wire protocol implementation in pure Turtle
});
```

### Modules

```javascript
// math.turtle
func square(x) {
    return x * x;
}
exports.square = square;

// main.turtle
import "math";
print math.square(5); // 25
```

## Standard Library

| Module   | Description |
|----------|-------------|
| `json`   | `parse()` and `stringify()` for JSON handling |
| `fs`     | File system — `open`, `read`, `write`, `close`, `exists`, `remove`, and more |
| `http`   | HTTP client — `request()` with GET, POST, and other methods |
| `net`    | HTTP server with routing (`get`, `post`, etc.) |
| `tcp`    | Raw TCP client/server connections |
| `os`     | OS operations — environment variables, process execution, `exit()` |
| `async`  | `setTimeout()` and `setInterval()` |
| `ascii`  | Character code utilities |

## Architecture

```
Source Code (.turtle)
        │
        ▼
    ┌────────┐
    │ Lexer  │   Tokenization
    └────┬───┘
         │
         ▼
    ┌────────┐
    │ Parser │   AST Generation
    └────┬───┘
         │
         ▼
  ┌──────────┐
  │ Compiler │   Bytecode Generation
  └────┬─────┘
       │
       ▼
    ┌──────┐
    │  VM  │     Stack-based Bytecode Execution
    └──────┘
```

The Go-powered VM handles scheduling, async I/O via goroutines, and moves concurrency complexity away from the user.

## Benchmarks

Turtle performs competitively against Go, Node.js, and Python across HTTP forwarding, database operations, and audio processing workloads. The website includes interactive benchmark charts comparing average response times across different load sizes.

Average response times in milliseconds (lower is better). Values are averages across 5 runs.

#### Forward Requests — HTTP proxy with caching and JSON parsing

| Requests | Go | Node.js | Python | **Turtle** |
|----------|---:|--------:|-------:|-----------:|
| 500      | 137 | 244 | 233 | **124** |
| 1,000    | 133 | 182 | 320 | **130** |
| 10,000   | 1,636 | 6,674 | 8,149 | **1,489** |

#### Database Operations — PostgreSQL queries

| Requests | Go | Node.js | Python | **Turtle** |
|----------|---:|--------:|-------:|-----------:|
| 500      | **138** | **132** | 143 | 189 |
| 1,000    | **297** | 231 | 264 | 550 |
| 5,000    | **1,608** | 3,323 | 3,501 | 9,042 |

#### Audio Operations — File I/O and processing

| Files | Go | Node.js | Python | **Turtle** |
|-------|---:|--------:|-------:|-----------:|
| 4     | 4,898 | 4,719 | **4,551** | 5,558 |
| 8     | 7,222 | 7,364 | 7,653 | **7,262** |
| 12    | **9,466** | 9,943 | 9,663 | 10,316 |

#### Combined Benchmark

| | Go | Node.js | Python | **Turtle** |
|-|---:|--------:|-------:|-----------:|
| Avg | **6,227** | 6,289 | 6,393 | 6,367 |

See the [interactive benchmark charts](https://turtle-interpreter.vercel.app/#benchmarks) for more detail.

## Try Online

No installation needed — write and run Turtle code directly in the browser at the [online playground](https://turtle-interpreter.vercel.app/#playground). The playground runs in safe mode on an AWS Lambda function powered by Turtle itself.

## License

© 2025 Turtle Lang. All rights reserved.
