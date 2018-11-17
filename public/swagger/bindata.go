// Code generated by go-bindata.
// sources:
// public/swagger/swagger.json
// public/swagger/swagger.yaml
// DO NOT EDIT!

package swagger

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
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
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

var _publicSwaggerSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x56\x4f\x6f\xe3\xb6\x13\xfd\x2a\xc4\xe4\x77\xf8\xb5\x50\x62\x65\x37\x0d\x5a\x15\x3d\x14\x4d\x83\x06\x41\xb1\x41\x9a\xee\xa5\x08\x8a\x31\x39\xb2\xb9\xa0\x48\x2d\xff\x38\x09\x0c\x7d\xf7\x62\x28\xc9\x96\x9d\x7a\x93\x00\xed\xad\x27\x1b\xe4\xf0\xcd\xcc\x7b\x8f\x23\xae\x21\x3c\xe0\x62\x41\x1e\x2a\x78\x77\x52\x42\x01\xda\xd6\x0e\xaa\x35\x44\x1d\x0d\x41\x05\x77\x4b\x12\xd2\x35\x8d\xb3\x22\x90\x5f\x69\x49\xe2\xc7\x9b\xab\x00\x05\x28\x0a\xd2\xeb\x36\x6a\x67\x73\x9c\x0e\x42\x07\x11\xff\x3e\xbe\x10\x0f\x4b\x2d\x97\x42\xba\x64\x94\x98\x93\x48\x81\x94\x98\x3f\x09\xb4\x4f\xce\x12\x14\xb0\x22\x1f\x7a\x28\xe8\x0a\x58\xba\x10\xa1\x02\xe3\x24\x1a\xfe\x5f\xbd\x2f\xcb\x53\x28\x20\xc8\x25\x35\x14\xa0\xfa\x03\x96\x31\xb6\x70\x5f\x80\x74\x36\xa4\x61\x0d\xdb\xd6\x68\x89\x5c\xd3\xec\x53\x70\x16\x8a\x9d\xa5\xc7\xc6\xec\xad\x2c\xdc\x7c\x3f\xe6\x98\xd7\xee\x0b\x68\xbd\x53\x49\xfe\x0b\xb8\x18\x97\x81\x29\x9e\xe5\x66\x70\xb6\xae\xb5\x21\x5e\xed\x78\x75\x41\x91\x7f\x42\x6a\x1a\xf4\x4f\x50\xc1\x85\x7b\xb0\xc6\xa1\x12\x6d\x9a\x1b\x2d\xc7\x53\x50\x80\x6b\xc9\x67\xf0\x2b\x05\x55\x4f\x0d\x1e\x8d\xfb\x5f\x8f\xa8\xc0\x29\x3d\x36\x14\xc9\x73\x33\x6b\xb0\xd8\xb0\xb4\x93\x00\xcd\xbc\x0f\xff\x77\x75\xbd\x25\x83\x51\xaf\x48\x70\xb4\x18\x42\x3c\x7d\x4e\xda\x93\x82\x2a\xfa\x44\x05\xc4\xa7\x96\x01\x43\xf4\xda\x2e\xa0\xbb\xe7\x88\xd0\x3a\x1b\x28\x37\xfa\xae\x2c\xf9\x67\x17\xf8\x92\xf1\xd4\xd0\x1b\xa9\x51\x5b\xcc\xe6\xeb\xf1\x38\x25\x74\x5d\x01\x67\xe5\xd9\x01\x00\xeb\xa2\xa8\x5d\xb2\xbb\xe7\xff\xe7\xa9\x86\x0a\x8e\x66\x8a\x6a\x6d\x35\x1f\x08\x33\xf2\xde\x79\xe8\x18\xef\x99\x8b\x78\x71\x16\xa2\xf3\xb8\x20\x06\x68\xb3\xfd\xd6\x10\x71\x91\xa3\xc6\xad\xfb\x62\x22\x0c\x2a\x25\xc6\x8d\x7d\xda\x7e\x6f\xb3\x64\x99\x34\x6d\x45\x93\x4c\xd4\x2d\xfa\x28\x98\x3a\x0a\xf1\xb9\x7a\x3d\xd0\x11\x2a\x6e\x65\xea\xe9\xcd\xd9\x59\xed\x7c\x73\xac\x30\xe2\xbe\x3d\x23\x3d\xc6\x59\x6b\x50\xdb\xde\x60\x87\xd4\x1e\x95\x66\xa0\x0b\xc6\xd9\x2f\x7b\x08\xda\xea\x5b\xa3\x09\x5b\x81\x7b\x41\x5e\x25\xef\x87\x6b\x78\x91\xe9\xd9\x1a\x95\xf2\x14\xc2\xd4\xf8\x5f\xa4\x5c\x62\x3c\x48\xf9\x4f\x18\xf3\xf4\x19\x39\x57\x9a\xfd\x38\x4f\x91\x36\x32\xe5\x81\xd3\xa7\x3c\x28\x80\x44\x16\xe7\xad\xec\x6e\x51\x0f\x5f\xa5\xab\x9b\xcb\xdf\x26\xe9\xff\x89\x4b\xf4\x65\x96\xfb\xd1\xfe\xe6\xf9\x32\x1c\x7b\xce\x50\xbf\x71\xb4\x89\xf8\x6f\xc4\xbc\x69\xc4\x74\xdc\xf8\x26\x98\x41\xfa\xf8\xc9\xd7\xf6\x57\x52\x1a\xb9\x34\xa1\x15\xd9\xa8\x6b\x4d\xbe\x12\xd3\xcf\xc8\xca\xaa\x93\x85\xc3\x93\x7c\xf4\x7b\xb1\xd2\xf4\xf0\x83\xa2\x1a\x93\x61\xdb\x0e\x5d\xb9\xf9\x27\x92\x83\x8d\x5b\xf2\x51\xf7\x24\x49\xa7\x68\xd2\xfb\xc0\xe5\xbe\x1c\x68\xa7\x09\x8f\x43\x4b\x52\xd7\x5a\x8a\x9c\x51\x30\x46\x21\xe8\xb1\x65\x17\x93\x12\x18\x04\x8a\x1e\x49\xac\xd0\x24\x3a\x81\x02\xe8\x11\x9b\x36\x37\xa4\xed\x0a\x8d\x56\x7f\xe6\x2d\xc8\x0c\x44\xd4\xe6\x15\x55\x88\x65\x6a\xd0\x1e\x7b\x42\x85\x73\x43\x9c\xd2\xa0\xcd\x35\x89\x4d\x4d\xd1\x89\xc8\x0f\x0f\x27\x65\xf2\x9e\xac\x24\xe1\xea\x3c\x05\x5a\xef\xe6\x86\x9a\xdd\x6a\x3e\x72\x15\x1c\x71\x75\x21\x9a\x14\x22\x3f\x44\xd0\x0a\x6d\x23\xf1\x1b\xa8\x2b\x40\xab\xd7\x94\x96\xac\xfe\x9c\xa6\x1a\x89\xda\xf9\xbe\x12\x1e\xd4\x5a\x26\x83\xfe\xb5\x45\xbd\xbf\x3c\xbd\xbc\xfe\x78\x7b\xcb\xe9\x1b\x8a\x53\x77\x6e\x74\xdc\x2f\x80\xe3\x44\xbf\x2b\xa4\xb3\x11\xb5\x65\x01\x2c\xeb\x15\xd1\x2a\xf4\x2a\xc7\x1c\xf3\x73\xce\x37\x3d\x6b\x38\x77\xa9\x1f\x91\x59\xca\x9d\x2a\xd8\x82\x0d\x85\x88\x4d\x0b\xd5\xe9\xd9\x37\xdf\x9e\x97\xdf\x95\xe7\xe7\x5d\xc1\x93\x2d\x1b\x16\xcd\xcd\xc4\x4c\x7c\x29\xd9\xe5\x11\x63\x0a\x2f\x53\xc6\x49\x7f\xb9\xbb\xbb\x11\xfd\x81\xec\xa2\xd1\x65\x2c\xee\xa8\xe3\x40\xd0\x1b\x0c\x76\x56\x96\xd0\xdf\xac\x69\xbe\x9f\xb3\x57\xc7\x09\x21\x1a\xbe\x56\x22\xdf\xab\xff\x0f\x97\x25\xdf\x9c\xaf\x76\x29\xe8\xef\xc7\x9e\x67\xb7\x96\x7d\xc9\x3e\xbd\x7b\xb6\x7a\x4e\xe4\x3c\xc0\xed\x48\xdf\xd0\x45\xb7\x37\xd4\x3e\x5c\x1f\x98\xf8\xdd\x5f\x01\x00\x00\xff\xff\x59\x1a\x02\x1a\xb8\x0b\x00\x00")

func publicSwaggerSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_publicSwaggerSwaggerJson,
		"public/swagger/swagger.json",
	)
}

func publicSwaggerSwaggerJson() (*asset, error) {
	bytes, err := publicSwaggerSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "public/swagger/swagger.json", size: 3000, mode: os.FileMode(420), modTime: time.Unix(1542407286, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _publicSwaggerSwaggerYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x57\xdf\x6b\xe4\x36\x10\x7e\xf7\x5f\x31\x6c\x0a\xe9\x8f\xdb\x5d\xe7\x2e\x0d\x3d\x97\x3e\x94\xa6\xa1\x21\x94\x84\x34\xbd\xd7\x32\x2b\x8d\x77\x75\xc8\x92\x4f\x1a\x6d\x12\x8e\xfe\xef\x45\xb2\xbd\x6b\x3b\xce\x65\x5b\x0a\xa5\x70\x6f\xcb\x68\x46\xfa\xe6\xfb\xbe\x19\xb3\xc2\x1a\x1f\x2a\xf2\x45\x36\x07\xac\x6b\xad\x04\xb2\xb2\x66\xf9\xde\x5b\x33\x0a\x3d\x54\x7a\x14\x59\xdb\xd5\x38\x67\x1e\x63\x92\x4a\x65\x54\x0c\xf8\x22\x03\x20\xe7\xac\x8b\x3f\x00\x24\x79\xe1\x54\x1d\x8f\x0a\xf8\x39\xc6\xc1\x91\xaf\xad\xf1\x04\x15\x49\x85\xc0\x8f\x35\xc1\x97\x92\x4a\x0c\x9a\x61\xab\xe8\xfe\xab\x54\x49\x0f\x58\xd5\x9a\x9a\x6b\x00\x84\x95\x54\x80\x32\x5b\xd4\x4a\xfe\xb1\x45\x1d\xa8\x3d\x91\xc4\xa8\x74\x01\xef\x62\x0c\x6c\x09\x97\xe7\x50\x05\xcf\xb0\x22\x40\x03\xca\x30\xad\xc9\xb5\xc9\x4a\x16\xf0\xe6\xe2\xe4\xe2\xea\xdd\xed\x6d\x1b\xaa\x88\xb1\x7b\x05\x80\x55\x45\x9e\xb1\xaa\x0b\x38\x59\x9c\x7e\xfb\xdd\x59\xfe\x36\x3f\x3b\xa3\x6f\xf2\xb7\x6d\x8a\x67\xe4\xe0\x0b\x98\x9d\xe6\xf9\x2c\xc5\x6a\x67\x6b\x72\xac\xc8\x0f\xc0\xee\xae\x1c\x70\x80\xa6\xcf\xdf\xdc\xd7\x24\x54\xa9\x44\xc3\x59\x2a\x7c\x05\xf4\x50\x3b\xf2\x9e\x24\xa0\x07\x04\xcf\x4e\x99\x35\xa4\x9e\x17\xbb\x5b\x3b\x7e\x26\x39\x81\x44\x6b\xd1\x96\x0e\x89\x7a\x06\x17\x6c\x42\x85\x66\xee\x08\x25\xae\x34\x45\x10\x1a\x4d\x42\x09\x3b\x94\x6c\x81\x37\xca\x83\x15\x22\x38\x47\x46\x24\xc2\x79\xb3\x7f\x37\xd1\xb1\xd2\x54\x4d\x20\x3d\x48\xa1\x49\xe8\x4a\x3e\x0b\x3b\x18\xf5\x21\x10\x28\x49\x86\x55\xa9\xc8\x41\x69\x5d\x83\xb2\x46\xc7\x4a\x04\x8d\xee\x29\xe0\x4f\xc0\x1c\xf9\x63\x12\xd0\xd0\x34\x28\x65\x72\x3f\xea\x9b\xbd\x17\x80\x5d\x4f\x8f\x11\xea\x58\x0f\x76\xf5\x9e\x04\x83\xb0\x86\x51\x99\xa8\xb1\x89\x96\x60\x34\x12\x9d\x4c\x39\x73\x65\x4a\xeb\xaa\x46\x06\x5c\xd9\xc0\x3d\xa6\x63\x1b\xc9\x37\x4f\x9b\xe8\x67\xbd\xe4\xe8\xae\xc1\x06\xce\xd0\xe6\xd3\xf8\xe3\xc3\xbf\xdc\xdd\xdd\xb4\x59\xc9\xb6\x9d\xad\xa3\x77\x3a\x9b\xb4\x1c\x0f\x1d\xdd\x83\xf6\xa2\xb7\xf7\x43\x36\xa9\x03\x2b\x8e\x59\xc7\xbf\xc6\x4d\x92\x16\xc9\xde\x06\xc5\x60\x4f\x6d\x8d\x5c\xac\x2d\x2e\x12\x5d\xdf\xa7\x35\xf3\x43\xbb\x73\x8e\xb3\x27\x14\x6c\xac\xe7\x02\xb4\x15\xa8\xd3\xcf\x37\x79\x7e\x92\x45\x25\x22\x1f\x03\x26\xee\x62\x9b\xca\x27\x46\x84\xad\xaa\x38\x2c\xe4\xb6\x4a\x10\xfc\x78\x73\xe9\x5f\xc1\xfd\x46\x89\x0d\x08\x1b\xb4\x8c\x76\x0f\x91\x83\xd5\x23\xa0\x79\xb4\x26\xba\xa3\x6d\xe1\x6e\xba\x3e\x03\xd8\x92\xf3\xe9\xa9\xd9\x2c\xab\x91\x37\x49\x93\xa5\x17\x1b\xaa\x70\xf9\xb1\x54\x9a\x62\xf4\xcf\x46\xa9\x35\x71\x27\x59\xf4\x61\xea\xfd\x52\x16\xd0\xa4\x1f\x75\x65\x5f\x77\x65\x6d\x6e\x8d\x0e\x2b\x62\x72\x3b\xc1\xe7\xc3\x36\x6f\x49\x23\xab\x2d\x41\x2c\x84\x5e\x25\x80\x32\xc5\x30\x60\xb0\xa2\x02\x46\x2f\x00\x38\xfa\x10\x94\x23\x39\x9a\x8b\x89\xd9\xea\x3e\x10\x3d\xf7\xcd\x5e\xe7\xf9\xac\x6f\xea\x01\xba\x8b\x08\x4a\xda\x7b\xa3\x2d\x4a\x92\xbd\xb4\xa6\xdf\x7e\x61\xf7\x62\xc4\xb7\xbf\xfe\x34\x3f\x7d\xe1\x7a\x63\x19\x4a\x1b\xcc\x4b\xb7\x7f\xe1\xa8\x2c\xe0\xf8\x68\xd9\xfb\x26\x2e\x93\xeb\x8e\xb3\x5e\x15\xf5\x98\xde\x30\xd7\xdd\x59\xa8\x2a\x74\x8f\x05\x9c\xb7\xdd\x40\x1d\x56\x5a\x89\x4e\xb8\x24\x3d\x5b\x87\xeb\x76\xc4\xeb\x68\xcf\xdd\x57\xa7\xfb\xb4\x77\x37\x57\x41\xb3\x8a\x5b\x70\x19\x97\xc8\x5c\x22\x63\x36\xd1\xe0\xef\x75\x7a\x2a\x69\xab\xcc\xbe\x2a\x69\x46\x9e\x27\x1d\xd5\xa0\x38\x42\x29\x0f\x35\xd1\x80\xf2\x68\x9b\x08\xea\x7c\x8f\xa9\x6f\x9d\x09\xdb\x94\xa8\xfd\xd8\x37\xbd\xd4\xda\x59\x19\x44\xbf\x7b\xa6\x07\x5e\xd6\x1a\x95\xf9\xa7\xc6\xba\xbe\xfa\x3b\x9a\xa1\x94\x1d\x2d\xed\x11\xe3\xba\x57\xb3\x3f\xeb\x34\x5c\x7e\x44\x29\xe3\x62\x7c\x3a\xbe\x03\x1c\x3f\x21\xa7\x0d\xd3\x29\x24\x55\x1c\x98\x55\x60\xda\xbd\x98\x96\x4a\x73\xd9\xa7\xe4\x12\xc8\x87\xca\x75\x79\x73\xf1\xdb\xe8\xca\x67\xa7\x7d\x9c\x76\xf8\xb0\xff\xf7\xb2\x09\xe4\x83\x65\xbb\xc7\xf5\x9a\xdc\xe1\x6b\xb7\xc9\x3f\xda\x15\x7e\x5e\xbc\xff\xbb\xc5\xdb\x2a\x97\xed\x8d\xfa\x2f\xfe\x71\x1a\x30\x7e\x7d\x35\xf1\x9f\xe9\xfa\x2a\xdb\xe1\x6e\x11\xb7\x90\x0a\x98\xbd\x5e\xe4\xb3\xec\xaf\x00\x00\x00\xff\xff\x95\x44\x62\xff\xd2\x0d\x00\x00")

func publicSwaggerSwaggerYamlBytes() ([]byte, error) {
	return bindataRead(
		_publicSwaggerSwaggerYaml,
		"public/swagger/swagger.yaml",
	)
}

func publicSwaggerSwaggerYaml() (*asset, error) {
	bytes, err := publicSwaggerSwaggerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "public/swagger/swagger.yaml", size: 3538, mode: os.FileMode(420), modTime: time.Unix(1542407286, 0)}
	a := &asset{bytes: bytes, info: info}
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
	if err != nil {
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
	"public/swagger/swagger.json": publicSwaggerSwaggerJson,
	"public/swagger/swagger.yaml": publicSwaggerSwaggerYaml,
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
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"public": &bintree{nil, map[string]*bintree{
		"swagger": &bintree{nil, map[string]*bintree{
			"swagger.json": &bintree{publicSwaggerSwaggerJson, map[string]*bintree{}},
			"swagger.yaml": &bintree{publicSwaggerSwaggerYaml, map[string]*bintree{}},
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

