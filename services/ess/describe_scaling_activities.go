package ess

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

func (client *Client) DescribeScalingActivities(request *DescribeScalingActivitiesRequest) (response *DescribeScalingActivitiesResponse, err error) {
	response = CreateDescribeScalingActivitiesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeScalingActivitiesWithChan(request *DescribeScalingActivitiesRequest) (<-chan *DescribeScalingActivitiesResponse, <-chan error) {
	responseChan := make(chan *DescribeScalingActivitiesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScalingActivities(request)
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

func (client *Client) DescribeScalingActivitiesWithCallback(request *DescribeScalingActivitiesRequest, callback func(response *DescribeScalingActivitiesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScalingActivitiesResponse
		var err error
		defer close(result)
		response, err = client.DescribeScalingActivities(request)
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

type DescribeScalingActivitiesRequest struct {
	*requests.RpcRequest
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	ScalingActivityId9   string           `position:"Query" name:"ScalingActivityId.9"`
	ScalingActivityId5   string           `position:"Query" name:"ScalingActivityId.5"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	ScalingActivityId6   string           `position:"Query" name:"ScalingActivityId.6"`
	ScalingActivityId7   string           `position:"Query" name:"ScalingActivityId.7"`
	ScalingActivityId8   string           `position:"Query" name:"ScalingActivityId.8"`
	ScalingActivityId1   string           `position:"Query" name:"ScalingActivityId.1"`
	ScalingActivityId2   string           `position:"Query" name:"ScalingActivityId.2"`
	ScalingActivityId3   string           `position:"Query" name:"ScalingActivityId.3"`
	ScalingActivityId4   string           `position:"Query" name:"ScalingActivityId.4"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ScalingActivityId20  string           `position:"Query" name:"ScalingActivityId.20"`
	ScalingGroupId       string           `position:"Query" name:"ScalingGroupId"`
	ScalingActivityId13  string           `position:"Query" name:"ScalingActivityId.13"`
	ScalingActivityId12  string           `position:"Query" name:"ScalingActivityId.12"`
	ScalingActivityId19  string           `position:"Query" name:"ScalingActivityId.19"`
	ScalingActivityId11  string           `position:"Query" name:"ScalingActivityId.11"`
	ScalingActivityId18  string           `position:"Query" name:"ScalingActivityId.18"`
	ScalingActivityId10  string           `position:"Query" name:"ScalingActivityId.10"`
	ScalingActivityId17  string           `position:"Query" name:"ScalingActivityId.17"`
	ScalingActivityId16  string           `position:"Query" name:"ScalingActivityId.16"`
	ScalingActivityId15  string           `position:"Query" name:"ScalingActivityId.15"`
	ScalingActivityId14  string           `position:"Query" name:"ScalingActivityId.14"`
	StatusCode           string           `position:"Query" name:"StatusCode"`
}

type DescribeScalingActivitiesResponse struct {
	*responses.BaseResponse
	TotalCount        int    `json:"TotalCount" xml:"TotalCount"`
	PageNumber        int    `json:"PageNumber" xml:"PageNumber"`
	PageSize          int    `json:"PageSize" xml:"PageSize"`
	RequestId         string `json:"RequestId" xml:"RequestId"`
	ScalingActivities struct {
		ScalingActivity []struct {
			ScalingActivityId   string `json:"ScalingActivityId" xml:"ScalingActivityId"`
			ScalingGroupId      string `json:"ScalingGroupId" xml:"ScalingGroupId"`
			Description         string `json:"Description" xml:"Description"`
			Cause               string `json:"Cause" xml:"Cause"`
			StartTime           string `json:"StartTime" xml:"StartTime"`
			EndTime             string `json:"EndTime" xml:"EndTime"`
			Progress            int    `json:"Progress" xml:"Progress"`
			StatusCode          string `json:"StatusCode" xml:"StatusCode"`
			StatusMessage       string `json:"StatusMessage" xml:"StatusMessage"`
			TotalCapacity       string `json:"TotalCapacity" xml:"TotalCapacity"`
			AttachedCapacity    string `json:"AttachedCapacity" xml:"AttachedCapacity"`
			AutoCreatedCapacity string `json:"AutoCreatedCapacity" xml:"AutoCreatedCapacity"`
		} `json:"ScalingActivity" xml:"ScalingActivity"`
	} `json:"ScalingActivities" xml:"ScalingActivities"`
}

func CreateDescribeScalingActivitiesRequest() (request *DescribeScalingActivitiesRequest) {
	request = &DescribeScalingActivitiesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "DescribeScalingActivities", "ess", "openAPI")
	return
}

func CreateDescribeScalingActivitiesResponse() (response *DescribeScalingActivitiesResponse) {
	response = &DescribeScalingActivitiesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
