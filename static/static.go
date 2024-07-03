package static

import "embed"

//go:embed sw/*.js wrapper/*.js
var Static embed.FS
