package string_utils

func ExamplePrintRot13Encoded() {
	PrintRot13Encoded("")
	PrintRot13Encoded("\n")
	PrintRot13Encoded("hello\n")
	PrintRot13Encoded("HELLO\n")
	PrintRot13Encoded("Hello!\n")
	// Output:
	// uryyb
	// URYYB
	// Uryyb!
}
