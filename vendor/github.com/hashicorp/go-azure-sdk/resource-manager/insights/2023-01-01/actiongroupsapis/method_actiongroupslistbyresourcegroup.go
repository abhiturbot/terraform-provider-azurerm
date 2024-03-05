package actiongroupsapis

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActionGroupsListByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ActionGroupResource
}

type ActionGroupsListByResourceGroupCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ActionGroupResource
}

// ActionGroupsListByResourceGroup ...
func (c ActionGroupsAPIsClient) ActionGroupsListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (result ActionGroupsListByResourceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/providers/Microsoft.Insights/actionGroups", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]ActionGroupResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ActionGroupsListByResourceGroupComplete retrieves all the results into a single object
func (c ActionGroupsAPIsClient) ActionGroupsListByResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId) (ActionGroupsListByResourceGroupCompleteResult, error) {
	return c.ActionGroupsListByResourceGroupCompleteMatchingPredicate(ctx, id, ActionGroupResourceOperationPredicate{})
}

// ActionGroupsListByResourceGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ActionGroupsAPIsClient) ActionGroupsListByResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, predicate ActionGroupResourceOperationPredicate) (result ActionGroupsListByResourceGroupCompleteResult, err error) {
	items := make([]ActionGroupResource, 0)

	resp, err := c.ActionGroupsListByResourceGroup(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ActionGroupsListByResourceGroupCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
