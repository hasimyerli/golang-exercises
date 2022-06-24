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

var stringList StringList

func init() {
	stringList = StringList{"a", "b"}
}

func main() {
	fmt.Println(stringList.Contains("a"))  //true
	fmt.Println(stringList.Contains("b"))  //true
	fmt.Println(stringList.Contains("ab")) //false
}
```