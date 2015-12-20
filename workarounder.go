package workarounder

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var re = regexp.MustCompile(`\[workaround for #[0-9]+\]`)

func findMatch(input string) (results [][]string) {
	return re.FindAllStringSubmatch(input, -1)
}

func makeFileList(path string) (fileList []string, err error) {
	err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		// Skip hidden directories
		if info.IsDir() && filepath.HasPrefix(info.Name(), ".") {
			return filepath.SkipDir
		}

		// Skip hidden files
		if filepath.HasPrefix(info.Name(), ".") {
			return nil
		}

		if !info.IsDir() {
			fileList = append(fileList, path)
		}

		return nil
	})
	return fileList, err
}

func searchFiles(fileList []string) error {
	for _, file := range fileList {
		fmt.Println("Reading: " + file)

		fh, err := os.Open(file)
		f := bufio.NewReader(fh)
		if err != nil {
			return errors.New("Failed opening file")
		}

		buf := make([]byte, 1024)
		for {
			buf, _, err = f.ReadLine()
			if err != nil {
				break
				// t.Fatal("Failed reading file")
			}

			if re.MatchString(string(buf)) {
				fmt.Printf("%v\n", findMatch(string(buf)))
			}
		}
		fh.Close()
	}
	return nil
}

func FindWorkarounds(path string) error {
	fileList, err := makeFileList(path)
	if err != nil {
		return err
	}
	return searchFiles(fileList)
}
