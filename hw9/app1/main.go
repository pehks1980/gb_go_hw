package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var (
	nRows  = flag.Int64("Number of rows", 0, "Number of rows to read from file(not given = all rows)")
	flCols = flag.String("columns", "", "columns list, separated by comma(not given = all columns)")
)

// struct member must start with capital otherwise marshal json doesnt work
type Row struct {
	Cols map[string]string `json:"cols"`
}

type JS struct {
	Fname string `json:"fname"`
	Rows  []Row  `json:"rows"`
}

func main() {
	flag.Parse()

	path, err := os.Getwd()

	if err != nil {
		log.Println(err)
	}

	filename := path + "/hw9/app1/" + "test.csv"

	//filename := strings.TrimSpace(flag.Arg(0)) // Считываем имя файла и очищаем ввод от пробелов

	// Пробуем открыть файл
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Не могу открыть файл: %v", err)
	}

	// Не забываем закрыть файл при выходе из функции
	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("Не могу закрыть файл: %v", err)
		}
	}()

	reader := csv.NewReader(file)
	// read cols names row
	fileCols, err := reader.Read()
	if err != nil {
		log.Fatalf("Не могу считать заголовочную строку csv-файла: %v", err)
	}

	fCols := strings.Split(*flCols, ",")

	colsMask := make(map[string]int)
	colsIdx := make(map[string]int)

	for i, col := range fileCols {
		//make key in map Mask
		colsMask[col] = 0
		colsIdx[col] = i
		// if no cols in flag mask all and continue
		if len(*flCols) == 0 {
			colsMask[col] = 1
			continue
		}

		for _, fl := range fCols {
			if col == fl {
				// if match make this key val 1
				colsMask[col] = 1
				break
			}
		}
	}
	// at the end we get map colsMask of keys columns, and matched cols will have 1 s values
	// colsIdx has key of columns and val index in row

	//prepare json struct
	var js JS

	js.Fname = filename
	//prepare json row array
	var rows []Row

	i := int64(1)
	// process file csv
	for {
		row, err := reader.Read()

		// Если в качестве ошибки получили EOF (End of file), значит
		// файл закончился раньше, чем мы считали максимально заданное кол-во строк.
		// прекращаем чтение
		if err == io.EOF {
			break
		}

		// Если произошла какая-то другая ошибка, прекращаем работу программы
		if err != nil {
			log.Fatalf("Ошибка при чтении файла: %v", err)
		}

		// Выводим значения строки в ожидаемом формате:
		fmt.Printf("---------- Строка %d ---------\n", i)
		//make map for JS cols
		jsCols := make(map[string]string)

		for col, mask := range colsMask {
			if mask == 1 {
				fmt.Printf("%s:\t%s\n", col, row[colsIdx[col]])
				//add key val to jsCols
				jsCols[col] = row[colsIdx[col]]

			}
		}
		// insert jsCols to Row struct
		jsRow := Row{Cols: jsCols}
		// build array of these Row structs
		rows = append(rows, jsRow)

		fmt.Printf("-----------------------------\n\n")

		// if we reached nRows
		if i == *nRows {
			break
		}
		i++
	}

	// add JS appended row to js struct
	js.Rows = rows

	fmt.Println(js)
	//dump json to text
	jsonTofile, _ := json.Marshal(&js)

	fmt.Println(string(jsonTofile))

	filenameO := path + "/hw9/app1/" + "test.json"
	//dump JS to file
	fileO, err := os.Create(filenameO)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = fileO.Close()
		// Обратите внимание: при закрытии файла также возможна ошибка
		if err != nil {
			log.Fatal(err)
		}
	}()

	_, err = fileO.Write(jsonTofile)

	// Как обычно, обрабатываем ошибку
	if err != nil {
		_ = fmt.Errorf("не могу записать в файл: %w", err)
		return
	}

}
