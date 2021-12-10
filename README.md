# Go Light Collections

Lightweight, efficient, and functionally rich implementation of List, Map, and Set for Golang, inspired by the Kotlin standard library.  The implementation uses Go's native slice and map as underlying data structures.



## glc package

The g2lc package contains an implementation of this framework using Go generics.

## Legacy

Due to the former lack of generics in Go, an implementation without generics was originally created.  That implementation, located in the **`legacy`** directory, uses the type `interface{}`.  I call this implementation *pseudo-generic*.  

An example is provided with a simple pattern in the [legacy/pkg/glc](https://github.com/pvillela/go-light-collections/tree/main/legacy/pkg/glc) package godoc documentation to show how to write adapters to create type-safe collections for specific types.  

The legacy framework also supports code generation for specific types.  See the [genex](https://github.com/pvillela/go-light-collections/tree/main/legacy/examples/genex) example. 

The package [legacy/examples/testgen](https://github.com/pvillela/go-light-collections/tree/main/legacy/examples/testgen) and its sub-packages contain the generation of tests for
specific concrete types.  These tests serve to test the correctness of the code generation
process.

The legacy library is extensively tested, including 100% test coverage for the core code and the execution of the tests with code generation.

## To-dos

Add missing library godoc comments

Update glc package documentation

