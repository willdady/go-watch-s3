# go-watch-s3

Watches a directory recursively for new files and uploads them to S3.

## Configuration

Configuration is controlled by environment variables. The following environment variables MUST be set.

```
WATCH_PATH=/absolute/path/to/files/
AWS_REGION=ap-southeast-2
AWS_S3_BUCKET=example.com
```

The following environment variables are OPTIONAL (default values shown).

```
AWS_S3_KEY_PREFIX=""
AWS_S3_STORAGE_CLASS=STANDARD
WATCH_INTERVAL=100
```
