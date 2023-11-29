# go-dev Dev Logging Support

The go-dev package provides a simple logging capability specifically intended to
allow developers to get a closer look at what is going on when debugging.

## Usage

Sample code:

```go
import dev "github.com/pjsoftware/go-dev"

dl := dev.InitLogging("./tests/log/tests.log")
dl.EnableLogging()

dl.Print("Logging output enabled")
dl.Printf("Printf also enabled: %d, %d, %d", 1,2,3)

dl.DisableLogging()
```

Standardised output:

```go
fn := dl.Enter("FuncName")
dl.ExitWithError(fn, 1, err)
dl.Exit(fn)
```
