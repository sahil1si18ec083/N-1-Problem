# N+1 Query Problem in Go (SQLite Example)

This repository demonstrates the **N+1 query problem** in backend systems and shows how to fix it using a **single optimized SQL query** with `LEFT JOIN` in Go (`database/sql` + SQLite).

---

## 📌 Problem Overview

We have two tables:

- `users`
- `posts` (each post belongs to a user)

Goal:  
👉 Fetch **all users along with their posts**

---

## ❌ What is the N+1 Query Problem?

The **N+1 query problem** occurs when:

1. One query is executed to fetch all parent records (users)
2. Then **N additional queries** are executed — one per parent — to fetch child records (posts)

### Query Pattern

```
1 query  → fetch users
N queries → fetch posts for each user
------------------------------
Total = 1 + N queries
```

This results in:
- Poor performance
- Excessive database round-trips
- Scalability issues as data grows

---

## ❌ N+1 Implementation (Inefficient)

### SQL Flow

```sql
SELECT id, email, name FROM users;       -- 1 query

FOR EACH user:
    SELECT title, body FROM posts WHERE userId = ?;  -- N queries
```

### Go Code (N+1)

```go
func FetchUsersWithPostsNPlusOne(allusers *[]UserResponse, db *sql.DB, totalQueries *int) {
	query := `SELECT Id, Email, Name FROM users`
	*totalQueries++

	rows, _ := db.Query(query)
	for rows.Next() {
		var u UserResponse
		rows.Scan(&u.Id, &u.Email, &u.Name)

		var posts []PostResponse
		GetAllPosts(db, &posts, u.Id, totalQueries)

		u.Posts = posts
		*allusers = append(*allusers, u)
	}
}
```

### Why this is bad ❌

- Database is queried repeatedly inside a loop
- Latency increases linearly with number of users
- Very inefficient in production systems

---

## ✅ Optimized Solution (Single Query)

### Key Idea

Use **one SQL query with `LEFT JOIN`** to fetch users and posts together, then aggregate the result in Go.

---

## ✅ Optimized SQL Query

```sql
SELECT
    U.id,
    U.email,
    P.id,
    P.title,
    P.body,
    U.name
FROM users U
LEFT JOIN posts P ON U.id = P.userId
ORDER BY U.id;
```

### Why `LEFT JOIN`?

- Ensures users **without posts** are still returned
- Post columns may be `NULL`

---

## ⚠️ Handling NULL values correctly

Because `LEFT JOIN` can return `NULL` values for post columns, Go must handle them explicitly.

### Solution

Use nullable SQL types:

- `sql.NullInt64`
- `sql.NullString`

These types contain:
- The actual value
- A `Valid` flag indicating whether the value is `NULL`

This prevents:
- Scan errors
- Fake posts with ID `0`
- Data corruption

---

## ✅ Optimized Go Implementation

```go
func FetchUsersWithPostsOptimized(allusers *[]UserResponse, db *sql.DB, totalQueries *int) {
	*totalQueries++

	resultmap := make(map[int64]UserResponse)

	query := `
	SELECT U.id, U.email, P.id, P.title, P.body, U.name
	FROM users U
	LEFT JOIN posts P ON U.id = P.userId
	ORDER BY U.id
	`

	rows, _ := db.Query(query)
	defer rows.Close()

	for rows.Next() {
		var j JoinResponse
		rows.Scan(&j.UserId, &j.Email, &j.PostId, &j.Title, &j.Body, &j.Name)

		user, exists := resultmap[j.UserId]
		if !exists {
			user = UserResponse{
				Id:    j.UserId,
				Email: j.Email,
				Name:  j.Name.String,
			}
		}

		if j.PostId.Valid {
			user.Posts = append(user.Posts, PostResponse{
				Id:    j.PostId.Int64,
				Title: j.Title.String,
				Body:  j.Body.String,
			})
		}

		resultmap[j.UserId] = user
	}

	for _, u := range resultmap {
		*allusers = append(*allusers, u)
	}
}
```

---

## 📊 Query Comparison

| Approach | Queries Executed |
|--------|------------------|
| N+1 | 1 + N |
| Optimized JOIN | **1** |

---

## ✅ Benefits of the Optimized Approach

- 🚀 Much faster execution
- 📉 Reduced database load
- 📦 Scales well with large datasets
- 💼 Production-ready pattern
- 🧠 Interview-friendly solution

---

## 🧠 Key Takeaways

- N+1 is a **design problem**, not a syntax error
- Queries inside loops are a red flag
- Use `LEFT JOIN` for optional relationships
- Always handle `NULL` values explicitly
- Aggregate joined data in memory

---

## 🏁 Conclusion

This project shows how a common backend performance issue can be fixed using:

- Better SQL design
- Proper handling of nullable fields
- Clean aggregation logic in Go

Understanding and fixing the N+1 query problem is a **core backend engineering skill**.
