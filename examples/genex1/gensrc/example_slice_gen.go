package gensrc

//go:generate ./gen_slice_0.sh Foo coll ../generated/coll
//go:generate ./gen_slice_01.sh Foo int coll ../generated/coll

//go:generate ./gen_slice_0.sh int coll ../generated/coll
//go:generate ./gen_slice_01.sh int string coll ../generated/coll
