```go
func main() {
    var m map[string]int
    fmt.Println(m["key"]) // This will panic if the map is not initialized
}
```