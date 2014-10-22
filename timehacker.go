package main

import(
  "net/http"
   "github.com/codegangsta/martini"
   "github.com/codegangsta/martini-contrib/binding"
   "strings"
   _ "github.com/lib/pq"
  "database/sql"
  "log"
)

type Post struct {
    Data   string    `form:"data" binding:"required"`
}

func main() {
  m := martini.Classic()
  p := []Post{}

  db, err := sql.Open("postgres", "user=art dbname=timehackerdb sslmode=disable")
  if err != nil {
    log.Fatal(err)
  } 
  
  rows, err := db.Query("SELECT * FROM user_data")
  if err != nil {
    log.Fatal(err)
  }

  for rows.Next() {
    var time string
    var data string
     rows.Scan(&time, &data)
    // if err != nil {
    //   println(err)
    // }
    println(time)
  }

  m.Get("/feedbacks", func() string {
    s := "["
    for _, q := range p {
      s = s + q.Data + ","
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