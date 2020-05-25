# ccp client demo

## api server 

```sh
env DOMAIN_ID=hz694 ENDPOINT=hz694.api.alicloudccp.com ACCESS_KEY_ID=$ACCESS_KEY_ID ACCESS_KEY_SECRET=$ACCESS_KEY_SECRET DRIVE_ID=1 go run main.go
```

## api documentation

endpoint: `http://localhost:8080`

获取所有文件

```sh
curl -X GET "http://localhost:8080/files"
```

上传文件 文件参数`file=$file`

```sh 
## files create
curl -X "POST" "http://localhost:8080/files/create" \
     -H 'Content-Type: multipart/form-data; charset=utf-8; boundary=__X_PAW_BOUNDARY__' \
     -F "file="
```

删除文件

```sh
## delete file
curl -X "DELETE" "http://localhost:8080/files?file_id=5ecbb8885e045e37ce8f4c7c86a755c97dde275e"
```

## links

* ccp dashboard https://ccp.console.aliyun.com/
* ccp docs https://help.aliyun.com/document_detail/134409.html
* ccp AccessKey调接口接入 https://help.aliyun.com/document_detail/142201.html
* sdk https://github.com/alibabacloud-go/ccppath-sdk
