
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

func (client *Client) AddTerrorismPipeline(request *AddTerrorismPipelineRequest) (response *AddTerrorismPipelineResponse, err error) {
response = CreateAddTerrorismPipelineResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) AddTerrorismPipelineWithChan(request *AddTerrorismPipelineRequest) (<-chan *AddTerrorismPipelineResponse, <-chan error) {
responseChan := make(chan *AddTerrorismPipelineResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.AddTerrorismPipeline(request)
responseChan <- response
errChan <- err
})
if err != nil {
errChan <- err
close(responseChan)
close(errChan)
}
return responseChan, errChan
}

func (client *Client) AddTerrorismPipelineWithCallback(request *AddTerrorismPipelineRequest, callback func(response *AddTerrorismPipelineResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *AddTerrorismPipelineResponse
var err error
defer close(result)
response, err = client.AddTerrorismPipeline(request)
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

type AddTerrorismPipelineRequest struct {
*requests.RpcRequest
                NotifyConfig  string `position:"Query" name:"NotifyConfig"`
                ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
                Priority  string `position:"Query" name:"Priority"`
                Name  string `position:"Query" name:"Name"`
                ResourceOwnerId  string `position:"Query" name:"ResourceOwnerId"`
                OwnerAccount  string `position:"Query" name:"OwnerAccount"`
                OwnerId  string `position:"Query" name:"OwnerId"`
}


type AddTerrorismPipelineResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            Pipeline struct {
            Id     string `json:"Id" xml:"Id"`
            Name     string `json:"Name" xml:"Name"`
            Priority     int `json:"Priority" xml:"Priority"`
            State     string `json:"State" xml:"State"`
            NotifyConfig struct {
            Topic     string `json:"Topic" xml:"Topic"`
            Queue     string `json:"Queue" xml:"Queue"`
            }  `json:"NotifyConfig" xml:"NotifyConfig"`
            }  `json:"Pipeline" xml:"Pipeline"`
}

func CreateAddTerrorismPipelineRequest() (request *AddTerrorismPipelineRequest) {
request = &AddTerrorismPipelineRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Mts", "2014-06-18", "AddTerrorismPipeline", "", "")
return
}

func CreateAddTerrorismPipelineResponse() (response *AddTerrorismPipelineResponse) {
response = &AddTerrorismPipelineResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}
