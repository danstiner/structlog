# go-structlog
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fdanstiner%2Fgo-structlog.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fdanstiner%2Fgo-structlog?ref=badge_shield)


Structured message template logs for Golang

## Usage

```go
package example

import "google.golang.org/genproto/googleapis/type/latlng"

var log = structlog.New(sink.Json{os.Stdout})

func Example() {
    log.Info("Processed {@position} in {elapsed} ms", latlng.LatLng{25, 132}, 34)
}
```

```json
{

}
```

## Contributing

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

Running it then should be as simple as:

```console
$ make
```

### Testing

``make test``

### Benchmarking

``make bench``
