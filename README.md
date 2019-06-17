# go-watch-s3

Watches a directory recursively for new files and uploads them to S3.

## Configuration

Configuration is controlled by environment variables.

**WATCH_PATH**

Absolute path to directory to watch. Must include trailing slash. e.g. _/absolute/path/to/files/_

**AWS_REGION**

AWS region e.g. _ap-southeast-2_

**AWS_S3_BUCKET**

Name of S3 bucket files will be uploaded to e.g. _mybucket.example.com_

**AWS_S3_KEY_PREFIX** (_Optional_)

String to prepend to file names when constructing keys. Defaults to an empty string.

**AWS_S3_STORAGE_CLASS** (_Optional_)

S3 storage class used for storing objects. Can be **STANDARD**, **STANDARD_IA** or **ONEZONE_IA**. Defaults to **STANDARD**.

**WATCH_INTERVAL** (_Optional_)

Duration in milliseconds specifying how often the watch directory should be polled for new files. Defaults to 500.

## Docker

Build Docker image by running the following from the repository root.

```bash
docker build -t willdady/go-watch-s3:latest .
```
