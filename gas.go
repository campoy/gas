package gas

import (
	"io"
	"sync"
)

var (
	fs   *FS
	lock sync.RWMutex
)

func init() {
	fs = GopathFS()
}

// Refresh the internal FS to reflect possible changes in the GOPATH env variable
func Refresh() {
	lock.Lock()
	defer lock.Unlock()
	fs = GopathFS()
}

// Open the file for reading or returns an error
//
// For more information, see the FS type
func Open(file string) (io.ReadCloser, error) {
	lock.RLock()
	defer lock.RUnlock()
	return fs.Open(file)
}

// Return the absolut filepath for the requested resource or return an error if not found
func Abs(file string) (string, error) {
	lock.RLock()
	defer lock.RUnlock()
	return fs.Abs(file)
}