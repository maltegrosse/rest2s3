# rest2s3

## Goal
simple rest interface to upload to minio/s3 bucket

## run 
`docker pull cspecht/rest2s3:latest`

## curl to post 
curl -F fileUpload=@file.png http://localhost:8080/upload

## Environment Variables
- MINIO_ENDPOINT
- MINIO_PORT
- MINIO_ACCESSKEY
- MINIO_SECRETKEY
- MINIO_BUCKET

## License
- under **[MIT license](http://opensource.org/licenses/mit-license.php)**
