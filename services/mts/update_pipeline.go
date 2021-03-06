package mts

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

func (client *Client) UpdatePipeline(request *UpdatePipelineRequest) (response *UpdatePipelineResponse, err error) {
	response = CreateUpdatePipelineResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) UpdatePipelineWithChan(request *UpdatePipelineRequest) (<-chan *UpdatePipelineResponse, <-chan error) {
	responseChan := make(chan *UpdatePipelineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdatePipeline(request)
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

func (client *Client) UpdatePipelineWithCallback(request *UpdatePipelineRequest, callback func(response *UpdatePipelineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdatePipelineResponse
		var err error
		defer close(result)
		response, err = client.UpdatePipeline(request)
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

type UpdatePipelineRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Role                 string           `position:"Query" name:"Role"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Name                 string           `position:"Query" name:"Name"`
	State                string           `position:"Query" name:"State"`
	NotifyConfig         string           `position:"Query" name:"NotifyConfig"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PipelineId           string           `position:"Query" name:"PipelineId"`
}

type UpdatePipelineResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Pipeline  struct {
		Id           string `json:"Id" xml:"Id"`
		Name         string `json:"Name" xml:"Name"`
		State        string `json:"State" xml:"State"`
		Speed        string `json:"Speed" xml:"Speed"`
		Role         string `json:"Role" xml:"Role"`
		NotifyConfig struct {
			Topic     string `json:"Topic" xml:"Topic"`
			QueueName string `json:"QueueName" xml:"QueueName"`
		} `json:"NotifyConfig" xml:"NotifyConfig"`
	} `json:"Pipeline" xml:"Pipeline"`
}

func CreateUpdatePipelineRequest() (request *UpdatePipelineRequest) {
	request = &UpdatePipelineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "UpdatePipeline", "mts", "openAPI")
	return
}

func CreateUpdatePipelineResponse() (response *UpdatePipelineResponse) {
	response = &UpdatePipelineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
