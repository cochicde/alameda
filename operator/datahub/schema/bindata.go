// Code generated by go-bindata. DO NOT EDIT.
// sources:
// kafka-topic.json
// kafka-consumergroup.json
// nginx.json

package schema


import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}


type asset struct {
	bytes []byte
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataKafkatopicjson = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x92\xb1\x6a\xc4\x30\x0c\x86\xf7\x3c\x85\xd0\x9c\xa5\xa5\x93\x5f\xa5" +
	"\x1c\x46\x75\xd4\xf6\x38\x2b\x49\x6d\x79\x08\x25\xef\x5e\x62\xc3\xa5\x26\x29\x85\xd2\xc5\x93\xf1\xaf\x4f\xbf\x7f" +
	"\x19\x7d\x76\x00\x00\x18\xdd\x3b\x0b\x59\x61\x25\x34\x50\xc4\x5c\x50\x7a\xf1\x8c\x06\x1e\xfa\x5d\x73\xa4\xfc\x36" +
	"\x85\x05\x0d\xe0\x8d\x5e\x6f\x84\xdf\x8a\xba\xcc\x1b\x8f\x3a\xcd\x57\x87\x59\x5f\x4b\x19\x85\x29\xa6\xc0\xc2\xa3" +
	"\x46\x34\xf0\x7c\x6f\xda\xdf\xcb\xdc\x48\xc2\x77\x6f\x5b\x8c\xfa\x1a\x71\x93\x4f\x32\xd6\x2e\xe7\x6e\x07\x57\xe7" +
	"\x53\x54\x0e\x36\xdf\xfb\x73\x36\xf0\x47\xba\x06\x1e\xd0\x80\x86\xc4\x3f\x50\x25\x84\xdd\x26\x8e\xf5\x17\x55\xd8" +
	"\x40\x4a\x3b\xf4\x74\x80\xd6\x63\xdf\x2f\x23\x6c\x67\x9c\xc9\xb5\x9c\xbf\xd1\xe8\xe4\x49\x78\x20\x1b\x1d\xf9\xb6" +
	"\x97\xe8\x64\x92\x96\x77\x6a\x48\x22\xcb\xbf\x66\x7f\xfc\x6b\xf6\x4a\xb9\x74\xb5\x7e\xe9\xd6\xaf\x00\x00\x00\xff" +
	"\xff\x30\xc9\x13\x11\x74\x05\x00\x00")

func bindataKafkatopicjsonBytes() ([]byte, error) {
	return bindataRead(
		_bindataKafkatopicjson,
		"kafka-topic.json",
	)
}



