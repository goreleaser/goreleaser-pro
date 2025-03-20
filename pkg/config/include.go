package config

// IncludeFromURL allows to import a config from an URL, passing headers if needed (e.g. authorization).
type IncludeFromURL struct { // pro only
	URL     string            `yaml:"url" json:"url"`
	Headers map[string]string `yaml:"headers,omitempty" json:"headers,omitempty"`
}

// IncludeFromFile allows to import a config from another file.
type IncludeFromFile struct { // pro only
	Path string `yaml:"path" json:"path"`
}

// Include allows to import a config file from either a file or an URL.
type Include struct { // pro only
	FromURL  IncludeFromURL  `yaml:"from_url,omitempty" json:"from_url,omitempty"`
	FromFile IncludeFromFile `yaml:"from_file,omitempty" json:"from_file,omitempty"`
	Content  string          `yaml:"-" json:"-"`
}

// IncludedMarkdown type alias.
type IncludedMarkdown Include

// MarshalYAML implements yaml.Marshaler.
func (i IncludedMarkdown) MarshalYAML() (interface{}, error) {
	if i.Content != "" {
		return i.Content, nil
	}
	type alias IncludedMarkdown
	return alias(i), nil
}
