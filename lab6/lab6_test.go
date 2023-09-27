package lab6_test

import (
	"lab6"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestTask3(t *testing.T) {
	dir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	subdir, err := os.MkdirTemp(dir, "subtest")
	if err != nil {
		t.Fatal(err)
	}

	files := []string{"A.txt", "B.txt", "C.txt", "AA.txt", "AB.txt", "AC.txt"}
	subfiles := []string{"D.txt", "E.txt", "F.txt", "AD.txt", "AE.txt", "AF.txt"}
	for _, file := range files {
		f, err := os.Create(filepath.Join(dir, file))
		if err != nil {
			t.Fatal(err)
		}
		f.Close()
	}
	for _, file := range subfiles {
		f, err := os.Create(filepath.Join(subdir, file))
		if err != nil {
			t.Fatal(err)
		}
		f.Close()
	}

	tests := []struct {
		name   string
		path   string
		letter string
		want   []string
	}{
		{
			name:   "starts with A",
			path:   dir,
			letter: "A",
			want:   []string{"A.txt", "AA.txt", "AB.txt", "AC.txt", "AD.txt", "AE.txt", "AF.txt"}},
		{
			name:   "one with B",
			path:   dir,
			letter: "B",
			want:   []string{"B.txt"},
		},
		{
			name:   "one in sub with D",
			path:   dir,
			letter: "D",
			want:   []string{"D.txt"}},
	}

	for _, tc := range tests {
		t.Run(tc.path+"_"+tc.letter, func(t *testing.T) {
			got := lab6.Task3(tc.path, tc.letter)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Task3(%q, %q) = %v, want %v", tc.path, tc.letter, got, tc.want)
			}
		})
	}
}

func TestTask5(t *testing.T) {
	dir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	subdir, err := os.MkdirTemp(dir, "subtest")
	if err != nil {
		t.Fatal(err)
	}

	subsubdir, err := os.MkdirTemp(subdir, "subsubtest")
	if err != nil {
		t.Fatal(err)
	}

	files := []string{"A.txt", "B.txt", "C.txt"}
	subfiles := []string{"D.txt", "E.txt", "F.txt"}
	subsubfiles := []string{"G.txt", "H.txt", "I.txt"}
	contents := []string{"hi hii world", "later world", "ohayo onichan"}
	for i, file := range files {
		f, err := os.Create(filepath.Join(dir, file))
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		f.WriteString(contents[i])
	}
	for i, file := range subfiles {
		f, err := os.Create(filepath.Join(subdir, file))
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		f.WriteString(contents[i])
	}
	for i, file := range subsubfiles {
		f, err := os.Create(filepath.Join(subsubdir, file))
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		f.WriteString(contents[i])
	}

	tests := []struct {
		name string
		path string
		word string
		want []string
	}{
		{
			name: "hii all files",
			path: dir,
			word: "hii",
			want: []string{"A.txt", "D.txt", "G.txt"},
		},
		{
			name: "ohayo all files",
			path: dir,
			word: "ohayo",
			want: []string{"C.txt", "F.txt", "I.txt"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.path+"_"+tc.word, func(t *testing.T) {
			got := lab6.Task5(tc.path, tc.word)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Task5(%q, %q) = %v, want %v", tc.path, tc.word, got, tc.want)
			}
		})
	}
}

func TestTask1(t *testing.T) {
	dir := "test_dir"
	expectedFile := "test_dir/file2.txt"
	expectedSize := int64(500)

	err := os.Mkdir(dir, 0755)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	file1, err := os.Create(dir + "/file1.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer file1.Close()

	file2, err := os.Create(expectedFile)
	if err != nil {
		t.Fatal(err)
	}
	defer file2.Close()

	file3, err := os.Create(dir + "/file3.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer file3.Close()

	file1.Write(make([]byte, 100))
	file2.Write(make([]byte, 500))
	file3.Write(make([]byte, 2))

	maxFile, maxSize := lab6.Task1(dir)

	if maxFile != expectedFile {
		t.Errorf("Expected maxFile to be %s, but got %s", expectedFile, maxFile)
	}

	if maxSize != expectedSize {
		t.Errorf("Expected maxSize to be %d, but got %d", expectedSize, maxSize)
	}
}
