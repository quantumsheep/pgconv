# Pgconv

Convert `pgx/v5`'s `pgtype` types to and from native Go types.

## How to install

```bash
go get -u github.com/quantumsheep/pgconv
```

## Usage

```go
import (
    "github.com/quantumsheep/pgconv"
)

text := pgconv.ToText("Hello, World!") // pgtype.Text
str := pgconv.FromText(text)           // string
```
