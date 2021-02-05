package benching

import (
	"math/rand"
	"testing"
	"time"
)

// Create slices and Maps to use in benchmark so we dont have to recreate them each itteration
var userSlice10000 []User
var userSlice100000 []User
var userSlice1000000 []User

var userMap10000 map[int]User
var userMap100000 map[int]User
var userMap1000000 map[int]User

// Names Some Non-Random name lists used to generate Random Users
var Names []string = []string{"Jasper", "Johan", "Edward", "Niel", "Percy", "Adam", "Grape", "Sam", "Redis", "Jennifer", "Jessica", "Angelica", "Amber", "Watch"}

// SirNames Some Non-Random name lists used to generate Random Users
var SirNames []string = []string{"Ericsson", "Redisson", "Edisson", "Tesla", "Bolmer", "Andersson", "Sword", "Fish", "Coder"}

var IDCounter int

type User struct {
	ID   int
	Name string
}

func init() {
	userMap10000 = generateMapUsers(10000)
	userMap100000 = generateMapUsers(100000)
	userMap1000000 = generateMapUsers(1000000)

	userSlice10000 = generateSliceUsers(10000)
	userSlice100000 = generateSliceUsers(100000)
	userSlice1000000 = generateSliceUsers(1000000)

	// Shuffle the slice
	rand.Shuffle(len(userSlice10000), func(i, j int) {
		userSlice10000[i], userSlice10000[j] = userSlice10000[j], userSlice10000[i]
	})
	rand.Shuffle(len(userSlice100000), func(i, j int) {
		userSlice100000[i], userSlice100000[j] = userSlice100000[j], userSlice100000[i]
	})
	rand.Shuffle(len(userSlice1000000), func(i, j int) {
		userSlice1000000[i], userSlice1000000[j] = userSlice1000000[j], userSlice1000000[i]
	})
}

// generateSliceUsers will generate X amount of users in a slice
func generateSliceUsers(x int) []User {
	IDCounter = 0
	users := make([]User, x)
	for i := 0; i <= x; i++ {
		users = append(users, generateRandomUser())
	}
	return users
}

// generateSliceUsers will generate X amount of users in a slice
func generateMapUsers(x int) map[int]User {
	IDCounter = 0
	users := make(map[int]User, x)
	for i := 0; i <= x; i++ {
		users[i] = generateRandomUser()
	}
	return users
}

// generate a RandomUser
func generateRandomUser() User {
	rand.Seed(time.Now().UnixNano())
	nameMax := len(Names)
	sirNameMax := len(SirNames)

	nameIndex := rand.Intn(nameMax-1) + 1
	sirNameIndex := rand.Intn(sirNameMax-1) + 1
	id := IDCounter
	IDCounter++
	return User{
		ID:   id,
		Name: Names[nameIndex] + " " + SirNames[sirNameIndex],
	}
}

// findUserInSlice is used to search a slice of users for a userID
func findUserInSlice(userID int, users []User) {
	// Now find the userID provided
	for _, user := range users {
		if user.ID == userID {
			return
		}
	}

}

// findUserInMap is used to search a Map of users for a particular ID
func findUserInMap(userID int, users map[int]User) {
	if _, ok := users[userID]; ok {
		return
	}
}

// BenchmarkFindUserInSlice1000000 will search in a slice of size 1000000 to find a user ID
func BenchmarkFindUserInSlice1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findUserInSlice(7777, userSlice1000000)
	}
}

// BenchmarkFindUserInSlice100000 will search in a slice of size 100000 to find a user ID
func BenchmarkFindUserInSlice100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findUserInSlice(7777, userSlice100000)
	}
}

// BenchmarkFindUserInSlice10000 will search in a slice of size 10000 to find a user ID
func BenchmarkFindUserInSlice10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findUserInSlice(7777, userSlice10000)
	}
}

// BenchmarkFindUserInMap1000000 will search in a Map of size 1000000 to find a user ID
func BenchmarkFindUserInMap1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findUserInMap(7777, userMap1000000)
	}
}

// BenchmarkFindUserInMap100000 will search in a Map of size 100000 to find a user ID
func BenchmarkFindUserInMap100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findUserInMap(7777, userMap100000)
	}
}

// BenchmarkFindUserInMap10000 will search in a Map of size 10000 to find a user ID
func BenchmarkFindUserInMap10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findUserInMap(7777, userMap100000)
	}
}
