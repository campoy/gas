package gas

import (
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type FS struct {
	searchPath []string
}

// Used to indicat that the file wasn't found in any possible location
type NotFound string

func (n NotFound) Error() string {
	return "The file " + string(n) + " wasn't found"
}

// Find the absolute path for the required file.
//
// The returned string is OS depended. If the desired file isn't present
// in any possible location returns NotFound error.
func (fs *FS) Abs(file string) (abs string, err error) {
	reqPath := filepath.FromSlash(path.Clean(file))

	for _, p := range fs.searchPath {
		abs = filepath.Join(p, "src", reqPath)
		var stat os.FileInfo
		stat, err = os.Stat(abs)
		if !os.IsNotExist(err) && !stat.IsDir() {
			return
		}
	}
	// if reach this point
	// all possible locations were tested
	// and no match was found
	abs = ""
	err = NotFound(reqPath)
	return
}

// Open the resource for reading
func (fs *FS) Open(file string) (r io.ReadCloser, err error) {
	abs, err := fs.Abs(file)
	if err != nil {
		return
	}

	r, err = os.Open(abs)
	return
}

// Create a new GopathFS instance
func GopathFS() *FS {
	fs := &FS{}
	fs.searchPath = strings.Split(os.Getenv("GOPATH"), ":")
	return fs
}