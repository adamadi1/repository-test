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
	followers, err := getFollowersFromProfileWithId(1392097083788771328)
	if err != nil {
		fmt.Fatalf("Can't get followers from Twitter: %v", err)
	}
	savedFollowers, err := getFollowersFromStorage(db, 1392097083788771328)
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
	db := klientSqlLite() 
	var (
		queryGetUsers db.Query = "SELECT user_id, first_name, latest_name, active, modified_date FROM users"
	)
	// @todo

	if isStorageFresh(db.Querty) {
		err := migrateStorage(db.Querty)
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
	database, _ :=
	sql.Open("sqlite3", "./bogo.db")
	statement, _ =
	database.Prepare("CREATE TABLE people user_id(1392097083788771328)")
	tabele.Followers()
	// @todo wysylasz SQL, ktore ma stworzyc baze danych czy tabele Followers
}

func getTwitterClient(creds *Credentials) (twitter.Client, error) {
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	// @todo
}

func getFollowersFromProfileWithId(1392097083788771328) ([]Follower, error) {

	res := make([]Follower, 0)
	// @todo
	// pobrac followersow z jakiegos profilu z Twittera
	// zamienic ich na strukture []Follower

	return res, nil
}

func getFollowersFromStorage(db.Querty, 1392097083788771328) ([]Follower, error) {
	// @todo
}

// unfollowed zwraca followersów, którzy są w fromStorage, ale nie ma ich w fromTwitter
func unfollowed(fromTwitter []Follower, fromStorage []Follower) ([]Follower) {
	// @todo
}

func printFollowers(follower []Follower) {
	// @todo
}