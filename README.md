# smartcardfyi-go

[![Go Reference](https://pkg.go.dev/badge/github.com/fyipedia/smartcardfyi-go.svg)](https://pkg.go.dev/github.com/fyipedia/smartcardfyi-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Go client for the [SmartCardFYI](https://smartcardfyi.com) REST API. Smart cards, EMV, APDU, Java Card. Zero external dependencies beyond stdlib.

> **Try the interactive tools at [smartcardfyi.com](https://smartcardfyi.com)** — explore, search, and discover.

## Install

```bash
go get github.com/fyipedia/smartcardfyi-go
```

## Quick Start

```go
package main

import (
    "fmt"
    smartcardfyi "github.com/fyipedia/smartcardfyi-go"
)

func main() {
    client := smartcardfyi.NewClient()

    // Search across all content
    result, err := client.Search("query")
    if err != nil {
        panic(err)
    }
    fmt.Println(result)
}
```

## API Methods

| Method | Description |
|--------|-------------|
| `Applications()` | List applications |
| `CardTypes()` | List card types |
| `Categories()` | List categories |
| `Certifications()` | List certifications |
| `Faqs()` | List faqs |
| `FormFactors()` | List form factors |
| `Glossary()` | List glossary |
| `Guides()` | List guides |
| `Manufacturers()` | List manufacturers |
| `Personalization()` | List personalization |
| `Platforms()` | List platforms |
| `Standards()` | List standards |
| `Tools()` | List tools |
| `Search(query)` | Search across all content |

## Also Available

| Platform | Package | Link |
|----------|---------|------|
| **Python** | `pip install smartcardfyi` | [PyPI](https://pypi.org/project/smartcardfyi/) |
| **npm** | `npm install smartcardfyi` | [npm](https://www.npmjs.com/package/smartcardfyi) |
| **Go** | `go get github.com/fyipedia/smartcardfyi-go` | [pkg.go.dev](https://pkg.go.dev/github.com/fyipedia/smartcardfyi-go) |
| **Rust** | `cargo add smartcardfyi` | [crates.io](https://crates.io/crates/smartcardfyi) |
| **Ruby** | `gem install smartcardfyi` | [rubygems](https://rubygems.org/gems/smartcardfyi) |

## Tag FYI Family

Part of the [FYIPedia](https://fyipedia.com) open-source developer tools ecosystem.

| Site | Domain | Focus |
|------|--------|-------|
| BarcodeFYI | [barcodefyi.com](https://barcodefyi.com) | Barcode formats, EAN, UPC, ISBN standards |
| BLEFYI | [blefyi.com](https://blefyi.com) | Bluetooth Low Energy, GATT, beacons |
| NFCFYI | [nfcfyi.com](https://nfcfyi.com) | NFC chips, tag types, NDEF, standards |
| QRCodeFYI | [qrcodefyi.com](https://qrcodefyi.com) | QR code types, versions, encoding modes |
| RFIDFYI | [rfidfyi.com](https://rfidfyi.com) | RFID tags, frequency bands, standards |
| **SmartCardFYI** | [smartcardfyi.com](https://smartcardfyi.com) | Smart cards, EMV, APDU, Java Card |

## Embed Widget

Embed [SmartCardFYI](https://smartcardfyi.com) widgets on any website with [smartcardfyi-embed](https://widget.smartcardfyi.com):

```html
<script src="https://cdn.jsdelivr.net/npm/smartcardfyi-embed@1/dist/embed.min.js"></script>
<div data-smartcardfyi="entity" data-slug="example"></div>
```

Zero dependencies · Shadow DOM · 4 themes (light/dark/sepia/auto) · [Widget docs](https://widget.smartcardfyi.com)

## License

MIT
