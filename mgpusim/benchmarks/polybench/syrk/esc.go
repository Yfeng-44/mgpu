// Code generated by "esc -o esc.go -pkg syrk -private kernels.hsaco"; DO NOT EDIT.

package syrk

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
		size:    9552,
		modtime: 1601877166,
		compressed: `
H4sIAAAAAAAC/+xa3W/TVht/7LiJY9LwcQW8F/itXqnvGEmTNOvc3oy0pR2iLYF2QDeh6sQ+SUMdO3Oc
qpk0aBFFXCCB+Ae43cWk/Qekk3Y7aUO75GI3XE7bzXa5TMc+J7FDXVrBxtj8k5LHfr4fn2PrfN06NzfD
c9xZoIjAj8CRiyPuPRMUh1x62uEpIMJZSIAEUQAQPHr9dIfzU5HyOWoXhG8H/ZTlQ+wGPPf9VOH91GtH
cgWZ8vtoHfyU2fEHtGP1XX5ua8I+7Lz5EVx6bmtRODgE9jxZ3X106ZCfCh47kcYvzE876qxt/uP0B5cv
QKxbG+MV5qdnix+5uicA4BDlo5pWUY0Uqmnkt9pAqVSlvKFkRqnf7yUAieqmUinpCrYaVdOYkBk+kbNn
5Ix8XbqALQPrjQlJllPyAqrhno4sy42Wtbay5qhI5H6xVSuZukdt2KNxdk0bdrTmkFFpokrP1cU6Nqbm
5CmftJuTk0tOvu5IC1bFSYZgl4SQxK6WWnXsEw9PF5YKK0vLxXOnh7tai9XP/A6UrqigVyvGxK6iK0hv
4gtVQ2PiWd0sIX2yWS5jy69F0mBaM6O5nndNsxbrSMWXmkif6LroyVWVSVxM4zJq6nZw6eq/t3Sk11f3
aPlu9cG154NrzwfXPtlyWPss+8BllbD9D6zKqAbXVDXsv7ya869WzY23uJrX9An4sKpp2HDf4ovlcgPb
1/ZIcCz/58dffsPxP34D8RdMY69+oezz+xtm9caymm/qdnXWqmqLLUMtWJVXznDK1HDRMuvdERMZzSGr
sogrNWzYbvIK8zVrmc06Fc1UN7DmyjNUXLSq68jGwQp+57R85vwqWsdly2RBZbn7Giw0a4uzxcuN7qPK
Kj3JFZ9knArm0caMjuyrprXmJu34zL03JqXTaekl43Qytj4Zib4wX+E8v5O9CZRI7h989cNPPB3Kc/0O
PUNcCBHigOC6/Up0Z3bc3vq34Ut4GCdzPX9Pv+a5PgZx/9xUEKKdTqfzd6w/EpF2jjiVSztkan4L7rcn
gd8R6DXXidwlmYsR8RGIsAUg5riH8CAO8Ud8HLaAj+YOR+99DvDNd4dFHuDG5l2Qbz8Zg3ttgO2nkVOu
bxIruSW0IQ65x8APATx4eht4J0aEF27yADk+wisAxWekSepwvz0oCBJA/RlpoCHYbguw/TQp8hAlNMoD
xydvPhaloQjxRWJHo7lYXFQeJ5JDA4SXcHVAknKxwYSShM0nxP8AQIzQCIB4JxrNbYqickeScvA1bG0m
EgpsxtqfcvfbiWRSInnEjgpnWB4/d7bbAFvt8M0JESJEiBAhQoQIESJEiLcLbK/5Gt1np9vDcJzSAUp3
qJzN+tm2/K+/d0xC71EG21f+YnD3eHNVYw1bE/Lc3LSczaYz6Yz8/5GGpY7gDRtbBtJHdH29lqpb5g2s
2iO6rskY53Eel0pIU1UFo2xexePjeDSPtdEsKo1m8Xheyb2P1HcAVB0ZFXnd3bvdj3/X4AAR9lpHIU/z
l6N+fozy88f8/EH29I/4+e+SPz7WOy9AcSpg3RDShmljSGsto9GqQbpiNNOrqLEK9J/wbQvSNt6wnTtU
q6qQVs1aDRs2pButmo1KkG6sNmzLvXIpTE5mVkad/7x3kfF/uqki3bfsuDK9vFCYPz/1utajYp7jC0Hn
GLprS/Di8z7kMWP9mtGMp19znvMarL8fBoDfOh2T2bN+zajcl5bYF/8E9c33vQeMHu+zF/roKXq+gu97
7xgd3LXf9TDsPesCwedjghykqG2EMQLOrQz01c/CjFGXmb4wdWrfDgjP6Afetvcg81+XPoTed2pgl/ab
9ebuwQ61X37J87sUYH+Cnn/Kv8T+jwAAAP//ThsohVAlAAA=
`,
	},
}

var _escDirs = map[string][]os.FileInfo{}
