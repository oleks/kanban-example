package main

import (
  "database/sql"
  "fmt"
  "log"
  "net/http"
  "os"
  "time"
  _ "github.com/go-sql-driver/mysql"
  "github.com/getsentry/sentry-go"
)

type kanban struct {
  db *sql.DB
}

func (k *kanban) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  switch r.URL.EscapedPath() {
    case "/v1/board/create":
      fmt.Fprintf(w, "Creating a board!\n")
      // Options: name
      // Result: board_id | error
      // Let board_id 0 be reserved for done tasks
    case "/v1/task/create":
      fmt.Fprintf(w, "Creating a task!\n")
      // Options: board_id, title
      // Result: task_id | error
    case "/v1/task/move":
      fmt.Fprintf(w, "Moving a task!\n")
      // Options: task_id, board_id
      // Result: ok | error
    default:
      fmt.Fprintf(w, "404\n")
  }
}

func main() {
  kanban := new(kanban)

  err := sentry.Init(sentry.ClientOptions{
    Dsn: os.Getenv("SENTRY_DSN"),
  })
  if err != nil {
    log.Fatalf("sentry.Init: %s", err)
  }
  defer sentry.Flush(2 * time.Second)

  db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/dolt")
  if err != nil {
    panic(err.Error())
  }
  kanban.db = db
  defer kanban.db.Close()

  var version string

  err2 := db.QueryRow("SELECT VERSION()").Scan(&version)

  if err2 != nil {
    log.Fatal(err2)
  }

  fmt.Println(version)

  fmt.Printf("Starting server at :8080\n")

  http.Handle("/", kanban)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
