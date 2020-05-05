# rint
 > Bring forth a pseudo-random integer from the darkness.

## Install
 ```go
go get github.com/tlboright/go-rint
```

## Usage
```go
rint.Init()
// seeds the pseudorandom number generator

rint.Gen(10)
//=> 6

rint.GenRange(4, 42)
//=> 37
```
## API

### Gen(max int) int
Generates a random number from 0 to the maximum (non-inclusive).

### GenRange(min, max int) int
Generates a random number from the minimum to the maximum (non-inclusive).

### Init()
Seeds the random number generator with [cryptographically secure random byte data] (https://golang.org/pkg/crypto/rand/) gathered from the environment the library is running in. 
