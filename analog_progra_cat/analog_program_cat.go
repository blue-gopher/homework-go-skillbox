package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

// пример команды запуска: go run main.go file1.txt file2.txt result.txt
func main() {
	fmt.Println("Программа аналог cat.")

	sliceFiles := os.Args
	countFiles := len(sliceFiles) - 1

	switch countFiles {
	case 1:
		file, err := os.Open(sliceFiles[countFiles])
		defer file.Close()
		if err != nil {
			log.Fatalf("Не удалось открыть файл %s\n", sliceFiles[countFiles])
		}

		fileInfo, err := file.Stat()
		if err != nil {
			log.Fatalf("Не удалось прочитать данные файла %s\n", sliceFiles[countFiles])
		}

		lenData := fileInfo.Size()
		data := make([]byte, lenData)

		_, err = file.Read(data)
		if err != nil && err != io.EOF {
			log.Fatalln("Не удалось записать данные в буфер")
		}

		fmt.Println(string(data))

	case 2:
		bb := bytes.Buffer{}

		for i := 1; i < 3; i++ {
			file, err := os.Open(sliceFiles[i])
			defer file.Close()
			if err != nil {
				log.Fatalf("Не удалось открыть файл %s\n", sliceFiles[i])
			}

			fileInfo, err := file.Stat()
			if err != nil {
				log.Fatalf("Не удалось прочитать данные файла %s\n", sliceFiles[i])
			}

			lenData := fileInfo.Size()
			data := make([]byte, lenData)

			_, err = file.Read(data)
			if err != nil && err != io.EOF {
				log.Fatalln("Не удалось записать данные в буфер")
			}

			bb.Write(data)

			if i == 1 {
				transfer := []byte("\n")
				bb.Write(transfer)
			}
		}

		fmt.Println(bb.String())

	case 3:
		resultFile, err := os.OpenFile(sliceFiles[countFiles], os.O_APPEND|os.O_WRONLY, 0666)
		defer resultFile.Close()
		if err != nil {
			log.Fatalf("Не удалось открыть файл %s\n", sliceFiles[countFiles])
		}

		for i := 1; i < 3; i++ {
			file, err := os.Open(sliceFiles[i])
			defer file.Close()
			if err != nil {
				log.Fatalf("Не удалось открыть файл %s\n", sliceFiles[i])
			}

			fileInfo, err := file.Stat()
			if err != nil {
				log.Fatalf("Не удалось прочитать данные файла %s\n", sliceFiles[i])
			}

			lenData := fileInfo.Size()
			data := make([]byte, lenData)

			_, err = file.Read(data)
			if err != nil && err != io.EOF {
				log.Fatalln("Не удалось записать данные в буфер")
			}

			resultFile.Write(data)

			if i == 1 {
				transfer := []byte("\n")
				resultFile.Write(transfer)
			}
		}

		fmt.Println("Данные записаны в файл", sliceFiles[countFiles])

	default:
		fmt.Println("Вы указали неверное количество файлов")
	}
}
