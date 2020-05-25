package ccp_utils

import (
	"io"
	"net/http"
	"os"

	ccpsdk "github.com/alibabacloud-go/ccppath-sdk/client"
)

var runtimeOptions = new(ccpsdk.RuntimeOptions).SetMaxAttempts(2)

func ListFiles(akClient *ccpsdk.Client) (*ccpsdk.CCPListFileResponse, error) {
	listFileRequestModel := new(ccpsdk.ListFileRequestModel)

	listFileRequest := new(ccpsdk.CCPListFileRequest)
	listFileRequest.SetParentFileId("root")
	listFileRequest.SetDriveId(os.Getenv("DRIVE_ID"))

	listFileRequestModel.SetBody(listFileRequest)

	response, err := akClient.ListFile(listFileRequestModel, runtimeOptions)
	if err != nil {
		return nil, err
	}

	return response.Body, nil

}

func CreateFile(akClient *ccpsdk.Client, filename string) (*ccpsdk.CCPCreateFileResponse, error) {
	createFileModel := new(ccpsdk.CreateFileRequestModel)
	createFileRequest := new(ccpsdk.CCPCreateFileRequest).
		SetDriveId(os.Getenv("DRIVE_ID")).
		SetName(filename).
		SetType("file").
		SetParentFileId("root").
		SetContentType("text/plain")
	createFileModel.SetBody(createFileRequest)

	response, err := akClient.CreateFile(createFileModel, runtimeOptions)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

func CompleteFile(akClient *ccpsdk.Client, createFileBody *ccpsdk.CCPCreateFileResponse, content io.Reader) (*ccpsdk.CCPCompleteFileResponse, error) {
	uploadUrl := createFileBody.PartInfoList[0].UploadUrl
	uploadId := createFileBody.UploadId
	fileId := createFileBody.FileId

	req, _ := http.NewRequest("PUT", *uploadUrl, content)
	req.Header.Add("Content-Type", "")
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	etag := res.Header.Get("ETag")

	uploadPartInfo := new(ccpsdk.UploadPartInfo)
	uploadPartInfo.SetEtag(etag)
	uploadPartInfo.SetPartNumber(1)

	completeFileModel := new(ccpsdk.CompleteFileRequestModel)
	completeFileRequest := new(ccpsdk.CCPCompleteFileRequest).
		SetDriveId(os.Getenv("DRIVE_ID")).
		SetFileId(*fileId).
		SetUploadId(*uploadId).
		SetPartInfoList([]*ccpsdk.UploadPartInfo{uploadPartInfo})
	completeFileModel.SetBody(completeFileRequest)

	completeResponse, err := akClient.CompleteFile(completeFileModel, runtimeOptions)
	if err != nil {
		return nil, err
	}
	return completeResponse.Body, nil
}
