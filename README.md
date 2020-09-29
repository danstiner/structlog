# structlog [![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci]

Structured logger for Go using the [message templates](https://messagetemplates.org/) format. Inspired by [Serilog](https://serilog.net/).

## Motivation

Structured logging makes it easier to search and extract useful information from your logs. In Go this usually means passing key-value pairs to your logger in addition to a string message. This library takes a more compact approach by embeddeding key names in the message string itself.

## Example

```go
import (
    "github.com/danstiner/structlog"
    "github.com/danstiner/structlog/sink"
)

type Coordinates struct {
    Latitude, Longitude float64
}

log := structlog.New(sink.Json{os.Stdout})

log.Info("Processed {@position} in {elapsed} ms", Coordinates{25, 132}, 34)
```

Formatted output:

```json
{
  "$level":"INFO",
  "$message":"Processed {\"Latitude\":25,\"Longitude\":132} in 34 ms",
  "$template":"Processed {@position} in {elapsed} ms",
  "$timestamp":"2228-03-22T12:34:56Z",
  "elapsed":34,
  "position":{"Latitude":25,"Longitude":132}
}
```

Some benefits of this approach:

- Simple syntax
- Values are included in both the message and structured data without repeating yourself
- No awkward syntax for passing key-value pairs as structs or arguments that alternate between keys and values
- The template is logged to enable both aggregation and use of `grep` to find which line of code a log originated from
- `@` can be used to serialize structs and other complex values as JSON

## Interface

```go
type Logger struct {
    Trace(template string, values ...interface{})
    Info(template string, values ...interface{})
    Error(template string, values ...interface{})
    Panic(template string, values ...interface{})

    With(key string, value interface{}) Logger
    WithFields(fields Fields) Logger
    WithError(err error) Logger
}

type Fields map[string]interface{}
```

[doc-img]: https://godoc.org/github.com/danstiner/structlog?status.svg
[doc]: https://godoc.org/github.com/danstiner/structlog
[ci-img]: https://travis-ci.org/danstiner/structlog.svg?branch=master
[ci]: https://travis-ci.org/danstiner/structlog
