package download

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	bar "github.com/hideaki10/poogo/common/bar"
)

func Download(filename, url string, header http.Header) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	req.Header = header

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	conetentLength, err := strconv.ParseInt(response.Header["Content-Length"][0], 10, 64)
	if err != nil {
		return err
	}

	br := bar.NewBar(conetentLength)
	br.Resize = func(bar *bar.Bar) error {
		fileInfo, err := os.Stat(filename)
		if err != nil {
			return err
		}
		bar.Size = fileInfo.Size()
		return nil
	}
	br.Start()
	go func() {
		defer response.Body.Close()
		io.Copy(file, response.Body)
	}()
	br.ShowProgress()
	// compare file size
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return fmt.Errorf("file is not exist")
	}

	if conetentLength != fileInfo.Size() {
		return fmt.Errorf("download failed")
	}

	return nil
}
