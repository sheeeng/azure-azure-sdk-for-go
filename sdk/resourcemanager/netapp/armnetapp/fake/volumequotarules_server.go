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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v6"
	"net/http"
	"net/url"
	"regexp"
)

// VolumeQuotaRulesServer is a fake server for instances of the armnetapp.VolumeQuotaRulesClient type.
type VolumeQuotaRulesServer struct {
	// BeginCreate is the fake for method VolumeQuotaRulesClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, volumeQuotaRuleName string, body armnetapp.VolumeQuotaRule, options *armnetapp.VolumeQuotaRulesClientBeginCreateOptions) (resp azfake.PollerResponder[armnetapp.VolumeQuotaRulesClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method VolumeQuotaRulesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, volumeQuotaRuleName string, options *armnetapp.VolumeQuotaRulesClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetapp.VolumeQuotaRulesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method VolumeQuotaRulesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, volumeQuotaRuleName string, options *armnetapp.VolumeQuotaRulesClientGetOptions) (resp azfake.Responder[armnetapp.VolumeQuotaRulesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByVolumePager is the fake for method VolumeQuotaRulesClient.NewListByVolumePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByVolumePager func(resourceGroupName string, accountName string, poolName string, volumeName string, options *armnetapp.VolumeQuotaRulesClientListByVolumeOptions) (resp azfake.PagerResponder[armnetapp.VolumeQuotaRulesClientListByVolumeResponse])

	// BeginUpdate is the fake for method VolumeQuotaRulesClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, volumeQuotaRuleName string, body armnetapp.VolumeQuotaRulePatch, options *armnetapp.VolumeQuotaRulesClientBeginUpdateOptions) (resp azfake.PollerResponder[armnetapp.VolumeQuotaRulesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewVolumeQuotaRulesServerTransport creates a new instance of VolumeQuotaRulesServerTransport with the provided implementation.
// The returned VolumeQuotaRulesServerTransport instance is connected to an instance of armnetapp.VolumeQuotaRulesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVolumeQuotaRulesServerTransport(srv *VolumeQuotaRulesServer) *VolumeQuotaRulesServerTransport {
	return &VolumeQuotaRulesServerTransport{
		srv:                  srv,
		beginCreate:          newTracker[azfake.PollerResponder[armnetapp.VolumeQuotaRulesClientCreateResponse]](),
		beginDelete:          newTracker[azfake.PollerResponder[armnetapp.VolumeQuotaRulesClientDeleteResponse]](),
		newListByVolumePager: newTracker[azfake.PagerResponder[armnetapp.VolumeQuotaRulesClientListByVolumeResponse]](),
		beginUpdate:          newTracker[azfake.PollerResponder[armnetapp.VolumeQuotaRulesClientUpdateResponse]](),
	}
}

