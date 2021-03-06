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

func (client *Client) DescribeLiveStreamFrameInfo(request *DescribeLiveStreamFrameInfoRequest) (response *DescribeLiveStreamFrameInfoResponse, err error) {
	response = CreateDescribeLiveStreamFrameInfoResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeLiveStreamFrameInfoWithChan(request *DescribeLiveStreamFrameInfoRequest) (<-chan *DescribeLiveStreamFrameInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveStreamFrameInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveStreamFrameInfo(request)
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

func (client *Client) DescribeLiveStreamFrameInfoWithCallback(request *DescribeLiveStreamFrameInfoRequest, callback func(response *DescribeLiveStreamFrameInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveStreamFrameInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveStreamFrameInfo(request)
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

type DescribeLiveStreamFrameInfoRequest struct {
	*requests.RpcRequest
	EndTime       string           `position:"Query" name:"EndTime"`
	StreamName    string           `position:"Query" name:"StreamName"`
	StartTime     string           `position:"Query" name:"StartTime"`
	DomainName    string           `position:"Query" name:"DomainName"`
	AppName       string           `position:"Query" name:"AppName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

type DescribeLiveStreamFrameInfoResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	FrameDataInfos struct {
		FrameDataModel []struct {
			Time       string  `json:"Time" xml:"Time"`
			Stream     string  `json:"Stream" xml:"Stream"`
			ClientAddr string  `json:"ClientAddr" xml:"ClientAddr"`
			Server     string  `json:"Server" xml:"Server"`
			AudioRate  float64 `json:"AudioRate" xml:"AudioRate"`
			AudioByte  float64 `json:"AudioByte" xml:"AudioByte"`
			FrameRate  float64 `json:"FrameRate" xml:"FrameRate"`
			FrameByte  float64 `json:"FrameByte" xml:"FrameByte"`
		} `json:"FrameDataModel" xml:"FrameDataModel"`
	} `json:"FrameDataInfos" xml:"FrameDataInfos"`
}

func CreateDescribeLiveStreamFrameInfoRequest() (request *DescribeLiveStreamFrameInfoRequest) {
	request = &DescribeLiveStreamFrameInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamFrameInfo", "", "")
	return
}

func CreateDescribeLiveStreamFrameInfoResponse() (response *DescribeLiveStreamFrameInfoResponse) {
	response = &DescribeLiveStreamFrameInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
