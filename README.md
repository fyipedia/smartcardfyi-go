# smartcardfyi-go

[![Go Reference](https://pkg.go.dev/badge/github.com/fyipedia/smartcardfyi-go.svg)](https://pkg.go.dev/github.com/fyipedia/smartcardfyi-go)

Go client for the [SmartCardFYI](https://smartcardfyi.com) API. Look up smart cards, platforms, standards, manufacturers, applications, form factors, and certifications. Zero dependencies — stdlib only.

## Install

```bash
go get github.com/fyipedia/smartcardfyi-go
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"

    smartcardfyi "github.com/fyipedia/smartcardfyi-go"
)

func main() {
    client := smartcardfyi.NewClient()

    result, err := client.Search("javacard")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Found %d results for 'javacard'\n", result.Total)

    card, err := client.Card("javacard-3-1")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s: %s\n", card.Name, card.Description)
}
```

## API Methods

| Method | Description |
|--------|-------------|
| `Search(query)` | Search cards, standards, and glossary |
| `Card(slug)` | Get smart card details |
| `Platform(slug)` | Get platform details |
| `Standard(slug)` | Get standard details |
| `Manufacturer(slug)` | Get manufacturer details |
| `Application(slug)` | Get application details |
| `FormFactor(slug)` | Get form factor details |
| `Certification(slug)` | Get certification details |
| `GlossaryTerm(slug)` | Get glossary term definition |
| `Compare(slugA, slugB)` | Compare two smart cards |
| `Random()` | Get a random smart card |

## REST API

Free, no authentication required. Base URL: `https://smartcardfyi.com/api`

```bash
curl https://smartcardfyi.com/api/search/?q=javacard
curl https://smartcardfyi.com/api/card/javacard-3-1/
curl https://smartcardfyi.com/api/random/
```

Full OpenAPI spec: https://smartcardfyi.com/api/openapi.json

## Also Available

| Language | Package | Install |
|----------|---------|---------|
| Python | [smartcardfyi](https://pypi.org/project/smartcardfyi/) | `pip install smartcardfyi` |
| TypeScript | [smartcardfyi](https://www.npmjs.com/package/smartcardfyi) | `npm install smartcardfyi` |
| Go | [smartcardfyi-go](https://pkg.go.dev/github.com/fyipedia/smartcardfyi-go) | `go get github.com/fyipedia/smartcardfyi-go` |
| Rust | [smartcardfyi](https://crates.io/crates/smartcardfyi) | `cargo add smartcardfyi` |
| Ruby | [smartcardfyi](https://rubygems.org/gems/smartcardfyi) | `gem install smartcardfyi` |

## Code FYI Family

| Site | Domain | Focus |
|------|--------|-------|
| BarcodeFYI | [barcodefyi.com](https://barcodefyi.com) | Barcode symbologies and standards |
| QRCodeFYI | [qrcodefyi.com](https://qrcodefyi.com) | QR code types and encoding |
| NFCFYI | [nfcfyi.com](https://nfcfyi.com) | NFC chips and standards |
| BLEFYI | [blefyi.com](https://blefyi.com) | Bluetooth Low Energy profiles |
| RFIDFYI | [rfidfyi.com](https://rfidfyi.com) | RFID tags and frequencies |
| SmartCardFYI | [smartcardfyi.com](https://smartcardfyi.com) | Smart card platforms and standards |

## License

MIT
