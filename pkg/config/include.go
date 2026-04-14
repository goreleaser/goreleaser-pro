package config

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/avast/retry-go/v4"
)

// IncludeFromURL allows to import a config from an URL, passing headers if needed (e.g. authorization).
type IncludeFromURL struct { // pro only
	URL     string            `yaml:"url" json:"url"`
	Headers map[string]string `yaml:"headers,omitempty" json:"headers,omitempty"`
}

// Reader ...
func (i IncludeFromURL) Reader() (io.ReadCloser, error) {
	url := i.URL
	if !strings.HasPrefix(url, "http") && strings.Contains(url, "/") { // assume it is a path within github
		// TODO: maybe deprecate this? it doesn't really save much space...
		url = fmt.Sprintf("https://raw.githubusercontent.com/%s", i.URL)
	}
	//nolint:noctx
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range i.Headers {
		req.Header.Add(k, os.ExpandEnv(v))
	}
	var body io.ReadCloser
	err = retry.Do(func() error {
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		if resp.StatusCode >= 500 || resp.StatusCode == http.StatusTooManyRequests {
			resp.Body.Close()
			return fmt.Errorf("%s %q: %s", "Get", i.URL, resp.Status)
		}
		if http.StatusOK != resp.StatusCode {
			resp.Body.Close()
			return retry.Unrecoverable(fmt.Errorf("%s %q: %s", "Get", i.URL, resp.Status))
		}
		body = resp.Body
		return nil
	},
		retry.Context(context.Background()),
		retry.Attempts(3),
		retry.Delay(2*time.Second),
		retry.MaxDelay(15*time.Second),
		retry.DelayType(retry.BackOffDelay),
		retry.LastErrorOnly(true),
	)
	return body, err
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

// Load the included content.
func (i Include) Load() (io.ReadCloser, error) {
	if i.FromFile.Path != "" {
		return os.Open(i.FromFile.Path)
	}

	if i.FromURL.URL != "" {
		return i.FromURL.Reader()
	}

	return io.NopCloser(strings.NewReader(i.Content)), nil
}

// IncludedMarkdown type alias.
type IncludedMarkdown Include

func (i IncludedMarkdown) Load() (io.ReadCloser, error) {
	return Include(i).Load()
}
