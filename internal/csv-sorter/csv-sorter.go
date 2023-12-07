package csvsorter

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ExecuteSortCSV(inputFile, outputFile, sortField, order string) error {
	records, err := readCSVFile(inputFile)
	if err != nil {
		return err
	}

	header, data, err := separateHeaderAndData(records)
	if err != nil {
		return err
	}

	sortIndex, err := findSortIndex(header, sortField)
	if err != nil {
		return err
	}

	compare := createCompareFunction(data, sortIndex, order)
	sort.SliceStable(data, compare)

	err = writeCSVFile(outputFile, header, data)
	if err != nil {
		return err
	}

	return nil
}

func readCSVFile(inputFile string) ([][]string, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func separateHeaderAndData(records [][]string) ([]string, [][]string, error) {
	if len(records) < 2 {
		return nil, nil, fmt.Errorf("Arquivo CSV não contém dados")
	}

	return records[0], records[1:], nil
}

func findSortIndex(header []string, sortField string) (int, error) {
	for i, col := range header {
		if col == sortField {
			return i, nil
		}
	}
	return -1, fmt.Errorf("Campo de ordenação '%s' não encontrado no cabeçalho", sortField)
}

func createCompareFunction(data [][]string, sortIndex int, order string) func(i, j int) bool {
	return func(i, j int) bool {
		valI, errI := strconv.Atoi(data[i][sortIndex])
		valJ, errJ := strconv.Atoi(data[j][sortIndex])

		if errI == nil && errJ == nil {
			if order == "asc" {
				return valI < valJ
			} else {
				return valI > valJ
			}
		}

		if order == "asc" {
			return strings.ToLower(data[i][sortIndex]) < strings.ToLower(data[j][sortIndex])
		} else {
			return strings.ToLower(data[i][sortIndex]) > strings.ToLower(data[j][sortIndex])
		}
	}
}

func writeCSVFile(outputFile string, header []string, data [][]string) error {
	outputFilePtr, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outputFilePtr.Close()

	writer := csv.NewWriter(outputFilePtr)
	defer writer.Flush()

	if err := writer.Write(header); err != nil {
		return err
	}

	for _, record := range data {
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}
