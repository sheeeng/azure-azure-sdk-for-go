//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sphere/armsphere"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// DeviceGroupsServer is a fake server for instances of the armsphere.DeviceGroupsClient type.
type DeviceGroupsServer struct {
	// BeginClaimDevices is the fake for method DeviceGroupsClient.BeginClaimDevices
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginClaimDevices func(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, claimDevicesRequest armsphere.ClaimDevicesRequest, options *armsphere.DeviceGroupsClientBeginClaimDevicesOptions) (resp azfake.PollerResponder[armsphere.DeviceGroupsClientClaimDevicesResponse], errResp azfake.ErrorResponder)

	// CountDevices is the fake for method DeviceGroupsClient.CountDevices
	// HTTP status codes to indicate success: http.StatusOK
	CountDevices func(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, options *armsphere.DeviceGroupsClientCountDevicesOptions) (resp azfake.Responder[armsphere.DeviceGroupsClientCountDevicesResponse], errResp azfake.ErrorResponder)

	// BeginCreateOrUpdate is the fake for method DeviceGroupsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, resource armsphere.DeviceGroup, options *armsphere.DeviceGroupsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armsphere.DeviceGroupsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method DeviceGroupsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, options *armsphere.DeviceGroupsClientBeginDeleteOptions) (resp azfake.PollerResponder[armsphere.DeviceGroupsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method DeviceGroupsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, options *armsphere.DeviceGroupsClientGetOptions) (resp azfake.Responder[armsphere.DeviceGroupsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByProductPager is the fake for method DeviceGroupsClient.NewListByProductPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByProductPager func(resourceGroupName string, catalogName string, productName string, options *armsphere.DeviceGroupsClientListByProductOptions) (resp azfake.PagerResponder[armsphere.DeviceGroupsClientListByProductResponse])

	// BeginUpdate is the fake for method DeviceGroupsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, properties armsphere.DeviceGroupUpdate, options *armsphere.DeviceGroupsClientBeginUpdateOptions) (resp azfake.PollerResponder[armsphere.DeviceGroupsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewDeviceGroupsServerTransport creates a new instance of DeviceGroupsServerTransport with the provided implementation.
// The returned DeviceGroupsServerTransport instance is connected to an instance of armsphere.DeviceGroupsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDeviceGroupsServerTransport(srv *DeviceGroupsServer) *DeviceGroupsServerTransport {
	return &DeviceGroupsServerTransport{
		srv:                   srv,
		beginClaimDevices:     newTracker[azfake.PollerResponder[armsphere.DeviceGroupsClientClaimDevicesResponse]](),
		beginCreateOrUpdate:   newTracker[azfake.PollerResponder[armsphere.DeviceGroupsClientCreateOrUpdateResponse]](),
		beginDelete:           newTracker[azfake.PollerResponder[armsphere.DeviceGroupsClientDeleteResponse]](),
		newListByProductPager: newTracker[azfake.PagerResponder[armsphere.DeviceGroupsClientListByProductResponse]](),
		beginUpdate:           newTracker[azfake.PollerResponder[armsphere.DeviceGroupsClientUpdateResponse]](),
	}
}

