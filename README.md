# Golang Notes

This notes come from this resources:

- [A Tour of Go](https://go.dev/tour)
- [Ultimate Go Tour](https://tour.ardanlabs.com/)
- [Fulltimegodev](https://fulltimegodev.com/)

Thanks for let me learn this awesome language, all credits for them.

# Table of Contents

- [Basics](#basics)
  - [Variables](#variables)
- [Methods and Interfaces](#methods-and-interfaces)
  - [Methods](#methods)
  - [Methods are functions](#methods-are-functions)
  - [Methods continued](#methods-continued)
  - [Pointer receivers](#pointer-receivers)
  - [Pointers and functions](#pointers-and-functions)
  - [Methods and pointer indirection](#methods-and-pointer-indirection)
  - [Methods and pointer indirection 2](#methods-and-pointer-indirection-2)
  - [Choosing a value or pointer receiver](#choosing-a-value-or-pointer-receiver)
  - [Interfaces](#interfaces)
  - [Interfaces are implemented implicitly](#interfaces-are-implemented-implicitly)
  - [Interface values](#interface-value)
  - [Interface values with nil underlying values](#interface-values-with-nil-underlying-values)
  - [Nil interface values](#nil-interface-values)
  - [The empty interface](#the-empty-interface)
  - [Type assertions](#type-assertions)
  - [Type switches](#type-switches)
  - [Stringers](#stringers)
  - [Errors](#errors)
  - [Readers](#readers)
- [Concurrency](#concurrency)
  - [Range and Close](#range-and-close)
  - [Select](#select)
  - [Default Selection](#default-selection)
  - [sync Mutex](#sync-mutex)

# Basics

## Pointers

Pointers serve the purpose of sharing values across program boundaries. There are several types of program boundaries. The most common one is between function calls. There is also a boundary between Goroutines which you have notes for later.

When a Go program starts up, the Go runtime creates a Goroutine. Goroutines are lightweight application level threads with many of the same semantics as operating system threads. Their job is to manage the physical execution of a distinct set of instructions. Every Go program has at least 1 Goroutine that you call the main Goroutine.

Each Goroutine is given its own block of memory called a stack. Each stack starts out as a 2048 byte (2k) allocation. It’s very small, but stacks can grow in size over time.

### Pass By Value

All data is moved around the program by value. This means as data is being passed across program boundaries, each function or Goroutine is given its own copy of the data. There are two types of data you will work with, the value itself (int, string, user) or the value's address. Addresses are data that need to be copied and stored across program boundaries.

The [following code](01-basics/04-pointers/pass-by-value.go) attempts to explain this more.

Output:

```shell
count:  Value Of[ 10 ]  Addr Of[ 0xc000050718 ]
inc1:   Value Of[ 11 ]  Addr Of[ 0xc000050710 ]
count:  Value Of[ 10 ]  Addr Of[ 0xc000050718 ]
inc2:   Value Of[ 0xc000050718 ]        Addr Of[ 0xc000050728 ] Points To[ 11 ]
count:  Value Of[ 11 ]  Addr Of[ 0xc000050718 ]
```

### Notes

- Use pointers to share data.
- Values in Go are always pass by value.
- "Value of", what's in the box. "Address of" (&), where is the box.
- The (\*) operator declares a pointer variable and the "Value that the pointer points to".

## Variables

### Zero Value Concept

Every single value you construct in Go is initialized at least to its zero value state unless you specify the initialization value at construction. The zero value is the setting of every bit in every byte to zero.

This is done for data integrity and it's not free. It takes time to push electrons through the machine to reset those bits, but you should always take integrity over performance.

```go
Zero Values:
Type Initialized Value
Boolean false
Integer 0
Floating Point 0
Complex 0i
String "" (empty string)
Pointer nil
```

### Declare and Initialize

The keyword var can be used to construct values to their zero value state for all types.

```go
var a int
var b string
var c float64
var d bool

fmt.Printf("Variables set to their Zero Value\n")
fmt.Printf("var a int \t %T [%v]\n", a, a)
fmt.Printf("var b string \t %T [%v]\n", b, b)
fmt.Printf("var c float64 \t %T [%v]\n", c, c)
fmt.Printf("var d bool \t %T [%v]\n", d, d)
```

Output:

```go
Variables set to their Zero Value
var a int        int [0]
var b string     string []
var c float64    float64 [0]
var d bool       bool [false]
```

Strings use the UTF8 character set, but are really just a collection of bytes.

A string is a two-word internal data structure in Go:

- The first word represents a pointer to a backing array of bytes
- The second word represents the length or the number of bytes in the backing array
- If the string is set to its zero value state, the the first word is nil and the second word is 0

Using the short variable declaration operator, you can declare, construct, and initialize a value all at the same time.

```go
aa := 10
bb := "hello"
cc := 3.14159
dd := true

fmt.Printf("Declare variable and initialize\n")
fmt.Printf("aa := 10 \t %T [%v]\n", aa, aa)
fmt.Printf("bb := \"hello\" \t %T [%v]\n", bb, bb)
fmt.Printf("cc := 3.14159 \t %T [%v]\n", cc, cc)
fmt.Printf("dd := true \t %T [%v]\n", dd, dd)
```

Output:

```go
Declare variable and initialize
aa := 10         int [10]
bb := "hello"    string [hello]
cc := 3.14159    float64 [3.14159]
dd := true       bool [true]
```

### Conversion vs Casting

Go doesn't have casting, but conversion. Instead of telling the compiler to map a set of bytes to a different representation, the bytes need to be copied to a new memory location for the new representation.

```go
fmt.Printf("Specify type and perform a conversion\n")
aaa := int32(10)
fmt.Printf("aaa := int32(10) %T [%v]\n", aaa, aaa)

Output:
Specify type and perform a conversion
aaa := int32(10) int32 [10]
```

Go does have a package in the standard library called unsafe if you need to perform an actual casting operation. You should really avoid that and be honest with yourself why you are considering using it. Performing a conversion provides the highest level of integrity for these types of operations.

### Notes

- The purpose of all programs and all parts of those programs is to transform data from one form to the other.
- Code primarily allocates, reads and writes to memory.
- Understanding type is crucial to writing good code and understanding code.
- If you don't understand the data, you don't understand the problem.
- You understand the problem better by understanding the data.
- When variables are being declared to their zero value, use the keyword var.
- When variables are being declared and initialized, use the short variable declaration operator.

[Code Example](01-basics/01-variables/variables.go)

⬆️ **[back to top](#table-of-contents)**

# Methods and Interfaces

## Methods

Go does not have classes. However, you can define methods on types.

A method is function with a special receiver argument.

The receiver appears in its own argument list between the `func` keyword and the method name.

In this example, the `Abs` method has a receiver of type `Vertex` named `v`.

[Code Example](15-methods-and-interfaces/01-methods/main.go)

⬆️ **[back to top](#table-of-contents)**

## Methods are functions

A method is just a function with a receiver argument.

Here's `Abs` written as a regular function with no change in functionality.

[Code Example](15-methods-and-interfaces/02-methods-are-functions/main.go)

⬆️ **[back to top](#table-of-contents)**

## Methods continued

You can declare a method on non-struct types, too.

In this example we see a numeric type `MyFloat` with an `Abs` method.

You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as `int`).

[Code Example](15-methods-and-interfaces/03-methods-continued/main.go)

⬆️ **[back to top](#table-of-contents)**

## Pointer receivers

You can declare methods with pointer receivers.

This means the receiver type has the literal syntax `*T` for some type `T`. (Also, `T` cannot itself be a pointer such as `*int`)

For example, the `Scale` method here is defined on `*Vertex`.

Methods with pointer receivers can modify the value to which the receiver points (as `Scale` does here). Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

Try removing the `*` from the declaration of the `Scale` function on line 16 and observe how the program's behavior changes.

With a value receiver, the `Scale` method operates on a copy of the original `Vertex` value. (This is the same behavior as for any other function argument). The `Scale` method must have a pointer receiver to change the `Vertex` value declared in the `main` function.

[Code Example](15-methods-and-interfaces/04-pointer-receivers/main.go)

⬆️ **[back to top](#table-of-contents)**

## Pointers and functions

Here we see the `Abs` and `Scale` methods rewritten as functions.

Again, try removing the `*` from line 16. Can you see why the behavior changes? What else did you need to change for the example to compile?

[Code Example](15-methods-and-interfaces/05-pointers-and-functions/main.go)

⬆️ **[back to top](#table-of-contents)**

## Methods and pointer indirection

Comparing the previous two programs, you might notice that functions with a pointer argument must take a pointer:

```go
var v Vertex
ScaleFunc(v, 5)  // Compile error!
ScaleFunc(&v, 5) // OK
```

while methods with pointer receivers take either a value or a pointer as the receiver when they are called:

```
var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
```

For the statement `v.Scale(5)`, even though `v` is a value and not a pointer, the method with the pointer receiver is called automatically. That is, as a convenience, Go interprets the statement `v.Scale(5)` as `(&v).Scale(5)` since the `Scale` method has a pointer receiver.

[Code Example](15-methods-and-interfaces/06-methods-and-pointer-indirection/main.go)

⬆️ **[back to top](#table-of-contents)**

## Methods and pointer indirection 2

The equivalent thing happens in the reverse direction.

Functions that take a value argument must take a value of that specific type:

```go
var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // Compile error!
```

while methods with value receivers take either a value or a pointer as the receiver when they are called:

```go
var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
```

In this case, the method call `p.Abs()` is interpreted as `(*p).Abs()`

[Code Example](15-methods-and-interfaces/07-methods-and-pointer-indirection-2/main.go)

⬆️ **[back to top](#table-of-contents)**

## Choosing a value or pointer receiver

There are two reasons to use a pointer receiver.

The first is so that the method can modify the value that its receiver points to.

The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.

In this example, both `Scale` and `Abs` are methods with receiver type `*Vertex`, even though the `Abs` method needn't modify its receiver.

In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both. (We'll see why over the next few pages.)

[Code Example](15-methods-and-interfaces/08-choosing-a-value-or-pointer-receiver/main.go)

⬆️ **[back to top](#table-of-contents)**

## Interfaces

An interface type is defined as a set of method signatures.

A value of interface type can hold any value that implements those methods.

**Note:** There is an error in the example code on line 23. `Vertex` (the value type) doesn't implement `Abser` because the `Abs` method is defined only on `*Vertex` (the pointer type).

[Code Example](15-methods-and-interfaces/09-interfaces/main.go)

⬆️ **[back to top](#table-of-contents)**

## Interfaces are implemented implicitly

A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.

Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.

[Code Example](15-methods-and-interfaces/10-interfaces-are-implemented-implicitly/main.go)

⬆️ **[back to top](#table-of-contents)**

## Interface values

Under the hood, interface values can be thought of as a tuple of a value and a concrete type:

```go
(value, type)
```

An interface value holds a value of a specific underlying concrete type.

Calling a method on an interface value executes the method of the same name on its underlying type.

[Code Example](15-methods-and-interfaces/11-interface-values/main.go)

⬆️ **[back to top](#table-of-contents)**

## Interface value with nil underlying values

If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.

In some languages this would trigger a null pointer exception, but in Go it is common to write methods that gracefully handle being called with a nil receiver (as with the method `M` in this example.)

Note that an interface value that holds a nil concrete value is itself non-nil.

[Code Example](15-methods-and-interfaces/12-interface-values-with-nil-underlying-values/main.go)

⬆️ **[back to top](#table-of-contents)**

## Nil interface values

A nil interface value holds neither value nor concrete type.

Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.

[Code Example](15-methods-and-interfaces/13-nil-interface-values/main.go)

⬆️ **[back to top](#table-of-contents)**

## The empty interface

The interface type that specifies zero methods is known as the `empty interface`:

```go
interface{}
```

An empty interface may hold values of any type. (Every type implements at least zero methods.)

Empty interfaces are used by code that handles values of unknown type. For example, `fmt.Print` takes any number of arguments of type `interface{}`.

[Code Example](15-methods-and-interfaces/14-the-empty-interface/main.go)

⬆️ **[back to top](#table-of-contents)**

## Type assertions

A type assertion provides access to an interface value's underlying concrete value.

```go
t := i.(T)
```

This statement asserts that the interface value `i` holds the concrete type `T` and assigns the underlying `T` value to the variable `t`.

If `i` does not hold a `T`, the statement will trigger a panic.

To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.

```go
t, ok := i.(T)
```

If `i` holds a `T`, the `t` will be the underlying value and `ok` will be true.

If not, `ok` will be false and `t` will be the zero value of type `T`, and no panic occurs.

Note the similarity between this syntax and that of reading from a map.

[Code Example](15-methods-and-interfaces/15-type-assertions/main.go)

⬆️ **[back to top](#table-of-contents)**

## Type switches

A type switch is a construct that permits several type assertions in series.

A type switch is like a regular switch statement, but the cases in a type switch specify types (not values), and those values are compared against the type of the value held by the given interface value.

```go
switch v := i.(type) {
    case T:
        // here v has type T
    case S:
        // here v has type S
    default:
        // no match; here v has the same type as i
}
```

The declaration in a type switch has the same syntax as a type assertion i.(T), but the specific type `T` is replaced with the keyword `type`.

This switch statement tests whether the interface value `i` holds a value of type `T` or `S`. In each of the `T` and `S` cases, the variable `v` will be of type `T` or `S` respectively and hold the value held by `i`. In the default case (where there is no match), the variable `v` is of the same interface type and value as `i`.

[Code Example](15-methods-and-interfaces/16-type-switches/main.go)

⬆️ **[back to top](#table-of-contents)**

## Stringers

One of the most ubiquitous interfaces is `Stringer` defined by the `fmt` package.

```go
type Stringer interface {
    String() string
}
```

A `Stringer` is a type that can describe itself as a string. The `fmt` package (and many others) look for this interface to print values.

[Code Example](15-methods-and-interfaces/17-stringers/main.go)

⬆️ **[back to top](#table-of-contents)**

## Errors

Go programs express error state with `error` values.

The `error` type is a built-in interface similar to `fmt.Stringer`:

```go
type error interface {
    Error() string
}
```

(As with `fmt.Stringer`, the `fmt` package looks for the `error` interface when printing values.)

Functions often return an `error` value, and calling code should handle errors by testing whether the error equals `nil`.

```go
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
```

A nil `error` denotes success; a non-nil `error` denotes failure.

[Code Example](15-methods-and-interfaces/18-errors/main.go)

⬆️ **[back to top](#table-of-contents)**

## Readers

The `io` package specifies the `io.Reader` interface, which represents the read end of a stream of data.

The Go standard library contains many implementations of this interface, including files, network connections, compressors, ciphers, and others.

The `io.Reader` interface has a `Read` method:

```go
func (T) Read(b []byte) (n int, err error)
```

`Read` populates the given byte slice with data and returns the number of bytes populated and an error value. It returns an `io.EOF` error when the stream ends.

The example code creates a `strings.Reader` and consumes its output 8 bytes at a time.

[Code Example](15-methods-and-interfaces/19-readers/main.go)

⬆️ **[back to top](#table-of-contents)**

# Concurrency

## Range and Close

A sender can `close` a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after

`v, ok := <-ch`

`ok` is `false` if there are no more values to receive and the channel is closed.

The loop `for i := range c` receives values from the channel repeatedly until it is closed.

Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.

Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.

[Code Example](14-concurrency/04-range-and-close/main.go)

⬆️ **[back to top](#table-of-contents)**

## Select

The `select` statement lets a goroutine wait on multiple communication operations.

A `select` blocks until one of its cases can run, the it executes that case. It chooses one at random if multiple are ready.

[Code Example](14-concurrency/05-select/main.go)

⬆️ **[back to top](#table-of-contents)**

## Default Selection

The `default` case in a `select` is run if no other case is ready.

Use a `default` case to try a send or receive without blocking:

```go
select {
case i := <-c:
    // use i
default:
    // receiving from c would block
}
```

[Code Example](14-concurrency/06-default-selection/main.go)

⬆️ **[back to top](#table-of-contents)**

## sync Mutex

We've seen how channels are great for communication among goroutines.

But what if we don't need communication? What if we just want to make sure only one goroutine can access a variable at a time to avoid conflicts?

This concept is called `mutual exclusion`, and the conventional name for the data structure that provides it is `mutex`.

Go' standard library provides mutual exclusion with `sync.Mutex` and its two methods:

- `Lock`
- `Unlock`

We can define a block of code to be executed in mutual exclusion by surrounding it with a call to `Lock` and `Unlock` as shown on the `Inc` method.

We can also use `defer` to ensure the mutex will be unlocked as in the `Value` method.

[Code Example](14-concurrency/07-sync-mutex/main.go)

⬆️ **[back to top](#table-of-contents)**
