# xcassets
[`xcassets`](https://pkg.go.dev/github.com/illyabusigin/apptools/xcassets?tab=doc "API documentation") package
-------------------------------------------------------------------------------------------

The `xcassets` package provides methods for declaring and generating Xcode assets. You can choose to generate one-off assets or your entire .xcassets folder.

Features include:
- Generate app icon, colors, launch images
- Support for remote images
- Functional approach 
- Strongly typed
- Built-in validation with human-readale errors
- Output to a string or file

See it in action:

## Application Icon

```go
package main

import (
	"fmt"
	"log"

	"github.com/illyabusigin/apptools/xcassets"
)

func main() {
	builder := AppIcon("AppIcon", func(b *AppIconBuilder) {
		b.File("./testdata/Icon.png")
		b.Phone().Configure(func(b *AppIconPhone) {
			b.Settings.File("./testdata/Icon.png") // override individual icon configurations
		})
		b.AppStore()
	})

    if err := builder.Validate(); err != nil {
        log.Fatal("Failed validation", err)
    }

	err := builder.SaveTo("./_test/", true)
}

```


## Launch Images

TBD

## Assets

```go
package main

import (
	"fmt"
	"log"

	"github.com/illyabusigin/apptools/xcassets"
)

func main() {
	logo := xcassets.Asset("Logo", func(b *xcassets.AssetBuilder) {
		b.Gamut.Any()
		b.Gamut.SRGBAndDisplayP3()
		b.Source.File("./path/to/file.png")

		b.Color(func(d *xcassets.AssetDefinition) {
			d.Devices.Universal()
			d.ColorSpace.SRGB()
			d.Appearance.Any()

            d.Source.File("./path/to/file.png")
		})
	})

	if err := logo.Validate(); err != nil {
        log.Fatal("Failed validation", err)
    }

	err := logo.SaveTo("./_test/", true)
}

```

## Colors

```go
package main

import (
	"fmt"
	"log"

	"github.com/illyabusigin/apptools/xcassets"
)

func main() {
	splashScreenColor := xcassets.Color("SplashScreenColor", func(b *xcassets.ColorBuilder) {
		b.Gamut.Any()
		b.Gamut.SRGBAndDisplayP3()
		b.Color(func(d *xcassets.ColorDefinition) {
			d.Devices.Universal()
			d.ColorSpace.SRGB()
			d.Appearance.Any()

			d.Hex("#262D44")
			d.Alpha(.4)
		})
	})

	output, err := splashScreenColor.Build()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)
}

```
