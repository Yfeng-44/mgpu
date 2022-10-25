// Code generated by "esc -o bindata.go -pkg mis -private kernels.hsaco"; DO NOT EDIT.

package mis

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return fis[0:limit], nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// _escFS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func _escFS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// _escDir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func _escDir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// _escFSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func _escFSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// _escFSMustByte is the same as _escFSByte, but panics if name is not present.
func _escFSMustByte(useLocal bool, name string) []byte {
	b, err := _escFSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// _escFSString is the string version of _escFSByte.
func _escFSString(useLocal bool, name string) (string, error) {
	b, err := _escFSByte(useLocal, name)
	return string(b), err
}

// _escFSMustString is the string version of _escFSMustByte.
func _escFSMustString(useLocal bool, name string) string {
	return string(_escFSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/kernels.hsaco": {
		name:    "kernels.hsaco",
		local:   "kernels.hsaco",
		size:    22352,
		modtime: 1617818674,
		compressed: `
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
8lBXAAA=
`,
	},
}

var _escDirs = map[string][]os.FileInfo{}