// VolumeQuotaRulesServerTransport connects instances of armnetapp.VolumeQuotaRulesClient to instances of VolumeQuotaRulesServer.
// Don't use this type directly, use NewVolumeQuotaRulesServerTransport instead.
type VolumeQuotaRulesServerTransport struct {
	srv                  *VolumeQuotaRulesServer
	beginCreate          *tracker[azfake.PollerResponder[armnetapp.VolumeQuotaRulesClientCreateResponse]]
	beginDelete          *tracker[azfake.PollerResponder[armnetapp.VolumeQuotaRulesClientDeleteResponse]]
	newListByVolumePager *tracker[azfake.PagerResponder[armnetapp.VolumeQuotaRulesClientListByVolumeResponse]]
	beginUpdate          *tracker[azfake.PollerResponder[armnetapp.VolumeQuotaRulesClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for VolumeQuotaRulesServerTransport.
func (v *VolumeQuotaRulesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VolumeQuotaRulesClient.BeginCreate":
		resp, err = v.dispatchBeginCreate(req)
	case "VolumeQuotaRulesClient.BeginDelete":
		resp, err = v.dispatchBeginDelete(req)
	case "VolumeQuotaRulesClient.Get":
		resp, err = v.dispatchGet(req)
	case "VolumeQuotaRulesClient.NewListByVolumePager":
		resp, err = v.dispatchNewListByVolumePager(req)
	case "VolumeQuotaRulesClient.BeginUpdate":
		resp, err = v.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VolumeQuotaRulesServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := v.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetApp/netAppAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/capacityPools/(?P<poolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/volumes/(?P<volumeName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/volumeQuotaRules/(?P<volumeQuotaRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetapp.VolumeQuotaRule](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		poolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("poolName")])
		if err != nil {
			return nil, err
		}
		volumeNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("volumeName")])
		if err != nil {
			return nil, err
		}
		volumeQuotaRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("volumeQuotaRuleName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginCreate(req.Context(), resourceGroupNameParam, accountNameParam, poolNameParam, volumeNameParam, volumeQuotaRuleNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		v.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		v.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		v.beginCreate.remove(req)
	}

	return resp, nil
}

func (v *VolumeQuotaRulesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if v.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := v.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetApp/netAppAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/capacityPools/(?P<poolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/volumes/(?P<volumeName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/volumeQuotaRules/(?P<volumeQuotaRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		poolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("poolName")])
		if err != nil {
			return nil, err
		}
		volumeNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("volumeName")])
		if err != nil {
			return nil, err
		}
		volumeQuotaRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("volumeQuotaRuleName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginDelete(req.Context(), resourceGroupNameParam, accountNameParam, poolNameParam, volumeNameParam, volumeQuotaRuleNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		v.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		v.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		v.beginDelete.remove(req)
	}

	return resp, nil
}

func (v *VolumeQuotaRulesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetApp/netAppAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/capacityPools/(?P<poolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/volumes/(?P<volumeName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/volumeQuotaRules/(?P<volumeQuotaRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	poolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("poolName")])
	if err != nil {
		return nil, err
	}
	volumeNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("volumeName")])
	if err != nil {
		return nil, err
	}
	volumeQuotaRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("volumeQuotaRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Get(req.Context(), resourceGroupNameParam, accountNameParam, poolNameParam, volumeNameParam, volumeQuotaRuleNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VolumeQuotaRule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VolumeQuotaRulesServerTransport) dispatchNewListByVolumePager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListByVolumePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByVolumePager not implemented")}
	}
	newListByVolumePager := v.newListByVolumePager.get(req)
	if newListByVolumePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetApp/netAppAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/capacityPools/(?P<poolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/volumes/(?P<volumeName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/volumeQuotaRules`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		poolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("poolName")])
		if err != nil {
			return nil, err
		}
		volumeNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("volumeName")])
		if err != nil {
			return nil, err
		}
		resp := v.srv.NewListByVolumePager(resourceGroupNameParam, accountNameParam, poolNameParam, volumeNameParam, nil)
		newListByVolumePager = &resp
		v.newListByVolumePager.add(req, newListByVolumePager)
	}
	resp, err := server.PagerResponderNext(newListByVolumePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		v.newListByVolumePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByVolumePager) {
		v.newListByVolumePager.remove(req)
	}
	return resp, nil
}

func (v *VolumeQuotaRulesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := v.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetApp/netAppAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/capacityPools/(?P<poolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/volumes/(?P<volumeName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/volumeQuotaRules/(?P<volumeQuotaRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetapp.VolumeQuotaRulePatch](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		poolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("poolName")])
		if err != nil {
			return nil, err
		}
		volumeNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("volumeName")])
		if err != nil {
			return nil, err
		}
		volumeQuotaRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("volumeQuotaRuleName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginUpdate(req.Context(), resourceGroupNameParam, accountNameParam, poolNameParam, volumeNameParam, volumeQuotaRuleNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		v.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		v.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		v.beginUpdate.remove(req)
	}

	return resp, nil
}