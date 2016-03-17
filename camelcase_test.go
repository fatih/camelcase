package camelcase

import "fmt"

func ExampleSplit() {

	for _, c := range []string{
		"",
		"lowercase",
		"Class",
		"MyClass",
		"MyC",
		"HTML",
		"PDFLoader",
		"AString",
		"SimpleXMLParser",
		"vimRPCPlugin",
		"GL11Version",
		"99Bottles",
		"May5",
		"BFG9000",
		"BöseÜberraschung",
		"Two  spaces",
		"BadUTF8\xe2\xe2\xa1",
	} {
		fmt.Printf("%#v\n", Split(c))
	}

	// Output:
	// []string{}
	// []string{"lowercase"}
	// []string{"Class"}
	// []string{"My", "Class"}
	// []string{"My", "C"}
	// []string{"HTML"}
	// []string{"PDF", "Loader"}
	// []string{"A", "String"}
	// []string{"Simple", "XML", "Parser"}
	// []string{"vim", "RPC", "Plugin"}
	// []string{"GL", "11", "Version"}
	// []string{"99", "Bottles"}
	// []string{"May", "5"}
	// []string{"BFG", "9000"}
	// []string{"Böse", "Überraschung"}
	// []string{"Two", "  ", "spaces"}
	// []string{"BadUTF8\xe2\xe2\xa1"}
}
