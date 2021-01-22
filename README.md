# Go Light Collections

Lightweight, efficient, and functionally rich implementation of List, Map, and Set for Golang, inspired by the Kotlin standard library.  The implementation uses Go's native slice and map as underlying data structures.

Due to the lack of generics in Go, the implementation uses the type interface{}.  I call the implementations *pseudo-generic*.  

An example is provided with a simple pattern in the [glc](https://github.com/pvillela/go-light-collections/tree/main/pkg/glc) package godoc documentation to show how to write adapters to create type-safe collections for specific types.  

This framework also supports code generation for specific types.  See the [genex](https://github.com/pvillela/go-light-collections/tree/main/examples/genex) example. 

The package [examples/gentest](https://github.com/pvillela/go-light-collections/tree/main/examples/gentests) and its sub-packages contain the generation of tests for
specific concrete types.  These tests serve to test the correctness of the code generation
process.


## To-dos

*   Add missing library godoc comments
*   Update collections package documentation
*   Write code generation documentation
*   Update code generation example documentation
