# structlog [![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci]

Structured logger for Go using the [message templates](https://messagetemplates.org/) format.

## Motivation

Structured logging is a great way to capture useful context when logging. [Serilog](https://serilog.net/) is a hugely popular C# library with a clean and concise message format, this library brings [that format]((https://messagetemplates.org/) to Go.

## Example

```go
import (
    "github.com/danstiner/structlog"
    "github.com/danstiner/structlog/sink"
)

type Position struct {
    Latitude, Longitude float64
}

log := structlog.New(sink.Json{os.Stdout})

log.Info("Processed {@position} in {elapsed} ms", Position{25, 132}, 34)
```

Output:

```json
{
  "$level":"INFO",
  "$message":"Processed {\"Latitude\":25,\"Longitude\":132} in 34 ms",
  "$template":"Processed {@position} in {elapsed} ms",
  "$timestamp":"2018-12-30T04:23:15Z",
  "elapsed":34,
  "position":{"Latitude":25,"Longitude":132}
}
```

Some benefits of this approach:

- No passing key-value structs or matching up pairs of key-value parameters
- Key names are derived from the message template
- The formatting is straightforward, a similar result could be achieved with `fmt.Sprintf`
- Optionally `@` can be used to format the value as a JSON string
- The template string can be used with `grep` to find the line of source code that logged

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

TODO Justify choice of methods https://winderresearch.com/logging-vs-tracing-vs-monitoring/

## Contributing

Contributions are welcomed, please open an issue or send a pull-request.

### Testing

``make test``

### Benchmarks

``make bench``

[doc-img]: https://godoc.org/github.com/danstiner/structlog?status.svg
[doc]: https://godoc.org/github.com/danstiner/structlog
[ci-img]: https://travis-ci.org/danstiner/go-structlog.svg?branch=master
[ci]: https://travis-ci.org/danstiner/go-structlog