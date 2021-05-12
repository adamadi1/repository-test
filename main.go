package main

import (
	"fmt"

	"github.com/mattn/go-sqlite3"
	"github.com/dghubble/go-twitter/twitter"
)

const followingProfileId int

type Follower struct {
	Name string
	Id string
}

func main() {
	db := initStorage()

	client := getTwitterClient()
	followers, err := getFollowersFromProfileWithId(followingProfileId)
	if err != nil {
		fmt.Fatalf("Can't get followers from Twitter: %v", err)
	}
	savedFollowers, err := getFollowersFromStorage(db, followingProfileId)
	if err != nil {
		fmt.Fatalf("Can't get followers from storage: %v", err)
	}

	unfollowed := unfollowed(followers, savedFollowers)
	printFollowers(unfollowed)

	err := saveFollowers(db, followers)
	if err != nil {
		fmt.Fatalf("Can't save new followers to storage: %v", err)
	}
}

func initStorage() {
	db := klientSqlLite() // @todo

	if isStorageFresh(db) {
		err := migrateStorage(db)
		if err != nil {
			fmt.Fatalf("Can't migrate db: %v", err)
		}
	}
	// @todo
}

func isStorageFresh(db) bool {
	// @todo sprawdzasz czy istnieje tabela Followers
}

func migrateStorage(db) error {
	// @todo wysylasz SQL, ktore ma stworzyc baze danych czy tabele Followers
}

func getTwitterClient() {
	// @todo
}

func getFollowersFromProfileWithId() ([]Follower, error) {
	res := make([]Follower, 0)
	// @todo
	// pobrac followersow z jakiegos profilu z Twittera
	// zamienic ich na strukture []Follower

	return res, nil
}

func getFollowersFromStorage(db, followingProfileId) ([]Follower, error) {
	// @todo
}

// unfollowed zwraca followersów, którzy są w fromStorage, ale nie ma ich w fromTwitter
func unfollowed(fromTwitter []Follower, fromStorage []Follower) ([]Follower) {
	// @todo
}

func printFollowers(follower []Follower) {
	// @todo
}