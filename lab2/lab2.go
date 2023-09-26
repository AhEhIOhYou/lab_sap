package lab2

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

/*
3. Реализовать функцию, которая из указанного каталога строит ассоциативный массив,
в котором ключами являются имена текстовых файлов, а значениями количество строк файла.
*/

func CreateArrData(dir string) map[string]int {

	res := make(map[string]int)

	err := filepath.WalkDir(dir, func(path string, dirInfo fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if !dirInfo.IsDir() && filepath.Ext(path) == ".txt" {
			res[filepath.Base(path)] = linesInFile(path)
		}

		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

	return res
}

func linesInFile(fileName string) int {
	f, _ := os.Open(fileName)
	scanner := bufio.NewScanner(f)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	counter := 0

	for scanner.Scan() {
		counter++
	}

	return counter
}
