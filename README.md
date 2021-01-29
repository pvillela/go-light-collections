# Go Light Collections

Lightweight, efficient, and functionally rich implementation of List, Map, and Set for Golang, inspired by the Kotlin standard library.  The implementation uses Go's native slice and map as underlying data structures.

Due to the lack of generics in Go, the implementation uses the type interface{}.  I call the implementations *pseudo-generic*.  

An example is provided with a simple pattern in the [glc](https://github.com/pvillela/go-light-collections/tree/main/pkg/glc) package godoc documentation to show how to write adapters to create type-safe collections for specific types.  

This framework also supports code generation for specific types.  See the [genex](https://github.com/pvillela/go-light-collections/tree/main/examples/genex) example. 

The package [examples/testgen](https://github.com/pvillela/go-light-collections/tree/main/examples/testgen) and its sub-packages contain the generation of tests for
specific concrete types.  These tests serve to test the correctness of the code generation
process.

This library is extensively tested, including 100% test coverage for the core code and the execution of the tests with code generation.

## g2lc branch

The g2lc package in the g2lc branch contains a draft implementation of this framework using Go2 generics.  The implementation compiles correctly with the go2go tool.  The go2go translation process needs to be applied to each package, with translation of a package taking place only after its dependencies have been translated.  In order to satisfy the go2go tool limitations, differnt collection types had to be moved to different packages.

## To-dos

*   Add missing library godoc comments
*   Update collections package documentation
*   Write code generation documentation
*   Update code generation example documentation
