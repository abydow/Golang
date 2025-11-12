# 🔑 Core Terminology {#core-terminology}

## **Package**
A package is a collection of source code files organized together. Every Go file belongs to exactly one package. The special `package main` is required for executable programs and serves as the entry point.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Gophers!")
}
```

## **Import**
The `import` statement allows you to include packages and access their exported functions, variables, and types.

```go
import (
    "fmt"           // Standard library
    "math"
    "github.com/user/mypackage"  // Third-party package
)
```

## **Identifier**
An identifier is a name used to identify variables, functions, types, or other user-defined items. Identifiers must start with a letter or underscore and can contain letters, digits, and underscores.

```go
var age int              // Valid identifier
var myVariable string    // Valid identifier
var _private bool        // Valid (starts with underscore)
// var 2invalid := 10    // Invalid (starts with digit)
```

## **Keyword**
Go has exactly **25 keywords** that cannot be used as identifiers. Some key keywords include:

```
break, case, chan, const, continue, default, defer, else, fallthrough,
for, func, go, goto, if, import, interface, map, package, range, return,
select, struct, switch, type, var
```
visit https://go.dev/ref/spec#Keywords for more info...

```go
// Exported
func PublicFunction() string {
    return "I am accessible everywhere"
}

// Unexported
func privateFunction() string {
    return "I am private to this package"
}
```

---

