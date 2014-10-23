package main

import(
  "net/http"
   "github.com/codegangsta/martini"
   "github.com/codegangsta/martini-contrib/binding"
   "strings"
   _ "github.com/lib/pq"
  "database/sql"
  "log"
  "time"
  "github.com/bitly/go-simplejson"
)

type Post struct {
    Data   string    `form:"data" binding:"required"`
}

type OnePost struct {
  Timestmp time.Time
  Data string
}

func printErr(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func main() {
  m := martini.Classic()
  p := []Post{}

  db, err := sql.Open("postgres", "user=art dbname=timehackerdb sslmode=disable")
  printErr(err)

  m.Get("/feedbacks", func() string {
    s := "["

rows, err := db.Query("SELECT * FROM user_data")
printErr(err)

for rows.Next() {
    var mytime time.Time
    var data string
    err := rows.Scan(&mytime, &data)
    printErr(err)

    _ := [mytime:Timestmp, data:Data]

println(mytime.Format("Mon Jan 2 15:04:05 -0700 MST 2006"))
    s = s + data + ","
  }

    s = strings.TrimRight(s, ",")
    s += "]"
    return s
  })

  m.Post("/feedback", binding.Bind(Post{}), func(post Post) string {
        // This function won't execute if there were errors
		p = append(p, post)
    return "OK"
    })

  http.ListenAndServe("127.0.0.1:8001", m)
  m.Run()

}