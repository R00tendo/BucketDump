![logo](https://github.com/R00tendo/BucketDump/assets/72181445/921bd9c0-7354-4f33-af3d-8c544120b266)

## Description
<b>A minimalistic tool to dump a whole S3 bucket to your local computer WIHTOUT AN API KEY. Note that for this program to work the main page of the bucket has to be bucket listing.</b>

## Installation
```
go install https://github.com/R00tendo/BucketDump@latest
```

## Help page
```
Usage of BucketDump:
  -bucket string
        Bucket to dump
  -save-dir string
        Directory to dump the files to (default "./out/")
```

## Examples
```
BucketDump -bucket randombucket -save-dir ./randombucket1

BucketDump -bucket randombucket
```
