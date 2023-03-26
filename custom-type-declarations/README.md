# Golang Custom Type Declarations

**Example:**
```go
package main

import "fmt"

type StringList []string


// Checks if a value exists in an list
func (s StringList) Contains(v string) bool {
	for _, i := range s {
		if i == v {
			return true
		}
	}

	return false
}

// Join list elements with a string
func (l StringList) Implode(separator string) string {
	s := ""
	for _, i := range l {
		s += i + separator
	}

	return strings.TrimSuffix(s, separator)
}

var fruits StringList

func init() {
	fruits = StringList{"apple", "grape", "cherry"}
}

func main() {
	//Example-1
	fmt.Println(fruits.Contains("apple")) //true
	fmt.Println(fruits.Contains("grape")) //true
	fmt.Println(fruits.Contains("pear"))  //false

	//Example-2
	fmt.Println(fruits.Implode(";")) // "apple;grape;pear"
}
```