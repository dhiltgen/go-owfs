# go-owfs
Golang library for interacting with an owfs 1-wire server

Still missing major functionality, but you can do a hello world directory listing with:

```go
package main

import (
        "github.com/dhiltgen/go-owfs"
        "log"
)

func main() {
        oc, err := owfs.NewClient("localhost:4304")
        if err != nil {
                log.Fatalf("Failed to connect: ", err)
        }
        defer oc.Close()
        sensors, err := oc.Dir("/")
        if err != nil {
                log.Fatalf("Failed to dir: ", err)
        }
        log.Println("Dir / ", sensors)
}
```

And you might see something like this:

```
2015/09/14 23:01:09 Dir /  [/10.F6D17C020800 /26.C29821010000 /81.C89F30000000 /1D.3ACA0F000000]
```
