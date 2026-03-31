package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// initDB 将环境变量读取和数据库连接逻辑封装
func initDB() {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
}

// Handler: 根路径
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello from go backend")
}

// Handler: 数据库健康检查
func pingDBHandler(w http.ResponseWriter, r *http.Request) {
	if db == nil || db.Ping() != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "db error")
		return
	}
	fmt.Fprintf(w, "db connected")
}

// Handler: 消息处理 (POST 插入, GET 查询)
func messagesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		content := r.URL.Query().Get("content")
		if content == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "content is empty")
			return
		}
		_, err := db.Exec("INSERT INTO messages (content) VALUES (?)", content)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "db error")
			return
		}
		fmt.Fprintf(w, "message saved")
	} else {
		rows, err := db.Query("SELECT id, content, created_at FROM messages")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "db error")
			return
		}
		defer rows.Close()

		for rows.Next() {
			var id int
			var content string
			var createdAt string
			rows.Scan(&id, &content, &createdAt)
			fmt.Fprintf(w, "id=%d content=%s time=%s\n", id, content, createdAt)
		}
	}
}

func main() {
	initDB()

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/ping-db", pingDBHandler)
	http.HandleFunc("/messages", messagesHandler)

	fmt.Println("Server starting at :5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
