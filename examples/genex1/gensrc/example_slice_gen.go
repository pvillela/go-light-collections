package gensrc

//go:generate ./gen_slice_0.sh Foo coll ../generated/coll
//go:generate ./gen_slice_1.sh Foo int coll ../generated/coll

//go:generate ./gen_slice_0.sh int coll ../generated/coll
//go:generate ./gen_slice_1.sh int string coll ../generated/coll
