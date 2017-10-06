package main

import "fmt"

func main() {
	alice := newPerson(aliceID, 4)
	bob := newPerson(bobID, 4)
	cathy := newPerson(cathyID, 4)
	dave := newPerson(daveID, 4)

	alice.suggest("Wednesday")
	dump()
	alice.discussWith(cathy)
	alice.discussWith(bob)

	bob.suggest("Tuesday")
	dump()
	bob.discussWith(dave)

	dave.suggest("Tuesday")
	dump()
	fmt.Printf("dave: %v", dave)

	cathy.suggest("Thursday")
	dump()
}

// import "fmt"

// func main() {
// 	fmt.Println("TODO: Write README, organise the code/tests")
// }
