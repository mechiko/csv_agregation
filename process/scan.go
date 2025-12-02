package process

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/mechiko/utility"
)

func (p *process) Scan() error {
	ar, err := readStringArray(p.File)
	if err != nil {
		return fmt.Errorf("ошибка получения массива данных %w", err)
	}

	for i, row := range ar {
		record, err := NewRecord(row)
		if err != nil {
			return fmt.Errorf("ошибка record %w", err)
		}
		if p.Gtin == "" {
			p.Gtin = record.Cis.Gtin
		}
		if p.Gtin != record.Cis.Gtin {
			return fmt.Errorf("gtin записи %s отличается от gtin первой строки %s [строка %d файла %s]", record.Cis.Gtin, p.Gtin, i+1, p.File)
		}
		if _, ok := p.KM[record.Cis.Cis]; ok {
			return fmt.Errorf("ошибка дубликат %d [%s] record  %w", i, record.Cis.Cis, err)
		}
		p.KM[record.Cis.Cis] = record.Cis.Serial
		p.arrKM = append(p.arrKM, record.Cis.Code)
		if _, ok := p.Koroba[record.Korob]; !ok {
			p.Koroba[record.Korob] = make([]*utility.CisInfo, 0)
			p.KorobaKeys = append(p.KorobaKeys, record.Korob)
		}
		p.Koroba[record.Korob] = append(p.Koroba[record.Korob], record.Cis)
		if _, ok := p.Palet[record.Palet]; !ok {
			p.Palet[record.Palet] = make(map[string]string)
		}
		p.Palet[record.Palet][record.Korob] = record.Korob
	}
	p.ListKoroba = make([][]string, 0)
	keysKorob := make([]string, 0, len(p.Koroba))
	for k := range p.Koroba {
		keysKorob = append(keysKorob, k)
	}
	slices.Sort(keysKorob)
	for _, key := range keysKorob {
		for _, cis := range p.Koroba[key] {
			r := []string{key, cis.Cis}
			p.ListKoroba = append(p.ListKoroba, r)
		}
	}
	p.ListPalet = make([][]string, 0)
	keysPalet := make([]string, 0, len(p.Palet))
	for k := range p.Palet {
		keysPalet = append(keysPalet, k)
	}
	slices.Sort(keysPalet)
	for _, key := range keysPalet {
		keys := make([]string, 0, len(p.Palet[key]))
		for k := range p.Palet[key] {
			keys = append(keys, k)
		}
		slices.Sort(keys)
		for _, kk := range keys {
			r := []string{key, kk}
			p.ListPalet = append(p.ListPalet, r)
		}
	}
	return nil
}

func readStringArray(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия файла %w", err)
	}
	defer f.Close()

	arr := make([][]string, 0)
	scanner := bufio.NewScanner(f)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	index := 0
	for scanner.Scan() {
		readTxt := scanner.Text()
		if readTxt != "" {
			txt := strings.Split(readTxt, "\t")
			if len(txt) != 3 {
				return nil, fmt.Errorf("полей в каждой строке [%d] файла должно быть три %s", index, filepath.Base(filePath))
			}
			arr = append(arr, txt)
		}
		index++
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("ошибка сканера %w", err)
	}
	return arr, nil
}
