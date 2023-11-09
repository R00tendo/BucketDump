![Logo](https://github.com/R00tendo/BucketDump/assets/72181445/d005bad0-1cd7-4dc3-af6b-cf62e6210590)

## Description
<b>A minimalistic tool to dump a whole S3 bucket to your local computer WIHTOUT AN API KEY. Note that for this program to work the main page of the bucket has to be bucket listing.</b>

## Installation
```
go install https://github.com/R00tendo/BucketDump@latest
```

## Usage
### Help page
```
Usage of BucketDump:
  -bucket string
        Bucket to dump
  -save-dir string
        Directory to dump the files to (default "./out/")
```
### Commands
```
BucketDump -bucket randombucket -save-dir ./randombucket1

BucketDump -bucket randombucket
```

## Example
<a href="https://asciinema.org/a/620026" target="_blank"><img src="https://asciinema.org/a/620026.svg" /></a>
