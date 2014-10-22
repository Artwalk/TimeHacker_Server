package main

import(
  "net/http"
   "github.com/codegangsta/martini"
   "github.com/codegangsta/martini-contrib/binding"
   "strings"
)

type Post struct {
    Data   string    `form:"data" binding:"required"`
}

func main() {
  m := martini.Classic()
  p := []Post{}

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