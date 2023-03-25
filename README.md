
# Validator

Various validators for Golang 



## Things you can validate

- [Iranian National ID](https://github.com/rootiens/validator/new/master?readme=1#national-id)


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

