// Code generated by "esc -o esc.go -pkg adi -private kernels.hsaco"; DO NOT EDIT.

package adi

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
		size:    26568,
		modtime: 1601819180,
		compressed: `
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
`,
	},
}

var _escDirs = map[string][]os.FileInfo{}