func bindataKafkatopicjson() (*asset, error) {
	bytes, err := bindataKafkatopicjsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "kafka-topic.json",
		size: 1396,
		md5checksum: "",
		mode: os.FileMode(438),
		modTime: time.Unix(1585192817, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataKafkaconsumergroupjson = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x97\xdd\x6a\xeb\x30\x0c\x80\xef\xfb\x14\xc6\xd7\xbd\x39\x87\x0d\x46" +
	"\x5e\x65\x14\xa3\x29\x6a\x67\xe2\xbf\x59\x76\x69\x36\xfa\xee\x23\xe9\xd6\xce\x6b\x07\x1d\x14\x32\xe7\x2a\x20\x7f" +
	"\x96\xf4\x45\x89\xc1\x6f\x0b\x21\x84\x90\x8c\xcf\x64\x41\x59\x4a\x20\x1b\x71\x08\x8e\x0b\x09\x9e\x0c\xc9\x46\xfc" +
	"\x5b\x9e\x62\x08\x89\x36\x3e\xf6\xb2\x11\xb2\x83\x75\x07\xf2\xcb\x62\xea\xc3\xc0\x4b\xf4\x8e\xb3\xa5\xa8\x36\xd1" +
	"\xe7\x20\x47\x60\x7f\xe0\xa4\x25\xe0\x1c\xc9\x92\x4b\x2c\x1b\xf1\x78\xdc\x7d\x2a\x3c\x72\x0e\x2c\x1d\x8b\xa8\x6f" +
	"\x19\x97\x25\x8b\xde\x64\xeb\xca\x74\x97\xd3\x9e\xa5\x1f\x9f\xcb\xcb\x4c\xa4\x97\xac\x23\xb5\xb2\x11\x29\x66\xfa" +
	"\x81\x3a\x14\x57\x83\x3b\x97\x2f\xab\xc0\x5a\x48\x70\x82\xee\xce\xa0\xfd\xf9\xbe\x2b\x5a\xe7\x00\x58\x6b\xff\x68" +
	"\x32\x27\x8a\xaa\xe2\x11\x24\x1f\x34\xd6\x2c\x00\x06\x2c\xb5\xa0\x18\xc1\xd4\x3d\x8a\x0b\x26\x35\xff\x1c\x91\xd8" +
	"\xe7\x88\xa4\xba\x07\xfe\x95\xcc\x1a\x0c\x5f\x69\xf3\x7f\x3a\x9b\x59\x88\x74\xda\xb5\xb3\x10\x89\x14\x8c\x46\xe0" +
	"\x09\x64\xee\x6f\xed\xc2\x81\x70\x56\x42\x56\xbb\x79\xf9\xc0\x6e\x36\x3e\x98\x39\x79\x5b\xf5\x89\x16\xbc\xd1\xd8" +
	"\xd7\xda\x3d\xb9\xe1\x8e\xa2\x68\x47\x98\x93\xf6\x6e\x0a\x8f\x9b\x7e\x50\x21\x2b\xa3\xad\x4e\xb5\x0e\xa4\x30\x19" +
	"\x9a\x25\xae\xdf\xc5\x92\xf5\xb1\x9f\xc9\x60\x3e\x64\x2a\x9f\xcd\x76\x28\x4c\xac\x58\xbf\x56\x7b\xf4\x7e\x3a\x84" +
	"\x2d\xfe\x25\x8f\x22\xb2\x5a\x94\xf1\xd5\x62\xff\x1e\x00\x00\xff\xff\x02\x03\x3b\x96\x37\x12\x00\x00")

func bindataKafkaconsumergroupjsonBytes() ([]byte, error) {
	return bindataRead(
		_bindataKafkaconsumergroupjson,
		"kafka-consumergroup.json",
	)
}



func bindataKafkaconsumergroupjson() (*asset, error) {
	bytes, err := bindataKafkaconsumergroupjsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "kafka-consumergroup.json",
		size: 4663,
		md5checksum: "",
		mode: os.FileMode(438),
		modTime: time.Unix(1585192817, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataNginxjson = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x97\xdb\x8e\x9b\x30\x10\x86\xef\xf3\x14\x96\xaf\x73\xd3\xaa\xad\x2a" +
	"\x5e\xa5\xaa\xac\xa9\x99\x12\x2b\x3e\xd5\x1e\x57\xa0\x8a\x77\xaf\x80\x55\x58\x02\xd1\x26\x5a\xb2\x1c\xae\x22\xcf" +
	"\xfc\xb6\xe7\xcb\x8f\x4f\xff\x0e\x8c\x31\xc6\xa3\x3c\xa1\x01\x61\x90\x80\x67\xac\x0b\xb6\x09\x82\x5f\x1a\x79\xc6" +
	"\x3e\x1d\xfb\x98\x04\xc2\xc2\x85\x8a\x67\x8c\xdb\x42\xd9\x92\xbf\x4a\x52\xe5\xb1\x4f\xb4\xf1\xba\x4b\x73\x83\x10" +
	"\x53\x40\x83\x96\x22\xcf\xd8\x8f\x4b\xa7\x7e\xbe\x56\x67\xc1\xe0\xd4\xd8\xdd\xe4\x4e\x27\x63\x87\xfd\xa7\xc7\x19" +
	"\x8d\x27\x75\x8a\x84\x41\xb4\xed\xe3\xb4\x36\xe0\x9f\xa4\x02\xe6\x3c\x63\x14\x12\xde\x50\x75\x45\x88\x86\x35\x0e" +
	"\xff\x9c\x81\x2c\x07\x82\x5e\xf4\x65\x24\xaa\xc7\xfd\xde\x40\x68\x7e\xa3\x07\xb9\x96\xfa\xaf\x23\x8f\xf2\x80\x06" +
	"\x83\x39\x88\x28\x41\x6f\xdb\x99\x09\x92\x35\x19\xf5\x28\x4e\xc0\xe8\x52\x90\x28\xce\xdf\xa3\x88\x18\xfe\x2a\x89" +
	"\x5b\xb6\xe7\x26\xcf\x6e\x4c\x7a\x04\xe6\x37\xe8\xb8\x7e\x9a\x5d\x80\x9c\x95\xcd\x77\x01\x12\xd0\x6b\x25\x21\xce" +
	"\x0b\xf3\xf9\x1e\x98\xaf\xb3\x6f\x00\x1e\xe5\xae\x80\x8c\xb2\xfb\xe2\x81\x72\x6d\x3c\xd7\x91\xf7\x2d\x26\x97\xe8" +
	"\xee\xf3\x74\x6e\xba\x27\xec\x0d\x17\x9c\x27\x9c\x40\x1f\xc7\x84\xa5\x77\xa1\x79\x21\x78\x97\x2f\xf1\xd9\xcd\x0b" +
	"\xb1\x79\x3b\xbc\xd3\x4a\x56\x5b\xad\x1e\x6d\xf3\x74\x16\x58\xa2\x4c\xa4\x9c\x5d\x82\x63\xce\x75\x2e\x7d\x12\x5a" +
	"\x19\x45\x5b\x35\x64\x40\xd2\x14\x8b\x71\xfb\x2c\x06\x8d\x0b\xd5\x4e\x8c\x79\x81\xd9\xbc\x37\xed\xbd\x45\x18\x08" +
	"\x85\xb2\xc2\x63\x90\x68\x09\x8a\x25\xf6\xe1\x39\x2e\x66\x27\x22\x2f\x02\x46\xef\x6c\x44\x41\x6a\x91\x3b\xcb\xb7" +
	"\x31\xc8\x75\x60\xa0\xf8\x79\x69\xd5\x87\xae\x5d\x1f\xfe\x07\x00\x00\xff\xff\x0c\xa9\xac\x03\xec\x14\x00\x00")

func bindataNginxjsonBytes() ([]byte, error) {
	return bindataRead(
		_bindataNginxjson,
		"nginx.json",
	)
}



func bindataNginxjson() (*asset, error) {
	bytes, err := bindataNginxjsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "nginx.json",
		size: 5356,
		md5checksum: "",
		mode: os.FileMode(438),
		modTime: time.Unix(1590721749, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}


//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"kafka-topic.json":         bindataKafkatopicjson,
	"kafka-consumergroup.json": bindataKafkaconsumergroupjson,
	"nginx.json":               bindataNginxjson,
}

//
// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
		}
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}


type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"kafka-consumergroup.json": {Func: bindataKafkaconsumergroupjson, Children: map[string]*bintree{}},
	"kafka-topic.json": {Func: bindataKafkatopicjson, Children: map[string]*bintree{}},
	"nginx.json": {Func: bindataNginxjson, Children: map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
