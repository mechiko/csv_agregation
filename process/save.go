package process

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func (p *process) Save(out string) error {
	fileNameKM := "Utility_" + p.NameFileWithoutExt + ".csv"
	fileNameKM = filepath.Join(out, fileNameKM)

	if err := saveKM(fileNameKM, p.arrKM); err != nil {
		return fmt.Errorf("error write file utility %w", err)
	}
	return nil
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
