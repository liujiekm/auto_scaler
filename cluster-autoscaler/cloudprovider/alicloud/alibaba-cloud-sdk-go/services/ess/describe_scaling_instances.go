/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ess

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/alicloud/alibaba-cloud-sdk-go/sdk/requests"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/alicloud/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeScalingInstances invokes the ess.DescribeScalingInstances API synchronously
// api document: https://help.aliyun.com/api/ess/describescalinginstances.html
func (client *Client) DescribeScalingInstances(request *DescribeScalingInstancesRequest) (response *DescribeScalingInstancesResponse, err error) {
	response = CreateDescribeScalingInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScalingInstancesWithChan invokes the ess.DescribeScalingInstances API asynchronously
// api document: https://help.aliyun.com/api/ess/describescalinginstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScalingInstancesWithChan(request *DescribeScalingInstancesRequest) (<-chan *DescribeScalingInstancesResponse, <-chan error) {
	responseChan := make(chan *DescribeScalingInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScalingInstances(request)
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

// DescribeScalingInstancesWithCallback invokes the ess.DescribeScalingInstances API asynchronously
// api document: https://help.aliyun.com/api/ess/describescalinginstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScalingInstancesWithCallback(request *DescribeScalingInstancesRequest, callback func(response *DescribeScalingInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScalingInstancesResponse
		var err error
		defer close(result)
		response, err = client.DescribeScalingInstances(request)
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

// DescribeScalingInstancesRequest is the request struct for api DescribeScalingInstances
type DescribeScalingInstancesRequest struct {
	*requests.RpcRequest
	InstanceId10           string           `position:"Query" name:"InstanceId.10"`
	ResourceOwnerId        requests.Integer `position:"Query" name:"ResourceOwnerId"`
	InstanceId12           string           `position:"Query" name:"InstanceId.12"`
	InstanceId11           string           `position:"Query" name:"InstanceId.11"`
	ScalingGroupId         string           `position:"Query" name:"ScalingGroupId"`
	LifecycleState         string           `position:"Query" name:"LifecycleState"`
	CreationType           string           `position:"Query" name:"CreationType"`
	PageNumber             requests.Integer `position:"Query" name:"PageNumber"`
	PageSize               requests.Integer `position:"Query" name:"PageSize"`
	InstanceId20           string           `position:"Query" name:"InstanceId.20"`
	InstanceId1            string           `position:"Query" name:"InstanceId.1"`
	InstanceId3            string           `position:"Query" name:"InstanceId.3"`
	ResourceOwnerAccount   string           `position:"Query" name:"ResourceOwnerAccount"`
	InstanceId2            string           `position:"Query" name:"InstanceId.2"`
	InstanceId5            string           `position:"Query" name:"InstanceId.5"`
	InstanceId4            string           `position:"Query" name:"InstanceId.4"`
	OwnerAccount           string           `position:"Query" name:"OwnerAccount"`
	InstanceId7            string           `position:"Query" name:"InstanceId.7"`
	InstanceId6            string           `position:"Query" name:"InstanceId.6"`
	InstanceId9            string           `position:"Query" name:"InstanceId.9"`
	InstanceId8            string           `position:"Query" name:"InstanceId.8"`
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
	ScalingConfigurationId string           `position:"Query" name:"ScalingConfigurationId"`
	HealthStatus           string           `position:"Query" name:"HealthStatus"`
	InstanceId18           string           `position:"Query" name:"InstanceId.18"`
	InstanceId17           string           `position:"Query" name:"InstanceId.17"`
	InstanceId19           string           `position:"Query" name:"InstanceId.19"`
	InstanceId14           string           `position:"Query" name:"InstanceId.14"`
	InstanceId13           string           `position:"Query" name:"InstanceId.13"`
	InstanceId16           string           `position:"Query" name:"InstanceId.16"`
	InstanceId15           string           `position:"Query" name:"InstanceId.15"`
}

// DescribeScalingInstancesResponse is the response struct for api DescribeScalingInstances
type DescribeScalingInstancesResponse struct {
	*responses.BaseResponse
	TotalCount       int              `json:"TotalCount" xml:"TotalCount"`
	PageNumber       int              `json:"PageNumber" xml:"PageNumber"`
	PageSize         int              `json:"PageSize" xml:"PageSize"`
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	ScalingInstances ScalingInstances `json:"ScalingInstances" xml:"ScalingInstances"`
}

// CreateDescribeScalingInstancesRequest creates a request to invoke DescribeScalingInstances API
func CreateDescribeScalingInstancesRequest() (request *DescribeScalingInstancesRequest) {
	request = &DescribeScalingInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "DescribeScalingInstances", "ess", "openAPI")
	return
}

// CreateDescribeScalingInstancesResponse creates a response to parse from DescribeScalingInstances response
func CreateDescribeScalingInstancesResponse() (response *DescribeScalingInstancesResponse) {
	response = &DescribeScalingInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

// ScalingInstances is a nested struct in ess response
type ScalingInstances struct {
	ScalingInstance []ScalingInstance `json:"ScalingInstance" xml:"ScalingInstance"`
}

// ScalingInstance is a nested struct in ess response
type ScalingInstance struct {
	InstanceId             string `json:"InstanceId" xml:"InstanceId"`
	ScalingConfigurationId string `json:"ScalingConfigurationId" xml:"ScalingConfigurationId"`
	ScalingGroupId         string `json:"ScalingGroupId" xml:"ScalingGroupId"`
	HealthStatus           string `json:"HealthStatus" xml:"HealthStatus"`
	LoadBalancerWeight     int    `json:"LoadBalancerWeight" xml:"LoadBalancerWeight"`
	LifecycleState         string `json:"LifecycleState" xml:"LifecycleState"`
	CreationTime           string `json:"CreationTime" xml:"CreationTime"`
	CreationType           string `json:"CreationType" xml:"CreationType"`
}
