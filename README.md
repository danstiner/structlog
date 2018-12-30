# structlog [![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci]

Structured logger for Go using the [message templates](https://messagetemplates.org/) format.

## Motivation

Structured logging makes it easier to search and extract useful information from your logs. In Go this is usually accomplished by passing key-value pairs to your logger in addition to a string message. This library instead embeddeds the key names in a message format string, an approach originally from the [Serilog](https://serilog.net/) library for C#.

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

- Values are included in both the message and structured data without repeating yourself
- No awkward syntax for passing key-value pairs as structs or arguments that alternate between keys and values
- The template is logged to enable both aggregation and use of `grep` to find which line of code a log originated from
- `@` can be used to format a value as a JSON string using reflection

## Interface

```go
type Logger struct {
    Trace(template string, values ...interface{})
    Info(template string, values ...interface{})
    Error(template string, values ...interface{})
    Panic(template string, values ...interface{})

    With(key string, value interface{}) Logger
}
```

## Contributing

Contributions are welcome, please open an issue or send a pull-request.

### Testing

``make test``

### Benchmarks

``make bench``

[doc-img]: https://godoc.org/github.com/danstiner/structlog?status.svg
[doc]: https://godoc.org/github.com/danstiner/structlog
[ci-img]: https://travis-ci.org/danstiner/go-structlog.svg?branch=master
[ci]: https://travis-ci.org/danstiner/go-structlog