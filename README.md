# GoSimpleCollections

Efficient and functionally rich implementation of List, Map, and Set for Golang, inspired on the Kotlin standard library.  The implementation uses Go's native slice and map as underlying data structures.

Due to the lack of generics in Go, the implementation uses the type interface{}.  I call the implementations *pseudo-generic*.  

Examples are provided to show how to write adapters to create type-safe collections for specific types using a simple pattern.  

In addition, type aliases/wrappers are used to facilitate code generation with replacement of the pseudo-generic types with specific types.

See the Godoc for additional details.
