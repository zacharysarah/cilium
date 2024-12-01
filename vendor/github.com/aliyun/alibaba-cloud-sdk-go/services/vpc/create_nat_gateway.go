package vpc

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

// CreateNatGateway invokes the vpc.CreateNatGateway API synchronously
func (client *Client) CreateNatGateway(request *CreateNatGatewayRequest) (response *CreateNatGatewayResponse, err error) {
	response = CreateCreateNatGatewayResponse()
	err = client.DoAction(request, response)
	return
}

// CreateNatGatewayWithChan invokes the vpc.CreateNatGateway API asynchronously
func (client *Client) CreateNatGatewayWithChan(request *CreateNatGatewayRequest) (<-chan *CreateNatGatewayResponse, <-chan error) {
	responseChan := make(chan *CreateNatGatewayResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateNatGateway(request)
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

// CreateNatGatewayWithCallback invokes the vpc.CreateNatGateway API asynchronously
func (client *Client) CreateNatGatewayWithCallback(request *CreateNatGatewayRequest, callback func(response *CreateNatGatewayResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateNatGatewayResponse
		var err error
		defer close(result)
		response, err = client.CreateNatGateway(request)
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

// CreateNatGatewayRequest is the request struct for api CreateNatGateway
type CreateNatGatewayRequest struct {
	*requests.RpcRequest
	ResourceOwnerId           requests.Integer                    `position:"Query" name:"ResourceOwnerId"`
	ClientToken               string                              `position:"Query" name:"ClientToken"`
	SecurityProtectionEnabled requests.Boolean                    `position:"Query" name:"SecurityProtectionEnabled"`
	SecurityGroupId           string                              `position:"Query" name:"SecurityGroupId"`
	Description               string                              `position:"Query" name:"Description"`
	NetworkType               string                              `position:"Query" name:"NetworkType"`
	Spec                      string                              `position:"Query" name:"Spec"`
	Duration                  string                              `position:"Query" name:"Duration"`
	IcmpReplyEnabled          requests.Boolean                    `position:"Query" name:"IcmpReplyEnabled"`
	NatType                   string                              `position:"Query" name:"NatType"`
	Tag                       *[]CreateNatGatewayTag              `position:"Query" name:"Tag"  type:"Repeated"`
	InstanceChargeType        string                              `position:"Query" name:"InstanceChargeType"`
	BandwidthPackage          *[]CreateNatGatewayBandwidthPackage `position:"Query" name:"BandwidthPackage"  type:"Repeated"`
	AutoPay                   requests.Boolean                    `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount      string                              `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount              string                              `position:"Query" name:"OwnerAccount"`
	PrivateLinkMode           string                              `position:"Query" name:"PrivateLinkMode"`
	OwnerId                   requests.Integer                    `position:"Query" name:"OwnerId"`
	IsCreateDefaultRoute      requests.Boolean                    `position:"Query" name:"IsCreateDefaultRoute"`
	VSwitchId                 string                              `position:"Query" name:"VSwitchId"`
	InternetChargeType        string                              `position:"Query" name:"InternetChargeType"`
	VpcId                     string                              `position:"Query" name:"VpcId"`
	Name                      string                              `position:"Query" name:"Name"`
	PrivateLinkEnabled        requests.Boolean                    `position:"Query" name:"PrivateLinkEnabled"`
	EipBindMode               string                              `position:"Query" name:"EipBindMode"`
	PricingCycle              string                              `position:"Query" name:"PricingCycle"`
	AccessMode                CreateNatGatewayAccessMode          `position:"Query" name:"AccessMode"  type:"Struct"`
}

// CreateNatGatewayTag is a repeated param struct in CreateNatGatewayRequest
type CreateNatGatewayTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateNatGatewayBandwidthPackage is a repeated param struct in CreateNatGatewayRequest
type CreateNatGatewayBandwidthPackage struct {
	Bandwidth          string `name:"Bandwidth"`
	Zone               string `name:"Zone"`
	InternetChargeType string `name:"InternetChargeType"`
	ISP                string `name:"ISP"`
	IpCount            string `name:"IpCount"`
}

// CreateNatGatewayAccessMode is a repeated param struct in CreateNatGatewayRequest
type CreateNatGatewayAccessMode struct {
	ModeValue  string `name:"ModeValue"`
	TunnelType string `name:"TunnelType"`
}

// CreateNatGatewayResponse is the response struct for api CreateNatGateway
type CreateNatGatewayResponse struct {
	*responses.BaseResponse
	NatGatewayId        string                                `json:"NatGatewayId" xml:"NatGatewayId"`
	RequestId           string                                `json:"RequestId" xml:"RequestId"`
	ForwardTableIds     ForwardTableIdsInCreateNatGateway     `json:"ForwardTableIds" xml:"ForwardTableIds"`
	SnatTableIds        SnatTableIdsInCreateNatGateway        `json:"SnatTableIds" xml:"SnatTableIds"`
	BandwidthPackageIds BandwidthPackageIdsInCreateNatGateway `json:"BandwidthPackageIds" xml:"BandwidthPackageIds"`
	FullNatTableIds     FullNatTableIdsInCreateNatGateway     `json:"FullNatTableIds" xml:"FullNatTableIds"`
}

// CreateCreateNatGatewayRequest creates a request to invoke CreateNatGateway API
func CreateCreateNatGatewayRequest() (request *CreateNatGatewayRequest) {
	request = &CreateNatGatewayRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CreateNatGateway", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateNatGatewayResponse creates a response to parse from CreateNatGateway response
func CreateCreateNatGatewayResponse() (response *CreateNatGatewayResponse) {
	response = &CreateNatGatewayResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
