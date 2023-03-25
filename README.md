
# Validator

Various validators for Golang 



## Things you can validate

- [Iranian National ID](https://github.com/rootiens/validator#national-id)
- [Iranian Phone Number](https://github.com/rootiens/validator#phone-number)


## Installation

Install my-project with npm

```bash
go get -u github.com/rootiens/validator
```
    

## Usage/Examples

### National ID

```go
package main

import (
     "github.com/rootiens/validator/id-card"
)

func main() {
  ok, err := idcard.IsValid("123456")

  if !ok {
    panic(err)
  }
}
```

### Phone Number

```go
package main

import (
     "fmt"
     "github.com/rootiens/validator/phone"
)

func main() {
  ok := phone.IsValid("123456")

  if !ok {
    fmt.Println("Invalid Phone")
  }
}
```




