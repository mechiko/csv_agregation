package process

import (
	"fmt"

	"github.com/mechiko/utility"
)

type Record struct {
	Cis   *utility.CisInfo
	Korob string
	Palet string
}

func NewRecord(row []string) (*Record, error) {
	if len(row) != 3 {
		return nil, fmt.Errorf("записей не равно 3")
	}
	cis, err := utility.ParseCisInfo(row[0])
	if err != nil {
		return nil, fmt.Errorf("получение КМ %w", err)
	}

	r := &Record{
		Cis:   cis,
		Korob: row[1],
		Palet: row[2],
	}
	return r, nil
}
