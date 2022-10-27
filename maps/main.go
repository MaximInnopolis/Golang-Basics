package main

import "fmt"

type User struct {
	id   int
	name string
}

func main() {
	var defaultMap map[int]string

	fmt.Printf("Type: %T, value: %#v\n", defaultMap, defaultMap)
	fmt.Printf("Length: %#v\n", len(defaultMap))

	fmt.Println("~~~~~~~~~~~~~~~~~~")

	mapByMake := make(map[int]string)
	fmt.Printf("Type: %T, value: %#v\n", mapByMake, mapByMake)

	fmt.Println("~~~~~~~~~~~~~~~~~~")

	mapByMakeCap := make(map[int]string, 2)
	fmt.Printf("Type: %T, value: %#v\n", mapByMakeCap, mapByMakeCap)
	fmt.Printf("Length: %#v\n", len(mapByMakeCap))

	fmt.Println("~~~~~~~~~~~~~~~~~~")

	mapByLiteral := map[int]string{21: "Max", 20: "Anton"}
	fmt.Printf("Type: %T, value: %#v\n", mapByLiteral, mapByLiteral)
	fmt.Printf("Length: %#v\n", len(mapByLiteral))

	fmt.Println("~~~~~~~~~~~~~~~~~~")

	mapNew := *new(map[int]string)
	fmt.Printf("Type: %T, value: %#v\n", mapNew, mapNew)

	fmt.Println("~~~~~~~~~~~~~~~~~~")

	value, ok := mapByLiteral[20]
	fmt.Printf("Value: %#v, isExist: %#v\n", value, ok)

	fmt.Println("~~~~~~~~~~~~~~~~~~")

	delete(mapByLiteral, 20)
	fmt.Printf("Type: %T, value: %#v\n", mapByLiteral, mapByLiteral)

	fmt.Println("~~~~~~~~~~~~~~~~~~")

	mapForIteration := map[string]int{"First": 1, "Second": 2, "Third": 3}

	for key, value := range mapForIteration {
		fmt.Printf("Key: %#v, value: %#v\n", key, value)
	}

	fmt.Println("~~~~~~~~~~~~~~~~~~")

	users := []User{
		{45, "Maxim"},
		{36, "Alex"},
		{89, "Slava"},
		{45, "Kolya"},
	}

	uniqueUsers := make(map[int]struct{}, len(users))

	for _, user := range users {
		if _, ok := uniqueUsers[user.id]; !ok {
			uniqueUsers[user.id] = struct{}{}
		}
	}

	fmt.Printf("Type: %T, value: %#v\n", uniqueUsers, uniqueUsers)

	fmt.Println("~~~~~~~~~~~~~~~~~~")

	fmt.Println(findInSlice(36, users)) //Search User by his id O(n)

	mapUsers := make(map[int]User, len(users))

	for _, user := range users {
		if _, ok := mapUsers[user.id]; !ok {
			mapUsers[user.id] = user
		}
	}
	fmt.Println(mapUsers)

	fmt.Println(findInMap(45, mapUsers)) //Search User by his id O(1)
}

func findInSlice(id int, users []User) *User {
	for _, user := range users {
		if user.id == id {
			return &user
		}
	}

	return nil
}

func findInMap(id int, users map[int]User) *User {
	if user, ok := users[id]; ok {
		return &user
	}

	return nil
}
