package ccp_utils

import (
	"log"
	"os"

	ccpsdk "github.com/alibabacloud-go/ccppath-sdk/client"
)

func ListFiles(akClient *ccpsdk.Client) (*ccpsdk.CCPListFileResponse, error)  {
	listFileRequestModel := new(ccpsdk.ListFileRequestModel)

	listFileRequest := new(ccpsdk.CCPListFileRequest)
	listFileRequest.SetParentFileId("root")
	listFileRequest.SetDriveId(os.Getenv("DRIVE_ID"))

	listFileRequestModel.SetBody(listFileRequest)

	runtimeOptions := new(ccpsdk.RuntimeOptions).SetMaxAttempts(2)

	response, err := akClient.ListFile(listFileRequestModel, runtimeOptions)
	if err != nil {
		return nil, err
	} else {
		return response.Body, nil
	}

}
