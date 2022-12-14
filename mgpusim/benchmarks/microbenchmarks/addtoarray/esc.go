// Code generated by "esc -o esc.go -pkg addtoarray -private kernels.hsaco"; DO NOT EDIT.

package addtoarray

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
		size:    9496,
		modtime: 1601287530,
		compressed: `
H4sIAAAAAAAC/+xaTW/jxBt/7LhN/tm/RE+ouxwwCFRYETdvdNNeaNrSbNWkm27YLgWt2Ik9cU39prFT
JRx2WZC4sBIrPgefgUbiwJ3zHjiwEnwAOBJkeyaxvXVftKDuwT8peex5XuZ5xo9nPC8PP2xu8hy3ChQZ
+BU472IuuGeM+TcDet0vq0EOVuH/kIdZABBCcnE64qI0R8s5qpeE+/koZf54ejOh+zi1uSgN63m+gkjL
Y9SGKGV6/AX1WHy3n7mKcA69sH8edp+5yixcHAJrTx6mjofob9koFUJ6OVp/vbXhi7Nn85qfD0G5ANlJ
bKys3tpotO8EslcB4AotR4aiymYBGYr3O3BQoaD2BrVihdpVsgB5KlsoFPJ7mDiaZa6IDJ+KpffEongv
v42JiXVnJS+KBXEHGXgqI4piS5OJtYZN+cBA5DDvFXWGRtfSQ5ILUaHVQ2XBF2wiU+0jdWrwlo3N9aa4
HuFOPPM9Kov3fG6dqL5LHk5wCxGChnl299HQxhGRhb5mutcXJgId7Yuofm3Cquuaaq6cyNpDeh9va6bC
2A3d6iJ9rd/rYRKV8jxgUncq5al1RSEdG8l4t4/0lYmJKV+WGSfABu6hvu4mR66ZCh4kR+4Fnhx2NTns
anLYa0O/6JwRnxXRv/QsbmqKgs2gOW/1eg52P052cGup+t/Xv3/J9X9yCfXvWOYpebFVO+eLkHp1aV61
+rqrNYimdIamXCfqC3u4bim4TSx70n17AwwiagerBjbdwPkbrMNoEKtvU9amNsBKwC9SdptoR8jFyQJR
4zR85uhddIR7xGKViuLkNdjpG51G+7YzaapSacrZi3AqlNFCg00duXctchg47dssv7+UlyQpf9p3gzfW
X8vMPvf9xIV+19iHwpfB/bfbP//Ihz4xuLDB6HgLKVKkSJHi5QRH+3HOn91lYp3581iFH+CJP9eLDis3
Q9dX/BliaG4qCLPj8Xj8MsbPAz/yhzaef/AQHh/DGL7xPM1B7vsMzI48ma+BH3mt8xPwI2+Y9OQ4XniQ
yWUeZQDKAN/98lUwGpb5DF8DaD/1TNvw+JgXhDKA/TRgPzpOMy5FihQpUqRIkSJFihQpUlwG2F7z7/9j
c/cA85TOUNqm0/3Ydjz8+ffY8uUpg+0r1xJWnZuaeYjJithsboilklSUiuI7iw6RF/HAxcRE+qKuHxkF
m1ifY9ld1HVFxLiKq7jbRYos1zAqVWW8vIwrVaxUSqhbKeHlaq18A8nvAsg6MlXxKNi1PY/9QOECNZy2
juK35ly0/G1/kSE73f+neDVh3R0k03IxSMrQdIYGSKrZlw6QcwD03yt3CUguHrj+HTI0GSTZMgxsuiA5
Q8NFXZCcA8clwVVAYwvzb+mWjPT4av1nG/s79dbW+ovnlRdONrRHkHQeYbJGFNPP0lzkY/nJaDGUn1zo
3AXL21cA4K/x2GL6LD8ZFWNu5WL1X6W2+Vg+Mzof0xdi9HV6ToKPvT+Mzp2YP1MshM+sQPI5lyQDBao7
WYlLOH8yE4ufVbNETRZj1dhU/zihekY/CD/7EIpvBPQ+TPubzAnPrxH2PYQnVH//jPbbTdD/g+pLZ+j/
EwAA//+EgvB6GCUAAA==
`,
	},
}

var _escDirs = map[string][]os.FileInfo{}
