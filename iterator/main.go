package main

import "fmt"
import "iterator/collection"

func main() {
	user1 := &collection.User{
		Name: "a",
		Age:  30,
	}
	user2 := &collection.User{
		Name: "b",
		Age:  20,
	}
	userCollection := &collection.UserCollection{
		Users: []*collection.User{user1, user2},
	}
	iterator := userCollection.CreateIterator()
	for iterator.HasNext() {
		user := iterator.GetNext()
		fmt.Printf("User is %+v\n", user)
	}
}
