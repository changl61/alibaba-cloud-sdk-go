package cs

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

func (client *Client) DescribeClusterScaledNode(request *DescribeClusterScaledNodeRequest) (response *DescribeClusterScaledNodeResponse, err error) {
	response = CreateDescribeClusterScaledNodeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeClusterScaledNodeWithChan(request *DescribeClusterScaledNodeRequest) (<-chan *DescribeClusterScaledNodeResponse, <-chan error) {
	responseChan := make(chan *DescribeClusterScaledNodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeClusterScaledNode(request)
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

func (client *Client) DescribeClusterScaledNodeWithCallback(request *DescribeClusterScaledNodeRequest, callback func(response *DescribeClusterScaledNodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeClusterScaledNodeResponse
		var err error
		defer close(result)
		response, err = client.DescribeClusterScaledNode(request)
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

type DescribeClusterScaledNodeRequest struct {
	*requests.RoaRequest
	ClusterId string `position:"Path" name:"ClusterId"`
}

type DescribeClusterScaledNodeResponse struct {
	*responses.BaseResponse
}

func CreateDescribeClusterScaledNodeRequest() (request *DescribeClusterScaledNodeRequest) {
	request = &DescribeClusterScaledNodeRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "DescribeClusterScaledNode", "/clusters/[ClusterId]/scaled_nodes/", "", "")
	request.Method = requests.GET
	return
}

func CreateDescribeClusterScaledNodeResponse() (response *DescribeClusterScaledNodeResponse) {
	response = &DescribeClusterScaledNodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
