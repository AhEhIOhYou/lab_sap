package lab6

import (
	"bufio"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Task3 - Построить список имен файлов, название которых начинается на заданную букву.
func Task3(path, letter string) []string {
	var files []string

	err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			name := d.Name()
			if strings.HasPrefix(name, letter) {
				files = append(files, name)
			}
		}
		return nil
	})
	if err != nil {
		log.Println(err)
		return nil
	}
	return files
}

// Task5 - Построить список имен файлов, в которых имеется заданное слово.
func Task5(dir string, word string) []string {
	var files []string
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			name := d.Name()
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			defer f.Close()
			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				line := scanner.Text()
				if strings.Contains(line, word) {
					files = append(files, name)
					break
				}
			}
			if err := scanner.Err(); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Println(err)
		return nil
	}
	return files
}
