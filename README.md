# Golang Notes

This notes come from 2 courses:

- [A Tour of Go](https://go.dev/tour)
- [Fulltimegodev](https://fulltimegodev.com/)

Thanks to Golang Team and AnthonyGG for let me learn this awesome language, all credits for them.

# Table of Contents

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
- [Concurrency](#concurrency)
  - [Range and Close](#range-and-close)
  - [Select](#select)
  - [Default Selection](#default-selection)
  - [sync Mutex](#sync-mutex)

# Methods and Interfaces

## Methods

Go does not have classes. However, you can define methods on types.

A method is function with a special _receiver_ argument.

The receiver appears in its own argument list between the `func` keyword and the method name.

In this example, the `Abs` method has a receiver of type `Vertex` named `v`.

[Code Example](15-methods-and-interfaces/01-methods/main.go)

## Methods are functions

A method is just a function with a receiver argument.

Here's `Abs` written as a regular function with no change in functionality.

## Methods continued

You can declare a method on non-struct types, too.

In this example we see a numeric type `MyFloat` with an `Abs` method.

You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as `int`).

[Code Example](15-methods-and-interfaces/03-methods-continued/main.go)

## Pointer receivers

You can declare methods with pointer receivers.

This means the receiver type has the literal syntax `*T` for some type `T`. (Also, `T` cannot itself be a pointer such as `*int`)

For example, the `Scale` method here is defined on `*Vertex`.

Methods with pointer receivers can modify the value to which the receiver points (as `Scale` does here). Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

Try removing the `*` from the declaration of the `Scale` function on line 16 and observe how the program's behavior changes.

With a value receiver, the `Scale` method operates on a copy of the original `Vertex` value. (This is the same behavior as for any other function argument). The `Scale` method must have a pointer receiver to change the `Vertex` value declared in the `main` function.

[Code Example](15-methods-and-interfaces/04-pointer-receivers/main.go)

## Pointers and functions

Here we see the `Abs` and `Scale` methods rewritten as functions.

Again, try removing the `*` from line 16. Can you see why the behavior changes? What else did you need to change for the example to compile?

[Code Example](15-methods-and-interfaces/05-pointers-and-functions/main.go)

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

## Choosing a value or pointer receiver

There are two reasons to use a pointer receiver.

The first is so that the method can modify the value that its receiver points to.

The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.

In this example, both `Scale` and `Abs` are methods with receiver type `*Vertex`, even though the `Abs` method needn't modify its receiver.

In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both. (We'll see why over the next few pages.)

[Code Example](15-methods-and-interfaces/08-choosing-a-value-or-pointer-receiver/main.go)

## Interfaces

An interface type is defined as a set of method signatures.

A value of interface type can hold any value that implements those methods.

**Note:** There is an error in the example code on line 23. `Vertex` (the value type) doesn't implement `Abser` because the `Abs` method is defined only on `*Vertex` (the pointer type).

[Code Example](15-methods-and-interfaces/09-interfaces/main.go)

## Interfaces are implemented implicitly

A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.

Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.

# Concurrency

## Range and Close

A sender can `close` a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after

`v, ok := <-ch`

`ok` is `false` if there are no more values to receive and the channel is closed.

The loop `for i := range c` receives values from the channel repeatedly until it is closed.

Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.

Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.

[Code Example](14-concurrency/04-range-and-close/main.go)

## Select

The `select` statement lets a goroutine wait on multiple communication operations.

A `select` blocks until one of its cases can run, the it executes that case. It chooses one at random if multiple are ready.

[Code Example](14-concurrency/05-select/main.go)

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

## sync Mutex

We've seen how channels are great for communication among goroutines.

But what if we don't need communication? What if we just want to make sure only one goroutine can access a variable at a time to avoid conflicts?

This concept is called _mutual exclusion_, and the conventional name for the data structure that provides it is _mutex_.

Go' standard library provides mutual exclusion with `sync.Mutex` and its two methods:

- `Lock`
- `Unlock`

We can define a block of code to be executed in mutual exclusion by surrounding it with a call to `Lock` and `Unlock` as shown on the `Inc` method.

We can also use `defer` to ensure the mutex will be unlocked as in the `Value` method.

[Code Example](14-concurrency/07-sync-mutex/main.go)
