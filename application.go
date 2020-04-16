package main

import (
    "bufio"
    "log"
    "net/http"
    "os"
    "math/rand"
    "fmt"
    "time"
)

func LinesInFile(fileName string) []string {
    f, _ := os.Open(fileName)
    
    scanner := bufio.NewScanner(f)
    result := []string{}
    
    for scanner.Scan() {
        line := scanner.Text()
    
        result = append(result, line)
    }
    return result
}

func main() {
    port := os.Getenv("PORT")
        if port == "" {
            port = "5000"
        }

        f, _ := os.Create("/var/log/golang/golang-server.log")
        defer f.Close()
        log.SetOutput(f)

        quotes := LinesInFile("quotes.txt")
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

            rand.Seed(time.Now().Unix())
            fmt.Fprintf(w, quotes[rand.Intn(len(quotes))])
            
        })

        log.Printf("Listening on port %s\n\n", port)
        http.ListenAndServe(":"+port, nil)
}
