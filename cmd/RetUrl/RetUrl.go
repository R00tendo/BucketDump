package RetUrl

import (
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/R00tendo/BucketDump/cmd/ErrorCheck"
)

func Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if resp.StatusCode != 200 {
		return []byte{}, errors.New("Code " + strconv.Itoa(resp.StatusCode) + ":  " + url)
	}
	ErrorCheck.Check(err, 1)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	ErrorCheck.Check(err, 1)

	return body, nil
}
