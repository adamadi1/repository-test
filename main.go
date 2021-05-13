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
	db := getDB()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	initStorage(ctx, db)

	client := getTwitterClient(ctx, credentials{clientID: "blabla", clientSecret: "blabla"})
	followers, err := getFollowersFromProfileWithId(client, followingProfileId)
	if err != nil {
		panic(fmt.Sprintf("Can't get followers from Twitter: %v", err))
	}

	savedFollowers, err := getFollowersFromStorage(db, followingProfileId)
	if err != nil {
		panic(fmt.Sprintf("Can't get followers from storage: %v", err))
	}

	unfollowed := unfollowed(followers, savedFollowers)
	printFollowers(unfollowed)

	err := saveFollowers(db, followers)
	if err != nil {
		fmt.Fatalf("Can't save new followers to storage: %v", err)
	}
}

func initStorage(ctx context.Context, db *sql.DB) {
	if !isStorageFresh(ctx, db) {
		return
	}

	err := migrateStorage(ctx, db)
	if err != nil {
		panic(err.Error())
	}
}

func isStorageFresh(ctx context.Context, db *sql.DB) bool {
	rows, err := db.QueryContext(ctx, "SHOW TABLES")
	if err != nil {
		panic(err.Error())
	}

	var table string
	for rows.Next() {
		err := rows.Scan(&table)
		if err != nil {
			panic(err.Error())
		}

		if table == followersTable {
			return true
		}
	}

	return false
}

func getDB() *sql.DB {
	database, err := sql.Open("sqlite3", "./bogo.db")
	if err != nil {
		panic(err.Error())
	}

	return database
}

func migrateStorage(ctx context.Context, db *sql.DB) error {
	const createTableQuery = "CREATE TABLE IF NOT EXISTS `?` (id DECIMAL, name TEXT)"
	_, err := db.ExecContext(ctx, createTableQuery, followersTable)

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

func getFollowersFromStorage(db.Querty, id int) ([]Follower, error) {
	// @todo
}

// unfollowed zwraca followersów, którzy są w fromStorage, ale nie ma ich w fromTwitter
func unfollowed(fromTwitter []Follower, fromStorage []Follower) []Follower {
	// @todo
}

func printFollowers(follower []Follower) {
	// @todo
}
