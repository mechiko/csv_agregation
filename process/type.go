package process

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mechiko/utility"
)

type process struct {
	NameFileWithoutExt string
	File               string
	Koroba             map[string][]*utility.CisInfo
	Palet              map[string]map[string]string
	KM                 map[string]string
	arrKM              []string
	ListKoroba         [][]string
	ListPalet          [][]string
	KorobaKeys         []string
}

func New(file string) (*process, error) {
	if file == "" {
		return nil, fmt.Errorf("file name empty")
	}
	if !utility.PathOrFileExists(file) {
		return nil, fmt.Errorf("file not found")
	}
	name := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
	p := &process{
		File:               file,
		NameFileWithoutExt: name,
		Koroba:             map[string][]*utility.CisInfo{},
		Palet:              map[string]map[string]string{},
		KM:                 make(map[string]string),
		arrKM:              make([]string, 0),
	}
	return p, nil
}
