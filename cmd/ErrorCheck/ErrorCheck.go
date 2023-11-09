package ErrorCheck

import (
	"os"

	"github.com/R00tendo/BucketDump/cmd/Log"
)

func Check(err error, severity int) {
	if err != nil {
		Log.Msg(err.Error(), "error")
		if severity == 1 {
			os.Exit(0)
		}
	}
}
