package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/R00tendo/BucketDump/cmd/DumpKeys"
	"github.com/R00tendo/BucketDump/cmd/GetKeys"
	"github.com/R00tendo/BucketDump/cmd/Log"
)

var args struct {
	bucketName string
	saveDir    string
}

func main() {
	fmt.Print("\033[33m" + `
╔══╗ ╔╗ ╔╗╔═══╗╔╗╔═╗╔═══╗╔════╗╔═══╗╔╗ ╔╗╔═╗╔═╗╔═══╗
║╔╗║ ║║ ║║║╔═╗║║║║╔╝║╔══╝║╔╗╔╗║╚╗╔╗║║║ ║║║║╚╝║║║╔═╗║  ╔═══════════╗
║╚╝╚╗║║ ║║║║ ╚╝║╚╝╝ ║╚══╗╚╝║║╚╝ ║║║║║║ ║║║╔╗╔╗║║╚═╝║  ║_ _o 0__0o
║╔═╗║║║ ║║║║ ╔╗║╔╗║ ║╔══╝  ║║   ║║║║║║ ║║║║║║║║║╔══╝  ║o0_0oO-oOo-O-0o-
║╚═╝║║╚═╝║║╚═╝║║║║╚╗║╚══╗ ╔╝╚╗ ╔╝╚╝║║╚═╝║║║║║║║║║     ╚═══════════╝ - o-O
╚═══╝╚═══╝╚═══╝╚╝╚═╝╚═══╝ ╚══╝ ╚═══╝╚═══╝╚╝╚╝╚╝╚╝                    - 0-0o_-
` + "\033[35m" + `With ❤️️   By @R00tendo` + "\033[33m" + `                                                0 o- O 0-
` + "\033[0m")
	flag.StringVar(&args.bucketName, "bucket", "", "Bucket to dump")
	flag.StringVar(&args.saveDir, "save-dir", "./out/", "Directory to dump the files to")
	flag.Parse()

	flag.VisitAll(func(curFlag *flag.Flag) {
		if len(curFlag.Value.String()) < 1 {
			flag.Usage()
			os.Exit(0)
		}
		if curFlag.Name == "out-dir" {
			if string(curFlag.Value.String()[len(curFlag.Value.String())-1]) != "/" {
				curFlag.Value.Set(curFlag.Value.String() + "/")
			}
		}
	})

	AWSUrl := fmt.Sprintf("https://%s.s3.amazonaws.com/", args.bucketName)

	keys := GetKeys.Get(AWSUrl)
	DumpKeys.Dump(keys, AWSUrl, args.saveDir)

	Log.Msg("All done! Take a loot at "+args.saveDir, "success")
}
