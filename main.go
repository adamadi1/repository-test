package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/oauth2/clientcredentials"
)

const followingProfileId int64 = 1392097083788771328
const followersTable = "followers"

type Follower struct {
	Name string
	Id   int64
}

type credentials struct {
	clientID     string
	clientSecret string
}

func main() {
	fmt.Println("Starting..")

	db := getDB()
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	initStorage(ctx, db)

	client := getTwitterClient(ctx, credentials{clientID: "blabla", clientSecret: "blabla"})
	followingNow, err := getFollowersFromProfileWithId(client, followingProfileId)
	if err != nil {
		panic(fmt.Sprintf("Can't get followers from Twitter: %v", err))
	}

	// followingNow := []Follower{
	// 	{
	// 		Id: 789,
	// 		Name: "Piotr",
	// 	},
	// }

	followingPrevious, err := getFollowersFromStorage(ctx, db, followingProfileId)
	if err != nil {
		panic(fmt.Sprintf("Can't get followers from storage: %v", err))
	}

	fmt.Println("")

	fmt.Println("Now followers:")
	printFollowers(followingNow)

	fmt.Println("Previous followers:")
	printFollowers(followingPrevious)

	difference := checkDifference(followingPrevious, followingNow)

	fmt.Println("Unfollowing:")
	printFollowers(difference)

	err = saveFollowers(ctx, db, followingProfileId, followingNow)
	if err != nil {
		panic(fmt.Sprintf("Can't save new followers to storage: %v", err))
	}
}

func saveFollowers(ctx context.Context, db *sql.DB, mainId int64, now []Follower) error {
	_, err := db.ExecContext(ctx, fmt.Sprintf("DELETE FROM `%s` WHERE main_id = ?", followersTable), mainId)
	if err != nil {
		return fmt.Errorf("can't remove old followers: %w", err)
	}

	for _, follower := range now {
		_, err := db.ExecContext(
			ctx,
			fmt.Sprintf("INSERT INTO `%s` VALUES (?, ?, ?)", followersTable),
			mainId, follower.Id, follower.Name,
		)

		if err != nil {
			return fmt.Errorf("can't add followers: %w (mainId: %d, follower: %+v)", err, mainId, follower)
		}
	}

	return nil
}

func initStorage(ctx context.Context, db *sql.DB) {
	err := migrateStorage(ctx, db)
	if err != nil {
		panic(err.Error())
	}
}

func getDB() *sql.DB {
	database, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err.Error())
	}

	return database
}

func migrateStorage(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS `%s` (main_id INTEGER, followers_id INTEGER, followers_name TEXT)",
		followersTable,
	), followersTable)

	return err
}

func getTwitterClient(ctx context.Context, creds credentials) *twitter.Client {
	config := &clientcredentials.Config{
		ClientID:     creds.clientID,
		ClientSecret: creds.clientSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}

	return twitter.NewClient(config.Client(ctx))
}

func getFollowersFromProfileWithId(client *twitter.Client, id int64) ([]Follower, error) {
	res := make([]Follower, 0)

	followers, _, err := client.Followers.List(&twitter.FollowerListParams{
		UserID: id,
		Cursor: 0,
		Count:  100,
	})

	if err != nil {
		return res, err
	}

	for _, twitterFollower := range followers.Users {
		res = append(res, Follower{Id: twitterFollower.ID, Name: twitterFollower.Name})
	}

	return res, nil
}

func getFollowersFromStorage(ctx context.Context, db *sql.DB, id int64) ([]Follower, error) {
	//
	// Followers
	// main_id | main_name | followers_id | followers_name
	// -------------------------------------------------------------------------------------------------
	// 123 | Krzysztof Gzocha | 456 | Followers1
	// 123 | Krzysztof Gzocha | 789 | Followers2

	rows, err := db.QueryContext(ctx, fmt.Sprintf("SELECT followers_id, followers_name FROM `%s` WHERE main_id = ?", followersTable), id)
	if err != nil {
		return nil, err
	}

	tabelaFollowersow := make([]Follower, 0)
	for rows.Next() {
		var followerId int64
		var followerName string

		err := rows.Scan(&followerId, &followerName)
		if err != nil {
			return tabelaFollowersow, err
		}

		tabelaFollowersow = append(tabelaFollowersow, Follower{
			Name: followerName,
			Id:   followerId,
		})
	}

	return tabelaFollowersow, nil
}

func checkDifference(now, previous []Follower) []Follower {
	wynik := make([]Follower, 0)

	for _, elementNow := range now {
		foundInPrevious := false

		for _, elementPrevious := range previous {
			if elementNow.Id == elementPrevious.Id {
				foundInPrevious = true
				break
			}
		}

		if !foundInPrevious {
			wynik = append(wynik, elementNow)
		}
	}

	return wynik
}

func printFollowers(followers []Follower) {
	for _, follower := range followers {
		fmt.Printf("- ID: %d, Name: %s\n", follower.Id, follower.Name)
	}
}
