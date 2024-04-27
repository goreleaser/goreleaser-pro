package config

import (
	"os"
	"time"
)

// Content describes the source and destination
// of one file to copy into a package.
type NFPMContent struct {
	Source      string               `yaml:"src,omitempty" json:"src,omitempty"`
	Destination string               `yaml:"dst" json:"dst"`
	Type        string               `yaml:"type,omitempty" json:"type,omitempty" jsonschema:"enum=symlink,enum=ghost,enum=config,enum=config|noreplace,enum=dir,enum=tree,enum=,default="`
	Packager    string               `yaml:"packager,omitempty" json:"packager,omitempty"`
	FileInfo    *NFPMContentFileInfo `yaml:"file_info,omitempty" json:"file_info,omitempty"`
	Expand      bool                 `yaml:"expand,omitempty" json:"expand,omitempty"`
}

type NFPMContentFileInfo struct {
	Owner string      `yaml:"owner,omitempty" json:"owner,omitempty"`
	Group string      `yaml:"group,omitempty" json:"group,omitempty"`
	Mode  os.FileMode `yaml:"mode,omitempty" json:"mode,omitempty"`
	MTime time.Time   `yaml:"mtime,omitempty" json:"mtime,omitempty"`
	Size  int64       `yaml:"-" json:"-"`
}

// Contents list of Content to process.
type NFPMContents []*NFPMContent
