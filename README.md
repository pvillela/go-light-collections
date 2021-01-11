# Go Light Collections

Lightweight, reasonably efficient, and functionally rich implementation of List, Map, and Set for Golang, inspired by the Kotlin standard library.  The implementation uses Go's native slice and map as underlying data structures.

Due to the lack of generics in Go, the implementation uses the type interface{}.  I call the implementations *pseudo-generic*.  

Examples are provided to show how to write adapters to create type-safe collections for specific types using a simple pattern.  

In addition, type aliases/wrappers are used to facilitate code generation with replacement of the pseudo-generic types with specific types.

See the package Godocs for additional details.

## To-dos

*   Update collections package documentation
    -   Mention code generation option
    -   Describe type name convention to support code generation
*   Add missing godoc comments
*   Add Map methods
*   Add Set methods
*   Update code generation example
