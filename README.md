# tlang

Trying to create an interpreter for my custom dynamically typed language

An algamation of different features from different languages mashed into one language

## Variables

Define variables with `let`

```
let myVar = 20
```

Re assign as you would in any C lang

```
myVar = 24
```

And some operators

```
myVar++
myVar--
myVar += 1
myVar -= 1
myVar /= 2
myVar *= 2
```

### Data types

- string `"hello world"`
- number `1337` (all numbers are actually 64 bit floats)
- bool `true` `false`
- dynamic arrays `[1, 2, 3]`
- maps `{ hello = "world" }`
- null `nil`

## Functions

All functions are variables, define them with `let` and `fn`

```
let sum = fn(a, b) {
    return a + b
}
```

Call functions like this

```
let short = sum(20, 42)

let longer = sum(20, b = 42)

let longest = sum(
    a = 20,
    b = 42,
)
```

We have closures, they are fucking awesome

```
let createMultiply = fn(by) {
    return fn(num) {
        num * by
    }
}

let multiplyBy3 = createMultiply(3)
let result = multiplyBy3(3) // 9
```

Functions are first-class served

```
let callPassed = fn(f) {
    f()
}

callPassed(fn() {
    // what's up
})

```

Immediately invoke functions for quick computations

```
let result = fn() {
    return 2 * 2
}()
```

## Comments

As you already have seen, define comments with `//`

```
// I am a comment, wow, pretty cool
```

## IFs

Write if statements like you would in rust

```
let cute = true
let evenMoreCute = true

if cute {
    // ...
} else if evenMoreCute {
    // ...
} else {
    // :(
}
```

Same line `if true return` are **NOT** supported
