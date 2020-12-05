# Haskell Study Notes
This doc contains Haskell study notes from the famous [learn you a Haskell](http://learnyouahaskell.com/introduction).

## Getting Started
- Install Haskell Platform on Ubuntu via:
```bash
$ sudo apt-get install haskell-platform
```

It comes with:
- [Glasgow Haskell Compiler](https://wiki.haskell.org/GHC/GHCi)
- Cabal build system
- Stack tool for developing projects
- support for profiling and code coverage analysis
- 35 core & widely-used packages

## Why Haskell?
- __Haskell is a purely functional programming language.__ 
In purely functional programming languages, a function has no side effects, the only thing that a function can do is calculate something and return it as a result - if a function is called twice with the same parameters, it is bound to have the same result - this is called __referential transparency__. It allows the compiler to reason about the program's behavior, but also allows you to deduce and prove that a function is correct, then build more complex functions by gluing simple functions together.

- __Haskell is lazy.__
Haskell won't execute functions and calculate things unless it is forced to show you a result.

- __Haskell is statically typed.__
Haskell uses a very good type system that has type reference, you don't have to explicitly label the type of variable in your code; also allowing your code to be more general. If a function you make takes two parameters and adds them together and you don't explicitly state their type, the function will work on any two parameters that act like numbers.

- __Haskell is elegant and concise.__
Because it uses a lot of high level concepts, Haskell programs are usually shorter, which means, compared to longer code, it is easier to maintain and less chance of bugs.

### Functions
- When a function doesn't take parameters, it is considered to be a _definition_ or _name_.

### Lists
- List is a homogenous data structure - it stores data of the same type.
- Common list operations:

- List concat with `++`
    ```haskell
    "ghci> Hello" ++ " " ++ "World" ++ "!"
    Hello World!
    ```
- List append with `:`
    ```haskell
    ghc1> 0:[1, 2, 3]
    [0,1,2,3]

    ghc1> 1:2:3:[]
    [1,2,3]
    ```
- Get element via index
    ```haskell
    ghc1> [1, 2, 3] !! 2
    3
    ```
- `head` takes a list and returns its head.
    ```haskell
    ghci> head [5,4,3,2,1]  
    5  
    ```

- `tail` takes a list and returns its tail. In other words, it chops off a list's head.
    ```haskell
    ghci> tail [5,4,3,2,1]
    [4,3,2,1]
    ```
- `last` takes a list and returns its last element.
    ```haskell
    ghci> last [5,4,3,2,1]
    1
    ```
- `init` takes a list and return everything except its last element.
    ```haskell
    ghci> init [5,4,3,2,1]
    [5,4,3,2]
    ```
- `take` takes number and a list. It extracts that many elements from the beginning of the list.
    ```haskell
    take 3 [5,4,3,2,1]
    [5,4,3]
    ```
- `drop` drops the number of elements from the beginning of a list.
```haskell
    drop 2 [5,4,3,2,1]
    [3,2,1]
```
- Texas ranges
    ```haskell
    ghci> [1..5]
    [1,2,3,4,5]

    ghci> [1,3..10]
    [1,3,5,7,9]
    ```

- Lists within lists can be of different length but cannot be of different types.
- Lists are compared in lexicographical order. First the heads are compared. If they are equal then the second elements are compared, etc. Example:
    ```haskell
    ghci> [3,4,2] > [3,4]
    True
    ``` 