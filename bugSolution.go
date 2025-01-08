```go
package main

import (


    "fmt"
)

func main() {
    // Initialize the map
    m := make(map[string]int)
    
    //Option 1: Check for key existence
    value, ok := m["key"]
    if ok {
        fmt.Println("Value:", value)
    } else {
        fmt.Println("Key not found")
    }
    
    // Option 2: Provide a default value using the comma ok idiom
    value, ok = m["anotherkey"]
    if !ok{
        value = 0 // or some other default value
    }
    fmt.Println("Value:", value)
    
    m["key"] = 10
    fmt.Println(m["key"])
}
```