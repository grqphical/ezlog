# ezlog - A Simple logging library for GO

## Installation
```bash
$   go get github.com/grqphical07/ezlog
```

## Usage

```go
import "github.com/grqphical07/ezlog"
impor "time"

func main() {
    ezlog.LogInfo("Information")
    ezlog.LogError("Error")

    logger := ezlog.NewLogger(1, "log.txt", "ApplicationName", time.RFC822, false, true)
}
```
