# logfmt-event

This is a simple go package that will format data into a logfmt standard. Uses RFC 3339 standard for time.

## Usage

```go
import "github.com/egrzeszczak/logfmtevt"


func main() {
    evt := logfmtevt.New([]logfmtevt.Pair{
        {Key: "type", Value: "Security"},
        {Key: "category", Value: "register"},
        {Key: "user", Value: "egrzeszczak"},
        {Key: "email", Value: "ernest.grzeszczak@protonmail.com"},
        {Key: "name", Value: "Ernest Grzeszczak"},
        {Key: "srcip", Value: "10.30.30.15"},
        {Key: "result", Value: "success"},
    })
    
    fmt.Println(evt.String())
}
```

