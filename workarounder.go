package workarounder

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"

	"github.com/olekukonko/tablewriter"
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

func findMatchInFiles(fileList []string) (data [][]string, err error) {
	for _, file := range fileList {
		// fmt.Println("Reading: " + file)

		fh, err := os.Open(file)
		f := bufio.NewReader(fh)
		if err != nil {
			return data, errors.New("Failed opening file")
		}

		buf := make([]byte, 1024)
		for {
			buf, _, err = f.ReadLine()
			if err == io.EOF {
				break
			}

			if err != nil {
				fmt.Printf("Failed reading file: %v\n, err: %v\n\n\n", file, err)
				break
			}

			results := findMatch(string(buf))
			for _, matches := range results {
				for _, match := range matches {
					// fmt.Printf("%v   (%v)\n", match, file)
					data = append(data, []string{match, file})
				}
			}
		}
		fh.Close()
	}
	return data, nil
}

func printTable(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Issue", "File"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}

func FindWorkarounds(path string) error {
	fileList, err := makeFileList(path)
	if err != nil {
		return err
	}
	data, err := findMatchInFiles(fileList)
	if err != nil {
		return err
	}
	printTable(data)
	return nil
}
