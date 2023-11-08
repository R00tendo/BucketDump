package ErrorCheck

import (
	"os"

	"github.com/R00tendo/BucketDump/cmd/Log"
)

func Check(err error) {
	if err != nil {
		Log.Msg(err.Error(), "error")
		os.Exit(0)
	}
}
