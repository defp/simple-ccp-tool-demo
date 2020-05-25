# ccp client demo

## api server 

```sh
env DOMAIN_ID=hz694 ENDPOINT=hz694.api.alicloudccp.com ACCESS_KEY_ID=$ACCESS_KEY_ID ACCESS_KEY_SECRET=$ACCESS_KEY_SECRET DRIVE_ID=1 go run main.go
```

## api documentation

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

## links

* ccp dashboard https://ccp.console.aliyun.com/
* ccp docs https://help.aliyun.com/document_detail/134409.html
* ccp AccessKey调接口接入 https://help.aliyun.com/document_detail/142201.html
* sdk https://github.com/alibabacloud-go/ccppath-sdk
