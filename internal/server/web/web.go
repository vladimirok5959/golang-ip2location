package web

import (
	_ "embed"
)

var (
	//go:embed index.html
	IndexHtml string

	//go:embed styles.css
	StylesCss string
)
