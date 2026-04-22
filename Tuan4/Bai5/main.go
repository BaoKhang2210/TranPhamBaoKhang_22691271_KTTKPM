package main
import (
    "fmt"
    "net/http"
)
func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Chào cậu! Đây là ứng dụng Go chạy trong Docker.")
    })
    fmt.Println("Server đang chạy tại cổng 8080...")
    http.ListenAndServe(":8080", nil)
}