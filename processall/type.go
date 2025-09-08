package processall

import (
	"github.com/mechiko/utility"
)

type process struct {
	NameFileWithoutExt string
	Koroba             map[string][]*utility.CisInfo
	Palet              map[string]map[string]string
	KM                 map[string]string
	arrKM              []string
	ListKoroba         [][]string
	ListPalet          [][]string
	KorobaKeys         []string
	Records            []*Record
}

func New() (*process, error) {
	p := &process{
		NameFileWithoutExt: "agregation",
		Koroba:             map[string][]*utility.CisInfo{},
		Palet:              map[string]map[string]string{},
		KM:                 make(map[string]string),
		arrKM:              make([]string, 0),
		Records:            make([]*Record, 0),
	}
	return p, nil
}
