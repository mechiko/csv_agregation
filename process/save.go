package process

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func (p *process) Save(out string) error {
	fileNameKM := "нанесение_" + p.NameFileWithoutExt + ".csv"
	index := 0
	for {
		cis := nextRecords(p.arrKM, index, 30000)
		if len(cis) == 0 {
			break
		}
		fn := fmt.Sprintf("%02d_%s", index+1, fileNameKM)
		fn = filepath.Join(out, fn)
		if err := saveKM(fn, cis); err != nil {
			return fmt.Errorf("error write file utility %w", err)
		}
		index++
	}

	return nil
}

// index from 0 startIndex 0
// nextRecords returns a batch of records starting from startIndex
// Returns empty slice when no more records are available
func nextRecords(arr []string, index int, count int) []string {
	startIndex := index * count
	if startIndex >= len(arr) {
		return []string{}
	}
	endIndex := startIndex + count
	// если последний индекс больше длины массива укорачиваем до размера массива
	if endIndex > len(arr) {
		endIndex = len(arr)
	}
	return arr[startIndex:endIndex]
}

// func saveTxt(name string, data [][]string) error {
// 	file, err := os.Create(name)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	for _, row := range data {
// 		_, err = fmt.Fprintf(file, "%s\n", strings.Join(row, "\t"))
// 		if err != nil {
// 			return fmt.Errorf("failed to write to file: %v", err)
// 		}
// 	}
// 	return nil
// }

func saveKM(name string, data []string) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, row := range data {
		_, err = fmt.Fprintf(file, "%s\n", row)
		if err != nil {
			return fmt.Errorf("failed to write to file: %v", err)
		}
	}
	return nil
}

func NewWriter(w io.Writer) (writer *csv.Writer) {
	writer = csv.NewWriter(w)
	writer.Comma = '\t'
	return
}

func saveCsvCustom(name string, data [][]string) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = '\t'
	writer.WriteAll(data) // calls Flush internally
	return writer.Error()
}
