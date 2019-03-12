# go-streamer
## Upload to S3 :-

`curl -F "file=@<file>" localhost:8192`

It returns ID of the uploaded file and the URL.
You can use this **ID** to download the file further.

---

## Download from S3:-

`curl -XGET localhost:8192/<ID>`