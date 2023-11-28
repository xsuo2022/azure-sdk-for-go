//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SecureScoreControlsClient contains the methods for the SecureScoreControls group.
// Don't use this type directly, use NewSecureScoreControlsClient() instead.
type SecureScoreControlsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSecureScoreControlsClient creates a new instance of SecureScoreControlsClient with the specified values.
//   - subscriptionID - Azure subscription ID
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSecureScoreControlsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SecureScoreControlsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SecureScoreControlsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Get all security controls within a scope
//
// Generated from API version 2020-01-01
//   - options - SecureScoreControlsClientListOptions contains the optional parameters for the SecureScoreControlsClient.NewListPager
//     method.
func (client *SecureScoreControlsClient) NewListPager(options *SecureScoreControlsClientListOptions) *runtime.Pager[SecureScoreControlsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SecureScoreControlsClientListResponse]{
		More: func(page SecureScoreControlsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SecureScoreControlsClientListResponse) (SecureScoreControlsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SecureScoreControlsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return SecureScoreControlsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *SecureScoreControlsClient) listCreateRequest(ctx context.Context, options *SecureScoreControlsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/secureScoreControls"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SecureScoreControlsClient) listHandleResponse(resp *http.Response) (SecureScoreControlsClientListResponse, error) {
	result := SecureScoreControlsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecureScoreControlList); err != nil {
		return SecureScoreControlsClientListResponse{}, err
	}
	return result, nil
}

// NewListBySecureScorePager - Get all security controls for a specific initiative within a scope
//
// Generated from API version 2020-01-01
//   - secureScoreName - The initiative name. For the ASC Default initiative, use 'ascScore' as in the sample request below.
//   - options - SecureScoreControlsClientListBySecureScoreOptions contains the optional parameters for the SecureScoreControlsClient.NewListBySecureScorePager
//     method.
func (client *SecureScoreControlsClient) NewListBySecureScorePager(secureScoreName string, options *SecureScoreControlsClientListBySecureScoreOptions) *runtime.Pager[SecureScoreControlsClientListBySecureScoreResponse] {
	return runtime.NewPager(runtime.PagingHandler[SecureScoreControlsClientListBySecureScoreResponse]{
		More: func(page SecureScoreControlsClientListBySecureScoreResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SecureScoreControlsClientListBySecureScoreResponse) (SecureScoreControlsClientListBySecureScoreResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SecureScoreControlsClient.NewListBySecureScorePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySecureScoreCreateRequest(ctx, secureScoreName, options)
			}, nil)
			if err != nil {
				return SecureScoreControlsClientListBySecureScoreResponse{}, err
			}
			return client.listBySecureScoreHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySecureScoreCreateRequest creates the ListBySecureScore request.
func (client *SecureScoreControlsClient) listBySecureScoreCreateRequest(ctx context.Context, secureScoreName string, options *SecureScoreControlsClientListBySecureScoreOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/secureScores/{secureScoreName}/secureScoreControls"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if secureScoreName == "" {
		return nil, errors.New("parameter secureScoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{secureScoreName}", url.PathEscape(secureScoreName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySecureScoreHandleResponse handles the ListBySecureScore response.
func (client *SecureScoreControlsClient) listBySecureScoreHandleResponse(resp *http.Response) (SecureScoreControlsClientListBySecureScoreResponse, error) {
	result := SecureScoreControlsClientListBySecureScoreResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecureScoreControlList); err != nil {
		return SecureScoreControlsClientListBySecureScoreResponse{}, err
	}
	return result, nil
}
