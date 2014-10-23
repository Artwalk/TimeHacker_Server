package main

import (
	"database/sql"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	_ "github.com/lib/pq"
)

type Post struct {
	Data string `form:"data" binding:"required"`
}

func printErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	m := martini.Classic()

	db, err := sql.Open("postgres", "user=art dbname=timehackerdb sslmode=disable")
	printErr(err)
	defer db.Close()

	m.Get("/feedbacks", func() string {
		s := "{"

		rows, err := db.Query("SELECT * FROM user_data")
		printErr(err)

		for rows.Next() {
			var mytime time.Time
			var data string
			err := rows.Scan(&mytime, &data)
			printErr(err)

			s = s + "\"" + mytime.Format("2006-01-02 15:04:05") + "\":" + data + ","
		}

		s = strings.TrimRight(s, ",")
		return s + "}"
	})

	m.Post("/feedback", binding.Bind(Post{}), func(post Post) string {
		_, err := db.Exec("INSERT INTO user_data(time, data) VALUES(now(), $1)", post.Data)
		printErr(err)

		return "ok"
	})

	http.ListenAndServe("127.0.0.1:8001", m)
	m.Run()
}
