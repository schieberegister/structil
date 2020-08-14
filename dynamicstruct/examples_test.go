package dynamicstruct

import (
	"fmt"

	"github.com/goldeneggg/structil"
)

func ExampleDynamicStruct_DecodeMap() {
	type Hoge struct {
		Key   string
		Value interface{}
	}

	hogePtr := &Hoge{
		Key:   "keystr",
		Value: "valuestr",
	}

	// Add struct fields using Builder
	b := NewBuilder().
		AddString("StringField").
		AddInt("IntField").
		AddFloat32("Float32Field").
		AddBool("BoolField").
		AddMap("MapField", SampleString, SampleFloat32).
		AddStructPtr("StructPtrField", hogePtr).
		AddSlice("SliceField", SampleInt)

	// Remove one field
	b = b.Remove("Float32Field")

	// Build returns a DynamicStruct
	ds := b.Build()

	// Decode to struct from map
	input := map[string]interface{}{
		"StringField":    "Abc Def",
		"IntField":       int(123),
		"BoolField":      true,
		"MapField":       map[string]float32{"mkey1": float32(1.23), "mkey2": float32(4.56)},
		"StructPtrField": hogePtr,
		"SliceField":     []int{111, 222},
	}
	dec, err := ds.DecodeMap(input)
	if err != nil {
		panic(err)
	}

	// Confirm decoded result using Getter
	g, err := structil.NewGetter(dec)
	if err != nil {
		panic(err)
	}
	fmt.Printf(
		"NumField: %d, String: %s, Int: %d, Bool: %v, Map: %+v, StructPtrField: %+v, SliceField: %+v\n",
		ds.NumField(),
		g.String("StringField"),
		g.Int("IntField"),
		g.Bool("BoolField"),
		g.Get("MapField"),
		g.Get("StructPtrField"),
		g.Get("SliceField"),
	)
	// Output:
	// NumField: 6, String: Abc Def, Int: 123, Bool: true, Map: map[mkey1:1.23 mkey2:4.56], StructPtrField: {Key:keystr Value:valuestr}, SliceField: [111 222]
}

func ExampleDynamicStruct_Definition() {
	type Hoge struct {
		Name   string
		Object interface{}
	}

	hogePtr := &Hoge{
		Name:   "this is name",
		Object: "this is object",
	}

	// Add struct fields with tag using Builder
	b := NewBuilder().
		AddString("MyStringNoTag").
		AddStringWithTag("MyString", `json:"my_string"`).
		AddIntWithTag("MyInt", `json:"my_int"`).
		AddFloat64WithTag("MyFloat64", `json:"my_float64"`).
		AddMapWithTag("MyMap", SampleString, SampleFloat32, `json:"my_map"`).
		AddStructPtrWithTag("MyStructPtr", hogePtr, `json:"my_struct_ptr"`).
		AddSliceWithTag("MySlice", SampleInt, `json:"my_slice"`)

	// Set struct name to DynamicStruct
	b.SetStructName("MyTestStruct")

	// Build returns a DynamicStruct
	ds := b.Build()

	// Print struct definition of built DynamicStruct
	// Fields are sorted by field name
	fmt.Println(ds.Definition())
	// Output:
	//type MyTestStruct struct {
	// 	MyFloat64 float64 `json:"my_float64"`
	// 	MyInt int `json:"my_int"`
	// 	MyMap map[string]float32 `json:"my_map"`
	// 	MySlice []int `json:"my_slice"`
	// 	MyString string `json:"my_string"`
	// 	MyStringNoTag string
	// 	MyStructPtr *struct { Name string; Object interface {} } `json:"my_struct_ptr"`
	//}
}
