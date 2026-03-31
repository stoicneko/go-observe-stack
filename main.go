package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	dbUser := os.Getenv("DB_USER") // 预期值: root
	dbPass := os.Getenv("DB_PASS") // 预期值: 123456
	dbHost := os.Getenv("DB_HOST") // 预期值: db
	dbPort := os.Getenv("DB_PORT") // 预期值: 3306
	dbName := os.Getenv("DB_NAME") // 预期值: myapp

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello from go backend")
	})

	http.HandleFunc("/ping-db", func(w http.ResponseWriter, r *http.Request) {
		if db.Ping() != nil {
			fmt.Fprintf(w, "db error")
			return
		}
		fmt.Fprintf(w, "db connected")
	})

	http.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// 1. 用 r.URL.Query().Get("content") 拿到 content
			content := r.URL.Query().Get("content")
			// 2. 如果 content 为空，返回错误信息，return
			if content == "" {
				fmt.Fprintf(w, "content is empty")
				return
			}
			// 3. 用 db.Exec("INSERT INTO messages (content) VALUES (?)", content) 插入
			_, err := db.Exec("INSERT INTO messages (content) VALUES (?)", content)
			// 4. 如果 err 不为 nil，返回错误信息，return
			if err != nil {
				fmt.Fprintf(w, "db error")
				return
			}
			// 5. 返回 "message saved"
			fmt.Fprintf(w, "message saved")
		} else {
			// 1. 用 db.Query("SELECT id, content, created_at FROM messages") 查询
			rows, err := db.Query("SELECT id, content, created_at FROM messages")
			// 2. 如果 err 不为 nil，返回错误信息，return
			if err != nil {
				fmt.Fprintf(w, "db error")
				return
			}
			// 3. 用 for rows.Next() 遍历
			//    每次循环里声明三个变量：var id int, var content string, var createdAt string
			//    用 rows.Scan(&id, &content, &createdAt) 读取
			//    用 fmt.Fprintf(w, "id=%d content=%s time=%s\n", id, content, createdAt) 输出
			for rows.Next() {
				var (
					id        int
					content   string
					createdAt string
				)
				rows.Scan(&id, &content, &createdAt)
				fmt.Fprintf(w, "id=%d content=%s time=%s\n", id, content, createdAt)
			}

		}
	})

	http.ListenAndServe(":5000", nil)
}
