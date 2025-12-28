package main

import (
	"database/sql"
	"fmt"
	"sort"

	_ "github.com/mattn/go-sqlite3"
)

type PostResponse struct {
	Id    int64  `json:"userId"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
type UserResponse struct {
	Id    int64          `json:"userId"`
	Email string         `json:"email"`
	Name  string         `json:"name"`
	Posts []PostResponse `json:"posts"`
}
type JoinResponse struct {
	UserId int64  `json:"userId"`
	Email  string `json:"email"`
	PostId int64  `json:"postId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Name   string `json:"name"`
}

func main() {
	totalQueries := 0
	var allusers []UserResponse
	var alljoinResponse []JoinResponse

	db, err := sql.Open("sqlite3", "app.db")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	// SeedData(db)

	// FetchUsersWithPostsNPlusOne(&allusers, db, &totalQueries)
	FetchUsersWithPostsOptimized(&allusers, db, &totalQueries, &alljoinResponse)
	// for _, val := range allusers {
	// 	fmt.Println(val)
	// }
	sort.Slice(allusers, func(i, j int) bool {
		return allusers[i].Id < allusers[j].Id
	})
	fmt.Println(allusers)
	fmt.Printf("query size is %d", totalQueries)

}
