package loader

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/net/webdav"
)

// Fileb0xLoader provides a Pongo2 loader for fileb0x template files.
type Fileb0xLoader struct {
	FS  webdav.FileSystem
	CTX context.Context
}

// Abs returns the absolute path to a template file.
func (f *Fileb0xLoader) Abs(base, name string) string {
	return name
}

// Get retrieves a reader for the specified path.
func (f *Fileb0xLoader) Get(path string) (io.Reader, error) {
	r, err := f.FS.OpenFile(
		f.CTX,
		fmt.Sprintf("templates/%s", path),
		os.O_RDONLY,
		0,
	)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}
