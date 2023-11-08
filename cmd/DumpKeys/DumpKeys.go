package DumpKeys

import (
	"os"
	"strings"

	"github.com/R00tendo/BucketDump/cmd/ErrorCheck"
	"github.com/R00tendo/BucketDump/cmd/Log"
	"github.com/R00tendo/BucketDump/cmd/RetUrl"
)

func Dump(keys []string, AWSUrl string, saveDir string) {
	for _, key := range keys {
		Log.Msg("Dumping:"+key, "info")
		keySegments := strings.Split(key, "/")
		os.MkdirAll(saveDir+strings.Join(keySegments[:len(keySegments)-1], "/"), 0600)

		keyUrl := AWSUrl + key

		content := RetUrl.Get(keyUrl)

		if stat, err := os.Stat(saveDir + key); err == nil {
			if stat.IsDir() {
				continue
			}
		}
		outputHandle, err := os.OpenFile(saveDir+key, os.O_CREATE|os.O_WRONLY, 0600)
		ErrorCheck.Check(err)

		_, err = outputHandle.Write(content)
		ErrorCheck.Check(err)
	}
}
