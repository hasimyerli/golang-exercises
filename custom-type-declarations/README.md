# Golang Custom Type Declarations

**Example:**
```go
package main

import "fmt"

type StringList []string

func (s StringList) Contains(v string) bool {
	for _, i := range s {
		if i == v {
			return true
		}
	}

	return false
}

var fruits StringList

func init() {
	fruits = StringList{"apple", "grape", "cherry"}
}

func main() {
	fmt.Println(fruits.Contains("apple")) //true
	fmt.Println(fruits.Contains("grape")) //true
	fmt.Println(fruits.Contains("pear"))  //false
}
```