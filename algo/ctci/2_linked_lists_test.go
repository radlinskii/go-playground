package ctci

func ExampleAppendHead() {
	l := makeList()
	l.appendHead(1)
	l.appendHead(2)
	l.appendHead(3)
	l.appendHead(4)
	l.print()

	// Output:
	// ->4->3->2->1->
}

func ExampleAppendTail() {
	l := makeList()
	l.appendTail(4)
	l.appendTail(3)
	l.appendTail(2)
	l.appendTail(1)
	l.print()

	// Output:
	// ->4->3->2->1->
}

func ExampleRemoveDuplicates1() {
	l := makeList()
	l.appendTail(4)
	l.appendTail(4)
	l.appendTail(3)
	l.appendTail(2)
	l.appendTail(2)
	l.appendTail(1)
	l.appendTail(1)
	l.removeDuplicates1()
	l.print()

	// Output:
	// ->4->3->2->1->
}
