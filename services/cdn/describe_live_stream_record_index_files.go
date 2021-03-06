package cdn

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) DescribeLiveStreamRecordIndexFiles(request *DescribeLiveStreamRecordIndexFilesRequest) (response *DescribeLiveStreamRecordIndexFilesResponse, err error) {
	response = CreateDescribeLiveStreamRecordIndexFilesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeLiveStreamRecordIndexFilesWithChan(request *DescribeLiveStreamRecordIndexFilesRequest) (<-chan *DescribeLiveStreamRecordIndexFilesResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveStreamRecordIndexFilesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveStreamRecordIndexFiles(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}

	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) DescribeLiveStreamRecordIndexFilesWithCallback(request *DescribeLiveStreamRecordIndexFilesRequest, callback func(response *DescribeLiveStreamRecordIndexFilesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveStreamRecordIndexFilesResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveStreamRecordIndexFiles(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

type DescribeLiveStreamRecordIndexFilesRequest struct {
	*requests.RpcRequest
	EndTime       string           `position:"Query" name:"EndTime"`
	StreamName    string           `position:"Query" name:"StreamName"`
	StartTime     string           `position:"Query" name:"StartTime"`
	DomainName    string           `position:"Query" name:"DomainName"`
	AppName       string           `position:"Query" name:"AppName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

type DescribeLiveStreamRecordIndexFilesResponse struct {
	*responses.BaseResponse
	RequestId           string `json:"RequestId" xml:"RequestId"`
	RecordIndexInfoList struct {
		RecordIndexInfo []struct {
			RecordId   string  `json:"RecordId" xml:"RecordId"`
			RecordUrl  string  `json:"RecordUrl" xml:"RecordUrl"`
			DomainName string  `json:"DomainName" xml:"DomainName"`
			AppName    string  `json:"AppName" xml:"AppName"`
			StreamName string  `json:"StreamName" xml:"StreamName"`
			OssObject  string  `json:"OssObject" xml:"OssObject"`
			StartTime  string  `json:"StartTime" xml:"StartTime"`
			EndTime    string  `json:"EndTime" xml:"EndTime"`
			Duration   float64 `json:"Duration" xml:"Duration"`
			Height     int     `json:"Height" xml:"Height"`
			Width      int     `json:"Width" xml:"Width"`
			CreateTime string  `json:"CreateTime" xml:"CreateTime"`
		} `json:"RecordIndexInfo" xml:"RecordIndexInfo"`
	} `json:"RecordIndexInfoList" xml:"RecordIndexInfoList"`
}

func CreateDescribeLiveStreamRecordIndexFilesRequest() (request *DescribeLiveStreamRecordIndexFilesRequest) {
	request = &DescribeLiveStreamRecordIndexFilesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamRecordIndexFiles", "", "")
	return
}

func CreateDescribeLiveStreamRecordIndexFilesResponse() (response *DescribeLiveStreamRecordIndexFilesResponse) {
	response = &DescribeLiveStreamRecordIndexFilesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
