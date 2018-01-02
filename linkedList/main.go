package main

import sLL "github.com/aatishrana/DataStructures/linkedList/SinglyLinkList"

func main() {
	list := sLL.New()
	list.Append(&sLL.Link{"First", nil})
	list.Append(&sLL.Link{"Second", nil})
	list.Print()

	list.Prepend(&sLL.Link{"Third", nil})
	list.Append(&sLL.Link{"Forth", nil})
	list.Print()
}
