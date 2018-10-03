// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/example.html (4.903kB)
// assets/expected.md (1.85kB)

package quip2md

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _assetsExampleHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x58\x51\x4f\x1b\x3d\xd6\xbe\xef\xaf\x38\x5f\xbe\x8b\x09\x52\x9b\x81\xf0\xb6\x85\x6a\x70\x15\xd8\xd2\xf0\x16\x0a\x2d\x04\xe8\x7b\x53\x79\x66\x4e\x62\x77\x3c\xf6\xd4\xf6\x24\x64\x57\x2b\xf5\x7e\xaf\xf7\x6e\x57\xe2\xb7\xf0\x53\xfa\x4b\x56\x1e\x27\x24\x24\x65\x02\x68\x5f\xad\x40\xa3\xe0\xe7\xf1\x73\x8e\xcf\xb1\x1f\xc6\x89\xd8\x06\xf0\x74\x27\x88\x13\xda\xd9\xeb\x5c\x1e\xe9\x2f\x67\x01\xf9\x54\xf2\x02\xac\x82\x23\xaa\xb3\x54\x8d\xe4\x1b\xf8\x8c\x7d\xd4\x28\x13\x84\xa3\xf3\x93\x28\x64\x1b\xe4\xd9\xb3\xa8\x98\x9b\xb9\x71\x45\x69\x1a\x40\x22\xa8\x31\x3b\x81\xe0\x12\x03\xf2\xf3\xc7\x3f\xa2\xb0\x58\x62\x0e\x8e\xb3\xf1\x02\xf3\x8c\x71\x03\xa9\x4a\xca\x1c\xa5\x05\x63\xa9\x4c\x0d\x50\x03\x96\x21\xe8\xdb\xd0\x89\x92\x43\xd4\x86\x2b\x09\x7d\xa5\xc1\x65\xf9\xc2\xaa\x17\xd3\x2c\x5b\xb0\xaf\x74\x4e\xad\xe5\x72\x00\x54\xa6\x60\xec\x58\x20\x14\x54\xd3\x1c\x2d\x6a\x03\x23\x6e\x19\x97\xc0\x50\x23\x50\x8d\x80\x57\x05\x26\x16\x53\xb7\xd6\x18\xc1\x94\x45\xa1\xb4\xc5\xb4\x05\xa7\x4c\x95\x22\x75\xf1\xc7\x20\x15\x08\x25\x07\xa8\x21\xc6\xe7\x50\x4a\x6e\xc1\xa2\xb1\x06\x8c\x27\xf5\x29\x17\xad\x5f\x2c\xf4\xe0\xb7\xb6\x18\x3e\xa8\x24\x32\xdf\x65\x83\x05\x66\x27\x4d\xdd\x3a\xa2\x98\xe4\x4a\x63\x14\xc6\xc4\xad\xdf\xba\xfa\x70\x09\x14\xa2\x92\x18\x74\x6b\xb3\x18\x85\x25\xa9\x96\x39\xd0\xb4\x60\x6e\x31\x2e\x3f\x70\xf3\xc0\x32\x2a\x21\xe2\xc4\xf0\xbc\x10\x18\x85\x7c\x26\x13\xa3\x1d\x21\xba\x72\x50\x17\xca\x80\xd2\xa0\x2c\x43\x3d\x29\x5c\x94\xa2\x20\xae\x60\x03\x13\x85\xee\xf3\xcd\xb5\x8a\xbf\x61\x62\xcd\x74\xb5\x77\xf6\x4e\xef\xb0\xff\xa1\x1d\x90\xee\xc6\x74\x7b\xb0\xf6\x1c\xda\x11\x3c\xfb\x12\x90\x6e\x3b\x0a\x59\xbb\x42\x37\xe7\xd0\x23\x4a\xbf\xcb\x80\x74\x37\xa3\x90\x6d\x2e\xa1\xbb\xa2\x73\x9e\x05\x64\xb7\x14\x02\x5d\xb7\x0e\xb9\xb1\x53\x62\xca\x87\x90\x52\x4b\x5f\x18\x4c\x2c\x57\xf2\x45\x95\xfb\x4e\xf0\x32\x20\x51\x29\xe6\xd3\xdb\xcb\x3b\x57\x01\x89\x04\x9f\x1b\x64\xef\x2e\xd7\x67\x7b\x31\x80\x21\x15\x25\xee\x04\x1b\x01\x89\x4c\x41\xe5\x12\x93\x44\x31\xd9\x55\x22\x75\xed\x88\x42\x47\x71\x39\xc4\x3a\x24\x51\x28\xf8\x82\x78\x67\xef\xec\x8b\x9e\x89\x2f\x49\x7a\x9c\x44\x9c\x1c\x58\x2a\x78\x62\x5c\x77\x56\xaa\x16\xf6\x90\x7e\xab\x51\xf5\x38\x89\x4a\xd2\x93\x29\x6a\xb7\x99\xa0\xd9\x29\x0a\xa4\xda\x00\xaf\x02\x41\x8c\x09\x2d\x0d\x42\xae\x8c\xbd\x3d\xe4\xd0\x17\x74\xa8\xb4\x3b\x87\x32\xb0\xa0\xd1\x4d\x87\x72\x2a\xb2\xe6\x76\xd9\xca\xec\x4e\xdf\x77\xf0\xa2\x26\x3b\x8f\x93\x6a\x6b\x9d\x5a\xcd\x33\xb4\x4c\xab\x72\xc0\xfc\x0e\x5b\xa9\xff\x8d\x89\xf8\xa4\x46\xdf\xe3\x24\x4a\x54\x8a\xe4\x40\x56\xab\xdf\x53\x29\x46\x61\x35\xb2\x52\x1f\x6d\x0f\x55\x8d\xbe\xc7\x49\x44\x81\x69\xec\xef\x34\x98\xb5\x85\x79\x13\x86\x03\x6e\x59\x19\xb7\x12\x95\x87\x89\xba\x12\x38\x6e\x90\xee\xb8\xa8\x2a\x97\x45\x21\x7d\x40\x60\xb5\x95\x8a\xba\xc0\x15\x4e\x22\x2a\xa5\xb2\xd4\xed\xf4\x09\xb7\x91\x6c\x03\x6b\x37\x1c\xb5\xe1\xa9\x67\xdb\xdb\xdd\x2f\x0d\xd2\xe5\x03\x26\xf8\x80\x59\x4c\xa3\x70\x36\x6d\x75\x2a\xe9\x3a\x0e\xf7\x6b\x52\xf1\x38\x39\xb0\x98\x57\x12\xce\x55\xab\x0f\xae\xda\xb1\x46\x9a\x99\x95\x31\x7a\xa2\xf7\x6a\xe6\x78\x05\xd5\x28\xed\x72\x24\xcf\x22\x27\x9a\xe7\x54\x8f\xc1\x94\x95\xfd\xfc\x4a\xbc\x14\x0b\x01\x8e\x7b\xeb\xe9\xf9\xca\x00\x9e\x45\x4e\xcb\x18\xde\x09\x74\xff\x7c\x1e\x24\x3e\xd0\x1f\xff\xf2\xb2\xa6\x42\x1e\x27\x1f\xd1\x38\xbb\x5a\xa1\x1e\x3a\xf9\xb9\x47\xca\x87\xe4\xae\xff\xfd\x31\xb4\x9f\x2e\x03\xf2\xb1\xcc\x63\xd4\x0f\xf4\xbf\x57\x8b\xfe\x77\x5e\xbc\xda\xec\x2d\xfa\x9f\xe6\xef\x7f\x7b\x98\xff\x79\xe6\xa3\xfc\x4f\x6d\x96\x27\x75\x55\xf2\xf8\x63\xfd\xef\xa4\xad\x37\xdf\xd7\xa8\x7a\xfc\x7f\xe5\x7f\xfb\x97\x5b\x87\x87\x35\xd9\x79\xfc\xe9\xfe\x47\xcb\xec\xf8\x75\x8d\xbe\xc7\x9f\xee\x7f\x22\xbd\xe8\xb6\x6b\xf4\x3d\xfe\x27\xf8\xdf\xe1\xee\x91\xec\xd4\x04\xf6\xf8\x5d\x53\x5b\x21\x39\xdc\x3b\x3d\xed\xad\xb4\x00\xcf\x7a\x92\xc7\xd8\xcd\x63\xfc\xeb\xca\x00\x9e\xf5\x68\x8f\x19\xad\x1f\x1f\x14\x35\x05\xf1\xf8\x7f\xcd\x63\x44\x22\xcd\x71\x40\xf6\x18\x26\xd9\x83\x0c\xe6\xf5\xa2\xc1\x9c\x1d\xd0\xdf\x3f\x2d\x1a\xcc\x70\xac\x69\xf7\x41\x06\xe3\x99\x8f\x32\x18\xc6\x7a\xb8\x7b\x2b\x9e\xb8\xd4\x31\xfd\xc5\xab\x5b\x45\x7b\xac\xcf\x74\x78\xf1\xa9\xbb\x5a\xdc\xd3\xee\xd8\xcd\x83\x8c\x62\x6b\xd4\x66\x7f\xd4\xb4\xd7\xe3\x4f\x37\x0a\x7b\xb1\xff\xee\xac\x46\xdf\xe3\x4f\x37\x8a\xbc\xd8\x3f\xc7\x1a\x7d\x8f\xff\x09\x46\x91\x64\x5d\x3c\xa8\x09\xec\xf1\x55\x46\x71\xdf\x31\xf8\xfc\x21\x1e\x5f\xba\x0b\xb0\xb2\x68\xa6\x47\x60\xfe\xaa\x36\xbe\xe8\x6d\xdf\x7f\xa9\x8b\x85\x4a\xb2\xef\x6e\xf2\xfc\x3b\xe9\x65\xf6\xee\xa5\x63\xfd\xab\xab\x46\x40\x63\x55\x5a\xb0\x8c\x9b\xb7\x01\x9c\xf2\x3c\x57\x12\xa8\xc9\x30\x85\x1c\x5b\xd0\xb8\x60\x3c\x61\xc0\x0d\x8c\x94\x36\xf8\x1c\x8c\x45\x2a\xaa\x0b\x2d\x14\x1c\xdd\x15\x2d\xe3\xa2\x1a\xe8\xe4\xb1\x56\x06\xdf\x36\xaa\x65\x1d\xc0\x80\x0e\x11\xb8\x05\x0a\xb9\x72\x3e\x10\x18\x60\x54\xbb\x3b\xac\xdb\x36\xb6\x05\x8d\x0e\xe4\x48\xad\xd3\x79\xee\x84\x28\xf4\x75\xc9\xab\xbf\xdf\xfe\xfc\xf1\x6f\xa8\x74\x7e\xfe\xf8\xa7\xbb\x73\x9e\x50\xab\x79\x92\xc1\x67\x65\x59\xbf\x34\xe6\x79\x14\x13\xf0\x87\x32\x5e\x6e\xea\x68\x34\x6a\x0d\x94\x4a\x35\xd2\xd4\x54\xbd\x1d\x29\x9d\x85\x55\x21\x4c\xd8\x7e\xb9\xde\xde\xda\x6a\x37\xc8\x19\x43\xb8\xe0\x06\xe1\x88\xca\xc0\xc0\x3e\x52\xed\x5b\x3e\x2b\xdb\x62\xb9\x5f\x63\x32\xbc\xb8\xb7\xdc\x77\x5a\xb7\xfd\xea\xf7\xdd\xc3\x80\xf8\x4d\x3c\x69\x9c\x9e\xef\x83\xed\x7c\x10\x73\xaf\x84\x1a\xad\x1d\x17\x9a\x3b\xcf\xfe\xff\xff\x0b\x4b\xa3\xc3\x98\xcb\x10\xe5\x10\x8a\xb1\x65\x4a\x6e\xfa\xfd\xe2\x1e\x3c\x2f\x94\xb6\x40\xcd\x58\x26\x5c\xcd\xc6\xab\x47\x35\x0a\x29\xf6\xa1\x40\xad\x52\x9e\x34\xf9\x1b\xe0\xd2\xae\xbd\xa9\xe0\x9b\x6b\xb8\xb9\x86\x11\xe3\x02\x6f\xae\xcf\x74\x89\xf3\xc3\xfe\xb7\xca\xa2\xd9\xe8\x38\x01\xae\x52\x60\xd4\x40\x41\x8d\xc1\xb4\xb1\xb6\x44\xa6\x23\xca\x6f\x53\x69\x19\x81\x58\x34\xf9\x94\x76\xf7\xe7\xfe\xc1\xe9\xf4\x44\x23\xb5\xf8\xd5\x52\x93\x35\x7d\x70\x9e\x34\x37\xd6\xd6\x56\x93\xda\x13\x92\x50\xaa\x80\x9d\x9b\xeb\x29\x79\x80\xf6\x2b\x0e\x51\xda\xaf\x0e\x69\xce\x48\x2d\x5d\xca\xaf\x7d\xa5\x71\x88\x7a\x32\x3c\x39\x90\x85\xc6\xe5\x6e\x8d\xaf\xb6\x8f\xee\xe9\x56\x41\x93\x8c\x0e\x10\x72\xca\xe5\x52\x8f\x1a\xfd\xdc\x36\x66\xa3\xfd\x52\x26\x15\xb1\xb9\x06\x7f\xbb\xad\x07\xdc\x5c\xf7\x73\xdb\x3a\x71\x7a\x42\x36\x1b\xdf\x27\x5f\x78\xe5\x93\x77\xc1\x49\xd9\xff\x3e\xcd\xed\xd9\x7f\x02\x00\x00\xff\xff\x38\xcf\x40\x0a\x27\x13\x00\x00")

func assetsExampleHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsExampleHtml,
		"assets/example.html",
	)
}

func assetsExampleHtml() (*asset, error) {
	bytes, err := assetsExampleHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/example.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x9, 0x6f, 0x90, 0xfe, 0xe9, 0xc4, 0x7d, 0x89, 0x17, 0xd1, 0xe7, 0x71, 0x54, 0xb6, 0x55, 0x7c, 0x33, 0x28, 0x73, 0x60, 0xf2, 0x23, 0xb9, 0xe6, 0x68, 0xd6, 0xe2, 0xfd, 0xf2, 0x79, 0x24, 0x20}}
	return a, nil
}

var _assetsExpectedMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x95\xdd\x6e\xdb\x36\x14\xc7\xef\xf9\x14\xff\xd9\x17\xb1\x84\x44\x5e\x9c\x75\x0b\x02\xac\x41\x5a\xac\x70\x80\xa5\x48\x9b\x6c\x1d\x10\x04\x0e\x25\x1d\x8b\x9c\x25\x52\x25\x8f\xac\x1a\xc3\x8c\xbe\xc4\xee\x36\xc0\xcf\xe2\x47\xe9\x93\x0c\x94\x9d\xaf\x35\x1b\x7a\x33\xdb\x80\xc8\xc3\xf3\xf5\xff\xd1\x3e\xee\xe3\x4d\xa3\x6b\xb0\xc5\x99\x74\xb3\xdc\xb6\xe6\x08\x6f\x69\x4a\x8e\x4c\x46\x38\xfb\xf9\x5c\x88\x4b\xa5\x3d\x72\x9b\x35\x15\x19\x86\x67\x69\x72\x0f\xe9\xc1\x8a\xe0\xee\x5c\x33\x6b\xe6\xe4\xbc\xb6\x06\x53\xeb\xba\xac\x7b\x6c\xf7\x6e\xb3\x26\x78\x65\x5d\x25\x99\xb5\x29\x20\x4d\x0e\xcf\x8b\x92\x50\x4b\x27\x2b\x62\x72\x1e\xad\x66\xa5\x0d\x14\x39\x82\x74\x04\xfa\x50\x53\xc6\x94\x87\xde\x52\x82\x6f\xea\xda\x3a\xa6\x3c\xc1\x85\xb2\x4d\x99\x87\xfa\x0b\x18\x8b\xd2\x9a\x82\x1c\x52\xda\x45\x63\x34\x83\xc9\xb3\x87\xdf\x38\x4d\xa5\x2e\x13\x21\x4e\xf2\x3c\x14\x8e\xe3\xca\x3a\x8a\xe3\xd0\x2d\x07\x35\xda\x40\x62\xe2\x29\xf4\xc1\x34\xe9\xfa\x29\x9c\xac\x55\xa8\x1a\x12\x21\x04\x80\x95\x34\x88\xbd\xae\xea\x92\xee\x83\x53\xe2\x96\x28\xb4\x2c\x43\x76\x0f\xeb\x60\x59\x91\xdb\x8a\x5b\x2e\x83\xa2\xc2\x2f\x97\xeb\x95\x4d\x7f\xa5\x8c\x7d\x22\x44\x1f\xe3\x7d\x21\xfa\x7d\x8c\x47\xe1\xd1\xc7\xf8\x60\xf3\x7c\xd1\x94\x25\x05\xc1\x3f\x6a\xcf\x42\xc4\x88\xe3\x17\xb6\xcc\xe3\x38\x2c\x4f\x59\x96\x3a\xf3\x61\x3d\xf9\xc9\xe4\xe4\x4a\x6d\x08\x83\x93\xba\x26\xe9\x3c\x74\x77\x8c\x94\x32\xd9\x78\x42\x65\x3d\xdf\x5d\x28\xa6\xa5\x9c\x5b\x17\xee\xd0\xec\x30\x1c\x85\x70\x34\xb7\x49\xa2\x89\x88\xb1\x5c\x5e\xb0\xd3\x33\x62\xe5\x6c\x53\xa8\xe5\x52\xc4\xb8\x39\x35\x5d\x91\x97\x36\xa7\x1b\x11\xe3\x6a\xbc\xa8\xbb\x90\xd9\xf5\x40\x31\xd7\xfe\x68\x38\x2c\x34\xab\x26\x4d\x32\x5b\x0d\x33\xfb\xa1\xa4\x45\x24\x62\x8c\x75\xa1\x4a\x5d\x28\xa6\x5c\xc4\x38\x65\xaa\x44\xb8\x5c\x11\xb2\xa5\x8e\xe4\xcc\x8b\x18\xe7\x4e\x57\xd2\x2d\xe0\x9b\x8e\x8c\x00\x62\x5c\x34\x29\x7e\x28\x29\x7c\xcf\x04\xba\x57\x8c\xd7\xe4\x03\x93\x87\x47\x1d\xad\xd7\x4d\x95\x92\xbb\xa3\xb5\x9f\xdc\xe3\x1a\x25\x0f\x78\x1d\x24\xff\x03\xb0\x6f\x92\x27\x88\x3d\x4b\xfe\x81\xec\xdb\xe4\x4b\x99\x7d\x97\x3c\x82\x76\x98\x3c\x81\x47\x26\x4f\xf0\xd1\xc9\xbf\x02\x7a\xa9\x28\x9b\x6d\xe9\x5c\xe1\xfa\x1e\xcf\xd5\x2f\xd7\x0f\xf8\x84\xdd\x3d\xa0\x49\xe7\xfa\xb9\xb6\x60\x7d\x2c\x2e\x58\xbe\x50\x5d\x70\x7d\x28\x6f\xd3\xdf\x9b\xc6\x32\x79\x21\x9e\xe3\xd3\xc7\x3f\xc7\xb6\x85\x4c\x6d\xc3\x60\xa5\xfd\xf1\x0e\x2e\x74\x55\x59\x03\xe9\x67\x94\xa3\xa2\x04\xbd\x77\x4a\x67\x0a\xda\xa3\xb5\xce\xd3\x2e\x3c\x93\x2c\xbb\x69\x82\x5a\x53\xf8\xed\xcd\x74\xd9\x19\x4e\xaa\xd4\x59\x4f\xc7\x3d\xf1\x1c\xa7\x28\xe4\x9c\xa0\x19\x12\x95\x0d\x74\x76\x3c\x94\x74\x61\x7c\x04\x6d\x9c\xa0\x77\x82\x8a\x24\x87\x2c\xbb\x21\x8d\xc4\xd4\x35\xba\xdb\x1f\x7f\xfa\xf8\x17\xba\x0e\xff\x40\x1c\x9f\x4b\x76\x3a\x9b\xe1\xad\x65\x35\x6d\xbc\xdf\x8d\x63\x5c\x5d\x2a\xc2\x3b\xed\x09\x67\xd2\xec\x78\xbc\x22\xe9\xee\x61\xb4\x6d\x9b\x14\xd6\xe6\x8e\x64\xee\x3b\x26\xad\x75\xb3\xe1\xfb\x4e\xfa\x70\xf4\xec\xeb\xd1\xe1\xe1\x28\xda\xde\x97\xcd\x49\x08\x71\x73\x73\x23\xfa\x5f\x0d\x1b\xef\x86\xa9\x36\x43\x32\x73\xd4\x0b\x56\xd6\x1c\x08\xa1\xab\x30\x02\x21\xfd\xc2\x64\xda\x0a\xd1\x2d\x90\xd3\x14\x35\x39\x9b\xeb\x6c\xa0\x8f\xa0\x0d\x47\x47\x42\xac\x57\x58\xaf\xd0\x2a\x5d\xd2\x7a\x75\xe9\x1a\x3a\xda\x9a\x36\x9f\xda\x69\xc3\x83\xde\x49\x88\xd4\x36\x87\x92\x1e\xb5\xf4\x9e\xf2\x5e\x24\x1e\x79\xca\x56\xea\xbb\x9a\x89\x2f\x89\xea\x81\x8e\xc4\x7a\xf5\xf8\x2d\x3e\xb7\xdc\xc6\x64\x8e\xc2\x70\x65\xe9\x67\x83\x4d\x39\x9d\x0d\xf6\xa3\x48\xfc\xb7\xc7\x28\x8a\x44\x69\x6d\x8d\xef\xd7\xab\x5b\xc7\x82\x78\x42\x73\x32\x3c\x09\x27\x83\x48\x74\x1e\x89\x6b\xcc\x64\x6a\x1d\xcd\xc9\x05\x5b\x60\x28\xf6\xf6\xf6\x36\xab\x5a\x66\x33\x59\x10\x2a\xa9\xcd\x1d\xc3\xde\xb4\xe2\x9e\x10\xd3\xc6\x64\xdd\xc1\x20\xc2\x6f\x9d\x02\xac\x57\xd3\x8a\x93\xf3\xc0\xa7\x34\x83\xde\xfb\xed\xdf\x63\xb5\x1d\x0e\x01\xcf\xef\x21\xef\xdf\x01\x00\x00\xff\xff\xc8\x1a\x53\x3c\x3a\x07\x00\x00")

func assetsExpectedMdBytes() ([]byte, error) {
	return bindataRead(
		_assetsExpectedMd,
		"assets/expected.md",
	)
}

func assetsExpectedMd() (*asset, error) {
	bytes, err := assetsExpectedMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/expected.md", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x68, 0x9d, 0x61, 0x5f, 0x68, 0x3b, 0x17, 0xbc, 0xeb, 0x16, 0xd6, 0x5a, 0xda, 0x34, 0xe5, 0x5b, 0xd4, 0xa4, 0x49, 0x2f, 0x1a, 0xf2, 0xe0, 0xe8, 0x12, 0x27, 0x9, 0x63, 0x60, 0xfe, 0xb2, 0x81}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"assets/example.html": assetsExampleHtml,

	"assets/expected.md": assetsExpectedMd,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"assets": &bintree{nil, map[string]*bintree{
		"example.html": &bintree{assetsExampleHtml, map[string]*bintree{}},
		"expected.md":  &bintree{assetsExpectedMd, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}