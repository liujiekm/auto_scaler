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

// ModifyScalingGroup invokes the ess.ModifyScalingGroup API synchronously
// api document: https://help.aliyun.com/api/ess/modifyscalinggroup.html
func (client *Client) ModifyScalingGroup(request *ModifyScalingGroupRequest) (response *ModifyScalingGroupResponse, err error) {
	response = CreateModifyScalingGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyScalingGroupWithChan invokes the ess.ModifyScalingGroup API asynchronously
// api document: https://help.aliyun.com/api/ess/modifyscalinggroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyScalingGroupWithChan(request *ModifyScalingGroupRequest) (<-chan *ModifyScalingGroupResponse, <-chan error) {
	responseChan := make(chan *ModifyScalingGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyScalingGroup(request)
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

// ModifyScalingGroupWithCallback invokes the ess.ModifyScalingGroup API asynchronously
// api document: https://help.aliyun.com/api/ess/modifyscalinggroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyScalingGroupWithCallback(request *ModifyScalingGroupRequest, callback func(response *ModifyScalingGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyScalingGroupResponse
		var err error
		defer close(result)
		response, err = client.ModifyScalingGroup(request)
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

// ModifyScalingGroupRequest is the request struct for api ModifyScalingGroup
type ModifyScalingGroupRequest struct {
	*requests.RpcRequest
	ResourceOwnerId              requests.Integer `position:"Query" name:"ResourceOwnerId"`
	HealthCheckType              string           `position:"Query" name:"HealthCheckType"`
	LaunchTemplateId             string           `position:"Query" name:"LaunchTemplateId"`
	ResourceOwnerAccount         string           `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupName             string           `position:"Query" name:"ScalingGroupName"`
	ScalingGroupId               string           `position:"Query" name:"ScalingGroupId"`
	OwnerAccount                 string           `position:"Query" name:"OwnerAccount"`
	ActiveScalingConfigurationId string           `position:"Query" name:"ActiveScalingConfigurationId"`
	MinSize                      requests.Integer `position:"Query" name:"MinSize"`
	OwnerId                      requests.Integer `position:"Query" name:"OwnerId"`
	LaunchTemplateVersion        string           `position:"Query" name:"LaunchTemplateVersion"`
	MaxSize                      requests.Integer `position:"Query" name:"MaxSize"`
	DefaultCooldown              requests.Integer `position:"Query" name:"DefaultCooldown"`
	RemovalPolicy1               string           `position:"Query" name:"RemovalPolicy.1"`
	RemovalPolicy2               string           `position:"Query" name:"RemovalPolicy.2"`
}

// ModifyScalingGroupResponse is the response struct for api ModifyScalingGroup
type ModifyScalingGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyScalingGroupRequest creates a request to invoke ModifyScalingGroup API
func CreateModifyScalingGroupRequest() (request *ModifyScalingGroupRequest) {
	request = &ModifyScalingGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "ModifyScalingGroup", "ess", "openAPI")
	return
}

// CreateModifyScalingGroupResponse creates a response to parse from ModifyScalingGroup response
func CreateModifyScalingGroupResponse() (response *ModifyScalingGroupResponse) {
	response = &ModifyScalingGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