// DeviceGroupsServerTransport connects instances of armsphere.DeviceGroupsClient to instances of DeviceGroupsServer.
// Don't use this type directly, use NewDeviceGroupsServerTransport instead.
type DeviceGroupsServerTransport struct {
	srv                   *DeviceGroupsServer
	beginClaimDevices     *tracker[azfake.PollerResponder[armsphere.DeviceGroupsClientClaimDevicesResponse]]
	beginCreateOrUpdate   *tracker[azfake.PollerResponder[armsphere.DeviceGroupsClientCreateOrUpdateResponse]]
	beginDelete           *tracker[azfake.PollerResponder[armsphere.DeviceGroupsClientDeleteResponse]]
	newListByProductPager *tracker[azfake.PagerResponder[armsphere.DeviceGroupsClientListByProductResponse]]
	beginUpdate           *tracker[azfake.PollerResponder[armsphere.DeviceGroupsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for DeviceGroupsServerTransport.
func (d *DeviceGroupsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DeviceGroupsClient.BeginClaimDevices":
		resp, err = d.dispatchBeginClaimDevices(req)
	case "DeviceGroupsClient.CountDevices":
		resp, err = d.dispatchCountDevices(req)
	case "DeviceGroupsClient.BeginCreateOrUpdate":
		resp, err = d.dispatchBeginCreateOrUpdate(req)
	case "DeviceGroupsClient.BeginDelete":
		resp, err = d.dispatchBeginDelete(req)
	case "DeviceGroupsClient.Get":
		resp, err = d.dispatchGet(req)
	case "DeviceGroupsClient.NewListByProductPager":
		resp, err = d.dispatchNewListByProductPager(req)
	case "DeviceGroupsClient.BeginUpdate":
		resp, err = d.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DeviceGroupsServerTransport) dispatchBeginClaimDevices(req *http.Request) (*http.Response, error) {
	if d.srv.BeginClaimDevices == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginClaimDevices not implemented")}
	}
	beginClaimDevices := d.beginClaimDevices.get(req)
	if beginClaimDevices == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureSphere/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/products/(?P<productName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deviceGroups/(?P<deviceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/claimDevices`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsphere.ClaimDevicesRequest](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
		if err != nil {
			return nil, err
		}
		productNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("productName")])
		if err != nil {
			return nil, err
		}
		deviceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginClaimDevices(req.Context(), resourceGroupNameParam, catalogNameParam, productNameParam, deviceGroupNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginClaimDevices = &respr
		d.beginClaimDevices.add(req, beginClaimDevices)
	}

	resp, err := server.PollerResponderNext(beginClaimDevices, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		d.beginClaimDevices.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginClaimDevices) {
		d.beginClaimDevices.remove(req)
	}

	return resp, nil
}

func (d *DeviceGroupsServerTransport) dispatchCountDevices(req *http.Request) (*http.Response, error) {
	if d.srv.CountDevices == nil {
		return nil, &nonRetriableError{errors.New("fake for method CountDevices not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureSphere/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/products/(?P<productName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deviceGroups/(?P<deviceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/countDevices`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
	if err != nil {
		return nil, err
	}
	productNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("productName")])
	if err != nil {
		return nil, err
	}
	deviceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.CountDevices(req.Context(), resourceGroupNameParam, catalogNameParam, productNameParam, deviceGroupNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CountDeviceResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DeviceGroupsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := d.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureSphere/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/products/(?P<productName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deviceGroups/(?P<deviceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsphere.DeviceGroup](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
		if err != nil {
			return nil, err
		}
		productNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("productName")])
		if err != nil {
			return nil, err
		}
		deviceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, catalogNameParam, productNameParam, deviceGroupNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		d.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		d.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		d.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (d *DeviceGroupsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if d.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := d.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureSphere/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/products/(?P<productName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deviceGroups/(?P<deviceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
		if err != nil {
			return nil, err
		}
		productNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("productName")])
		if err != nil {
			return nil, err
		}
		deviceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginDelete(req.Context(), resourceGroupNameParam, catalogNameParam, productNameParam, deviceGroupNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		d.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		d.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		d.beginDelete.remove(req)
	}

	return resp, nil
}

func (d *DeviceGroupsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureSphere/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/products/(?P<productName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deviceGroups/(?P<deviceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
	if err != nil {
		return nil, err
	}
	productNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("productName")])
	if err != nil {
		return nil, err
	}
	deviceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Get(req.Context(), resourceGroupNameParam, catalogNameParam, productNameParam, deviceGroupNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DeviceGroup, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DeviceGroupsServerTransport) dispatchNewListByProductPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListByProductPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByProductPager not implemented")}
	}
	newListByProductPager := d.newListByProductPager.get(req)
	if newListByProductPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureSphere/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/products/(?P<productName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deviceGroups`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		skipUnescaped, err := url.QueryUnescape(qp.Get("$skip"))
		if err != nil {
			return nil, err
		}
		skipParam, err := parseOptional(skipUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		maxpagesizeUnescaped, err := url.QueryUnescape(qp.Get("$maxpagesize"))
		if err != nil {
			return nil, err
		}
		maxpagesizeParam, err := parseOptional(maxpagesizeUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
		if err != nil {
			return nil, err
		}
		productNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("productName")])
		if err != nil {
			return nil, err
		}
		var options *armsphere.DeviceGroupsClientListByProductOptions
		if filterParam != nil || topParam != nil || skipParam != nil || maxpagesizeParam != nil {
			options = &armsphere.DeviceGroupsClientListByProductOptions{
				Filter:      filterParam,
				Top:         topParam,
				Skip:        skipParam,
				Maxpagesize: maxpagesizeParam,
			}
		}
		resp := d.srv.NewListByProductPager(resourceGroupNameParam, catalogNameParam, productNameParam, options)
		newListByProductPager = &resp
		d.newListByProductPager.add(req, newListByProductPager)
		server.PagerResponderInjectNextLinks(newListByProductPager, req, func(page *armsphere.DeviceGroupsClientListByProductResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByProductPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListByProductPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByProductPager) {
		d.newListByProductPager.remove(req)
	}
	return resp, nil
}

func (d *DeviceGroupsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := d.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureSphere/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/products/(?P<productName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deviceGroups/(?P<deviceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsphere.DeviceGroupUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
		if err != nil {
			return nil, err
		}
		productNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("productName")])
		if err != nil {
			return nil, err
		}
		deviceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginUpdate(req.Context(), resourceGroupNameParam, catalogNameParam, productNameParam, deviceGroupNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		d.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		d.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		d.beginUpdate.remove(req)
	}

	return resp, nil
}