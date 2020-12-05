# Haskell Study Notes
This doc contains Haskell study notes from the famous [learn you a Haskell](http://learnyouahaskell.com/introduction).

## Getting Started
- Install Haskell Platform on Ubuntu via:
```bash
$ sudo apt-get install haskell-platform
```

## Why Haskell?
- __Haskell is a purely functional programming language.__ 
In purely functional programming languages, a function has no side effects, the only thing that a function can do is calculate something and return it as a result - if a function is called twice with the same parameters, it is bound to have the same result - this is called __referential transparency__. It allows the compiler to reason about the program's behavior, but also allows you to deduce and prove that a function is correct, then build more complex functions by gluing simple functions together.

- __Haskell is lazy.__
Haskell won't execute functions and calculate things unless it is forced to show you a result.

- __Haskell is statically typed.__
Haskell uses a very good type system that has type reference, you don't have to explicitly label the type of variable in your code; also allowing your code to be more general. If a function you make takes two parameters and adds them together and you don't explicitly state their type, the function will work on any two parameters that act like numbers.

- __Haskell is elegant and concise.__
Because it uses a lot of high level concepts, Haskell programs are usually shorter, which means, compared to longer code, it is easier to maintain and less chance of bugs.