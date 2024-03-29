package main

import (
    "fmt"
    "net/http"
    "strconv"
)

func main() {
    http.HandleFunc("/append", func(w http.ResponseWriter, r *http.Request) {
        value := r.URL.Query().Get("value")
        // Validate value here (e.g., allowed characters, operator logic)
        // ...

        // Update the expression in your server-side storage
        // (e.g., in-memory variable, database)
        expression := updateExpression(value)

        fmt.Fprintf(w, expression)
    })

    http.ListenAndServe(":8080", nil)
}

func updateExpression(value string) string {
    // Implement logic to update the expression based on value
    // ...

    return updatedExpression
}
