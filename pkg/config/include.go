package config

import (
	"io"
	"os"
)

// IncludeFromURL allows to import a config from an URL, passing headers if needed (e.g. authorization).
type IncludeFromURL struct { // pro only
	URL     string            `yaml:"url" json:"url"`
	Headers map[string]string `yaml:"headers,omitempty" json:"headers,omitempty"`
}

// IncludeFromFile allows to import a config from another file.
type IncludeFromFile struct { // pro only
	Path string `yaml:"path" json:"path"`
}

// Reader ...
func (i IncludeFromFile) Reader() (io.ReadCloser, error) {
	return os.Open(i.Path)
}

// Include allows to import a config file from either a file or an URL.
type Include struct { // pro only
	FromURL  IncludeFromURL  `yaml:"from_url,omitempty" json:"from_url,omitempty"`
	FromFile IncludeFromFile `yaml:"from_file,omitempty" json:"from_file,omitempty"`
	Content  string          `yaml:"-" json:"-"`
}

// String ...
func (i Include) String() string {
	if p := i.FromFile.Path; p != "" {
		return p
	}
	return i.FromURL.URL
}

// IncludedMarkdown type alias.
type IncludedMarkdown Include
