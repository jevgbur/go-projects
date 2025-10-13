package main 

func main() {
	list := SinglyLinkedList{} 

	list.InsertLast("Hello")
	list.InsertLast("My")
	list.InsertLast("World")

	list.Print()
	
	list.DeleteLast()
	list.Print() 
	
	list.DeleteLast()
        list.Print()

	list.DeleteLast()
        list.Print()

	list.InsertLast("pear")
	list.Print()

	list.InsertFirst("apple")
	list.Print()

	list.InsertFirst("avocado")
	list.Print() 

	list.DeleteFirst()
	list.Print()


	nodeToDelete := list.head.next // We know that "pear" is the second node.
	list.DeleteAt(nodeToDelete)

	list.Print()
}
