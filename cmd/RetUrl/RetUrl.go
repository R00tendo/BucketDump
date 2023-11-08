package RetUrl

import (
	"io"
	"net/http"

	"github.com/R00tendo/BucketDump/cmd/ErrorCheck"
)

func Get(url string) []byte {
	resp, err := http.Get(url)
	ErrorCheck.Check(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	ErrorCheck.Check(err)

	return body
}
