package main

import (
	"database/sql"
	"fmt"
)

func FetchUsersWithPostsNPlusOne(allusers *[]UserResponse, db *sql.DB, totalQueries *int) {
	query := `SELECT Id,Email,Name from users`
	*totalQueries++
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var u UserResponse
		err = rows.Scan(&u.Id, &u.Email, &u.Name)
		if err != nil {
			panic(err)
		}
		var allposts []PostResponse
		GetAllPosts(db, &allposts, u.Id, totalQueries)
		u.Posts = allposts

		*allusers = append(*allusers, u)
	}

}

func GetAllPosts(db *sql.DB, allposts *[]PostResponse, id int64, totalQueries *int) {
	query := `SELECT title, body, id FROM posts WHERE userId = ?`
	*totalQueries++
	rows, err := db.Query(query, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var p PostResponse
		err = rows.Scan(&p.Title, &p.Body, &p.Id)
		if err != nil {
			panic(err)
		}

		*allposts = append(*allposts, p)
	}
}

func FetchUsersWithPostsOptimized(allusers *[]UserResponse, db *sql.DB, totalQueries *int,
	alljoinResponse *[]JoinResponse) {
	resultmap := make(map[int64]UserResponse)
	joinQuery := `
	Select U.id as userId, U.Email as email, P.id as PostId, P.title as title, P.body as body, U.Name as Name  from users as U left join posts as P on U.id = p.userId order by U.id
	`
	rows, err := db.Query(joinQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var j JoinResponse
		err = rows.Scan(&j.UserId, &j.Email, &j.PostId, &j.Title, &j.Body, &j.Name)
		fmt.Println(j.UserId)
		if err != nil {
			panic(err)
		}
		*alljoinResponse = append(*alljoinResponse, j)
		val, exists := resultmap[j.UserId]

		if !exists {
			u := UserResponse{Id: j.UserId, Email: j.Email, Name: j.Name}
			p := PostResponse{Id: j.PostId, Title: j.Title, Body: j.Body}
			u.Posts = append(u.Posts, p)
			resultmap[j.UserId] = u
		} else {
			p := PostResponse{Id: j.PostId, Title: j.Title, Body: j.Body}
			val.Posts = append(val.Posts, p)
			resultmap[j.UserId] = val
		}
	}

	for _, val := range resultmap {
		*allusers = append(*allusers, val)
	}

}
