// Code generated by go-bindata.
// sources:
// assets/badge.svg
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
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
	info  os.FileInfo
}

type bindataFileInfo struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
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
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _assetsBadgeSvg = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x54\x5d\x6f\xda\x3c\x14\xbe\xe7\x57\x1c\xb9\x7a\xa5\xb7\x12\x10\x87\x0c\x4a\xbb\x98\x5e\xac\x2a\xbb\x99\x34\xed\x83\x7b\x27\x71\x12\xaf\xc6\x41\xb6\x21\xd0\xaa\xff\x7d\xc7\x1e\xb4\x54\x22\xa3\xd2\xc4\x45\xe2\xc3\xf3\x71\xce\x79\xac\xa4\xb7\xdb\xa5\x82\x8d\x30\x56\x36\x9a\x91\x78\x48\x09\x08\x9d\x37\x85\xd4\x15\x23\x3f\x7f\xdc\x0f\xa6\xe4\x76\xd6\x4b\xed\xa6\x02\x44\x6a\xcb\x48\xed\xdc\xea\x26\x8a\xda\xb6\x1d\xb6\xc9\xb0\x31\x55\x34\xa2\x94\x46\x88\x20\xd0\xca\xc2\xd5\x8c\x60\x81\x40\x2d\x64\x55\x3b\x46\xa6\x63\x32\xeb\x01\xa4\x4a\x6a\xc1\xcd\xdc\xf0\x42\x0a\xed\x40\x16\x8c\x64\x04\xb6\x23\x46\x10\xbc\xc3\x47\x4c\xe9\x7f\x01\x8a\x60\xeb\x9a\x15\x34\x65\x69\x85\x0b\xff\xfb\xf3\x20\x6f\x54\x63\x18\xb9\xc8\xb2\x6c\x5f\x69\x56\x3c\x97\x6e\xc7\xc8\x30\x26\x10\x9d\xe2\xc6\x9d\xc8\x34\x7a\xdb\x51\xa8\x2d\xb9\x7d\x08\xad\xf1\x43\x27\x46\xe4\xae\x6b\x2e\x30\x5b\x46\x12\x02\xa5\x54\x0a\xfb\x2a\xcb\xf2\x45\xdb\x0b\x85\xb7\x0a\xfc\x2b\x23\x6b\xa3\xfe\xbf\xe0\x97\x07\xd9\x15\x77\xf5\x81\x37\x1e\xa3\x14\x7a\x7e\xa1\x40\x6b\xf4\xd8\x7c\x98\x7c\xa6\x8f\xaf\x03\x1d\x61\x9f\x9e\x60\xf8\xc9\xaf\x01\x9e\x9f\x0f\x1c\x44\x7b\x52\x72\xdd\x45\x0a\xd6\xd9\xe5\x5b\x8f\xe9\xf8\x15\x9e\x46\xd5\xbe\xd7\xe3\x49\x9c\xd8\xba\x01\xd7\x79\xed\x97\xbe\x94\x45\xa1\x04\x8e\xda\x68\x37\x28\xf9\x52\x2a\xdc\xe5\x9d\xf8\xc5\x17\x6b\xf8\xce\xb5\xed\x2f\x84\x29\xb8\xe6\xfd\xb9\xd0\x62\xc3\xfb\x16\x6b\x03\x2b\x8c\x2c\xf7\x1c\x2b\x1f\x05\xc6\x31\x3a\x2c\xc0\xab\xc3\x36\xa4\x8e\xf1\xe3\xf3\xea\x65\x8f\x34\xf6\xbf\x3f\xc7\xa3\xe0\x12\x1f\xe5\x4e\xa1\x4a\x10\x6c\x43\x0e\x37\x59\xa3\x8a\x8f\x64\x36\xff\x3a\x87\x85\xb7\x93\x39\x77\x78\x97\xe1\x9b\xb0\x6b\xe5\xd0\xdb\xa4\x91\xf7\xea\xb2\x9d\xfc\xa3\xea\x69\xd9\x64\x72\x7e\x9a\x99\x0f\xf3\x5e\x2a\xa1\xf9\x52\x60\x9e\x7f\xed\x33\x19\x77\xe2\x4f\x13\x26\xf1\x3b\x3b\xd8\x8f\x74\xc6\x7f\x42\x3b\xd0\xa7\xe1\x53\xfa\x0e\xf7\xbc\x16\xf9\x83\x28\x00\xf7\xea\xa5\xef\xb8\x3b\xbb\x86\xab\xeb\xf3\xb4\x70\x9b\x53\xff\x41\x9a\xf5\x7e\x07\x00\x00\xff\xff\xd6\xb7\x25\xad\xe0\x04\x00\x00")

func assetsBadgeSvgBytes() ([]byte, error) {
	return bindataRead(
		_assetsBadgeSvg,
		"assets/badge.svg",
	)
}

func assetsBadgeSvg() (*asset, error) {
	bytes, err := assetsBadgeSvgBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/badge.svg", size: 1248, mode: os.FileMode(420), modTime: time.Unix(1445646455, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if (err != nil) {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"assets/badge.svg": assetsBadgeSvg,
}

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
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"assets": &bintree{nil, map[string]*bintree{
		"badge.svg": &bintree{assetsBadgeSvg, map[string]*bintree{
		}},
	}},
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
        err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
        if err != nil {
                return err
        }
        return nil
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

