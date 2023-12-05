//go:build !noembed

package assets

import "embed"

//go:embed templates
var TmplFiles embed.FS
