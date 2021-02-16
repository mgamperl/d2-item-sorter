module d2-item-sorter

go 1.15

require (
	github.com/fsnotify/fsnotify v1.4.9
	github.com/gofiber/fiber/v2 v2.5.0
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/mkideal/cli v0.2.3
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/nokka/d2s v0.0.0
	github.com/stretchr/testify v1.5.1 // indirect
	golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/nokka/d2s v0.0.0 => ../d2s
