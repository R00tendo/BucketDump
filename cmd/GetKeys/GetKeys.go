package GetKeys

import (
	"os"
	"strings"

	"github.com/R00tendo/BucketDump/cmd/Log"
	"github.com/R00tendo/BucketDump/cmd/RetUrl"
)

func Get(AWSUrl string) []string {
	var keys []string

	body := string(RetUrl.Get(AWSUrl))
	if !strings.Contains(body, "<ListBucketResult xmlns") {
		Log.Msg("Bucket doesn't have listing enabled", "error")
		os.Exit(0)
	}
	segments := strings.Split(body, "<Key>")
	for _, segment := range segments {
		if strings.Contains(segment, "</Key>") {
			keys = append(keys, strings.Split(segment, "</Key>")[0])
		}
	}

	return keys
}
