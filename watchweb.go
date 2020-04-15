// Log the response code of a website
// Author = ScaredOS
// 4/15/2020
// ToDo:
//  Add more options:
//    Keyword check (Check for keyword, if not there, return error to log)\
//    Set Interval of check (Time to sleep for)
//    Set where to write the log to
package main
import (
  "fmt"
  "log"
  "net/http"
  "time"
  "os"
  "strconv"
)


func main() {
  url := os.Args[1]
  f, errr := os.Create("watchweb.log")
  if errr != nil {log.Fatal(errr)}
    for {
      resp, err := http.Get(url)
      if err != nil {log.Fatal(err)}
      stat := strconv.Itoa(resp.StatusCode)
      _, err = f.WriteString(fmt.Sprintf("[%s] Response code on (%s) - %s\n", stat, url, time.Now().Format("2006-01-02 15:04:05")))
      fmt.Printf("[%s] Response code on (%s) - %s\n", stat, url, time.Now().Format("2006-01-02 15:04:05"))
      if err != nil {log.Fatal(err)}
      time.Sleep(300 * time.Second)
    }
}
