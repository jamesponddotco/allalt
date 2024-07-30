// Package prompt holds the prompt to be used by the LLM.
package prompt

import (
	_ "embed"
)

//go:embed data/prompt.txt
var System string
