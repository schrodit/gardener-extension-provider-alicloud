package slb

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

// AddListenerWhiteListItem invokes the slb.AddListenerWhiteListItem API synchronously
// api document: https://help.aliyun.com/api/slb/addlistenerwhitelistitem.html
func (client *Client) AddListenerWhiteListItem(request *AddListenerWhiteListItemRequest) (response *AddListenerWhiteListItemResponse, err error) {
	response = CreateAddListenerWhiteListItemResponse()
	err = client.DoAction(request, response)
	return
}

// AddListenerWhiteListItemWithChan invokes the slb.AddListenerWhiteListItem API asynchronously
// api document: https://help.aliyun.com/api/slb/addlistenerwhitelistitem.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddListenerWhiteListItemWithChan(request *AddListenerWhiteListItemRequest) (<-chan *AddListenerWhiteListItemResponse, <-chan error) {
	responseChan := make(chan *AddListenerWhiteListItemResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddListenerWhiteListItem(request)
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

// AddListenerWhiteListItemWithCallback invokes the slb.AddListenerWhiteListItem API asynchronously
// api document: https://help.aliyun.com/api/slb/addlistenerwhitelistitem.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddListenerWhiteListItemWithCallback(request *AddListenerWhiteListItemRequest, callback func(response *AddListenerWhiteListItemResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddListenerWhiteListItemResponse
		var err error
		defer close(result)
		response, err = client.AddListenerWhiteListItem(request)
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

// AddListenerWhiteListItemRequest is the request struct for api AddListenerWhiteListItem
type AddListenerWhiteListItemRequest struct {
	*requests.RpcRequest
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SourceItems          string           `position:"Query" name:"SourceItems"`
	ListenerPort         requests.Integer `position:"Query" name:"ListenerPort"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ListenerProtocol     string           `position:"Query" name:"ListenerProtocol"`
	Tags                 string           `position:"Query" name:"Tags"`
	LoadBalancerId       string           `position:"Query" name:"LoadBalancerId"`
}

// AddListenerWhiteListItemResponse is the response struct for api AddListenerWhiteListItem
type AddListenerWhiteListItemResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddListenerWhiteListItemRequest creates a request to invoke AddListenerWhiteListItem API
func CreateAddListenerWhiteListItemRequest() (request *AddListenerWhiteListItemRequest) {
	request = &AddListenerWhiteListItemRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "AddListenerWhiteListItem", "slb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddListenerWhiteListItemResponse creates a response to parse from AddListenerWhiteListItem response
func CreateAddListenerWhiteListItemResponse() (response *AddListenerWhiteListItemResponse) {
	response = &AddListenerWhiteListItemResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
