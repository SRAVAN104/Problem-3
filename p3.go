// main.go
package main

import (
    "fmt"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    currentTime := time.Now()
    formattedTime := currentTime.Format("Monday, January 2, 2006 15:04:05 MST")
    utcTime := currentTime.UTC()
    formattedUTCTime := utcTime.Format("Monday, January 2, 2006 15:04:05 UTC")
    
    html := fmt.Sprintf(`
        <!DOCTYPE html>
        <html>
        <head>
            <title>Current Date & Time</title>
            <meta http-equiv="refresh" content="1">
        </head>
        <body>
            <h1>Current Date and Time</h1>
            <p>Local: %s</p>
            <p>UTC: %s</p>
        </body>
        </html>`, formattedTime, formattedUTCTime)
    
    w.Header().Set("Content-Type", "text/html")
    fmt.Fprintf(w, html)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server starting on port 8080...")
    http.ListenAndServe(":8080", nil)
}