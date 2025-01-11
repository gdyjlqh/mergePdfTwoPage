package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("please input 2 params,param1:need to merge pdf files dir,param2:out put pdf file")
	}

	pdfFiles, err := getPDFFiles(args[1])
	outFile := args[2]

	tmpOutFile := outFile + "._temp.pdf"

	if err != nil {
		fmt.Errorf("error:%s\n", err)
	}
	api.MergeCreateFile(pdfFiles, tmpOutFile, false, nil)

	nup, err := api.PDFGridConfig(2, 1, "", nil)
	api.NUpFile([]string{tmpOutFile}, outFile, nil, nup, nil)
}

// 获取指定目录下的所有PDF文件
func getPDFFiles(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".pdf") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
